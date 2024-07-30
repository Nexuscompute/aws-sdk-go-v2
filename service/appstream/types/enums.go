// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type AccessEndpointType string

// Enum values for AccessEndpointType
const (
	AccessEndpointTypeStreaming AccessEndpointType = "STREAMING"
)

// Values returns all known values for AccessEndpointType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (AccessEndpointType) Values() []AccessEndpointType {
	return []AccessEndpointType{
		"STREAMING",
	}
}

type Action string

// Enum values for Action
const (
	ActionClipboardCopyFromLocalDevice Action = "CLIPBOARD_COPY_FROM_LOCAL_DEVICE"
	ActionClipboardCopyToLocalDevice   Action = "CLIPBOARD_COPY_TO_LOCAL_DEVICE"
	ActionFileUpload                   Action = "FILE_UPLOAD"
	ActionFileDownload                 Action = "FILE_DOWNLOAD"
	ActionPrintingToLocalDevice        Action = "PRINTING_TO_LOCAL_DEVICE"
	ActionDomainPasswordSignin         Action = "DOMAIN_PASSWORD_SIGNIN"
	ActionDomainSmartCardSignin        Action = "DOMAIN_SMART_CARD_SIGNIN"
)

// Values returns all known values for Action. Note that this can be expanded in
// the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (Action) Values() []Action {
	return []Action{
		"CLIPBOARD_COPY_FROM_LOCAL_DEVICE",
		"CLIPBOARD_COPY_TO_LOCAL_DEVICE",
		"FILE_UPLOAD",
		"FILE_DOWNLOAD",
		"PRINTING_TO_LOCAL_DEVICE",
		"DOMAIN_PASSWORD_SIGNIN",
		"DOMAIN_SMART_CARD_SIGNIN",
	}
}

type AppBlockBuilderAttribute string

// Enum values for AppBlockBuilderAttribute
const (
	AppBlockBuilderAttributeIamRoleArn                       AppBlockBuilderAttribute = "IAM_ROLE_ARN"
	AppBlockBuilderAttributeAccessEndpoints                  AppBlockBuilderAttribute = "ACCESS_ENDPOINTS"
	AppBlockBuilderAttributeVpcConfigurationSecurityGroupIds AppBlockBuilderAttribute = "VPC_CONFIGURATION_SECURITY_GROUP_IDS"
)

// Values returns all known values for AppBlockBuilderAttribute. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (AppBlockBuilderAttribute) Values() []AppBlockBuilderAttribute {
	return []AppBlockBuilderAttribute{
		"IAM_ROLE_ARN",
		"ACCESS_ENDPOINTS",
		"VPC_CONFIGURATION_SECURITY_GROUP_IDS",
	}
}

type AppBlockBuilderPlatformType string

// Enum values for AppBlockBuilderPlatformType
const (
	AppBlockBuilderPlatformTypeWindowsServer2019 AppBlockBuilderPlatformType = "WINDOWS_SERVER_2019"
)

// Values returns all known values for AppBlockBuilderPlatformType. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (AppBlockBuilderPlatformType) Values() []AppBlockBuilderPlatformType {
	return []AppBlockBuilderPlatformType{
		"WINDOWS_SERVER_2019",
	}
}

type AppBlockBuilderState string

// Enum values for AppBlockBuilderState
const (
	AppBlockBuilderStateStarting AppBlockBuilderState = "STARTING"
	AppBlockBuilderStateRunning  AppBlockBuilderState = "RUNNING"
	AppBlockBuilderStateStopping AppBlockBuilderState = "STOPPING"
	AppBlockBuilderStateStopped  AppBlockBuilderState = "STOPPED"
)

// Values returns all known values for AppBlockBuilderState. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (AppBlockBuilderState) Values() []AppBlockBuilderState {
	return []AppBlockBuilderState{
		"STARTING",
		"RUNNING",
		"STOPPING",
		"STOPPED",
	}
}

type AppBlockBuilderStateChangeReasonCode string

// Enum values for AppBlockBuilderStateChangeReasonCode
const (
	AppBlockBuilderStateChangeReasonCodeInternalError AppBlockBuilderStateChangeReasonCode = "INTERNAL_ERROR"
)

