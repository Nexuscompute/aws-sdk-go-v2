// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type Action string

// Enum values for Action
const (
	ActionAllow Action = "ALLOW"
	ActionBlock Action = "BLOCK"
	ActionAlert Action = "ALERT"
)

// Values returns all known values for Action. Note that this can be expanded in
// the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (Action) Values() []Action {
	return []Action{
		"ALLOW",
		"BLOCK",
		"ALERT",
	}
}

type AutodefinedReverseFlag string

// Enum values for AutodefinedReverseFlag
const (
	AutodefinedReverseFlagEnable                  AutodefinedReverseFlag = "ENABLE"
	AutodefinedReverseFlagDisable                 AutodefinedReverseFlag = "DISABLE"
	AutodefinedReverseFlagUseLocalResourceSetting AutodefinedReverseFlag = "USE_LOCAL_RESOURCE_SETTING"
)

// Values returns all known values for AutodefinedReverseFlag. Note that this can
// be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (AutodefinedReverseFlag) Values() []AutodefinedReverseFlag {
	return []AutodefinedReverseFlag{
		"ENABLE",
		"DISABLE",
		"USE_LOCAL_RESOURCE_SETTING",
	}
}

type BlockOverrideDnsType string

// Enum values for BlockOverrideDnsType
const (
	BlockOverrideDnsTypeCname BlockOverrideDnsType = "CNAME"
)

// Values returns all known values for BlockOverrideDnsType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (BlockOverrideDnsType) Values() []BlockOverrideDnsType {
	return []BlockOverrideDnsType{
		"CNAME",
	}
}

type BlockResponse string

// Enum values for BlockResponse
const (
	BlockResponseNodata   BlockResponse = "NODATA"
	BlockResponseNxdomain BlockResponse = "NXDOMAIN"
	BlockResponseOverride BlockResponse = "OVERRIDE"
)

// Values returns all known values for BlockResponse. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (BlockResponse) Values() []BlockResponse {
	return []BlockResponse{
		"NODATA",
		"NXDOMAIN",
		"OVERRIDE",
	}
}

type ConfidenceThreshold string

// Enum values for ConfidenceThreshold
const (
	ConfidenceThresholdLow    ConfidenceThreshold = "LOW"
	ConfidenceThresholdMedium ConfidenceThreshold = "MEDIUM"
	ConfidenceThresholdHigh   ConfidenceThreshold = "HIGH"
)

// Values returns all known values for ConfidenceThreshold. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ConfidenceThreshold) Values() []ConfidenceThreshold {
	return []ConfidenceThreshold{
		"LOW",
		"MEDIUM",
		"HIGH",
	}
}

type DnsThreatProtection string

// Enum values for DnsThreatProtection
const (
	DnsThreatProtectionDga          DnsThreatProtection = "DGA"
	DnsThreatProtectionDnsTunneling DnsThreatProtection = "DNS_TUNNELING"
)

// Values returns all known values for DnsThreatProtection. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (DnsThreatProtection) Values() []DnsThreatProtection {
	return []DnsThreatProtection{
		"DGA",
		"DNS_TUNNELING",
	}
}

type FirewallDomainImportOperation string

// Enum values for FirewallDomainImportOperation
const (
	FirewallDomainImportOperationReplace FirewallDomainImportOperation = "REPLACE"
)

// Values returns all known values for FirewallDomainImportOperation. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (FirewallDomainImportOperation) Values() []FirewallDomainImportOperation {
	return []FirewallDomainImportOperation{
		"REPLACE",
	}
}

type FirewallDomainListStatus string

// Enum values for FirewallDomainListStatus
const (
	FirewallDomainListStatusComplete             FirewallDomainListStatus = "COMPLETE"
	FirewallDomainListStatusCompleteImportFailed FirewallDomainListStatus = "COMPLETE_IMPORT_FAILED"
	FirewallDomainListStatusImporting            FirewallDomainListStatus = "IMPORTING"
	FirewallDomainListStatusDeleting             FirewallDomainListStatus = "DELETING"
	FirewallDomainListStatusUpdating             FirewallDomainListStatus = "UPDATING"
)

