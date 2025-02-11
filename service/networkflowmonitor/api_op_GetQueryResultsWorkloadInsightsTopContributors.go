// Code generated by smithy-go-codegen DO NOT EDIT.

package networkflowmonitor

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/networkflowmonitor/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Return the data for a query with the Network Flow Monitor query interface. You
// specify the query that you want to return results for by providing a query ID
// and a monitor name. This query returns the top contributors for a specific
// monitor.
//
// Create a query ID for this call by calling the corresponding API call to start
// the query, StartQueryWorkloadInsightsTopContributors . Use the scope ID that was
// returned for your account by CreateScope .
//
// Top contributors in Network Flow Monitor are network flows with the highest
// values for a specific metric type, related to a scope (for workload insights) or
// a monitor.
func (c *Client) GetQueryResultsWorkloadInsightsTopContributors(ctx context.Context, params *GetQueryResultsWorkloadInsightsTopContributorsInput, optFns ...func(*Options)) (*GetQueryResultsWorkloadInsightsTopContributorsOutput, error) {
	if params == nil {
		params = &GetQueryResultsWorkloadInsightsTopContributorsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetQueryResultsWorkloadInsightsTopContributors", params, optFns, c.addOperationGetQueryResultsWorkloadInsightsTopContributorsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetQueryResultsWorkloadInsightsTopContributorsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetQueryResultsWorkloadInsightsTopContributorsInput struct {

	// The identifier for the query. A query ID is an internally-generated identifier
	// for a specific query returned from an API call to start a query.
	//
	// This member is required.
	QueryId *string

	// The identifier for the scope that includes the resources you want to get data
	// results for. A scope ID is an internally-generated identifier that includes all
	// the resources for a specific root account.
	//
	// This member is required.
	ScopeId *string

	// The number of query results that you want to return with this call.
	MaxResults *int32

	// The token for the next set of results. You receive this token from a previous
	// call.
	NextToken *string

	noSmithyDocumentSerde
}

type GetQueryResultsWorkloadInsightsTopContributorsOutput struct {

	// The token for the next set of results. You receive this token from a previous
	// call.
	NextToken *string

	// The top contributor network flows overall for a specific metric type, for
	// example, the number of retransmissions.
	TopContributors []types.WorkloadInsightsTopContributorsRow

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetQueryResultsWorkloadInsightsTopContributorsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetQueryResultsWorkloadInsightsTopContributors{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetQueryResultsWorkloadInsightsTopContributors{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetQueryResultsWorkloadInsightsTopContributors"); err != nil {
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
	if err = addOpGetQueryResultsWorkloadInsightsTopContributorsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetQueryResultsWorkloadInsightsTopContributors(options.Region), middleware.Before); err != nil {
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

// GetQueryResultsWorkloadInsightsTopContributorsPaginatorOptions is the paginator
// options for GetQueryResultsWorkloadInsightsTopContributors
type GetQueryResultsWorkloadInsightsTopContributorsPaginatorOptions struct {
	// The number of query results that you want to return with this call.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetQueryResultsWorkloadInsightsTopContributorsPaginator is a paginator for
// GetQueryResultsWorkloadInsightsTopContributors
type GetQueryResultsWorkloadInsightsTopContributorsPaginator struct {
	options   GetQueryResultsWorkloadInsightsTopContributorsPaginatorOptions
	client    GetQueryResultsWorkloadInsightsTopContributorsAPIClient
	params    *GetQueryResultsWorkloadInsightsTopContributorsInput
	nextToken *string
	firstPage bool
}

// NewGetQueryResultsWorkloadInsightsTopContributorsPaginator returns a new
// GetQueryResultsWorkloadInsightsTopContributorsPaginator
func NewGetQueryResultsWorkloadInsightsTopContributorsPaginator(client GetQueryResultsWorkloadInsightsTopContributorsAPIClient, params *GetQueryResultsWorkloadInsightsTopContributorsInput, optFns ...func(*GetQueryResultsWorkloadInsightsTopContributorsPaginatorOptions)) *GetQueryResultsWorkloadInsightsTopContributorsPaginator {
	if params == nil {
		params = &GetQueryResultsWorkloadInsightsTopContributorsInput{}
	}

	options := GetQueryResultsWorkloadInsightsTopContributorsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetQueryResultsWorkloadInsightsTopContributorsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetQueryResultsWorkloadInsightsTopContributorsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetQueryResultsWorkloadInsightsTopContributors page.
func (p *GetQueryResultsWorkloadInsightsTopContributorsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetQueryResultsWorkloadInsightsTopContributorsOutput, error) {
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
	result, err := p.client.GetQueryResultsWorkloadInsightsTopContributors(ctx, &params, optFns...)
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

// GetQueryResultsWorkloadInsightsTopContributorsAPIClient is a client that
// implements the GetQueryResultsWorkloadInsightsTopContributors operation.
type GetQueryResultsWorkloadInsightsTopContributorsAPIClient interface {
	GetQueryResultsWorkloadInsightsTopContributors(context.Context, *GetQueryResultsWorkloadInsightsTopContributorsInput, ...func(*Options)) (*GetQueryResultsWorkloadInsightsTopContributorsOutput, error)
}

var _ GetQueryResultsWorkloadInsightsTopContributorsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opGetQueryResultsWorkloadInsightsTopContributors(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetQueryResultsWorkloadInsightsTopContributors",
	}
}
