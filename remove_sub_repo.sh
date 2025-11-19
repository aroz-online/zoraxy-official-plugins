#!/bin/bash

# List all directories in src/
plugins=($(ls -d src/*/ 2>/dev/null | sed 's|src/||' | sed 's|/||'))

if [ ${#plugins[@]} -eq 0 ]; then
    echo "No plugins found in src/"
    exit 1
fi

echo "Available plugins:"
for i in "${!plugins[@]}"; do
    echo "$((i+1)). ${plugins[$i]}"
done

echo "Enter the number of the plugin to remove, or 'q' to exit:"
read choice

if [ "$choice" = "q" ]; then
    exit 0
fi

if ! [[ "$choice" =~ ^[0-9]+$ ]] || [ "$choice" -lt 1 ] || [ "$choice" -gt ${#plugins[@]} ]; then
    echo "Invalid choice."
    exit 1
fi

plugin=${plugins[$((choice-1))]}

echo "Removing submodule src/$plugin"
git submodule deinit -f "src/$plugin"
git rm -f "src/$plugin"
rm -rf ".git/modules/src/$plugin"

# Remove entry from .gitmodules
if [ -f ".gitmodules" ]; then
    sed -i "/\[submodule \"src\/$plugin\"\]/,/^$/d" .gitmodules
fi

echo "Submodule removed."
