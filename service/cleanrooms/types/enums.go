// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type AccessDeniedExceptionReason string

// Enum values for AccessDeniedExceptionReason
const (
	AccessDeniedExceptionReasonInsufficientPermissions AccessDeniedExceptionReason = "INSUFFICIENT_PERMISSIONS"
)

// Values returns all known values for AccessDeniedExceptionReason. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (AccessDeniedExceptionReason) Values() []AccessDeniedExceptionReason {
	return []AccessDeniedExceptionReason{
		"INSUFFICIENT_PERMISSIONS",
	}
}

type AdditionalAnalyses string

// Enum values for AdditionalAnalyses
const (
	AdditionalAnalysesAllowed    AdditionalAnalyses = "ALLOWED"
	AdditionalAnalysesRequired   AdditionalAnalyses = "REQUIRED"
	AdditionalAnalysesNotAllowed AdditionalAnalyses = "NOT_ALLOWED"
)

// Values returns all known values for AdditionalAnalyses. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (AdditionalAnalyses) Values() []AdditionalAnalyses {
	return []AdditionalAnalyses{
		"ALLOWED",
		"REQUIRED",
		"NOT_ALLOWED",
	}
}

type AggregateFunctionName string

// Enum values for AggregateFunctionName
const (
	AggregateFunctionNameSum           AggregateFunctionName = "SUM"
	AggregateFunctionNameSumDistinct   AggregateFunctionName = "SUM_DISTINCT"
	AggregateFunctionNameCount         AggregateFunctionName = "COUNT"
	AggregateFunctionNameCountDistinct AggregateFunctionName = "COUNT_DISTINCT"
	AggregateFunctionNameAvg           AggregateFunctionName = "AVG"
)

// Values returns all known values for AggregateFunctionName. Note that this can
// be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (AggregateFunctionName) Values() []AggregateFunctionName {
	return []AggregateFunctionName{
		"SUM",
		"SUM_DISTINCT",
		"COUNT",
		"COUNT_DISTINCT",
		"AVG",
	}
}

type AggregationType string

// Enum values for AggregationType
const (
	AggregationTypeCountDistinct AggregationType = "COUNT_DISTINCT"
)

// Values returns all known values for AggregationType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (AggregationType) Values() []AggregationType {
	return []AggregationType{
		"COUNT_DISTINCT",
	}
}

type AnalysisFormat string

// Enum values for AnalysisFormat
const (
	AnalysisFormatSql AnalysisFormat = "SQL"
)

// Values returns all known values for AnalysisFormat. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (AnalysisFormat) Values() []AnalysisFormat {
	return []AnalysisFormat{
		"SQL",
	}
}

type AnalysisMethod string

// Enum values for AnalysisMethod
const (
	AnalysisMethodDirectQuery AnalysisMethod = "DIRECT_QUERY"
)

// Values returns all known values for AnalysisMethod. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (AnalysisMethod) Values() []AnalysisMethod {
	return []AnalysisMethod{
		"DIRECT_QUERY",
	}
}

type AnalysisRuleType string

// Enum values for AnalysisRuleType
const (
	AnalysisRuleTypeAggregation    AnalysisRuleType = "AGGREGATION"
	AnalysisRuleTypeList           AnalysisRuleType = "LIST"
	AnalysisRuleTypeCustom         AnalysisRuleType = "CUSTOM"
	AnalysisRuleTypeIdMappingTable AnalysisRuleType = "ID_MAPPING_TABLE"
)

// Values returns all known values for AnalysisRuleType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (AnalysisRuleType) Values() []AnalysisRuleType {
	return []AnalysisRuleType{
		"AGGREGATION",
		"LIST",
		"CUSTOM",
		"ID_MAPPING_TABLE",
	}
}

type AnalysisTemplateValidationStatus string

// Enum values for AnalysisTemplateValidationStatus
const (
	AnalysisTemplateValidationStatusValid            AnalysisTemplateValidationStatus = "VALID"
	AnalysisTemplateValidationStatusInvalid          AnalysisTemplateValidationStatus = "INVALID"
	AnalysisTemplateValidationStatusUnableToValidate AnalysisTemplateValidationStatus = "UNABLE_TO_VALIDATE"
)

