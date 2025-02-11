// Code generated by smithy-go-codegen DO NOT EDIT.

package qbusiness

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/qbusiness/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Get information about an Amazon Q Business data source connector
// synchronization.
func (c *Client) ListDataSourceSyncJobs(ctx context.Context, params *ListDataSourceSyncJobsInput, optFns ...func(*Options)) (*ListDataSourceSyncJobsOutput, error) {
	if params == nil {
		params = &ListDataSourceSyncJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDataSourceSyncJobs", params, optFns, c.addOperationListDataSourceSyncJobsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDataSourceSyncJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDataSourceSyncJobsInput struct {

	// The identifier of the Amazon Q Business application connected to the data
	// source.
	//
	// This member is required.
	ApplicationId *string

	//  The identifier of the data source connector.
	//
	// This member is required.
	DataSourceId *string

	// The identifier of the index used with the Amazon Q Business data source
	// connector.
	//
	// This member is required.
	IndexId *string

	//  The end time of the data source connector sync.
	EndTime *time.Time

	// The maximum number of synchronization jobs to return in the response.
	MaxResults *int32

	// If the maxResults response was incpmplete because there is more data to
	// retriever, Amazon Q Business returns a pagination token in the response. You can
	// use this pagination token to retrieve the next set of responses.
	NextToken *string

	//  The start time of the data source connector sync.
	StartTime *time.Time

	// Only returns synchronization jobs with the Status field equal to the specified
	// status.
	StatusFilter types.DataSourceSyncJobStatus

	noSmithyDocumentSerde
}

type ListDataSourceSyncJobsOutput struct {

	// A history of synchronization jobs for the data source connector.
	History []types.DataSourceSyncJob

	// If the response is truncated, Amazon Q Business returns this token. You can use
	// this token in any subsequent request to retrieve the next set of jobs.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListDataSourceSyncJobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListDataSourceSyncJobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListDataSourceSyncJobs{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListDataSourceSyncJobs"); err != nil {
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
	if err = addOpListDataSourceSyncJobsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListDataSourceSyncJobs(options.Region), middleware.Before); err != nil {
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

// ListDataSourceSyncJobsPaginatorOptions is the paginator options for
// ListDataSourceSyncJobs
type ListDataSourceSyncJobsPaginatorOptions struct {
	// The maximum number of synchronization jobs to return in the response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListDataSourceSyncJobsPaginator is a paginator for ListDataSourceSyncJobs
type ListDataSourceSyncJobsPaginator struct {
	options   ListDataSourceSyncJobsPaginatorOptions
	client    ListDataSourceSyncJobsAPIClient
	params    *ListDataSourceSyncJobsInput
	nextToken *string
	firstPage bool
}

// NewListDataSourceSyncJobsPaginator returns a new ListDataSourceSyncJobsPaginator
func NewListDataSourceSyncJobsPaginator(client ListDataSourceSyncJobsAPIClient, params *ListDataSourceSyncJobsInput, optFns ...func(*ListDataSourceSyncJobsPaginatorOptions)) *ListDataSourceSyncJobsPaginator {
	if params == nil {
		params = &ListDataSourceSyncJobsInput{}
	}

	options := ListDataSourceSyncJobsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListDataSourceSyncJobsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListDataSourceSyncJobsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListDataSourceSyncJobs page.
func (p *ListDataSourceSyncJobsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListDataSourceSyncJobsOutput, error) {
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
	result, err := p.client.ListDataSourceSyncJobs(ctx, &params, optFns...)
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

// ListDataSourceSyncJobsAPIClient is a client that implements the
// ListDataSourceSyncJobs operation.
type ListDataSourceSyncJobsAPIClient interface {
	ListDataSourceSyncJobs(context.Context, *ListDataSourceSyncJobsInput, ...func(*Options)) (*ListDataSourceSyncJobsOutput, error)
}

var _ ListDataSourceSyncJobsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListDataSourceSyncJobs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListDataSourceSyncJobs",
	}
}