// Values returns all known values for AppBlockBuilderStateChangeReasonCode. Note
// that this can be expanded in the future, and so it is only as up to date as the
// client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (AppBlockBuilderStateChangeReasonCode) Values() []AppBlockBuilderStateChangeReasonCode {
	return []AppBlockBuilderStateChangeReasonCode{
		"INTERNAL_ERROR",
	}
}

type AppBlockState string

// Enum values for AppBlockState
const (
	AppBlockStateInactive AppBlockState = "INACTIVE"
	AppBlockStateActive   AppBlockState = "ACTIVE"
)

// Values returns all known values for AppBlockState. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (AppBlockState) Values() []AppBlockState {
	return []AppBlockState{
		"INACTIVE",
		"ACTIVE",
	}
}

type ApplicationAttribute string

// Enum values for ApplicationAttribute
const (
	ApplicationAttributeLaunchParameters ApplicationAttribute = "LAUNCH_PARAMETERS"
	ApplicationAttributeWorkingDirectory ApplicationAttribute = "WORKING_DIRECTORY"
)

// Values returns all known values for ApplicationAttribute. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ApplicationAttribute) Values() []ApplicationAttribute {
	return []ApplicationAttribute{
		"LAUNCH_PARAMETERS",
		"WORKING_DIRECTORY",
	}
}

type AppVisibility string

// Enum values for AppVisibility
const (
	AppVisibilityAll        AppVisibility = "ALL"
	AppVisibilityAssociated AppVisibility = "ASSOCIATED"
)

// Values returns all known values for AppVisibility. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (AppVisibility) Values() []AppVisibility {
	return []AppVisibility{
		"ALL",
		"ASSOCIATED",
	}
}

type AuthenticationType string

// Enum values for AuthenticationType
const (
	AuthenticationTypeApi      AuthenticationType = "API"
	AuthenticationTypeSaml     AuthenticationType = "SAML"
	AuthenticationTypeUserpool AuthenticationType = "USERPOOL"
	AuthenticationTypeAwsAd    AuthenticationType = "AWS_AD"
)

// Values returns all known values for AuthenticationType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (AuthenticationType) Values() []AuthenticationType {
	return []AuthenticationType{
		"API",
		"SAML",
		"USERPOOL",
		"AWS_AD",
	}
}

type CertificateBasedAuthStatus string

// Enum values for CertificateBasedAuthStatus
const (
	CertificateBasedAuthStatusDisabled                        CertificateBasedAuthStatus = "DISABLED"
	CertificateBasedAuthStatusEnabled                         CertificateBasedAuthStatus = "ENABLED"
	CertificateBasedAuthStatusEnabledNoDirectoryLoginFallback CertificateBasedAuthStatus = "ENABLED_NO_DIRECTORY_LOGIN_FALLBACK"
)

// Values returns all known values for CertificateBasedAuthStatus. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (CertificateBasedAuthStatus) Values() []CertificateBasedAuthStatus {
	return []CertificateBasedAuthStatus{
		"DISABLED",
		"ENABLED",
		"ENABLED_NO_DIRECTORY_LOGIN_FALLBACK",
	}
}

type FleetAttribute string

// Enum values for FleetAttribute
const (
	FleetAttributeVpcConfiguration                 FleetAttribute = "VPC_CONFIGURATION"
	FleetAttributeVpcConfigurationSecurityGroupIds FleetAttribute = "VPC_CONFIGURATION_SECURITY_GROUP_IDS"
	FleetAttributeDomainJoinInfo                   FleetAttribute = "DOMAIN_JOIN_INFO"
	FleetAttributeIamRoleArn                       FleetAttribute = "IAM_ROLE_ARN"
	FleetAttributeUsbDeviceFilterStrings           FleetAttribute = "USB_DEVICE_FILTER_STRINGS"
	FleetAttributeSessionScriptS3Location          FleetAttribute = "SESSION_SCRIPT_S3_LOCATION"
	FleetAttributeMaxSessionsPerInstance           FleetAttribute = "MAX_SESSIONS_PER_INSTANCE"
)

// Values returns all known values for FleetAttribute. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (FleetAttribute) Values() []FleetAttribute {
	return []FleetAttribute{
		"VPC_CONFIGURATION",
		"VPC_CONFIGURATION_SECURITY_GROUP_IDS",
		"DOMAIN_JOIN_INFO",
		"IAM_ROLE_ARN",
		"USB_DEVICE_FILTER_STRINGS",
		"SESSION_SCRIPT_S3_LOCATION",
		"MAX_SESSIONS_PER_INSTANCE",
	}
}

