// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type AccessMethodType string

// Enum values for AccessMethodType
const (
	AccessMethodTypeCaRepository        AccessMethodType = "CA_REPOSITORY"
	AccessMethodTypeResourcePkiManifest AccessMethodType = "RESOURCE_PKI_MANIFEST"
	AccessMethodTypeResourcePkiNotify   AccessMethodType = "RESOURCE_PKI_NOTIFY"
)

// Values returns all known values for AccessMethodType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (AccessMethodType) Values() []AccessMethodType {
	return []AccessMethodType{
		"CA_REPOSITORY",
		"RESOURCE_PKI_MANIFEST",
		"RESOURCE_PKI_NOTIFY",
	}
}

type ActionType string

// Enum values for ActionType
const (
	ActionTypeIssueCertificate ActionType = "IssueCertificate"
	ActionTypeGetCertificate   ActionType = "GetCertificate"
	ActionTypeListPermissions  ActionType = "ListPermissions"
)

// Values returns all known values for ActionType. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ActionType) Values() []ActionType {
	return []ActionType{
		"IssueCertificate",
		"GetCertificate",
		"ListPermissions",
	}
}

type AuditReportResponseFormat string

// Enum values for AuditReportResponseFormat
const (
	AuditReportResponseFormatJson AuditReportResponseFormat = "JSON"
	AuditReportResponseFormatCsv  AuditReportResponseFormat = "CSV"
)

// Values returns all known values for AuditReportResponseFormat. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (AuditReportResponseFormat) Values() []AuditReportResponseFormat {
	return []AuditReportResponseFormat{
		"JSON",
		"CSV",
	}
}

type AuditReportStatus string

// Enum values for AuditReportStatus
const (
	AuditReportStatusCreating AuditReportStatus = "CREATING"
	AuditReportStatusSuccess  AuditReportStatus = "SUCCESS"
	AuditReportStatusFailed   AuditReportStatus = "FAILED"
)

// Values returns all known values for AuditReportStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (AuditReportStatus) Values() []AuditReportStatus {
	return []AuditReportStatus{
		"CREATING",
		"SUCCESS",
		"FAILED",
	}
}

type CertificateAuthorityStatus string

// Enum values for CertificateAuthorityStatus
const (
	CertificateAuthorityStatusCreating           CertificateAuthorityStatus = "CREATING"
	CertificateAuthorityStatusPendingCertificate CertificateAuthorityStatus = "PENDING_CERTIFICATE"
	CertificateAuthorityStatusActive             CertificateAuthorityStatus = "ACTIVE"
	CertificateAuthorityStatusDeleted            CertificateAuthorityStatus = "DELETED"
	CertificateAuthorityStatusDisabled           CertificateAuthorityStatus = "DISABLED"
	CertificateAuthorityStatusExpired            CertificateAuthorityStatus = "EXPIRED"
	CertificateAuthorityStatusFailed             CertificateAuthorityStatus = "FAILED"
)

// Values returns all known values for CertificateAuthorityStatus. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (CertificateAuthorityStatus) Values() []CertificateAuthorityStatus {
	return []CertificateAuthorityStatus{
		"CREATING",
		"PENDING_CERTIFICATE",
		"ACTIVE",
		"DELETED",
		"DISABLED",
		"EXPIRED",
		"FAILED",
	}
}

type CertificateAuthorityType string

// Enum values for CertificateAuthorityType
const (
	CertificateAuthorityTypeRoot        CertificateAuthorityType = "ROOT"
	CertificateAuthorityTypeSubordinate CertificateAuthorityType = "SUBORDINATE"
)

// Values returns all known values for CertificateAuthorityType. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (CertificateAuthorityType) Values() []CertificateAuthorityType {
	return []CertificateAuthorityType{
		"ROOT",
		"SUBORDINATE",
	}
}

type CertificateAuthorityUsageMode string

// Enum values for CertificateAuthorityUsageMode
const (
	CertificateAuthorityUsageModeGeneralPurpose        CertificateAuthorityUsageMode = "GENERAL_PURPOSE"
	CertificateAuthorityUsageModeShortLivedCertificate CertificateAuthorityUsageMode = "SHORT_LIVED_CERTIFICATE"
)

// Values returns all known values for CertificateAuthorityUsageMode. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (CertificateAuthorityUsageMode) Values() []CertificateAuthorityUsageMode {
	return []CertificateAuthorityUsageMode{
		"GENERAL_PURPOSE",
		"SHORT_LIVED_CERTIFICATE",
	}
}

type CrlType string

