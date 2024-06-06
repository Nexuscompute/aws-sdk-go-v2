// Code generated by smithy-go-codegen DO NOT EDIT.

package outposts

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/outposts/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the items in the catalog.
//
// Use filters to return specific results. If you specify multiple filters, the
// results include only the resources that match all of the specified filters. For
// a filter where you can specify multiple values, the results include items that
// match any of the values that you specify for the filter.
func (c *Client) ListCatalogItems(ctx context.Context, params *ListCatalogItemsInput, optFns ...func(*Options)) (*ListCatalogItemsOutput, error) {
	if params == nil {
		params = &ListCatalogItemsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListCatalogItems", params, optFns, c.addOperationListCatalogItemsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListCatalogItemsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListCatalogItemsInput struct {

	// Filters the results by EC2 family (for example, M5).
	EC2FamilyFilter []string

	// Filters the results by item class.
	ItemClassFilter []types.CatalogItemClass

	// The maximum page size.
	MaxResults *int32

	// The pagination token.
	NextToken *string

	// Filters the results by storage option.
	SupportedStorageFilter []types.SupportedStorageEnum

	noSmithyDocumentSerde
}

type ListCatalogItemsOutput struct {

	// Information about the catalog items.
	CatalogItems []types.CatalogItem

	// The pagination token.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListCatalogItemsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListCatalogItems{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListCatalogItems{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListCatalogItems"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListCatalogItems(options.Region), middleware.Before); err != nil {
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

// ListCatalogItemsAPIClient is a client that implements the ListCatalogItems
// operation.
type ListCatalogItemsAPIClient interface {
	ListCatalogItems(context.Context, *ListCatalogItemsInput, ...func(*Options)) (*ListCatalogItemsOutput, error)
}

var _ ListCatalogItemsAPIClient = (*Client)(nil)

// ListCatalogItemsPaginatorOptions is the paginator options for ListCatalogItems
type ListCatalogItemsPaginatorOptions struct {
	// The maximum page size.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListCatalogItemsPaginator is a paginator for ListCatalogItems
type ListCatalogItemsPaginator struct {
	options   ListCatalogItemsPaginatorOptions
	client    ListCatalogItemsAPIClient
	params    *ListCatalogItemsInput
	nextToken *string
	firstPage bool
}

// NewListCatalogItemsPaginator returns a new ListCatalogItemsPaginator
func NewListCatalogItemsPaginator(client ListCatalogItemsAPIClient, params *ListCatalogItemsInput, optFns ...func(*ListCatalogItemsPaginatorOptions)) *ListCatalogItemsPaginator {
	if params == nil {
		params = &ListCatalogItemsInput{}
	}

	options := ListCatalogItemsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListCatalogItemsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListCatalogItemsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListCatalogItems page.
func (p *ListCatalogItemsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListCatalogItemsOutput, error) {
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

	result, err := p.client.ListCatalogItems(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListCatalogItems(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListCatalogItems",
	}
}
