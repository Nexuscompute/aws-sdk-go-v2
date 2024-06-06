// Code generated by smithy-go-codegen DO NOT EDIT.

package xray

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/xray/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a list of traces specified by ID. Each trace is a collection of
// segment documents that originates from a single request. Use GetTraceSummaries
// to get a list of trace IDs.
func (c *Client) BatchGetTraces(ctx context.Context, params *BatchGetTracesInput, optFns ...func(*Options)) (*BatchGetTracesOutput, error) {
	if params == nil {
		params = &BatchGetTracesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchGetTraces", params, optFns, c.addOperationBatchGetTracesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchGetTracesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchGetTracesInput struct {

	// Specify the trace IDs of requests for which to retrieve segments.
	//
	// This member is required.
	TraceIds []string

	// Pagination token.
	NextToken *string

	noSmithyDocumentSerde
}

type BatchGetTracesOutput struct {

	// Pagination token.
	NextToken *string

	// Full traces for the specified requests.
	Traces []types.Trace

	// Trace IDs of requests that haven't been processed.
	UnprocessedTraceIds []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationBatchGetTracesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpBatchGetTraces{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpBatchGetTraces{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "BatchGetTraces"); err != nil {
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
	if err = addOpBatchGetTracesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchGetTraces(options.Region), middleware.Before); err != nil {
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

// BatchGetTracesAPIClient is a client that implements the BatchGetTraces
// operation.
type BatchGetTracesAPIClient interface {
	BatchGetTraces(context.Context, *BatchGetTracesInput, ...func(*Options)) (*BatchGetTracesOutput, error)
}

var _ BatchGetTracesAPIClient = (*Client)(nil)

// BatchGetTracesPaginatorOptions is the paginator options for BatchGetTraces
type BatchGetTracesPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// BatchGetTracesPaginator is a paginator for BatchGetTraces
type BatchGetTracesPaginator struct {
	options   BatchGetTracesPaginatorOptions
	client    BatchGetTracesAPIClient
	params    *BatchGetTracesInput
	nextToken *string
	firstPage bool
}

// NewBatchGetTracesPaginator returns a new BatchGetTracesPaginator
func NewBatchGetTracesPaginator(client BatchGetTracesAPIClient, params *BatchGetTracesInput, optFns ...func(*BatchGetTracesPaginatorOptions)) *BatchGetTracesPaginator {
	if params == nil {
		params = &BatchGetTracesInput{}
	}

	options := BatchGetTracesPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	return &BatchGetTracesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *BatchGetTracesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next BatchGetTraces page.
func (p *BatchGetTracesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*BatchGetTracesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	result, err := p.client.BatchGetTraces(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opBatchGetTraces(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "BatchGetTraces",
	}
}
