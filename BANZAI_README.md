<p align="center">
  <h3 align="center">Starlink Prometheus Exporter Monitoring Stack</h3>
</p>

---

A [Starlink](https://www.starlink.com/) exporter for Prometheus. Not affiliated with or acting on behalf of Starlink(™)

This is a fork of [clarkzjw/starlink_exporter](https://github.com/clarkzjw/starlink_exporter) with Debian packaging, security hardening, and additional metrics.

[![build](https://github.com/clarkzjw/starlink_exporter/actions/workflows/build.yaml/badge.svg)](https://github.com/clarkzjw/starlink_exporter/actions/workflows/build.yaml)
[![License](https://img.shields.io/github/license/clarkzjw/starlink_exporter)](/LICENSE)
[![Release](https://img.shields.io/github/release/brendanbank/starlink_exporter.svg)](https://github.com/brendanbank/starlink_exporter/releases/latest)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/clarkzjw/starlink_exporter)

Original repositories:

- https://github.com/danopstech/starlink_exporter
- https://github.com/clarkzjw/starlink_exporter

---

Starlink gRPC protobuf for Golang: [clarkzjw/starlink-grpc-golang](https://github.com/clarkzjw/starlink-grpc-golang)

Starlink dish firmware tracking website: https://starlinktrack.com/firmware/dishy

---

## Installation

### Debian/Ubuntu (recommended)

Pre-built `.deb` packages are published with each release for amd64, arm64, and armhf (Raspberry Pi).

```bash
wget https://github.com/brendanbank/starlink_exporter/releases/latest/download/starlink-exporter_<version>_amd64.deb
sudo dpkg -i starlink-exporter_<version>_amd64.deb
```

The service is enabled and started automatically after installation.

See [PACKAGING.md](PACKAGING.md) for full build and publish instructions.

### Binaries

For pre-built binaries see the [releases](https://github.com/brendanbank/starlink_exporter/releases).

```bash
./starlink_exporter [flags]
```

### Docker Compose Stack

Use the `docker-compose.yaml` in the [contrib](./contrib) directory.

```bash
docker-compose up -d
```

---

## Usage

### Flags

`starlink_exporter` is configured through optional command line flags:

```bash
$ ./starlink_exporter -h
Usage of ./starlink_exporter:
  -address string
        IP address and port to reach dish (default "192.168.100.1:9200")
  -port string
        listening port to expose metrics on (default "9817")
```

### Service management

```bash
sudo systemctl status starlink-exporter
sudo journalctl -u starlink-exporter -f
```

---

## Fork Changes

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

All metrics include a `device_id` label for identifying individual dishes when multiple units are monitored.

### Security Hardening

The service runs as a dedicated `starlink-exporter` system user (not root), with the following systemd security directives:

- `NoNewPrivileges=true`
- `PrivateTmp=true`
- `ProtectSystem=strict`
- `ProtectHome=true`

A state directory `/var/lib/starlink-exporter` is created on install and removed on purge.

### Build System

- `Makefile` updated for cross-compilation targeting Linux, macOS, and Windows (amd64, arm64, armhf)
- Optional UPX compression for binary size reduction
- Docker image building and Docker Hub publishing removed; releases are binary archives + `.deb` only
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

## Grafana Dashboard

<p align="center">
	<img src="https://github.com/clarkzjw/starlink_exporter/raw/main/static/Screenshot.jpeg" width="95%">
</p>
