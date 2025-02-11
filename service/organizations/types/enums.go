// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type AccessDeniedForDependencyExceptionReason string

// Enum values for AccessDeniedForDependencyExceptionReason
const (
	AccessDeniedForDependencyExceptionReasonAccessDeniedDuringCreateServiceLinkedRole AccessDeniedForDependencyExceptionReason = "ACCESS_DENIED_DURING_CREATE_SERVICE_LINKED_ROLE"
)

// Values returns all known values for AccessDeniedForDependencyExceptionReason.
// Note that this can be expanded in the future, and so it is only as up to date as
// the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (AccessDeniedForDependencyExceptionReason) Values() []AccessDeniedForDependencyExceptionReason {
	return []AccessDeniedForDependencyExceptionReason{
		"ACCESS_DENIED_DURING_CREATE_SERVICE_LINKED_ROLE",
	}
}

type AccountJoinedMethod string

// Enum values for AccountJoinedMethod
const (
	AccountJoinedMethodInvited AccountJoinedMethod = "INVITED"
	AccountJoinedMethodCreated AccountJoinedMethod = "CREATED"
)

// Values returns all known values for AccountJoinedMethod. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (AccountJoinedMethod) Values() []AccountJoinedMethod {
	return []AccountJoinedMethod{
		"INVITED",
		"CREATED",
	}
}

type AccountStatus string

// Enum values for AccountStatus
const (
	AccountStatusActive         AccountStatus = "ACTIVE"
	AccountStatusSuspended      AccountStatus = "SUSPENDED"
	AccountStatusPendingClosure AccountStatus = "PENDING_CLOSURE"
)

// Values returns all known values for AccountStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (AccountStatus) Values() []AccountStatus {
	return []AccountStatus{
		"ACTIVE",
		"SUSPENDED",
		"PENDING_CLOSURE",
	}
}

type ActionType string

// Enum values for ActionType
const (
	ActionTypeInviteAccountToOrganization       ActionType = "INVITE"
	ActionTypeEnableAllFeatures                 ActionType = "ENABLE_ALL_FEATURES"
	ActionTypeApproveAllFeatures                ActionType = "APPROVE_ALL_FEATURES"
	ActionTypeAddOrganizationsServiceLinkedRole ActionType = "ADD_ORGANIZATIONS_SERVICE_LINKED_ROLE"
)

// Values returns all known values for ActionType. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ActionType) Values() []ActionType {
	return []ActionType{
		"INVITE",
		"ENABLE_ALL_FEATURES",
		"APPROVE_ALL_FEATURES",
		"ADD_ORGANIZATIONS_SERVICE_LINKED_ROLE",
	}
}

type ChildType string

// Enum values for ChildType
const (
	ChildTypeAccount            ChildType = "ACCOUNT"
	ChildTypeOrganizationalUnit ChildType = "ORGANIZATIONAL_UNIT"
)

// Values returns all known values for ChildType. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ChildType) Values() []ChildType {
	return []ChildType{
		"ACCOUNT",
		"ORGANIZATIONAL_UNIT",
	}
}

type ConstraintViolationExceptionReason string