type FleetErrorCode string

// Enum values for FleetErrorCode
const (
	FleetErrorCodeIamServiceRoleMissingEniDescribeAction            FleetErrorCode = "IAM_SERVICE_ROLE_MISSING_ENI_DESCRIBE_ACTION"
	FleetErrorCodeIamServiceRoleMissingEniCreateAction              FleetErrorCode = "IAM_SERVICE_ROLE_MISSING_ENI_CREATE_ACTION"
	FleetErrorCodeIamServiceRoleMissingEniDeleteAction              FleetErrorCode = "IAM_SERVICE_ROLE_MISSING_ENI_DELETE_ACTION"
	FleetErrorCodeNetworkInterfaceLimitExceeded                     FleetErrorCode = "NETWORK_INTERFACE_LIMIT_EXCEEDED"
	FleetErrorCodeInternalServiceError                              FleetErrorCode = "INTERNAL_SERVICE_ERROR"
	FleetErrorCodeIamServiceRoleIsMissing                           FleetErrorCode = "IAM_SERVICE_ROLE_IS_MISSING"
	FleetErrorCodeMachineRoleIsMissing                              FleetErrorCode = "MACHINE_ROLE_IS_MISSING"
	FleetErrorCodeStsDisabledInRegion                               FleetErrorCode = "STS_DISABLED_IN_REGION"
	FleetErrorCodeSubnetHasInsufficientIpAddresses                  FleetErrorCode = "SUBNET_HAS_INSUFFICIENT_IP_ADDRESSES"
	FleetErrorCodeIamServiceRoleMissingDescribeSubnetAction         FleetErrorCode = "IAM_SERVICE_ROLE_MISSING_DESCRIBE_SUBNET_ACTION"
	FleetErrorCodeSubnetNotFound                                    FleetErrorCode = "SUBNET_NOT_FOUND"
	FleetErrorCodeImageNotFound                                     FleetErrorCode = "IMAGE_NOT_FOUND"
	FleetErrorCodeInvalidSubnetConfiguration                        FleetErrorCode = "INVALID_SUBNET_CONFIGURATION"
	FleetErrorCodeSecurityGroupsNotFound                            FleetErrorCode = "SECURITY_GROUPS_NOT_FOUND"
	FleetErrorCodeIgwNotAttached                                    FleetErrorCode = "IGW_NOT_ATTACHED"
	FleetErrorCodeIamServiceRoleMissingDescribeSecurityGroupsAction FleetErrorCode = "IAM_SERVICE_ROLE_MISSING_DESCRIBE_SECURITY_GROUPS_ACTION"
	FleetErrorCodeFleetStopped                                      FleetErrorCode = "FLEET_STOPPED"
	FleetErrorCodeFleetInstanceProvisioningFailure                  FleetErrorCode = "FLEET_INSTANCE_PROVISIONING_FAILURE"
	FleetErrorCodeDomainJoinErrorFileNotFound                       FleetErrorCode = "DOMAIN_JOIN_ERROR_FILE_NOT_FOUND"
	FleetErrorCodeDomainJoinErrorAccessDenied                       FleetErrorCode = "DOMAIN_JOIN_ERROR_ACCESS_DENIED"
	FleetErrorCodeDomainJoinErrorLogonFailure                       FleetErrorCode = "DOMAIN_JOIN_ERROR_LOGON_FAILURE"
	FleetErrorCodeDomainJoinErrorInvalidParameter                   FleetErrorCode = "DOMAIN_JOIN_ERROR_INVALID_PARAMETER"
	FleetErrorCodeDomainJoinErrorMoreData                           FleetErrorCode = "DOMAIN_JOIN_ERROR_MORE_DATA"
	FleetErrorCodeDomainJoinErrorNoSuchDomain                       FleetErrorCode = "DOMAIN_JOIN_ERROR_NO_SUCH_DOMAIN"
	FleetErrorCodeDomainJoinErrorNotSupported                       FleetErrorCode = "DOMAIN_JOIN_ERROR_NOT_SUPPORTED"
	FleetErrorCodeDomainJoinNerrInvalidWorkgroupName                FleetErrorCode = "DOMAIN_JOIN_NERR_INVALID_WORKGROUP_NAME"
	FleetErrorCodeDomainJoinNerrWorkstationNotStarted               FleetErrorCode = "DOMAIN_JOIN_NERR_WORKSTATION_NOT_STARTED"
	FleetErrorCodeDomainJoinErrorDsMachineAccountQuotaExceeded      FleetErrorCode = "DOMAIN_JOIN_ERROR_DS_MACHINE_ACCOUNT_QUOTA_EXCEEDED"
	FleetErrorCodeDomainJoinNerrPasswordExpired                     FleetErrorCode = "DOMAIN_JOIN_NERR_PASSWORD_EXPIRED"
	FleetErrorCodeDomainJoinInternalServiceError                    FleetErrorCode = "DOMAIN_JOIN_INTERNAL_SERVICE_ERROR"
)

