# Changelog

All notable changes to this fork are documented here.

Upstream changes from [clarkzjw/starlink_exporter](https://github.com/clarkzjw/starlink_exporter) are not repeated.

---

## [v0.0.8] — 2026-01-03

- Post-install script now enables and starts the `starlink-exporter` service automatically after `.deb` installation

## [v0.0.7] — 2026-01-03

- Dedicated `starlink-exporter` system user and group (service no longer runs as root)
- Systemd security hardening: `NoNewPrivileges`, `PrivateTmp`, `ProtectSystem=strict`, `ProtectHome`
- State directory `/var/lib/starlink-exporter` created on install, removed on purge
- Debian workflow updated to append installation instructions to GoReleaser release notes
- Docker Hub authentication removed from CI

## [v0.0.3] — 2026-01-03

- Docker image templates and manifests removed from `.goreleaser.yml`

## [v0.0.1] — 2026-01-03

- Initial Debian packaging for amd64, arm64, armhf
- `Makefile` updated for cross-compilation (Linux, macOS, Windows) with optional UPX compression
- `starlink-exporter.service` systemd unit file added
- GitHub Actions workflow for building and publishing `.deb` packages
- New Prometheus metrics:
  - `starlink_dish_latitude`, `starlink_dish_longitude`, `starlink_dish_altitude`
  - `starlink_dish_boresight_azimuth_deg_diff`, `starlink_dish_boresight_elevation_deg_diff`
  - `starlink_dish_tilt_angle_deg`
  - `starlink_dish_snr_above_noise_floor`, `starlink_dish_snr_persistently_low`
  - `starlink_dish_initialization_duration_seconds`
  - `starlink_dish_power_supply_connected`
  - Additional obstruction detail and quaternion orientation metrics
- `device_id` label added to all metrics
