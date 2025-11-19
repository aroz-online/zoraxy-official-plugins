#!/bin/bash

# Prompt the user for the GitHub repository URL
echo "Enter the GitHub repository URL for the subrepo:"
read repo_url

# Extract the repository name from the URL (assuming format https://github.com/user/repo)
repo_name=$(basename "$repo_url" .git)

# Add the submodule to the src folder
git submodule add "$repo_url" "src/$repo_name"

echo "Subrepo added successfully to src/$repo_name"

# Check if the submodule has a makefile
if [ ! -f "src/$repo_name/makefile" ]; then
    echo "The submodule does not have a makefile. Do you want to automatically generate one? (y/n)"
    read answer
    if [ "$answer" = "y" ] || [ "$answer" = "Y" ]; then
        cat > "src/$repo_name/makefile" << EOF
.PHONY: all

all: ${repo_name}_linux_386 ${repo_name}_linux_amd64 ${repo_name}_linux_arm ${repo_name}_linux_arm64 ${repo_name}_linux_mipsle ${repo_name}_linux_riscv64 ${repo_name}_windows_amd64.exe

${repo_name}_linux_386:
	mkdir -p ./build
	CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -o ${repo_name}_linux_386
	mv ${repo_name}_linux_386 ./build/

${repo_name}_linux_amd64:
	mkdir -p ./build
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${repo_name}_linux_amd64
	mv ${repo_name}_linux_amd64 ./build/

${repo_name}_linux_arm:
	mkdir -p ./build
	CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -o ${repo_name}_linux_arm
	mv ${repo_name}_linux_arm ./build/

${repo_name}_linux_arm64:
	mkdir -p ./build
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o ${repo_name}_linux_arm64
	mv ${repo_name}_linux_arm64 ./build/

${repo_name}_linux_mipsle:
	mkdir -p ./build
	CGO_ENABLED=0 GOOS=linux GOARCH=mipsle go build -o ${repo_name}_linux_mipsle
	mv ${repo_name}_linux_mipsle ./build/

${repo_name}_linux_riscv64:
	mkdir -p ./build
	CGO_ENABLED=0 GOOS=linux GOARCH=riscv64 go build -o ${repo_name}_linux_riscv64
	mv ${repo_name}_linux_riscv64 ./build/

${repo_name}_windows_amd64.exe:
	mkdir -p ./build
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ${repo_name}_windows_amd64.exe
	mv ${repo_name}_windows_amd64.exe ./build/

.PHONY: all ${repo_name}_linux_386 ${repo_name}_linux_amd64 ${repo_name}_linux_arm ${repo_name}_linux_arm64 ${repo_name}_linux_mipsle ${repo_name}_linux_riscv64 ${repo_name}_windows_amd64.exe
EOF
        echo "Makefile generated for src/$repo_name"
    fi
fi
