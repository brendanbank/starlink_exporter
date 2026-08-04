package exporter

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"

	device "github.com/clarkzjw/starlink-grpc-golang/pkg/spacex.com/api/device"
)

// Exporter collects Starlink stats from the Dish and exports them using
// the prometheus metrics package.
type Exporter struct {
	Conn   *grpc.ClientConn
	Client device.DeviceClient

	DishID      string
	CountryCode string

	// locationDenied is set when the dish refuses GetLocation by policy, so
	// the exporter stops asking on every scrape.
	locationDenied bool
}

// New returns an initialized Exporter.
func New(address string) (*Exporter, error) {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("connect to Starlink dish gRPC interface failed: %s", err.Error())
	}

	defer func() {
		if err != nil {
			if err = conn.Close(); err != nil {
				log.Errorf("Failed to close gRPC client: %s", err.Error())
			}
		}
	}()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	client := device.NewDeviceClient(conn)
	resp, err := client.Handle(ctx, &device.Request{
		Request: &device.Request_GetDeviceInfo{},
	})
	if err != nil {
		return nil, fmt.Errorf("gRPC GetDeviceInfo failed: %s", err.Error())
	}

	deviceInfo := resp.GetGetDeviceInfo().GetDeviceInfo()
	if deviceInfo == nil {
		return nil, fmt.Errorf("gRPC GetDeviceInfo failed: deviceInfo is nil")
	}

	return &Exporter{
		Conn:        conn,
		Client:      client,
		DishID:      deviceInfo.GetId(),
		CountryCode: deviceInfo.GetCountryCode(),
	}, nil
}

// Describe describes all the metrics ever exported by the Starlink exporter.
// It implements prometheus.Collector.
func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {

	// WiFi
	ch <- dishConfig

	// DeviceInfo
	ch <- dishMobilityClass
	ch <- userClassOfService
	ch <- dishReadyState
	ch <- dishInitializationDurationSeconds
	ch <- dishInfo
	ch <- SoftwarePartitionsEqual
	ch <- dishSoftwareUpdateState
	ch <- dishDisablementCode
	ch <- dishIsDev
	ch <- dishBootCount
	ch <- dishAntiRollbackVersion
	ch <- dishIsHit

	// Quaternion
	ch <- dishNed2dishQuaternionQScalar
	ch <- dishNed2dishQuaternionQX
	ch <- dishNed2dishQuaternionQY
	ch <- dishNed2dishQuaternionQZ

	// BootInfo
	ch <- dishBootInfo

	// DeviceState
	ch <- dishUp
	ch <- dishUptimeSeconds
	// ch <- dishScrapeDurationSeconds

	// DishOutage
	ch <- dishOutage
	ch <- dishOutageDidSwitch

	// DishGpsStats
	ch <- dishGpsValid
	ch <- dishGpsSats
	ch <- dishPntFilterConvergenceState
	ch <- dishGpsNoSatsAfterTtff
	ch <- dishGpsInhibited

	// DishLocation
	ch <- dishLatitude
	ch <- dishLongitude
	ch <- dishAltitude

	// DishStatus
	ch <- dishSecondsToFirstNonemptySlot
	ch <- dishPopPingDropRatio
	ch <- dishDownlinkThroughputBytes
	ch <- dishUplinkThroughputBytes
	ch <- dishPopPingLatencySeconds
	ch <- dishStowRequested
	ch <- dishBoreSightAzimuthDeg
	ch <- dishBoreSightElevationDeg
	ch <- dishEthSpeedMbps
	ch <- dishSnrAboveNoiseFloor
	ch <- dishSnrPersistentlyLow
	ch <- dishDlBandwidthRestrictedReason
	ch <- dishUlBandwidthRestrictedReason
	ch <- dishRebootReason
	ch <- dishSwupdateRebootReady
	ch <- dishSoftwareUpdateProgress
	ch <- dishSoftwareUpdateRequiresReboot
	ch <- dishSecondsUntilSwupdateRebootPossible
	ch <- dishIsCellDisabled
	ch <- dishHasActuators
	ch <- dishIsMovingFastPersisted
	ch <- dishHighPowerTestMode
	ch <- dishHasSignedCals
	ch <- dishUserDebugModeEnabled
	ch <- dishMacFlag
	ch <- dishNatFlag
	ch <- dishAccountShard
	ch <- dishConnectedRouters
	ch <- dishDownstreamRouters

	// Power accessories
	ch <- dishPlcReceiving
	ch <- dishPlcStateOfCharge
	ch <- dishPlcBatteryHealth
	ch <- dishPlcThermalThrottleLevel
	ch <- dishPlcSafetyModeActive
	ch <- dishPlcPermanentFailure
	ch <- dishPlcTimeToEmpty
	ch <- dishPlcTimeToFull
	ch <- dishPlcPortPower
	ch <- dishBatteryStateOfCharge
	ch <- dishBatteryIsCharging
	ch <- dishBatteryPowerSource
	ch <- dishUpsuDishPower
	ch <- dishUpsuRouterPower
	ch <- dishUpsuUptimeSeconds
	ch <- dishApsDishPower
	ch <- dishApsUptimeSeconds

	// DishAlerts
	ch <- dishPowerSupplyThermalThrottle
	ch <- dishIsPowerSaveIdle
	ch <- dishLowMotorCurrent
	ch <- dishLowerSignalThanPredicted
	ch <- dishObstructionMapReset
	ch <- dishAlertRoaming
	ch <- dishAlertMotorsStuck
	ch <- dishAlertThermalThrottle
	ch <- dishAlertThermalShutdown
	ch <- dishAlertMastNotNearVertical
	ch <- dishUnexpectedLocation
	ch <- dishSlowEthernetSpeeds
	ch <- dishInstallPending
	ch <- dishIsHeating
	ch <- dishNoEthernetLink
	ch <- dishAlertDbfTelemStale
	ch <- dishAlertDishWaterDetected
	ch <- dishAlertRouterWaterDetected
	ch <- dishAlertUpsuRouterPortSlow
	ch <- dishSlowEthernetSpeeds100

	// DishAlignment
	ch <- dishAlignmentStats
	ch <- dishBoresightAzimuthDiffDeg
	ch <- dishBoresightElevationDiffDeg
	ch <- dishTiltAngleDeg

	// DishObstructions
	ch <- dishPatchesValid
	ch <- dishCurrentlyObstructed
	ch <- dishTimeObstructed
	ch <- dishFractionObstructionRatio
	ch <- dishValidSeconds
	ch <- dishProlongedObstructionDurationSeconds
	ch <- dishProlongedObstructionIntervalSeconds
	ch <- dishProlongedObstructionValid
	ch <- dishObstructionMap

	// DishLocation
	ch <- dishLocationInfo

	// Diagnostics
	ch <- dishGpsTimeS
	ch <- dishHardwareSelfTest
	ch <- dishHardwareSelfTestCode
	ch <- dishStowed
	ch <- dishOverageRateLimited

	// Power
	ch <- dishPowerWatt
	ch <- dishPowerWattAvg15min

	// History
	ch <- dishHistorySamples
	ch <- dishHistoryPopPingDropRateAvg
	ch <- dishHistoryFullDropSeconds
	ch <- dishHistoryPartialDropSeconds
	ch <- dishHistoryLongestFullDropSeconds
	ch <- dishHistoryPopPingLatencySecondsAvg
	ch <- dishHistoryPopPingLatencySecondsMin
	ch <- dishHistoryPopPingLatencySecondsMax
	ch <- dishHistoryDownlinkBpsAvg
	ch <- dishHistoryDownlinkBpsMax
	ch <- dishHistoryUplinkBpsAvg
	ch <- dishHistoryUplinkBpsMax
	ch <- dishHistoryDownlinkBytes
	ch <- dishHistoryUplinkBytes
	ch <- dishHistoryPowerWattMin
	ch <- dishHistoryPowerWattMax
	ch <- dishHistoryOutageCount
	ch <- dishHistoryOutageSeconds
	ch <- dishHistoryOutageMaxSeconds
	ch <- dishHistoryOutageSwitchCount
}

