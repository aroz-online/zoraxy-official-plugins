# Zoraxy Official Plugins
The official plugin list for Zoraxy

**Warning! This repo is still in very early stage of development and things might not work as expected. Please use this with your own risk.**



## Introduction

Since Zoraxy v3.2.0, we introduced a new plugin system and moved some of the features from the Zoraxy main branch into the new plugin system. This is generally a more flexible way to handle features that not everyone uses and allow thirds party developers to add their plugins easily. 



## Installation

### Use the Zoraxy Plugin Manager

Follow the Zoraxy in-system plugin manager guideline to download and install plugins.

### Manual Installation
If you want to install a plugin without using the in-system plugin manager, follow these steps.

- Locate the plugin you want to download in the `app` folder
- Clone & build or download the prebuild binary in the target repo
- Create a folder inside Zoraxy plugin folder named identical to the plugin ID (you can check this with `./plugin_filename -introspec`)
- Place the plugin binary (and other required file and folder structure) into the folder with the plugin name
- Restart Zoraxy

## Contribute Plugin

To add a third-party/community plugin to the plugin store (so Zoraxy instances can discover it via `index2.json`), follow these steps.

1. Prepare the plugin repository (maintainer/developer)

	The plugin repository root should contain three optional/required files that the indexer will read directly from the remote repo's default branch (`main` or `master`):

	- `.introspect` (required): JSON output produced by your plugin's `introspect` step. This should match the structure used by Zoraxy's plugin introspection (fields like `id`, `name`, `author`, `author_contact`, `description`, `url`, `type`, `version_major`, `version_minor`, `version_patch`, etc.).
	- `.releaseurl` (required): a plain-text URL (or path) that will be used as the base to compose per-arch release download URLs (the indexer appends `/{pluginID}_{arch}` and adds `.exe` for windows). If omitted, the download URLs in `index2.json` will be empty.
	- `icon.png` (optional): plugin icon. If missing the repository fallback is `dummy.png` in this repo's `directories/icon`.

	Example files in plugin repo root:

	- `.introspect` (example JSON)

	```json
	{
	  "id": "org.aroz.zoraxy.ztnc",
	  "name": "ztnc",
	  "author": "aroz.org",
	  "author_contact": "noreply@aroz.org",
	  "description": "UI for ZeroTier Network Controller",
	  "url": "https://github.com/aroz-online/ztnc",
	  "type": 1,
	  "version_major": 0,
	  "version_minor": 0,
	  "version_patch": 0
	}
	```

	- `.releaseurl` (example content):

	```text
	https://github.com/aroz-online/ztnc/releases/latest/download
	```

2. Add a registry entry in this repository (contributor)

	Create a JSON file under `apps/` with a short name (one file per plugin). The file must include at least the `repo_url`; recommended fields are `author`, `contact` and an optional `min_zoraxy_version`.

	Example `apps/org.aroz.zoraxy.ztnc.json`:

	```json
	{
	  "repo_url": "https://github.com/aroz-online/ztnc",
	  "author": "aroz.org",
	  "contact": "noreply@aroz.org",
	  "min_zoraxy_version": "3.0.0"
	}
	```

3. Validate locally (maintainer/contributor)

	- Run the indexer locally to ensure the plugin metadata is picked up and `directories/index2.json` is generated correctly:

	```powershell
	cd tools\dirupdate2
	go run .
	# The tool writes to directories/index2.json
	```

	- Inspect `directories/index2.json`. Fields that could not be retrieved (missing `.introspect`, `.releaseurl`, `icon.png`) will be empty strings or empty objects.

4. Create a Pull Request

	- Fork this repository, add your `apps/<plugin>.json` into `apps/`, commit and open a pull request against `aroz-online/zoraxy-official-plugins`.
	- In the PR description include the plugin repository URL, author contact, and confirm that the remote repo contains `.introspect` and (optionally) `.releaseurl` and `icon.png` at the repository root.

Notes

- `index2.json` uses the same structure as the original `index.json` but does not include checksum hashes. Download URL composition depends on `.releaseurl` being present in the plugin repository.
- If `.introspect` is missing in the remote repo the indexer cannot populate the plugin's introspection metadata and will skip or produce an entry with empty fields.
- If you need help producing `.introspect` output for your plugin, check the `tools/dirupdate/mod/zoraxy_plugin` README or consult existing plugin repos for examples.



## License

Please refer to the LICENSE file for more information.
