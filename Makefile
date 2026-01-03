BINARY_NAME=starlink_exporter
CMD_PATH=./cmd/starlink_exporter
LDFLAGS=-s -w -extldflags '-static'

all:
	go build -o $(BINARY_NAME) $(CMD_PATH)/main.go

# Build for current platform
build:
	go build -o $(BINARY_NAME) $(CMD_PATH)

# Cross-compile for Linux ARM64 (aarch64) - optimized
build-linux-arm64:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags="$(LDFLAGS)" -trimpath -o $(BINARY_NAME)_linux_arm64 $(CMD_PATH)

# Cross-compile for Linux AMD64 - optimized
build-linux-amd64:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="$(LDFLAGS)" -trimpath -o $(BINARY_NAME)_linux_amd64 $(CMD_PATH)

# Cross-compile for Linux ARM (32-bit) - optimized
build-linux-arm:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=7 go build -ldflags="$(LDFLAGS)" -trimpath -o $(BINARY_NAME)_linux_arm $(CMD_PATH)

# Build all Linux architectures
build-linux-all: build-linux-amd64 build-linux-arm64 build-linux-arm

# Cross-compile for macOS AMD64 - optimized
build-darwin-amd64:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -trimpath -o $(BINARY_NAME)_darwin_amd64 $(CMD_PATH)

# Cross-compile for macOS ARM64 (Apple Silicon) - optimized
build-darwin-arm64:
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" -trimpath -o $(BINARY_NAME)_darwin_arm64 $(CMD_PATH)

# Cross-compile for Windows AMD64 - optimized
build-windows-amd64:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -trimpath -o $(BINARY_NAME)_windows_amd64.exe $(CMD_PATH)

# Build all platforms (Linux, macOS, Windows)
build-all: build-linux-amd64 build-linux-arm64 build-linux-arm build-darwin-amd64 build-darwin-arm64 build-windows-amd64

# UPX compressed builds (requires upx to be installed)
build-linux-arm64-upx: build-linux-arm64
	upx --best --lzma $(BINARY_NAME)_linux_arm64

build-linux-amd64-upx: build-linux-amd64
	upx --best --lzma $(BINARY_NAME)_linux_amd64

build-linux-all-upx: build-linux-all
	upx --best --lzma $(BINARY_NAME)_linux_amd64 $(BINARY_NAME)_linux_arm64 $(BINARY_NAME)_linux_arm

# Debian packaging
build-deb:
	@./scripts/build-deb.sh

# Publish to GitHub Packages (requires gh CLI)
publish-deb:
	@./scripts/publish-deb.sh

clean:
	rm -f $(BINARY_NAME) $(BINARY_NAME)_linux_* $(BINARY_NAME)_darwin_* $(BINARY_NAME)_windows_*
	rm -rf build/

.PHONY: all build build-linux-arm64 build-linux-amd64 build-linux-arm build-linux-all build-darwin-amd64 build-darwin-arm64 build-windows-amd64 build-all build-linux-arm64-upx build-linux-amd64-upx build-linux-all-upx build-deb publish-deb clean
