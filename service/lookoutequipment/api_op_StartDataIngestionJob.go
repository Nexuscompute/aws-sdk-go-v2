// Code generated by smithy-go-codegen DO NOT EDIT.

package lookoutequipment

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/lookoutequipment/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Starts a data ingestion job. Amazon Lookout for Equipment returns the job
// status.
func (c *Client) StartDataIngestionJob(ctx context.Context, params *StartDataIngestionJobInput, optFns ...func(*Options)) (*StartDataIngestionJobOutput, error) {
	if params == nil {
		params = &StartDataIngestionJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartDataIngestionJob", params, optFns, c.addOperationStartDataIngestionJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartDataIngestionJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartDataIngestionJobInput struct {

	//  A unique identifier for the request. If you do not set the client request
	// token, Amazon Lookout for Equipment generates one.
	//
	// This member is required.
	ClientToken *string

	// The name of the dataset being used by the data ingestion job.
	//
	// This member is required.
	DatasetName *string

	//  Specifies information for the input data for the data ingestion job, including
	// dataset S3 location.
	//
	// This member is required.
	IngestionInputConfiguration *types.IngestionInputConfiguration

	//  The Amazon Resource Name (ARN) of a role with permission to access the data
	// source for the data ingestion job.
	//
	// This member is required.
	RoleArn *string

	noSmithyDocumentSerde
}

type StartDataIngestionJobOutput struct {

	// Indicates the job ID of the data ingestion job.
	JobId *string

	// Indicates the status of the StartDataIngestionJob operation.
	Status types.IngestionJobStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartDataIngestionJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpStartDataIngestionJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpStartDataIngestionJob{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "StartDataIngestionJob"); err != nil {
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
	if err = addIdempotencyToken_opStartDataIngestionJobMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpStartDataIngestionJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartDataIngestionJob(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpStartDataIngestionJob struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpStartDataIngestionJob) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpStartDataIngestionJob) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*StartDataIngestionJobInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *StartDataIngestionJobInput ")
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
func addIdempotencyToken_opStartDataIngestionJobMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpStartDataIngestionJob{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opStartDataIngestionJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "StartDataIngestionJob",
	}
}
