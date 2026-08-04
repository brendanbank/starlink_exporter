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
	dishPntFilterConvergenceState = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "pnt_filter_convergence_state"),
		"Convergence state of the position/navigation/timing filter",
		[]string{"device_id", "pnt_filter_convergence_state"}, nil,
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
	dishNoEthernetLink = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "alert_no_ethernet_link"),
		"No ethernet link detected",
		[]string{"device_id"}, nil,
	)
	dishAlertDbfTelemStale = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "alert_dbf_telem_stale"),
		"Digital beamforming telemetry is stale",
		[]string{"device_id"}, nil,
	)
	dishAlertDishWaterDetected = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "alert_dish_water_detected"),
		"Water ingress detected in the dish",
		[]string{"device_id"}, nil,
	)
	dishAlertRouterWaterDetected = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "alert_router_water_detected"),
		"Water ingress detected in the router",
		[]string{"device_id"}, nil,
	)
	dishAlertUpsuRouterPortSlow = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "alert_upsu_router_port_slow"),
		"UPSU router port is running below the expected link speed",
		[]string{"device_id"}, nil,
	)
	dishSlowEthernetSpeeds100 = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "alert_slow_eth_speeds_100"),
		"Ethernet link negotiated at 100Mbps or below",
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
	// dishObstructionMap
	// The map itself rides in the data label as a base64 PNG. The collection
	// time is the metric's value rather than a label: a label would mint a new
	// series, and a new copy of that image, on every single scrape.
	dishObstructionMap = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "obstruction_map"),
		"Obstruction map as a base64 PNG in the data label, valued with the unix time it was collected",
		[]string{
			"device_id",
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

	// Bandwidth restrictions
	dishDlBandwidthRestrictedReason = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "dl_bandwidth_restricted_reason"),
		"Reason the downlink bandwidth is being restricted (1 = NO_LIMIT, 2 = POLICY_LIMIT, 5 = OVERAGE_LIMIT)",
		[]string{"device_id", "reason"}, nil,
	)
	dishUlBandwidthRestrictedReason = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "ul_bandwidth_restricted_reason"),
		"Reason the uplink bandwidth is being restricted (1 = NO_LIMIT, 2 = POLICY_LIMIT, 5 = OVERAGE_LIMIT)",
		[]string{"device_id", "reason"}, nil,
	)

	// Software update
	dishRebootReason = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "reboot_reason"),
		"Reason for the pending or last reboot",
		[]string{"device_id", "reboot_reason"}, nil,
	)
	dishSwupdateRebootReady = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "swupdate_reboot_ready"),
		"A software update is staged and the dish is ready to reboot",
		[]string{"device_id"}, nil,
	)
	dishSoftwareUpdateProgress = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "software_update_progress"),
		"Progress of the running software update (0-1)",
		[]string{"device_id"}, nil,
	)
	dishSoftwareUpdateRequiresReboot = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "software_update_requires_reboot"),
		"The staged software update requires a reboot",
		[]string{"device_id"}, nil,
	)
	dishSecondsUntilSwupdateRebootPossible = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "seconds_until_swupdate_reboot_possible"),
		"Seconds until the dish is allowed to reboot for a software update (-1 = unknown)",
		[]string{"device_id"}, nil,
	)

	// Misc device state
	dishIsCellDisabled = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "is_cell_disabled"),
		"Service is disabled in the cell the dish is located in",
		[]string{"device_id"}, nil,
	)
	dishHasActuators = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "has_actuators"),
		"Whether the dish has motorized actuators",
		[]string{"device_id", "has_actuators"}, nil,
	)
	dishIsMovingFastPersisted = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "is_moving_fast_persisted"),
		"Dish has persistently been moving faster than its service plan allows",
		[]string{"device_id"}, nil,
	)
	dishHighPowerTestMode = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "high_power_test_mode"),
		"Dish is in high power test mode",
		[]string{"device_id"}, nil,
	)
	dishHasSignedCals = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "has_signed_cals"),
		"Dish has signed calibration data",
		[]string{"device_id"}, nil,
	)
	dishUserDebugModeEnabled = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "user_debug_mode_enabled"),
		"User debug mode is enabled on the dish",
		[]string{"device_id"}, nil,
	)
	dishMacFlag = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "mac_flag"),
		"Dish MAC flag",
		[]string{"device_id"}, nil,
	)
	dishNatFlag = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "nat_flag"),
		"NAT state reported by the dish",
		[]string{"device_id", "nat_flag"}, nil,
	)
	dishAccountShard = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "account_shard"),
		"Account shard the dish is served by",
		[]string{"device_id", "account_shard"}, nil,
	)
	dishConnectedRouters = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "connected_routers"),
		"Number of routers connected to the dish",
		[]string{"device_id"}, nil,
	)
	dishDownstreamRouters = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "downstream_routers"),
		"Number of downstream routers known to the dish",
		[]string{"device_id"}, nil,
	)
	dishGpsNoSatsAfterTtff = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "gps_no_sats_after_ttff"),
		"No GPS satellites seen after time to first fix",
		[]string{"device_id"}, nil,
	)
	dishGpsInhibited = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "gps_inhibited"),
		"GPS is inhibited by configuration",
		[]string{"device_id"}, nil,
	)

	// PLC (Starlink battery / power line accessory)
	dishPlcReceiving = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "plc_receiving"),
		"Dish is receiving PLC telemetry",
		[]string{"device_id"}, nil,
	)
	dishPlcStateOfCharge = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "plc_state_of_charge"),
		"PLC battery state of charge",
		[]string{"device_id"}, nil,
	)
	dishPlcBatteryHealth = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "plc_battery_health"),
		"PLC battery health",
		[]string{"device_id"}, nil,
	)
	dishPlcThermalThrottleLevel = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "plc_thermal_throttle_level"),
		"PLC thermal throttle level",
		[]string{"device_id"}, nil,
	)
	dishPlcSafetyModeActive = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "plc_safety_mode_active"),
		"PLC safety mode is active",
		[]string{"device_id"}, nil,
	)
	dishPlcPermanentFailure = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "plc_permanent_failure"),
		"PLC reported a permanent failure",
		[]string{"device_id"}, nil,
	)
	dishPlcTimeToEmpty = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "plc_average_time_to_empty"),
		"PLC average time to empty as reported by the accessory",
		[]string{"device_id"}, nil,
	)
	dishPlcTimeToFull = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "plc_average_time_to_full"),
		"PLC average time to full as reported by the accessory",
		[]string{"device_id"}, nil,
	)
	dishPlcPortPower = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "plc_port_power"),
		"Power delivered on a PLC port",
		[]string{"device_id", "port", "status"}, nil,
	)

	// Battery
	dishBatteryStateOfCharge = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "battery_state_of_charge"),
		"Dish battery state of charge",
		[]string{"device_id"}, nil,
	)
	dishBatteryIsCharging = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "battery_is_charging"),
		"Dish battery is charging",
		[]string{"device_id"}, nil,
	)
	dishBatteryPowerSource = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "battery_power_source"),
		"Power source the dish is running from",
		[]string{"device_id", "power_source"}, nil,
	)

	// UPSU / APS power supplies
	dishUpsuDishPower = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "upsu_dish_power_watt"),
		"Dish power draw reported by the UPSU",
		[]string{"device_id"}, nil,
	)
	dishUpsuRouterPower = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "upsu_router_power_watt"),
		"Router power draw reported by the UPSU",
		[]string{"device_id"}, nil,
	)
	dishUpsuUptimeSeconds = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "upsu_uptime_seconds"),
		"UPSU uptime",
		[]string{"device_id"}, nil,
	)
	dishApsDishPower = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "aps_dish_power_watt"),
		"Dish power draw reported by the APS",
		[]string{"device_id"}, nil,
	)
	dishApsUptimeSeconds = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "aps_uptime_seconds"),
		"APS uptime",
		[]string{"device_id"}, nil,
	)

	// Diagnostics
	dishHardwareSelfTest = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "hardware_self_test"),
		"Result of the dish hardware self test",
		[]string{"device_id", "result"}, nil,
	)
	dishHardwareSelfTestCode = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "hardware_self_test_code"),
		"Hardware self test failure code reported by the dish",
		[]string{"device_id", "code"}, nil,
	)
	dishStowed = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "stowed"),
		"Dish is stowed",
		[]string{"device_id"}, nil,
	)
	dishOverageRateLimited = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "overage_rate_limited"),
		"Dish is rate limited because the data allowance was exceeded",
		[]string{"device_id"}, nil,
	)

	// History (aggregates over the ~15 minute sample window returned by GetHistory)
	dishHistorySamples = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "history_samples"),
		"Number of one second samples aggregated from the dish history buffer",
		[]string{"device_id"}, nil,
	)
	dishHistoryPopPingDropRateAvg = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "history_pop_ping_drop_rate_avg"),
		"Mean ping drop rate over the history window",
		[]string{"device_id"}, nil,
	)
	dishHistoryFullDropSeconds = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "history_full_drop_seconds"),
		"Number of seconds in the history window with a 100% ping drop rate",
		[]string{"device_id"}, nil,
	)
	dishHistoryPartialDropSeconds = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "history_partial_drop_seconds"),
		"Number of seconds in the history window with a partial ping drop rate",
		[]string{"device_id"}, nil,
	)
	dishHistoryLongestFullDropSeconds = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "history_longest_full_drop_seconds"),
		"Longest run of consecutive seconds with a 100% ping drop rate in the history window",
		[]string{"device_id"}, nil,
	)
	dishHistoryPopPingLatencySecondsAvg = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "history_pop_ping_latency_seconds_avg"),
		"Mean pop ping latency over the history window, ignoring fully dropped seconds",
		[]string{"device_id"}, nil,
	)
	dishHistoryPopPingLatencySecondsMin = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "history_pop_ping_latency_seconds_min"),
		"Minimum pop ping latency over the history window, ignoring fully dropped seconds",
		[]string{"device_id"}, nil,
	)
	dishHistoryPopPingLatencySecondsMax = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "history_pop_ping_latency_seconds_max"),
		"Maximum pop ping latency over the history window, ignoring fully dropped seconds",
		[]string{"device_id"}, nil,
	)
	dishHistoryDownlinkBpsAvg = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "history_downlink_throughput_bps_avg"),
		"Mean downlink throughput over the history window",
		[]string{"device_id"}, nil,
	)
	dishHistoryDownlinkBpsMax = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "history_downlink_throughput_bps_max"),
		"Peak downlink throughput over the history window",
		[]string{"device_id"}, nil,
	)
	dishHistoryUplinkBpsAvg = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "history_uplink_throughput_bps_avg"),
		"Mean uplink throughput over the history window",
		[]string{"device_id"}, nil,
	)
	dishHistoryUplinkBpsMax = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "history_uplink_throughput_bps_max"),
		"Peak uplink throughput over the history window",
		[]string{"device_id"}, nil,
	)
	dishHistoryDownlinkBytes = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "history_downlink_bytes"),
		"Bytes downloaded during the history window",
		[]string{"device_id"}, nil,
	)
	dishHistoryUplinkBytes = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "history_uplink_bytes"),
		"Bytes uploaded during the history window",
		[]string{"device_id"}, nil,
	)
	dishHistoryPowerWattMin = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "history_power_watt_min"),
		"Minimum power draw over the history window",
		[]string{"device_id"}, nil,
	)
	dishHistoryPowerWattMax = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "history_power_watt_max"),
		"Peak power draw over the history window",
		[]string{"device_id"}, nil,
	)
	dishHistoryOutageCount = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "history_outage_count"),
		"Number of outages in the history window, by cause",
		[]string{"device_id", "cause"}, nil,
	)
	dishHistoryOutageSeconds = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "history_outage_seconds"),
		"Total outage duration in the history window, by cause",
		[]string{"device_id", "cause"}, nil,
	)
	dishHistoryOutageMaxSeconds = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "history_outage_max_seconds"),
		"Longest single outage in the history window",
		[]string{"device_id"}, nil,
	)
	dishHistoryOutageSwitchCount = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "history_outage_switch_count"),
		"Number of outages in the history window caused by a satellite or beam switch",
		[]string{"device_id"}, nil,
	)
)
