// Code generated by smithy-go-codegen DO NOT EDIT.

package datazone

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/datazone/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Starts the metadata generation run.
func (c *Client) StartMetadataGenerationRun(ctx context.Context, params *StartMetadataGenerationRunInput, optFns ...func(*Options)) (*StartMetadataGenerationRunOutput, error) {
	if params == nil {
		params = &StartMetadataGenerationRunInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartMetadataGenerationRun", params, optFns, c.addOperationStartMetadataGenerationRunMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartMetadataGenerationRunOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartMetadataGenerationRunInput struct {

	// The ID of the Amazon DataZone domain where you want to start a metadata
	// generation run.
	//
	// This member is required.
	DomainIdentifier *string

	// The ID of the project that owns the asset for which you want to start a
	// metadata generation run.
	//
	// This member is required.
	OwningProjectIdentifier *string

	// The asset for which you want to start a metadata generation run.
	//
	// This member is required.
	Target *types.MetadataGenerationRunTarget

	// The type of the metadata generation run.
	//
	// This member is required.
	Type types.MetadataGenerationRunType

	// A unique, case-sensitive identifier to ensure idempotency of the request. This
	// field is automatically populated if not provided.
	ClientToken *string

	noSmithyDocumentSerde
}

type StartMetadataGenerationRunOutput struct {

	// The ID of the Amazon DataZone domain in which the metadata generation run was
	// started.
	//
	// This member is required.
	DomainId *string

	// The ID of the metadata generation run.
	//
	// This member is required.
	Id *string

	// The timestamp at which the metadata generation run was started.
	CreatedAt *time.Time

	// The ID of the user who started the metadata generation run.
	CreatedBy *string

	// The ID of the project that owns the asset for which the metadata generation run
	// was started.
	OwningProjectId *string

	// The status of the metadata generation run.
	Status types.MetadataGenerationRunStatus

	// The type of the metadata generation run.
	Type types.MetadataGenerationRunType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartMetadataGenerationRunMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartMetadataGenerationRun{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartMetadataGenerationRun{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "StartMetadataGenerationRun"); err != nil {
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
	if err = addIdempotencyToken_opStartMetadataGenerationRunMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpStartMetadataGenerationRunValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartMetadataGenerationRun(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpStartMetadataGenerationRun struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpStartMetadataGenerationRun) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpStartMetadataGenerationRun) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*StartMetadataGenerationRunInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *StartMetadataGenerationRunInput ")
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
func addIdempotencyToken_opStartMetadataGenerationRunMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpStartMetadataGenerationRun{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opStartMetadataGenerationRun(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "StartMetadataGenerationRun",
	}
}