// Collect fetches the stats from Starlink dish and delivers them as Prometheus metrics.
// It implements prometheus.Collector.
func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
	ok := true

	// Status, obstructions, alerts and alignment all come out of the same
	// GetStatus response, so ask the dish for it once per scrape.
	status, err := e.getStatus()
	if err != nil {
		log.Errorf("gRPC GetStatus failed: %s", err.Error())
		ok = false
	} else {
		e.collectDishStatus(ch, status)
		e.collectDishObstructionStatus(ch, status)
		e.collectDishAlerts(ch, status)
		e.collectAlignmentStats(ch, status)
	}

	ok = e.collectDishLocation(ch) && ok
	ok = e.collectDishObstructionMap(ch) && ok
	ok = e.collectDishConfig(ch) && ok
	ok = e.collectDishDiagnostics(ch) && ok
	ok = e.collectDishHistory(ch) && ok

	if ok {
		ch <- prometheus.MustNewConstMetric(
			dishUp, prometheus.GaugeValue, 1.0, e.DishID,
		)
	} else {
		ch <- prometheus.MustNewConstMetric(
			dishUp, prometheus.GaugeValue, 0.0, e.DishID,
		)
	}
}

func (e *Exporter) getStatus() (*device.DishGetStatusResponse, error) {
	req := &device.Request{
		Request: &device.Request_GetStatus{},
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()
	resp, err := e.Client.Handle(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.GetDishGetStatus(), nil
}

func (e *Exporter) collectDishStatus(ch chan<- prometheus.Metric, dishStatus *device.DishGetStatusResponse) {
	dishI := dishStatus.GetDeviceInfo()
	dishB := dishI.GetBoot()
	dishS := dishStatus.GetDeviceState()
	dishG := dishStatus.GetGpsStats()
	dishO := dishStatus.GetOutage()
	dishR := dishStatus.GetReadyStates()
	dishInit := dishStatus.GetInitializationDurationSeconds()
	dishQuaternion := dishStatus.GetNed2DishQuaternion()

	ch <- prometheus.MustNewConstMetric(
		dishNed2dishQuaternionQScalar, prometheus.GaugeValue, float64(dishQuaternion.QScalar), e.DishID)

	ch <- prometheus.MustNewConstMetric(
		dishNed2dishQuaternionQX, prometheus.GaugeValue, float64(dishQuaternion.QX), e.DishID)

	ch <- prometheus.MustNewConstMetric(
		dishNed2dishQuaternionQY, prometheus.GaugeValue, float64(dishQuaternion.QY), e.DishID)

	ch <- prometheus.MustNewConstMetric(
		dishNed2dishQuaternionQZ, prometheus.GaugeValue, float64(dishQuaternion.QZ), e.DishID)

	ch <- prometheus.MustNewConstMetric(
		dishInitializationDurationSeconds, prometheus.GaugeValue, 1.00,
		e.DishID,
		fmt.Sprint(dishInit.GetAttitudeInitialization()),
		fmt.Sprint(dishInit.GetBurstDetected()),
		fmt.Sprint(dishInit.GetEkfConverged()),
		fmt.Sprint(dishInit.GetFirstCplane()),
		fmt.Sprint(dishInit.GetFirstPopPing()),
		fmt.Sprint(dishInit.GetGpsValid()),
		fmt.Sprint(dishInit.GetInitialNetworkEntry()),
		fmt.Sprint(dishInit.GetNetworkSchedule()),
		fmt.Sprint(dishInit.GetRfReady()),
		fmt.Sprint(dishInit.GetStableConnection()),
	)

	ch <- prometheus.MustNewConstMetric(
		dishReadyState, prometheus.GaugeValue, 1.00,
		e.DishID,
		fmt.Sprint(dishR.GetCady()),
		fmt.Sprint(dishR.GetScp()),
		fmt.Sprint(dishR.GetL1L2()),
		fmt.Sprint(dishR.GetXphy()),
		fmt.Sprint(dishR.GetAap()),
		fmt.Sprint(dishR.GetRf()))

	ch <- prometheus.MustNewConstMetric(
		userClassOfService, prometheus.GaugeValue, 1.00,
		e.DishID,
		dishStatus.GetClassOfService().String())
	ch <- prometheus.MustNewConstMetric(
		dishMobilityClass, prometheus.GaugeValue, 1.00,
		e.DishID,
		dishStatus.GetMobilityClass().String())
	ch <- prometheus.MustNewConstMetric(
		dishInfo, prometheus.GaugeValue, 1.00,
		dishI.GetId(),
		dishI.GetBuildId(),
		dishI.GetHardwareVersion(),
		dishI.GetSoftwareVersion(),
		fmt.Sprint(dishI.GetGenerationNumber()),
		dishI.GetCountryCode(),
		fmt.Sprint(dishI.GetBootcount()),
		fmt.Sprint(dishI.GetUtcOffsetS()),
	)
	ch <- prometheus.MustNewConstMetric(
		SoftwarePartitionsEqual, prometheus.GaugeValue, flool(dishI.GetSoftwarePartitionsEqual()), e.DishID,
	)
	ch <- prometheus.MustNewConstMetric(
		dishSoftwareUpdateState, prometheus.GaugeValue, 1.00, e.DishID, dishStatus.GetSoftwareUpdateState().String(),
	)
	ch <- prometheus.MustNewConstMetric(
		dishDisablementCode, prometheus.GaugeValue, 1.00, e.DishID, dishStatus.GetDisablementCode().String(),
	)
	ch <- prometheus.MustNewConstMetric(
		dishIsDev, prometheus.GaugeValue, flool(dishI.GetIsDev()), e.DishID,
	)
	ch <- prometheus.MustNewConstMetric(
		dishBootCount, prometheus.CounterValue, float64(dishI.GetBootcount()), e.DishID,
	)
	ch <- prometheus.MustNewConstMetric(
		dishAntiRollbackVersion, prometheus.CounterValue, float64(dishI.GetAntiRollbackVersion()), e.DishID,
	)
	ch <- prometheus.MustNewConstMetric(
		dishIsHit, prometheus.GaugeValue, flool(dishI.GetIsHitl()), e.DishID,
	)
	ch <- prometheus.MustNewConstMetric(
		dishBootInfo, prometheus.GaugeValue, 1.00,
		e.DishID,
		fmt.Sprint(dishB.GetCountByReason()),
		fmt.Sprint(dishB.GetCountByReasonDelta()),
		fmt.Sprint(dishB.GetLastReason()),
		fmt.Sprint(dishB.GetLastCount()),
	)
	ch <- prometheus.MustNewConstMetric(
		dishUptimeSeconds, prometheus.CounterValue, float64(dishS.GetUptimeS()), e.DishID,
	)
	ch <- prometheus.MustNewConstMetric(
		dishOutage, prometheus.GaugeValue, float64(dishO.GetDurationNs()),
		e.DishID,
		fmt.Sprint(dishO.GetStartTimestampNs()),
		dishO.GetCause().String(),
	)
	ch <- prometheus.MustNewConstMetric(
		dishOutageDidSwitch, prometheus.GaugeValue, flool(dishO.GetDidSwitch()), e.DishID,
	)
	ch <- prometheus.MustNewConstMetric(
		dishGpsValid, prometheus.GaugeValue, flool(dishG.GetGpsValid()), e.DishID,
	)
	ch <- prometheus.MustNewConstMetric(
		dishGpsSats, prometheus.GaugeValue, float64(dishG.GetGpsSats()), e.DishID,
	)
	ch <- prometheus.MustNewConstMetric(
		dishPntFilterConvergenceState, prometheus.GaugeValue, float64(dishG.GetPntFilterConvergenceState()),
		e.DishID,
		dishG.GetPntFilterConvergenceState().String(),
	)
	ch <- prometheus.MustNewConstMetric(
		dishSecondsToFirstNonemptySlot, prometheus.GaugeValue, float64(dishStatus.GetSecondsToFirstNonemptySlot()), e.DishID,
	)
	ch <- prometheus.MustNewConstMetric(
		dishPopPingDropRatio, prometheus.GaugeValue, float64(dishStatus.GetPopPingDropRate()), e.DishID,
	)
	ch <- prometheus.MustNewConstMetric(
		dishDownlinkThroughputBytes, prometheus.GaugeValue, float64(dishStatus.GetDownlinkThroughputBps()), e.DishID,
	)
	ch <- prometheus.MustNewConstMetric(
		dishUplinkThroughputBytes, prometheus.GaugeValue, float64(dishStatus.GetUplinkThroughputBps()), e.DishID,
	)
	ch <- prometheus.MustNewConstMetric(
		dishPopPingLatencySeconds, prometheus.GaugeValue, float64(dishStatus.GetPopPingLatencyMs()/1000), e.DishID,
	)
	ch <- prometheus.MustNewConstMetric(
		dishStowRequested, prometheus.GaugeValue, flool(dishStatus.GetStowRequested()), e.DishID,
	)
	ch <- prometheus.MustNewConstMetric(
		dishBoreSightAzimuthDeg, prometheus.GaugeValue, float64(dishStatus.GetBoresightAzimuthDeg()), e.DishID,
	)
	ch <- prometheus.MustNewConstMetric(
		dishBoreSightElevationDeg, prometheus.GaugeValue, float64(dishStatus.GetBoresightElevationDeg()), e.DishID,
	)
	ch <- prometheus.MustNewConstMetric(
		dishEthSpeedMbps, prometheus.UntypedValue, float64(dishStatus.GetEthSpeedMbps()), e.DishID,
	)
	ch <- prometheus.MustNewConstMetric(
		dishSnrAboveNoiseFloor, prometheus.GaugeValue, flool(!dishStatus.GetIsSnrAboveNoiseFloor()), e.DishID,
	)
	ch <- prometheus.MustNewConstMetric(
		dishSnrPersistentlyLow, prometheus.GaugeValue, flool(dishStatus.GetIsSnrPersistentlyLow()), e.DishID,
	)
	// ch <- prometheus.MustNewConstMetric(
	// 	dishPhyRxBeamSnrAvg, prometheus.GaugeValue, float64(dishStatus.GetPhyRxBeamSnrAvg()),
	// )
	// ch <- prometheus.MustNewConstMetric(
	// 	dishTemperateCenter, prometheus.GaugeValue, float64(dishStatus.GetTCenter()),
	// )

	ch <- prometheus.MustNewConstMetric(
		dishGpsNoSatsAfterTtff, prometheus.GaugeValue, flool(dishG.GetNoSatsAfterTtff()), e.DishID,
	)
	ch <- prometheus.MustNewConstMetric(
		dishGpsInhibited, prometheus.GaugeValue, flool(dishG.GetInhibitGps()), e.DishID,
	)

	// Bandwidth restrictions: the enum value is 0 for NO_LIMIT, so the value is
	// usable on its own and the label says why.
	ch <- prometheus.MustNewConstMetric(
		dishDlBandwidthRestrictedReason, prometheus.GaugeValue, float64(dishStatus.GetDlBandwidthRestrictedReason()),
		e.DishID, dishStatus.GetDlBandwidthRestrictedReason().String(),
	)
	ch <- prometheus.MustNewConstMetric(
		dishUlBandwidthRestrictedReason, prometheus.GaugeValue, float64(dishStatus.GetUlBandwidthRestrictedReason()),
		e.DishID, dishStatus.GetUlBandwidthRestrictedReason().String(),
	)

	swUpdate := dishStatus.GetSoftwareUpdateStats()
	ch <- prometheus.MustNewConstMetric(
		dishRebootReason, prometheus.GaugeValue, float64(dishStatus.GetRebootReason()),
		e.DishID, dishStatus.GetRebootReason().String(),
	)
	ch <- prometheus.MustNewConstMetric(
		dishSwupdateRebootReady, prometheus.GaugeValue, flool(dishStatus.GetSwupdateRebootReady()), e.DishID,
	)
	ch <- prometheus.MustNewConstMetric(
		dishSoftwareUpdateProgress, prometheus.GaugeValue, float64(swUpdate.GetSoftwareUpdateProgress()), e.DishID,
	)
	ch <- prometheus.MustNewConstMetric(
		dishSoftwareUpdateRequiresReboot, prometheus.GaugeValue, flool(swUpdate.GetUpdateRequiresReboot()), e.DishID,
	)
	ch <- prometheus.MustNewConstMetric(
		dishSecondsUntilSwupdateRebootPossible, prometheus.GaugeValue,
		float64(dishStatus.GetSecondsUntilSwupdateRebootPossible()), e.DishID,
	)

	ch <- prometheus.MustNewConstMetric(
		dishIsCellDisabled, prometheus.GaugeValue, flool(dishStatus.GetIsCellDisabled()), e.DishID,
	)
	ch <- prometheus.MustNewConstMetric(
		dishHasActuators, prometheus.GaugeValue, float64(dishStatus.GetHasActuators()),
		e.DishID, dishStatus.GetHasActuators().String(),
	)
	ch <- prometheus.MustNewConstMetric(
		dishIsMovingFastPersisted, prometheus.GaugeValue, flool(dishStatus.GetIsMovingFastPersisted()), e.DishID,
	)
	ch <- prometheus.MustNewConstMetric(
		dishHighPowerTestMode, prometheus.GaugeValue, flool(dishStatus.GetHighPowerTestMode()), e.DishID,
	)
	ch <- prometheus.MustNewConstMetric(
		dishHasSignedCals, prometheus.GaugeValue, flool(dishStatus.GetHasSignedCals()), e.DishID,
	)
	ch <- prometheus.MustNewConstMetric(
		dishUserDebugModeEnabled, prometheus.GaugeValue, flool(dishStatus.GetUserDebugModeEnabled()), e.DishID,
	)
	ch <- prometheus.MustNewConstMetric(
		dishMacFlag, prometheus.GaugeValue, flool(dishStatus.GetMacFlag()), e.DishID,
	)
	ch <- prometheus.MustNewConstMetric(
		dishNatFlag, prometheus.GaugeValue, float64(dishStatus.GetNatFlag()),
		e.DishID, dishStatus.GetNatFlag().String(),
	)
	ch <- prometheus.MustNewConstMetric(
		dishAccountShard, prometheus.GaugeValue, float64(dishStatus.GetAccountShard()),
		e.DishID, dishStatus.GetAccountShard().String(),
	)
	ch <- prometheus.MustNewConstMetric(
		dishConnectedRouters, prometheus.GaugeValue, float64(len(dishStatus.GetConnectedRouters())), e.DishID,
	)
	ch <- prometheus.MustNewConstMetric(
		dishDownstreamRouters, prometheus.GaugeValue, float64(len(dishStatus.GetDownstreamRouters())), e.DishID,
	)

	e.collectDishPowerAccessories(ch, dishStatus)
}

// collectDishPowerAccessories reports the PLC, battery and UPSU/APS telemetry.
// Dishes without those accessories report zeroed structs.
func (e *Exporter) collectDishPowerAccessories(ch chan<- prometheus.Metric, dishStatus *device.DishGetStatusResponse) {
	plc := dishStatus.GetPlcStats()
	ch <- prometheus.MustNewConstMetric(
		dishPlcReceiving, prometheus.GaugeValue, flool(plc.GetReceivingPlc()), e.DishID,
	)
	ch <- prometheus.MustNewConstMetric(
		dishPlcStateOfCharge, prometheus.GaugeValue, float64(plc.GetStateOfCharge()), e.DishID,
	)
	ch <- prometheus.MustNewConstMetric(
		dishPlcBatteryHealth, prometheus.GaugeValue, float64(plc.GetBatteryHealth()), e.DishID,
	)
	ch <- prometheus.MustNewConstMetric(
		dishPlcThermalThrottleLevel, prometheus.GaugeValue, float64(plc.GetThermalThrottleLevel()), e.DishID,
	)
	ch <- prometheus.MustNewConstMetric(
		dishPlcSafetyModeActive, prometheus.GaugeValue, flool(plc.GetSafetyModeActive()), e.DishID,
	)
	ch <- prometheus.MustNewConstMetric(
		dishPlcPermanentFailure, prometheus.GaugeValue, flool(plc.GetPermanentFailure()), e.DishID,
	)
	ch <- prometheus.MustNewConstMetric(
		dishPlcTimeToEmpty, prometheus.GaugeValue, float64(plc.GetAverageTimeToEmpty()), e.DishID,
	)
	ch <- prometheus.MustNewConstMetric(
		dishPlcTimeToFull, prometheus.GaugeValue, float64(plc.GetAverageTimeToFull()), e.DishID,
	)
	for port, stats := range map[string]*device.PLCPortStats{
		"1": plc.GetPort_1Stats(),
		"2": plc.GetPort_2Stats(),
		"3": plc.GetPort_3Stats(),
	} {
		ch <- prometheus.MustNewConstMetric(
			dishPlcPortPower, prometheus.GaugeValue, float64(stats.GetPower()),
			e.DishID, port, stats.GetStatus().String(),
		)
	}

	battery := dishStatus.GetBatteryStats()
	ch <- prometheus.MustNewConstMetric(
		dishBatteryStateOfCharge, prometheus.GaugeValue, float64(battery.GetStateOfCharge()), e.DishID,
	)
	ch <- prometheus.MustNewConstMetric(
		dishBatteryIsCharging, prometheus.GaugeValue, flool(battery.GetIsCharging()), e.DishID,
	)
	ch <- prometheus.MustNewConstMetric(
		dishBatteryPowerSource, prometheus.GaugeValue, float64(battery.GetPowerSource()),
		e.DishID, battery.GetPowerSource().String(),
	)

	upsu := dishStatus.GetUpsuStats()
	ch <- prometheus.MustNewConstMetric(
		dishUpsuDishPower, prometheus.GaugeValue, float64(upsu.GetDishPower()), e.DishID,
	)
	ch <- prometheus.MustNewConstMetric(
		dishUpsuRouterPower, prometheus.GaugeValue, float64(upsu.GetRouterPower()), e.DishID,
	)
	ch <- prometheus.MustNewConstMetric(
		dishUpsuUptimeSeconds, prometheus.GaugeValue, float64(upsu.GetUptime()), e.DishID,
	)

	aps := dishStatus.GetApsStats()
	ch <- prometheus.MustNewConstMetric(
		dishApsDishPower, prometheus.GaugeValue, float64(aps.GetDishPower()), e.DishID,
	)
	ch <- prometheus.MustNewConstMetric(
		dishApsUptimeSeconds, prometheus.GaugeValue, float64(aps.GetUptime()), e.DishID,
	)
}

func (e *Exporter) collectDishConfig(ch chan<- prometheus.Metric) bool {
	req := &device.Request{
		Request: &device.Request_DishGetConfig{},
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()
	resp, err := e.Client.Handle(ctx, req)
	if err != nil {
		log.Errorf("gRPC DishGetConfig failed: %s", err.Error())
		return false
	}

	dishC := resp.GetDishGetConfig()
	ch <- prometheus.MustNewConstMetric(
		dishConfig, prometheus.GaugeValue, 1.00,
		e.DishID,
		dishC.GetDishConfig().GetSnowMeltMode().String(),
		dishC.GetDishConfig().GetLocationRequestMode().String(),
		dishC.GetDishConfig().GetLevelDishMode().String(),
		fmt.Sprint(dishC.GetDishConfig().GetPowerSaveMode()),
	)
	return true
}

func (e *Exporter) collectDishLocation(ch chan<- prometheus.Metric) bool {
	if e.locationDenied {
		return true
	}

	req := &device.Request{
		Request: &device.Request_GetLocation{},
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()

	resp, err := e.Client.Handle(ctx, req)
	if err != nil {
		// Location access is denied by policy in a growing number of regions.
		// That is a permanent answer, so say so once and stop asking.
		if status.Code(err) == codes.PermissionDenied {
			log.Warnf("gRPC GetLocation denied, no location metrics will be exported: %s", err.Error())
			e.locationDenied = true
			return true
		}
		log.Errorf("gRPC GetLocation failed: %s", err.Error())
		// don't return false since location service might not be enabled
		return true
	}

	locationInfo := resp.GetGetLocation()

	locationSource := locationInfo.GetSource().String()
	sigmaM := locationInfo.GetSigmaM()
	horizontalSpeedMps := locationInfo.GetHorizontalSpeedMps()
	verticalSpeedMps := locationInfo.GetVerticalSpeedMps()

	lla := locationInfo.GetLla()
	lat := lla.GetLat()
	lon := lla.GetLon()
	alt := lla.GetAlt()

	ch <- prometheus.MustNewConstMetric(
		dishLocationInfo, prometheus.GaugeValue, 1.00,
		e.DishID,
		locationSource,
		fmt.Sprintf("%.6f", lat),
		fmt.Sprintf("%.6f", lon),
		fmt.Sprintf("%.3f", alt),
		fmt.Sprintf("%.6f", sigmaM),
		fmt.Sprintf("%.6f", horizontalSpeedMps),
		fmt.Sprintf("%.6f", verticalSpeedMps),
	)

	ch <- prometheus.MustNewConstMetric(
		dishLatitude, prometheus.GaugeValue, lat, e.DishID,
	)
	ch <- prometheus.MustNewConstMetric(
		dishLongitude, prometheus.GaugeValue, lon, e.DishID,
	)
	ch <- prometheus.MustNewConstMetric(
		dishAltitude, prometheus.GaugeValue, float64(alt), e.DishID,
	)

	return true
}

func (e *Exporter) collectDishObstructionStatus(ch chan<- prometheus.Metric, dishStatus *device.DishGetStatusResponse) {
	obstructions := dishStatus.GetObstructionStats()

	ch <- prometheus.MustNewConstMetric(
		dishCurrentlyObstructed, prometheus.GaugeValue, flool(obstructions.GetCurrentlyObstructed()), e.DishID,
	)
	ch <- prometheus.MustNewConstMetric(
		dishFractionObstructionRatio, prometheus.GaugeValue, float64(obstructions.GetFractionObstructed()), e.DishID,
	)
	ch <- prometheus.MustNewConstMetric(
		dishTimeObstructed, prometheus.GaugeValue, float64(obstructions.GetTimeObstructed()), e.DishID,
	)
	ch <- prometheus.MustNewConstMetric(
		dishValidSeconds, prometheus.CounterValue, float64(obstructions.GetValidS()), e.DishID,
	)
	ch <- prometheus.MustNewConstMetric(
		dishPatchesValid, prometheus.GaugeValue, float64(obstructions.GetPatchesValid()), e.DishID,
	)
	ch <- prometheus.MustNewConstMetric(
		dishProlongedObstructionDurationSeconds, prometheus.GaugeValue, float64(obstructions.GetAvgProlongedObstructionDurationS()), e.DishID,
	)
	ch <- prometheus.MustNewConstMetric(
		dishProlongedObstructionIntervalSeconds, prometheus.GaugeValue, float64(obstructions.GetAvgProlongedObstructionIntervalS()), e.DishID,
	)
	ch <- prometheus.MustNewConstMetric(
		dishProlongedObstructionValid, prometheus.GaugeValue, flool(obstructions.GetAvgProlongedObstructionValid()), e.DishID,
	)
}

func (e *Exporter) collectDishObstructionMap(ch chan<- prometheus.Metric) bool {
	req := &device.Request{
		Request: &device.Request_DishGetObstructionMap{},
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()
	resp, err := e.Client.Handle(ctx, req)
	if err != nil {
		log.Errorf("gRPC GetStatus failed: %s", err.Error())
		return false
	}

	obstructionMap := resp.GetDishGetObstructionMap()

	rows := int(obstructionMap.NumRows)
	cols := int(obstructionMap.NumCols)
	referenceFrame := obstructionMap.GetMapReferenceFrame().String()
	data := obstructionMap.Snr

	upLeft := image.Point{0, 0}
	lowRight := image.Point{cols, rows}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	for x := 0; x < cols; x++ {
		for y := 0; y < rows; y++ {
			snr := data[y*cols+x]
			if snr > 1 {
				// shouldn't happen
				snr = 1.0
			}
			if snr == -1 {
				// background
				img.Set(x, y, color.Black)
			} else if snr > 0 {
				// use the same image color style as in starlink-grpc-tools
				// https://github.com/sparky8512/starlink-grpc-tools/blob/a3860e0a73d0b2280eed92eb8a2a97de0ea5fe43/dish_obstruction_map.py#L59-L87
				r := 255
				g := snr * 255
				b := snr * 255
				alpha := 255
				img.Set(x, y, color.RGBA{uint8(r), uint8(g), uint8(b), uint8(alpha)})
			}
		}
	}

	// Encode the image to PNG format in a buffer
	var buf bytes.Buffer
	if err := png.Encode(&buf, img); err != nil {
		fmt.Printf("Failed to encode image: %s", err.Error())
	}

	b64 := base64.StdEncoding.EncodeToString(buf.Bytes())

	ch <- prometheus.MustNewConstMetric(
		dishObstructionMap, prometheus.GaugeValue, float64(time.Now().Unix()),
		e.DishID,
		fmt.Sprint(obstructionMap.GetNumRows()),
		fmt.Sprint(obstructionMap.GetNumCols()),
		fmt.Sprint(obstructionMap.GetMaxThetaDeg()),
		referenceFrame,
		fmt.Sprintf("data:image/png;base64,%s", b64),
	)

	return true
}

func (e *Exporter) collectDishDiagnostics(ch chan<- prometheus.Metric) bool {
	req := &device.Request{
		Request: &device.Request_GetDiagnostics{},
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()
	resp, err := e.Client.Handle(ctx, req)
	if err != nil {
		log.Errorf("gRPC GetDiagnostics failed: %s", err.Error())
		return false
	}

	diagnostics := resp.GetDishGetDiagnostics()

	ch <- prometheus.MustNewConstMetric(
		dishGpsTimeS, prometheus.GaugeValue,
		float64(diagnostics.GetLocation().GetGpsTimeS()), e.DishID,
	)
	ch <- prometheus.MustNewConstMetric(
		dishHardwareSelfTest, prometheus.GaugeValue, float64(diagnostics.GetHardwareSelfTest()),
		e.DishID, diagnostics.GetHardwareSelfTest().String(),
	)
	for _, code := range diagnostics.GetHardwareSelfTestCodes() {
		ch <- prometheus.MustNewConstMetric(
			dishHardwareSelfTestCode, prometheus.GaugeValue, float64(code),
			e.DishID, code.String(),
		)
	}
	ch <- prometheus.MustNewConstMetric(
		dishStowed, prometheus.GaugeValue, flool(diagnostics.GetStowed()), e.DishID,
	)
	ch <- prometheus.MustNewConstMetric(
		dishOverageRateLimited, prometheus.GaugeValue, flool(diagnostics.GetOverageRateLimited()), e.DishID,
	)

	return true
}

func (e *Exporter) collectDishAlerts(ch chan<- prometheus.Metric, dishStatus *device.DishGetStatusResponse) {
	alerts := dishStatus.GetAlerts()

	ch <- prometheus.MustNewConstMetric(
		dishAlertMotorsStuck, prometheus.GaugeValue, flool(alerts.GetMotorsStuck()), e.DishID,
	)
	ch <- prometheus.MustNewConstMetric(
		dishPowerSupplyThermalThrottle, prometheus.GaugeValue, flool(alerts.GetPowerSupplyThermalThrottle()), e.DishID,
	)
	ch <- prometheus.MustNewConstMetric(
		dishIsPowerSaveIdle, prometheus.GaugeValue, flool(alerts.GetIsPowerSaveIdle()), e.DishID,
	)
	// ch <- prometheus.MustNewConstMetric(
	// 	dishMovingWhileNotMobile, prometheus.GaugeValue, flool(alerts.GetMovingWhileNotMobile()),
	// )
	// ch <- prometheus.MustNewConstMetric(
	// 	dishMovingTooFastForPolicy, prometheus.GaugeValue, flool(alerts.GetMovingTooFastForPolicy()),
	// )
	ch <- prometheus.MustNewConstMetric(
		dishLowMotorCurrent, prometheus.GaugeValue, flool(alerts.GetLowMotorCurrent()), e.DishID,
	)
	ch <- prometheus.MustNewConstMetric(
		dishLowerSignalThanPredicted, prometheus.GaugeValue, flool(alerts.GetLowerSignalThanPredicted()), e.DishID,
	)
	ch <- prometheus.MustNewConstMetric(
		dishObstructionMapReset, prometheus.GaugeValue, flool(alerts.GetObstructionMapReset()), e.DishID,
	)
	ch <- prometheus.MustNewConstMetric(
		dishAlertThermalThrottle, prometheus.GaugeValue, flool(alerts.GetThermalThrottle()), e.DishID,
	)
	ch <- prometheus.MustNewConstMetric(
		dishAlertThermalShutdown, prometheus.GaugeValue, flool(alerts.GetThermalShutdown()), e.DishID,
	)
	ch <- prometheus.MustNewConstMetric(
		dishAlertMastNotNearVertical, prometheus.GaugeValue, flool(alerts.GetMastNotNearVertical()), e.DishID,
	)
	ch <- prometheus.MustNewConstMetric(
		dishUnexpectedLocation, prometheus.GaugeValue, flool(alerts.GetUnexpectedLocation()), e.DishID,
	)
	ch <- prometheus.MustNewConstMetric(
		dishSlowEthernetSpeeds, prometheus.GaugeValue, flool(alerts.GetSlowEthernetSpeeds()), e.DishID,
	)
	ch <- prometheus.MustNewConstMetric(
		dishAlertRoaming, prometheus.GaugeValue, flool(alerts.GetRoaming()), e.DishID,
	)
	ch <- prometheus.MustNewConstMetric(
		dishInstallPending, prometheus.GaugeValue, flool(alerts.GetInstallPending()), e.DishID,
	)
	ch <- prometheus.MustNewConstMetric(
		dishIsHeating, prometheus.GaugeValue, flool(alerts.GetIsHeating()), e.DishID,
	)
	ch <- prometheus.MustNewConstMetric(
		dishNoEthernetLink, prometheus.GaugeValue, flool(alerts.GetNoEthernetLink()), e.DishID,
	)
	ch <- prometheus.MustNewConstMetric(
		dishAlertDbfTelemStale, prometheus.GaugeValue, flool(alerts.GetDbfTelemStale()), e.DishID,
	)
	ch <- prometheus.MustNewConstMetric(
		dishAlertDishWaterDetected, prometheus.GaugeValue, flool(alerts.GetDishWaterDetected()), e.DishID,
	)
	ch <- prometheus.MustNewConstMetric(
		dishAlertRouterWaterDetected, prometheus.GaugeValue, flool(alerts.GetRouterWaterDetected()), e.DishID,
	)
	ch <- prometheus.MustNewConstMetric(
		dishAlertUpsuRouterPortSlow, prometheus.GaugeValue, flool(alerts.GetUpsuRouterPortSlow()), e.DishID,
	)
	ch <- prometheus.MustNewConstMetric(
		dishSlowEthernetSpeeds100, prometheus.GaugeValue, flool(alerts.GetSlowEthernetSpeeds_100()), e.DishID,
	)
}

func (e *Exporter) collectAlignmentStats(ch chan<- prometheus.Metric, dishStatus *device.DishGetStatusResponse) {
	alignmentStats := dishStatus.GetAlignmentStats()

	ch <- prometheus.MustNewConstMetric(
		dishAlignmentStats, prometheus.GaugeValue, 1.00,
		e.DishID,
		fmt.Sprint(alignmentStats.GetHasActuators()),
		fmt.Sprint(alignmentStats.GetActuatorState()),
		fmt.Sprint(alignmentStats.GetTiltAngleDeg()),
		fmt.Sprint(alignmentStats.GetBoresightAzimuthDeg()),
		fmt.Sprint(alignmentStats.GetBoresightElevationDeg()),
		fmt.Sprint(alignmentStats.GetAttitudeEstimationState()),
		fmt.Sprint(alignmentStats.GetAttitudeUncertaintyDeg()),
		fmt.Sprint(alignmentStats.GetDesiredBoresightAzimuthDeg()),
		fmt.Sprint(alignmentStats.GetDesiredBoresightElevationDeg()),
	)

	// Calculate difference between desired and actual boresight angles
	azimuthDiff := alignmentStats.GetDesiredBoresightAzimuthDeg() - alignmentStats.GetBoresightAzimuthDeg()
	elevationDiff := alignmentStats.GetDesiredBoresightElevationDeg() - alignmentStats.GetBoresightElevationDeg()

	ch <- prometheus.MustNewConstMetric(
		dishBoresightAzimuthDiffDeg, prometheus.GaugeValue, float64(azimuthDiff), e.DishID,
	)
	ch <- prometheus.MustNewConstMetric(
		dishBoresightElevationDiffDeg, prometheus.GaugeValue, float64(elevationDiff), e.DishID,
	)
	ch <- prometheus.MustNewConstMetric(
		dishTiltAngleDeg, prometheus.GaugeValue, float64(alignmentStats.GetTiltAngleDeg()), e.DishID,
	)
}

// collectDishHistory aggregates the per second ring buffers the dish keeps for
// the last ~15 minutes. Scraping only reads the instantaneous status values, so
// without this everything that happens between two scrapes is invisible.
func (e *Exporter) collectDishHistory(ch chan<- prometheus.Metric) bool {
	req := &device.Request{
		Request: &device.Request_GetHistory{},
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	resp, err := e.Client.Handle(ctx, req)
	if err != nil {
		log.Errorf("gRPC GetHistory failed: %s", err.Error())
		return false
	}
	history := resp.GetDishGetHistory()
	powerHistory := history.GetPowerIn()

	sampleRange, _, _ := computeSampleRange(history, -1)
	if len(sampleRange) == 0 {
		log.Warn("gRPC GetHistory returned no samples")
		return true
	}

	latestRange, _, _ := computeSampleRange(history, 1)
	if len(latestRange) > 0 && len(powerHistory) > 0 {
		ch <- prometheus.MustNewConstMetric(
			dishPowerWatt, prometheus.GaugeValue, float64(powerHistory[latestRange[0]]), e.DishID,
		)
	}

	ch <- prometheus.MustNewConstMetric(
		dishHistorySamples, prometheus.GaugeValue, float64(len(sampleRange)), e.DishID,
	)

	dropRate := history.GetPopPingDropRate()
	latency := history.GetPopPingLatencyMs()
	downlink := history.GetDownlinkThroughputBps()
	uplink := history.GetUplinkThroughputBps()

	var (
		dropSum, dlSum, ulSum, powerSum    float64
		latencySum                         float64
		latencySamples                     int
		fullDrop, partialDrop              int
		currentDropRun, longestDropRun     int
		dlMax, ulMax, latencyMax, powerMax float64
		latencyMin, powerMin               = math.Inf(1), math.Inf(1)
		samples                            = len(sampleRange)
	)

	for _, i := range sampleRange {
		if i < len(dropRate) {
			drop := float64(dropRate[i])
			dropSum += drop
			switch {
			case drop >= 1:
				fullDrop++
				currentDropRun++
				if currentDropRun > longestDropRun {
					longestDropRun = currentDropRun
				}
			case drop > 0:
				partialDrop++
				currentDropRun = 0
			default:
				currentDropRun = 0
			}

			// Latency during a fully dropped second is not meaningful.
			if drop < 1 && i < len(latency) {
				l := float64(latency[i])
				latencySum += l
				latencySamples++
				if l > latencyMax {
					latencyMax = l
				}
				if l < latencyMin {
					latencyMin = l
				}
			}
		}
		if i < len(downlink) {
			d := float64(downlink[i])
			dlSum += d
			if d > dlMax {
				dlMax = d
			}
		}
		if i < len(uplink) {
			u := float64(uplink[i])
			ulSum += u
			if u > ulMax {
				ulMax = u
			}
		}
		if i < len(powerHistory) {
			p := float64(powerHistory[i])
			powerSum += p
			if p > powerMax {
				powerMax = p
			}
			if p < powerMin {
				powerMin = p
			}
		}
	}

	if math.IsInf(latencyMin, 1) {
		latencyMin = 0
	}
	if math.IsInf(powerMin, 1) {
		powerMin = 0
	}
	latencyAvg := 0.0
	if latencySamples > 0 {
		latencyAvg = latencySum / float64(latencySamples)
	}

	ch <- prometheus.MustNewConstMetric(
		dishHistoryPopPingDropRateAvg, prometheus.GaugeValue, dropSum/float64(samples), e.DishID,
	)
	ch <- prometheus.MustNewConstMetric(
		dishHistoryFullDropSeconds, prometheus.GaugeValue, float64(fullDrop), e.DishID,
	)
	ch <- prometheus.MustNewConstMetric(
		dishHistoryPartialDropSeconds, prometheus.GaugeValue, float64(partialDrop), e.DishID,
	)
	ch <- prometheus.MustNewConstMetric(
		dishHistoryLongestFullDropSeconds, prometheus.GaugeValue, float64(longestDropRun), e.DishID,
	)
	ch <- prometheus.MustNewConstMetric(
		dishHistoryPopPingLatencySecondsAvg, prometheus.GaugeValue, latencyAvg/1000, e.DishID,
	)
	ch <- prometheus.MustNewConstMetric(
		dishHistoryPopPingLatencySecondsMin, prometheus.GaugeValue, latencyMin/1000, e.DishID,
	)
	ch <- prometheus.MustNewConstMetric(
		dishHistoryPopPingLatencySecondsMax, prometheus.GaugeValue, latencyMax/1000, e.DishID,
	)
	ch <- prometheus.MustNewConstMetric(
		dishHistoryDownlinkBpsAvg, prometheus.GaugeValue, dlSum/float64(samples), e.DishID,
	)
	ch <- prometheus.MustNewConstMetric(
		dishHistoryDownlinkBpsMax, prometheus.GaugeValue, dlMax, e.DishID,
	)
	ch <- prometheus.MustNewConstMetric(
		dishHistoryUplinkBpsAvg, prometheus.GaugeValue, ulSum/float64(samples), e.DishID,
	)
	ch <- prometheus.MustNewConstMetric(
		dishHistoryUplinkBpsMax, prometheus.GaugeValue, ulMax, e.DishID,
	)
	// Samples are one second apart, so the sum of the rates is the volume.
	ch <- prometheus.MustNewConstMetric(
		dishHistoryDownlinkBytes, prometheus.GaugeValue, dlSum/8, e.DishID,
	)
	ch <- prometheus.MustNewConstMetric(
		dishHistoryUplinkBytes, prometheus.GaugeValue, ulSum/8, e.DishID,
	)
	ch <- prometheus.MustNewConstMetric(
		dishPowerWattAvg15min, prometheus.GaugeValue, powerSum/float64(samples), e.DishID,
	)
	ch <- prometheus.MustNewConstMetric(
		dishHistoryPowerWattMin, prometheus.GaugeValue, powerMin, e.DishID,
	)
	ch <- prometheus.MustNewConstMetric(
		dishHistoryPowerWattMax, prometheus.GaugeValue, powerMax, e.DishID,
	)

	e.collectHistoryOutages(ch, history)

	return true
}

func (e *Exporter) collectHistoryOutages(ch chan<- prometheus.Metric, history *device.DishGetHistoryResponse) {
	count := map[string]int{}
	duration := map[string]float64{}
	maxDuration := 0.0
	switched := 0

	for _, outage := range history.GetOutages() {
		cause := outage.GetCause().String()
		seconds := float64(outage.GetDurationNs()) / float64(time.Second)
		count[cause]++
		duration[cause] += seconds
		if seconds > maxDuration {
			maxDuration = seconds
		}
		if outage.GetDidSwitch() {
			switched++
		}
	}

	for cause, n := range count {
		ch <- prometheus.MustNewConstMetric(
			dishHistoryOutageCount, prometheus.GaugeValue, float64(n), e.DishID, cause,
		)
		ch <- prometheus.MustNewConstMetric(
			dishHistoryOutageSeconds, prometheus.GaugeValue, duration[cause], e.DishID, cause,
		)
	}
	ch <- prometheus.MustNewConstMetric(
		dishHistoryOutageMaxSeconds, prometheus.GaugeValue, maxDuration, e.DishID,
	)
	ch <- prometheus.MustNewConstMetric(
		dishHistoryOutageSwitchCount, prometheus.GaugeValue, float64(switched), e.DishID,
	)
}

func flool(b bool) float64 {
	if b {
		return 1.00
	}
	return 0.00
}

// https://github.com/sparky8512/starlink-grpc-tools/blob/a3860e0a73d0b2280eed92eb8a2a97de0ea5fe43/starlink_grpc.py#L1038-L1090
func computeSampleRange(history *device.DishGetHistoryResponse, parseSamples int) ([]int, int, int) {
	current := int(history.Current)
	samples := len(history.PopPingDropRate)
	if samples == 0 {
		return []int{}, 0, 0
	}

	// Adjust parseSamples if needed
	if parseSamples < 0 || samples < parseSamples {
		parseSamples = samples
	}

	// A dish that booted less than a full buffer ago has not filled the ring
	// yet; asking for more samples than it has produced would wrap into
	// negative offsets.
	if uint64(parseSamples) > history.GetCurrent() {
		parseSamples = int(history.GetCurrent())
	}
	if parseSamples <= 0 {
		return []int{}, 0, current
	}

	// Calculate start position
	start := current - parseSamples

	if start == current {
		return []int{}, 0, current
	}

	// Calculate ring buffer offsets
	endOffset := current % samples
	startOffset := start % samples

	// Create a slice to hold the range of sample indices
	var sampleRange []int

	// Set the range for the requested set of samples
	if startOffset < endOffset {
		// Continuous range
		for i := startOffset; i < endOffset; i++ {
			sampleRange = append(sampleRange, i)
		}
	} else {
		// Wrap-around range
		for i := startOffset; i < samples; i++ {
			sampleRange = append(sampleRange, i)
		}
		for i := 0; i < endOffset; i++ {
			sampleRange = append(sampleRange, i)
		}
	}

	return sampleRange, current - start, current
}
