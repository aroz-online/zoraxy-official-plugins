.PHONY: all

all: zoraxy_linux_386 zoraxy_linux_amd64 zoraxy_linux_arm zoraxy_linux_arm64 zoraxy_linux_mipsle zoraxy_linux_riscv64 zoraxy_windows_amd64.exe

zoraxy_linux_386:
	mkdir -p ./build
	CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -o ztnc_linux_386
	mv ztnc_linux_386 ./build/

zoraxy_linux_amd64:
	mkdir -p ./build
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ztnc_linux_amd64
	mv ztnc_linux_amd64 ./build/

zoraxy_linux_arm:
	mkdir -p ./build
	CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -o ztnc_linux_arm
	mv ztnc_linux_arm ./build/

zoraxy_linux_arm64:
	mkdir -p ./build
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o ztnc_linux_arm64
	mv ztnc_linux_arm64 ./build/

zoraxy_linux_mipsle:
	mkdir -p ./build
	CGO_ENABLED=0 GOOS=linux GOARCH=mipsle go build -o ztnc_linux_mipsle
	mv ztnc_linux_mipsle ./build/

zoraxy_linux_riscv64:
	mkdir -p ./build
	CGO_ENABLED=0 GOOS=linux GOARCH=riscv64 go build -o ztnc_linux_riscv64
	mv ztnc_linux_riscv64 ./build/

zoraxy_windows_amd64.exe:
	mkdir -p ./build
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ztnc_windows_amd64.exe
	mv ztnc_windows_amd64.exe ./build/

.PHONY: all zoraxy_linux_386 zoraxy_linux_amd64 zoraxy_linux_arm zoraxy_linux_arm64 zoraxy_linux_mipsle zoraxy_linux_riscv64 zoraxy_windows_amd64.exe