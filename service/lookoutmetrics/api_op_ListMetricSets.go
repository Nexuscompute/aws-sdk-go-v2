// Code generated by smithy-go-codegen DO NOT EDIT.

package lookoutmetrics

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/lookoutmetrics/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the datasets in the current AWS Region.
//
// Amazon Lookout for Metrics API actions are eventually consistent. If you do a
// read operation on a resource immediately after creating or modifying it, use
// retries to allow time for the write operation to complete.
func (c *Client) ListMetricSets(ctx context.Context, params *ListMetricSetsInput, optFns ...func(*Options)) (*ListMetricSetsOutput, error) {
	if params == nil {
		params = &ListMetricSetsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListMetricSets", params, optFns, c.addOperationListMetricSetsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListMetricSetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListMetricSetsInput struct {

	// The ARN of the anomaly detector containing the metrics sets to list.
	AnomalyDetectorArn *string

	// The maximum number of results to return.
	MaxResults *int32

	// If the result of the previous request was truncated, the response includes a
	// NextToken . To retrieve the next set of results, use the token in the next
	// request. Tokens expire after 24 hours.
	NextToken *string

	noSmithyDocumentSerde
}

type ListMetricSetsOutput struct {

	// A list of the datasets in the AWS Region, with configuration details for each.
	MetricSetSummaryList []types.MetricSetSummary

	// If the response is truncated, the list call returns this token. To retrieve the
	// next set of results, use the token in the next list request.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListMetricSetsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListMetricSets{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListMetricSets{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListMetricSets"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListMetricSets(options.Region), middleware.Before); err != nil {
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

// ListMetricSetsPaginatorOptions is the paginator options for ListMetricSets
type ListMetricSetsPaginatorOptions struct {
	// The maximum number of results to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListMetricSetsPaginator is a paginator for ListMetricSets
type ListMetricSetsPaginator struct {
	options   ListMetricSetsPaginatorOptions
	client    ListMetricSetsAPIClient
	params    *ListMetricSetsInput
	nextToken *string
	firstPage bool
}

// NewListMetricSetsPaginator returns a new ListMetricSetsPaginator
func NewListMetricSetsPaginator(client ListMetricSetsAPIClient, params *ListMetricSetsInput, optFns ...func(*ListMetricSetsPaginatorOptions)) *ListMetricSetsPaginator {
	if params == nil {
		params = &ListMetricSetsInput{}
	}

	options := ListMetricSetsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListMetricSetsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListMetricSetsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListMetricSets page.
func (p *ListMetricSetsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListMetricSetsOutput, error) {
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
	result, err := p.client.ListMetricSets(ctx, &params, optFns...)
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

// ListMetricSetsAPIClient is a client that implements the ListMetricSets
// operation.
type ListMetricSetsAPIClient interface {
	ListMetricSets(context.Context, *ListMetricSetsInput, ...func(*Options)) (*ListMetricSetsOutput, error)
}

var _ ListMetricSetsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListMetricSets(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListMetricSets",
	}
}
