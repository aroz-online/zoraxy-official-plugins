#!/bin/bash

# Define the repository URL and target directory
REPO_URL="https://github.com/tobychui/zoraxy"
TARGET_DIR="./zoraxy"
PLUGIN_SRC_DIR="$TARGET_DIR/src/mod/plugins/zoraxy_plugin"
PLUGIN_DEST_DIR="./mod"

# Check if the target directory already exists and remove it
if [ -d "$TARGET_DIR" ]; then
    echo "Directory $TARGET_DIR already exists. Removing it..."
    rm -rf "$TARGET_DIR"
    echo "Old directory $TARGET_DIR removed."
fi

# Clone the repository into the target directory
echo "Cloning repository $REPO_URL into $TARGET_DIR..."
git clone --branch main --depth 1 "$REPO_URL" "$TARGET_DIR"

# Check if the clone operation was successful
if [ $? -eq 0 ]; then
    echo "Repository successfully cloned into $TARGET_DIR."
else
    echo "Failed to clone repository."
    exit 1
fi

# Check if the old plugin library copy exists in the destination directory
if [ -d "$PLUGIN_DEST_DIR/zoraxy_plugin" ]; then
    echo "Old plugin copy found in $PLUGIN_DEST_DIR. Removing it..."
    rm -rf "$PLUGIN_DEST_DIR/zoraxy_plugin"
    echo "Old plugin copy removed."
fi

# Check if the plugin library directory exists
if [ -d "$PLUGIN_SRC_DIR" ]; then
    echo "Copying $PLUGIN_SRC_DIR to $PLUGIN_DEST_DIR..."
    mkdir -p "$PLUGIN_DEST_DIR"
    cp -r "$PLUGIN_SRC_DIR" "$PLUGIN_DEST_DIR"
    echo "Plugin library successfully copied to $PLUGIN_DEST_DIR."
else
    echo "Plugin library directory $PLUGIN_SRC_DIR does not exist."
    exit 1
fi

# Remove the cloned repository directory after copying the plugin library
echo "Removing the cloned repository directory $TARGET_DIR..."
rm -rf "$TARGET_DIR"
if [ $? -eq 0 ]; then
    echo "Cloned repository directory $TARGET_DIR removed successfully."
else
    echo "Failed to remove the cloned repository directory $TARGET_DIR."
    exit 1
fi

# Run 'go mod tidy' to ensure dependencies are up to date
echo "Running 'go mod tidy'..."
go mod tidy
if [ $? -eq 0 ]; then
    echo "'go mod tidy' completed successfully."
else
    echo "Failed to run 'go mod tidy'."
    exit 1
fi

# Build the Go project
echo "Building the Go project..."
go build -o dirupdate
if [ $? -eq 0 ]; then
    echo "Go project built successfully."
else
    echo "Failed to build the Go project."
    exit 1
fi

# Check if the built executable exists and run it
if [ -f "./dirupdate" ]; then
    echo "Running './dirupdate'..."
    ./dirupdate
elif [ -f "./dirupdate.exe" ]; then
    echo "Running './dirupdate.exe'..."
    ./dirupdate.exe
else
    echo "Executable not found. Build might have failed."
    exit 1
fi