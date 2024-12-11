// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/document"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Completes registration of a passkey authenticator for the current user. Your
// application provides data from a successful registration request with the data
// from the output of a [StartWebAuthnRegistration].
//
// Authorize this action with a signed-in user's access token. It must include the
// scope aws.cognito.signin.user.admin .
//
// [StartWebAuthnRegistration]: https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_StartWebAuthnRegistration.html
func (c *Client) CompleteWebAuthnRegistration(ctx context.Context, params *CompleteWebAuthnRegistrationInput, optFns ...func(*Options)) (*CompleteWebAuthnRegistrationOutput, error) {
	if params == nil {
		params = &CompleteWebAuthnRegistrationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CompleteWebAuthnRegistration", params, optFns, c.addOperationCompleteWebAuthnRegistrationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CompleteWebAuthnRegistrationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CompleteWebAuthnRegistrationInput struct {

	// A valid access token that Amazon Cognito issued to the user whose passkey
	// registration you want to complete.
	//
	// This member is required.
	AccessToken *string

	// A [RegistrationResponseJSON] public-key credential response from the user's passkey provider.
	//
	// [RegistrationResponseJSON]: https://www.w3.org/TR/webauthn-3/#dictdef-registrationresponsejson
	//
	// This member is required.
	Credential document.Interface

	noSmithyDocumentSerde
}

type CompleteWebAuthnRegistrationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCompleteWebAuthnRegistrationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCompleteWebAuthnRegistration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCompleteWebAuthnRegistration{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CompleteWebAuthnRegistration"); err != nil {
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
	if err = addOpCompleteWebAuthnRegistrationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCompleteWebAuthnRegistration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCompleteWebAuthnRegistration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CompleteWebAuthnRegistration",
	}
}
