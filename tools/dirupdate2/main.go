package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type AppConfig struct {
	RepoURL          string `json:"repo_url"`
	Author           string `json:"author"`
	Contact          string `json:"contact"`
	MinZoraxyVersion string `json:"min_zoraxy_version"`

	// Optionally override the default artifact names used to construct download URLs
	ArtifactNames map[string]string `json:"artifact_names,omitempty"`
}

type PluginIntrospect struct {
	ID                    string      `json:"id"`
	Name                  string      `json:"name"`
	Author                string      `json:"author"`
	AuthorContact         string      `json:"author_contact"`
	Description           string      `json:"description"`
	URL                   string      `json:"url"`
	Type                  int         `json:"type"`
	VersionMajor          int         `json:"version_major"`
	VersionMinor          int         `json:"version_minor"`
	VersionPatch          int         `json:"version_patch"`
	Preview      		  bool        `json:"preview,omitempty"`
	StaticCapturePaths    interface{} `json:"static_capture_paths"`
	StaticCaptureIngress  string      `json:"static_capture_ingress"`
	DynamicCaptureSniff   string      `json:"dynamic_capture_sniff"`
	DynamicCaptureIngress string      `json:"dynamic_capture_ingress"`
	UIPath                string      `json:"ui_path"`
	SubscriptionPath      string      `json:"subscription_path"`
	SubscriptionsEvents   interface{} `json:"subscriptions_events"`
	PermittedAPIEndpoints interface{} `json:"permitted_api_endpoints"`
}

type IndexEntry struct {
	IconPath         string            `json:"IconPath"`
	PluginIntroSpect *PluginIntrospect `json:"PluginIntroSpect"`
	DownloadURLs     map[string]string `json:"DownloadURLs"`
}

func main() {
	fmt.Println(`
 ____  ____  ___  ___   _  ____  __   ___  __   __  ____________  ______
 /_  / / __ \/ _ \/ _ | | |/_/\ \/ /  / _ \/ /  / / / / ___/  _/ |/ / __/
  / /_/ /_/ / , _/ __ |_>  <   \  /  / ___/ /__/ /_/ / (_ // //    /\ \  
 /___/\____/_/|_/_/ |_/_/|_|   /_/  /_/  /____/\____/\___/___/_/|_/___/  
                                                                         
 ----------------------------
 This script updates the index2.json file based on the apps' .introspect and .releaseurl files.
 Make sure you have internet connectivity!

 `)
	appsDir := "../../apps"
	files, err := os.ReadDir(appsDir)
	if err != nil {
		fmt.Println("Error reading apps directory:", err)
		return
	}

	index := []IndexEntry{}

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".json" {
			fmt.Println("--------------------------------------------------")
			fmt.Println("Processing", file.Name())
			config, err := readAppConfig(filepath.Join(appsDir, file.Name()))
			if err != nil {
				fmt.Println("Error reading config:", err)
				continue
			}

			entry, err := processApp(config)
			if err != nil {
				fmt.Println("Error processing app:", err)
				// We skip this app if critical info is missing
				continue
			}
			index = append(index, *entry)
		}
	}
	fmt.Println("--------------------------------------------------")
	outputFile := "../../directories/index2.json"
	err = writeIndex(outputFile, index)
	if err != nil {
		fmt.Println("Error writing index:", err)
	} else {
		fmt.Println("Successfully updated index2.json")
	}
}

func readAppConfig(path string) (*AppConfig, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var config AppConfig
	err = json.Unmarshal(data, &config)
	return &config, err
}

func processApp(config *AppConfig) (*IndexEntry, error) {
	// Convert github URL to raw URL base
	rawBase, err := getRawBaseURL(config.RepoURL)
	if err != nil {
		return nil, err
	}

	// 1. Fetch .introspect (Critical)
	introspectBytes, err := fetchURL(rawBase + ".introspect")
	if err != nil {
		return nil, fmt.Errorf("failed to fetch .introspect: %v", err)
	}
	var introspect PluginIntrospect
	err = json.Unmarshal(introspectBytes, &introspect)
	if err != nil {
		fmt.Println("Warning: .introspect not found or failed to fetch")
		//return nil, fmt.Errorf("failed to parse .introspect: %v", err)
	}

	// 2. Fetch .releaseurl (Critical)
	var releaseBase string
	releaseURLBytes, err := fetchURL(rawBase + ".releaseurl")
	if err == nil {
		releaseBase = strings.TrimSpace(string(releaseURLBytes))
	} else {
		fmt.Println("Warning: .releaseurl not found or failed to fetch")
	}

	// 3. Check icon.png
	iconURL := rawBase + "icon.png"
	if !checkURLExists(iconURL) {
		iconURL = "" // No icon available
	}

	// Construct DownloadURLs
	downloadURLs := make(map[string]string)
	if releaseBase != "" && len(config.ArtifactNames) > 0 {
		// Use custom artifact names if provided
		for arch, filename := range config.ArtifactNames {
			if !strings.HasSuffix(releaseBase, "/") {
				releaseBase += "/"
			}
			downloadURLs[arch] = releaseBase + filename
		}
	} else if releaseBase != "" {
		archs := []string{"linux_amd64", "linux_386", "linux_arm", "linux_arm64", "linux_mipsle", "linux_riscv64", "windows_amd64"}

		// Ensure releaseBase ends with / if not empty
		if !strings.HasSuffix(releaseBase, "/") {
			releaseBase += "/"
		}

		for _, arch := range archs {
			filename := introspect.ID + "_" + arch
			if strings.Contains(introspect.ID, ".") {
				//The introspec ID is a inverse domain name like org.aroz.zoraxy.xxxx
				//We will trim off the domain and keeping only the application name
				parts := strings.Split(introspect.ID, ".")
				filename = parts[len(parts)-1] + "_" + arch
			}
			if strings.HasPrefix(arch, "windows") {
				filename += ".exe"
			}
			downloadURLs[arch] = releaseBase + filename
		}
	}

	return &IndexEntry{
		IconPath:         iconURL,
		PluginIntroSpect: &introspect,
		DownloadURLs:     downloadURLs,
	}, nil
}

func getRawBaseURL(repoURL string) (string, error) {
	repoURL = strings.TrimSuffix(repoURL, "/")
	parts := strings.Split(repoURL, "/")
	if len(parts) < 5 {
		return "", fmt.Errorf("invalid repo url")
	}
	// https://github.com/user/repo
	user := parts[len(parts)-2]
	repo := parts[len(parts)-1]

	// Try main
	mainURL := fmt.Sprintf("https://raw.githubusercontent.com/%s/%s/refs/heads/main/", user, repo)
	fmt.Println("Trying main URL:", mainURL+".introspect")
	if checkURLExists(mainURL + ".introspect") {
		fmt.Println("Found main URL:", mainURL+".introspect")
		return mainURL, nil
	}

	// Try master
	masterURL := fmt.Sprintf("https://raw.githubusercontent.com/%s/%s/refs/heads/master/", user, repo)
	fmt.Println("Trying master URL:", masterURL+".introspect")
	if checkURLExists(masterURL + ".introspect") {
		fmt.Println("Found master URL:", masterURL+".introspect")
		return masterURL, nil
	}

	return "", fmt.Errorf("could not determine raw base URL (tried main and master)")
}

func fetchURL(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("status code %d", resp.StatusCode)
	}

	return io.ReadAll(resp.Body)
}

func checkURLExists(url string) bool {
	resp, err := http.Get(url)
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	return resp.StatusCode == 200
}

func writeIndex(path string, index []IndexEntry) error {
	data, err := json.MarshalIndent(index, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}