// Enum values for ConstraintViolationExceptionReason
const (
	ConstraintViolationExceptionReasonAccountNumberLimitExceeded                             ConstraintViolationExceptionReason = "ACCOUNT_NUMBER_LIMIT_EXCEEDED"
	ConstraintViolationExceptionReasonHandshakeRateLimitExceeded                             ConstraintViolationExceptionReason = "HANDSHAKE_RATE_LIMIT_EXCEEDED"
	ConstraintViolationExceptionReasonOuNumberLimitExceeded                                  ConstraintViolationExceptionReason = "OU_NUMBER_LIMIT_EXCEEDED"
	ConstraintViolationExceptionReasonOuDepthLimitExceeded                                   ConstraintViolationExceptionReason = "OU_DEPTH_LIMIT_EXCEEDED"
	ConstraintViolationExceptionReasonPolicyNumberLimitExceeded                              ConstraintViolationExceptionReason = "POLICY_NUMBER_LIMIT_EXCEEDED"
	ConstraintViolationExceptionReasonPolicyContentLimitExceeded                             ConstraintViolationExceptionReason = "POLICY_CONTENT_LIMIT_EXCEEDED"
	ConstraintViolationExceptionReasonMaxPolicyTypeAttachmentLimitExceeded                   ConstraintViolationExceptionReason = "MAX_POLICY_TYPE_ATTACHMENT_LIMIT_EXCEEDED"
	ConstraintViolationExceptionReasonMinPolicyTypeAttachmentLimitExceeded                   ConstraintViolationExceptionReason = "MIN_POLICY_TYPE_ATTACHMENT_LIMIT_EXCEEDED"
	ConstraintViolationExceptionReasonAccountCannotLeaveOrganization                         ConstraintViolationExceptionReason = "ACCOUNT_CANNOT_LEAVE_ORGANIZATION"
	ConstraintViolationExceptionReasonAccountCannotLeaveWithoutEula                          ConstraintViolationExceptionReason = "ACCOUNT_CANNOT_LEAVE_WITHOUT_EULA"
	ConstraintViolationExceptionReasonAccountCannotLeaveWithoutPhoneVerification             ConstraintViolationExceptionReason = "ACCOUNT_CANNOT_LEAVE_WITHOUT_PHONE_VERIFICATION"
	ConstraintViolationExceptionReasonMasterAccountPaymentInstrumentRequired                 ConstraintViolationExceptionReason = "MASTER_ACCOUNT_PAYMENT_INSTRUMENT_REQUIRED"
	ConstraintViolationExceptionReasonMemberAccountPaymentInstrumentRequired                 ConstraintViolationExceptionReason = "MEMBER_ACCOUNT_PAYMENT_INSTRUMENT_REQUIRED"
	ConstraintViolationExceptionReasonAccountCreationRateLimitExceeded                       ConstraintViolationExceptionReason = "ACCOUNT_CREATION_RATE_LIMIT_EXCEEDED"
	ConstraintViolationExceptionReasonMasterAccountAddressDoesNotMatchMarketplace            ConstraintViolationExceptionReason = "MASTER_ACCOUNT_ADDRESS_DOES_NOT_MATCH_MARKETPLACE"
	ConstraintViolationExceptionReasonMasterAccountMissingContactInfo                        ConstraintViolationExceptionReason = "MASTER_ACCOUNT_MISSING_CONTACT_INFO"
	ConstraintViolationExceptionReasonMasterAccountNotGovcloudEnabled                        ConstraintViolationExceptionReason = "MASTER_ACCOUNT_NOT_GOVCLOUD_ENABLED"
	ConstraintViolationExceptionReasonOrganizationNotInAllFeaturesMode                       ConstraintViolationExceptionReason = "ORGANIZATION_NOT_IN_ALL_FEATURES_MODE"
	ConstraintViolationExceptionReasonCreateOrganizationInBillingModeUnsupportedRegion       ConstraintViolationExceptionReason = "CREATE_ORGANIZATION_IN_BILLING_MODE_UNSUPPORTED_REGION"
	ConstraintViolationExceptionReasonEmailVerificationCodeExpired                           ConstraintViolationExceptionReason = "EMAIL_VERIFICATION_CODE_EXPIRED"
	ConstraintViolationExceptionReasonWaitPeriodActive                                       ConstraintViolationExceptionReason = "WAIT_PERIOD_ACTIVE"
	ConstraintViolationExceptionReasonMaxTagLimitExceeded                                    ConstraintViolationExceptionReason = "MAX_TAG_LIMIT_EXCEEDED"
	ConstraintViolationExceptionReasonTagPolicyViolation                                     ConstraintViolationExceptionReason = "TAG_POLICY_VIOLATION"
	ConstraintViolationExceptionReasonMaxDelegatedAdministratorsForServiceLimitExceeded      ConstraintViolationExceptionReason = "MAX_DELEGATED_ADMINISTRATORS_FOR_SERVICE_LIMIT_EXCEEDED"
	ConstraintViolationExceptionReasonCannotRegisterMasterAsDelegatedAdministrator           ConstraintViolationExceptionReason = "CANNOT_REGISTER_MASTER_AS_DELEGATED_ADMINISTRATOR"
	ConstraintViolationExceptionReasonCannotRemoveDelegatedAdministratorFromOrg              ConstraintViolationExceptionReason = "CANNOT_REMOVE_DELEGATED_ADMINISTRATOR_FROM_ORG"
	ConstraintViolationExceptionReasonDelegatedAdministratorExistsForThisService             ConstraintViolationExceptionReason = "DELEGATED_ADMINISTRATOR_EXISTS_FOR_THIS_SERVICE"
	ConstraintViolationExceptionReasonMasterAccountMissingBusinessLicense                    ConstraintViolationExceptionReason = "MASTER_ACCOUNT_MISSING_BUSINESS_LICENSE"
	ConstraintViolationExceptionReasonCannotCloseManagementAccount                           ConstraintViolationExceptionReason = "CANNOT_CLOSE_MANAGEMENT_ACCOUNT"
	ConstraintViolationExceptionReasonCloseAccountQuotaExceeded                              ConstraintViolationExceptionReason = "CLOSE_ACCOUNT_QUOTA_EXCEEDED"
	ConstraintViolationExceptionReasonCloseAccountRequestsLimitExceeded                      ConstraintViolationExceptionReason = "CLOSE_ACCOUNT_REQUESTS_LIMIT_EXCEEDED"
	ConstraintViolationExceptionReasonServiceAccessNotEnabled                                ConstraintViolationExceptionReason = "SERVICE_ACCESS_NOT_ENABLED"
	ConstraintViolationExceptionReasonInvalidPaymentInstrument                               ConstraintViolationExceptionReason = "INVALID_PAYMENT_INSTRUMENT"
	ConstraintViolationExceptionReasonAccountCreationNotComplete                             ConstraintViolationExceptionReason = "ACCOUNT_CREATION_NOT_COMPLETE"
	ConstraintViolationExceptionReasonCannotRegisterSuspendedAccountAsDelegatedAdministrator ConstraintViolationExceptionReason = "CANNOT_REGISTER_SUSPENDED_ACCOUNT_AS_DELEGATED_ADMINISTRATOR"
	ConstraintViolationExceptionReasonAllFeaturesMigrationOrganizationSizeLimitExceeded      ConstraintViolationExceptionReason = "ALL_FEATURES_MIGRATION_ORGANIZATION_SIZE_LIMIT_EXCEEDED"
)

