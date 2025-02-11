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

// List the cluster policy configurations.
func (c *Client) ListClusterSchedulerConfigs(ctx context.Context, params *ListClusterSchedulerConfigsInput, optFns ...func(*Options)) (*ListClusterSchedulerConfigsOutput, error) {
	if params == nil {
		params = &ListClusterSchedulerConfigsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListClusterSchedulerConfigs", params, optFns, c.addOperationListClusterSchedulerConfigsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListClusterSchedulerConfigsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListClusterSchedulerConfigsInput struct {

	// Filter for ARN of the cluster.
	ClusterArn *string

	// Filter for after this creation time. The input for this parameter is a Unix
	// timestamp. To convert a date and time into a Unix timestamp, see [EpochConverter].
	//
	// [EpochConverter]: https://www.epochconverter.com/
	CreatedAfter *time.Time

	// Filter for before this creation time. The input for this parameter is a Unix
	// timestamp. To convert a date and time into a Unix timestamp, see [EpochConverter].
	//
	// [EpochConverter]: https://www.epochconverter.com/
	CreatedBefore *time.Time

	// The maximum number of cluster policies to list.
	MaxResults *int32

	// Filter for name containing this string.
	NameContains *string

	// If the previous response was truncated, you will receive this token. Use it in
	// your next request to receive the next set of results.
	NextToken *string

	// Filter for sorting the list by a given value. For example, sort by name,
	// creation time, or status.
	SortBy types.SortClusterSchedulerConfigBy

	// The order of the list. By default, listed in Descending order according to by
	// SortBy . To change the list order, you can specify SortOrder to be Ascending .
	SortOrder types.SortOrder

	// Filter for status.
	Status types.SchedulerResourceStatus

	noSmithyDocumentSerde
}

type ListClusterSchedulerConfigsOutput struct {

	// Summaries of the cluster policies.
	ClusterSchedulerConfigSummaries []types.ClusterSchedulerConfigSummary

	// If the previous response was truncated, you will receive this token. Use it in
	// your next request to receive the next set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListClusterSchedulerConfigsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListClusterSchedulerConfigs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListClusterSchedulerConfigs{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListClusterSchedulerConfigs"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListClusterSchedulerConfigs(options.Region), middleware.Before); err != nil {
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

// ListClusterSchedulerConfigsPaginatorOptions is the paginator options for
// ListClusterSchedulerConfigs
type ListClusterSchedulerConfigsPaginatorOptions struct {
	// The maximum number of cluster policies to list.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListClusterSchedulerConfigsPaginator is a paginator for
// ListClusterSchedulerConfigs
type ListClusterSchedulerConfigsPaginator struct {
	options   ListClusterSchedulerConfigsPaginatorOptions
	client    ListClusterSchedulerConfigsAPIClient
	params    *ListClusterSchedulerConfigsInput
	nextToken *string
	firstPage bool
}

// NewListClusterSchedulerConfigsPaginator returns a new
// ListClusterSchedulerConfigsPaginator
func NewListClusterSchedulerConfigsPaginator(client ListClusterSchedulerConfigsAPIClient, params *ListClusterSchedulerConfigsInput, optFns ...func(*ListClusterSchedulerConfigsPaginatorOptions)) *ListClusterSchedulerConfigsPaginator {
	if params == nil {
		params = &ListClusterSchedulerConfigsInput{}
	}

	options := ListClusterSchedulerConfigsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListClusterSchedulerConfigsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListClusterSchedulerConfigsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListClusterSchedulerConfigs page.
func (p *ListClusterSchedulerConfigsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListClusterSchedulerConfigsOutput, error) {
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
	result, err := p.client.ListClusterSchedulerConfigs(ctx, &params, optFns...)
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

// ListClusterSchedulerConfigsAPIClient is a client that implements the
// ListClusterSchedulerConfigs operation.
type ListClusterSchedulerConfigsAPIClient interface {
	ListClusterSchedulerConfigs(context.Context, *ListClusterSchedulerConfigsInput, ...func(*Options)) (*ListClusterSchedulerConfigsOutput, error)
}

var _ ListClusterSchedulerConfigsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListClusterSchedulerConfigs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListClusterSchedulerConfigs",
	}
}
