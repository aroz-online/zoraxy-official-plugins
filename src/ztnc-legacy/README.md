## Global Area Network Plugin

This plugin implements a user interface for ZeroTier Network Controller in Zoraxy. 

This is identical to the original Zoraxy GAN function that provide ZeroTier network controller user interface in Zoraxy. However, since most user are using Zoraxy in docker and ignoring that feature, this has been moved out to an independent plugin for easier maintenance. 

## Build from source

```
go mod tidy
go build
```

After the build is done, create a folder under Zoraxy plugin directory (usually under ./plugins) named ztnc and put the binary into it. So after the process is completed, you should be able to see the ztnc inside the plugin folder with structure like this.

```
(Assuming you are currently in the working directory of the zoraxy executable)
$ ls ./plugins/ztnc/
ztnc

$ sudo ./zoraxy -port=:8000
```

## Setup

1. Get your zerotier authtoken from your zerotier installation
2. Create a file named "authtoken.secret" and paste the authtoken into the file
3. Start ztnc from Zoraxy plugin manager



## License

AGPL