// Code generated by smithy-go-codegen DO NOT EDIT.

package bedrock

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/bedrock/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns a list of model customization jobs that you have submitted. You can
// filter the jobs to return based on one or more criteria.
//
// For more information, see [Custom models] in the Amazon Bedrock User Guide.
//
// [Custom models]: https://docs.aws.amazon.com/bedrock/latest/userguide/custom-models.html
func (c *Client) ListModelCustomizationJobs(ctx context.Context, params *ListModelCustomizationJobsInput, optFns ...func(*Options)) (*ListModelCustomizationJobsOutput, error) {
	if params == nil {
		params = &ListModelCustomizationJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListModelCustomizationJobs", params, optFns, c.addOperationListModelCustomizationJobsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListModelCustomizationJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListModelCustomizationJobsInput struct {

	// Return customization jobs created after the specified time.
	CreationTimeAfter *time.Time

	// Return customization jobs created before the specified time.
	CreationTimeBefore *time.Time

	// Maximum number of results to return in the response.
	MaxResults *int32

	// Return customization jobs only if the job name contains these characters.
	NameContains *string

	// Continuation token from the previous response, for Amazon Bedrock to list the
	// next set of results.
	NextToken *string

	// The field to sort by in the returned list of jobs.
	SortBy types.SortJobsBy

	// The sort order of the results.
	SortOrder types.SortOrder

	// Return customization jobs with the specified status.
	StatusEquals types.FineTuningJobStatus

	noSmithyDocumentSerde
}

type ListModelCustomizationJobsOutput struct {

	// Job summaries.
	ModelCustomizationJobSummaries []types.ModelCustomizationJobSummary

	// Page continuation token to use in the next request.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListModelCustomizationJobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListModelCustomizationJobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListModelCustomizationJobs{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListModelCustomizationJobs"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListModelCustomizationJobs(options.Region), middleware.Before); err != nil {
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

// ListModelCustomizationJobsAPIClient is a client that implements the
// ListModelCustomizationJobs operation.
type ListModelCustomizationJobsAPIClient interface {
	ListModelCustomizationJobs(context.Context, *ListModelCustomizationJobsInput, ...func(*Options)) (*ListModelCustomizationJobsOutput, error)
}

var _ ListModelCustomizationJobsAPIClient = (*Client)(nil)

// ListModelCustomizationJobsPaginatorOptions is the paginator options for
// ListModelCustomizationJobs
type ListModelCustomizationJobsPaginatorOptions struct {
	// Maximum number of results to return in the response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListModelCustomizationJobsPaginator is a paginator for
// ListModelCustomizationJobs
type ListModelCustomizationJobsPaginator struct {
	options   ListModelCustomizationJobsPaginatorOptions
	client    ListModelCustomizationJobsAPIClient
	params    *ListModelCustomizationJobsInput
	nextToken *string
	firstPage bool
}

// NewListModelCustomizationJobsPaginator returns a new
// ListModelCustomizationJobsPaginator
func NewListModelCustomizationJobsPaginator(client ListModelCustomizationJobsAPIClient, params *ListModelCustomizationJobsInput, optFns ...func(*ListModelCustomizationJobsPaginatorOptions)) *ListModelCustomizationJobsPaginator {
	if params == nil {
		params = &ListModelCustomizationJobsInput{}
	}

	options := ListModelCustomizationJobsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListModelCustomizationJobsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListModelCustomizationJobsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListModelCustomizationJobs page.
func (p *ListModelCustomizationJobsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListModelCustomizationJobsOutput, error) {
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

	result, err := p.client.ListModelCustomizationJobs(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListModelCustomizationJobs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListModelCustomizationJobs",
	}
}
