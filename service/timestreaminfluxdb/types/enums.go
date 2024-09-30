// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type DbInstanceType string

// Enum values for DbInstanceType
const (
	DbInstanceTypeDbInfluxMedium   DbInstanceType = "db.influx.medium"
	DbInstanceTypeDbInfluxLarge    DbInstanceType = "db.influx.large"
	DbInstanceTypeDbInfluxXlarge   DbInstanceType = "db.influx.xlarge"
	DbInstanceTypeDbInflux2xlarge  DbInstanceType = "db.influx.2xlarge"
	DbInstanceTypeDbInflux4xlarge  DbInstanceType = "db.influx.4xlarge"
	DbInstanceTypeDbInflux8xlarge  DbInstanceType = "db.influx.8xlarge"
	DbInstanceTypeDbInflux12xlarge DbInstanceType = "db.influx.12xlarge"
	DbInstanceTypeDbInflux16xlarge DbInstanceType = "db.influx.16xlarge"
)

// Values returns all known values for DbInstanceType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (DbInstanceType) Values() []DbInstanceType {
	return []DbInstanceType{
		"db.influx.medium",
		"db.influx.large",
		"db.influx.xlarge",
		"db.influx.2xlarge",
		"db.influx.4xlarge",
		"db.influx.8xlarge",
		"db.influx.12xlarge",
		"db.influx.16xlarge",
	}
}

type DbStorageType string

// Enum values for DbStorageType
const (
	DbStorageTypeInfluxIoIncludedT1 DbStorageType = "InfluxIOIncludedT1"
	DbStorageTypeInfluxIoIncludedT2 DbStorageType = "InfluxIOIncludedT2"
	DbStorageTypeInfluxIoIncludedT3 DbStorageType = "InfluxIOIncludedT3"
)

// Values returns all known values for DbStorageType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (DbStorageType) Values() []DbStorageType {
	return []DbStorageType{
		"InfluxIOIncludedT1",
		"InfluxIOIncludedT2",
		"InfluxIOIncludedT3",
	}
}

type DeploymentType string

// Enum values for DeploymentType
const (
	DeploymentTypeSingleAz           DeploymentType = "SINGLE_AZ"
	DeploymentTypeWithMultiazStandby DeploymentType = "WITH_MULTIAZ_STANDBY"
)

// Values returns all known values for DeploymentType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (DeploymentType) Values() []DeploymentType {
	return []DeploymentType{
		"SINGLE_AZ",
		"WITH_MULTIAZ_STANDBY",
	}
}

type DurationType string

// Enum values for DurationType
const (
	DurationTypeHours        DurationType = "hours"
	DurationTypeMinutes      DurationType = "minutes"
	DurationTypeSeconds      DurationType = "seconds"
	DurationTypeMilliseconds DurationType = "milliseconds"
)

// Values returns all known values for DurationType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (DurationType) Values() []DurationType {
	return []DurationType{
		"hours",
		"minutes",
		"seconds",
		"milliseconds",
	}
}

type LogLevel string

// Enum values for LogLevel
const (
	LogLevelDebug LogLevel = "debug"
	LogLevelInfo  LogLevel = "info"
	LogLevelError LogLevel = "error"
)

// Values returns all known values for LogLevel. Note that this can be expanded in
// the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (LogLevel) Values() []LogLevel {
	return []LogLevel{
		"debug",
		"info",
		"error",
	}
}

type Status string

// Enum values for Status
const (
	StatusCreating               Status = "CREATING"
	StatusAvailable              Status = "AVAILABLE"
	StatusDeleting               Status = "DELETING"
	StatusModifying              Status = "MODIFYING"
	StatusUpdating               Status = "UPDATING"
	StatusDeleted                Status = "DELETED"
	StatusFailed                 Status = "FAILED"
	StatusUpdatingDeploymentType Status = "UPDATING_DEPLOYMENT_TYPE"
	StatusUpdatingInstanceType   Status = "UPDATING_INSTANCE_TYPE"
)

// Values returns all known values for Status. Note that this can be expanded in
// the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (Status) Values() []Status {
	return []Status{
		"CREATING",
		"AVAILABLE",
		"DELETING",
		"MODIFYING",
		"UPDATING",
		"DELETED",
		"FAILED",
		"UPDATING_DEPLOYMENT_TYPE",
		"UPDATING_INSTANCE_TYPE",
	}
}

type TracingType string

// Enum values for TracingType
const (
	TracingTypeLog    TracingType = "log"
	TracingTypeJaeger TracingType = "jaeger"
)

// Values returns all known values for TracingType. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (TracingType) Values() []TracingType {
	return []TracingType{
		"log",
		"jaeger",
	}
}

type ValidationExceptionReason string

// Enum values for ValidationExceptionReason
const (
	ValidationExceptionReasonFieldValidationFailed ValidationExceptionReason = "FIELD_VALIDATION_FAILED"
	ValidationExceptionReasonOther                 ValidationExceptionReason = "OTHER"
)

// Values returns all known values for ValidationExceptionReason. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ValidationExceptionReason) Values() []ValidationExceptionReason {
	return []ValidationExceptionReason{
		"FIELD_VALIDATION_FAILED",
		"OTHER",
	}
}
