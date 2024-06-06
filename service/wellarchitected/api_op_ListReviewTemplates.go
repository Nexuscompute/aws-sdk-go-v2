// Code generated by smithy-go-codegen DO NOT EDIT.

package wellarchitected

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/wellarchitected/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// List review templates.
func (c *Client) ListReviewTemplates(ctx context.Context, params *ListReviewTemplatesInput, optFns ...func(*Options)) (*ListReviewTemplatesOutput, error) {
	if params == nil {
		params = &ListReviewTemplatesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListReviewTemplates", params, optFns, c.addOperationListReviewTemplatesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListReviewTemplatesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListReviewTemplatesInput struct {

	// The maximum number of results to return for this request.
	MaxResults *int32

	// The token to use to retrieve the next set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListReviewTemplatesOutput struct {

	// The token to use to retrieve the next set of results.
	NextToken *string

	// List of review templates.
	ReviewTemplates []types.ReviewTemplateSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListReviewTemplatesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListReviewTemplates{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListReviewTemplates{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListReviewTemplates"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListReviewTemplates(options.Region), middleware.Before); err != nil {
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

// ListReviewTemplatesAPIClient is a client that implements the
// ListReviewTemplates operation.
type ListReviewTemplatesAPIClient interface {
	ListReviewTemplates(context.Context, *ListReviewTemplatesInput, ...func(*Options)) (*ListReviewTemplatesOutput, error)
}

var _ ListReviewTemplatesAPIClient = (*Client)(nil)

// ListReviewTemplatesPaginatorOptions is the paginator options for
// ListReviewTemplates
type ListReviewTemplatesPaginatorOptions struct {
	// The maximum number of results to return for this request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListReviewTemplatesPaginator is a paginator for ListReviewTemplates
type ListReviewTemplatesPaginator struct {
	options   ListReviewTemplatesPaginatorOptions
	client    ListReviewTemplatesAPIClient
	params    *ListReviewTemplatesInput
	nextToken *string
	firstPage bool
}

// NewListReviewTemplatesPaginator returns a new ListReviewTemplatesPaginator
func NewListReviewTemplatesPaginator(client ListReviewTemplatesAPIClient, params *ListReviewTemplatesInput, optFns ...func(*ListReviewTemplatesPaginatorOptions)) *ListReviewTemplatesPaginator {
	if params == nil {
		params = &ListReviewTemplatesInput{}
	}

	options := ListReviewTemplatesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListReviewTemplatesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListReviewTemplatesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListReviewTemplates page.
func (p *ListReviewTemplatesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListReviewTemplatesOutput, error) {
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

	result, err := p.client.ListReviewTemplates(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListReviewTemplates(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListReviewTemplates",
	}
}
