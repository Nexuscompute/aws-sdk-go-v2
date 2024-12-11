// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This action might generate an SMS text message. Starting June 1, 2021, US
// telecom carriers require you to register an origination phone number before you
// can send SMS messages to US phone numbers. If you use SMS text messages in
// Amazon Cognito, you must register a phone number with [Amazon Pinpoint]. Amazon Cognito uses the
// registered number automatically. Otherwise, Amazon Cognito users who must
// receive SMS messages might not be able to sign up, activate their accounts, or
// sign in.
//
// If you have never used SMS text messages with Amazon Cognito or any other
// Amazon Web Services service, Amazon Simple Notification Service might place your
// account in the SMS sandbox. In [sandbox mode], you can send messages only to verified phone
// numbers. After you test your app while in the sandbox environment, you can move
// out of the sandbox and into production. For more information, see [SMS message settings for Amazon Cognito user pools]in the Amazon
// Cognito Developer Guide.
//
// Creates a new Amazon Cognito user pool. This operation sets basic and advanced
// configuration options. You can create a user pool in the Amazon Cognito console
// to your preferences and use the output of [DescribeUserPool]to generate requests from that
// baseline.
//
// If you don't provide a value for an attribute, Amazon Cognito sets it to its
// default value.
//
// Amazon Cognito evaluates Identity and Access Management (IAM) policies in
// requests for this API operation. For this operation, you must use IAM
// credentials to authorize requests, and you must grant yourself the corresponding
// IAM permission in a policy.
//
// # Learn more
//
// [Signing Amazon Web Services API Requests]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]
//
// [SMS message settings for Amazon Cognito user pools]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-sms-settings.html
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
// [DescribeUserPool]: https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_DescribeUserPool.html
// [sandbox mode]: https://docs.aws.amazon.com/sns/latest/dg/sns-sms-sandbox.html
// [Signing Amazon Web Services API Requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-signing.html
// [Amazon Pinpoint]: https://console.aws.amazon.com/pinpoint/home/
func (c *Client) CreateUserPool(ctx context.Context, params *CreateUserPoolInput, optFns ...func(*Options)) (*CreateUserPoolOutput, error) {
	if params == nil {
		params = &CreateUserPoolInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateUserPool", params, optFns, c.addOperationCreateUserPoolMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateUserPoolOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the request to create a user pool.
type CreateUserPoolInput struct {

	// A friendlhy name for your user pool.
	//
	// This member is required.
	PoolName *string

	// The available verified method a user can use to recover their password when
	// they call ForgotPassword . You can use this setting to define a preferred method
	// when a user has more than one method available. With this setting, SMS doesn't
	// qualify for a valid password recovery mechanism if the user also has SMS
	// multi-factor authentication (MFA) activated. In the absence of this setting,
	// Amazon Cognito uses the legacy behavior to determine the recovery method where
	// SMS is preferred through email.
	AccountRecoverySetting *types.AccountRecoverySettingType

	// The configuration for [AdminCreateUser] requests. Includes the template for the invitation
	// message for new users, the duration of temporary passwords, and permitting
	// self-service sign-up.
	//
	// [AdminCreateUser]: https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_AdminCreateUser.html
	AdminCreateUserConfig *types.AdminCreateUserConfigType

	// Attributes supported as an alias for this user pool. Possible values:
	// phone_number, email, or preferred_username. For more information about alias
	// attributes, see [Customizing sign-in attributes].
	//
	// [Customizing sign-in attributes]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-settings-attributes.html#user-pool-settings-aliases
	AliasAttributes []types.AliasAttributeType

	// The attributes that you want your user pool to automatically verify. Possible
	// values: email, phone_number. For more information see [Verifying contact information at sign-up].
	//
	// [Verifying contact information at sign-up]: https://docs.aws.amazon.com/cognito/latest/developerguide/signing-up-users-in-your-app.html#allowing-users-to-sign-up-and-confirm-themselves
	AutoVerifiedAttributes []types.VerifiedAttributeType

	// When active, DeletionProtection prevents accidental deletion of your user pool.
	// Before you can delete a user pool that you have protected against deletion, you
	// must deactivate this feature.
	//
	// When you try to delete a protected user pool in a DeleteUserPool API request,
	// Amazon Cognito returns an InvalidParameterException error. To delete a
	// protected user pool, send a new DeleteUserPool request after you deactivate
	// deletion protection in an UpdateUserPool API request.
	DeletionProtection types.DeletionProtectionType

	// The device-remembering configuration for a user pool. Device remembering or
	// device tracking is a "Remember me on this device" option for user pools that
	// perform authentication with the device key of a trusted device in the back end,
	// instead of a user-provided MFA code. For more information about device
	// authentication, see [Working with user devices in your user pool]. A null value indicates that you have deactivated device
	// remembering in your user pool.
	//
	// When you provide a value for any DeviceConfiguration field, you activate the
	// Amazon Cognito device-remembering feature. For more infor
	//
	// [Working with user devices in your user pool]: https://docs.aws.amazon.com/cognito/latest/developerguide/amazon-cognito-user-pools-device-tracking.html
	DeviceConfiguration *types.DeviceConfigurationType

	// The email configuration of your user pool. The email configuration type sets
	// your preferred sending method, Amazon Web Services Region, and sender for
	// messages from your user pool.
	EmailConfiguration *types.EmailConfigurationType

	// This parameter is no longer used. See [VerificationMessageTemplateType].
	//
	// [VerificationMessageTemplateType]: https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_VerificationMessageTemplateType.html
	EmailVerificationMessage *string

	// This parameter is no longer used. See [VerificationMessageTemplateType].
	//
	// [VerificationMessageTemplateType]: https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_VerificationMessageTemplateType.html
	EmailVerificationSubject *string

	// A collection of user pool Lambda triggers. Amazon Cognito invokes triggers at
	// several possible stages of authentication operations. Triggers can modify the
	// outcome of the operations that invoked them.
	LambdaConfig *types.LambdaConfigType

	// Sets multi-factor authentication (MFA) to be on, off, or optional. When ON , all
	// users must set up MFA before they can sign in. When OPTIONAL , your application
	// must make a client-side determination of whether a user wants to register an MFA
	// device. For user pools with adaptive authentication with threat protection,
	// choose OPTIONAL .
	MfaConfiguration types.UserPoolMfaType

	// The password policy and sign-in policy in the user pool. The password policy
	// sets options like password complexity requirements and password history. The
	// sign-in policy sets the options available to applications in [choice-based authentication].
	//
	// [choice-based authentication]: https://docs.aws.amazon.com/cognito/latest/developerguide/authentication-flows-selection-sdk.html#authentication-flows-selection-choice
	Policies *types.UserPoolPolicyType

	// An array of attributes for the new user pool. You can add custom attributes and
	// modify the properties of default attributes. The specifications in this
	// parameter set the required attributes in your user pool. For more information,
	// see [Working with user attributes].
	//
	// [Working with user attributes]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-settings-attributes.html
	Schema []types.SchemaAttributeType

	// A string representing the SMS authentication message.
	SmsAuthenticationMessage *string

	// The SMS configuration with the settings that your Amazon Cognito user pool must
	// use to send an SMS message from your Amazon Web Services account through Amazon
	// Simple Notification Service. To send SMS messages with Amazon SNS in the Amazon
	// Web Services Region that you want, the Amazon Cognito user pool uses an Identity
	// and Access Management (IAM) role in your Amazon Web Services account. For more
	// information see [SMS message settings].
	//
	// [SMS message settings]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-sms-settings.html
	SmsConfiguration *types.SmsConfigurationType

	// This parameter is no longer used. See [VerificationMessageTemplateType].
	//
	// [VerificationMessageTemplateType]: https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_VerificationMessageTemplateType.html
	SmsVerificationMessage *string

	// The settings for updates to user attributes. These settings include the
	// property AttributesRequireVerificationBeforeUpdate , a user-pool setting that
	// tells Amazon Cognito how to handle changes to the value of your users' email
	// address and phone number attributes. For more information, see [Verifying updates to email addresses and phone numbers].
	//
	// [Verifying updates to email addresses and phone numbers]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-settings-email-phone-verification.html#user-pool-settings-verifications-verify-attribute-updates
	UserAttributeUpdateSettings *types.UserAttributeUpdateSettingsType

	// User pool add-ons. Contains settings for activation of advanced security
	// features. To log user security information but take no action, set to AUDIT . To
	// configure automatic security responses to risky traffic to your user pool, set
	// to ENFORCED .
	//
	// For more information, see [Adding advanced security to a user pool].
	//
	// [Adding advanced security to a user pool]: https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pool-settings-advanced-security.html
	UserPoolAddOns *types.UserPoolAddOnsType

	// The tag keys and values to assign to the user pool. A tag is a label that you
	// can use to categorize and manage user pools in different ways, such as by
	// purpose, owner, environment, or other criteria.
	UserPoolTags map[string]string

	// The user pool [feature plan], or tier. This parameter determines the eligibility of the user
	// pool for features like managed login, access-token customization, and threat
	// protection. Defaults to ESSENTIALS .
	//
	// [feature plan]: https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-sign-in-feature-plans.html
	UserPoolTier types.UserPoolTierType

	// Specifies whether a user can use an email address or phone number as a username
	// when they sign up. For more information, see [Customizing sign-in attributes].
	//
	// [Customizing sign-in attributes]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-settings-attributes.html#user-pool-settings-aliases
	UsernameAttributes []types.UsernameAttributeType

	// Sets the case sensitivity option for sign-in usernames. When CaseSensitive is
	// false (case insensitive), users can sign in with any combination of capital and
	// lowercase letters. For example, username , USERNAME , or UserName , or for
	// email, email@example.com or EMaiL@eXamplE.Com . For most use cases, set case
	// sensitivity to false as a best practice. When usernames and email addresses are
	// case insensitive, Amazon Cognito treats any variation in case as the same user,
	// and prevents a case variation from being assigned to the same attribute for a
	// different user.
	//
	// When CaseSensitive is true (case sensitive), Amazon Cognito interprets USERNAME
	// and UserName as distinct users.
	//
	// This configuration is immutable after you set it.
	UsernameConfiguration *types.UsernameConfigurationType

	// The template for the verification message that your user pool delivers to users
	// who set an email address or phone number attribute.
	//
	// Set the email message type that corresponds to your DefaultEmailOption
	// selection. For CONFIRM_WITH_LINK , specify an EmailMessageByLink and leave
	// EmailMessage blank. For CONFIRM_WITH_CODE , specify an EmailMessage and leave
	// EmailMessageByLink blank. When you supply both parameters with either choice,
	// Amazon Cognito returns an error.
	VerificationMessageTemplate *types.VerificationMessageTemplateType

	noSmithyDocumentSerde
}

// Represents the response from the server for the request to create a user pool.
type CreateUserPoolOutput struct {

	// The details of the created user pool.
	UserPool *types.UserPoolType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateUserPoolMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateUserPool{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateUserPool{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateUserPool"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addSpanRetryLoop(stack, options); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = addOpCreateUserPoolValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateUserPool(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	if err = addSpanInitializeStart(stack); err != nil {
		return err
	}
	if err = addSpanInitializeEnd(stack); err != nil {
		return err
	}
	if err = addSpanBuildRequestStart(stack); err != nil {
		return err
	}
	if err = addSpanBuildRequestEnd(stack); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opCreateUserPool(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateUserPool",
	}
}
