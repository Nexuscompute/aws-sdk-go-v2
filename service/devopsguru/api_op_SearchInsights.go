// Code generated by smithy-go-codegen DO NOT EDIT.

package devopsguru

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/devopsguru/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

//	Returns a list of insights in your Amazon Web Services account. You can
//
// specify which insights are returned by their start time, one or more statuses (
// ONGOING or CLOSED ), one or more severities ( LOW , MEDIUM , and HIGH ), and
// type ( REACTIVE or PROACTIVE ).
//
// Use the Filters parameter to specify status and severity search parameters. Use
// the Type parameter to specify REACTIVE or PROACTIVE in your search.
func (c *Client) SearchInsights(ctx context.Context, params *SearchInsightsInput, optFns ...func(*Options)) (*SearchInsightsOutput, error) {
	if params == nil {
		params = &SearchInsightsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SearchInsights", params, optFns, c.addOperationSearchInsightsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SearchInsightsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SearchInsightsInput struct {

	//  The start of the time range passed in. Returned insights occurred after this
	// time.
	//
	// This member is required.
	StartTimeRange *types.StartTimeRange

	//  The type of insights you are searching for ( REACTIVE or PROACTIVE ).
	//
	// This member is required.
	Type types.InsightType

	//  A SearchInsightsFilters object that is used to set the severity and status
	// filters on your insight search.
	Filters *types.SearchInsightsFilters

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int32

	// The pagination token to use to retrieve the next page of results for this
	// operation. If this value is null, it retrieves the first page.
	NextToken *string

	noSmithyDocumentSerde
}

type SearchInsightsOutput struct {

	// The pagination token to use to retrieve the next page of results for this
	// operation. If there are no more pages, this value is null.
	NextToken *string

	//  The returned proactive insights.
	ProactiveInsights []types.ProactiveInsightSummary

	//  The returned reactive insights.
	ReactiveInsights []types.ReactiveInsightSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSearchInsightsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpSearchInsights{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpSearchInsights{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "SearchInsights"); err != nil {
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
	if err = addOpSearchInsightsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSearchInsights(options.Region), middleware.Before); err != nil {
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

// SearchInsightsAPIClient is a client that implements the SearchInsights
// operation.
type SearchInsightsAPIClient interface {
	SearchInsights(context.Context, *SearchInsightsInput, ...func(*Options)) (*SearchInsightsOutput, error)
}

var _ SearchInsightsAPIClient = (*Client)(nil)

// SearchInsightsPaginatorOptions is the paginator options for SearchInsights
type SearchInsightsPaginatorOptions struct {
	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// SearchInsightsPaginator is a paginator for SearchInsights
type SearchInsightsPaginator struct {
	options   SearchInsightsPaginatorOptions
	client    SearchInsightsAPIClient
	params    *SearchInsightsInput
	nextToken *string
	firstPage bool
}

// NewSearchInsightsPaginator returns a new SearchInsightsPaginator
func NewSearchInsightsPaginator(client SearchInsightsAPIClient, params *SearchInsightsInput, optFns ...func(*SearchInsightsPaginatorOptions)) *SearchInsightsPaginator {
	if params == nil {
		params = &SearchInsightsInput{}
	}

	options := SearchInsightsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &SearchInsightsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *SearchInsightsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next SearchInsights page.
func (p *SearchInsightsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*SearchInsightsOutput, error) {
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

	result, err := p.client.SearchInsights(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opSearchInsights(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "SearchInsights",
	}
}
