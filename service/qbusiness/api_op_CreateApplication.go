// Code generated by smithy-go-codegen DO NOT EDIT.

package qbusiness

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/qbusiness/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an Amazon Q Business application.
//
// There are new tiers for Amazon Q Business. Not all features in Amazon Q
// Business Pro are also available in Amazon Q Business Lite. For information on
// what's included in Amazon Q Business Lite and what's included in Amazon Q
// Business Pro, see [Amazon Q Business tiers]. You must use the Amazon Q Business console to assign
// subscription tiers to users.
//
// An Amazon Q Apps service linked role will be created if it's absent in the
// Amazon Web Services account when QAppsConfiguration is enabled in the request.
// For more information, see [Using service-linked roles for Q Apps].
//
// When you create an application, Amazon Q Business may securely transmit data
// for processing from your selected Amazon Web Services region, but within your
// geography. For more information, see [Cross region inference in Amazon Q Business].
//
// [Amazon Q Business tiers]: https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/tiers.html#user-sub-tiers
// [Using service-linked roles for Q Apps]: https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/using-service-linked-roles-qapps.html
// [Cross region inference in Amazon Q Business]: https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/cross-region-inference.html
func (c *Client) CreateApplication(ctx context.Context, params *CreateApplicationInput, optFns ...func(*Options)) (*CreateApplicationOutput, error) {
	if params == nil {
		params = &CreateApplicationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateApplication", params, optFns, c.addOperationCreateApplicationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateApplicationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateApplicationInput struct {

	// A name for the Amazon Q Business application.
	//
	// This member is required.
	DisplayName *string

	// An option to allow end users to upload files directly during chat.
	AttachmentsConfiguration *types.AttachmentsConfiguration

	// The OIDC client ID for a Amazon Q Business application.
	ClientIdsForOIDC []string

	// A token that you provide to identify the request to create your Amazon Q
	// Business application.
	ClientToken *string

	// A description for the Amazon Q Business application.
	Description *string

	// The identifier of the KMS key that is used to encrypt your data. Amazon Q
	// Business doesn't support asymmetric keys.
	EncryptionConfiguration *types.EncryptionConfiguration

	// The Amazon Resource Name (ARN) of an identity provider being used by an Amazon
	// Q Business application.
	IamIdentityProviderArn *string

	//  The Amazon Resource Name (ARN) of the IAM Identity Center instance you are
	// either creating for—or connecting to—your Amazon Q Business application.
	IdentityCenterInstanceArn *string

	// The authentication type being used by a Amazon Q Business application.
	IdentityType types.IdentityType

	// Configuration information about chat response personalization. For more
	// information, see [Personalizing chat responses]
	//
	// [Personalizing chat responses]: https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/personalizing-chat-responses.html
	PersonalizationConfiguration *types.PersonalizationConfiguration

	// An option to allow end users to create and use Amazon Q Apps in the web
	// experience.
	QAppsConfiguration *types.QAppsConfiguration

	// The Amazon QuickSight configuration for an Amazon Q Business application that
	// uses QuickSight for authentication. This configuration is required if your
	// application uses QuickSight as the identity provider. For more information, see [Creating an Amazon QuickSight integrated application]
	// .
	//
	// [Creating an Amazon QuickSight integrated application]: https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/create-quicksight-integrated-application.html
	QuickSightConfiguration *types.QuickSightConfiguration

	//  The Amazon Resource Name (ARN) of an IAM role with permissions to access your
	// Amazon CloudWatch logs and metrics. If this property is not specified, Amazon Q
	// Business will create a [service linked role (SLR)]and use it as the application's role.
	//
	// [service linked role (SLR)]: https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/using-service-linked-roles.html#slr-permissions
	RoleArn *string

	// A list of key-value pairs that identify or categorize your Amazon Q Business
	// application. You can also use tags to help control access to the application.
	// Tag keys and values can consist of Unicode letters, digits, white space, and any
	// of the following symbols: _ . : / = + - @.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateApplicationOutput struct {

	//  The Amazon Resource Name (ARN) of the Amazon Q Business application.
	ApplicationArn *string

	// The identifier of the Amazon Q Business application.
	ApplicationId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateApplicationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateApplication{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateApplication{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateApplication"); err != nil {
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
	if err = addIdempotencyToken_opCreateApplicationMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateApplicationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateApplication(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateApplication struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateApplication) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateApplication) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateApplicationInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateApplicationInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateApplicationMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateApplication{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateApplication(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateApplication",
	}
}
