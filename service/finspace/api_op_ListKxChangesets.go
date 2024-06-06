// Code generated by smithy-go-codegen DO NOT EDIT.

package finspace

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/finspace/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of all the changesets for a database.
func (c *Client) ListKxChangesets(ctx context.Context, params *ListKxChangesetsInput, optFns ...func(*Options)) (*ListKxChangesetsOutput, error) {
	if params == nil {
		params = &ListKxChangesetsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListKxChangesets", params, optFns, c.addOperationListKxChangesetsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListKxChangesetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListKxChangesetsInput struct {

	// The name of the kdb database.
	//
	// This member is required.
	DatabaseName *string

	// A unique identifier for the kdb environment.
	//
	// This member is required.
	EnvironmentId *string

	// The maximum number of results to return in this request.
	MaxResults int32

	// A token that indicates where a results page should begin.
	NextToken *string

	noSmithyDocumentSerde
}

type ListKxChangesetsOutput struct {

	// A list of changesets for a database.
	KxChangesets []types.KxChangesetListEntry

	// A token that indicates where a results page should begin.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListKxChangesetsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListKxChangesets{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListKxChangesets{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListKxChangesets"); err != nil {
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
	if err = addRestJsonContentTypeCustomization(stack); err != nil {
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
	if err = addOpListKxChangesetsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListKxChangesets(options.Region), middleware.Before); err != nil {
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

// ListKxChangesetsAPIClient is a client that implements the ListKxChangesets
// operation.
type ListKxChangesetsAPIClient interface {
	ListKxChangesets(context.Context, *ListKxChangesetsInput, ...func(*Options)) (*ListKxChangesetsOutput, error)
}

var _ ListKxChangesetsAPIClient = (*Client)(nil)

// ListKxChangesetsPaginatorOptions is the paginator options for ListKxChangesets
type ListKxChangesetsPaginatorOptions struct {
	// The maximum number of results to return in this request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListKxChangesetsPaginator is a paginator for ListKxChangesets
type ListKxChangesetsPaginator struct {
	options   ListKxChangesetsPaginatorOptions
	client    ListKxChangesetsAPIClient
	params    *ListKxChangesetsInput
	nextToken *string
	firstPage bool
}

// NewListKxChangesetsPaginator returns a new ListKxChangesetsPaginator
func NewListKxChangesetsPaginator(client ListKxChangesetsAPIClient, params *ListKxChangesetsInput, optFns ...func(*ListKxChangesetsPaginatorOptions)) *ListKxChangesetsPaginator {
	if params == nil {
		params = &ListKxChangesetsInput{}
	}

	options := ListKxChangesetsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListKxChangesetsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListKxChangesetsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListKxChangesets page.
func (p *ListKxChangesetsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListKxChangesetsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListKxChangesets(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListKxChangesets(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListKxChangesets",
	}
}
