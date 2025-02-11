// Code generated by smithy-go-codegen DO NOT EDIT.

package deadline

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a limit that manages the distribution of shared resources, such as
// floating licenses. A limit can throttle work assignments, help manage workloads,
// and track current usage. Before you use a limit, you must associate the limit
// with one or more queues.
//
// You must add the amountRequirementName to a step in a job template to declare
// the limit requirement.
func (c *Client) CreateLimit(ctx context.Context, params *CreateLimitInput, optFns ...func(*Options)) (*CreateLimitOutput, error) {
	if params == nil {
		params = &CreateLimitInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateLimit", params, optFns, c.addOperationCreateLimitMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateLimitOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateLimitInput struct {

	// The value that you specify as the name in the amounts field of the
	// hostRequirements in a step of a job template to declare the limit requirement.
	//
	// This member is required.
	AmountRequirementName *string

	// The display name of the limit.
	//
	// This field can store any content. Escape or encode this content before
	// displaying it on a webpage or any other system that might interpret the content
	// of this field.
	//
	// This member is required.
	DisplayName *string

	// The farm ID of the farm that contains the limit.
	//
	// This member is required.
	FarmId *string

	// The maximum number of resources constrained by this limit. When all of the
	// resources are in use, steps that require the limit won't be scheduled until the
	// resource is available.
	//
	// The maxCount must not be 0. If the value is -1, there is no restriction on the
	// number of resources that can be acquired for this limit.
	//
	// This member is required.
	MaxCount *int32

	// The unique token which the server uses to recognize retries of the same request.
	ClientToken *string

	// A description of the limit. A description helps you identify the purpose of the
	// limit.
	//
	// This field can store any content. Escape or encode this content before
	// displaying it on a webpage or any other system that might interpret the content
	// of this field.
	Description *string

	noSmithyDocumentSerde
}

type CreateLimitOutput struct {

	// A unique identifier for the limit. Use this identifier in other operations,
	// such as CreateQueueLimitAssociation and DeleteLimit .
	//
	// This member is required.
	LimitId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateLimitMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateLimit{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateLimit{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateLimit"); err != nil {
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
	if err = addEndpointPrefix_opCreateLimitMiddleware(stack); err != nil {
		return err
	}
	if err = addIdempotencyToken_opCreateLimitMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateLimitValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateLimit(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opCreateLimitMiddleware struct {
}

func (*endpointPrefix_opCreateLimitMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opCreateLimitMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleFinalize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "management." + req.URL.Host

	return next.HandleFinalize(ctx, in)
}
func addEndpointPrefix_opCreateLimitMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opCreateLimitMiddleware{}, "ResolveEndpointV2", middleware.After)
}

type idempotencyToken_initializeOpCreateLimit struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateLimit) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateLimit) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateLimitInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateLimitInput ")
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
func addIdempotencyToken_opCreateLimitMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateLimit{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateLimit(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateLimit",
	}
}
