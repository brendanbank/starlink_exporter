# Changelog

All notable changes to this fork are documented here.

Upstream changes from [clarkzjw/starlink_exporter](https://github.com/clarkzjw/starlink_exporter) are not repeated.

---

## [Unreleased]

- Debian packages are now built by GoReleaser's `nfpms` section instead of a
  separate workflow and hand-rolled `dpkg-deb` script
- Removed the `Build and Release Debian Packages` workflow, which raced the
  GoReleaser run and failed to upload packages when GoReleaser took longer than
  its fixed 30-second wait
- Package files are now owned by `root:root`; previously everything except the
  binary was installed owned by the CI runner's uid/gid
- Packages now ship an `md5sums` control file, so `dpkg -V` can verify them
- Removed the unused `debian/` debhelper configuration, which was never used by
  the build, and `scripts/build-deb.sh` / `scripts/publish-deb.sh`
- `make build-deb` now shells out to goreleaser and writes to `dist/`
- Upgrades now restart the service. The postinstall script ran `systemctl
  start`, which is a no-op when the service is already running, so an upgrade
  left the old binary running until the next manual restart or reboot. It now
  runs `try-restart` on upgrade and `start` on fresh install

## [v0.0.9] — 2026-08-02

- Bumped `clarkzjw/starlink-grpc-golang` from `v1.0.20250818` to `v1.0.20260526` (dish protoset `2026.05.26.mr80668`), pulling in grpc `1.79.1` and protobuf `1.36.10`
- New Prometheus metrics:
  - `starlink_dish_pnt_filter_convergence_state`
  - `starlink_dish_alert_no_ethernet_link`
- `starlink_dish_outage_duration` now resolves the `SKY_SEARCH` and `INHIBIT_RF` outage causes instead of leaving them unnamed

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