// Enum values for CrlType
const (
	CrlTypeComplete    CrlType = "COMPLETE"
	CrlTypePartitioned CrlType = "PARTITIONED"
)

// Values returns all known values for CrlType. Note that this can be expanded in
// the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (CrlType) Values() []CrlType {
	return []CrlType{
		"COMPLETE",
		"PARTITIONED",
	}
}

type ExtendedKeyUsageType string

// Enum values for ExtendedKeyUsageType
const (
	ExtendedKeyUsageTypeServerAuth              ExtendedKeyUsageType = "SERVER_AUTH"
	ExtendedKeyUsageTypeClientAuth              ExtendedKeyUsageType = "CLIENT_AUTH"
	ExtendedKeyUsageTypeCodeSigning             ExtendedKeyUsageType = "CODE_SIGNING"
	ExtendedKeyUsageTypeEmailProtection         ExtendedKeyUsageType = "EMAIL_PROTECTION"
	ExtendedKeyUsageTypeTimeStamping            ExtendedKeyUsageType = "TIME_STAMPING"
	ExtendedKeyUsageTypeOcspSigning             ExtendedKeyUsageType = "OCSP_SIGNING"
	ExtendedKeyUsageTypeSmartCardLogin          ExtendedKeyUsageType = "SMART_CARD_LOGIN"
	ExtendedKeyUsageTypeDocumentSigning         ExtendedKeyUsageType = "DOCUMENT_SIGNING"
	ExtendedKeyUsageTypeCertificateTransparency ExtendedKeyUsageType = "CERTIFICATE_TRANSPARENCY"
)

// Values returns all known values for ExtendedKeyUsageType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ExtendedKeyUsageType) Values() []ExtendedKeyUsageType {
	return []ExtendedKeyUsageType{
		"SERVER_AUTH",
		"CLIENT_AUTH",
		"CODE_SIGNING",
		"EMAIL_PROTECTION",
		"TIME_STAMPING",
		"OCSP_SIGNING",
		"SMART_CARD_LOGIN",
		"DOCUMENT_SIGNING",
		"CERTIFICATE_TRANSPARENCY",
	}
}

type FailureReason string

// Enum values for FailureReason
const (
	FailureReasonRequestTimedOut      FailureReason = "REQUEST_TIMED_OUT"
	FailureReasonUnsupportedAlgorithm FailureReason = "UNSUPPORTED_ALGORITHM"
	FailureReasonOther                FailureReason = "OTHER"
)

// Values returns all known values for FailureReason. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (FailureReason) Values() []FailureReason {
	return []FailureReason{
		"REQUEST_TIMED_OUT",
		"UNSUPPORTED_ALGORITHM",
		"OTHER",
	}
}

type KeyAlgorithm string

// Enum values for KeyAlgorithm
const (
	KeyAlgorithmRsa2048      KeyAlgorithm = "RSA_2048"
	KeyAlgorithmRsa4096      KeyAlgorithm = "RSA_4096"
	KeyAlgorithmEcPrime256v1 KeyAlgorithm = "EC_prime256v1"
	KeyAlgorithmEcSecp384r1  KeyAlgorithm = "EC_secp384r1"
	KeyAlgorithmSm2          KeyAlgorithm = "SM2"
)

// Values returns all known values for KeyAlgorithm. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (KeyAlgorithm) Values() []KeyAlgorithm {
	return []KeyAlgorithm{
		"RSA_2048",
		"RSA_4096",
		"EC_prime256v1",
		"EC_secp384r1",
		"SM2",
	}
}

type KeyStorageSecurityStandard string

// Enum values for KeyStorageSecurityStandard
const (
	KeyStorageSecurityStandardFips1402Level2OrHigher KeyStorageSecurityStandard = "FIPS_140_2_LEVEL_2_OR_HIGHER"
	KeyStorageSecurityStandardFips1402Level3OrHigher KeyStorageSecurityStandard = "FIPS_140_2_LEVEL_3_OR_HIGHER"
	KeyStorageSecurityStandardCcpcLevel1OrHigher     KeyStorageSecurityStandard = "CCPC_LEVEL_1_OR_HIGHER"
)

// Values returns all known values for KeyStorageSecurityStandard. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (KeyStorageSecurityStandard) Values() []KeyStorageSecurityStandard {
	return []KeyStorageSecurityStandard{
		"FIPS_140_2_LEVEL_2_OR_HIGHER",
		"FIPS_140_2_LEVEL_3_OR_HIGHER",
		"CCPC_LEVEL_1_OR_HIGHER",
	}
}

type PolicyQualifierId string

// Enum values for PolicyQualifierId
const (
	PolicyQualifierIdCps PolicyQualifierId = "CPS"
)

