// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all asset bundle import jobs that have taken place in the last 14 days.
// Jobs created more than 14 days ago are deleted forever and are not returned. If
// you are using the same job ID for multiple jobs, ListAssetBundleImportJobs only
// returns the most recent job that uses the repeated job ID.
func (c *Client) ListAssetBundleImportJobs(ctx context.Context, params *ListAssetBundleImportJobsInput, optFns ...func(*Options)) (*ListAssetBundleImportJobsOutput, error) {
	if params == nil {
		params = &ListAssetBundleImportJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAssetBundleImportJobs", params, optFns, c.addOperationListAssetBundleImportJobsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAssetBundleImportJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAssetBundleImportJobsInput struct {

	// The ID of the Amazon Web Services account that the import jobs were executed in.
	//
	// This member is required.
	AwsAccountId *string

	// The maximum number of results to be returned per request.
	MaxResults *int32

	// The token for the next set of results, or null if there are no more results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListAssetBundleImportJobsOutput struct {

	// A list of import job summaries.
	AssetBundleImportJobSummaryList []types.AssetBundleImportJobSummary

	// The token for the next set of results, or null if there are no more results.
	NextToken *string

	// The Amazon Web Services request ID for this operation.
	RequestId *string

	// The HTTP status of the response.
	Status int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAssetBundleImportJobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListAssetBundleImportJobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListAssetBundleImportJobs{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListAssetBundleImportJobs"); err != nil {
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
	if err = addOpListAssetBundleImportJobsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAssetBundleImportJobs(options.Region), middleware.Before); err != nil {
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

// ListAssetBundleImportJobsPaginatorOptions is the paginator options for
// ListAssetBundleImportJobs
type ListAssetBundleImportJobsPaginatorOptions struct {
	// The maximum number of results to be returned per request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAssetBundleImportJobsPaginator is a paginator for ListAssetBundleImportJobs
type ListAssetBundleImportJobsPaginator struct {
	options   ListAssetBundleImportJobsPaginatorOptions
	client    ListAssetBundleImportJobsAPIClient
	params    *ListAssetBundleImportJobsInput
	nextToken *string
	firstPage bool
}

// NewListAssetBundleImportJobsPaginator returns a new
// ListAssetBundleImportJobsPaginator
func NewListAssetBundleImportJobsPaginator(client ListAssetBundleImportJobsAPIClient, params *ListAssetBundleImportJobsInput, optFns ...func(*ListAssetBundleImportJobsPaginatorOptions)) *ListAssetBundleImportJobsPaginator {
	if params == nil {
		params = &ListAssetBundleImportJobsInput{}
	}

	options := ListAssetBundleImportJobsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAssetBundleImportJobsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAssetBundleImportJobsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListAssetBundleImportJobs page.
func (p *ListAssetBundleImportJobsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAssetBundleImportJobsOutput, error) {
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
	result, err := p.client.ListAssetBundleImportJobs(ctx, &params, optFns...)
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

// ListAssetBundleImportJobsAPIClient is a client that implements the
// ListAssetBundleImportJobs operation.
type ListAssetBundleImportJobsAPIClient interface {
	ListAssetBundleImportJobs(context.Context, *ListAssetBundleImportJobsInput, ...func(*Options)) (*ListAssetBundleImportJobsOutput, error)
}

var _ ListAssetBundleImportJobsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListAssetBundleImportJobs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListAssetBundleImportJobs",
	}
}
