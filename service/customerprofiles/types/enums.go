// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type AttributeDimensionType string

// Enum values for AttributeDimensionType
const (
	AttributeDimensionTypeInclusive          AttributeDimensionType = "INCLUSIVE"
	AttributeDimensionTypeExclusive          AttributeDimensionType = "EXCLUSIVE"
	AttributeDimensionTypeContains           AttributeDimensionType = "CONTAINS"
	AttributeDimensionTypeBeginsWith         AttributeDimensionType = "BEGINS_WITH"
	AttributeDimensionTypeEndsWith           AttributeDimensionType = "ENDS_WITH"
	AttributeDimensionTypeBefore             AttributeDimensionType = "BEFORE"
	AttributeDimensionTypeAfter              AttributeDimensionType = "AFTER"
	AttributeDimensionTypeBetween            AttributeDimensionType = "BETWEEN"
	AttributeDimensionTypeNotBetween         AttributeDimensionType = "NOT_BETWEEN"
	AttributeDimensionTypeOn                 AttributeDimensionType = "ON"
	AttributeDimensionTypeGreaterThan        AttributeDimensionType = "GREATER_THAN"
	AttributeDimensionTypeLessThan           AttributeDimensionType = "LESS_THAN"
	AttributeDimensionTypeGreaterThanOrEqual AttributeDimensionType = "GREATER_THAN_OR_EQUAL"
	AttributeDimensionTypeLessThanOrEqual    AttributeDimensionType = "LESS_THAN_OR_EQUAL"
	AttributeDimensionTypeEqual              AttributeDimensionType = "EQUAL"
)

// Values returns all known values for AttributeDimensionType. Note that this can
// be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (AttributeDimensionType) Values() []AttributeDimensionType {
	return []AttributeDimensionType{
		"INCLUSIVE",
		"EXCLUSIVE",
		"CONTAINS",
		"BEGINS_WITH",
		"ENDS_WITH",
		"BEFORE",
		"AFTER",
		"BETWEEN",
		"NOT_BETWEEN",
		"ON",
		"GREATER_THAN",
		"LESS_THAN",
		"GREATER_THAN_OR_EQUAL",
		"LESS_THAN_OR_EQUAL",
		"EQUAL",
	}
}

type AttributeMatchingModel string

// Enum values for AttributeMatchingModel
const (
	AttributeMatchingModelOneToOne   AttributeMatchingModel = "ONE_TO_ONE"
	AttributeMatchingModelManyToMany AttributeMatchingModel = "MANY_TO_MANY"
)

// Values returns all known values for AttributeMatchingModel. Note that this can
// be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (AttributeMatchingModel) Values() []AttributeMatchingModel {
	return []AttributeMatchingModel{
		"ONE_TO_ONE",
		"MANY_TO_MANY",
	}
}

type ConflictResolvingModel string

// Enum values for ConflictResolvingModel
const (
	ConflictResolvingModelRecency ConflictResolvingModel = "RECENCY"
	ConflictResolvingModelSource  ConflictResolvingModel = "SOURCE"
)

// Values returns all known values for ConflictResolvingModel. Note that this can
// be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ConflictResolvingModel) Values() []ConflictResolvingModel {
	return []ConflictResolvingModel{
		"RECENCY",
		"SOURCE",
	}
}

type DataFormat string

// Enum values for DataFormat
const (
	DataFormatCsv   DataFormat = "CSV"
	DataFormatJsonl DataFormat = "JSONL"
	DataFormatOrc   DataFormat = "ORC"
)

// Values returns all known values for DataFormat. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (DataFormat) Values() []DataFormat {
	return []DataFormat{
		"CSV",
		"JSONL",
		"ORC",
	}
}

type DataPullMode string

// Enum values for DataPullMode
const (
	DataPullModeIncremental DataPullMode = "Incremental"
	DataPullModeComplete    DataPullMode = "Complete"
)

// Values returns all known values for DataPullMode. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (DataPullMode) Values() []DataPullMode {
	return []DataPullMode{
		"Incremental",
		"Complete",
	}
}

type DateDimensionType string

// Enum values for DateDimensionType
const (
	DateDimensionTypeBefore     DateDimensionType = "BEFORE"
	DateDimensionTypeAfter      DateDimensionType = "AFTER"
	DateDimensionTypeBetween    DateDimensionType = "BETWEEN"
	DateDimensionTypeNotBetween DateDimensionType = "NOT_BETWEEN"
	DateDimensionTypeOn         DateDimensionType = "ON"
)

// Values returns all known values for DateDimensionType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (DateDimensionType) Values() []DateDimensionType {
	return []DateDimensionType{
		"BEFORE",
		"AFTER",
		"BETWEEN",
		"NOT_BETWEEN",
		"ON",
	}
}

type EstimateStatus string

// Enum values for EstimateStatus
const (
	EstimateStatusRunning   EstimateStatus = "RUNNING"
	EstimateStatusSucceeded EstimateStatus = "SUCCEEDED"
	EstimateStatusFailed    EstimateStatus = "FAILED"
)

