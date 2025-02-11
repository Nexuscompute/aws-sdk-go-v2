// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type AggregationPeriod string

// Enum values for AggregationPeriod
const (
	AggregationPeriodOneHour AggregationPeriod = "OneHour"
	AggregationPeriodOneDay  AggregationPeriod = "OneDay"
	AggregationPeriodOneWeek AggregationPeriod = "OneWeek"
)

// Values returns all known values for AggregationPeriod. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (AggregationPeriod) Values() []AggregationPeriod {
	return []AggregationPeriod{
		"OneHour",
		"OneDay",
		"OneWeek",
	}
}

type ApplicationConfigType string

// Enum values for ApplicationConfigType
const (
	ApplicationConfigTypeSemtechGeoLocation ApplicationConfigType = "SemtechGeolocation"
)

// Values returns all known values for ApplicationConfigType. Note that this can
// be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ApplicationConfigType) Values() []ApplicationConfigType {
	return []ApplicationConfigType{
		"SemtechGeolocation",
	}
}

type BatteryLevel string

// Enum values for BatteryLevel
const (
	BatteryLevelNormal   BatteryLevel = "normal"
	BatteryLevelLow      BatteryLevel = "low"
	BatteryLevelCritical BatteryLevel = "critical"
)

// Values returns all known values for BatteryLevel. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (BatteryLevel) Values() []BatteryLevel {
	return []BatteryLevel{
		"normal",
		"low",
		"critical",
	}
}

type ConnectionStatus string

// Enum values for ConnectionStatus
const (
	ConnectionStatusConnected    ConnectionStatus = "Connected"
	ConnectionStatusDisconnected ConnectionStatus = "Disconnected"
)

// Values returns all known values for ConnectionStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ConnectionStatus) Values() []ConnectionStatus {
	return []ConnectionStatus{
		"Connected",
		"Disconnected",
	}
}

type DeviceProfileType string

// Enum values for DeviceProfileType
const (
	DeviceProfileTypeSidewalk DeviceProfileType = "Sidewalk"
	DeviceProfileTypeLoRaWAN  DeviceProfileType = "LoRaWAN"
)

// Values returns all known values for DeviceProfileType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (DeviceProfileType) Values() []DeviceProfileType {
	return []DeviceProfileType{
		"Sidewalk",
		"LoRaWAN",
	}
}

type DeviceState string

// Enum values for DeviceState
const (
	DeviceStateProvisioned           DeviceState = "Provisioned"
	DeviceStateRegisterednotseen     DeviceState = "RegisteredNotSeen"
	DeviceStateRegisteredreachable   DeviceState = "RegisteredReachable"
	DeviceStateRegisteredunreachable DeviceState = "RegisteredUnreachable"
)

// Values returns all known values for DeviceState. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (DeviceState) Values() []DeviceState {
	return []DeviceState{
		"Provisioned",
		"RegisteredNotSeen",
		"RegisteredReachable",
		"RegisteredUnreachable",
	}
}

type DimensionName string

// Enum values for DimensionName
const (
	DimensionNameDeviceId  DimensionName = "DeviceId"
	DimensionNameGatewayId DimensionName = "GatewayId"
)

// Values returns all known values for DimensionName. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (DimensionName) Values() []DimensionName {
	return []DimensionName{
		"DeviceId",
		"GatewayId",
	}
}

type DlClass string

// Enum values for DlClass
const (
	DlClassClassB DlClass = "ClassB"
	DlClassClassC DlClass = "ClassC"
)

// Values returns all known values for DlClass. Note that this can be expanded in
// the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (DlClass) Values() []DlClass {
	return []DlClass{
		"ClassB",
		"ClassC",
	}
}

type DownlinkMode string

// Enum values for DownlinkMode
const (
	DownlinkModeSequential         DownlinkMode = "SEQUENTIAL"
	DownlinkModeConcurrent         DownlinkMode = "CONCURRENT"
	DownlinkModeUsingUplinkGateway DownlinkMode = "USING_UPLINK_GATEWAY"
)