// Values returns all known values for ConstraintViolationExceptionReason. Note
// that this can be expanded in the future, and so it is only as up to date as the
// client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ConstraintViolationExceptionReason) Values() []ConstraintViolationExceptionReason {
	return []ConstraintViolationExceptionReason{
		"ACCOUNT_NUMBER_LIMIT_EXCEEDED",
		"HANDSHAKE_RATE_LIMIT_EXCEEDED",
		"OU_NUMBER_LIMIT_EXCEEDED",
		"OU_DEPTH_LIMIT_EXCEEDED",
		"POLICY_NUMBER_LIMIT_EXCEEDED",
		"POLICY_CONTENT_LIMIT_EXCEEDED",
		"MAX_POLICY_TYPE_ATTACHMENT_LIMIT_EXCEEDED",
		"MIN_POLICY_TYPE_ATTACHMENT_LIMIT_EXCEEDED",
		"ACCOUNT_CANNOT_LEAVE_ORGANIZATION",
		"ACCOUNT_CANNOT_LEAVE_WITHOUT_EULA",
		"ACCOUNT_CANNOT_LEAVE_WITHOUT_PHONE_VERIFICATION",
		"MASTER_ACCOUNT_PAYMENT_INSTRUMENT_REQUIRED",
		"MEMBER_ACCOUNT_PAYMENT_INSTRUMENT_REQUIRED",
		"ACCOUNT_CREATION_RATE_LIMIT_EXCEEDED",
		"MASTER_ACCOUNT_ADDRESS_DOES_NOT_MATCH_MARKETPLACE",
		"MASTER_ACCOUNT_MISSING_CONTACT_INFO",
		"MASTER_ACCOUNT_NOT_GOVCLOUD_ENABLED",
		"ORGANIZATION_NOT_IN_ALL_FEATURES_MODE",
		"CREATE_ORGANIZATION_IN_BILLING_MODE_UNSUPPORTED_REGION",
		"EMAIL_VERIFICATION_CODE_EXPIRED",
		"WAIT_PERIOD_ACTIVE",
		"MAX_TAG_LIMIT_EXCEEDED",
		"TAG_POLICY_VIOLATION",
		"MAX_DELEGATED_ADMINISTRATORS_FOR_SERVICE_LIMIT_EXCEEDED",
		"CANNOT_REGISTER_MASTER_AS_DELEGATED_ADMINISTRATOR",
		"CANNOT_REMOVE_DELEGATED_ADMINISTRATOR_FROM_ORG",
		"DELEGATED_ADMINISTRATOR_EXISTS_FOR_THIS_SERVICE",
		"MASTER_ACCOUNT_MISSING_BUSINESS_LICENSE",
		"CANNOT_CLOSE_MANAGEMENT_ACCOUNT",
		"CLOSE_ACCOUNT_QUOTA_EXCEEDED",
		"CLOSE_ACCOUNT_REQUESTS_LIMIT_EXCEEDED",
		"SERVICE_ACCESS_NOT_ENABLED",
		"INVALID_PAYMENT_INSTRUMENT",
		"ACCOUNT_CREATION_NOT_COMPLETE",
		"CANNOT_REGISTER_SUSPENDED_ACCOUNT_AS_DELEGATED_ADMINISTRATOR",
		"ALL_FEATURES_MIGRATION_ORGANIZATION_SIZE_LIMIT_EXCEEDED",
	}
}

