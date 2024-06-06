// Code generated by smithy-go-codegen DO NOT EDIT.

package codestarconnections

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/codestarconnections/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the repository links created for connections in your account.
func (c *Client) ListRepositoryLinks(ctx context.Context, params *ListRepositoryLinksInput, optFns ...func(*Options)) (*ListRepositoryLinksOutput, error) {
	if params == nil {
		params = &ListRepositoryLinksInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListRepositoryLinks", params, optFns, c.addOperationListRepositoryLinksMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListRepositoryLinksOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListRepositoryLinksInput struct {

	//  A non-zero, non-negative integer used to limit the number of returned results.
	MaxResults int32

	//  An enumeration token that, when provided in a request, returns the next batch
	// of the results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListRepositoryLinksOutput struct {

	// Lists the repository links called by the list repository links operation.
	//
	// This member is required.
	RepositoryLinks []types.RepositoryLinkInfo

	// An enumeration token that allows the operation to batch the results of the
	// operation.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListRepositoryLinksMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpListRepositoryLinks{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpListRepositoryLinks{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListRepositoryLinks"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListRepositoryLinks(options.Region), middleware.Before); err != nil {
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

// ListRepositoryLinksAPIClient is a client that implements the
// ListRepositoryLinks operation.
type ListRepositoryLinksAPIClient interface {
	ListRepositoryLinks(context.Context, *ListRepositoryLinksInput, ...func(*Options)) (*ListRepositoryLinksOutput, error)
}

var _ ListRepositoryLinksAPIClient = (*Client)(nil)

// ListRepositoryLinksPaginatorOptions is the paginator options for
// ListRepositoryLinks
type ListRepositoryLinksPaginatorOptions struct {
	//  A non-zero, non-negative integer used to limit the number of returned results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListRepositoryLinksPaginator is a paginator for ListRepositoryLinks
type ListRepositoryLinksPaginator struct {
	options   ListRepositoryLinksPaginatorOptions
	client    ListRepositoryLinksAPIClient
	params    *ListRepositoryLinksInput
	nextToken *string
	firstPage bool
}

// NewListRepositoryLinksPaginator returns a new ListRepositoryLinksPaginator
func NewListRepositoryLinksPaginator(client ListRepositoryLinksAPIClient, params *ListRepositoryLinksInput, optFns ...func(*ListRepositoryLinksPaginatorOptions)) *ListRepositoryLinksPaginator {
	if params == nil {
		params = &ListRepositoryLinksInput{}
	}

	options := ListRepositoryLinksPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListRepositoryLinksPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListRepositoryLinksPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListRepositoryLinks page.
func (p *ListRepositoryLinksPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListRepositoryLinksOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListRepositoryLinks(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListRepositoryLinks(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListRepositoryLinks",
	}
}