// Values returns all known values for EstimateStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (EstimateStatus) Values() []EstimateStatus {
	return []EstimateStatus{
		"RUNNING",
		"SUCCEEDED",
		"FAILED",
	}
}

type EventStreamDestinationStatus string

// Enum values for EventStreamDestinationStatus
const (
	EventStreamDestinationStatusHealthy   EventStreamDestinationStatus = "HEALTHY"
	EventStreamDestinationStatusUnhealthy EventStreamDestinationStatus = "UNHEALTHY"
)

// Values returns all known values for EventStreamDestinationStatus. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (EventStreamDestinationStatus) Values() []EventStreamDestinationStatus {
	return []EventStreamDestinationStatus{
		"HEALTHY",
		"UNHEALTHY",
	}
}

type EventStreamState string

// Enum values for EventStreamState
const (
	EventStreamStateRunning EventStreamState = "RUNNING"
	EventStreamStateStopped EventStreamState = "STOPPED"
)

// Values returns all known values for EventStreamState. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (EventStreamState) Values() []EventStreamState {
	return []EventStreamState{
		"RUNNING",
		"STOPPED",
	}
}

type FieldContentType string

// Enum values for FieldContentType
const (
	FieldContentTypeString       FieldContentType = "STRING"
	FieldContentTypeNumber       FieldContentType = "NUMBER"
	FieldContentTypePhoneNumber  FieldContentType = "PHONE_NUMBER"
	FieldContentTypeEmailAddress FieldContentType = "EMAIL_ADDRESS"
	FieldContentTypeName         FieldContentType = "NAME"
)

// Values returns all known values for FieldContentType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (FieldContentType) Values() []FieldContentType {
	return []FieldContentType{
		"STRING",
		"NUMBER",
		"PHONE_NUMBER",
		"EMAIL_ADDRESS",
		"NAME",
	}
}

type FilterDimensionType string

// Enum values for FilterDimensionType
const (
	FilterDimensionTypeInclusive          FilterDimensionType = "INCLUSIVE"
	FilterDimensionTypeExclusive          FilterDimensionType = "EXCLUSIVE"
	FilterDimensionTypeContains           FilterDimensionType = "CONTAINS"
	FilterDimensionTypeBeginsWith         FilterDimensionType = "BEGINS_WITH"
	FilterDimensionTypeEndsWith           FilterDimensionType = "ENDS_WITH"
	FilterDimensionTypeBefore             FilterDimensionType = "BEFORE"
	FilterDimensionTypeAfter              FilterDimensionType = "AFTER"
	FilterDimensionTypeBetween            FilterDimensionType = "BETWEEN"
	FilterDimensionTypeNotBetween         FilterDimensionType = "NOT_BETWEEN"
	FilterDimensionTypeOn                 FilterDimensionType = "ON"
	FilterDimensionTypeGreaterThan        FilterDimensionType = "GREATER_THAN"
	FilterDimensionTypeLessThan           FilterDimensionType = "LESS_THAN"
	FilterDimensionTypeGreaterThanOrEqual FilterDimensionType = "GREATER_THAN_OR_EQUAL"
	FilterDimensionTypeLessThanOrEqual    FilterDimensionType = "LESS_THAN_OR_EQUAL"
	FilterDimensionTypeEqual              FilterDimensionType = "EQUAL"
)

// Values returns all known values for FilterDimensionType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (FilterDimensionType) Values() []FilterDimensionType {
	return []FilterDimensionType{
		"INCLUSIVE",
		"EXCLUSIVE",
		"CONTAINS",
		"BEGINS_WITH",
		"ENDS_WITH",
		"BEFORE",
		"AFTER",
		"BETWEEN",
		"NOT_BETWEEN",
		"ON",
		"GREATER_THAN",
		"LESS_THAN",
		"GREATER_THAN_OR_EQUAL",
		"LESS_THAN_OR_EQUAL",
		"EQUAL",
	}
}

type Gender string

// Enum values for Gender
const (
	GenderMale        Gender = "MALE"
	GenderFemale      Gender = "FEMALE"
	GenderUnspecified Gender = "UNSPECIFIED"
)

// Values returns all known values for Gender. Note that this can be expanded in
// the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (Gender) Values() []Gender {
	return []Gender{
		"MALE",
		"FEMALE",
		"UNSPECIFIED",
	}
}

type IdentityResolutionJobStatus string

// Enum values for IdentityResolutionJobStatus
const (
	IdentityResolutionJobStatusPending        IdentityResolutionJobStatus = "PENDING"
	IdentityResolutionJobStatusPreprocessing  IdentityResolutionJobStatus = "PREPROCESSING"
	IdentityResolutionJobStatusFindMatching   IdentityResolutionJobStatus = "FIND_MATCHING"
	IdentityResolutionJobStatusMerging        IdentityResolutionJobStatus = "MERGING"
	IdentityResolutionJobStatusCompleted      IdentityResolutionJobStatus = "COMPLETED"
	IdentityResolutionJobStatusPartialSuccess IdentityResolutionJobStatus = "PARTIAL_SUCCESS"
	IdentityResolutionJobStatusFailed         IdentityResolutionJobStatus = "FAILED"
)

