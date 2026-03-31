# Starlink Exporter — Fork Changes

This is a fork of [clarkzjw/starlink_exporter](https://github.com/clarkzjw/starlink_exporter), a Prometheus exporter for Starlink dish metrics via gRPC.

---

## What's Different in This Fork

### New Prometheus Metrics

The following metrics have been added on top of upstream:

**Location**
- `starlink_dish_latitude` — dish latitude (degrees)
- `starlink_dish_longitude` — dish longitude (degrees)
- `starlink_dish_altitude` — dish altitude (meters)

**Alignment**
- `starlink_dish_boresight_azimuth_deg_diff` — difference between desired and actual boresight azimuth
- `starlink_dish_boresight_elevation_deg_diff` — difference between desired and actual boresight elevation
- `starlink_dish_tilt_angle_deg` — dish tilt angle

**Signal Quality**
- `starlink_dish_snr_above_noise_floor` — whether SNR is above the noise floor
- `starlink_dish_snr_persistently_low` — whether SNR is persistently low

**Dish Status**
- `starlink_dish_initialization_duration_seconds` — time taken to initialize the dish
- `starlink_dish_power_supply_connected` — power supply connectivity status
- Additional obstruction detail metrics
- Quaternion orientation values (ned2dish)

**Device ID Label**

All metrics now include a `device_id` label for identifying individual dishes when multiple units are monitored.

---

### Debian/Ubuntu Packaging

Pre-built `.deb` packages are published with each release for:
- `amd64` (x86_64)
- `arm64` (64-bit ARM)
- `armhf` (Raspberry Pi / 32-bit ARM)

Install from a release:

```bash
wget https://github.com/brendanbank/starlink_exporter/releases/latest/download/starlink-exporter_<version>_amd64.deb
sudo dpkg -i starlink-exporter_<version>_amd64.deb
```

See [PACKAGING.md](PACKAGING.md) for full build and publish instructions.

---

### Systemd Service

A `starlink-exporter.service` unit file is included. After `.deb` installation the service is enabled and started automatically.

```bash
sudo systemctl status starlink-exporter
sudo journalctl -u starlink-exporter -f
```

---

### Security Hardening

The service runs as a dedicated `starlink-exporter` system user (not root), with the following systemd security directives:

- `NoNewPrivileges=true`
- `PrivateTmp=true`
- `ProtectSystem=strict`
- `ProtectHome=true`

A state directory `/var/lib/starlink-exporter` is created on install and removed on purge.

---

### Build System

- `Makefile` updated for cross-compilation targeting Linux, macOS, and Windows (amd64, arm64, armhf)
- Optional UPX compression for binary size reduction
- Docker image building and Docker Hub publishing removed; releases are binary + `.deb` only
- GitHub Actions workflow builds and attaches `.deb` packages to the GoReleaser release

---

## Release History

| Version | Key Changes |
|---------|-------------|
| v0.0.8  | Auto-enable and start service on `.deb` install |
| v0.0.6/v0.0.7 | Dedicated system user, systemd security hardening |
| v0.0.3  | Docker build removed from GoReleaser |
| v0.0.1/v0.0.2 | Initial Debian packaging, cross-compilation Makefile, service file |
| (pre-tag) | New metrics: location, alignment, SNR, device ID label |

---

## Upstream

- Original fork base: [clarkzjw/starlink_exporter](https://github.com/clarkzjw/starlink_exporter)
- Starlink gRPC protobuf: [clarkzjw/starlink-grpc-golang](https://github.com/clarkzjw/starlink-grpc-golang)