// Values returns all known values for FirewallDomainListStatus. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (FirewallDomainListStatus) Values() []FirewallDomainListStatus {
	return []FirewallDomainListStatus{
		"COMPLETE",
		"COMPLETE_IMPORT_FAILED",
		"IMPORTING",
		"DELETING",
		"UPDATING",
	}
}

type FirewallDomainRedirectionAction string

// Enum values for FirewallDomainRedirectionAction
const (
	FirewallDomainRedirectionActionInspectRedirectionDomain FirewallDomainRedirectionAction = "INSPECT_REDIRECTION_DOMAIN"
	FirewallDomainRedirectionActionTrustRedirectionDomain   FirewallDomainRedirectionAction = "TRUST_REDIRECTION_DOMAIN"
)

// Values returns all known values for FirewallDomainRedirectionAction. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (FirewallDomainRedirectionAction) Values() []FirewallDomainRedirectionAction {
	return []FirewallDomainRedirectionAction{
		"INSPECT_REDIRECTION_DOMAIN",
		"TRUST_REDIRECTION_DOMAIN",
	}
}

type FirewallDomainUpdateOperation string

// Enum values for FirewallDomainUpdateOperation
const (
	FirewallDomainUpdateOperationAdd     FirewallDomainUpdateOperation = "ADD"
	FirewallDomainUpdateOperationRemove  FirewallDomainUpdateOperation = "REMOVE"
	FirewallDomainUpdateOperationReplace FirewallDomainUpdateOperation = "REPLACE"
)

// Values returns all known values for FirewallDomainUpdateOperation. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (FirewallDomainUpdateOperation) Values() []FirewallDomainUpdateOperation {
	return []FirewallDomainUpdateOperation{
		"ADD",
		"REMOVE",
		"REPLACE",
	}
}

type FirewallFailOpenStatus string

// Enum values for FirewallFailOpenStatus
const (
	FirewallFailOpenStatusEnabled                 FirewallFailOpenStatus = "ENABLED"
	FirewallFailOpenStatusDisabled                FirewallFailOpenStatus = "DISABLED"
	FirewallFailOpenStatusUseLocalResourceSetting FirewallFailOpenStatus = "USE_LOCAL_RESOURCE_SETTING"
)

// Values returns all known values for FirewallFailOpenStatus. Note that this can
// be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (FirewallFailOpenStatus) Values() []FirewallFailOpenStatus {
	return []FirewallFailOpenStatus{
		"ENABLED",
		"DISABLED",
		"USE_LOCAL_RESOURCE_SETTING",
	}
}

type FirewallRuleGroupAssociationStatus string

// Enum values for FirewallRuleGroupAssociationStatus
const (
	FirewallRuleGroupAssociationStatusComplete FirewallRuleGroupAssociationStatus = "COMPLETE"
	FirewallRuleGroupAssociationStatusDeleting FirewallRuleGroupAssociationStatus = "DELETING"
	FirewallRuleGroupAssociationStatusUpdating FirewallRuleGroupAssociationStatus = "UPDATING"
)

// Values returns all known values for FirewallRuleGroupAssociationStatus. Note
// that this can be expanded in the future, and so it is only as up to date as the
// client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (FirewallRuleGroupAssociationStatus) Values() []FirewallRuleGroupAssociationStatus {
	return []FirewallRuleGroupAssociationStatus{
		"COMPLETE",
		"DELETING",
		"UPDATING",
	}
}

type FirewallRuleGroupStatus string

// Enum values for FirewallRuleGroupStatus
const (
	FirewallRuleGroupStatusComplete FirewallRuleGroupStatus = "COMPLETE"
	FirewallRuleGroupStatusDeleting FirewallRuleGroupStatus = "DELETING"
	FirewallRuleGroupStatusUpdating FirewallRuleGroupStatus = "UPDATING"
)