// Values returns all known values for DownlinkMode. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (DownlinkMode) Values() []DownlinkMode {
	return []DownlinkMode{
		"SEQUENTIAL",
		"CONCURRENT",
		"USING_UPLINK_GATEWAY",
	}
}

type Event string

// Enum values for Event
const (
	EventDiscovered  Event = "discovered"
	EventLost        Event = "lost"
	EventAck         Event = "ack"
	EventNack        Event = "nack"
	EventPassthrough Event = "passthrough"
)

// Values returns all known values for Event. Note that this can be expanded in
// the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (Event) Values() []Event {
	return []Event{
		"discovered",
		"lost",
		"ack",
		"nack",
		"passthrough",
	}
}

type EventNotificationPartnerType string

// Enum values for EventNotificationPartnerType
const (
	EventNotificationPartnerTypeSidewalk EventNotificationPartnerType = "Sidewalk"
)

// Values returns all known values for EventNotificationPartnerType. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (EventNotificationPartnerType) Values() []EventNotificationPartnerType {
	return []EventNotificationPartnerType{
		"Sidewalk",
	}
}

type EventNotificationResourceType string

// Enum values for EventNotificationResourceType
const (
	EventNotificationResourceTypeFuotaTask       EventNotificationResourceType = "FuotaTask"
	EventNotificationResourceTypeSidewalkAccount EventNotificationResourceType = "SidewalkAccount"
	EventNotificationResourceTypeWirelessDevice  EventNotificationResourceType = "WirelessDevice"
	EventNotificationResourceTypeWirelessGateway EventNotificationResourceType = "WirelessGateway"
)

// Values returns all known values for EventNotificationResourceType. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (EventNotificationResourceType) Values() []EventNotificationResourceType {
	return []EventNotificationResourceType{
		"FuotaTask",
		"SidewalkAccount",
		"WirelessDevice",
		"WirelessGateway",
	}
}

type EventNotificationTopicStatus string

// Enum values for EventNotificationTopicStatus
const (
	EventNotificationTopicStatusEnabled  EventNotificationTopicStatus = "Enabled"
	EventNotificationTopicStatusDisabled EventNotificationTopicStatus = "Disabled"
)

// Values returns all known values for EventNotificationTopicStatus. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (EventNotificationTopicStatus) Values() []EventNotificationTopicStatus {
	return []EventNotificationTopicStatus{
		"Enabled",
		"Disabled",
	}
}

type ExpressionType string

// Enum values for ExpressionType
const (
	ExpressionTypeRuleName  ExpressionType = "RuleName"
	ExpressionTypeMqttTopic ExpressionType = "MqttTopic"
)

// Values returns all known values for ExpressionType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ExpressionType) Values() []ExpressionType {
	return []ExpressionType{
		"RuleName",
		"MqttTopic",
	}
}

type FuotaDeviceStatus string

// Enum values for FuotaDeviceStatus
const (
	FuotaDeviceStatusInitial                        FuotaDeviceStatus = "Initial"
	FuotaDeviceStatusPackageNotSupported            FuotaDeviceStatus = "Package_Not_Supported"
	FuotaDeviceStatusFragAlgoUnsupported            FuotaDeviceStatus = "FragAlgo_unsupported"
	FuotaDeviceStatusNotEnoughMemory                FuotaDeviceStatus = "Not_enough_memory"
	FuotaDeviceStatusFragIndexUnsupported           FuotaDeviceStatus = "FragIndex_unsupported"
	FuotaDeviceStatusWrongDescriptor                FuotaDeviceStatus = "Wrong_descriptor"
	FuotaDeviceStatusSessionCntReplay               FuotaDeviceStatus = "SessionCnt_replay"
	FuotaDeviceStatusMissingFrag                    FuotaDeviceStatus = "MissingFrag"
	FuotaDeviceStatusMemoryError                    FuotaDeviceStatus = "MemoryError"
	FuotaDeviceStatusMICError                       FuotaDeviceStatus = "MICError"
	FuotaDeviceStatusSuccessful                     FuotaDeviceStatus = "Successful"
	FuotaDeviceStatusDeviceExistInConflictFuotaTask FuotaDeviceStatus = "Device_exist_in_conflict_fuota_task"
)

