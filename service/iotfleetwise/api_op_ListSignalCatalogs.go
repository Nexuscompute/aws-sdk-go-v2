// Code generated by smithy-go-codegen DO NOT EDIT.

package iotfleetwise

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/iotfleetwise/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

//	Lists all the created signal catalogs in an Amazon Web Services account.
//
// You can use to list information about each signal (node) specified in a signal
// catalog.
//
// This API operation uses pagination. Specify the nextToken parameter in the
// request to return more results.
func (c *Client) ListSignalCatalogs(ctx context.Context, params *ListSignalCatalogsInput, optFns ...func(*Options)) (*ListSignalCatalogsOutput, error) {
	if params == nil {
		params = &ListSignalCatalogsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListSignalCatalogs", params, optFns, c.addOperationListSignalCatalogsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListSignalCatalogsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSignalCatalogsInput struct {

	// The maximum number of items to return, between 1 and 100, inclusive.
	MaxResults *int32

	// A pagination token for the next set of results.
	//
	// If the results of a search are large, only a portion of the results are
	// returned, and a nextToken pagination token is returned in the response. To
	// retrieve the next set of results, reissue the search request and include the
	// returned token. When all results have been returned, the response does not
	// contain a pagination token value.
	NextToken *string

	noSmithyDocumentSerde
}

type ListSignalCatalogsOutput struct {

	//  The token to retrieve the next set of results, or null if there are no more
	// results.
	NextToken *string

	//  A list of information about each signal catalog.
	Summaries []types.SignalCatalogSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListSignalCatalogsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpListSignalCatalogs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpListSignalCatalogs{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListSignalCatalogs"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListSignalCatalogs(options.Region), middleware.Before); err != nil {
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

// ListSignalCatalogsPaginatorOptions is the paginator options for
// ListSignalCatalogs
type ListSignalCatalogsPaginatorOptions struct {
	// The maximum number of items to return, between 1 and 100, inclusive.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListSignalCatalogsPaginator is a paginator for ListSignalCatalogs
type ListSignalCatalogsPaginator struct {
	options   ListSignalCatalogsPaginatorOptions
	client    ListSignalCatalogsAPIClient
	params    *ListSignalCatalogsInput
	nextToken *string
	firstPage bool
}

// NewListSignalCatalogsPaginator returns a new ListSignalCatalogsPaginator
func NewListSignalCatalogsPaginator(client ListSignalCatalogsAPIClient, params *ListSignalCatalogsInput, optFns ...func(*ListSignalCatalogsPaginatorOptions)) *ListSignalCatalogsPaginator {
	if params == nil {
		params = &ListSignalCatalogsInput{}
	}

	options := ListSignalCatalogsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListSignalCatalogsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListSignalCatalogsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListSignalCatalogs page.
func (p *ListSignalCatalogsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListSignalCatalogsOutput, error) {
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
	result, err := p.client.ListSignalCatalogs(ctx, &params, optFns...)
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

// ListSignalCatalogsAPIClient is a client that implements the ListSignalCatalogs
// operation.
type ListSignalCatalogsAPIClient interface {
	ListSignalCatalogs(context.Context, *ListSignalCatalogsInput, ...func(*Options)) (*ListSignalCatalogsOutput, error)
}

var _ ListSignalCatalogsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListSignalCatalogs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListSignalCatalogs",
	}
}
