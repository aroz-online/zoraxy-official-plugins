.PHONY: all foldertree clean dirupdate 
# Build all platforms and architectures
# The build process is done in the plugin source folder of src/*/.
# Each plugin source folder should contain its own makefile that handles the specific build process for that plugin

all:
	@echo "Cleaning up previous builds..."
	$(MAKE) clean
	@echo "Gathering and building all plugins..."
	$(MAKE) foldertree
	$(MAKE) dirupdate
	@echo "All plugins built successfully."
	

dirupdate:
	@echo "Updating directories..."
	@cd ./tools/dirupdate2 && go run .

clean: 
	rm -f directories/index2.json
