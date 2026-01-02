package exporter

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	// ned2dishQuaternion
	dishNed2dishQuaternionQScalar = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "ned2dish_quaternion_q_scalar"),
		"ned2dishQuaternion qScalar",
		[]string{"device_id"}, nil,
	)
	dishNed2dishQuaternionQX = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "ned2dish_quaternion_q_x"),
		"ned2dishQuaternion qX",
		[]string{"device_id"}, nil,
	)
	dishNed2dishQuaternionQY = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "ned2dish_quaternion_q_y"),
		"ned2dishQuaternion qY",
		[]string{"device_id"}, nil,
	)
	dishNed2dishQuaternionQZ = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "ned2dish_quaternion_q_z"),
		"ned2dishQuaternion qZ",
		[]string{"device_id"}, nil,
	)

	// Location Info
	dishLocationInfo = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "location_info"),
		"Dish Location Info",
		[]string{
			"device_id",
			"location_source",
			"lat",
			"lon",
			"alt",
			"sigmaM",
			"horizontalSpeedMps",
			"verticalSpeedMps"}, nil,
	)
	dishLatitude = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "latitude"),
		"Dish Latitude",
		[]string{"device_id"}, nil,
	)
	dishLongitude = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "longitude"),
		"Dish Longitude",
		[]string{"device_id"}, nil,
	)
	dishAltitude = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "altitude"),
		"Dish Altitude",
		[]string{"device_id"}, nil,
	)

	// DeviceInfo
	dishInfo = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "info"),
		"Running software versions and IDs of hardware",
		[]string{
			"device_id",
			"build_id",
			"hardware_version",
			"software_version",
			"generationNumber",
			"country_code",
			"bootcount",
			"utc_offset"}, nil,
	)
	dishInitializationDurationSeconds = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "initialization_duration_seconds"),
		"Initialization duration in seconds",
		[]string{
			"device_id",
			"attitudeInitialization",
			"burstDetected",
			"ekfConverged",
			"firstCplane",
			"firstPopPing",
			"gpsValid",
			"initialNetworkEntry",
			"networkSchedule",
			"rfReady",
			"stableConnection",
		}, nil,
	)
	dishConfig = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "dish_config"),
		"Dish Config",
		[]string{
			"device_id",
			"snow_melt_mode",
			"location_request_mode",
			"level_dish_mode",
			"power_save_mode",
		}, nil,
	)
	SoftwarePartitionsEqual = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "software_partitions_equal"),
		"Starlink Dish Software Partitions Equal.",
		[]string{"device_id"}, nil,
	)
	dishMobilityClass = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "mobility_class"),
		"Dish mobility class",
		[]string{"device_id", "mobility_class"}, nil,
	)
	userClassOfService = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "class_of_service"),
		"User class of service",
		[]string{"device_id", "class_of_service"}, nil,
	)
	dishReadyState = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "ready_state"),
		"Dish ready states",
		[]string{
			"device_id",
			"cady",
			"scp",
			"l1l2",
			"xphy",
			"aap",
			"rf",
		}, nil,
	)
	dishIsDev = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "is_dev"),
		"Starlink Dish is Dev.",
		[]string{"device_id"}, nil,
	)
	dishBootCount = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "boot_count"),
		"Starlink Dish boot count.",
		[]string{"device_id"}, nil,
	)
	dishAntiRollbackVersion = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "anti_rollback_version"),
		"Starlink Dish Anti Rollback Version.",
		[]string{"device_id"}, nil,
	)
	dishIsHit = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "is_hit"),
		"Starlink Dish is Hit.",
		[]string{"device_id"}, nil,
	)
	// BootInfo
	dishBootInfo = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "info_debug"),
		"Debug Dish Info",
		[]string{
			"device_id",
			"count_by_reason",
			"count_by_reason_delta",
			"last_reason",
			"last_count"}, nil,
	)
	dishAlignmentStats = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "alignment_stats"),
		"Starlink Dish Alignment Stats",
		[]string{
			"device_id",
			"hasActuators",
			"actuatorState",
			"tiltAngleDeg",
			"boresightAzimuthDeg",
			"boresightElevationDeg",
			"attitudeEstimationState",
			"attitudeUncertaintyDeg",
			"desiredBoresightAzimuthDeg",
			"desiredBoresightElevationDeg"}, nil,
	)
	dishBoresightAzimuthDiffDeg = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "boresight_azimuth_diff_deg"),
		"Difference between desired and actual boresight azimuth in degrees",
		[]string{"device_id"}, nil,
	)
	dishBoresightElevationDiffDeg = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "boresight_elevation_diff_deg"),
		"Difference between desired and actual boresight elevation in degrees",
		[]string{"device_id"}, nil,
	)
	dishTiltAngleDeg = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "tilt_angle_deg"),
		"Dish tilt angle in degrees from vertical",
		[]string{"device_id"}, nil,
	)
	// DeviceState
	dishUp = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "up"),
		"Was the last query of Starlink dish successful.",
		[]string{"device_id"}, nil,
	)
	// dishScrapeDurationSeconds = prometheus.NewDesc(
	// 	prometheus.BuildFQName(namespace, "dish", "scrape_duration_seconds"),
	// 	"Time to scrape metrics from starlink dish",
	// 	nil, nil,
	// )
	dishUptimeSeconds = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "uptime_seconds"),
		"Dish running time",
		[]string{"device_id"}, nil,
	)
	// DishOutages
	dishOutage = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "outage_duration"),
		"Starlink Dish Outage Information",
		[]string{"device_id", "start_time", "cause"}, nil,
	)
	dishOutageDidSwitch = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "outage_did_switch"),
		"Starlink Dish Outage Information",
		[]string{"device_id"}, nil,
	)
	dishSoftwareUpdateState = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "software_update_state"),
		"Starlink Dish Software Update State",
		[]string{"device_id", "software_update_state"}, nil,
	)
	dishDisablementCode = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "disablement_code"),
		"Starlink Dish Disablement Code",
		[]string{"device_id", "disablement_code"}, nil,
	)
	// DishGpsStats
	dishGpsValid = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "gps_valid"),
		"GPS Status",
		[]string{"device_id"}, nil,
	)
	dishGpsSats = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "gps_sats"),
		"Number of GPS Sats.",
		[]string{"device_id"}, nil,
	)
	// DishStatus
	dishSecondsToFirstNonemptySlot = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "first_nonempty_slot_seconds"),
		"Seconds to next non empty slot",
		[]string{"device_id"}, nil,
	)
	dishPopPingDropRatio = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "pop_ping_drop_ratio"),
		"Percent of pings dropped",
		[]string{"device_id"}, nil,
	)
	dishDownlinkThroughputBytes = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "downlink_throughput_bytes"),
		"Amount of bandwidth in bytes per second download",
		[]string{"device_id"}, nil,
	)
	dishUplinkThroughputBytes = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "uplink_throughput_bytes"),
		"Amount of bandwidth in bytes per second upload",
		[]string{"device_id"}, nil,
	)
	dishPopPingLatencySeconds = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "pop_ping_latency_seconds"),
		"Latency of connection in seconds",
		[]string{"device_id"}, nil,
	)
	dishStowRequested = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "dish_stow_requested"),
		"stow requested",
		[]string{"device_id"}, nil,
	)
	dishBoreSightAzimuthDeg = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "bore_sight_azimuth_deg"),
		"azimuth in degrees",
		[]string{"device_id"}, nil,
	)
	dishBoreSightElevationDeg = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "bore_sight_elevation_deg"),
		"elevation in degrees",
		[]string{"device_id"}, nil,
	)
	// dishPhyRxBeamSnrAvg = prometheus.NewDesc(
	// 	prometheus.BuildFQName(namespace, "dish", "phy_rx_beam_snr_avg"),
	// 	"physical rx beam snr average",
	// 	nil, nil,
	// )
	// dishTemperateCenter = prometheus.NewDesc(
	// 	prometheus.BuildFQName(namespace, "dish", "tCenter"),
	// 	"Temperature center",
	// 	nil, nil,
	// )
	dishEthSpeedMbps = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "eth_speed"),
		"ethernet speed",
		[]string{"device_id"}, nil,
	)
	dishSnrAboveNoiseFloor = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "snr_above_noise_floor"),
		"SNR is below noise floor (1 = poor signal problem, 0 = good signal)",
		[]string{"device_id"}, nil,
	)
	dishSnrPersistentlyLow = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "snr_persistently_low"),
		"SNR is persistently low (1 = chronic problem, 0 = OK)",
		[]string{"device_id"}, nil,
	)
	// DishAlerts
	dishPowerSupplyThermalThrottle = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "alert_power_supply_thermal_throttle"),
		"Status of power supply thermal throttling",
		[]string{"device_id"}, nil,
	)
	dishIsPowerSaveIdle = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "alert_power_save_idle"),
		"Status of power save idle",
		[]string{"device_id"}, nil,
	)
	// dishMovingWhileNotMobile = prometheus.NewDesc(
	// 	prometheus.BuildFQName(namespace, "dish", "alert_moving_while_not_mobile"),
	// 	"Status of moving while not mobile",
	// 	nil, nil,
	// )
	// dishMovingTooFastForPolicy = prometheus.NewDesc(
	// 	prometheus.BuildFQName(namespace, "dish", "alert_moving_too_fast_for_policy"),
	// 	"Status of moving too fast for policy",
	// 	nil, nil,
	// )
	dishLowMotorCurrent = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "alert_low_motor_current"),
		"Status of low motor current",
		[]string{"device_id"}, nil,
	)
	dishLowerSignalThanPredicted = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "alert_lower_signal_than_predicted"),
		"Status of lower signal than predicted",
		[]string{"device_id"}, nil,
	)
	dishObstructionMapReset = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "alert_obstruction_map_reset"),
		"Status of obstruction map reset",
		[]string{"device_id"}, nil,
	)
	dishAlertRoaming = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "alert_roaming"),
		"Status of roaming",
		[]string{"device_id"}, nil,
	)
	dishAlertMotorsStuck = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "alert_motors_stuck"),
		"Status of motor stuck",
		[]string{"device_id"}, nil,
	)
	dishAlertThermalThrottle = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "alert_thermal_throttle"),
		"Status of thermal throttling",
		[]string{"device_id"}, nil,
	)
	dishAlertThermalShutdown = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "alert_thermal_shutdown"),
		"Status of thermal shutdown",
		[]string{"device_id"}, nil,
	)
	dishAlertMastNotNearVertical = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "alert_mast_not_near_vertical"),
		"Status of mast position",
		[]string{"device_id"}, nil,
	)
	dishUnexpectedLocation = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "alert_unexpected_location"),
		"Status of location",
		[]string{"device_id"}, nil,
	)
	dishSlowEthernetSpeeds = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "alert_slow_eth_speeds"),
		"Status of ethernet",
		[]string{"device_id"}, nil,
	)
	dishInstallPending = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "alert_install_pending"),
		"Installation Pending",
		[]string{"device_id"}, nil,
	)
	dishIsHeating = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "alert_is_heating"),
		"Is Heating",
		[]string{"device_id"}, nil,
	)
	// DishObstructions
	dishPatchesValid = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "patches_valid"),
		"Number of valid patches",
		[]string{"device_id"}, nil,
	)
	dishCurrentlyObstructed = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "currently_obstructed"),
		"Status of view of the sky",
		[]string{"device_id"}, nil,
	)
	dishTimeObstructed = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "time_obstructed"),
		"Time obstructed ratio",
		[]string{"device_id"}, nil,
	)
	dishFractionObstructionRatio = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "fraction_obstruction_ratio"),
		"Percentage of obstruction",
		[]string{"device_id"}, nil,
	)
	dishValidSeconds = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "valid_seconds"),
		"Unknown",
		[]string{"device_id"}, nil,
	)
	dishProlongedObstructionDurationSeconds = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "prolonged_obstruction_duration_seconds"),
		"Average in seconds of prolonged obstructions",
		[]string{"device_id"}, nil,
	)
	dishProlongedObstructionIntervalSeconds = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "prolonged_obstruction_interval_seconds"),
		"Average prolonged obstruction interval in seconds",
		[]string{"device_id"}, nil,
	)
	dishProlongedObstructionValid = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "prolonged_obstruction_valid"),
		"Average prolonged obstruction is valid",
		[]string{"device_id"}, nil,
	)
	dishWedgeFractionObstructionRatio = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "wedge_fraction_obstruction_ratio"),
		"Percentage of obstruction per wedge section",
		[]string{"wedge", "wedge_name"}, nil,
	)
	dishWedgeAbsFractionObstructionRatio = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "wedge_abs_fraction_obstruction_ratio"),
		"Percentage of Absolute fraction per wedge section",
		[]string{"wedge", "wedge_name"}, nil,
	)

	// dishObstructionMap
	dishObstructionMap = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "obstruction_map"),
		"Obstruction Map",
		[]string{
			"device_id",
			"timestamp",
			"num_rows",
			"num_cols",
			// "min_elevation_deg",
			"max_theta_deg",
			"reference_frame",
			"data"}, nil,
	)

	// diagnostics
	dishGpsTimeS = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "gps_time_s"),
		"GPS Time",
		[]string{"device_id"}, nil,
	)
	// TODO:
	// Find a Golang package to convert GPS time to UTC time
	// dishUTCTime = prometheus.NewDesc(
	// 	prometheus.BuildFQName(namespace, "dish", "utc_time"),
	// 	"UTC Time",
	// 	nil, nil,
	// )

	// dishPowerWatt
	dishPowerWatt = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "power_watt_current"),
		"Current Power Usage in Watt",
		[]string{"device_id"}, nil,
	)
	dishPowerWattAvg15min = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "power_watt_avg_15min"),
		"Average Power Usage in Watt over 15 minutes",
		[]string{"device_id"}, nil,
	)
)