// Values returns all known values for AnalysisTemplateValidationStatus. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (AnalysisTemplateValidationStatus) Values() []AnalysisTemplateValidationStatus {
	return []AnalysisTemplateValidationStatus{
		"VALID",
		"INVALID",
		"UNABLE_TO_VALIDATE",
	}
}

type AnalysisTemplateValidationType string

// Enum values for AnalysisTemplateValidationType
const (
	AnalysisTemplateValidationTypeDifferentialPrivacy AnalysisTemplateValidationType = "DIFFERENTIAL_PRIVACY"
)

// Values returns all known values for AnalysisTemplateValidationType. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (AnalysisTemplateValidationType) Values() []AnalysisTemplateValidationType {
	return []AnalysisTemplateValidationType{
		"DIFFERENTIAL_PRIVACY",
	}
}

type AnalysisType string

// Enum values for AnalysisType
const (
	AnalysisTypeDirectAnalysis     AnalysisType = "DIRECT_ANALYSIS"
	AnalysisTypeAdditionalAnalysis AnalysisType = "ADDITIONAL_ANALYSIS"
)

// Values returns all known values for AnalysisType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (AnalysisType) Values() []AnalysisType {
	return []AnalysisType{
		"DIRECT_ANALYSIS",
		"ADDITIONAL_ANALYSIS",
	}
}

type AnalyticsEngine string

// Enum values for AnalyticsEngine
const (
	AnalyticsEngineSpark         AnalyticsEngine = "SPARK"
	AnalyticsEngineCleanRoomsSql AnalyticsEngine = "CLEAN_ROOMS_SQL"
)

// Values returns all known values for AnalyticsEngine. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (AnalyticsEngine) Values() []AnalyticsEngine {
	return []AnalyticsEngine{
		"SPARK",
		"CLEAN_ROOMS_SQL",
	}
}

type CollaborationQueryLogStatus string

// Enum values for CollaborationQueryLogStatus
const (
	CollaborationQueryLogStatusEnabled  CollaborationQueryLogStatus = "ENABLED"
	CollaborationQueryLogStatusDisabled CollaborationQueryLogStatus = "DISABLED"
)

// Values returns all known values for CollaborationQueryLogStatus. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (CollaborationQueryLogStatus) Values() []CollaborationQueryLogStatus {
	return []CollaborationQueryLogStatus{
		"ENABLED",
		"DISABLED",
	}
}

type ConfiguredTableAnalysisRuleType string

// Enum values for ConfiguredTableAnalysisRuleType
const (
	ConfiguredTableAnalysisRuleTypeAggregation ConfiguredTableAnalysisRuleType = "AGGREGATION"
	ConfiguredTableAnalysisRuleTypeList        ConfiguredTableAnalysisRuleType = "LIST"
	ConfiguredTableAnalysisRuleTypeCustom      ConfiguredTableAnalysisRuleType = "CUSTOM"
)

// Values returns all known values for ConfiguredTableAnalysisRuleType. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ConfiguredTableAnalysisRuleType) Values() []ConfiguredTableAnalysisRuleType {
	return []ConfiguredTableAnalysisRuleType{
		"AGGREGATION",
		"LIST",
		"CUSTOM",
	}
}

type ConfiguredTableAssociationAnalysisRuleType string

// Enum values for ConfiguredTableAssociationAnalysisRuleType
const (
	ConfiguredTableAssociationAnalysisRuleTypeAggregation ConfiguredTableAssociationAnalysisRuleType = "AGGREGATION"
	ConfiguredTableAssociationAnalysisRuleTypeList        ConfiguredTableAssociationAnalysisRuleType = "LIST"
	ConfiguredTableAssociationAnalysisRuleTypeCustom      ConfiguredTableAssociationAnalysisRuleType = "CUSTOM"
)

// Values returns all known values for ConfiguredTableAssociationAnalysisRuleType.
// Note that this can be expanded in the future, and so it is only as up to date as
// the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ConfiguredTableAssociationAnalysisRuleType) Values() []ConfiguredTableAssociationAnalysisRuleType {
	return []ConfiguredTableAssociationAnalysisRuleType{
		"AGGREGATION",
		"LIST",
		"CUSTOM",
	}
}