// Values returns all known values for FuotaDeviceStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (FuotaDeviceStatus) Values() []FuotaDeviceStatus {
	return []FuotaDeviceStatus{
		"Initial",
		"Package_Not_Supported",
		"FragAlgo_unsupported",
		"Not_enough_memory",
		"FragIndex_unsupported",
		"Wrong_descriptor",
		"SessionCnt_replay",
		"MissingFrag",
		"MemoryError",
		"MICError",
		"Successful",
		"Device_exist_in_conflict_fuota_task",
	}
}

type FuotaTaskEvent string

// Enum values for FuotaTaskEvent
const (
	FuotaTaskEventFuota FuotaTaskEvent = "Fuota"
)

// Values returns all known values for FuotaTaskEvent. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (FuotaTaskEvent) Values() []FuotaTaskEvent {
	return []FuotaTaskEvent{
		"Fuota",
	}
}

type FuotaTaskStatus string

// Enum values for FuotaTaskStatus
const (
	FuotaTaskStatusPending             FuotaTaskStatus = "Pending"
	FuotaTaskStatusFuotaSessionWaiting FuotaTaskStatus = "FuotaSession_Waiting"
	FuotaTaskStatusInFuotaSession      FuotaTaskStatus = "In_FuotaSession"
	FuotaTaskStatusFuotaDone           FuotaTaskStatus = "FuotaDone"
	FuotaTaskStatusDeleteWaiting       FuotaTaskStatus = "Delete_Waiting"
)

// Values returns all known values for FuotaTaskStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (FuotaTaskStatus) Values() []FuotaTaskStatus {
	return []FuotaTaskStatus{
		"Pending",
		"FuotaSession_Waiting",
		"In_FuotaSession",
		"FuotaDone",
		"Delete_Waiting",
	}
}

type FuotaTaskType string

// Enum values for FuotaTaskType
const (
	FuotaTaskTypeLoRaWAN FuotaTaskType = "LoRaWAN"
)

// Values returns all known values for FuotaTaskType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (FuotaTaskType) Values() []FuotaTaskType {
	return []FuotaTaskType{
		"LoRaWAN",
	}
}

type IdentifierType string

// Enum values for IdentifierType
const (
	IdentifierTypePartnerAccountId  IdentifierType = "PartnerAccountId"
	IdentifierTypeDevEui            IdentifierType = "DevEui"
	IdentifierTypeFuotaTaskId       IdentifierType = "FuotaTaskId"
	IdentifierTypeGatewayEui        IdentifierType = "GatewayEui"
	IdentifierTypeWirelessDeviceId  IdentifierType = "WirelessDeviceId"
	IdentifierTypeWirelessGatewayId IdentifierType = "WirelessGatewayId"
)

// Values returns all known values for IdentifierType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (IdentifierType) Values() []IdentifierType {
	return []IdentifierType{
		"PartnerAccountId",
		"DevEui",
		"FuotaTaskId",
		"GatewayEui",
		"WirelessDeviceId",
		"WirelessGatewayId",
	}
}

type ImportTaskStatus string

// Enum values for ImportTaskStatus
const (
	ImportTaskStatusInitializing ImportTaskStatus = "INITIALIZING"
	ImportTaskStatusInitialized  ImportTaskStatus = "INITIALIZED"
	ImportTaskStatusPending      ImportTaskStatus = "PENDING"
	ImportTaskStatusComplete     ImportTaskStatus = "COMPLETE"
	ImportTaskStatusFailed       ImportTaskStatus = "FAILED"
	ImportTaskStatusDeleting     ImportTaskStatus = "DELETING"
)

// Values returns all known values for ImportTaskStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ImportTaskStatus) Values() []ImportTaskStatus {
	return []ImportTaskStatus{
		"INITIALIZING",
		"INITIALIZED",
		"PENDING",
		"COMPLETE",
		"FAILED",
		"DELETING",
	}
}

type LogLevel string

