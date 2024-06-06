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

// Gets a set of statistics for queues or farms. Before you can call the
// GetSessionStatisticsAggregation operation, you must first call the
// StartSessionsStatisticsAggregation operation. Statistics are available for 1
// hour after you call the StartSessionsStatisticsAggregation operation.
func (c *Client) GetSessionsStatisticsAggregation(ctx context.Context, params *GetSessionsStatisticsAggregationInput, optFns ...func(*Options)) (*GetSessionsStatisticsAggregationOutput, error) {
	if params == nil {
		params = &GetSessionsStatisticsAggregationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetSessionsStatisticsAggregation", params, optFns, c.addOperationGetSessionsStatisticsAggregationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetSessionsStatisticsAggregationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetSessionsStatisticsAggregationInput struct {

	// The identifier returned by the StartSessionsStatisticsAggregation operation
	// that identifies the aggregated statistics.
	//
	// This member is required.
	AggregationId *string

	// The identifier of the farm to include in the statistics. This should be the
	// same as the farm ID used in the call to the StartSessionsStatisticsAggregation
	// operation.
	//
	// This member is required.
	FarmId *string

	// The maximum number of results to return. Use this parameter with NextToken to
	// get results as a set of sequential pages.
	MaxResults *int32

	// The token for the next set of results, or null to start from the beginning.
	NextToken *string

	noSmithyDocumentSerde
}

type GetSessionsStatisticsAggregationOutput struct {

	// The status of the aggregated results.
	//
	// This member is required.
	Status types.SessionsStatisticsAggregationStatus

	// If Deadline Cloud returns nextToken , then there are more results available. The
	// value of nextToken is a unique pagination token for each page. To retrieve the
	// next page, call the operation again using the returned token. Keep all other
	// arguments unchanged. If no results remain, then nextToken is set to null . Each
	// pagination token expires after 24 hours. If you provide a token that isn't
	// valid, then you receive an HTTP 400 ValidationException error.
	NextToken *string

	// The statistics for the specified fleets or queues.
	Statistics []types.Statistics

	// A message that describes the status.
	StatusMessage *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetSessionsStatisticsAggregationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetSessionsStatisticsAggregation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetSessionsStatisticsAggregation{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetSessionsStatisticsAggregation"); err != nil {
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
	if err = addEndpointPrefix_opGetSessionsStatisticsAggregationMiddleware(stack); err != nil {
		return err
	}
	if err = addOpGetSessionsStatisticsAggregationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetSessionsStatisticsAggregation(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opGetSessionsStatisticsAggregationMiddleware struct {
}

func (*endpointPrefix_opGetSessionsStatisticsAggregationMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opGetSessionsStatisticsAggregationMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
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
func addEndpointPrefix_opGetSessionsStatisticsAggregationMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opGetSessionsStatisticsAggregationMiddleware{}, "ResolveEndpointV2", middleware.After)
}

// GetSessionsStatisticsAggregationAPIClient is a client that implements the
// GetSessionsStatisticsAggregation operation.
type GetSessionsStatisticsAggregationAPIClient interface {
	GetSessionsStatisticsAggregation(context.Context, *GetSessionsStatisticsAggregationInput, ...func(*Options)) (*GetSessionsStatisticsAggregationOutput, error)
}

var _ GetSessionsStatisticsAggregationAPIClient = (*Client)(nil)

// GetSessionsStatisticsAggregationPaginatorOptions is the paginator options for
// GetSessionsStatisticsAggregation
type GetSessionsStatisticsAggregationPaginatorOptions struct {
	// The maximum number of results to return. Use this parameter with NextToken to
	// get results as a set of sequential pages.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetSessionsStatisticsAggregationPaginator is a paginator for
// GetSessionsStatisticsAggregation
type GetSessionsStatisticsAggregationPaginator struct {
	options   GetSessionsStatisticsAggregationPaginatorOptions
	client    GetSessionsStatisticsAggregationAPIClient
	params    *GetSessionsStatisticsAggregationInput
	nextToken *string
	firstPage bool
}

// NewGetSessionsStatisticsAggregationPaginator returns a new
// GetSessionsStatisticsAggregationPaginator
func NewGetSessionsStatisticsAggregationPaginator(client GetSessionsStatisticsAggregationAPIClient, params *GetSessionsStatisticsAggregationInput, optFns ...func(*GetSessionsStatisticsAggregationPaginatorOptions)) *GetSessionsStatisticsAggregationPaginator {
	if params == nil {
		params = &GetSessionsStatisticsAggregationInput{}
	}

	options := GetSessionsStatisticsAggregationPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetSessionsStatisticsAggregationPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetSessionsStatisticsAggregationPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetSessionsStatisticsAggregation page.
func (p *GetSessionsStatisticsAggregationPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetSessionsStatisticsAggregationOutput, error) {
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

	result, err := p.client.GetSessionsStatisticsAggregation(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetSessionsStatisticsAggregation(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetSessionsStatisticsAggregation",
	}
}