type CreateAccountFailureReason string

// Enum values for CreateAccountFailureReason
const (
	CreateAccountFailureReasonAccountLimitExceeded                             CreateAccountFailureReason = "ACCOUNT_LIMIT_EXCEEDED"
	CreateAccountFailureReasonEmailAlreadyExists                               CreateAccountFailureReason = "EMAIL_ALREADY_EXISTS"
	CreateAccountFailureReasonInvalidAddress                                   CreateAccountFailureReason = "INVALID_ADDRESS"
	CreateAccountFailureReasonInvalidEmail                                     CreateAccountFailureReason = "INVALID_EMAIL"
	CreateAccountFailureReasonConcurrentAccountModification                    CreateAccountFailureReason = "CONCURRENT_ACCOUNT_MODIFICATION"
	CreateAccountFailureReasonInternalFailure                                  CreateAccountFailureReason = "INTERNAL_FAILURE"
	CreateAccountFailureReasonGovcloudAccountAlreadyExists                     CreateAccountFailureReason = "GOVCLOUD_ACCOUNT_ALREADY_EXISTS"
	CreateAccountFailureReasonMissingBusinessValidation                        CreateAccountFailureReason = "MISSING_BUSINESS_VALIDATION"
	CreateAccountFailureReasonFailedBusinessValidation                         CreateAccountFailureReason = "FAILED_BUSINESS_VALIDATION"
	CreateAccountFailureReasonPendingBusinessVALIDATIONv                       CreateAccountFailureReason = "PENDING_BUSINESS_VALIDATION"
	CreateAccountFailureReasonInvalidIdentityForBusinessValidation             CreateAccountFailureReason = "INVALID_IDENTITY_FOR_BUSINESS_VALIDATION"
	CreateAccountFailureReasonUnknownBusinessValidation                        CreateAccountFailureReason = "UNKNOWN_BUSINESS_VALIDATION"
	CreateAccountFailureReasonMissingPaymentInstrument                         CreateAccountFailureReason = "MISSING_PAYMENT_INSTRUMENT"
	CreateAccountFailureReasonInvalidPaymentInstrument                         CreateAccountFailureReason = "INVALID_PAYMENT_INSTRUMENT"
	CreateAccountFailureReasonUpdateExistingResourcePolicyWithTagsNotSupported CreateAccountFailureReason = "UPDATE_EXISTING_RESOURCE_POLICY_WITH_TAGS_NOT_SUPPORTED"
)

// Values returns all known values for CreateAccountFailureReason. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (CreateAccountFailureReason) Values() []CreateAccountFailureReason {
	return []CreateAccountFailureReason{
		"ACCOUNT_LIMIT_EXCEEDED",
		"EMAIL_ALREADY_EXISTS",
		"INVALID_ADDRESS",
		"INVALID_EMAIL",
		"CONCURRENT_ACCOUNT_MODIFICATION",
		"INTERNAL_FAILURE",
		"GOVCLOUD_ACCOUNT_ALREADY_EXISTS",
		"MISSING_BUSINESS_VALIDATION",
		"FAILED_BUSINESS_VALIDATION",
		"PENDING_BUSINESS_VALIDATION",
		"INVALID_IDENTITY_FOR_BUSINESS_VALIDATION",
		"UNKNOWN_BUSINESS_VALIDATION",
		"MISSING_PAYMENT_INSTRUMENT",
		"INVALID_PAYMENT_INSTRUMENT",
		"UPDATE_EXISTING_RESOURCE_POLICY_WITH_TAGS_NOT_SUPPORTED",
	}
}

type CreateAccountState string

// Enum values for CreateAccountState
const (
	CreateAccountStateInProgress CreateAccountState = "IN_PROGRESS"
	CreateAccountStateSucceeded  CreateAccountState = "SUCCEEDED"
	CreateAccountStateFailed     CreateAccountState = "FAILED"
)