// Enum values for LogLevel
const (
	LogLevelInfo     LogLevel = "INFO"
	LogLevelError    LogLevel = "ERROR"
	LogLevelDisabled LogLevel = "DISABLED"
)

// Values returns all known values for LogLevel. Note that this can be expanded in
// the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (LogLevel) Values() []LogLevel {
	return []LogLevel{
		"INFO",
		"ERROR",
		"DISABLED",
	}
}

type MessageType string

// Enum values for MessageType
const (
	MessageTypeCustomCommandIdNotify MessageType = "CUSTOM_COMMAND_ID_NOTIFY"
	MessageTypeCustomCommandIdGet    MessageType = "CUSTOM_COMMAND_ID_GET"
	MessageTypeCustomCommandIdSet    MessageType = "CUSTOM_COMMAND_ID_SET"
	MessageTypeCustomCommandIdResp   MessageType = "CUSTOM_COMMAND_ID_RESP"
)

// Values returns all known values for MessageType. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (MessageType) Values() []MessageType {
	return []MessageType{
		"CUSTOM_COMMAND_ID_NOTIFY",
		"CUSTOM_COMMAND_ID_GET",
		"CUSTOM_COMMAND_ID_SET",
		"CUSTOM_COMMAND_ID_RESP",
	}
}

type MetricName string

// Enum values for MetricName
const (
	MetricNameDeviceRSSI                     MetricName = "DeviceRSSI"
	MetricNameDeviceSNR                      MetricName = "DeviceSNR"
	MetricNameDeviceRoamingRSSI              MetricName = "DeviceRoamingRSSI"
	MetricNameDeviceRoamingSNR               MetricName = "DeviceRoamingSNR"
	MetricNameDeviceUplinkCount              MetricName = "DeviceUplinkCount"
	MetricNameDeviceDownlinkCount            MetricName = "DeviceDownlinkCount"
	MetricNameDeviceUplinkLostCount          MetricName = "DeviceUplinkLostCount"
	MetricNameDeviceUplinkLostRate           MetricName = "DeviceUplinkLostRate"
	MetricNameDeviceJoinRequestCount         MetricName = "DeviceJoinRequestCount"
	MetricNameDeviceJoinAcceptCount          MetricName = "DeviceJoinAcceptCount"
	MetricNameDeviceRoamingUplinkCount       MetricName = "DeviceRoamingUplinkCount"
	MetricNameDeviceRoamingDownlinkCount     MetricName = "DeviceRoamingDownlinkCount"
	MetricNameGatewayUpTime                  MetricName = "GatewayUpTime"
	MetricNameGatewayDownTime                MetricName = "GatewayDownTime"
	MetricNameGatewayRSSI                    MetricName = "GatewayRSSI"
	MetricNameGatewaySNR                     MetricName = "GatewaySNR"
	MetricNameGatewayUplinkCount             MetricName = "GatewayUplinkCount"
	MetricNameGatewayDownlinkCount           MetricName = "GatewayDownlinkCount"
	MetricNameGatewayJoinRequestCount        MetricName = "GatewayJoinRequestCount"
	MetricNameGatewayJoinAcceptCount         MetricName = "GatewayJoinAcceptCount"
	MetricNameAwsAccountUplinkCount          MetricName = "AwsAccountUplinkCount"
	MetricNameAwsAccountDownlinkCount        MetricName = "AwsAccountDownlinkCount"
	MetricNameAwsAccountUplinkLostCount      MetricName = "AwsAccountUplinkLostCount"
	MetricNameAwsAccountUplinkLostRate       MetricName = "AwsAccountUplinkLostRate"
	MetricNameAwsAccountJoinRequestCount     MetricName = "AwsAccountJoinRequestCount"
	MetricNameAwsAccountJoinAcceptCount      MetricName = "AwsAccountJoinAcceptCount"
	MetricNameAwsAccountRoamingUplinkCount   MetricName = "AwsAccountRoamingUplinkCount"
	MetricNameAwsAccountRoamingDownlinkCount MetricName = "AwsAccountRoamingDownlinkCount"
	MetricNameAwsAccountDeviceCount          MetricName = "AwsAccountDeviceCount"
	MetricNameAwsAccountGatewayCount         MetricName = "AwsAccountGatewayCount"
	MetricNameAwsAccountActiveDeviceCount    MetricName = "AwsAccountActiveDeviceCount"
	MetricNameAwsAccountActiveGatewayCount   MetricName = "AwsAccountActiveGatewayCount"
)

