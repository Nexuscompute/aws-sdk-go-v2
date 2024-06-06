// Code generated by smithy-go-codegen DO NOT EDIT.

package batch

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/batch/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of Batch jobs.
//
// You must specify only one of the following items:
//
//   - A job queue ID to return a list of jobs in that job queue
//
//   - A multi-node parallel job ID to return a list of nodes for that job
//
//   - An array job ID to return a list of the children for that job
//
// You can filter the results by job status with the jobStatus parameter. If you
// don't specify a status, only RUNNING jobs are returned.
func (c *Client) ListJobs(ctx context.Context, params *ListJobsInput, optFns ...func(*Options)) (*ListJobsOutput, error) {
	if params == nil {
		params = &ListJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListJobs", params, optFns, c.addOperationListJobsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for ListJobs .
type ListJobsInput struct {

	// The job ID for an array job. Specifying an array job ID with this parameter
	// lists all child jobs from within the specified array.
	ArrayJobId *string

	// The filter to apply to the query. Only one filter can be used at a time. When
	// the filter is used, jobStatus is ignored. The filter doesn't apply to child
	// jobs in an array or multi-node parallel (MNP) jobs. The results are sorted by
	// the createdAt field, with the most recent jobs being first.
	//
	// JOB_NAME The value of the filter is a case-insensitive match for the job name.
	// If the value ends with an asterisk (*), the filter matches any job name that
	// begins with the string before the '*'. This corresponds to the jobName value.
	// For example, test1 matches both Test1 and test1 , and test1* matches both test1
	// and Test10 . When the JOB_NAME filter is used, the results are grouped by the
	// job name and version.
	//
	// JOB_DEFINITION The value for the filter is the name or Amazon Resource Name
	// (ARN) of the job definition. This corresponds to the jobDefinition value. The
	// value is case sensitive. When the value for the filter is the job definition
	// name, the results include all the jobs that used any revision of that job
	// definition name. If the value ends with an asterisk (*), the filter matches any
	// job definition name that begins with the string before the '*'. For example, jd1
	// matches only jd1 , and jd1* matches both jd1 and jd1A . The version of the job
	// definition that's used doesn't affect the sort order. When the JOB_DEFINITION
	// filter is used and the ARN is used (which is in the form
	// arn:${Partition}:batch:${Region}:${Account}:job-definition/${JobDefinitionName}:${Revision}
	// ), the results include jobs that used the specified revision of the job
	// definition. Asterisk (*) isn't supported when the ARN is used.
	//
	// BEFORE_CREATED_AT The value for the filter is the time that's before the job
	// was created. This corresponds to the createdAt value. The value is a string
	// representation of the number of milliseconds since 00:00:00 UTC (midnight) on
	// January 1, 1970.
	//
	// AFTER_CREATED_AT The value for the filter is the time that's after the job was
	// created. This corresponds to the createdAt value. The value is a string
	// representation of the number of milliseconds since 00:00:00 UTC (midnight) on
	// January 1, 1970.
	Filters []types.KeyValuesPair

	// The name or full Amazon Resource Name (ARN) of the job queue used to list jobs.
	JobQueue *string

	// The job status used to filter jobs in the specified queue. If the filters
	// parameter is specified, the jobStatus parameter is ignored and jobs with any
	// status are returned. If you don't specify a status, only RUNNING jobs are
	// returned.
	JobStatus types.JobStatus

	// The maximum number of results returned by ListJobs in a paginated output. When
	// this parameter is used, ListJobs returns up to maxResults results in a single
	// page and a nextToken response element, if applicable. The remaining results of
	// the initial request can be seen by sending another ListJobs request with the
	// returned nextToken value.
	//
	// The following outlines key parameters and limitations:
	//
	//   - The minimum value is 1.
	//
	//   - When --job-status is used, Batch returns up to 1000 values.
	//
	//   - When --filters is used, Batch returns up to 100 values.
	//
	//   - If neither parameter is used, then ListJobs returns up to 1000 results (jobs
	//   that are in the RUNNING status) and a nextToken value, if applicable.
	MaxResults *int32

	// The job ID for a multi-node parallel job. Specifying a multi-node parallel job
	// ID with this parameter lists all nodes that are associated with the specified
	// job.
	MultiNodeJobId *string

	// The nextToken value returned from a previous paginated ListJobs request where
	// maxResults was used and the results exceeded the value of that parameter.
	// Pagination continues from the end of the previous results that returned the
	// nextToken value. This value is null when there are no more results to return.
	//
	// Treat this token as an opaque identifier that's only used to retrieve the next
	// items in a list and not for other programmatic purposes.
	NextToken *string

	noSmithyDocumentSerde
}

type ListJobsOutput struct {

	// A list of job summaries that match the request.
	//
	// This member is required.
	JobSummaryList []types.JobSummary

	// The nextToken value to include in a future ListJobs request. When the results
	// of a ListJobs request exceed maxResults , this value can be used to retrieve the
	// next page of results. This value is null when there are no more results to
	// return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListJobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListJobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListJobs{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListJobs"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListJobs(options.Region), middleware.Before); err != nil {
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

// ListJobsAPIClient is a client that implements the ListJobs operation.
type ListJobsAPIClient interface {
	ListJobs(context.Context, *ListJobsInput, ...func(*Options)) (*ListJobsOutput, error)
}

var _ ListJobsAPIClient = (*Client)(nil)

// ListJobsPaginatorOptions is the paginator options for ListJobs
type ListJobsPaginatorOptions struct {
	// The maximum number of results returned by ListJobs in a paginated output. When
	// this parameter is used, ListJobs returns up to maxResults results in a single
	// page and a nextToken response element, if applicable. The remaining results of
	// the initial request can be seen by sending another ListJobs request with the
	// returned nextToken value.
	//
	// The following outlines key parameters and limitations:
	//
	//   - The minimum value is 1.
	//
	//   - When --job-status is used, Batch returns up to 1000 values.
	//
	//   - When --filters is used, Batch returns up to 100 values.
	//
	//   - If neither parameter is used, then ListJobs returns up to 1000 results (jobs
	//   that are in the RUNNING status) and a nextToken value, if applicable.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListJobsPaginator is a paginator for ListJobs
type ListJobsPaginator struct {
	options   ListJobsPaginatorOptions
	client    ListJobsAPIClient
	params    *ListJobsInput
	nextToken *string
	firstPage bool
}

// NewListJobsPaginator returns a new ListJobsPaginator
func NewListJobsPaginator(client ListJobsAPIClient, params *ListJobsInput, optFns ...func(*ListJobsPaginatorOptions)) *ListJobsPaginator {
	if params == nil {
		params = &ListJobsInput{}
	}

	options := ListJobsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListJobsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListJobsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListJobs page.
func (p *ListJobsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListJobsOutput, error) {
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

	result, err := p.client.ListJobs(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListJobs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListJobs",
	}
}