type ConflictExceptionReason string

// Enum values for ConflictExceptionReason
const (
	ConflictExceptionReasonAlreadyExists     ConflictExceptionReason = "ALREADY_EXISTS"
	ConflictExceptionReasonSubresourcesExist ConflictExceptionReason = "SUBRESOURCES_EXIST"
	ConflictExceptionReasonInvalidState      ConflictExceptionReason = "INVALID_STATE"
)

// Values returns all known values for ConflictExceptionReason. Note that this can
// be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ConflictExceptionReason) Values() []ConflictExceptionReason {
	return []ConflictExceptionReason{
		"ALREADY_EXISTS",
		"SUBRESOURCES_EXIST",
		"INVALID_STATE",
	}
}

type DifferentialPrivacyAggregationType string

// Enum values for DifferentialPrivacyAggregationType
const (
	DifferentialPrivacyAggregationTypeAvg           DifferentialPrivacyAggregationType = "AVG"
	DifferentialPrivacyAggregationTypeCount         DifferentialPrivacyAggregationType = "COUNT"
	DifferentialPrivacyAggregationTypeCountDistinct DifferentialPrivacyAggregationType = "COUNT_DISTINCT"
	DifferentialPrivacyAggregationTypeSum           DifferentialPrivacyAggregationType = "SUM"
	DifferentialPrivacyAggregationTypeStddev        DifferentialPrivacyAggregationType = "STDDEV"
)

// Values returns all known values for DifferentialPrivacyAggregationType. Note
// that this can be expanded in the future, and so it is only as up to date as the
// client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (DifferentialPrivacyAggregationType) Values() []DifferentialPrivacyAggregationType {
	return []DifferentialPrivacyAggregationType{
		"AVG",
		"COUNT",
		"COUNT_DISTINCT",
		"SUM",
		"STDDEV",
	}
}

type FilterableMemberStatus string

// Enum values for FilterableMemberStatus
const (
	FilterableMemberStatusInvited FilterableMemberStatus = "INVITED"
	FilterableMemberStatusActive  FilterableMemberStatus = "ACTIVE"
)

// Values returns all known values for FilterableMemberStatus. Note that this can
// be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (FilterableMemberStatus) Values() []FilterableMemberStatus {
	return []FilterableMemberStatus{
		"INVITED",
		"ACTIVE",
	}
}

type IdNamespaceType string

// Enum values for IdNamespaceType
const (
	IdNamespaceTypeSource IdNamespaceType = "SOURCE"
	IdNamespaceTypeTarget IdNamespaceType = "TARGET"
)

// Values returns all known values for IdNamespaceType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (IdNamespaceType) Values() []IdNamespaceType {
	return []IdNamespaceType{
		"SOURCE",
		"TARGET",
	}
}

type JoinOperator string

// Enum values for JoinOperator
const (
	JoinOperatorOr  JoinOperator = "OR"
	JoinOperatorAnd JoinOperator = "AND"
)

// Values returns all known values for JoinOperator. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (JoinOperator) Values() []JoinOperator {
	return []JoinOperator{
		"OR",
		"AND",
	}
}

type JoinRequiredOption string

// Enum values for JoinRequiredOption
const (
	JoinRequiredOptionQueryRunner JoinRequiredOption = "QUERY_RUNNER"
)

// Values returns all known values for JoinRequiredOption. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (JoinRequiredOption) Values() []JoinRequiredOption {
	return []JoinRequiredOption{
		"QUERY_RUNNER",
	}
}

type MemberAbility string

// Enum values for MemberAbility
const (
	MemberAbilityCanQuery          MemberAbility = "CAN_QUERY"
	MemberAbilityCanReceiveResults MemberAbility = "CAN_RECEIVE_RESULTS"
)

// Values returns all known values for MemberAbility. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (MemberAbility) Values() []MemberAbility {
	return []MemberAbility{
		"CAN_QUERY",
		"CAN_RECEIVE_RESULTS",
	}
}

type MembershipQueryLogStatus string

// Enum values for MembershipQueryLogStatus
const (
	MembershipQueryLogStatusEnabled  MembershipQueryLogStatus = "ENABLED"
	MembershipQueryLogStatusDisabled MembershipQueryLogStatus = "DISABLED"
)