// Values returns all known values for MetricName. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (MetricName) Values() []MetricName {
	return []MetricName{
		"DeviceRSSI",
		"DeviceSNR",
		"DeviceRoamingRSSI",
		"DeviceRoamingSNR",
		"DeviceUplinkCount",
		"DeviceDownlinkCount",
		"DeviceUplinkLostCount",
		"DeviceUplinkLostRate",
		"DeviceJoinRequestCount",
		"DeviceJoinAcceptCount",
		"DeviceRoamingUplinkCount",
		"DeviceRoamingDownlinkCount",
		"GatewayUpTime",
		"GatewayDownTime",
		"GatewayRSSI",
		"GatewaySNR",
		"GatewayUplinkCount",
		"GatewayDownlinkCount",
		"GatewayJoinRequestCount",
		"GatewayJoinAcceptCount",
		"AwsAccountUplinkCount",
		"AwsAccountDownlinkCount",
		"AwsAccountUplinkLostCount",
		"AwsAccountUplinkLostRate",
		"AwsAccountJoinRequestCount",
		"AwsAccountJoinAcceptCount",
		"AwsAccountRoamingUplinkCount",
		"AwsAccountRoamingDownlinkCount",
		"AwsAccountDeviceCount",
		"AwsAccountGatewayCount",
		"AwsAccountActiveDeviceCount",
		"AwsAccountActiveGatewayCount",
	}
}

type MetricQueryStatus string

// Enum values for MetricQueryStatus
const (
	MetricQueryStatusSucceeded MetricQueryStatus = "Succeeded"
	MetricQueryStatusFailed    MetricQueryStatus = "Failed"
)

// Values returns all known values for MetricQueryStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (MetricQueryStatus) Values() []MetricQueryStatus {
	return []MetricQueryStatus{
		"Succeeded",
		"Failed",
	}
}

type MulticastFrameInfo string

// Enum values for MulticastFrameInfo
const (
	MulticastFrameInfoEnabled  MulticastFrameInfo = "ENABLED"
	MulticastFrameInfoDisabled MulticastFrameInfo = "DISABLED"
)

// Values returns all known values for MulticastFrameInfo. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (MulticastFrameInfo) Values() []MulticastFrameInfo {
	return []MulticastFrameInfo{
		"ENABLED",
		"DISABLED",
	}
}

type OnboardStatus string

// Enum values for OnboardStatus
const (
	OnboardStatusInitialized OnboardStatus = "INITIALIZED"
	OnboardStatusPending     OnboardStatus = "PENDING"
	OnboardStatusOnboarded   OnboardStatus = "ONBOARDED"
	OnboardStatusFailed      OnboardStatus = "FAILED"
)

// Values returns all known values for OnboardStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (OnboardStatus) Values() []OnboardStatus {
	return []OnboardStatus{
		"INITIALIZED",
		"PENDING",
		"ONBOARDED",
		"FAILED",
	}
}

type PartnerType string

// Enum values for PartnerType
const (
	PartnerTypeSidewalk PartnerType = "Sidewalk"
)

// Values returns all known values for PartnerType. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (PartnerType) Values() []PartnerType {
	return []PartnerType{
		"Sidewalk",
	}
}

type PositionConfigurationFec string

// Enum values for PositionConfigurationFec
const (
	PositionConfigurationFecRose PositionConfigurationFec = "ROSE"
	PositionConfigurationFecNone PositionConfigurationFec = "NONE"
)

// Values returns all known values for PositionConfigurationFec. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (PositionConfigurationFec) Values() []PositionConfigurationFec {
	return []PositionConfigurationFec{
		"ROSE",
		"NONE",
	}
}

