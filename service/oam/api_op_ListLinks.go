// Code generated by smithy-go-codegen DO NOT EDIT.

package oam

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/oam/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Use this operation in a source account to return a list of links to monitoring
// account sinks that this source account has.
//
// To find a list of links for one monitoring account sink, use [ListAttachedLinks] from within the
// monitoring account.
//
// [ListAttachedLinks]: https://docs.aws.amazon.com/OAM/latest/APIReference/API_ListAttachedLinks.html
func (c *Client) ListLinks(ctx context.Context, params *ListLinksInput, optFns ...func(*Options)) (*ListLinksOutput, error) {
	if params == nil {
		params = &ListLinksInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListLinks", params, optFns, c.addOperationListLinksMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListLinksOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListLinksInput struct {

	// Limits the number of returned links to the specified number.
	MaxResults *int32

	// The token for the next set of items to return. You received this token from a
	// previous call.
	NextToken *string

	noSmithyDocumentSerde
}

type ListLinksOutput struct {

	// An array of structures that contain the information about the returned links.
	//
	// This member is required.
	Items []types.ListLinksItem

	// The token to use when requesting the next set of links.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListLinksMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListLinks{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListLinks{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListLinks"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListLinks(options.Region), middleware.Before); err != nil {
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

// ListLinksPaginatorOptions is the paginator options for ListLinks
type ListLinksPaginatorOptions struct {
	// Limits the number of returned links to the specified number.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListLinksPaginator is a paginator for ListLinks
type ListLinksPaginator struct {
	options   ListLinksPaginatorOptions
	client    ListLinksAPIClient
	params    *ListLinksInput
	nextToken *string
	firstPage bool
}

// NewListLinksPaginator returns a new ListLinksPaginator
func NewListLinksPaginator(client ListLinksAPIClient, params *ListLinksInput, optFns ...func(*ListLinksPaginatorOptions)) *ListLinksPaginator {
	if params == nil {
		params = &ListLinksInput{}
	}

	options := ListLinksPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListLinksPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListLinksPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListLinks page.
func (p *ListLinksPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListLinksOutput, error) {
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
	result, err := p.client.ListLinks(ctx, &params, optFns...)
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

// ListLinksAPIClient is a client that implements the ListLinks operation.
type ListLinksAPIClient interface {
	ListLinks(context.Context, *ListLinksInput, ...func(*Options)) (*ListLinksOutput, error)
}

var _ ListLinksAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListLinks(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListLinks",
	}
}
