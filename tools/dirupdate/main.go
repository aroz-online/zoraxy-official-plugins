package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"

	plugin "zoraxy.aroz.org/tools/dirupdate/mod/zoraxy_plugin"
)

type Checksums struct {
	LinuxAmd64   string `json:"linux_amd64"`
	Linux386     string `json:"linux_386"`
	LinuxArm     string `json:"linux_arm"`
	LinuxArm64   string `json:"linux_arm64"`
	LinuxMipsle  string `json:"linux_mipsle"`
	LinuxRiscv64 string `json:"linux_riscv64"`
	WindowsAmd64 string `json:"windows_amd64"`
}

type PluginDirInfo struct {
	IconPath         string
	PluginIntroSpect plugin.IntroSpect //Plugin introspect information
	ChecksumsSHA256  Checksums         //Checksums for the plugin binary
	DownloadURLs     map[string]string //Download URLs for different platforms
}

const (
	// The URL to download the directory information
	DIR_INFO_INDEX_URL = "https://raw.githubusercontent.com/aroz-online/zoraxy-official-plugins/main/"
	DIR_INFO_ICON_URL  = DIR_INFO_INDEX_URL + "directories/icon/"

	// The URL to download the latest version of the plugin
	DOWNLOAD_MAIN_URL = "https://github.com/aroz-online/zoraxy-official-plugins/releases/latest/download/"
)

func main() {
	// Check if ./src exists
	if _, err := os.Stat("./src"); os.IsNotExist(err) {
		// Change directory to ../../
		err := os.Chdir("../../")
		if err != nil {
			fmt.Println("Error changing directory:", err)
			return
		}
	}

	//Create a icon folder in ./directories
	if err := os.MkdirAll("./directories/icon", os.ModePerm); err != nil {
		fmt.Println("Error creating icon directory:", err)
		return
	}

	// Read the contents of the current directory
	entries, err := os.ReadDir("./src")
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	// Iterate through each entry in the current directory
	newDirectories := []*PluginDirInfo{}
	for _, entry := range entries {
		if entry.IsDir() {
			fmt.Println("Generating plugin directory for :", filepath.Join("./src", entry.Name()))
			//Check if the plugin exists in dist folder.
			distPath := filepath.Join("./dist", entry.Name())
			if _, err := os.Stat(distPath); os.IsNotExist(err) {
				fmt.Println("Plugin does not exist in dist folder:", distPath)
				continue
			}

			thisPluginDirInfo := &PluginDirInfo{
				ChecksumsSHA256:  Checksums{},
				PluginIntroSpect: plugin.IntroSpect{},
				DownloadURLs:     make(map[string]string),
			}

			//Check if the current platform exists in the dist folder.
			platformBinaryName := getPlatformBinaryNameFromFolderName(entry.Name())
			platformBinaryPath := filepath.Join(distPath, platformBinaryName)
			if _, err := os.Stat(platformBinaryPath); os.IsNotExist(err) {
				fmt.Println("Platform binary does not exist:", platformBinaryPath)
				continue
			}

			//Get plugin introspect
			pluginSpec, err := getPluginSpec(platformBinaryPath)
			if err != nil {
				fmt.Println("Error getting plugin spec:", err)
				continue
			}
			thisPluginDirInfo.PluginIntroSpect = *pluginSpec

			//Generate checksum for the dist folder
			thisChecksumList, err := generateChecksumForDistFolder(distPath)
			if err != nil {
				fmt.Println("Error generating checksum:", err)
				continue
			}
			thisPluginDirInfo.ChecksumsSHA256 = thisChecksumList

			downloadURLs, err := generateDownloadURLs(distPath)
			if err != nil {
				fmt.Println("Error generating download URLs:", err)
				continue
			}
			thisPluginDirInfo.DownloadURLs = downloadURLs

			//Check if icon.png exists in this folder. If it does, copy it to ./directories/icon
			iconPath := filepath.Join("./src", entry.Name(), "icon.png")
			if _, err := os.Stat(iconPath); err == nil {
				destPath := filepath.Join("./directories/icon", entry.Name()+".png")
				srcFile, err := os.Open(iconPath)
				if err != nil {
					fmt.Println("Error opening source file:", err)
					continue
				}
				defer srcFile.Close()

				destFile, err := os.Create(destPath)
				if err != nil {
					fmt.Println("Error creating destination file:", err)
					continue
				}
				defer destFile.Close()

				_, err = io.Copy(destFile, srcFile)
				if err != nil {
					fmt.Println("Error copying file:", err)
				}

				thisPluginDirInfo.IconPath = DIR_INFO_ICON_URL + entry.Name() + ".png"
			}

			newDirectories = append(newDirectories, thisPluginDirInfo)
		}
	}
	if err != nil {
		fmt.Println("Error walking the directory:", err)
	}

	//Write the new directories to ./directories/index.json
	js, err := json.MarshalIndent(newDirectories, "", " ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}
	if err := os.WriteFile("./directories/index.json", js, os.ModePerm); err != nil {
		fmt.Println("Error writing index.json:", err)
		return
	}

}
