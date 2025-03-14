# Zoraxy Official Plugins
The offical plugin list for Zoraxy



## Introduction

Since Zoraxy v3.2.0, we introduced a new plugin system and moved some of the features from the Zoraxy main branch into the new plugin system. This is generally a more flexible way to handle features that not everyone uses and allow thirds party developers to add their plugins easily. 



## Installation

Here are the steps to install plugin to Zoraxy

1. Locate the plugin folder (The plugin folder is usually located at the same directory as your Zoraxy executable)
2. Create a folder with the name of your plugin executable. For example, if you have a plugin executable named `ztnc.exe` (Windows) or `ztnc` (Linux), create a folder named "ztnc" under the plugin folder so you now have `./plugins/ztnc/ztnc(.exe)`
3. Refresh Zoraxy Plugin list or restart Zoraxy. Your should be able to see your new plugin in the Plugin Manager menu. 



**Notes on the plugin folder location**

Plugins are designed to be installed to the `plugins/{plugin_name}/` directory relative to the Zoraxy working directory. For example, if your Zoraxy binary executable is located at `/home/user/zoraxy/`, the plugin folder can usually be found at `/home/user/zoraxy/plugins`. For some special installation where the `cwd` is set to an alternative path, you might need to install them to the `{cwd}/plugins` folder.

## Build From Source

To build the plugins, you will need the latest Go compiler and GNU make. It is recommended that you build on the machine that you plan to run them as some of them might require CGO. 

#### Build Single Plugin

To build just one plugin, go into the respective plugin directory and run go build. Assume the plugin you want to build is `ztnc`

```bash
cd plugins/ztnc/
go mod tidy
go build

# and you should see ztnc.exe or ztnc (depends on the platform you are using)
```

#### Build All Plugins

You can build all the plugin with a simple makefile call

```bash
cd plugins
make
```

