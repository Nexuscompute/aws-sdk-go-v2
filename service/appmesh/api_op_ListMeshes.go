// Code generated by smithy-go-codegen DO NOT EDIT.

package appmesh

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/appmesh/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of existing service meshes.
func (c *Client) ListMeshes(ctx context.Context, params *ListMeshesInput, optFns ...func(*Options)) (*ListMeshesOutput, error) {
	if params == nil {
		params = &ListMeshesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListMeshes", params, optFns, c.addOperationListMeshesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListMeshesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListMeshesInput struct {

	// The maximum number of results returned by ListMeshes in paginated output. When
	// you use this parameter, ListMeshes returns only limit results in a single page
	// along with a nextToken response element. You can see the remaining results of
	// the initial request by sending another ListMeshes request with the returned
	// nextToken value. This value can be between 1 and 100. If you don't use this
	// parameter, ListMeshes returns up to 100 results and a nextToken value if
	// applicable.
	Limit *int32

	// The nextToken value returned from a previous paginated ListMeshes request where
	// limit was used and the results exceeded the value of that parameter. Pagination
	// continues from the end of the previous results that returned the nextToken
	// value.
	//
	// This token should be treated as an opaque identifier that is used only to
	// retrieve the next items in a list and not for other programmatic purposes.
	NextToken *string

	noSmithyDocumentSerde
}

type ListMeshesOutput struct {

	// The list of existing service meshes.
	//
	// This member is required.
	Meshes []types.MeshRef

	// The nextToken value to include in a future ListMeshes request. When the results
	// of a ListMeshes request exceed limit , you can use this value to retrieve the
	// next page of results. This value is null when there are no more results to
	// return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListMeshesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListMeshes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListMeshes{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListMeshes"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListMeshes(options.Region), middleware.Before); err != nil {
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

// ListMeshesAPIClient is a client that implements the ListMeshes operation.
type ListMeshesAPIClient interface {
	ListMeshes(context.Context, *ListMeshesInput, ...func(*Options)) (*ListMeshesOutput, error)
}

var _ ListMeshesAPIClient = (*Client)(nil)

// ListMeshesPaginatorOptions is the paginator options for ListMeshes
type ListMeshesPaginatorOptions struct {
	// The maximum number of results returned by ListMeshes in paginated output. When
	// you use this parameter, ListMeshes returns only limit results in a single page
	// along with a nextToken response element. You can see the remaining results of
	// the initial request by sending another ListMeshes request with the returned
	// nextToken value. This value can be between 1 and 100. If you don't use this
	// parameter, ListMeshes returns up to 100 results and a nextToken value if
	// applicable.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListMeshesPaginator is a paginator for ListMeshes
type ListMeshesPaginator struct {
	options   ListMeshesPaginatorOptions
	client    ListMeshesAPIClient
	params    *ListMeshesInput
	nextToken *string
	firstPage bool
}

// NewListMeshesPaginator returns a new ListMeshesPaginator
func NewListMeshesPaginator(client ListMeshesAPIClient, params *ListMeshesInput, optFns ...func(*ListMeshesPaginatorOptions)) *ListMeshesPaginator {
	if params == nil {
		params = &ListMeshesInput{}
	}

	options := ListMeshesPaginatorOptions{}
	if params.Limit != nil {
		options.Limit = *params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListMeshesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListMeshesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListMeshes page.
func (p *ListMeshesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListMeshesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.Limit = limit

	result, err := p.client.ListMeshes(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListMeshes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListMeshes",
	}
}