type PositionConfigurationStatus string

// Enum values for PositionConfigurationStatus
const (
	PositionConfigurationStatusEnabled  PositionConfigurationStatus = "Enabled"
	PositionConfigurationStatusDisabled PositionConfigurationStatus = "Disabled"
)

// Values returns all known values for PositionConfigurationStatus. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (PositionConfigurationStatus) Values() []PositionConfigurationStatus {
	return []PositionConfigurationStatus{
		"Enabled",
		"Disabled",
	}
}

type PositioningConfigStatus string

// Enum values for PositioningConfigStatus
const (
	PositioningConfigStatusEnabled  PositioningConfigStatus = "Enabled"
	PositioningConfigStatusDisabled PositioningConfigStatus = "Disabled"
)

// Values returns all known values for PositioningConfigStatus. Note that this can
// be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (PositioningConfigStatus) Values() []PositioningConfigStatus {
	return []PositioningConfigStatus{
		"Enabled",
		"Disabled",
	}
}

type PositionResourceType string

// Enum values for PositionResourceType
const (
	PositionResourceTypeWirelessDevice  PositionResourceType = "WirelessDevice"
	PositionResourceTypeWirelessGateway PositionResourceType = "WirelessGateway"
)

// Values returns all known values for PositionResourceType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (PositionResourceType) Values() []PositionResourceType {
	return []PositionResourceType{
		"WirelessDevice",
		"WirelessGateway",
	}
}

type PositionSolverProvider string

// Enum values for PositionSolverProvider
const (
	PositionSolverProviderSemtech PositionSolverProvider = "Semtech"
)

// Values returns all known values for PositionSolverProvider. Note that this can
// be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (PositionSolverProvider) Values() []PositionSolverProvider {
	return []PositionSolverProvider{
		"Semtech",
	}
}

type PositionSolverType string

// Enum values for PositionSolverType
const (
	PositionSolverTypeGnss PositionSolverType = "GNSS"
)

// Values returns all known values for PositionSolverType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (PositionSolverType) Values() []PositionSolverType {
	return []PositionSolverType{
		"GNSS",
	}
}

type SigningAlg string

// Enum values for SigningAlg
const (
	SigningAlgEd25519 SigningAlg = "Ed25519"
	SigningAlgP256r1  SigningAlg = "P256r1"
)

// Values returns all known values for SigningAlg. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (SigningAlg) Values() []SigningAlg {
	return []SigningAlg{
		"Ed25519",
		"P256r1",
	}
}

type SummaryMetricConfigurationStatus string

// Enum values for SummaryMetricConfigurationStatus
const (
	SummaryMetricConfigurationStatusEnabled  SummaryMetricConfigurationStatus = "Enabled"
	SummaryMetricConfigurationStatusDisabled SummaryMetricConfigurationStatus = "Disabled"
)

// Values returns all known values for SummaryMetricConfigurationStatus. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (SummaryMetricConfigurationStatus) Values() []SummaryMetricConfigurationStatus {
	return []SummaryMetricConfigurationStatus{
		"Enabled",
		"Disabled",
	}
}

type SupportedRfRegion string

// Enum values for SupportedRfRegion
const (
	SupportedRfRegionEu868  SupportedRfRegion = "EU868"
	SupportedRfRegionUs915  SupportedRfRegion = "US915"
	SupportedRfRegionAu915  SupportedRfRegion = "AU915"
	SupportedRfRegionAs9231 SupportedRfRegion = "AS923-1"
	SupportedRfRegionAs9232 SupportedRfRegion = "AS923-2"
	SupportedRfRegionAs9233 SupportedRfRegion = "AS923-3"
	SupportedRfRegionAs9234 SupportedRfRegion = "AS923-4"
	SupportedRfRegionEu433  SupportedRfRegion = "EU433"
	SupportedRfRegionCn470  SupportedRfRegion = "CN470"
	SupportedRfRegionCn779  SupportedRfRegion = "CN779"
	SupportedRfRegionRu864  SupportedRfRegion = "RU864"
	SupportedRfRegionKr920  SupportedRfRegion = "KR920"
	SupportedRfRegionIn865  SupportedRfRegion = "IN865"
)