// Values returns all known values for IdentityResolutionJobStatus. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (IdentityResolutionJobStatus) Values() []IdentityResolutionJobStatus {
	return []IdentityResolutionJobStatus{
		"PENDING",
		"PREPROCESSING",
		"FIND_MATCHING",
		"MERGING",
		"COMPLETED",
		"PARTIAL_SUCCESS",
		"FAILED",
	}
}

type Include string

// Enum values for Include
const (
	IncludeAll  Include = "ALL"
	IncludeAny  Include = "ANY"
	IncludeNone Include = "NONE"
)

// Values returns all known values for Include. Note that this can be expanded in
// the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (Include) Values() []Include {
	return []Include{
		"ALL",
		"ANY",
		"NONE",
	}
}

type IncludeOptions string

// Enum values for IncludeOptions
const (
	IncludeOptionsAll  IncludeOptions = "ALL"
	IncludeOptionsAny  IncludeOptions = "ANY"
	IncludeOptionsNone IncludeOptions = "NONE"
)

// Values returns all known values for IncludeOptions. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (IncludeOptions) Values() []IncludeOptions {
	return []IncludeOptions{
		"ALL",
		"ANY",
		"NONE",
	}
}

type JobScheduleDayOfTheWeek string

// Enum values for JobScheduleDayOfTheWeek
const (
	JobScheduleDayOfTheWeekSunday    JobScheduleDayOfTheWeek = "SUNDAY"
	JobScheduleDayOfTheWeekMonday    JobScheduleDayOfTheWeek = "MONDAY"
	JobScheduleDayOfTheWeekTuesday   JobScheduleDayOfTheWeek = "TUESDAY"
	JobScheduleDayOfTheWeekWednesday JobScheduleDayOfTheWeek = "WEDNESDAY"
	JobScheduleDayOfTheWeekThursday  JobScheduleDayOfTheWeek = "THURSDAY"
	JobScheduleDayOfTheWeekFriday    JobScheduleDayOfTheWeek = "FRIDAY"
	JobScheduleDayOfTheWeekSaturday  JobScheduleDayOfTheWeek = "SATURDAY"
)

// Values returns all known values for JobScheduleDayOfTheWeek. Note that this can
// be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (JobScheduleDayOfTheWeek) Values() []JobScheduleDayOfTheWeek {
	return []JobScheduleDayOfTheWeek{
		"SUNDAY",
		"MONDAY",
		"TUESDAY",
		"WEDNESDAY",
		"THURSDAY",
		"FRIDAY",
		"SATURDAY",
	}
}

type LogicalOperator string

// Enum values for LogicalOperator
const (
	LogicalOperatorAnd LogicalOperator = "AND"
	LogicalOperatorOr  LogicalOperator = "OR"
)

// Values returns all known values for LogicalOperator. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (LogicalOperator) Values() []LogicalOperator {
	return []LogicalOperator{
		"AND",
		"OR",
	}
}

type MarketoConnectorOperator string

// Enum values for MarketoConnectorOperator
const (
	MarketoConnectorOperatorProjection          MarketoConnectorOperator = "PROJECTION"
	MarketoConnectorOperatorLessThan            MarketoConnectorOperator = "LESS_THAN"
	MarketoConnectorOperatorGreaterThan         MarketoConnectorOperator = "GREATER_THAN"
	MarketoConnectorOperatorBetween             MarketoConnectorOperator = "BETWEEN"
	MarketoConnectorOperatorAddition            MarketoConnectorOperator = "ADDITION"
	MarketoConnectorOperatorMultiplication      MarketoConnectorOperator = "MULTIPLICATION"
	MarketoConnectorOperatorDivision            MarketoConnectorOperator = "DIVISION"
	MarketoConnectorOperatorSubtraction         MarketoConnectorOperator = "SUBTRACTION"
	MarketoConnectorOperatorMaskAll             MarketoConnectorOperator = "MASK_ALL"
	MarketoConnectorOperatorMaskFirstN          MarketoConnectorOperator = "MASK_FIRST_N"
	MarketoConnectorOperatorMaskLastN           MarketoConnectorOperator = "MASK_LAST_N"
	MarketoConnectorOperatorValidateNonNull     MarketoConnectorOperator = "VALIDATE_NON_NULL"
	MarketoConnectorOperatorValidateNonZero     MarketoConnectorOperator = "VALIDATE_NON_ZERO"
	MarketoConnectorOperatorValidateNonNegative MarketoConnectorOperator = "VALIDATE_NON_NEGATIVE"
	MarketoConnectorOperatorValidateNumeric     MarketoConnectorOperator = "VALIDATE_NUMERIC"
	MarketoConnectorOperatorNoOp                MarketoConnectorOperator = "NO_OP"
)

