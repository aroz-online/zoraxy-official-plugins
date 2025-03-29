.PHONY: all

# Build all platforms and architectures
# The build process is done in the plugin source folder of src/*/.
# Each plugin source folder should contain its own makefile that handles the specific build process for that plugin

all:
	rm -rf dist
	@for dir in src/*/; do \
		$(MAKE) -C $$dir; \
		platform=$$(uname -s | tr '[:upper:]' '[:lower:]'); \
		arch=$$(uname -m); \
		folder_name=$$(basename $$dir); \
		mkdir -p dist/$$folder_name; \
		mv "$$dir/build/"* "dist/$$folder_name/"; \
		done