// Code generated by smithy-go-codegen DO NOT EDIT.

package kendraranking

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/kendraranking/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a rescore execution plan. A rescore execution plan is an Amazon Kendra
// Intelligent Ranking resource used for provisioning the Rescore API. You set the
// number of capacity units that you require for Amazon Kendra Intelligent Ranking
// to rescore or re-rank a search service's results.
//
// For an example of using the CreateRescoreExecutionPlan API, including using the
// Python and Java SDKs, see [Semantically ranking a search service's results].
//
// [Semantically ranking a search service's results]: https://docs.aws.amazon.com/kendra/latest/dg/search-service-rerank.html
func (c *Client) CreateRescoreExecutionPlan(ctx context.Context, params *CreateRescoreExecutionPlanInput, optFns ...func(*Options)) (*CreateRescoreExecutionPlanOutput, error) {
	if params == nil {
		params = &CreateRescoreExecutionPlanInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateRescoreExecutionPlan", params, optFns, c.addOperationCreateRescoreExecutionPlanMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateRescoreExecutionPlanOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateRescoreExecutionPlanInput struct {

	// A name for the rescore execution plan.
	//
	// This member is required.
	Name *string

	// You can set additional capacity units to meet the needs of your rescore
	// execution plan. You are given a single capacity unit by default. If you want to
	// use the default capacity, you don't set additional capacity units. For more
	// information on the default capacity and additional capacity units, see [Adjusting capacity].
	//
	// [Adjusting capacity]: https://docs.aws.amazon.com/kendra/latest/dg/adjusting-capacity.html
	CapacityUnits *types.CapacityUnitsConfiguration

	// A token that you provide to identify the request to create a rescore execution
	// plan. Multiple calls to the CreateRescoreExecutionPlanRequest API with the same
	// client token will create only one rescore execution plan.
	ClientToken *string

	// A description for the rescore execution plan.
	Description *string

	// A list of key-value pairs that identify or categorize your rescore execution
	// plan. You can also use tags to help control access to the rescore execution
	// plan. Tag keys and values can consist of Unicode letters, digits, white space,
	// and any of the following symbols: _ . : / = + - @.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateRescoreExecutionPlanOutput struct {

	// The Amazon Resource Name (ARN) of the rescore execution plan.
	//
	// This member is required.
	Arn *string

	// The identifier of the rescore execution plan.
	//
	// This member is required.
	Id *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateRescoreExecutionPlanMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpCreateRescoreExecutionPlan{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpCreateRescoreExecutionPlan{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateRescoreExecutionPlan"); err != nil {
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
	if err = addIdempotencyToken_opCreateRescoreExecutionPlanMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateRescoreExecutionPlanValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateRescoreExecutionPlan(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateRescoreExecutionPlan struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateRescoreExecutionPlan) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateRescoreExecutionPlan) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateRescoreExecutionPlanInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateRescoreExecutionPlanInput ")
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
func addIdempotencyToken_opCreateRescoreExecutionPlanMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateRescoreExecutionPlan{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateRescoreExecutionPlan(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateRescoreExecutionPlan",
	}
}
