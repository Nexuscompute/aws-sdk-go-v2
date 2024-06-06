// Code generated by smithy-go-codegen DO NOT EDIT.

package personalize

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/personalize/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of dataset export jobs that use the given dataset. When a
// dataset is not specified, all the dataset export jobs associated with the
// account are listed. The response provides the properties for each dataset export
// job, including the Amazon Resource Name (ARN). For more information on dataset
// export jobs, see [CreateDatasetExportJob]. For more information on datasets, see [CreateDataset].
//
// [CreateDataset]: https://docs.aws.amazon.com/personalize/latest/dg/API_CreateDataset.html
// [CreateDatasetExportJob]: https://docs.aws.amazon.com/personalize/latest/dg/API_CreateDatasetExportJob.html
func (c *Client) ListDatasetExportJobs(ctx context.Context, params *ListDatasetExportJobsInput, optFns ...func(*Options)) (*ListDatasetExportJobsOutput, error) {
	if params == nil {
		params = &ListDatasetExportJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDatasetExportJobs", params, optFns, c.addOperationListDatasetExportJobsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDatasetExportJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDatasetExportJobsInput struct {

	// The Amazon Resource Name (ARN) of the dataset to list the dataset export jobs
	// for.
	DatasetArn *string

	// The maximum number of dataset export jobs to return.
	MaxResults *int32

	// A token returned from the previous call to ListDatasetExportJobs for getting
	// the next set of dataset export jobs (if they exist).
	NextToken *string

	noSmithyDocumentSerde
}

type ListDatasetExportJobsOutput struct {

	// The list of dataset export jobs.
	DatasetExportJobs []types.DatasetExportJobSummary

	// A token for getting the next set of dataset export jobs (if they exist).
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListDatasetExportJobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListDatasetExportJobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListDatasetExportJobs{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListDatasetExportJobs"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListDatasetExportJobs(options.Region), middleware.Before); err != nil {
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

// ListDatasetExportJobsAPIClient is a client that implements the
// ListDatasetExportJobs operation.
type ListDatasetExportJobsAPIClient interface {
	ListDatasetExportJobs(context.Context, *ListDatasetExportJobsInput, ...func(*Options)) (*ListDatasetExportJobsOutput, error)
}

var _ ListDatasetExportJobsAPIClient = (*Client)(nil)

// ListDatasetExportJobsPaginatorOptions is the paginator options for
// ListDatasetExportJobs
type ListDatasetExportJobsPaginatorOptions struct {
	// The maximum number of dataset export jobs to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListDatasetExportJobsPaginator is a paginator for ListDatasetExportJobs
type ListDatasetExportJobsPaginator struct {
	options   ListDatasetExportJobsPaginatorOptions
	client    ListDatasetExportJobsAPIClient
	params    *ListDatasetExportJobsInput
	nextToken *string
	firstPage bool
}

// NewListDatasetExportJobsPaginator returns a new ListDatasetExportJobsPaginator
func NewListDatasetExportJobsPaginator(client ListDatasetExportJobsAPIClient, params *ListDatasetExportJobsInput, optFns ...func(*ListDatasetExportJobsPaginatorOptions)) *ListDatasetExportJobsPaginator {
	if params == nil {
		params = &ListDatasetExportJobsInput{}
	}

	options := ListDatasetExportJobsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListDatasetExportJobsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListDatasetExportJobsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListDatasetExportJobs page.
func (p *ListDatasetExportJobsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListDatasetExportJobsOutput, error) {
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

	result, err := p.client.ListDatasetExportJobs(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListDatasetExportJobs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListDatasetExportJobs",
	}
}