// Values returns all known values for CreateAccountState. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (CreateAccountState) Values() []CreateAccountState {
	return []CreateAccountState{
		"IN_PROGRESS",
		"SUCCEEDED",
		"FAILED",
	}
}

type EffectivePolicyType string

// Enum values for EffectivePolicyType
const (
	EffectivePolicyTypeTagPolicy              EffectivePolicyType = "TAG_POLICY"
	EffectivePolicyTypeBackupPolicy           EffectivePolicyType = "BACKUP_POLICY"
	EffectivePolicyTypeAiservicesOptOutPolicy EffectivePolicyType = "AISERVICES_OPT_OUT_POLICY"
	EffectivePolicyTypeChatbotPolicy          EffectivePolicyType = "CHATBOT_POLICY"
	EffectivePolicyTypeDeclarativePolicyEc2   EffectivePolicyType = "DECLARATIVE_POLICY_EC2"
)

// Values returns all known values for EffectivePolicyType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (EffectivePolicyType) Values() []EffectivePolicyType {
	return []EffectivePolicyType{
		"TAG_POLICY",
		"BACKUP_POLICY",
		"AISERVICES_OPT_OUT_POLICY",
		"CHATBOT_POLICY",
		"DECLARATIVE_POLICY_EC2",
	}
}

type HandshakeConstraintViolationExceptionReason string

// Enum values for HandshakeConstraintViolationExceptionReason
const (
	HandshakeConstraintViolationExceptionReasonAccountNumberLimitExceeded                       HandshakeConstraintViolationExceptionReason = "ACCOUNT_NUMBER_LIMIT_EXCEEDED"
	HandshakeConstraintViolationExceptionReasonHandshakeRateLimitExceeded                       HandshakeConstraintViolationExceptionReason = "HANDSHAKE_RATE_LIMIT_EXCEEDED"
	HandshakeConstraintViolationExceptionReasonAlreadyInAnOrganization                          HandshakeConstraintViolationExceptionReason = "ALREADY_IN_AN_ORGANIZATION"
	HandshakeConstraintViolationExceptionReasonOrganizationAlreadyHasAllFeatures                HandshakeConstraintViolationExceptionReason = "ORGANIZATION_ALREADY_HAS_ALL_FEATURES"
	HandshakeConstraintViolationExceptionReasonOrganizationIsAlreadyPendingAllFeaturesMigration HandshakeConstraintViolationExceptionReason = "ORGANIZATION_IS_ALREADY_PENDING_ALL_FEATURES_MIGRATION"
	HandshakeConstraintViolationExceptionReasonInviteDisabledDuringEnableAllFeatures            HandshakeConstraintViolationExceptionReason = "INVITE_DISABLED_DURING_ENABLE_ALL_FEATURES"
	HandshakeConstraintViolationExceptionReasonPaymentInstrumentRequired                        HandshakeConstraintViolationExceptionReason = "PAYMENT_INSTRUMENT_REQUIRED"
	HandshakeConstraintViolationExceptionReasonOrganizationFromDifferentSellerOfRecord          HandshakeConstraintViolationExceptionReason = "ORGANIZATION_FROM_DIFFERENT_SELLER_OF_RECORD"
	HandshakeConstraintViolationExceptionReasonOrganizationMembershipChangeRateLimitExceeded    HandshakeConstraintViolationExceptionReason = "ORGANIZATION_MEMBERSHIP_CHANGE_RATE_LIMIT_EXCEEDED"
	HandshakeConstraintViolationExceptionReasonManagementAccountEmailNotVerified                HandshakeConstraintViolationExceptionReason = "MANAGEMENT_ACCOUNT_EMAIL_NOT_VERIFIED"
)

// Values returns all known values for
// HandshakeConstraintViolationExceptionReason. Note that this can be expanded in
// the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (HandshakeConstraintViolationExceptionReason) Values() []HandshakeConstraintViolationExceptionReason {
	return []HandshakeConstraintViolationExceptionReason{
		"ACCOUNT_NUMBER_LIMIT_EXCEEDED",
		"HANDSHAKE_RATE_LIMIT_EXCEEDED",
		"ALREADY_IN_AN_ORGANIZATION",
		"ORGANIZATION_ALREADY_HAS_ALL_FEATURES",
		"ORGANIZATION_IS_ALREADY_PENDING_ALL_FEATURES_MIGRATION",
		"INVITE_DISABLED_DURING_ENABLE_ALL_FEATURES",
		"PAYMENT_INSTRUMENT_REQUIRED",
		"ORGANIZATION_FROM_DIFFERENT_SELLER_OF_RECORD",
		"ORGANIZATION_MEMBERSHIP_CHANGE_RATE_LIMIT_EXCEEDED",
		"MANAGEMENT_ACCOUNT_EMAIL_NOT_VERIFIED",
	}
}