// Values returns all known values for FleetErrorCode. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (FleetErrorCode) Values() []FleetErrorCode {
	return []FleetErrorCode{
		"IAM_SERVICE_ROLE_MISSING_ENI_DESCRIBE_ACTION",
		"IAM_SERVICE_ROLE_MISSING_ENI_CREATE_ACTION",
		"IAM_SERVICE_ROLE_MISSING_ENI_DELETE_ACTION",
		"NETWORK_INTERFACE_LIMIT_EXCEEDED",
		"INTERNAL_SERVICE_ERROR",
		"IAM_SERVICE_ROLE_IS_MISSING",
		"MACHINE_ROLE_IS_MISSING",
		"STS_DISABLED_IN_REGION",
		"SUBNET_HAS_INSUFFICIENT_IP_ADDRESSES",
		"IAM_SERVICE_ROLE_MISSING_DESCRIBE_SUBNET_ACTION",
		"SUBNET_NOT_FOUND",
		"IMAGE_NOT_FOUND",
		"INVALID_SUBNET_CONFIGURATION",
		"SECURITY_GROUPS_NOT_FOUND",
		"IGW_NOT_ATTACHED",
		"IAM_SERVICE_ROLE_MISSING_DESCRIBE_SECURITY_GROUPS_ACTION",
		"FLEET_STOPPED",
		"FLEET_INSTANCE_PROVISIONING_FAILURE",
		"DOMAIN_JOIN_ERROR_FILE_NOT_FOUND",
		"DOMAIN_JOIN_ERROR_ACCESS_DENIED",
		"DOMAIN_JOIN_ERROR_LOGON_FAILURE",
		"DOMAIN_JOIN_ERROR_INVALID_PARAMETER",
		"DOMAIN_JOIN_ERROR_MORE_DATA",
		"DOMAIN_JOIN_ERROR_NO_SUCH_DOMAIN",
		"DOMAIN_JOIN_ERROR_NOT_SUPPORTED",
		"DOMAIN_JOIN_NERR_INVALID_WORKGROUP_NAME",
		"DOMAIN_JOIN_NERR_WORKSTATION_NOT_STARTED",
		"DOMAIN_JOIN_ERROR_DS_MACHINE_ACCOUNT_QUOTA_EXCEEDED",
		"DOMAIN_JOIN_NERR_PASSWORD_EXPIRED",
		"DOMAIN_JOIN_INTERNAL_SERVICE_ERROR",
	}
}

type FleetState string

// Enum values for FleetState
const (
	FleetStateStarting FleetState = "STARTING"
	FleetStateRunning  FleetState = "RUNNING"
	FleetStateStopping FleetState = "STOPPING"
	FleetStateStopped  FleetState = "STOPPED"
)

// Values returns all known values for FleetState. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (FleetState) Values() []FleetState {
	return []FleetState{
		"STARTING",
		"RUNNING",
		"STOPPING",
		"STOPPED",
	}
}

type FleetType string

// Enum values for FleetType
const (
	FleetTypeAlwaysOn FleetType = "ALWAYS_ON"
	FleetTypeOnDemand FleetType = "ON_DEMAND"
	FleetTypeElastic  FleetType = "ELASTIC"
)

// Values returns all known values for FleetType. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (FleetType) Values() []FleetType {
	return []FleetType{
		"ALWAYS_ON",
		"ON_DEMAND",
		"ELASTIC",
	}
}

type ImageBuilderState string

