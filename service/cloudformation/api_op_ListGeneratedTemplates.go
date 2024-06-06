// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudformation

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists your generated templates in this Region.
func (c *Client) ListGeneratedTemplates(ctx context.Context, params *ListGeneratedTemplatesInput, optFns ...func(*Options)) (*ListGeneratedTemplatesOutput, error) {
	if params == nil {
		params = &ListGeneratedTemplatesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListGeneratedTemplates", params, optFns, c.addOperationListGeneratedTemplatesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListGeneratedTemplatesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListGeneratedTemplatesInput struct {

	//  If the number of available results exceeds this maximum, the response includes
	// a NextToken value that you can use for the NextToken parameter to get the next
	// set of results. By default the ListGeneratedTemplates API action will return at
	// most 50 results in each response. The maximum value is 100.
	MaxResults *int32

	// A string that identifies the next page of resource scan results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListGeneratedTemplatesOutput struct {

	// If the request doesn't return all the remaining results, NextToken is set to a
	// token. To retrieve the next set of results, call ListGeneratedTemplates again
	// and use that value for the NextToken parameter. If the request returns all
	// results, NextToken is set to an empty string.
	NextToken *string

	// A list of summaries of the generated templates.
	Summaries []types.TemplateSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListGeneratedTemplatesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpListGeneratedTemplates{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpListGeneratedTemplates{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListGeneratedTemplates"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListGeneratedTemplates(options.Region), middleware.Before); err != nil {
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

// ListGeneratedTemplatesAPIClient is a client that implements the
// ListGeneratedTemplates operation.
type ListGeneratedTemplatesAPIClient interface {
	ListGeneratedTemplates(context.Context, *ListGeneratedTemplatesInput, ...func(*Options)) (*ListGeneratedTemplatesOutput, error)
}

var _ ListGeneratedTemplatesAPIClient = (*Client)(nil)

// ListGeneratedTemplatesPaginatorOptions is the paginator options for
// ListGeneratedTemplates
type ListGeneratedTemplatesPaginatorOptions struct {
	//  If the number of available results exceeds this maximum, the response includes
	// a NextToken value that you can use for the NextToken parameter to get the next
	// set of results. By default the ListGeneratedTemplates API action will return at
	// most 50 results in each response. The maximum value is 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListGeneratedTemplatesPaginator is a paginator for ListGeneratedTemplates
type ListGeneratedTemplatesPaginator struct {
	options   ListGeneratedTemplatesPaginatorOptions
	client    ListGeneratedTemplatesAPIClient
	params    *ListGeneratedTemplatesInput
	nextToken *string
	firstPage bool
}

// NewListGeneratedTemplatesPaginator returns a new ListGeneratedTemplatesPaginator
func NewListGeneratedTemplatesPaginator(client ListGeneratedTemplatesAPIClient, params *ListGeneratedTemplatesInput, optFns ...func(*ListGeneratedTemplatesPaginatorOptions)) *ListGeneratedTemplatesPaginator {
	if params == nil {
		params = &ListGeneratedTemplatesInput{}
	}

	options := ListGeneratedTemplatesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListGeneratedTemplatesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListGeneratedTemplatesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListGeneratedTemplates page.
func (p *ListGeneratedTemplatesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListGeneratedTemplatesOutput, error) {
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

	result, err := p.client.ListGeneratedTemplates(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListGeneratedTemplates(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListGeneratedTemplates",
	}
}