type HandshakePartyType string

// Enum values for HandshakePartyType
const (
	HandshakePartyTypeAccount      HandshakePartyType = "ACCOUNT"
	HandshakePartyTypeOrganization HandshakePartyType = "ORGANIZATION"
	HandshakePartyTypeEmail        HandshakePartyType = "EMAIL"
)

// Values returns all known values for HandshakePartyType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (HandshakePartyType) Values() []HandshakePartyType {
	return []HandshakePartyType{
		"ACCOUNT",
		"ORGANIZATION",
		"EMAIL",
	}
}

type HandshakeResourceType string

// Enum values for HandshakeResourceType
const (
	HandshakeResourceTypeAccount                HandshakeResourceType = "ACCOUNT"
	HandshakeResourceTypeOrganization           HandshakeResourceType = "ORGANIZATION"
	HandshakeResourceTypeOrganizationFeatureSet HandshakeResourceType = "ORGANIZATION_FEATURE_SET"
	HandshakeResourceTypeEmail                  HandshakeResourceType = "EMAIL"
	HandshakeResourceTypeMasterEmail            HandshakeResourceType = "MASTER_EMAIL"
	HandshakeResourceTypeMasterName             HandshakeResourceType = "MASTER_NAME"
	HandshakeResourceTypeNotes                  HandshakeResourceType = "NOTES"
	HandshakeResourceTypeParentHandshake        HandshakeResourceType = "PARENT_HANDSHAKE"
)

// Values returns all known values for HandshakeResourceType. Note that this can
// be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (HandshakeResourceType) Values() []HandshakeResourceType {
	return []HandshakeResourceType{
		"ACCOUNT",
		"ORGANIZATION",
		"ORGANIZATION_FEATURE_SET",
		"EMAIL",
		"MASTER_EMAIL",
		"MASTER_NAME",
		"NOTES",
		"PARENT_HANDSHAKE",
	}
}

type HandshakeState string

// Enum values for HandshakeState
const (
	HandshakeStateRequested HandshakeState = "REQUESTED"
	HandshakeStateOpen      HandshakeState = "OPEN"
	HandshakeStateCanceled  HandshakeState = "CANCELED"
	HandshakeStateAccepted  HandshakeState = "ACCEPTED"
	HandshakeStateDeclined  HandshakeState = "DECLINED"
	HandshakeStateExpired   HandshakeState = "EXPIRED"
)

// Values returns all known values for HandshakeState. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (HandshakeState) Values() []HandshakeState {
	return []HandshakeState{
		"REQUESTED",
		"OPEN",
		"CANCELED",
		"ACCEPTED",
		"DECLINED",
		"EXPIRED",
	}
}

type IAMUserAccessToBilling string

// Enum values for IAMUserAccessToBilling
const (
	IAMUserAccessToBillingAllow IAMUserAccessToBilling = "ALLOW"
	IAMUserAccessToBillingDeny  IAMUserAccessToBilling = "DENY"
)

// Values returns all known values for IAMUserAccessToBilling. Note that this can
// be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (IAMUserAccessToBilling) Values() []IAMUserAccessToBilling {
	return []IAMUserAccessToBilling{
		"ALLOW",
		"DENY",
	}
}

type InvalidInputExceptionReason string

