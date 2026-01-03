# Debian Package Build & Release Guide

This guide explains how to build and publish Debian packages for the Starlink Exporter.

## Quick Start

### Build packages locally

```bash
make build-deb
```

This creates `.deb` packages for AMD64, ARM64, and ARMHF in `build/deb/`.

### Test installation locally

```bash
sudo dpkg -i build/deb/starlink-exporter_1.0.0_amd64.deb
sudo systemctl start starlink-exporter
sudo systemctl status starlink-exporter
```

### Publish to GitHub Releases

```bash
VERSION=1.0.0 make publish-deb
```

**Requirements**: GitHub CLI (`gh`) must be installed and authenticated.

## Manual Build Process

### 1. Set version (optional)

```bash
export VERSION=1.0.0
```

### 2. Build packages

```bash
./scripts/build-deb.sh
```

### 3. Verify packages

```bash
ls -lh build/deb/
dpkg -I build/deb/starlink-exporter_1.0.0_amd64.deb
dpkg -c build/deb/starlink-exporter_1.0.0_amd64.deb
```

### 4. Publish to GitHub

```bash
./scripts/publish-deb.sh
```

## Automated Release via GitHub Actions

### Using Git Tags

```bash
# Create and push a tag
git tag v1.0.0
git push origin v1.0.0
```

The GitHub Action will automatically:
1. Build packages for all architectures
2. Create a GitHub release
3. Upload packages to the release
4. Generate installation instructions

### Manual Trigger

1. Go to GitHub Actions tab
2. Select "Build and Release Debian Packages"
3. Click "Run workflow"
4. Enter version number (e.g., `1.0.0`)
5. Click "Run workflow"

## Package Details

### Architectures Supported

- **amd64** - Intel/AMD 64-bit (most servers and desktops)
- **arm64** - ARM 64-bit (Raspberry Pi 4, newer ARM devices)
- **armhf** - ARM 32-bit (older Raspberry Pi models)

### Package Contents

```
/usr/bin/starlink_exporter                    # Main binary
/lib/systemd/system/starlink-exporter.service # Systemd service
/usr/share/doc/starlink-exporter/LICENSE      # License
/usr/share/doc/starlink-exporter/README.md    # Documentation
```

### Post-installation

The package automatically:
- Enables the systemd service
- Does NOT start the service (user must start manually)

### Pre-removal

The package automatically:
- Stops the service
- Disables the service

## User Installation Guide

### From GitHub Releases

```bash
# Determine your architecture
dpkg --print-architecture

# Download package (replace <arch> with amd64, arm64, or armhf)
wget https://github.com/brendanbank/starlink_exporter/releases/download/v1.0.0/starlink-exporter_1.0.0_<arch>.deb

# Install
sudo dpkg -i starlink-exporter_1.0.0_<arch>.deb

# Start service
sudo systemctl start starlink-exporter
sudo systemctl enable starlink-exporter

# Check status
sudo systemctl status starlink-exporter

# View logs
sudo journalctl -u starlink-exporter -f
```

### Uninstallation

```bash
sudo systemctl stop starlink-exporter
sudo apt remove starlink-exporter
# or
sudo dpkg -r starlink-exporter
```

## Troubleshooting

### Build fails with "go: command not found"

Install Go 1.24 or later:
```bash
sudo apt install golang-go
# or download from https://go.dev/dl/
```

### dpkg-deb not found

Install dpkg:
```bash
sudo apt install dpkg
```

### GitHub CLI (gh) not found

Install gh:
```bash
# Ubuntu/Debian
sudo apt install gh

# Or see: https://cli.github.com/
```

### Permission denied when running scripts

Make scripts executable:
```bash
chmod +x scripts/*.sh
```

### Package already exists in release

The publish script uses `--clobber` to replace existing files:
```bash
gh release upload v1.0.0 build/deb/*.deb --clobber
```

## Versioning

Version numbers should follow semantic versioning (MAJOR.MINOR.PATCH):

- **MAJOR**: Breaking changes
- **MINOR**: New features, backward compatible
- **PATCH**: Bug fixes

Examples: `1.0.0`, `1.2.3`, `2.0.0-beta1`

## Testing Packages

### Test installation on Docker

```bash
# AMD64
docker run -it --rm -v $(pwd)/build/deb:/packages ubuntu:22.04 bash
apt update && apt install -y /packages/starlink-exporter_1.0.0_amd64.deb

# ARM64
docker run -it --rm --platform linux/arm64 -v $(pwd)/build/deb:/packages ubuntu:22.04 bash
apt update && apt install -y /packages/starlink-exporter_1.0.0_arm64.deb
```

### Verify package quality

```bash
# Check for lintian issues
lintian build/deb/starlink-exporter_1.0.0_amd64.deb

# Install lintian if needed
sudo apt install lintian
```

## Files Structure

```
.
├── debian/                           # Debian package metadata
│   ├── changelog                     # Version history
│   ├── compat                        # Debhelper compatibility
│   ├── control                       # Package dependencies/description
│   ├── rules                         # Build rules
│   ├── starlink-exporter.install     # File installation mapping
│   ├── starlink-exporter.postinst    # Post-install script
│   ├── starlink-exporter.prerm       # Pre-removal script
│   └── README.md                     # Packaging documentation
├── scripts/
│   ├── build-deb.sh                  # Build packages for all archs
│   └── publish-deb.sh                # Publish to GitHub
├── .github/
│   └── workflows/
│       └── release.yml               # GitHub Actions workflow
└── Makefile                          # Build targets
```

## CI/CD Integration

The GitHub Actions workflow (`.github/workflows/release.yml`) runs on:
- Git tags matching `v*` pattern
- Manual trigger from GitHub UI

It automatically:
- Builds all architectures
- Creates SHA256 checksums
- Creates GitHub release with installation instructions
- Uploads all packages and checksums

## Best Practices

1. **Always test locally before publishing**
   ```bash
   make build-deb
   sudo dpkg -i build/deb/starlink-exporter_*_$(dpkg --print-architecture).deb
   ```

2. **Use semantic versioning**
   - Increment PATCH for bug fixes
   - Increment MINOR for new features
   - Increment MAJOR for breaking changes

3. **Update debian/changelog**
   - Document changes for each release

4. **Test on target architectures**
   - Use Docker or actual hardware

5. **Verify checksums**
   - Always publish checksums.txt with releases

## Support

For issues with packaging:
- Check [debian/README.md](debian/README.md)
- Review build logs: `./scripts/build-deb.sh 2>&1 | tee build.log`
- Check GitHub Actions logs for automated builds

For application issues:
- See main [README.md](README.md)
- Open an issue on GitHub
