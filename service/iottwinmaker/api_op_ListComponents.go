// Code generated by smithy-go-codegen DO NOT EDIT.

package iottwinmaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/iottwinmaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This API lists the components of an entity.
func (c *Client) ListComponents(ctx context.Context, params *ListComponentsInput, optFns ...func(*Options)) (*ListComponentsOutput, error) {
	if params == nil {
		params = &ListComponentsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListComponents", params, optFns, c.addOperationListComponentsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListComponentsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListComponentsInput struct {

	// The ID for the entity whose metadata (component/properties) is returned by the
	// operation.
	//
	// This member is required.
	EntityId *string

	// The workspace ID.
	//
	// This member is required.
	WorkspaceId *string

	// This string specifies the path to the composite component, starting from the
	// top-level component.
	ComponentPath *string

	// The maximum number of results returned at one time. The default is 25.
	MaxResults *int32

	// The string that specifies the next page of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListComponentsOutput struct {

	// A list of objects that contain information about the components.
	//
	// This member is required.
	ComponentSummaries []types.ComponentSummary

	// The string that specifies the next page of component results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListComponentsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListComponents{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListComponents{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListComponents"); err != nil {
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
	if err = addEndpointPrefix_opListComponentsMiddleware(stack); err != nil {
		return err
	}
	if err = addOpListComponentsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListComponents(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opListComponentsMiddleware struct {
}

func (*endpointPrefix_opListComponentsMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opListComponentsMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleFinalize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "api." + req.URL.Host

	return next.HandleFinalize(ctx, in)
}
func addEndpointPrefix_opListComponentsMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opListComponentsMiddleware{}, "ResolveEndpointV2", middleware.After)
}

// ListComponentsAPIClient is a client that implements the ListComponents
// operation.
type ListComponentsAPIClient interface {
	ListComponents(context.Context, *ListComponentsInput, ...func(*Options)) (*ListComponentsOutput, error)
}

var _ ListComponentsAPIClient = (*Client)(nil)

// ListComponentsPaginatorOptions is the paginator options for ListComponents
type ListComponentsPaginatorOptions struct {
	// The maximum number of results returned at one time. The default is 25.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListComponentsPaginator is a paginator for ListComponents
type ListComponentsPaginator struct {
	options   ListComponentsPaginatorOptions
	client    ListComponentsAPIClient
	params    *ListComponentsInput
	nextToken *string
	firstPage bool
}

// NewListComponentsPaginator returns a new ListComponentsPaginator
func NewListComponentsPaginator(client ListComponentsAPIClient, params *ListComponentsInput, optFns ...func(*ListComponentsPaginatorOptions)) *ListComponentsPaginator {
	if params == nil {
		params = &ListComponentsInput{}
	}

	options := ListComponentsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListComponentsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListComponentsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListComponents page.
func (p *ListComponentsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListComponentsOutput, error) {
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

	result, err := p.client.ListComponents(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListComponents(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListComponents",
	}
}