// Values returns all known values for FirewallRuleGroupStatus. Note that this can
// be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (FirewallRuleGroupStatus) Values() []FirewallRuleGroupStatus {
	return []FirewallRuleGroupStatus{
		"COMPLETE",
		"DELETING",
		"UPDATING",
	}
}

type IpAddressStatus string

// Enum values for IpAddressStatus
const (
	IpAddressStatusCreating               IpAddressStatus = "CREATING"
	IpAddressStatusFailedCreation         IpAddressStatus = "FAILED_CREATION"
	IpAddressStatusAttaching              IpAddressStatus = "ATTACHING"
	IpAddressStatusAttached               IpAddressStatus = "ATTACHED"
	IpAddressStatusRemapDetaching         IpAddressStatus = "REMAP_DETACHING"
	IpAddressStatusRemapAttaching         IpAddressStatus = "REMAP_ATTACHING"
	IpAddressStatusDetaching              IpAddressStatus = "DETACHING"
	IpAddressStatusFailedResourceGone     IpAddressStatus = "FAILED_RESOURCE_GONE"
	IpAddressStatusDeleting               IpAddressStatus = "DELETING"
	IpAddressStatusDeleteFailedFasExpired IpAddressStatus = "DELETE_FAILED_FAS_EXPIRED"
	IpAddressStatusUpdating               IpAddressStatus = "UPDATING"
	IpAddressStatusUpdateFailed           IpAddressStatus = "UPDATE_FAILED"
)

// Values returns all known values for IpAddressStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (IpAddressStatus) Values() []IpAddressStatus {
	return []IpAddressStatus{
		"CREATING",
		"FAILED_CREATION",
		"ATTACHING",
		"ATTACHED",
		"REMAP_DETACHING",
		"REMAP_ATTACHING",
		"DETACHING",
		"FAILED_RESOURCE_GONE",
		"DELETING",
		"DELETE_FAILED_FAS_EXPIRED",
		"UPDATING",
		"UPDATE_FAILED",
	}
}

type MutationProtectionStatus string

// Enum values for MutationProtectionStatus
const (
	MutationProtectionStatusEnabled  MutationProtectionStatus = "ENABLED"
	MutationProtectionStatusDisabled MutationProtectionStatus = "DISABLED"
)

// Values returns all known values for MutationProtectionStatus. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (MutationProtectionStatus) Values() []MutationProtectionStatus {
	return []MutationProtectionStatus{
		"ENABLED",
		"DISABLED",
	}
}

type OutpostResolverStatus string

// Enum values for OutpostResolverStatus
const (
	OutpostResolverStatusCreating       OutpostResolverStatus = "CREATING"
	OutpostResolverStatusOperational    OutpostResolverStatus = "OPERATIONAL"
	OutpostResolverStatusUpdating       OutpostResolverStatus = "UPDATING"
	OutpostResolverStatusDeleting       OutpostResolverStatus = "DELETING"
	OutpostResolverStatusActionNeeded   OutpostResolverStatus = "ACTION_NEEDED"
	OutpostResolverStatusFailedCreation OutpostResolverStatus = "FAILED_CREATION"
	OutpostResolverStatusFailedDeletion OutpostResolverStatus = "FAILED_DELETION"
)

// Values returns all known values for OutpostResolverStatus. Note that this can
// be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (OutpostResolverStatus) Values() []OutpostResolverStatus {
	return []OutpostResolverStatus{
		"CREATING",
		"OPERATIONAL",
		"UPDATING",
		"DELETING",
		"ACTION_NEEDED",
		"FAILED_CREATION",
		"FAILED_DELETION",
	}
}

type Protocol string

// Enum values for Protocol
const (
	ProtocolDoh     Protocol = "DoH"
	ProtocolDo53    Protocol = "Do53"
	ProtocolDohfips Protocol = "DoH-FIPS"
)

// Values returns all known values for Protocol. Note that this can be expanded in
// the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (Protocol) Values() []Protocol {
	return []Protocol{
		"DoH",
		"Do53",
		"DoH-FIPS",
	}
}