// Values returns all known values for MembershipQueryLogStatus. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (MembershipQueryLogStatus) Values() []MembershipQueryLogStatus {
	return []MembershipQueryLogStatus{
		"ENABLED",
		"DISABLED",
	}
}

type MembershipStatus string

// Enum values for MembershipStatus
const (
	MembershipStatusActive               MembershipStatus = "ACTIVE"
	MembershipStatusRemoved              MembershipStatus = "REMOVED"
	MembershipStatusCollaborationDeleted MembershipStatus = "COLLABORATION_DELETED"
)

// Values returns all known values for MembershipStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (MembershipStatus) Values() []MembershipStatus {
	return []MembershipStatus{
		"ACTIVE",
		"REMOVED",
		"COLLABORATION_DELETED",
	}
}

type MemberStatus string

// Enum values for MemberStatus
const (
	MemberStatusInvited MemberStatus = "INVITED"
	MemberStatusActive  MemberStatus = "ACTIVE"
	MemberStatusLeft    MemberStatus = "LEFT"
	MemberStatusRemoved MemberStatus = "REMOVED"
)

// Values returns all known values for MemberStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (MemberStatus) Values() []MemberStatus {
	return []MemberStatus{
		"INVITED",
		"ACTIVE",
		"LEFT",
		"REMOVED",
	}
}

type ParameterType string

// Enum values for ParameterType
const (
	ParameterTypeSmallint        ParameterType = "SMALLINT"
	ParameterTypeInteger         ParameterType = "INTEGER"
	ParameterTypeBigint          ParameterType = "BIGINT"
	ParameterTypeDecimal         ParameterType = "DECIMAL"
	ParameterTypeReal            ParameterType = "REAL"
	ParameterTypeDoublePrecision ParameterType = "DOUBLE_PRECISION"
	ParameterTypeBoolean         ParameterType = "BOOLEAN"
	ParameterTypeChar            ParameterType = "CHAR"
	ParameterTypeVarchar         ParameterType = "VARCHAR"
	ParameterTypeDate            ParameterType = "DATE"
	ParameterTypeTimestamp       ParameterType = "TIMESTAMP"
	ParameterTypeTimestamptz     ParameterType = "TIMESTAMPTZ"
	ParameterTypeTime            ParameterType = "TIME"
	ParameterTypeTimetz          ParameterType = "TIMETZ"
	ParameterTypeVarbyte         ParameterType = "VARBYTE"
	ParameterTypeBinary          ParameterType = "BINARY"
	ParameterTypeByte            ParameterType = "BYTE"
	ParameterTypeCharacter       ParameterType = "CHARACTER"
	ParameterTypeDouble          ParameterType = "DOUBLE"
	ParameterTypeFloat           ParameterType = "FLOAT"
	ParameterTypeInt             ParameterType = "INT"
	ParameterTypeLong            ParameterType = "LONG"
	ParameterTypeNumeric         ParameterType = "NUMERIC"
	ParameterTypeShort           ParameterType = "SHORT"
	ParameterTypeString          ParameterType = "STRING"
	ParameterTypeTimestampLtz    ParameterType = "TIMESTAMP_LTZ"
	ParameterTypeTimestampNtz    ParameterType = "TIMESTAMP_NTZ"
	ParameterTypeTinyint         ParameterType = "TINYINT"
)

// Values returns all known values for ParameterType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ParameterType) Values() []ParameterType {
	return []ParameterType{
		"SMALLINT",
		"INTEGER",
		"BIGINT",
		"DECIMAL",
		"REAL",
		"DOUBLE_PRECISION",
		"BOOLEAN",
		"CHAR",
		"VARCHAR",
		"DATE",
		"TIMESTAMP",
		"TIMESTAMPTZ",
		"TIME",
		"TIMETZ",
		"VARBYTE",
		"BINARY",
		"BYTE",
		"CHARACTER",
		"DOUBLE",
		"FLOAT",
		"INT",
		"LONG",
		"NUMERIC",
		"SHORT",
		"STRING",
		"TIMESTAMP_LTZ",
		"TIMESTAMP_NTZ",
		"TINYINT",
	}
}

