// Code generated by smithy-go-codegen DO NOT EDIT.

package ssmincidents

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/ssmincidents/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// List all related items for an incident record.
func (c *Client) ListRelatedItems(ctx context.Context, params *ListRelatedItemsInput, optFns ...func(*Options)) (*ListRelatedItemsOutput, error) {
	if params == nil {
		params = &ListRelatedItemsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListRelatedItems", params, optFns, c.addOperationListRelatedItemsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListRelatedItemsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListRelatedItemsInput struct {

	// The Amazon Resource Name (ARN) of the incident record containing the listed
	// related items.
	//
	// This member is required.
	IncidentRecordArn *string

	// The maximum number of related items per page.
	MaxResults *int32

	// The pagination token for the next set of items to return. (You received this
	// token from a previous call.)
	NextToken *string

	noSmithyDocumentSerde
}

type ListRelatedItemsOutput struct {

	// Details about each related item.
	//
	// This member is required.
	RelatedItems []types.RelatedItem

	// The pagination token to use when requesting the next set of items. If there are
	// no additional items to return, the string is null.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListRelatedItemsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListRelatedItems{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListRelatedItems{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListRelatedItems"); err != nil {
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
	if err = addOpListRelatedItemsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListRelatedItems(options.Region), middleware.Before); err != nil {
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

// ListRelatedItemsAPIClient is a client that implements the ListRelatedItems
// operation.
type ListRelatedItemsAPIClient interface {
	ListRelatedItems(context.Context, *ListRelatedItemsInput, ...func(*Options)) (*ListRelatedItemsOutput, error)
}

var _ ListRelatedItemsAPIClient = (*Client)(nil)

// ListRelatedItemsPaginatorOptions is the paginator options for ListRelatedItems
type ListRelatedItemsPaginatorOptions struct {
	// The maximum number of related items per page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListRelatedItemsPaginator is a paginator for ListRelatedItems
type ListRelatedItemsPaginator struct {
	options   ListRelatedItemsPaginatorOptions
	client    ListRelatedItemsAPIClient
	params    *ListRelatedItemsInput
	nextToken *string
	firstPage bool
}

// NewListRelatedItemsPaginator returns a new ListRelatedItemsPaginator
func NewListRelatedItemsPaginator(client ListRelatedItemsAPIClient, params *ListRelatedItemsInput, optFns ...func(*ListRelatedItemsPaginatorOptions)) *ListRelatedItemsPaginator {
	if params == nil {
		params = &ListRelatedItemsInput{}
	}

	options := ListRelatedItemsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListRelatedItemsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListRelatedItemsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListRelatedItems page.
func (p *ListRelatedItemsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListRelatedItemsOutput, error) {
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

	result, err := p.client.ListRelatedItems(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListRelatedItems(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListRelatedItems",
	}
}
