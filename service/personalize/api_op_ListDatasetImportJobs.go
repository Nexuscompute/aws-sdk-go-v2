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

// Returns a list of dataset import jobs that use the given dataset. When a
// dataset is not specified, all the dataset import jobs associated with the
// account are listed. The response provides the properties for each dataset import
// job, including the Amazon Resource Name (ARN). For more information on dataset
// import jobs, see [CreateDatasetImportJob]. For more information on datasets, see [CreateDataset].
//
// [CreateDataset]: https://docs.aws.amazon.com/personalize/latest/dg/API_CreateDataset.html
// [CreateDatasetImportJob]: https://docs.aws.amazon.com/personalize/latest/dg/API_CreateDatasetImportJob.html
func (c *Client) ListDatasetImportJobs(ctx context.Context, params *ListDatasetImportJobsInput, optFns ...func(*Options)) (*ListDatasetImportJobsOutput, error) {
	if params == nil {
		params = &ListDatasetImportJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDatasetImportJobs", params, optFns, c.addOperationListDatasetImportJobsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDatasetImportJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDatasetImportJobsInput struct {

	// The Amazon Resource Name (ARN) of the dataset to list the dataset import jobs
	// for.
	DatasetArn *string

	// The maximum number of dataset import jobs to return.
	MaxResults *int32

	// A token returned from the previous call to ListDatasetImportJobs for getting
	// the next set of dataset import jobs (if they exist).
	NextToken *string

	noSmithyDocumentSerde
}

type ListDatasetImportJobsOutput struct {

	// The list of dataset import jobs.
	DatasetImportJobs []types.DatasetImportJobSummary

	// A token for getting the next set of dataset import jobs (if they exist).
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListDatasetImportJobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListDatasetImportJobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListDatasetImportJobs{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListDatasetImportJobs"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListDatasetImportJobs(options.Region), middleware.Before); err != nil {
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

// ListDatasetImportJobsAPIClient is a client that implements the
// ListDatasetImportJobs operation.
type ListDatasetImportJobsAPIClient interface {
	ListDatasetImportJobs(context.Context, *ListDatasetImportJobsInput, ...func(*Options)) (*ListDatasetImportJobsOutput, error)
}

var _ ListDatasetImportJobsAPIClient = (*Client)(nil)

// ListDatasetImportJobsPaginatorOptions is the paginator options for
// ListDatasetImportJobs
type ListDatasetImportJobsPaginatorOptions struct {
	// The maximum number of dataset import jobs to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListDatasetImportJobsPaginator is a paginator for ListDatasetImportJobs
type ListDatasetImportJobsPaginator struct {
	options   ListDatasetImportJobsPaginatorOptions
	client    ListDatasetImportJobsAPIClient
	params    *ListDatasetImportJobsInput
	nextToken *string
	firstPage bool
}

// NewListDatasetImportJobsPaginator returns a new ListDatasetImportJobsPaginator
func NewListDatasetImportJobsPaginator(client ListDatasetImportJobsAPIClient, params *ListDatasetImportJobsInput, optFns ...func(*ListDatasetImportJobsPaginatorOptions)) *ListDatasetImportJobsPaginator {
	if params == nil {
		params = &ListDatasetImportJobsInput{}
	}

	options := ListDatasetImportJobsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListDatasetImportJobsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListDatasetImportJobsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListDatasetImportJobs page.
func (p *ListDatasetImportJobsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListDatasetImportJobsOutput, error) {
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

	result, err := p.client.ListDatasetImportJobs(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListDatasetImportJobs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListDatasetImportJobs",
	}
}