// Values returns all known values for PolicyQualifierId. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (PolicyQualifierId) Values() []PolicyQualifierId {
	return []PolicyQualifierId{
		"CPS",
	}
}

type ResourceOwner string

// Enum values for ResourceOwner
const (
	ResourceOwnerSelf          ResourceOwner = "SELF"
	ResourceOwnerOtherAccounts ResourceOwner = "OTHER_ACCOUNTS"
)

// Values returns all known values for ResourceOwner. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ResourceOwner) Values() []ResourceOwner {
	return []ResourceOwner{
		"SELF",
		"OTHER_ACCOUNTS",
	}
}

type RevocationReason string

// Enum values for RevocationReason
const (
	RevocationReasonUnspecified                    RevocationReason = "UNSPECIFIED"
	RevocationReasonKeyCompromise                  RevocationReason = "KEY_COMPROMISE"
	RevocationReasonCertificateAuthorityCompromise RevocationReason = "CERTIFICATE_AUTHORITY_COMPROMISE"
	RevocationReasonAffiliationChanged             RevocationReason = "AFFILIATION_CHANGED"
	RevocationReasonSuperseded                     RevocationReason = "SUPERSEDED"
	RevocationReasonCessationOfOperation           RevocationReason = "CESSATION_OF_OPERATION"
	RevocationReasonPrivilegeWithdrawn             RevocationReason = "PRIVILEGE_WITHDRAWN"
	RevocationReasonAACompromise                   RevocationReason = "A_A_COMPROMISE"
)

// Values returns all known values for RevocationReason. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (RevocationReason) Values() []RevocationReason {
	return []RevocationReason{
		"UNSPECIFIED",
		"KEY_COMPROMISE",
		"CERTIFICATE_AUTHORITY_COMPROMISE",
		"AFFILIATION_CHANGED",
		"SUPERSEDED",
		"CESSATION_OF_OPERATION",
		"PRIVILEGE_WITHDRAWN",
		"A_A_COMPROMISE",
	}
}

type S3ObjectAcl string

// Enum values for S3ObjectAcl
const (
	S3ObjectAclPublicRead             S3ObjectAcl = "PUBLIC_READ"
	S3ObjectAclBucketOwnerFullControl S3ObjectAcl = "BUCKET_OWNER_FULL_CONTROL"
)

// Values returns all known values for S3ObjectAcl. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (S3ObjectAcl) Values() []S3ObjectAcl {
	return []S3ObjectAcl{
		"PUBLIC_READ",
		"BUCKET_OWNER_FULL_CONTROL",
	}
}

type SigningAlgorithm string

// Enum values for SigningAlgorithm
const (
	SigningAlgorithmSha256withecdsa SigningAlgorithm = "SHA256WITHECDSA"
	SigningAlgorithmSha384withecdsa SigningAlgorithm = "SHA384WITHECDSA"
	SigningAlgorithmSha512withecdsa SigningAlgorithm = "SHA512WITHECDSA"
	SigningAlgorithmSha256withrsa   SigningAlgorithm = "SHA256WITHRSA"
	SigningAlgorithmSha384withrsa   SigningAlgorithm = "SHA384WITHRSA"
	SigningAlgorithmSha512withrsa   SigningAlgorithm = "SHA512WITHRSA"
	SigningAlgorithmSm3withsm2      SigningAlgorithm = "SM3WITHSM2"
)

// Values returns all known values for SigningAlgorithm. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (SigningAlgorithm) Values() []SigningAlgorithm {
	return []SigningAlgorithm{
		"SHA256WITHECDSA",
		"SHA384WITHECDSA",
		"SHA512WITHECDSA",
		"SHA256WITHRSA",
		"SHA384WITHRSA",
		"SHA512WITHRSA",
		"SM3WITHSM2",
	}
}

type ValidityPeriodType string

// Enum values for ValidityPeriodType
const (
	ValidityPeriodTypeEndDate  ValidityPeriodType = "END_DATE"
	ValidityPeriodTypeAbsolute ValidityPeriodType = "ABSOLUTE"
	ValidityPeriodTypeDays     ValidityPeriodType = "DAYS"
	ValidityPeriodTypeMonths   ValidityPeriodType = "MONTHS"
	ValidityPeriodTypeYears    ValidityPeriodType = "YEARS"
)

// Values returns all known values for ValidityPeriodType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ValidityPeriodType) Values() []ValidityPeriodType {
	return []ValidityPeriodType{
		"END_DATE",
		"ABSOLUTE",
		"DAYS",
		"MONTHS",
		"YEARS",
	}
}
