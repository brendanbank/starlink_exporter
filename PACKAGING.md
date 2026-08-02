# Debian Package Build & Release Guide

Debian packages are built by [GoReleaser](https://goreleaser.com) via its `nfpms`
section, configured in [`.goreleaser.yml`](.goreleaser.yml). The same run that
produces the release archives also produces the `.deb` files and uploads
everything to the GitHub release, so there is no separate packaging step.

## Quick Start

### Build packages locally

```bash
make build-deb
```

This runs `goreleaser release --snapshot --clean --skip=publish`, creating
`.deb` packages for `amd64`, `arm64`, and `armhf` in `dist/` alongside the
release archives. Requires [goreleaser](https://goreleaser.com/install/)
(`brew install goreleaser`).

### Inspect a package

```bash
dpkg -I dist/starlink-exporter_*_amd64.deb   # control metadata
dpkg -c dist/starlink-exporter_*_amd64.deb   # file listing
```

### Test installation

```bash
sudo dpkg -i dist/starlink-exporter_*_$(dpkg --print-architecture).deb
sudo systemctl status starlink-exporter
```

On a non-Debian workstation, a container works well:

```bash
docker run --rm -v "$PWD/dist:/pkg:ro" debian:bookworm \
  bash -c 'dpkg -i /pkg/starlink-exporter_*_amd64.deb; dpkg -V starlink-exporter'
```

## Publishing a Release

Push a tag. The [`goreleaser` workflow](.github/workflows/release.yaml) builds
archives and `.deb` packages, creates the GitHub release, and uploads all
artifacts:

```bash
git tag -s v1.0.0 -m "v1.0.0"
git push origin v1.0.0
```

The release notes footer with Debian install instructions is generated from the
`release.footer` template in `.goreleaser.yml`.

## Package Layout

| Path | Mode | Source |
| --- | --- | --- |
| `/usr/bin/starlink_exporter` | 0755 | goreleaser build output |
| `/lib/systemd/system/starlink-exporter.service` | 0644 | `starlink-exporter.service` |
| `/usr/share/doc/starlink-exporter/LICENSE` | 0644 | `LICENSE` |
| `/usr/share/doc/starlink-exporter/README.md` | 0644 | `README.md` |

All files are owned by `root:root`.

## Maintainer Scripts

Lives in [`packaging/deb/`](packaging/deb/), wired up via the `scripts:` block
in `.goreleaser.yml`.

| Script | Runs | Does |
| --- | --- | --- |
| `postinstall.sh` | after install/upgrade | creates the `starlink-exporter` system user and group, creates `/var/lib/starlink-exporter`, reloads systemd, enables and starts the service |
| `preremove.sh` | before removal | stops and disables the service |
| `postremove.sh` | on purge | removes `/var/lib/starlink-exporter` and deletes the user and group |

On `remove` (as opposed to `purge`) the user and state directory are kept
deliberately, so journal ownership and any local state survive a reinstall.

## User Installation Guide

### From GitHub Releases

```bash
# Determine your architecture
dpkg --print-architecture

# Download package (replace <version> and <arch> with amd64, arm64, or armhf)
wget https://github.com/brendanbank/starlink_exporter/releases/download/v<version>/starlink-exporter_<version>_<arch>.deb

# Install — the service is enabled and started automatically
sudo dpkg -i starlink-exporter_<version>_<arch>.deb

# Check status
sudo systemctl status starlink-exporter

# View logs
sudo journalctl -u starlink-exporter -f
```

### Uninstallation

```bash
sudo apt remove starlink-exporter    # keeps user and /var/lib/starlink-exporter
sudo apt purge starlink-exporter     # also removes them
```

## Troubleshooting

### `goreleaser: command not found`

```bash
brew install goreleaser
# or see https://goreleaser.com/install/
```

### Release ran but no `.deb` appeared

Check that the `nfpms` section in `.goreleaser.yml` still resolves the
maintainer scripts under `packaging/deb/`; goreleaser fails the run if a
referenced script is missing. Validate the config without building:

```bash
goreleaser check
```

### Service fails to start after install

The unit runs as the unprivileged `starlink-exporter` user with
`ProtectSystem=strict`, writing only to `/var/lib/starlink-exporter`. If the
postinstall script was interrupted, that directory or the user may be missing:

```bash
getent passwd starlink-exporter
ls -ld /var/lib/starlink-exporter
sudo dpkg-reconfigure starlink-exporter
```