// Values returns all known values for MarketoConnectorOperator. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (MarketoConnectorOperator) Values() []MarketoConnectorOperator {
	return []MarketoConnectorOperator{
		"PROJECTION",
		"LESS_THAN",
		"GREATER_THAN",
		"BETWEEN",
		"ADDITION",
		"MULTIPLICATION",
		"DIVISION",
		"SUBTRACTION",
		"MASK_ALL",
		"MASK_FIRST_N",
		"MASK_LAST_N",
		"VALIDATE_NON_NULL",
		"VALIDATE_NON_ZERO",
		"VALIDATE_NON_NEGATIVE",
		"VALIDATE_NUMERIC",
		"NO_OP",
	}
}

type MatchType string

// Enum values for MatchType
const (
	MatchTypeRuleBasedMatching MatchType = "RULE_BASED_MATCHING"
	MatchTypeMlBasedMatching   MatchType = "ML_BASED_MATCHING"
)

// Values returns all known values for MatchType. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (MatchType) Values() []MatchType {
	return []MatchType{
		"RULE_BASED_MATCHING",
		"ML_BASED_MATCHING",
	}
}

type Operator string

// Enum values for Operator
const (
	OperatorEqualTo     Operator = "EQUAL_TO"
	OperatorGreaterThan Operator = "GREATER_THAN"
	OperatorLessThan    Operator = "LESS_THAN"
	OperatorNotEqualTo  Operator = "NOT_EQUAL_TO"
)

// Values returns all known values for Operator. Note that this can be expanded in
// the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (Operator) Values() []Operator {
	return []Operator{
		"EQUAL_TO",
		"GREATER_THAN",
		"LESS_THAN",
		"NOT_EQUAL_TO",
	}
}

type OperatorPropertiesKeys string

// Enum values for OperatorPropertiesKeys
const (
	OperatorPropertiesKeysValue                    OperatorPropertiesKeys = "VALUE"
	OperatorPropertiesKeysValues                   OperatorPropertiesKeys = "VALUES"
	OperatorPropertiesKeysDataType                 OperatorPropertiesKeys = "DATA_TYPE"
	OperatorPropertiesKeysUpperBound               OperatorPropertiesKeys = "UPPER_BOUND"
	OperatorPropertiesKeysLowerBound               OperatorPropertiesKeys = "LOWER_BOUND"
	OperatorPropertiesKeysSourceDataType           OperatorPropertiesKeys = "SOURCE_DATA_TYPE"
	OperatorPropertiesKeysDestinationDataType      OperatorPropertiesKeys = "DESTINATION_DATA_TYPE"
	OperatorPropertiesKeysValidationAction         OperatorPropertiesKeys = "VALIDATION_ACTION"
	OperatorPropertiesKeysMaskValue                OperatorPropertiesKeys = "MASK_VALUE"
	OperatorPropertiesKeysMaskLength               OperatorPropertiesKeys = "MASK_LENGTH"
	OperatorPropertiesKeysTruncateLength           OperatorPropertiesKeys = "TRUNCATE_LENGTH"
	OperatorPropertiesKeysMathOperationFieldsOrder OperatorPropertiesKeys = "MATH_OPERATION_FIELDS_ORDER"
	OperatorPropertiesKeysConcatFormat             OperatorPropertiesKeys = "CONCAT_FORMAT"
	OperatorPropertiesKeysSubfieldCategoryMap      OperatorPropertiesKeys = "SUBFIELD_CATEGORY_MAP"
)

// Values returns all known values for OperatorPropertiesKeys. Note that this can
// be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (OperatorPropertiesKeys) Values() []OperatorPropertiesKeys {
	return []OperatorPropertiesKeys{
		"VALUE",
		"VALUES",
		"DATA_TYPE",
		"UPPER_BOUND",
		"LOWER_BOUND",
		"SOURCE_DATA_TYPE",
		"DESTINATION_DATA_TYPE",
		"VALIDATION_ACTION",
		"MASK_VALUE",
		"MASK_LENGTH",
		"TRUNCATE_LENGTH",
		"MATH_OPERATION_FIELDS_ORDER",
		"CONCAT_FORMAT",
		"SUBFIELD_CATEGORY_MAP",
	}
}

type PartyType string

// Enum values for PartyType
const (
	PartyTypeIndividual PartyType = "INDIVIDUAL"
	PartyTypeBusiness   PartyType = "BUSINESS"
	PartyTypeOther      PartyType = "OTHER"
)

// Values returns all known values for PartyType. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (PartyType) Values() []PartyType {
	return []PartyType{
		"INDIVIDUAL",
		"BUSINESS",
		"OTHER",
	}
}

type QueryResult string

// Enum values for QueryResult
const (
	QueryResultPresent QueryResult = "PRESENT"
	QueryResultAbsent  QueryResult = "ABSENT"
)

// Values returns all known values for QueryResult. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (QueryResult) Values() []QueryResult {
	return []QueryResult{
		"PRESENT",
		"ABSENT",
	}
}

type RangeUnit string

// Enum values for RangeUnit
const (
	RangeUnitDays RangeUnit = "DAYS"
)

// Values returns all known values for RangeUnit. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (RangeUnit) Values() []RangeUnit {
	return []RangeUnit{
		"DAYS",
	}
}

type RuleBasedMatchingStatus string