type PrivacyBudgetTemplateAutoRefresh string

// Enum values for PrivacyBudgetTemplateAutoRefresh
const (
	PrivacyBudgetTemplateAutoRefreshCalendarMonth PrivacyBudgetTemplateAutoRefresh = "CALENDAR_MONTH"
	PrivacyBudgetTemplateAutoRefreshNone          PrivacyBudgetTemplateAutoRefresh = "NONE"
)

// Values returns all known values for PrivacyBudgetTemplateAutoRefresh. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (PrivacyBudgetTemplateAutoRefresh) Values() []PrivacyBudgetTemplateAutoRefresh {
	return []PrivacyBudgetTemplateAutoRefresh{
		"CALENDAR_MONTH",
		"NONE",
	}
}

type PrivacyBudgetType string

// Enum values for PrivacyBudgetType
const (
	PrivacyBudgetTypeDifferentialPrivacy PrivacyBudgetType = "DIFFERENTIAL_PRIVACY"
)

// Values returns all known values for PrivacyBudgetType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (PrivacyBudgetType) Values() []PrivacyBudgetType {
	return []PrivacyBudgetType{
		"DIFFERENTIAL_PRIVACY",
	}
}

type ProtectedQueryStatus string

// Enum values for ProtectedQueryStatus
const (
	ProtectedQueryStatusSubmitted  ProtectedQueryStatus = "SUBMITTED"
	ProtectedQueryStatusStarted    ProtectedQueryStatus = "STARTED"
	ProtectedQueryStatusCancelled  ProtectedQueryStatus = "CANCELLED"
	ProtectedQueryStatusCancelling ProtectedQueryStatus = "CANCELLING"
	ProtectedQueryStatusFailed     ProtectedQueryStatus = "FAILED"
	ProtectedQueryStatusSuccess    ProtectedQueryStatus = "SUCCESS"
	ProtectedQueryStatusTimedOut   ProtectedQueryStatus = "TIMED_OUT"
)

// Values returns all known values for ProtectedQueryStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ProtectedQueryStatus) Values() []ProtectedQueryStatus {
	return []ProtectedQueryStatus{
		"SUBMITTED",
		"STARTED",
		"CANCELLED",
		"CANCELLING",
		"FAILED",
		"SUCCESS",
		"TIMED_OUT",
	}
}

type ProtectedQueryType string

// Enum values for ProtectedQueryType
const (
	ProtectedQueryTypeSql ProtectedQueryType = "SQL"
)

// Values returns all known values for ProtectedQueryType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ProtectedQueryType) Values() []ProtectedQueryType {
	return []ProtectedQueryType{
		"SQL",
	}
}

type ResourceType string

// Enum values for ResourceType
const (
	ResourceTypeConfiguredTable            ResourceType = "CONFIGURED_TABLE"
	ResourceTypeCollaboration              ResourceType = "COLLABORATION"
	ResourceTypeMembership                 ResourceType = "MEMBERSHIP"
	ResourceTypeConfiguredTableAssociation ResourceType = "CONFIGURED_TABLE_ASSOCIATION"
)

// Values returns all known values for ResourceType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ResourceType) Values() []ResourceType {
	return []ResourceType{
		"CONFIGURED_TABLE",
		"COLLABORATION",
		"MEMBERSHIP",
		"CONFIGURED_TABLE_ASSOCIATION",
	}
}

type ResultFormat string

// Enum values for ResultFormat
const (
	ResultFormatCsv     ResultFormat = "CSV"
	ResultFormatParquet ResultFormat = "PARQUET"
)

// Values returns all known values for ResultFormat. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ResultFormat) Values() []ResultFormat {
	return []ResultFormat{
		"CSV",
		"PARQUET",
	}
}

type ScalarFunctions string