// Enum values for InvalidInputExceptionReason
const (
	InvalidInputExceptionReasonInvalidPartyTypeTarget                InvalidInputExceptionReason = "INVALID_PARTY_TYPE_TARGET"
	InvalidInputExceptionReasonInvalidSyntaxOrganization             InvalidInputExceptionReason = "INVALID_SYNTAX_ORGANIZATION_ARN"
	InvalidInputExceptionReasonInvalidSyntaxPolicy                   InvalidInputExceptionReason = "INVALID_SYNTAX_POLICY_ID"
	InvalidInputExceptionReasonInvalidEnum                           InvalidInputExceptionReason = "INVALID_ENUM"
	InvalidInputExceptionReasonInvalidEnumPolicyType                 InvalidInputExceptionReason = "INVALID_ENUM_POLICY_TYPE"
	InvalidInputExceptionReasonInvalidListMember                     InvalidInputExceptionReason = "INVALID_LIST_MEMBER"
	InvalidInputExceptionReasonMaxLengthExceeded                     InvalidInputExceptionReason = "MAX_LENGTH_EXCEEDED"
	InvalidInputExceptionReasonMaxValueExceeded                      InvalidInputExceptionReason = "MAX_VALUE_EXCEEDED"
	InvalidInputExceptionReasonMinLengthExceeded                     InvalidInputExceptionReason = "MIN_LENGTH_EXCEEDED"
	InvalidInputExceptionReasonMinValueExceeded                      InvalidInputExceptionReason = "MIN_VALUE_EXCEEDED"
	InvalidInputExceptionReasonImmutablePolicy                       InvalidInputExceptionReason = "IMMUTABLE_POLICY"
	InvalidInputExceptionReasonInvalidPattern                        InvalidInputExceptionReason = "INVALID_PATTERN"
	InvalidInputExceptionReasonInvalidPatternTargetId                InvalidInputExceptionReason = "INVALID_PATTERN_TARGET_ID"
	InvalidInputExceptionReasonInputRequired                         InvalidInputExceptionReason = "INPUT_REQUIRED"
	InvalidInputExceptionReasonInvalidPaginationToken                InvalidInputExceptionReason = "INVALID_NEXT_TOKEN"
	InvalidInputExceptionReasonMaxFilterLimitExceeded                InvalidInputExceptionReason = "MAX_LIMIT_EXCEEDED_FILTER"
	InvalidInputExceptionReasonMovingAccountBetweenDifferentRoots    InvalidInputExceptionReason = "MOVING_ACCOUNT_BETWEEN_DIFFERENT_ROOTS"
	InvalidInputExceptionReasonInvalidFullNameTarget                 InvalidInputExceptionReason = "INVALID_FULL_NAME_TARGET"
	InvalidInputExceptionReasonUnrecognizedServicePrincipal          InvalidInputExceptionReason = "UNRECOGNIZED_SERVICE_PRINCIPAL"
	InvalidInputExceptionReasonInvalidRoleName                       InvalidInputExceptionReason = "INVALID_ROLE_NAME"
	InvalidInputExceptionReasonInvalidSystemTagsParameter            InvalidInputExceptionReason = "INVALID_SYSTEM_TAGS_PARAMETER"
	InvalidInputExceptionReasonDuplicateTagKey                       InvalidInputExceptionReason = "DUPLICATE_TAG_KEY"
	InvalidInputExceptionReasonTargetNotSupported                    InvalidInputExceptionReason = "TARGET_NOT_SUPPORTED"
	InvalidInputExceptionReasonInvalidEmailAddressTarget             InvalidInputExceptionReason = "INVALID_EMAIL_ADDRESS_TARGET"
	InvalidInputExceptionReasonInvalidResourcePolicyJson             InvalidInputExceptionReason = "INVALID_RESOURCE_POLICY_JSON"
	InvalidInputExceptionReasonInvalidPrincipal                      InvalidInputExceptionReason = "INVALID_PRINCIPAL"
	InvalidInputExceptionReasonUnsupportedActionInResourcePolicy     InvalidInputExceptionReason = "UNSUPPORTED_ACTION_IN_RESOURCE_POLICY"
	InvalidInputExceptionReasonUnsupportedPolicyTypeInResourcePolicy InvalidInputExceptionReason = "UNSUPPORTED_POLICY_TYPE_IN_RESOURCE_POLICY"
	InvalidInputExceptionReasonUnsupportedResourceInResourcePolicy   InvalidInputExceptionReason = "UNSUPPORTED_RESOURCE_IN_RESOURCE_POLICY"
	InvalidInputExceptionReasonNonDetachablePolicy                   InvalidInputExceptionReason = "NON_DETACHABLE_POLICY"
)