// Enum values for ImageBuilderState
const (
	ImageBuilderStatePending              ImageBuilderState = "PENDING"
	ImageBuilderStateUpdatingAgent        ImageBuilderState = "UPDATING_AGENT"
	ImageBuilderStateRunning              ImageBuilderState = "RUNNING"
	ImageBuilderStateStopping             ImageBuilderState = "STOPPING"
	ImageBuilderStateStopped              ImageBuilderState = "STOPPED"
	ImageBuilderStateRebooting            ImageBuilderState = "REBOOTING"
	ImageBuilderStateSnapshotting         ImageBuilderState = "SNAPSHOTTING"
	ImageBuilderStateDeleting             ImageBuilderState = "DELETING"
	ImageBuilderStateFailed               ImageBuilderState = "FAILED"
	ImageBuilderStateUpdating             ImageBuilderState = "UPDATING"
	ImageBuilderStatePendingQualification ImageBuilderState = "PENDING_QUALIFICATION"
)

// Values returns all known values for ImageBuilderState. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ImageBuilderState) Values() []ImageBuilderState {
	return []ImageBuilderState{
		"PENDING",
		"UPDATING_AGENT",
		"RUNNING",
		"STOPPING",
		"STOPPED",
		"REBOOTING",
		"SNAPSHOTTING",
		"DELETING",
		"FAILED",
		"UPDATING",
		"PENDING_QUALIFICATION",
	}
}

type ImageBuilderStateChangeReasonCode string

// Enum values for ImageBuilderStateChangeReasonCode
const (
	ImageBuilderStateChangeReasonCodeInternalError    ImageBuilderStateChangeReasonCode = "INTERNAL_ERROR"
	ImageBuilderStateChangeReasonCodeImageUnavailable ImageBuilderStateChangeReasonCode = "IMAGE_UNAVAILABLE"
)

// Values returns all known values for ImageBuilderStateChangeReasonCode. Note
// that this can be expanded in the future, and so it is only as up to date as the
// client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ImageBuilderStateChangeReasonCode) Values() []ImageBuilderStateChangeReasonCode {
	return []ImageBuilderStateChangeReasonCode{
		"INTERNAL_ERROR",
		"IMAGE_UNAVAILABLE",
	}
}

type ImageState string

// Enum values for ImageState
const (
	ImageStatePending   ImageState = "PENDING"
	ImageStateAvailable ImageState = "AVAILABLE"
	ImageStateFailed    ImageState = "FAILED"
	ImageStateCopying   ImageState = "COPYING"
	ImageStateDeleting  ImageState = "DELETING"
	ImageStateCreating  ImageState = "CREATING"
	ImageStateImporting ImageState = "IMPORTING"
)

// Values returns all known values for ImageState. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ImageState) Values() []ImageState {
	return []ImageState{
		"PENDING",
		"AVAILABLE",
		"FAILED",
		"COPYING",
		"DELETING",
		"CREATING",
		"IMPORTING",
	}
}

type ImageStateChangeReasonCode string

// Enum values for ImageStateChangeReasonCode
const (
	ImageStateChangeReasonCodeInternalError            ImageStateChangeReasonCode = "INTERNAL_ERROR"
	ImageStateChangeReasonCodeImageBuilderNotAvailable ImageStateChangeReasonCode = "IMAGE_BUILDER_NOT_AVAILABLE"
	ImageStateChangeReasonCodeImageCopyFailure         ImageStateChangeReasonCode = "IMAGE_COPY_FAILURE"
)

// Values returns all known values for ImageStateChangeReasonCode. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ImageStateChangeReasonCode) Values() []ImageStateChangeReasonCode {
	return []ImageStateChangeReasonCode{
		"INTERNAL_ERROR",
		"IMAGE_BUILDER_NOT_AVAILABLE",
		"IMAGE_COPY_FAILURE",
	}
}

type MessageAction string

// Enum values for MessageAction
const (
	MessageActionSuppress MessageAction = "SUPPRESS"
	MessageActionResend   MessageAction = "RESEND"
)

// Values returns all known values for MessageAction. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (MessageAction) Values() []MessageAction {
	return []MessageAction{
		"SUPPRESS",
		"RESEND",
	}
}

type PackagingType string

// Enum values for PackagingType
const (
	PackagingTypeCustom     PackagingType = "CUSTOM"
	PackagingTypeAppstream2 PackagingType = "APPSTREAM2"
)

// Values returns all known values for PackagingType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (PackagingType) Values() []PackagingType {
	return []PackagingType{
		"CUSTOM",
		"APPSTREAM2",
	}
}

type Permission string

