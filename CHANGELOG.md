# Changelog

All notable changes to this fork are documented here.

Upstream changes from [clarkzjw/starlink_exporter](https://github.com/clarkzjw/starlink_exporter) are not repeated.

---

## [Unreleased]

- New Prometheus metrics for dish data the exporter fetched but discarded:
  - History aggregates from `GetHistory` (`starlink_dish_history_*`): ping drop
    seconds, latency min/mean/max, throughput mean/peak, bytes transferred,
    power min/max, and outage count/duration per cause. Only `PowerIn` was used
    before, so 15 minutes of per second samples were thrown away on every scrape
  - Bandwidth restriction reasons (`dl`/`ul_bandwidth_restricted_reason`),
    software update state (`reboot_reason`, `swupdate_reboot_ready`,
    `software_update_progress`, `software_update_requires_reboot`,
    `seconds_until_swupdate_reboot_possible`) and device flags
    (`is_cell_disabled`, `has_actuators`, `is_moving_fast_persisted`,
    `high_power_test_mode`, `has_signed_cals`, `user_debug_mode_enabled`,
    `mac_flag`, `nat_flag`, `account_shard`, `connected_routers`,
    `downstream_routers`)
  - PLC, battery, UPSU and APS power accessory telemetry
  - Diagnostics beyond GPS time: `hardware_self_test`,
    `hardware_self_test_code`, `stowed`, `overage_rate_limited`
  - Alerts: `dbf_telem_stale`, `dish_water_detected`, `router_water_detected`,
    `upsu_router_port_slow`, `slow_eth_speeds_100`
  - GPS: `gps_no_sats_after_ttff`, `gps_inhibited`
- `GetStatus` is now requested once per scrape instead of four times; the status,
  obstruction, alert and alignment collectors share one response
- Removed `starlink_dish_wedge_fraction_obstruction_ratio` and
  `starlink_dish_wedge_abs_fraction_obstruction_ratio`, which were described but
  never collected, so they never appeared in `/metrics`
- `computeSampleRange` no longer returns negative ring buffer offsets when the
  dish has been up for less than a full history buffer
- `GetLocation` denial by policy is now logged once and location collection is
  disabled, instead of logging an error on every scrape
- New provisioned Grafana dashboard, `contrib/config/grafana/provisioning/dashboards/Starlink-exporter.json`,
  covering the metrics above: outages and reliability, throughput history,
  service state, obstruction detail and pointing/GPS. It selects its Prometheus
  datasource through a `datasource` variable, so it can be provisioned or
  imported into any Grafana. The older `Starlink.json` still ships unchanged and
  still refers to metrics the exporter no longer emits
- `contrib/docker-compose.yaml` installs `marcusolsson-dynamictext-panel`, which
  the obstruction map panel needs
- README corrections: the boresight difference metrics were documented under
  names that do not exist (`boresight_azimuth_deg_diff` rather than
  `boresight_azimuth_diff_deg`), `starlink_dish_power_supply_connected` was
  documented but has never been collected, and
  `starlink_dish_snr_above_noise_floor` was described as the inverse of what it
  reports. The `/health` and `/infrequentMetrics` endpoints, the
  `STARLINK_GRPC_ADDR_PORT` and `IFACE` environment variables, and the fact that
  dishes now commonly refuse `GetLocation` are documented for the first time.
  `starlink_dish_public_ip_pop` gets its own section: what each label is derived
  from, that `dig` is an undeclared runtime dependency for the PoP labels, and
  that the lookup has to egress through the dish to describe the dish.
  Badges point at this fork rather than upstream, and the release table that
  stopped at v0.0.8 is replaced by a link to this file

## [v0.0.10] — 2026-08-02

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
