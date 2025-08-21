.PHONY: all

all: ztnc-legacylinux_386 ztnc-legacylinux_amd64 ztnc-legacylinux_arm ztnc-legacylinux_arm64 ztnc-legacylinux_mipsle ztnc-legacylinux_riscv64 ztnc-legacywindows_amd64.exe

ztnc-legacylinux_386:
	mkdir -p ./build
	CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -o ztnc-legacy_linux_386
	mv ztnc-legacy_linux_386 ./build/

ztnc-legacylinux_amd64:
	mkdir -p ./build
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ztnc-legacy_linux_amd64
	mv ztnc-legacy_linux_amd64 ./build/

ztnc-legacylinux_arm:
	mkdir -p ./build
	CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -o ztnc-legacy_linux_arm
	mv ztnc-legacy_linux_arm ./build/

ztnc-legacylinux_arm64:
	mkdir -p ./build
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o ztnc-legacy_linux_arm64
	mv ztnc-legacy_linux_arm64 ./build/

ztnc-legacylinux_mipsle:
	mkdir -p ./build
	CGO_ENABLED=0 GOOS=linux GOARCH=mipsle go build -o ztnc-legacy_linux_mipsle
	mv ztnc-legacy_linux_mipsle ./build/

ztnc-legacylinux_riscv64:
	mkdir -p ./build
	CGO_ENABLED=0 GOOS=linux GOARCH=riscv64 go build -o ztnc-legacy_linux_riscv64
	mv ztnc-legacy_linux_riscv64 ./build/

ztnc-legacywindows_amd64.exe:
	mkdir -p ./build
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ztnc-legacy_windows_amd64.exe
	mv ztnc-legacy_windows_amd64.exe ./build/

.PHONY: all ztnc-legacylinux_386 ztnc-legacylinux_amd64 ztnc-legacylinux_arm ztnc-legacylinux_arm64 ztnc-legacylinux_mipsle ztnc-legacylinux_riscv64 ztnc-legacywindows_amd64.exe