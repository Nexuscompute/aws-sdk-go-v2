// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatch

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of metric streams in this account.
func (c *Client) ListMetricStreams(ctx context.Context, params *ListMetricStreamsInput, optFns ...func(*Options)) (*ListMetricStreamsOutput, error) {
	if params == nil {
		params = &ListMetricStreamsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListMetricStreams", params, optFns, c.addOperationListMetricStreamsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListMetricStreamsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListMetricStreamsInput struct {

	// The maximum number of results to return in one operation.
	MaxResults *int32

	// Include this value, if it was returned by the previous call, to get the next
	// set of metric streams.
	NextToken *string

	noSmithyDocumentSerde
}

type ListMetricStreamsOutput struct {

	// The array of metric stream information.
	Entries []types.MetricStreamEntry

	// The token that marks the start of the next batch of returned results. You can
	// use this token in a subsequent operation to get the next batch of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListMetricStreamsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpListMetricStreams{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpListMetricStreams{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListMetricStreams"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListMetricStreams(options.Region), middleware.Before); err != nil {
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

// ListMetricStreamsAPIClient is a client that implements the ListMetricStreams
// operation.
type ListMetricStreamsAPIClient interface {
	ListMetricStreams(context.Context, *ListMetricStreamsInput, ...func(*Options)) (*ListMetricStreamsOutput, error)
}

var _ ListMetricStreamsAPIClient = (*Client)(nil)

// ListMetricStreamsPaginatorOptions is the paginator options for ListMetricStreams
type ListMetricStreamsPaginatorOptions struct {
	// The maximum number of results to return in one operation.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListMetricStreamsPaginator is a paginator for ListMetricStreams
type ListMetricStreamsPaginator struct {
	options   ListMetricStreamsPaginatorOptions
	client    ListMetricStreamsAPIClient
	params    *ListMetricStreamsInput
	nextToken *string
	firstPage bool
}

// NewListMetricStreamsPaginator returns a new ListMetricStreamsPaginator
func NewListMetricStreamsPaginator(client ListMetricStreamsAPIClient, params *ListMetricStreamsInput, optFns ...func(*ListMetricStreamsPaginatorOptions)) *ListMetricStreamsPaginator {
	if params == nil {
		params = &ListMetricStreamsInput{}
	}

	options := ListMetricStreamsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListMetricStreamsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListMetricStreamsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListMetricStreams page.
func (p *ListMetricStreamsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListMetricStreamsOutput, error) {
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

	result, err := p.client.ListMetricStreams(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListMetricStreams(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListMetricStreams",
	}
}
