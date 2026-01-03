# Debian Packaging for Starlink Exporter

This directory contains the Debian packaging configuration for building `.deb` packages.

## Building Packages

### Prerequisites

- Go 1.24 or later
- `dpkg-deb` (usually pre-installed on Debian/Ubuntu)

### Build all architectures

```bash
make build-deb
```

This will create packages for:
- `amd64` - x86_64 processors
- `arm64` - 64-bit ARM processors (Raspberry Pi 4, newer)
- `armhf` - 32-bit ARM processors (older Raspberry Pi)

Packages will be created in `build/deb/`:
- `starlink-exporter_1.0.0_amd64.deb`
- `starlink-exporter_1.0.0_arm64.deb`
- `starlink-exporter_1.0.0_armhf.deb`

### Specify version

```bash
VERSION=2.0.0 make build-deb
```

## Publishing to GitHub

### Using Make

```bash
# Build and publish
make build-deb
VERSION=1.0.0 make publish-deb
```

### Automatic via GitHub Actions

Push a tag to automatically build and publish:

```bash
git tag v1.0.0
git push origin v1.0.0
```

Or trigger manually from GitHub Actions tab with custom version.

## Package Contents

Each package includes:
- Binary: `/usr/bin/starlink_exporter`
- Systemd service: `/lib/systemd/system/starlink-exporter.service`
- Documentation: `/usr/share/doc/starlink-exporter/`

## Installation

### From GitHub Releases

```bash
# AMD64
wget https://github.com/brendanbank/starlink_exporter/releases/download/v1.0.0/starlink-exporter_1.0.0_amd64.deb
sudo dpkg -i starlink-exporter_1.0.0_amd64.deb

# ARM64
wget https://github.com/brendanbank/starlink_exporter/releases/download/v1.0.0/starlink-exporter_1.0.0_arm64.deb
sudo dpkg -i starlink-exporter_1.0.0_arm64.deb

# ARMHF (Raspberry Pi)
wget https://github.com/brendanbank/starlink_exporter/releases/download/v1.0.0/starlink-exporter_1.0.0_armhf.deb
sudo dpkg -i starlink-exporter_1.0.0_armhf.deb
```

### From local build

```bash
sudo dpkg -i build/deb/starlink-exporter_1.0.0_<arch>.deb
```

## Post-Installation

Start the service:
```bash
sudo systemctl start starlink-exporter
sudo systemctl enable starlink-exporter
```

Check status:
```bash
sudo systemctl status starlink-exporter
```

View logs:
```bash
sudo journalctl -u starlink-exporter -f
```

## Uninstallation

```bash
sudo apt remove starlink-exporter
# or
sudo dpkg -r starlink-exporter
```

## Package Structure

```
/
├── usr/
│   ├── bin/
│   │   └── starlink_exporter          # Main binary
│   └── share/
│       └── doc/
│           └── starlink-exporter/
│               ├── LICENSE             # License file
│               └── README.md           # Documentation
└── lib/
    └── systemd/
        └── system/
            └── starlink-exporter.service  # Systemd unit
```

## Maintainer Scripts

- `postinst`: Runs after installation - enables systemd service
- `prerm`: Runs before removal - stops and disables service

## Verifying Package

```bash
# List contents
dpkg -c starlink-exporter_1.0.0_amd64.deb

# Show package info
dpkg -I starlink-exporter_1.0.0_amd64.deb

# After installation, verify files
dpkg -L starlink-exporter
```
