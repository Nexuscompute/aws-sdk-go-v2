// Code generated by smithy-go-codegen DO NOT EDIT.

package workspacesweb

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/workspacesweb/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates browser settings.
func (c *Client) UpdateBrowserSettings(ctx context.Context, params *UpdateBrowserSettingsInput, optFns ...func(*Options)) (*UpdateBrowserSettingsOutput, error) {
	if params == nil {
		params = &UpdateBrowserSettingsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateBrowserSettings", params, optFns, c.addOperationUpdateBrowserSettingsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateBrowserSettingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateBrowserSettingsInput struct {

	// The ARN of the browser settings.
	//
	// This member is required.
	BrowserSettingsArn *string

	// A JSON string containing Chrome Enterprise policies that will be applied to all
	// streaming sessions.
	BrowserPolicy *string

	// A unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request. Idempotency ensures that an API request completes only once.
	// With an idempotent request, if the original request completes successfully,
	// subsequent retries with the same client token return the result from the
	// original successful request.
	//
	// If you do not specify a client token, one is automatically generated by the
	// Amazon Web Services SDK.
	ClientToken *string

	noSmithyDocumentSerde
}

type UpdateBrowserSettingsOutput struct {

	// The browser settings.
	//
	// This member is required.
	BrowserSettings *types.BrowserSettings

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateBrowserSettingsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateBrowserSettings{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateBrowserSettings{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateBrowserSettings"); err != nil {
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
	if err = addIdempotencyToken_opUpdateBrowserSettingsMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpUpdateBrowserSettingsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateBrowserSettings(options.Region), middleware.Before); err != nil {
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
	return nil
}

type idempotencyToken_initializeOpUpdateBrowserSettings struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpUpdateBrowserSettings) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpUpdateBrowserSettings) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*UpdateBrowserSettingsInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *UpdateBrowserSettingsInput ")
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
func addIdempotencyToken_opUpdateBrowserSettingsMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpUpdateBrowserSettings{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opUpdateBrowserSettings(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateBrowserSettings",
	}
}
