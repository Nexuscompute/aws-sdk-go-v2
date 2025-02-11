// Code generated by smithy-go-codegen DO NOT EDIT.

package imagebuilder

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/imagebuilder/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// List resources that the runtime instance of the image lifecycle identified for
// lifecycle actions.
func (c *Client) ListLifecycleExecutionResources(ctx context.Context, params *ListLifecycleExecutionResourcesInput, optFns ...func(*Options)) (*ListLifecycleExecutionResourcesOutput, error) {
	if params == nil {
		params = &ListLifecycleExecutionResourcesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListLifecycleExecutionResources", params, optFns, c.addOperationListLifecycleExecutionResourcesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListLifecycleExecutionResourcesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListLifecycleExecutionResourcesInput struct {

	// Use the unique identifier for a runtime instance of the lifecycle policy to get
	// runtime details.
	//
	// This member is required.
	LifecycleExecutionId *string

	// The maximum items to return in a request.
	MaxResults *int32

	// A token to specify where to start paginating. This is the nextToken from a
	// previously truncated response.
	NextToken *string

	// You can leave this empty to get a list of Image Builder resources that were
	// identified for lifecycle actions.
	//
	// To get a list of associated resources that are impacted for an individual
	// resource (the parent), specify its Amazon Resource Name (ARN). Associated
	// resources are produced from your image and distributed when you run a build,
	// such as AMIs or container images stored in ECR repositories.
	ParentResourceId *string

	noSmithyDocumentSerde
}

type ListLifecycleExecutionResourcesOutput struct {

	// Runtime details for the specified runtime instance of the lifecycle policy.
	LifecycleExecutionId *string

	// The current state of the lifecycle runtime instance.
	LifecycleExecutionState *types.LifecycleExecutionState

	// The next token used for paginated responses. When this field isn't empty, there
	// are additional elements that the service hasn't included in this request. Use
	// this token with the next request to retrieve additional objects.
	NextToken *string

	// A list of resources that were identified for lifecycle actions.
	Resources []types.LifecycleExecutionResource

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListLifecycleExecutionResourcesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListLifecycleExecutionResources{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListLifecycleExecutionResources{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListLifecycleExecutionResources"); err != nil {
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
	if err = addOpListLifecycleExecutionResourcesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListLifecycleExecutionResources(options.Region), middleware.Before); err != nil {
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

// ListLifecycleExecutionResourcesPaginatorOptions is the paginator options for
// ListLifecycleExecutionResources
type ListLifecycleExecutionResourcesPaginatorOptions struct {
	// The maximum items to return in a request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListLifecycleExecutionResourcesPaginator is a paginator for
// ListLifecycleExecutionResources
type ListLifecycleExecutionResourcesPaginator struct {
	options   ListLifecycleExecutionResourcesPaginatorOptions
	client    ListLifecycleExecutionResourcesAPIClient
	params    *ListLifecycleExecutionResourcesInput
	nextToken *string
	firstPage bool
}

// NewListLifecycleExecutionResourcesPaginator returns a new
// ListLifecycleExecutionResourcesPaginator
func NewListLifecycleExecutionResourcesPaginator(client ListLifecycleExecutionResourcesAPIClient, params *ListLifecycleExecutionResourcesInput, optFns ...func(*ListLifecycleExecutionResourcesPaginatorOptions)) *ListLifecycleExecutionResourcesPaginator {
	if params == nil {
		params = &ListLifecycleExecutionResourcesInput{}
	}

	options := ListLifecycleExecutionResourcesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListLifecycleExecutionResourcesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListLifecycleExecutionResourcesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListLifecycleExecutionResources page.
func (p *ListLifecycleExecutionResourcesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListLifecycleExecutionResourcesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	optFns = append([]func(*Options){
		addIsPaginatorUserAgent,
	}, optFns...)
	result, err := p.client.ListLifecycleExecutionResources(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

// ListLifecycleExecutionResourcesAPIClient is a client that implements the
// ListLifecycleExecutionResources operation.
type ListLifecycleExecutionResourcesAPIClient interface {
	ListLifecycleExecutionResources(context.Context, *ListLifecycleExecutionResourcesInput, ...func(*Options)) (*ListLifecycleExecutionResourcesOutput, error)
}

var _ ListLifecycleExecutionResourcesAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListLifecycleExecutionResources(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListLifecycleExecutionResources",
	}
}