// Enum values for ScalarFunctions
const (
	ScalarFunctionsAbs         ScalarFunctions = "ABS"
	ScalarFunctionsCast        ScalarFunctions = "CAST"
	ScalarFunctionsCeiling     ScalarFunctions = "CEILING"
	ScalarFunctionsCoalesce    ScalarFunctions = "COALESCE"
	ScalarFunctionsConvert     ScalarFunctions = "CONVERT"
	ScalarFunctionsCurrentDate ScalarFunctions = "CURRENT_DATE"
	ScalarFunctionsDateadd     ScalarFunctions = "DATEADD"
	ScalarFunctionsExtract     ScalarFunctions = "EXTRACT"
	ScalarFunctionsFloor       ScalarFunctions = "FLOOR"
	ScalarFunctionsGetdate     ScalarFunctions = "GETDATE"
	ScalarFunctionsLn          ScalarFunctions = "LN"
	ScalarFunctionsLog         ScalarFunctions = "LOG"
	ScalarFunctionsLower       ScalarFunctions = "LOWER"
	ScalarFunctionsRound       ScalarFunctions = "ROUND"
	ScalarFunctionsRtrim       ScalarFunctions = "RTRIM"
	ScalarFunctionsSqrt        ScalarFunctions = "SQRT"
	ScalarFunctionsSubstring   ScalarFunctions = "SUBSTRING"
	ScalarFunctionsToChar      ScalarFunctions = "TO_CHAR"
	ScalarFunctionsToDate      ScalarFunctions = "TO_DATE"
	ScalarFunctionsToNumber    ScalarFunctions = "TO_NUMBER"
	ScalarFunctionsToTimestamp ScalarFunctions = "TO_TIMESTAMP"
	ScalarFunctionsTrim        ScalarFunctions = "TRIM"
	ScalarFunctionsTrunc       ScalarFunctions = "TRUNC"
	ScalarFunctionsUpper       ScalarFunctions = "UPPER"
)

// Values returns all known values for ScalarFunctions. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ScalarFunctions) Values() []ScalarFunctions {
	return []ScalarFunctions{
		"ABS",
		"CAST",
		"CEILING",
		"COALESCE",
		"CONVERT",
		"CURRENT_DATE",
		"DATEADD",
		"EXTRACT",
		"FLOOR",
		"GETDATE",
		"LN",
		"LOG",
		"LOWER",
		"ROUND",
		"RTRIM",
		"SQRT",
		"SUBSTRING",
		"TO_CHAR",
		"TO_DATE",
		"TO_NUMBER",
		"TO_TIMESTAMP",
		"TRIM",
		"TRUNC",
		"UPPER",
	}
}

type SchemaConfiguration string

// Enum values for SchemaConfiguration
const (
	SchemaConfigurationDifferentialPrivacy SchemaConfiguration = "DIFFERENTIAL_PRIVACY"
)

// Values returns all known values for SchemaConfiguration. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (SchemaConfiguration) Values() []SchemaConfiguration {
	return []SchemaConfiguration{
		"DIFFERENTIAL_PRIVACY",
	}
}

type SchemaStatus string

// Enum values for SchemaStatus
const (
	SchemaStatusReady    SchemaStatus = "READY"
	SchemaStatusNotReady SchemaStatus = "NOT_READY"
)

// Values returns all known values for SchemaStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (SchemaStatus) Values() []SchemaStatus {
	return []SchemaStatus{
		"READY",
		"NOT_READY",
	}
}

type SchemaStatusReasonCode string

// Enum values for SchemaStatusReasonCode
const (
	SchemaStatusReasonCodeAnalysisRuleMissing                    SchemaStatusReasonCode = "ANALYSIS_RULE_MISSING"
	SchemaStatusReasonCodeAnalysisTemplatesNotConfigured         SchemaStatusReasonCode = "ANALYSIS_TEMPLATES_NOT_CONFIGURED"
	SchemaStatusReasonCodeAnalysisProvidersNotConfigured         SchemaStatusReasonCode = "ANALYSIS_PROVIDERS_NOT_CONFIGURED"
	SchemaStatusReasonCodeDifferentialPrivacyPolicyNotConfigured SchemaStatusReasonCode = "DIFFERENTIAL_PRIVACY_POLICY_NOT_CONFIGURED"
	SchemaStatusReasonCodeIdMappingTableNotPopulated             SchemaStatusReasonCode = "ID_MAPPING_TABLE_NOT_POPULATED"
	SchemaStatusReasonCodeCollaborationAnalysisRuleNotConfigured SchemaStatusReasonCode = "COLLABORATION_ANALYSIS_RULE_NOT_CONFIGURED"
	SchemaStatusReasonCodeAdditionalAnalysesNotConfigured        SchemaStatusReasonCode = "ADDITIONAL_ANALYSES_NOT_CONFIGURED"
	SchemaStatusReasonCodeResultReceiversNotConfigured           SchemaStatusReasonCode = "RESULT_RECEIVERS_NOT_CONFIGURED"
	SchemaStatusReasonCodeAdditionalAnalysesNotAllowed           SchemaStatusReasonCode = "ADDITIONAL_ANALYSES_NOT_ALLOWED"
	SchemaStatusReasonCodeResultReceiversNotAllowed              SchemaStatusReasonCode = "RESULT_RECEIVERS_NOT_ALLOWED"
	SchemaStatusReasonCodeAnalysisRuleTypesNotCompatible         SchemaStatusReasonCode = "ANALYSIS_RULE_TYPES_NOT_COMPATIBLE"
)