type ResolverAutodefinedReverseStatus string

// Enum values for ResolverAutodefinedReverseStatus
const (
	ResolverAutodefinedReverseStatusEnabling                          ResolverAutodefinedReverseStatus = "ENABLING"
	ResolverAutodefinedReverseStatusEnabled                           ResolverAutodefinedReverseStatus = "ENABLED"
	ResolverAutodefinedReverseStatusDisabling                         ResolverAutodefinedReverseStatus = "DISABLING"
	ResolverAutodefinedReverseStatusDisabled                          ResolverAutodefinedReverseStatus = "DISABLED"
	ResolverAutodefinedReverseStatusUpdatingToUseLocalResourceSetting ResolverAutodefinedReverseStatus = "UPDATING_TO_USE_LOCAL_RESOURCE_SETTING"
	ResolverAutodefinedReverseStatusUseLocalResourceSetting           ResolverAutodefinedReverseStatus = "USE_LOCAL_RESOURCE_SETTING"
)

// Values returns all known values for ResolverAutodefinedReverseStatus. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ResolverAutodefinedReverseStatus) Values() []ResolverAutodefinedReverseStatus {
	return []ResolverAutodefinedReverseStatus{
		"ENABLING",
		"ENABLED",
		"DISABLING",
		"DISABLED",
		"UPDATING_TO_USE_LOCAL_RESOURCE_SETTING",
		"USE_LOCAL_RESOURCE_SETTING",
	}
}

type ResolverDNSSECValidationStatus string

// Enum values for ResolverDNSSECValidationStatus
const (
	ResolverDNSSECValidationStatusEnabling                        ResolverDNSSECValidationStatus = "ENABLING"
	ResolverDNSSECValidationStatusEnabled                         ResolverDNSSECValidationStatus = "ENABLED"
	ResolverDNSSECValidationStatusDisabling                       ResolverDNSSECValidationStatus = "DISABLING"
	ResolverDNSSECValidationStatusDisabled                        ResolverDNSSECValidationStatus = "DISABLED"
	ResolverDNSSECValidationStatusUpdateToUseLocalResourceSetting ResolverDNSSECValidationStatus = "UPDATING_TO_USE_LOCAL_RESOURCE_SETTING"
	ResolverDNSSECValidationStatusUseLocalResourceSetting         ResolverDNSSECValidationStatus = "USE_LOCAL_RESOURCE_SETTING"
)

// Values returns all known values for ResolverDNSSECValidationStatus. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ResolverDNSSECValidationStatus) Values() []ResolverDNSSECValidationStatus {
	return []ResolverDNSSECValidationStatus{
		"ENABLING",
		"ENABLED",
		"DISABLING",
		"DISABLED",
		"UPDATING_TO_USE_LOCAL_RESOURCE_SETTING",
		"USE_LOCAL_RESOURCE_SETTING",
	}
}

type ResolverEndpointDirection string

// Enum values for ResolverEndpointDirection
const (
	ResolverEndpointDirectionInbound  ResolverEndpointDirection = "INBOUND"
	ResolverEndpointDirectionOutbound ResolverEndpointDirection = "OUTBOUND"
)

// Values returns all known values for ResolverEndpointDirection. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ResolverEndpointDirection) Values() []ResolverEndpointDirection {
	return []ResolverEndpointDirection{
		"INBOUND",
		"OUTBOUND",
	}
}

type ResolverEndpointStatus string

// Enum values for ResolverEndpointStatus
const (
	ResolverEndpointStatusCreating       ResolverEndpointStatus = "CREATING"
	ResolverEndpointStatusOperational    ResolverEndpointStatus = "OPERATIONAL"
	ResolverEndpointStatusUpdating       ResolverEndpointStatus = "UPDATING"
	ResolverEndpointStatusAutoRecovering ResolverEndpointStatus = "AUTO_RECOVERING"
	ResolverEndpointStatusActionNeeded   ResolverEndpointStatus = "ACTION_NEEDED"
	ResolverEndpointStatusDeleting       ResolverEndpointStatus = "DELETING"
)