// Enum values for RuleBasedMatchingStatus
const (
	RuleBasedMatchingStatusPending    RuleBasedMatchingStatus = "PENDING"
	RuleBasedMatchingStatusInProgress RuleBasedMatchingStatus = "IN_PROGRESS"
	RuleBasedMatchingStatusActive     RuleBasedMatchingStatus = "ACTIVE"
)

// Values returns all known values for RuleBasedMatchingStatus. Note that this can
// be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (RuleBasedMatchingStatus) Values() []RuleBasedMatchingStatus {
	return []RuleBasedMatchingStatus{
		"PENDING",
		"IN_PROGRESS",
		"ACTIVE",
	}
}

type S3ConnectorOperator string

// Enum values for S3ConnectorOperator
const (
	S3ConnectorOperatorProjection           S3ConnectorOperator = "PROJECTION"
	S3ConnectorOperatorLessThan             S3ConnectorOperator = "LESS_THAN"
	S3ConnectorOperatorGreaterThan          S3ConnectorOperator = "GREATER_THAN"
	S3ConnectorOperatorBetween              S3ConnectorOperator = "BETWEEN"
	S3ConnectorOperatorLessThanOrEqualTo    S3ConnectorOperator = "LESS_THAN_OR_EQUAL_TO"
	S3ConnectorOperatorGreaterThanOrEqualTo S3ConnectorOperator = "GREATER_THAN_OR_EQUAL_TO"
	S3ConnectorOperatorEqualTo              S3ConnectorOperator = "EQUAL_TO"
	S3ConnectorOperatorNotEqualTo           S3ConnectorOperator = "NOT_EQUAL_TO"
	S3ConnectorOperatorAddition             S3ConnectorOperator = "ADDITION"
	S3ConnectorOperatorMultiplication       S3ConnectorOperator = "MULTIPLICATION"
	S3ConnectorOperatorDivision             S3ConnectorOperator = "DIVISION"
	S3ConnectorOperatorSubtraction          S3ConnectorOperator = "SUBTRACTION"
	S3ConnectorOperatorMaskAll              S3ConnectorOperator = "MASK_ALL"
	S3ConnectorOperatorMaskFirstN           S3ConnectorOperator = "MASK_FIRST_N"
	S3ConnectorOperatorMaskLastN            S3ConnectorOperator = "MASK_LAST_N"
	S3ConnectorOperatorValidateNonNull      S3ConnectorOperator = "VALIDATE_NON_NULL"
	S3ConnectorOperatorValidateNonZero      S3ConnectorOperator = "VALIDATE_NON_ZERO"
	S3ConnectorOperatorValidateNonNegative  S3ConnectorOperator = "VALIDATE_NON_NEGATIVE"
	S3ConnectorOperatorValidateNumeric      S3ConnectorOperator = "VALIDATE_NUMERIC"
	S3ConnectorOperatorNoOp                 S3ConnectorOperator = "NO_OP"
)

// Values returns all known values for S3ConnectorOperator. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (S3ConnectorOperator) Values() []S3ConnectorOperator {
	return []S3ConnectorOperator{
		"PROJECTION",
		"LESS_THAN",
		"GREATER_THAN",
		"BETWEEN",
		"LESS_THAN_OR_EQUAL_TO",
		"GREATER_THAN_OR_EQUAL_TO",
		"EQUAL_TO",
		"NOT_EQUAL_TO",
		"ADDITION",
		"MULTIPLICATION",
		"DIVISION",
		"SUBTRACTION",
		"MASK_ALL",
		"MASK_FIRST_N",
		"MASK_LAST_N",
		"VALIDATE_NON_NULL",
		"VALIDATE_NON_ZERO",
		"VALIDATE_NON_NEGATIVE",
		"VALIDATE_NUMERIC",
		"NO_OP",
	}
}

type SalesforceConnectorOperator string

// Enum values for SalesforceConnectorOperator
const (
	SalesforceConnectorOperatorProjection           SalesforceConnectorOperator = "PROJECTION"
	SalesforceConnectorOperatorLessThan             SalesforceConnectorOperator = "LESS_THAN"
	SalesforceConnectorOperatorContains             SalesforceConnectorOperator = "CONTAINS"
	SalesforceConnectorOperatorGreaterThan          SalesforceConnectorOperator = "GREATER_THAN"
	SalesforceConnectorOperatorBetween              SalesforceConnectorOperator = "BETWEEN"
	SalesforceConnectorOperatorLessThanOrEqualTo    SalesforceConnectorOperator = "LESS_THAN_OR_EQUAL_TO"
	SalesforceConnectorOperatorGreaterThanOrEqualTo SalesforceConnectorOperator = "GREATER_THAN_OR_EQUAL_TO"
	SalesforceConnectorOperatorEqualTo              SalesforceConnectorOperator = "EQUAL_TO"
	SalesforceConnectorOperatorNotEqualTo           SalesforceConnectorOperator = "NOT_EQUAL_TO"
	SalesforceConnectorOperatorAddition             SalesforceConnectorOperator = "ADDITION"
	SalesforceConnectorOperatorMultiplication       SalesforceConnectorOperator = "MULTIPLICATION"
	SalesforceConnectorOperatorDivision             SalesforceConnectorOperator = "DIVISION"
	SalesforceConnectorOperatorSubtraction          SalesforceConnectorOperator = "SUBTRACTION"
	SalesforceConnectorOperatorMaskAll              SalesforceConnectorOperator = "MASK_ALL"
	SalesforceConnectorOperatorMaskFirstN           SalesforceConnectorOperator = "MASK_FIRST_N"
	SalesforceConnectorOperatorMaskLastN            SalesforceConnectorOperator = "MASK_LAST_N"
	SalesforceConnectorOperatorValidateNonNull      SalesforceConnectorOperator = "VALIDATE_NON_NULL"
	SalesforceConnectorOperatorValidateNonZero      SalesforceConnectorOperator = "VALIDATE_NON_ZERO"
	SalesforceConnectorOperatorValidateNonNegative  SalesforceConnectorOperator = "VALIDATE_NON_NEGATIVE"
	SalesforceConnectorOperatorValidateNumeric      SalesforceConnectorOperator = "VALIDATE_NUMERIC"
	SalesforceConnectorOperatorNoOp                 SalesforceConnectorOperator = "NO_OP"
)