// Values returns all known values for SupportedRfRegion. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (SupportedRfRegion) Values() []SupportedRfRegion {
	return []SupportedRfRegion{
		"EU868",
		"US915",
		"AU915",
		"AS923-1",
		"AS923-2",
		"AS923-3",
		"AS923-4",
		"EU433",
		"CN470",
		"CN779",
		"RU864",
		"KR920",
		"IN865",
	}
}

type WirelessDeviceEvent string

// Enum values for WirelessDeviceEvent
const (
	WirelessDeviceEventJoin         WirelessDeviceEvent = "Join"
	WirelessDeviceEventRejoin       WirelessDeviceEvent = "Rejoin"
	WirelessDeviceEventUplinkData   WirelessDeviceEvent = "Uplink_Data"
	WirelessDeviceEventDownlinkData WirelessDeviceEvent = "Downlink_Data"
	WirelessDeviceEventRegistration WirelessDeviceEvent = "Registration"
)

// Values returns all known values for WirelessDeviceEvent. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (WirelessDeviceEvent) Values() []WirelessDeviceEvent {
	return []WirelessDeviceEvent{
		"Join",
		"Rejoin",
		"Uplink_Data",
		"Downlink_Data",
		"Registration",
	}
}

type WirelessDeviceFrameInfo string

// Enum values for WirelessDeviceFrameInfo
const (
	WirelessDeviceFrameInfoEnabled  WirelessDeviceFrameInfo = "ENABLED"
	WirelessDeviceFrameInfoDisabled WirelessDeviceFrameInfo = "DISABLED"
)

// Values returns all known values for WirelessDeviceFrameInfo. Note that this can
// be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (WirelessDeviceFrameInfo) Values() []WirelessDeviceFrameInfo {
	return []WirelessDeviceFrameInfo{
		"ENABLED",
		"DISABLED",
	}
}

type WirelessDeviceIdType string

// Enum values for WirelessDeviceIdType
const (
	WirelessDeviceIdTypeWirelessDeviceId        WirelessDeviceIdType = "WirelessDeviceId"
	WirelessDeviceIdTypeDevEui                  WirelessDeviceIdType = "DevEui"
	WirelessDeviceIdTypeThingName               WirelessDeviceIdType = "ThingName"
	WirelessDeviceIdTypeSidewalkManufacturingSn WirelessDeviceIdType = "SidewalkManufacturingSn"
)

// Values returns all known values for WirelessDeviceIdType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (WirelessDeviceIdType) Values() []WirelessDeviceIdType {
	return []WirelessDeviceIdType{
		"WirelessDeviceId",
		"DevEui",
		"ThingName",
		"SidewalkManufacturingSn",
	}
}

type WirelessDeviceSidewalkStatus string

// Enum values for WirelessDeviceSidewalkStatus
const (
	WirelessDeviceSidewalkStatusProvisioned WirelessDeviceSidewalkStatus = "PROVISIONED"
	WirelessDeviceSidewalkStatusRegistered  WirelessDeviceSidewalkStatus = "REGISTERED"
	WirelessDeviceSidewalkStatusActivated   WirelessDeviceSidewalkStatus = "ACTIVATED"
	WirelessDeviceSidewalkStatusUnknown     WirelessDeviceSidewalkStatus = "UNKNOWN"
)

// Values returns all known values for WirelessDeviceSidewalkStatus. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (WirelessDeviceSidewalkStatus) Values() []WirelessDeviceSidewalkStatus {
	return []WirelessDeviceSidewalkStatus{
		"PROVISIONED",
		"REGISTERED",
		"ACTIVATED",
		"UNKNOWN",
	}
}

type WirelessDeviceType string

// Enum values for WirelessDeviceType
const (
	WirelessDeviceTypeSidewalk WirelessDeviceType = "Sidewalk"
	WirelessDeviceTypeLoRaWAN  WirelessDeviceType = "LoRaWAN"
)