// Values returns all known values for ResolverEndpointStatus. Note that this can
// be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ResolverEndpointStatus) Values() []ResolverEndpointStatus {
	return []ResolverEndpointStatus{
		"CREATING",
		"OPERATIONAL",
		"UPDATING",
		"AUTO_RECOVERING",
		"ACTION_NEEDED",
		"DELETING",
	}
}

type ResolverEndpointType string

// Enum values for ResolverEndpointType
const (
	ResolverEndpointTypeIpv6      ResolverEndpointType = "IPV6"
	ResolverEndpointTypeIpv4      ResolverEndpointType = "IPV4"
	ResolverEndpointTypeDualstack ResolverEndpointType = "DUALSTACK"
)

// Values returns all known values for ResolverEndpointType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ResolverEndpointType) Values() []ResolverEndpointType {
	return []ResolverEndpointType{
		"IPV6",
		"IPV4",
		"DUALSTACK",
	}
}

type ResolverQueryLogConfigAssociationError string

// Enum values for ResolverQueryLogConfigAssociationError
const (
	ResolverQueryLogConfigAssociationErrorNone                 ResolverQueryLogConfigAssociationError = "NONE"
	ResolverQueryLogConfigAssociationErrorDestinationNotFound  ResolverQueryLogConfigAssociationError = "DESTINATION_NOT_FOUND"
	ResolverQueryLogConfigAssociationErrorAccessDenied         ResolverQueryLogConfigAssociationError = "ACCESS_DENIED"
	ResolverQueryLogConfigAssociationErrorInternalServiceError ResolverQueryLogConfigAssociationError = "INTERNAL_SERVICE_ERROR"
)

// Values returns all known values for ResolverQueryLogConfigAssociationError.
// Note that this can be expanded in the future, and so it is only as up to date as
// the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ResolverQueryLogConfigAssociationError) Values() []ResolverQueryLogConfigAssociationError {
	return []ResolverQueryLogConfigAssociationError{
		"NONE",
		"DESTINATION_NOT_FOUND",
		"ACCESS_DENIED",
		"INTERNAL_SERVICE_ERROR",
	}
}

type ResolverQueryLogConfigAssociationStatus string

// Enum values for ResolverQueryLogConfigAssociationStatus
const (
	ResolverQueryLogConfigAssociationStatusCreating     ResolverQueryLogConfigAssociationStatus = "CREATING"
	ResolverQueryLogConfigAssociationStatusActive       ResolverQueryLogConfigAssociationStatus = "ACTIVE"
	ResolverQueryLogConfigAssociationStatusActionNeeded ResolverQueryLogConfigAssociationStatus = "ACTION_NEEDED"
	ResolverQueryLogConfigAssociationStatusDeleting     ResolverQueryLogConfigAssociationStatus = "DELETING"
	ResolverQueryLogConfigAssociationStatusFailed       ResolverQueryLogConfigAssociationStatus = "FAILED"
)

// Values returns all known values for ResolverQueryLogConfigAssociationStatus.
// Note that this can be expanded in the future, and so it is only as up to date as
// the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ResolverQueryLogConfigAssociationStatus) Values() []ResolverQueryLogConfigAssociationStatus {
	return []ResolverQueryLogConfigAssociationStatus{
		"CREATING",
		"ACTIVE",
		"ACTION_NEEDED",
		"DELETING",
		"FAILED",
	}
}

type ResolverQueryLogConfigStatus string

// Enum values for ResolverQueryLogConfigStatus
const (
	ResolverQueryLogConfigStatusCreating ResolverQueryLogConfigStatus = "CREATING"
	ResolverQueryLogConfigStatusCreated  ResolverQueryLogConfigStatus = "CREATED"
	ResolverQueryLogConfigStatusDeleting ResolverQueryLogConfigStatus = "DELETING"
	ResolverQueryLogConfigStatusFailed   ResolverQueryLogConfigStatus = "FAILED"
)