// Values returns all known values for SalesforceConnectorOperator. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (SalesforceConnectorOperator) Values() []SalesforceConnectorOperator {
	return []SalesforceConnectorOperator{
		"PROJECTION",
		"LESS_THAN",
		"CONTAINS",
		"GREATER_THAN",
		"BETWEEN",
		"LESS_THAN_OR_EQUAL_TO",
		"GREATER_THAN_OR_EQUAL_TO",
		"EQUAL_TO",
		"NOT_EQUAL_TO",
		"ADDITION",
		"MULTIPLICATION",
		"DIVISION",
		"SUBTRACTION",
		"MASK_ALL",
		"MASK_FIRST_N",
		"MASK_LAST_N",
		"VALIDATE_NON_NULL",
		"VALIDATE_NON_ZERO",
		"VALIDATE_NON_NEGATIVE",
		"VALIDATE_NUMERIC",
		"NO_OP",
	}
}

type SegmentSnapshotStatus string

// Enum values for SegmentSnapshotStatus
const (
	SegmentSnapshotStatusCompleted  SegmentSnapshotStatus = "COMPLETED"
	SegmentSnapshotStatusInProgress SegmentSnapshotStatus = "IN_PROGRESS"
	SegmentSnapshotStatusFailed     SegmentSnapshotStatus = "FAILED"
)

// Values returns all known values for SegmentSnapshotStatus. Note that this can
// be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (SegmentSnapshotStatus) Values() []SegmentSnapshotStatus {
	return []SegmentSnapshotStatus{
		"COMPLETED",
		"IN_PROGRESS",
		"FAILED",
	}
}

type ServiceNowConnectorOperator string

// Enum values for ServiceNowConnectorOperator
const (
	ServiceNowConnectorOperatorProjection           ServiceNowConnectorOperator = "PROJECTION"
	ServiceNowConnectorOperatorContains             ServiceNowConnectorOperator = "CONTAINS"
	ServiceNowConnectorOperatorLessThan             ServiceNowConnectorOperator = "LESS_THAN"
	ServiceNowConnectorOperatorGreaterThan          ServiceNowConnectorOperator = "GREATER_THAN"
	ServiceNowConnectorOperatorBetween              ServiceNowConnectorOperator = "BETWEEN"
	ServiceNowConnectorOperatorLessThanOrEqualTo    ServiceNowConnectorOperator = "LESS_THAN_OR_EQUAL_TO"
	ServiceNowConnectorOperatorGreaterThanOrEqualTo ServiceNowConnectorOperator = "GREATER_THAN_OR_EQUAL_TO"
	ServiceNowConnectorOperatorEqualTo              ServiceNowConnectorOperator = "EQUAL_TO"
	ServiceNowConnectorOperatorNotEqualTo           ServiceNowConnectorOperator = "NOT_EQUAL_TO"
	ServiceNowConnectorOperatorAddition             ServiceNowConnectorOperator = "ADDITION"
	ServiceNowConnectorOperatorMultiplication       ServiceNowConnectorOperator = "MULTIPLICATION"
	ServiceNowConnectorOperatorDivision             ServiceNowConnectorOperator = "DIVISION"
	ServiceNowConnectorOperatorSubtraction          ServiceNowConnectorOperator = "SUBTRACTION"
	ServiceNowConnectorOperatorMaskAll              ServiceNowConnectorOperator = "MASK_ALL"
	ServiceNowConnectorOperatorMaskFirstN           ServiceNowConnectorOperator = "MASK_FIRST_N"
	ServiceNowConnectorOperatorMaskLastN            ServiceNowConnectorOperator = "MASK_LAST_N"
	ServiceNowConnectorOperatorValidateNonNull      ServiceNowConnectorOperator = "VALIDATE_NON_NULL"
	ServiceNowConnectorOperatorValidateNonZero      ServiceNowConnectorOperator = "VALIDATE_NON_ZERO"
	ServiceNowConnectorOperatorValidateNonNegative  ServiceNowConnectorOperator = "VALIDATE_NON_NEGATIVE"
	ServiceNowConnectorOperatorValidateNumeric      ServiceNowConnectorOperator = "VALIDATE_NUMERIC"
	ServiceNowConnectorOperatorNoOp                 ServiceNowConnectorOperator = "NO_OP"
)

