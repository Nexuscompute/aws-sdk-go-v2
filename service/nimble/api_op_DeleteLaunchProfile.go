// Code generated by smithy-go-codegen DO NOT EDIT.

package nimble

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/nimble/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Permanently delete a launch profile.
//
// Deprecated: AWS has deprecated this service. It is no longer available for use.
func (c *Client) DeleteLaunchProfile(ctx context.Context, params *DeleteLaunchProfileInput, optFns ...func(*Options)) (*DeleteLaunchProfileOutput, error) {
	if params == nil {
		params = &DeleteLaunchProfileInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteLaunchProfile", params, optFns, c.addOperationDeleteLaunchProfileMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteLaunchProfileOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteLaunchProfileInput struct {

	// The ID of the launch profile used to control access from the streaming session.
	//
	// This member is required.
	//
	// Deprecated: AWS has deprecated this service. It is no longer available for use.
	LaunchProfileId *string

	// The studio ID.
	//
	// This member is required.
	//
	// Deprecated: AWS has deprecated this service. It is no longer available for use.
	StudioId *string

	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request. If you don’t specify a client token, the Amazon Web Services SDK
	// automatically generates a client token and uses it for the request to ensure
	// idempotency.
	//
	// Deprecated: AWS has deprecated this service. It is no longer available for use.
	ClientToken *string

	noSmithyDocumentSerde
}

type DeleteLaunchProfileOutput struct {

	// The launch profile.
	//
	// Deprecated: AWS has deprecated this service. It is no longer available for use.
	LaunchProfile *types.LaunchProfile

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteLaunchProfileMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteLaunchProfile{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteLaunchProfile{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeleteLaunchProfile"); err != nil {
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
	if err = addIdempotencyToken_opDeleteLaunchProfileMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpDeleteLaunchProfileValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteLaunchProfile(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpDeleteLaunchProfile struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpDeleteLaunchProfile) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpDeleteLaunchProfile) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*DeleteLaunchProfileInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *DeleteLaunchProfileInput ")
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
func addIdempotencyToken_opDeleteLaunchProfileMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpDeleteLaunchProfile{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opDeleteLaunchProfile(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeleteLaunchProfile",
	}
}
