// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Sets the specified user's password in a user pool. This operation
// administratively sets a temporary or permanent password for a user. With this
// operation, you can bypass self-service password changes and permit immediate
// sign-in with the password that you set. To do this, set Permanent to true .
//
// You can also set a new temporary password in this request, send it to a user,
// and require them to choose a new password on their next sign-in. To do this, set
// Permanent to false .
//
// If the password is temporary, the user's Status becomes FORCE_CHANGE_PASSWORD .
// When the user next tries to sign in, the InitiateAuth or AdminInitiateAuth
// response includes the NEW_PASSWORD_REQUIRED challenge. If the user doesn't sign
// in before the temporary password expires, they can no longer sign in and you
// must repeat this operation to set a temporary or permanent password for them.
//
// After the user sets a new password, or if you set a permanent password, their
// status becomes Confirmed .
//
// AdminSetUserPassword can set a password for the user profile that Amazon
// Cognito creates for third-party federated users. When you set a password, the
// federated user's status changes from EXTERNAL_PROVIDER to CONFIRMED . A user in
// this state can sign in as a federated user, and initiate authentication flows in
// the API like a linked native user. They can also modify their password and
// attributes in token-authenticated API requests like ChangePassword and
// UpdateUserAttributes . As a best security practice and to keep users in sync
// with your external IdP, don't set passwords on federated user profiles. To set
// up a federated user for native sign-in with a linked native user, refer to [Linking federated users to an existing user profile].
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
// [Linking federated users to an existing user profile]: https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-identity-federation-consolidate-users.html
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
// [Signing Amazon Web Services API Requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-signing.html
func (c *Client) AdminSetUserPassword(ctx context.Context, params *AdminSetUserPasswordInput, optFns ...func(*Options)) (*AdminSetUserPasswordOutput, error) {
	if params == nil {
		params = &AdminSetUserPasswordInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AdminSetUserPassword", params, optFns, c.addOperationAdminSetUserPasswordMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AdminSetUserPasswordOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AdminSetUserPasswordInput struct {

	// The new temporary or permanent password that you want to set for the user. You
	// can't remove the password for a user who already has a password so that they can
	// only sign in with passwordless methods. In this scenario, you must create a new
	// user without a password.
	//
	// This member is required.
	Password *string

	// The ID of the user pool where you want to set the user's password.
	//
	// This member is required.
	UserPoolId *string

	// The username of the user that you want to query or modify. The value of this
	// parameter is typically your user's username, but it can be any of their alias
	// attributes. If username isn't an alias attribute in your user pool, this value
	// must be the sub of a local user or the username of a user from a third-party
	// IdP.
	//
	// This member is required.
	Username *string

	// Set to true to set a password that the user can immediately sign in with. Set
	// to false to set a temporary password that the user must change on their next
	// sign-in.
	Permanent bool

	noSmithyDocumentSerde
}

type AdminSetUserPasswordOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAdminSetUserPasswordMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAdminSetUserPassword{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAdminSetUserPassword{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "AdminSetUserPassword"); err != nil {
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
	if err = addOpAdminSetUserPasswordValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAdminSetUserPassword(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAdminSetUserPassword(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "AdminSetUserPassword",
	}
}
