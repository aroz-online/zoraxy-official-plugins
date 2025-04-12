package main

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	plugin "zoraxy.aroz.org/tools/dirupdate/mod/zoraxy_plugin"
)

// / getPlatformBinaryNameFromFolderName returns the binary name for the given folder name.
func getPlatformBinaryNameFromFolderName(folderName string) string {
	folderName = filepath.Base(folderName)
	binaryName := folderName + "_" + runtime.GOOS + "_" + runtime.GOARCH
	if runtime.GOOS == "windows" {
		binaryName += ".exe"
	}
	return binaryName
}

// GetPluginEntryPoint returns the plugin entry point
func getPluginSpec(entryPoint string) (*plugin.IntroSpect, error) {
	pluginSpec := plugin.IntroSpect{}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, entryPoint, "-introspect")
	output, err := cmd.Output()
	if ctx.Err() == context.DeadlineExceeded {
		return nil, fmt.Errorf("plugin introspect timed out")
	}
	if err != nil {
		return nil, err
	}

	// Assuming the output is JSON and needs to be unmarshaled into pluginSpec
	err = json.Unmarshal(output, &pluginSpec)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal plugin spec: %v", err)
	}

	return &pluginSpec, nil
}

// generateDownloadURLs generates the download URLs for the given folder path
func generateDownloadURLs(folderpath string) (map[string]string, error) {
	downloadURLs := make(map[string]string)
	entries, err := os.ReadDir(folderpath)
	if err != nil {
		return nil, fmt.Errorf("failed to read directory: %v", err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		//filePath := filepath.Join(folderpath, entry.Name())
		fileURL := fmt.Sprintf(DOWNLOAD_MAIN_URL + entry.Name()) // Replace with actual URL generation logic
		fmt.Println("Download URL for file", entry.Name(), ":", fileURL)
		// Add the file URL to the map with the file name as the key
		key := filepath.Base(entry.Name())
		if runtime.GOOS == "windows" {
			key = strings.TrimSuffix(key, ".exe")
		}

		//Also trim the plugin name from the file name
		key = strings.TrimPrefix(key, filepath.Base(folderpath)+"_")
		downloadURLs[key] = fileURL
	}
	return downloadURLs, nil
}

// GeneratePluginDirInfo generates the plugin directory info for the given folder name
func generateChecksumForDistFolder(folderpath string) (Checksums, error) {
	checksums := Checksums{}

	entries, err := os.ReadDir(folderpath)
	if err != nil {
		return checksums, fmt.Errorf("failed to read directory: %v", err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		filePath := filepath.Join(folderpath, entry.Name())

		// Calculate the checksum of the file
		file, err := os.Open(filePath)
		if err != nil {
			return checksums, fmt.Errorf("failed to open file: %v", err)
		}
		defer file.Close()

		hash := sha256.New()
		if _, err := io.Copy(hash, file); err != nil {
			return checksums, fmt.Errorf("failed to calculate checksum: %v", err)
		}
		checksum := fmt.Sprintf("%x", hash.Sum(nil))
		fmt.Println("Checksum for file", filePath, ":", checksum)

		// Determine the correct field in the Checksums struct
		baseName := filepath.Base(filePath)
		pluginName := filepath.Base(folderpath)
		switch {
		case baseName == pluginName+"_linux_amd64":
			checksums.LinuxAmd64 = checksum
		case baseName == pluginName+"_linux_386":
			checksums.Linux386 = checksum
		case baseName == pluginName+"_linux_arm":
			checksums.LinuxArm = checksum
		case baseName == pluginName+"_linux_arm64":
			checksums.LinuxArm64 = checksum
		case baseName == pluginName+"_linux_mipsle":
			checksums.LinuxMipsle = checksum
		case baseName == pluginName+"_linux_riscv64":
			checksums.LinuxRiscv64 = checksum
		case baseName == pluginName+"_windows_amd64.exe":
			checksums.WindowsAmd64 = checksum
		}
	}

	return checksums, nil
}
