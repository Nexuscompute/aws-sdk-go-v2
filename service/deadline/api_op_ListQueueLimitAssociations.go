// Code generated by smithy-go-codegen DO NOT EDIT.

package deadline

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/deadline/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets a list of the associations between queues and limits defined in a farm.
func (c *Client) ListQueueLimitAssociations(ctx context.Context, params *ListQueueLimitAssociationsInput, optFns ...func(*Options)) (*ListQueueLimitAssociationsOutput, error) {
	if params == nil {
		params = &ListQueueLimitAssociationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListQueueLimitAssociations", params, optFns, c.addOperationListQueueLimitAssociationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListQueueLimitAssociationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListQueueLimitAssociationsInput struct {

	// The unique identifier of the farm that contains the limits and associations.
	//
	// This member is required.
	FarmId *string

	// Specifies that the operation should return only the queue limit associations
	// for the specified limit. If you specify both the queueId and the limitId , only
	// the specified limit is returned if it exists.
	LimitId *string

	// The maximum number of associations to return in each page of results.
	MaxResults *int32

	// The token for the next set of results, or null to start from the beginning.
	NextToken *string

	// Specifies that the operation should return only the queue limit associations
	// for the specified queue. If you specify both the queueId and the limitId , only
	// the specified limit is returned if it exists.
	QueueId *string

	noSmithyDocumentSerde
}

type ListQueueLimitAssociationsOutput struct {

	// A list of associations between limits and queues in the farm specified in the
	// request.
	//
	// This member is required.
	QueueLimitAssociations []types.QueueLimitAssociationSummary

	// If Deadline Cloud returns nextToken , then there are more results available. The
	// value of nextToken is a unique pagination token for each page. To retrieve the
	// next page, call the operation again using the returned token. Keep all other
	// arguments unchanged. If no results remain, then nextToken is set to null . Each
	// pagination token expires after 24 hours. If you provide a token that isn't
	// valid, then you receive an HTTP 400 ValidationException error.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListQueueLimitAssociationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListQueueLimitAssociations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListQueueLimitAssociations{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListQueueLimitAssociations"); err != nil {
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
	if err = addEndpointPrefix_opListQueueLimitAssociationsMiddleware(stack); err != nil {
		return err
	}
	if err = addOpListQueueLimitAssociationsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListQueueLimitAssociations(options.Region), middleware.Before); err != nil {
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

// ListQueueLimitAssociationsPaginatorOptions is the paginator options for
// ListQueueLimitAssociations
type ListQueueLimitAssociationsPaginatorOptions struct {
	// The maximum number of associations to return in each page of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListQueueLimitAssociationsPaginator is a paginator for
// ListQueueLimitAssociations
type ListQueueLimitAssociationsPaginator struct {
	options   ListQueueLimitAssociationsPaginatorOptions
	client    ListQueueLimitAssociationsAPIClient
	params    *ListQueueLimitAssociationsInput
	nextToken *string
	firstPage bool
}

// NewListQueueLimitAssociationsPaginator returns a new
// ListQueueLimitAssociationsPaginator
func NewListQueueLimitAssociationsPaginator(client ListQueueLimitAssociationsAPIClient, params *ListQueueLimitAssociationsInput, optFns ...func(*ListQueueLimitAssociationsPaginatorOptions)) *ListQueueLimitAssociationsPaginator {
	if params == nil {
		params = &ListQueueLimitAssociationsInput{}
	}

	options := ListQueueLimitAssociationsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListQueueLimitAssociationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListQueueLimitAssociationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListQueueLimitAssociations page.
func (p *ListQueueLimitAssociationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListQueueLimitAssociationsOutput, error) {
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
	result, err := p.client.ListQueueLimitAssociations(ctx, &params, optFns...)
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

type endpointPrefix_opListQueueLimitAssociationsMiddleware struct {
}

func (*endpointPrefix_opListQueueLimitAssociationsMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opListQueueLimitAssociationsMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
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
func addEndpointPrefix_opListQueueLimitAssociationsMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opListQueueLimitAssociationsMiddleware{}, "ResolveEndpointV2", middleware.After)
}

// ListQueueLimitAssociationsAPIClient is a client that implements the
// ListQueueLimitAssociations operation.
type ListQueueLimitAssociationsAPIClient interface {
	ListQueueLimitAssociations(context.Context, *ListQueueLimitAssociationsInput, ...func(*Options)) (*ListQueueLimitAssociationsOutput, error)
}

var _ ListQueueLimitAssociationsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListQueueLimitAssociations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListQueueLimitAssociations",
	}
}