// Values returns all known values for SchemaStatusReasonCode. Note that this can
// be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (SchemaStatusReasonCode) Values() []SchemaStatusReasonCode {
	return []SchemaStatusReasonCode{
		"ANALYSIS_RULE_MISSING",
		"ANALYSIS_TEMPLATES_NOT_CONFIGURED",
		"ANALYSIS_PROVIDERS_NOT_CONFIGURED",
		"DIFFERENTIAL_PRIVACY_POLICY_NOT_CONFIGURED",
		"ID_MAPPING_TABLE_NOT_POPULATED",
		"COLLABORATION_ANALYSIS_RULE_NOT_CONFIGURED",
		"ADDITIONAL_ANALYSES_NOT_CONFIGURED",
		"RESULT_RECEIVERS_NOT_CONFIGURED",
		"ADDITIONAL_ANALYSES_NOT_ALLOWED",
		"RESULT_RECEIVERS_NOT_ALLOWED",
		"ANALYSIS_RULE_TYPES_NOT_COMPATIBLE",
	}
}

type SchemaType string

// Enum values for SchemaType
const (
	SchemaTypeTable          SchemaType = "TABLE"
	SchemaTypeIdMappingTable SchemaType = "ID_MAPPING_TABLE"
)

// Values returns all known values for SchemaType. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (SchemaType) Values() []SchemaType {
	return []SchemaType{
		"TABLE",
		"ID_MAPPING_TABLE",
	}
}

type TargetProtectedQueryStatus string

// Enum values for TargetProtectedQueryStatus
const (
	TargetProtectedQueryStatusCancelled TargetProtectedQueryStatus = "CANCELLED"
)

// Values returns all known values for TargetProtectedQueryStatus. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (TargetProtectedQueryStatus) Values() []TargetProtectedQueryStatus {
	return []TargetProtectedQueryStatus{
		"CANCELLED",
	}
}

type ValidationExceptionReason string

// Enum values for ValidationExceptionReason
const (
	ValidationExceptionReasonFieldValidationFailed   ValidationExceptionReason = "FIELD_VALIDATION_FAILED"
	ValidationExceptionReasonInvalidConfiguration    ValidationExceptionReason = "INVALID_CONFIGURATION"
	ValidationExceptionReasonInvalidQuery            ValidationExceptionReason = "INVALID_QUERY"
	ValidationExceptionReasonIamSynchronizationDelay ValidationExceptionReason = "IAM_SYNCHRONIZATION_DELAY"
)

// Values returns all known values for ValidationExceptionReason. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ValidationExceptionReason) Values() []ValidationExceptionReason {
	return []ValidationExceptionReason{
		"FIELD_VALIDATION_FAILED",
		"INVALID_CONFIGURATION",
		"INVALID_QUERY",
		"IAM_SYNCHRONIZATION_DELAY",
	}
}

type WorkerComputeType string

// Enum values for WorkerComputeType
const (
	WorkerComputeTypeCr1x WorkerComputeType = "CR.1X"
	WorkerComputeTypeCr4x WorkerComputeType = "CR.4X"
)

// Values returns all known values for WorkerComputeType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (WorkerComputeType) Values() []WorkerComputeType {
	return []WorkerComputeType{
		"CR.1X",
		"CR.4X",
	}
}
