// Code generated by smithy-go-codegen DO NOT EDIT.

package chimesdkmediapipelines

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/chimesdkmediapipelines/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an Amazon Kinesis Video Stream pool for use with media stream pipelines.
//
// If a meeting uses an opt-in Region as its [MediaRegion], the KVS stream must be in that same
// Region. For example, if a meeting uses the af-south-1 Region, the KVS stream
// must also be in af-south-1 . However, if the meeting uses a Region that AWS
// turns on by default, the KVS stream can be in any available Region, including an
// opt-in Region. For example, if the meeting uses ca-central-1 , the KVS stream
// can be in eu-west-2 , us-east-1 , af-south-1 , or any other Region that the
// Amazon Chime SDK supports.
//
// To learn which AWS Region a meeting uses, call the [GetMeeting] API and use the [MediaRegion] parameter
// from the response.
//
// For more information about opt-in Regions, refer to [Available Regions] in the Amazon Chime SDK
// Developer Guide, and [Specify which AWS Regions your account can use], in the AWS Account Management Reference Guide.
//
// [GetMeeting]: https://docs.aws.amazon.com/chime-sdk/latest/APIReference/API_meeting-chime_GetMeeting.html
// [Specify which AWS Regions your account can use]: https://docs.aws.amazon.com/accounts/latest/reference/manage-acct-regions.html#rande-manage-enable.html
// [Available Regions]: https://docs.aws.amazon.com/chime-sdk/latest/dg/sdk-available-regions.html
// [MediaRegion]: https://docs.aws.amazon.com/chime-sdk/latest/APIReference/API_meeting-chime_CreateMeeting.html#chimesdk-meeting-chime_CreateMeeting-request-MediaRegion
func (c *Client) CreateMediaPipelineKinesisVideoStreamPool(ctx context.Context, params *CreateMediaPipelineKinesisVideoStreamPoolInput, optFns ...func(*Options)) (*CreateMediaPipelineKinesisVideoStreamPoolOutput, error) {
	if params == nil {
		params = &CreateMediaPipelineKinesisVideoStreamPoolInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateMediaPipelineKinesisVideoStreamPool", params, optFns, c.addOperationCreateMediaPipelineKinesisVideoStreamPoolMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateMediaPipelineKinesisVideoStreamPoolOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateMediaPipelineKinesisVideoStreamPoolInput struct {

	// The name of the pool.
	//
	// This member is required.
	PoolName *string

	// The configuration settings for the stream.
	//
	// This member is required.
	StreamConfiguration *types.KinesisVideoStreamConfiguration

	// The token assigned to the client making the request.
	ClientRequestToken *string

	// The tags assigned to the stream pool.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateMediaPipelineKinesisVideoStreamPoolOutput struct {

	// The configuration for applying the streams to the pool.
	KinesisVideoStreamPoolConfiguration *types.KinesisVideoStreamPoolConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateMediaPipelineKinesisVideoStreamPoolMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateMediaPipelineKinesisVideoStreamPool{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateMediaPipelineKinesisVideoStreamPool{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateMediaPipelineKinesisVideoStreamPool"); err != nil {
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
	if err = addIdempotencyToken_opCreateMediaPipelineKinesisVideoStreamPoolMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateMediaPipelineKinesisVideoStreamPoolValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateMediaPipelineKinesisVideoStreamPool(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateMediaPipelineKinesisVideoStreamPool struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateMediaPipelineKinesisVideoStreamPool) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateMediaPipelineKinesisVideoStreamPool) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateMediaPipelineKinesisVideoStreamPoolInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateMediaPipelineKinesisVideoStreamPoolInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateMediaPipelineKinesisVideoStreamPoolMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateMediaPipelineKinesisVideoStreamPool{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateMediaPipelineKinesisVideoStreamPool(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateMediaPipelineKinesisVideoStreamPool",
	}
}
