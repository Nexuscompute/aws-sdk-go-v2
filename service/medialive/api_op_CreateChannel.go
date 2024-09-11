// Code generated by smithy-go-codegen DO NOT EDIT.

package medialive

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/medialive/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new channel
func (c *Client) CreateChannel(ctx context.Context, params *CreateChannelInput, optFns ...func(*Options)) (*CreateChannelOutput, error) {
	if params == nil {
		params = &CreateChannelInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateChannel", params, optFns, c.addOperationCreateChannelMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateChannelOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to create a channel
type CreateChannelInput struct {

	// The Elemental Anywhere settings for this channel.
	AnywhereSettings *types.AnywhereSettings

	// Specification of CDI inputs for this channel
	CdiInputSpecification *types.CdiInputSpecification

	// The class for this channel. STANDARD for a channel with two pipelines or
	// SINGLE_PIPELINE for a channel with one pipeline.
	ChannelClass types.ChannelClass

	// Placeholder documentation for __listOfOutputDestination
	Destinations []types.OutputDestination

	// Encoder Settings
	EncoderSettings *types.EncoderSettings

	// List of input attachments for channel.
	InputAttachments []types.InputAttachment

	// Specification of network and file inputs for this channel
	InputSpecification *types.InputSpecification

	// The log level to write to CloudWatch Logs.
	LogLevel types.LogLevel

	// Maintenance settings for this channel.
	Maintenance *types.MaintenanceCreateSettings

	// Name of channel.
	Name *string

	// Unique request ID to be specified. This is needed to prevent retries from
	// creating multiple resources.
	RequestId *string

	// Deprecated field that's only usable by whitelisted customers.
	//
	// Deprecated: This member has been deprecated.
	Reserved *string

	// An optional Amazon Resource Name (ARN) of the role to assume when running the
	// Channel.
	RoleArn *string

	// A collection of key-value pairs.
	Tags map[string]string

	// Settings for the VPC outputs
	Vpc *types.VpcOutputSettings

	noSmithyDocumentSerde
}

// Placeholder documentation for CreateChannelResponse
type CreateChannelOutput struct {

	// Placeholder documentation for Channel
	Channel *types.Channel

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateChannelMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateChannel{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateChannel{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateChannel"); err != nil {
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
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = addIdempotencyToken_opCreateChannelMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateChannelValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateChannel(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateChannel struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateChannel) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateChannel) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateChannelInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateChannelInput ")
	}

	if input.RequestId == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.RequestId = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateChannelMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateChannel{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateChannel(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateChannel",
	}
}