// Values returns all known values for WirelessDeviceType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (WirelessDeviceType) Values() []WirelessDeviceType {
	return []WirelessDeviceType{
		"Sidewalk",
		"LoRaWAN",
	}
}

type WirelessGatewayEvent string

// Enum values for WirelessGatewayEvent
const (
	WirelessGatewayEventCupsRequest WirelessGatewayEvent = "CUPS_Request"
	WirelessGatewayEventCertificate WirelessGatewayEvent = "Certificate"
)

// Values returns all known values for WirelessGatewayEvent. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (WirelessGatewayEvent) Values() []WirelessGatewayEvent {
	return []WirelessGatewayEvent{
		"CUPS_Request",
		"Certificate",
	}
}

type WirelessGatewayIdType string

// Enum values for WirelessGatewayIdType
const (
	WirelessGatewayIdTypeGatewayEui        WirelessGatewayIdType = "GatewayEui"
	WirelessGatewayIdTypeWirelessGatewayId WirelessGatewayIdType = "WirelessGatewayId"
	WirelessGatewayIdTypeThingName         WirelessGatewayIdType = "ThingName"
)

// Values returns all known values for WirelessGatewayIdType. Note that this can
// be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (WirelessGatewayIdType) Values() []WirelessGatewayIdType {
	return []WirelessGatewayIdType{
		"GatewayEui",
		"WirelessGatewayId",
		"ThingName",
	}
}

type WirelessGatewayServiceType string

// Enum values for WirelessGatewayServiceType
const (
	WirelessGatewayServiceTypeCups WirelessGatewayServiceType = "CUPS"
	WirelessGatewayServiceTypeLns  WirelessGatewayServiceType = "LNS"
)

// Values returns all known values for WirelessGatewayServiceType. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (WirelessGatewayServiceType) Values() []WirelessGatewayServiceType {
	return []WirelessGatewayServiceType{
		"CUPS",
		"LNS",
	}
}

type WirelessGatewayTaskDefinitionType string

// Enum values for WirelessGatewayTaskDefinitionType
const (
	WirelessGatewayTaskDefinitionTypeUpdate WirelessGatewayTaskDefinitionType = "UPDATE"
)

// Values returns all known values for WirelessGatewayTaskDefinitionType. Note
// that this can be expanded in the future, and so it is only as up to date as the
// client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (WirelessGatewayTaskDefinitionType) Values() []WirelessGatewayTaskDefinitionType {
	return []WirelessGatewayTaskDefinitionType{
		"UPDATE",
	}
}

type WirelessGatewayTaskStatus string

// Enum values for WirelessGatewayTaskStatus
const (
	WirelessGatewayTaskStatusPending     WirelessGatewayTaskStatus = "PENDING"
	WirelessGatewayTaskStatusInProgress  WirelessGatewayTaskStatus = "IN_PROGRESS"
	WirelessGatewayTaskStatusFirstRetry  WirelessGatewayTaskStatus = "FIRST_RETRY"
	WirelessGatewayTaskStatusSecondRetry WirelessGatewayTaskStatus = "SECOND_RETRY"
	WirelessGatewayTaskStatusCompleted   WirelessGatewayTaskStatus = "COMPLETED"
	WirelessGatewayTaskStatusFailed      WirelessGatewayTaskStatus = "FAILED"
)

// Values returns all known values for WirelessGatewayTaskStatus. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (WirelessGatewayTaskStatus) Values() []WirelessGatewayTaskStatus {
	return []WirelessGatewayTaskStatus{
		"PENDING",
		"IN_PROGRESS",
		"FIRST_RETRY",
		"SECOND_RETRY",
		"COMPLETED",
		"FAILED",
	}
}

type WirelessGatewayType string

// Enum values for WirelessGatewayType
const (
	WirelessGatewayTypeLoRaWAN WirelessGatewayType = "LoRaWAN"
)

// Values returns all known values for WirelessGatewayType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (WirelessGatewayType) Values() []WirelessGatewayType {
	return []WirelessGatewayType{
		"LoRaWAN",
	}
}