// Values returns all known values for InvalidInputExceptionReason. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (InvalidInputExceptionReason) Values() []InvalidInputExceptionReason {
	return []InvalidInputExceptionReason{
		"INVALID_PARTY_TYPE_TARGET",
		"INVALID_SYNTAX_ORGANIZATION_ARN",
		"INVALID_SYNTAX_POLICY_ID",
		"INVALID_ENUM",
		"INVALID_ENUM_POLICY_TYPE",
		"INVALID_LIST_MEMBER",
		"MAX_LENGTH_EXCEEDED",
		"MAX_VALUE_EXCEEDED",
		"MIN_LENGTH_EXCEEDED",
		"MIN_VALUE_EXCEEDED",
		"IMMUTABLE_POLICY",
		"INVALID_PATTERN",
		"INVALID_PATTERN_TARGET_ID",
		"INPUT_REQUIRED",
		"INVALID_NEXT_TOKEN",
		"MAX_LIMIT_EXCEEDED_FILTER",
		"MOVING_ACCOUNT_BETWEEN_DIFFERENT_ROOTS",
		"INVALID_FULL_NAME_TARGET",
		"UNRECOGNIZED_SERVICE_PRINCIPAL",
		"INVALID_ROLE_NAME",
		"INVALID_SYSTEM_TAGS_PARAMETER",
		"DUPLICATE_TAG_KEY",
		"TARGET_NOT_SUPPORTED",
		"INVALID_EMAIL_ADDRESS_TARGET",
		"INVALID_RESOURCE_POLICY_JSON",
		"INVALID_PRINCIPAL",
		"UNSUPPORTED_ACTION_IN_RESOURCE_POLICY",
		"UNSUPPORTED_POLICY_TYPE_IN_RESOURCE_POLICY",
		"UNSUPPORTED_RESOURCE_IN_RESOURCE_POLICY",
		"NON_DETACHABLE_POLICY",
	}
}

type OrganizationFeatureSet string

// Enum values for OrganizationFeatureSet
const (
	OrganizationFeatureSetAll                 OrganizationFeatureSet = "ALL"
	OrganizationFeatureSetConsolidatedBilling OrganizationFeatureSet = "CONSOLIDATED_BILLING"
)

// Values returns all known values for OrganizationFeatureSet. Note that this can
// be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (OrganizationFeatureSet) Values() []OrganizationFeatureSet {
	return []OrganizationFeatureSet{
		"ALL",
		"CONSOLIDATED_BILLING",
	}
}

type ParentType string

// Enum values for ParentType
const (
	ParentTypeRoot               ParentType = "ROOT"
	ParentTypeOrganizationalUnit ParentType = "ORGANIZATIONAL_UNIT"
)

// Values returns all known values for ParentType. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ParentType) Values() []ParentType {
	return []ParentType{
		"ROOT",
		"ORGANIZATIONAL_UNIT",
	}
}

type PolicyType string

// Enum values for PolicyType
const (
	PolicyTypeServiceControlPolicy   PolicyType = "SERVICE_CONTROL_POLICY"
	PolicyTypeResourceControlPolicy  PolicyType = "RESOURCE_CONTROL_POLICY"
	PolicyTypeTagPolicy              PolicyType = "TAG_POLICY"
	PolicyTypeBackupPolicy           PolicyType = "BACKUP_POLICY"
	PolicyTypeAiservicesOptOutPolicy PolicyType = "AISERVICES_OPT_OUT_POLICY"
	PolicyTypeChatbotPolicy          PolicyType = "CHATBOT_POLICY"
	PolicyTypeDeclarativePolicyEc2   PolicyType = "DECLARATIVE_POLICY_EC2"
)

// Values returns all known values for PolicyType. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (PolicyType) Values() []PolicyType {
	return []PolicyType{
		"SERVICE_CONTROL_POLICY",
		"RESOURCE_CONTROL_POLICY",
		"TAG_POLICY",
		"BACKUP_POLICY",
		"AISERVICES_OPT_OUT_POLICY",
		"CHATBOT_POLICY",
		"DECLARATIVE_POLICY_EC2",
	}
}

type PolicyTypeStatus string

// Enum values for PolicyTypeStatus
const (
	PolicyTypeStatusEnabled        PolicyTypeStatus = "ENABLED"
	PolicyTypeStatusPendingEnable  PolicyTypeStatus = "PENDING_ENABLE"
	PolicyTypeStatusPendingDisable PolicyTypeStatus = "PENDING_DISABLE"
)

// Values returns all known values for PolicyTypeStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (PolicyTypeStatus) Values() []PolicyTypeStatus {
	return []PolicyTypeStatus{
		"ENABLED",
		"PENDING_ENABLE",
		"PENDING_DISABLE",
	}
}

type TargetType string

// Enum values for TargetType
const (
	TargetTypeAccount            TargetType = "ACCOUNT"
	TargetTypeOrganizationalUnit TargetType = "ORGANIZATIONAL_UNIT"
	TargetTypeRoot               TargetType = "ROOT"
)

// Values returns all known values for TargetType. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (TargetType) Values() []TargetType {
	return []TargetType{
		"ACCOUNT",
		"ORGANIZATIONAL_UNIT",
		"ROOT",
	}
}