// Values returns all known values for ServiceNowConnectorOperator. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ServiceNowConnectorOperator) Values() []ServiceNowConnectorOperator {
	return []ServiceNowConnectorOperator{
		"PROJECTION",
		"CONTAINS",
		"LESS_THAN",
		"GREATER_THAN",
		"BETWEEN",
		"LESS_THAN_OR_EQUAL_TO",
		"GREATER_THAN_OR_EQUAL_TO",
		"EQUAL_TO",
		"NOT_EQUAL_TO",
		"ADDITION",
		"MULTIPLICATION",
		"DIVISION",
		"SUBTRACTION",
		"MASK_ALL",
		"MASK_FIRST_N",
		"MASK_LAST_N",
		"VALIDATE_NON_NULL",
		"VALIDATE_NON_ZERO",
		"VALIDATE_NON_NEGATIVE",
		"VALIDATE_NUMERIC",
		"NO_OP",
	}
}

type SourceConnectorType string

// Enum values for SourceConnectorType
const (
	SourceConnectorTypeSalesforce SourceConnectorType = "Salesforce"
	SourceConnectorTypeMarketo    SourceConnectorType = "Marketo"
	SourceConnectorTypeZendesk    SourceConnectorType = "Zendesk"
	SourceConnectorTypeServicenow SourceConnectorType = "Servicenow"
	SourceConnectorTypeS3         SourceConnectorType = "S3"
)

// Values returns all known values for SourceConnectorType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (SourceConnectorType) Values() []SourceConnectorType {
	return []SourceConnectorType{
		"Salesforce",
		"Marketo",
		"Zendesk",
		"Servicenow",
		"S3",
	}
}

type StandardIdentifier string

// Enum values for StandardIdentifier
const (
	StandardIdentifierProfile             StandardIdentifier = "PROFILE"
	StandardIdentifierAsset               StandardIdentifier = "ASSET"
	StandardIdentifierCase                StandardIdentifier = "CASE"
	StandardIdentifierOrder               StandardIdentifier = "ORDER"
	StandardIdentifierCommunicationRecord StandardIdentifier = "COMMUNICATION_RECORD"
	StandardIdentifierUnique              StandardIdentifier = "UNIQUE"
	StandardIdentifierSecondary           StandardIdentifier = "SECONDARY"
	StandardIdentifierLookupOnly          StandardIdentifier = "LOOKUP_ONLY"
	StandardIdentifierNewOnly             StandardIdentifier = "NEW_ONLY"
)

// Values returns all known values for StandardIdentifier. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (StandardIdentifier) Values() []StandardIdentifier {
	return []StandardIdentifier{
		"PROFILE",
		"ASSET",
		"CASE",
		"ORDER",
		"COMMUNICATION_RECORD",
		"UNIQUE",
		"SECONDARY",
		"LOOKUP_ONLY",
		"NEW_ONLY",
	}
}

type Statistic string

// Enum values for Statistic
const (
	StatisticFirstOccurrence Statistic = "FIRST_OCCURRENCE"
	StatisticLastOccurrence  Statistic = "LAST_OCCURRENCE"
	StatisticCount           Statistic = "COUNT"
	StatisticSum             Statistic = "SUM"
	StatisticMinimum         Statistic = "MINIMUM"
	StatisticMaximum         Statistic = "MAXIMUM"
	StatisticAverage         Statistic = "AVERAGE"
	StatisticMaxOccurrence   Statistic = "MAX_OCCURRENCE"
)

// Values returns all known values for Statistic. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (Statistic) Values() []Statistic {
	return []Statistic{
		"FIRST_OCCURRENCE",
		"LAST_OCCURRENCE",
		"COUNT",
		"SUM",
		"MINIMUM",
		"MAXIMUM",
		"AVERAGE",
		"MAX_OCCURRENCE",
	}
}

type Status string

// Enum values for Status
const (
	StatusNotStarted Status = "NOT_STARTED"
	StatusInProgress Status = "IN_PROGRESS"
	StatusComplete   Status = "COMPLETE"
	StatusFailed     Status = "FAILED"
	StatusSplit      Status = "SPLIT"
	StatusRetry      Status = "RETRY"
	StatusCancelled  Status = "CANCELLED"
)

// Values returns all known values for Status. Note that this can be expanded in
// the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (Status) Values() []Status {
	return []Status{
		"NOT_STARTED",
		"IN_PROGRESS",
		"COMPLETE",
		"FAILED",
		"SPLIT",
		"RETRY",
		"CANCELLED",
	}
}

type StringDimensionType string