// Values returns all known values for ResolverQueryLogConfigStatus. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ResolverQueryLogConfigStatus) Values() []ResolverQueryLogConfigStatus {
	return []ResolverQueryLogConfigStatus{
		"CREATING",
		"CREATED",
		"DELETING",
		"FAILED",
	}
}

type ResolverRuleAssociationStatus string

// Enum values for ResolverRuleAssociationStatus
const (
	ResolverRuleAssociationStatusCreating   ResolverRuleAssociationStatus = "CREATING"
	ResolverRuleAssociationStatusComplete   ResolverRuleAssociationStatus = "COMPLETE"
	ResolverRuleAssociationStatusDeleting   ResolverRuleAssociationStatus = "DELETING"
	ResolverRuleAssociationStatusFailed     ResolverRuleAssociationStatus = "FAILED"
	ResolverRuleAssociationStatusOverridden ResolverRuleAssociationStatus = "OVERRIDDEN"
)

// Values returns all known values for ResolverRuleAssociationStatus. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ResolverRuleAssociationStatus) Values() []ResolverRuleAssociationStatus {
	return []ResolverRuleAssociationStatus{
		"CREATING",
		"COMPLETE",
		"DELETING",
		"FAILED",
		"OVERRIDDEN",
	}
}

type ResolverRuleStatus string

// Enum values for ResolverRuleStatus
const (
	ResolverRuleStatusComplete ResolverRuleStatus = "COMPLETE"
	ResolverRuleStatusDeleting ResolverRuleStatus = "DELETING"
	ResolverRuleStatusUpdating ResolverRuleStatus = "UPDATING"
	ResolverRuleStatusFailed   ResolverRuleStatus = "FAILED"
)

// Values returns all known values for ResolverRuleStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ResolverRuleStatus) Values() []ResolverRuleStatus {
	return []ResolverRuleStatus{
		"COMPLETE",
		"DELETING",
		"UPDATING",
		"FAILED",
	}
}

type RuleTypeOption string

// Enum values for RuleTypeOption
const (
	RuleTypeOptionForward   RuleTypeOption = "FORWARD"
	RuleTypeOptionSystem    RuleTypeOption = "SYSTEM"
	RuleTypeOptionRecursive RuleTypeOption = "RECURSIVE"
)

// Values returns all known values for RuleTypeOption. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (RuleTypeOption) Values() []RuleTypeOption {
	return []RuleTypeOption{
		"FORWARD",
		"SYSTEM",
		"RECURSIVE",
	}
}

type ShareStatus string

// Enum values for ShareStatus
const (
	ShareStatusNotShared    ShareStatus = "NOT_SHARED"
	ShareStatusSharedWithMe ShareStatus = "SHARED_WITH_ME"
	ShareStatusSharedByMe   ShareStatus = "SHARED_BY_ME"
)

// Values returns all known values for ShareStatus. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ShareStatus) Values() []ShareStatus {
	return []ShareStatus{
		"NOT_SHARED",
		"SHARED_WITH_ME",
		"SHARED_BY_ME",
	}
}

type SortOrder string

// Enum values for SortOrder
const (
	SortOrderAscending  SortOrder = "ASCENDING"
	SortOrderDescending SortOrder = "DESCENDING"
)

// Values returns all known values for SortOrder. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (SortOrder) Values() []SortOrder {
	return []SortOrder{
		"ASCENDING",
		"DESCENDING",
	}
}

type Validation string

// Enum values for Validation
const (
	ValidationEnable                  Validation = "ENABLE"
	ValidationDisable                 Validation = "DISABLE"
	ValidationUseLocalResourceSetting Validation = "USE_LOCAL_RESOURCE_SETTING"
)

// Values returns all known values for Validation. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (Validation) Values() []Validation {
	return []Validation{
		"ENABLE",
		"DISABLE",
		"USE_LOCAL_RESOURCE_SETTING",
	}
}
