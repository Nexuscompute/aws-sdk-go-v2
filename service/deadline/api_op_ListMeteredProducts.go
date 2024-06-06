// Code generated by smithy-go-codegen DO NOT EDIT.

package deadline

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/deadline/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists metered products.
func (c *Client) ListMeteredProducts(ctx context.Context, params *ListMeteredProductsInput, optFns ...func(*Options)) (*ListMeteredProductsOutput, error) {
	if params == nil {
		params = &ListMeteredProductsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListMeteredProducts", params, optFns, c.addOperationListMeteredProductsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListMeteredProductsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListMeteredProductsInput struct {

	// The license endpoint ID to include on the list of metered products.
	//
	// This member is required.
	LicenseEndpointId *string

	// The maximum number of results to return. Use this parameter with NextToken to
	// get results as a set of sequential pages.
	MaxResults *int32

	// The token for the next set of results, or null to start from the beginning.
	NextToken *string

	noSmithyDocumentSerde
}

type ListMeteredProductsOutput struct {

	// The metered products to list.
	//
	// This member is required.
	MeteredProducts []types.MeteredProductSummary

	// If Deadline Cloud returns nextToken , then there are more results available. The
	// value of nextToken is a unique pagination token for each page. To retrieve the
	// next page, call the operation again using the returned token. Keep all other
	// arguments unchanged. If no results remain, then nextToken is set to null . Each
	// pagination token expires after 24 hours. If you provide a token that isn't
	// valid, then you receive an HTTP 400 ValidationException error.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListMeteredProductsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListMeteredProducts{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListMeteredProducts{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListMeteredProducts"); err != nil {
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
	if err = addEndpointPrefix_opListMeteredProductsMiddleware(stack); err != nil {
		return err
	}
	if err = addOpListMeteredProductsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListMeteredProducts(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opListMeteredProductsMiddleware struct {
}

func (*endpointPrefix_opListMeteredProductsMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opListMeteredProductsMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleFinalize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "management." + req.URL.Host

	return next.HandleFinalize(ctx, in)
}
func addEndpointPrefix_opListMeteredProductsMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opListMeteredProductsMiddleware{}, "ResolveEndpointV2", middleware.After)
}

// ListMeteredProductsAPIClient is a client that implements the
// ListMeteredProducts operation.
type ListMeteredProductsAPIClient interface {
	ListMeteredProducts(context.Context, *ListMeteredProductsInput, ...func(*Options)) (*ListMeteredProductsOutput, error)
}

var _ ListMeteredProductsAPIClient = (*Client)(nil)

// ListMeteredProductsPaginatorOptions is the paginator options for
// ListMeteredProducts
type ListMeteredProductsPaginatorOptions struct {
	// The maximum number of results to return. Use this parameter with NextToken to
	// get results as a set of sequential pages.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListMeteredProductsPaginator is a paginator for ListMeteredProducts
type ListMeteredProductsPaginator struct {
	options   ListMeteredProductsPaginatorOptions
	client    ListMeteredProductsAPIClient
	params    *ListMeteredProductsInput
	nextToken *string
	firstPage bool
}

// NewListMeteredProductsPaginator returns a new ListMeteredProductsPaginator
func NewListMeteredProductsPaginator(client ListMeteredProductsAPIClient, params *ListMeteredProductsInput, optFns ...func(*ListMeteredProductsPaginatorOptions)) *ListMeteredProductsPaginator {
	if params == nil {
		params = &ListMeteredProductsInput{}
	}

	options := ListMeteredProductsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListMeteredProductsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListMeteredProductsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListMeteredProducts page.
func (p *ListMeteredProductsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListMeteredProductsOutput, error) {
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

	result, err := p.client.ListMeteredProducts(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListMeteredProducts(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListMeteredProducts",
	}
}