// Enum values for StringDimensionType
const (
	StringDimensionTypeInclusive  StringDimensionType = "INCLUSIVE"
	StringDimensionTypeExclusive  StringDimensionType = "EXCLUSIVE"
	StringDimensionTypeContains   StringDimensionType = "CONTAINS"
	StringDimensionTypeBeginsWith StringDimensionType = "BEGINS_WITH"
	StringDimensionTypeEndsWith   StringDimensionType = "ENDS_WITH"
)

// Values returns all known values for StringDimensionType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (StringDimensionType) Values() []StringDimensionType {
	return []StringDimensionType{
		"INCLUSIVE",
		"EXCLUSIVE",
		"CONTAINS",
		"BEGINS_WITH",
		"ENDS_WITH",
	}
}

type TaskType string

// Enum values for TaskType
const (
	TaskTypeArithmetic TaskType = "Arithmetic"
	TaskTypeFilter     TaskType = "Filter"
	TaskTypeMap        TaskType = "Map"
	TaskTypeMask       TaskType = "Mask"
	TaskTypeMerge      TaskType = "Merge"
	TaskTypeTruncate   TaskType = "Truncate"
	TaskTypeValidate   TaskType = "Validate"
)

// Values returns all known values for TaskType. Note that this can be expanded in
// the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (TaskType) Values() []TaskType {
	return []TaskType{
		"Arithmetic",
		"Filter",
		"Map",
		"Mask",
		"Merge",
		"Truncate",
		"Validate",
	}
}

type TriggerType string

// Enum values for TriggerType
const (
	TriggerTypeScheduled TriggerType = "Scheduled"
	TriggerTypeEvent     TriggerType = "Event"
	TriggerTypeOndemand  TriggerType = "OnDemand"
)

// Values returns all known values for TriggerType. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (TriggerType) Values() []TriggerType {
	return []TriggerType{
		"Scheduled",
		"Event",
		"OnDemand",
	}
}

type Type string

// Enum values for Type
const (
	TypeAll  Type = "ALL"
	TypeAny  Type = "ANY"
	TypeNone Type = "NONE"
)

// Values returns all known values for Type. Note that this can be expanded in the
// future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (Type) Values() []Type {
	return []Type{
		"ALL",
		"ANY",
		"NONE",
	}
}

type Unit string

// Enum values for Unit
const (
	UnitDays Unit = "DAYS"
)

// Values returns all known values for Unit. Note that this can be expanded in the
// future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (Unit) Values() []Unit {
	return []Unit{
		"DAYS",
	}
}

type WorkflowType string

// Enum values for WorkflowType
const (
	WorkflowTypeAppflowIntegration WorkflowType = "APPFLOW_INTEGRATION"
)

// Values returns all known values for WorkflowType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (WorkflowType) Values() []WorkflowType {
	return []WorkflowType{
		"APPFLOW_INTEGRATION",
	}
}

type ZendeskConnectorOperator string

// Enum values for ZendeskConnectorOperator
const (
	ZendeskConnectorOperatorProjection          ZendeskConnectorOperator = "PROJECTION"
	ZendeskConnectorOperatorGreaterThan         ZendeskConnectorOperator = "GREATER_THAN"
	ZendeskConnectorOperatorAddition            ZendeskConnectorOperator = "ADDITION"
	ZendeskConnectorOperatorMultiplication      ZendeskConnectorOperator = "MULTIPLICATION"
	ZendeskConnectorOperatorDivision            ZendeskConnectorOperator = "DIVISION"
	ZendeskConnectorOperatorSubtraction         ZendeskConnectorOperator = "SUBTRACTION"
	ZendeskConnectorOperatorMaskAll             ZendeskConnectorOperator = "MASK_ALL"
	ZendeskConnectorOperatorMaskFirstN          ZendeskConnectorOperator = "MASK_FIRST_N"
	ZendeskConnectorOperatorMaskLastN           ZendeskConnectorOperator = "MASK_LAST_N"
	ZendeskConnectorOperatorValidateNonNull     ZendeskConnectorOperator = "VALIDATE_NON_NULL"
	ZendeskConnectorOperatorValidateNonZero     ZendeskConnectorOperator = "VALIDATE_NON_ZERO"
	ZendeskConnectorOperatorValidateNonNegative ZendeskConnectorOperator = "VALIDATE_NON_NEGATIVE"
	ZendeskConnectorOperatorValidateNumeric     ZendeskConnectorOperator = "VALIDATE_NUMERIC"
	ZendeskConnectorOperatorNoOp                ZendeskConnectorOperator = "NO_OP"
)

// Values returns all known values for ZendeskConnectorOperator. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ZendeskConnectorOperator) Values() []ZendeskConnectorOperator {
	return []ZendeskConnectorOperator{
		"PROJECTION",
		"GREATER_THAN",
		"ADDITION",
		"MULTIPLICATION",
		"DIVISION",
		"SUBTRACTION",
		"MASK_ALL",
		"MASK_FIRST_N",
		"MASK_LAST_N",
		"VALIDATE_NON_NULL",
		"VALIDATE_NON_ZERO",
		"VALIDATE_NON_NEGATIVE",
		"VALIDATE_NUMERIC",
		"NO_OP",
	}
}
