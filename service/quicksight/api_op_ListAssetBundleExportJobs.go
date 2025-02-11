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

// Lists all asset bundle export jobs that have been taken place in the last 14
// days. Jobs created more than 14 days ago are deleted forever and are not
// returned. If you are using the same job ID for multiple jobs,
// ListAssetBundleExportJobs only returns the most recent job that uses the
// repeated job ID.
func (c *Client) ListAssetBundleExportJobs(ctx context.Context, params *ListAssetBundleExportJobsInput, optFns ...func(*Options)) (*ListAssetBundleExportJobsOutput, error) {
	if params == nil {
		params = &ListAssetBundleExportJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAssetBundleExportJobs", params, optFns, c.addOperationListAssetBundleExportJobsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAssetBundleExportJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAssetBundleExportJobsInput struct {

	// The ID of the Amazon Web Services account that the export jobs were executed
	// in.
	//
	// This member is required.
	AwsAccountId *string

	// The maximum number of results to be returned per request.
	MaxResults *int32

	// The token for the next set of results, or null if there are no more results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListAssetBundleExportJobsOutput struct {

	// A list of export job summaries.
	AssetBundleExportJobSummaryList []types.AssetBundleExportJobSummary

	// The token for the next set of results, or null if there are no more results.
	NextToken *string

	// The Amazon Web Services request ID for this operation.
	RequestId *string

	// The HTTP status of the request.
	Status int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAssetBundleExportJobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListAssetBundleExportJobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListAssetBundleExportJobs{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListAssetBundleExportJobs"); err != nil {
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
	if err = addOpListAssetBundleExportJobsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAssetBundleExportJobs(options.Region), middleware.Before); err != nil {
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

// ListAssetBundleExportJobsPaginatorOptions is the paginator options for
// ListAssetBundleExportJobs
type ListAssetBundleExportJobsPaginatorOptions struct {
	// The maximum number of results to be returned per request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAssetBundleExportJobsPaginator is a paginator for ListAssetBundleExportJobs
type ListAssetBundleExportJobsPaginator struct {
	options   ListAssetBundleExportJobsPaginatorOptions
	client    ListAssetBundleExportJobsAPIClient
	params    *ListAssetBundleExportJobsInput
	nextToken *string
	firstPage bool
}

// NewListAssetBundleExportJobsPaginator returns a new
// ListAssetBundleExportJobsPaginator
func NewListAssetBundleExportJobsPaginator(client ListAssetBundleExportJobsAPIClient, params *ListAssetBundleExportJobsInput, optFns ...func(*ListAssetBundleExportJobsPaginatorOptions)) *ListAssetBundleExportJobsPaginator {
	if params == nil {
		params = &ListAssetBundleExportJobsInput{}
	}

	options := ListAssetBundleExportJobsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAssetBundleExportJobsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAssetBundleExportJobsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListAssetBundleExportJobs page.
func (p *ListAssetBundleExportJobsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAssetBundleExportJobsOutput, error) {
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
	result, err := p.client.ListAssetBundleExportJobs(ctx, &params, optFns...)
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

// ListAssetBundleExportJobsAPIClient is a client that implements the
// ListAssetBundleExportJobs operation.
type ListAssetBundleExportJobsAPIClient interface {
	ListAssetBundleExportJobs(context.Context, *ListAssetBundleExportJobsInput, ...func(*Options)) (*ListAssetBundleExportJobsOutput, error)
}

var _ ListAssetBundleExportJobsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListAssetBundleExportJobs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListAssetBundleExportJobs",
	}
}