// Enum values for Permission
const (
	PermissionEnabled  Permission = "ENABLED"
	PermissionDisabled Permission = "DISABLED"
)

// Values returns all known values for Permission. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (Permission) Values() []Permission {
	return []Permission{
		"ENABLED",
		"DISABLED",
	}
}

type PlatformType string

// Enum values for PlatformType
const (
	PlatformTypeWindows           PlatformType = "WINDOWS"
	PlatformTypeWindowsServer2016 PlatformType = "WINDOWS_SERVER_2016"
	PlatformTypeWindowsServer2019 PlatformType = "WINDOWS_SERVER_2019"
	PlatformTypeWindowsServer2022 PlatformType = "WINDOWS_SERVER_2022"
	PlatformTypeAmazonLinux2      PlatformType = "AMAZON_LINUX2"
	PlatformTypeRhel8             PlatformType = "RHEL8"
)

// Values returns all known values for PlatformType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (PlatformType) Values() []PlatformType {
	return []PlatformType{
		"WINDOWS",
		"WINDOWS_SERVER_2016",
		"WINDOWS_SERVER_2019",
		"WINDOWS_SERVER_2022",
		"AMAZON_LINUX2",
		"RHEL8",
	}
}

type PreferredProtocol string

// Enum values for PreferredProtocol
const (
	PreferredProtocolTcp PreferredProtocol = "TCP"
	PreferredProtocolUdp PreferredProtocol = "UDP"
)

// Values returns all known values for PreferredProtocol. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (PreferredProtocol) Values() []PreferredProtocol {
	return []PreferredProtocol{
		"TCP",
		"UDP",
	}
}

type SessionConnectionState string

// Enum values for SessionConnectionState
const (
	SessionConnectionStateConnected    SessionConnectionState = "CONNECTED"
	SessionConnectionStateNotConnected SessionConnectionState = "NOT_CONNECTED"
)

// Values returns all known values for SessionConnectionState. Note that this can
// be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (SessionConnectionState) Values() []SessionConnectionState {
	return []SessionConnectionState{
		"CONNECTED",
		"NOT_CONNECTED",
	}
}

type SessionState string

// Enum values for SessionState
const (
	SessionStateActive  SessionState = "ACTIVE"
	SessionStatePending SessionState = "PENDING"
	SessionStateExpired SessionState = "EXPIRED"
)

// Values returns all known values for SessionState. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (SessionState) Values() []SessionState {
	return []SessionState{
		"ACTIVE",
		"PENDING",
		"EXPIRED",
	}
}

type StackAttribute string

// Enum values for StackAttribute
const (
	StackAttributeStorageConnectors           StackAttribute = "STORAGE_CONNECTORS"
	StackAttributeStorageConnectorHomefolders StackAttribute = "STORAGE_CONNECTOR_HOMEFOLDERS"
	StackAttributeStorageConnectorGoogleDrive StackAttribute = "STORAGE_CONNECTOR_GOOGLE_DRIVE"
	StackAttributeStorageConnectorOneDrive    StackAttribute = "STORAGE_CONNECTOR_ONE_DRIVE"
	StackAttributeRedirectUrl                 StackAttribute = "REDIRECT_URL"
	StackAttributeFeedbackUrl                 StackAttribute = "FEEDBACK_URL"
	StackAttributeThemeName                   StackAttribute = "THEME_NAME"
	StackAttributeUserSettings                StackAttribute = "USER_SETTINGS"
	StackAttributeEmbedHostDomains            StackAttribute = "EMBED_HOST_DOMAINS"
	StackAttributeIamRoleArn                  StackAttribute = "IAM_ROLE_ARN"
	StackAttributeAccessEndpoints             StackAttribute = "ACCESS_ENDPOINTS"
	StackAttributeStreamingExperienceSettings StackAttribute = "STREAMING_EXPERIENCE_SETTINGS"
)

// Values returns all known values for StackAttribute. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (StackAttribute) Values() []StackAttribute {
	return []StackAttribute{
		"STORAGE_CONNECTORS",
		"STORAGE_CONNECTOR_HOMEFOLDERS",
		"STORAGE_CONNECTOR_GOOGLE_DRIVE",
		"STORAGE_CONNECTOR_ONE_DRIVE",
		"REDIRECT_URL",
		"FEEDBACK_URL",
		"THEME_NAME",
		"USER_SETTINGS",
		"EMBED_HOST_DOMAINS",
		"IAM_ROLE_ARN",
		"ACCESS_ENDPOINTS",
		"STREAMING_EXPERIENCE_SETTINGS",
	}
}

