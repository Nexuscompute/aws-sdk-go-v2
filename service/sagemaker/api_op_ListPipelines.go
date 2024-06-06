// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Gets a list of pipelines.
func (c *Client) ListPipelines(ctx context.Context, params *ListPipelinesInput, optFns ...func(*Options)) (*ListPipelinesOutput, error) {
	if params == nil {
		params = &ListPipelinesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListPipelines", params, optFns, c.addOperationListPipelinesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListPipelinesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPipelinesInput struct {

	// A filter that returns the pipelines that were created after a specified time.
	CreatedAfter *time.Time

	// A filter that returns the pipelines that were created before a specified time.
	CreatedBefore *time.Time

	// The maximum number of pipelines to return in the response.
	MaxResults *int32

	// If the result of the previous ListPipelines request was truncated, the response
	// includes a NextToken . To retrieve the next set of pipelines, use the token in
	// the next request.
	NextToken *string

	// The prefix of the pipeline name.
	PipelineNamePrefix *string

	// The field by which to sort results. The default is CreatedTime .
	SortBy types.SortPipelinesBy

	// The sort order for results.
	SortOrder types.SortOrder

	noSmithyDocumentSerde
}

type ListPipelinesOutput struct {

	// If the result of the previous ListPipelines request was truncated, the response
	// includes a NextToken . To retrieve the next set of pipelines, use the token in
	// the next request.
	NextToken *string

	// Contains a sorted list of PipelineSummary objects matching the specified
	// filters. Each PipelineSummary consists of PipelineArn, PipelineName,
	// ExperimentName, PipelineDescription, CreationTime, LastModifiedTime,
	// LastRunTime, and RoleArn. This list can be empty.
	PipelineSummaries []types.PipelineSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListPipelinesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListPipelines{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListPipelines{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListPipelines"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListPipelines(options.Region), middleware.Before); err != nil {
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

// ListPipelinesAPIClient is a client that implements the ListPipelines operation.
type ListPipelinesAPIClient interface {
	ListPipelines(context.Context, *ListPipelinesInput, ...func(*Options)) (*ListPipelinesOutput, error)
}

var _ ListPipelinesAPIClient = (*Client)(nil)

// ListPipelinesPaginatorOptions is the paginator options for ListPipelines
type ListPipelinesPaginatorOptions struct {
	// The maximum number of pipelines to return in the response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListPipelinesPaginator is a paginator for ListPipelines
type ListPipelinesPaginator struct {
	options   ListPipelinesPaginatorOptions
	client    ListPipelinesAPIClient
	params    *ListPipelinesInput
	nextToken *string
	firstPage bool
}

// NewListPipelinesPaginator returns a new ListPipelinesPaginator
func NewListPipelinesPaginator(client ListPipelinesAPIClient, params *ListPipelinesInput, optFns ...func(*ListPipelinesPaginatorOptions)) *ListPipelinesPaginator {
	if params == nil {
		params = &ListPipelinesInput{}
	}

	options := ListPipelinesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListPipelinesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListPipelinesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListPipelines page.
func (p *ListPipelinesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListPipelinesOutput, error) {
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

	result, err := p.client.ListPipelines(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListPipelines(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListPipelines",
	}
}
