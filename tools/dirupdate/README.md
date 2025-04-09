# Directory Info Update Tool

This is an internal tool to update the directory info of Zoraxy official plugin list. It is design for CICD workflow to build plugin index for the Zoraxy plugin manager, but you can also run it manually if you are hosting your own plugin store.

### IMPORTANT

Before running this tool, make sure all the plugins has already been built using the make file in the root directory. 

### Usage

To update the `/directories` folder, simply run the update bash script as follows.

```
./update.sh
```

This will also build the go code that generate the directories checksum. Make sure you have go compiler installed on your development environment. 

### Procedures

 This tool basically do the following tasks in one go program.

1. Cone the zoraxy main branch
2. Extract the zoraxy_plugin directory from the main branch and move it to `./mod `
3. Remove the zoraxy source code
4. Use the new `zoraxy_plugin` introspect definition to extract introspect JSON from pluigins in the `./dist` folder
5.  Generate checksum for each of the plugin binaries
6. Copy the icon.png for a given plugin to `./directories/icon/{plugin_name}.png`
7. Repeat step 4 - 6 until all plugins specification and checksum has been extracted and calculated
8. Write the plugin information into index.json