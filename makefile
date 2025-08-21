.PHONY: all build clean dirupdate dirupdate-noclone

# Build all platforms and architectures
# The build process is done in the plugin source folder of src/*/.
# Each plugin source folder should contain its own makefile that handles the specific build process for that plugin

all:
	@echo "Cleaning up previous builds..."
	$(MAKE) clean
	@echo "Building all plugins..."
	$(MAKE) build
	$(MAKE) dirupdate
	@echo "All plugins built successfully."
	
build:
	rm -rf dist
	@for dir in src/*/; do \
		$(MAKE) -C $$dir; \
		platform=$$(uname -s | tr '[:upper:]' '[:lower:]'); \
		arch=$$(uname -m); \
		folder_name=$$(basename $$dir); \
		mkdir -p dist/$$folder_name; \
		mv "$$dir/build/"* "dist/$$folder_name/"; \
		done

dirupdate:
	@echo "Updating directories..."
	@cd ./tools/dirupdate && ./update.sh

dirupdate-noclone:
	@echo "Updating directories without cloning..."
	@cd ./tools/dirupdate && ./update.sh -noclone

clean: 
	rm -rf dist