type StackErrorCode string

// Enum values for StackErrorCode
const (
	StackErrorCodeStorageConnectorError StackErrorCode = "STORAGE_CONNECTOR_ERROR"
	StackErrorCodeInternalServiceError  StackErrorCode = "INTERNAL_SERVICE_ERROR"
)

// Values returns all known values for StackErrorCode. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (StackErrorCode) Values() []StackErrorCode {
	return []StackErrorCode{
		"STORAGE_CONNECTOR_ERROR",
		"INTERNAL_SERVICE_ERROR",
	}
}

type StorageConnectorType string

// Enum values for StorageConnectorType
const (
	StorageConnectorTypeHomefolders StorageConnectorType = "HOMEFOLDERS"
	StorageConnectorTypeGoogleDrive StorageConnectorType = "GOOGLE_DRIVE"
	StorageConnectorTypeOneDrive    StorageConnectorType = "ONE_DRIVE"
)

// Values returns all known values for StorageConnectorType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (StorageConnectorType) Values() []StorageConnectorType {
	return []StorageConnectorType{
		"HOMEFOLDERS",
		"GOOGLE_DRIVE",
		"ONE_DRIVE",
	}
}

type StreamView string

// Enum values for StreamView
const (
	StreamViewApp     StreamView = "APP"
	StreamViewDesktop StreamView = "DESKTOP"
)

// Values returns all known values for StreamView. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (StreamView) Values() []StreamView {
	return []StreamView{
		"APP",
		"DESKTOP",
	}
}

type UsageReportExecutionErrorCode string

// Enum values for UsageReportExecutionErrorCode
const (
	UsageReportExecutionErrorCodeResourceNotFound     UsageReportExecutionErrorCode = "RESOURCE_NOT_FOUND"
	UsageReportExecutionErrorCodeAccessDenied         UsageReportExecutionErrorCode = "ACCESS_DENIED"
	UsageReportExecutionErrorCodeInternalServiceError UsageReportExecutionErrorCode = "INTERNAL_SERVICE_ERROR"
)

// Values returns all known values for UsageReportExecutionErrorCode. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (UsageReportExecutionErrorCode) Values() []UsageReportExecutionErrorCode {
	return []UsageReportExecutionErrorCode{
		"RESOURCE_NOT_FOUND",
		"ACCESS_DENIED",
		"INTERNAL_SERVICE_ERROR",
	}
}

type UsageReportSchedule string

// Enum values for UsageReportSchedule
const (
	UsageReportScheduleDaily UsageReportSchedule = "DAILY"
)

// Values returns all known values for UsageReportSchedule. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (UsageReportSchedule) Values() []UsageReportSchedule {
	return []UsageReportSchedule{
		"DAILY",
	}
}

type UserStackAssociationErrorCode string

// Enum values for UserStackAssociationErrorCode
const (
	UserStackAssociationErrorCodeStackNotFound     UserStackAssociationErrorCode = "STACK_NOT_FOUND"
	UserStackAssociationErrorCodeUserNameNotFound  UserStackAssociationErrorCode = "USER_NAME_NOT_FOUND"
	UserStackAssociationErrorCodeDirectoryNotFound UserStackAssociationErrorCode = "DIRECTORY_NOT_FOUND"
	UserStackAssociationErrorCodeInternalError     UserStackAssociationErrorCode = "INTERNAL_ERROR"
)

// Values returns all known values for UserStackAssociationErrorCode. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (UserStackAssociationErrorCode) Values() []UserStackAssociationErrorCode {
	return []UserStackAssociationErrorCode{
		"STACK_NOT_FOUND",
		"USER_NAME_NOT_FOUND",
		"DIRECTORY_NOT_FOUND",
		"INTERNAL_ERROR",
	}
}

type VisibilityType string

// Enum values for VisibilityType
const (
	VisibilityTypePublic  VisibilityType = "PUBLIC"
	VisibilityTypePrivate VisibilityType = "PRIVATE"
	VisibilityTypeShared  VisibilityType = "SHARED"
)

// Values returns all known values for VisibilityType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (VisibilityType) Values() []VisibilityType {
	return []VisibilityType{
		"PUBLIC",
		"PRIVATE",
		"SHARED",
	}
}
