// Code generated by smithy-go-codegen DO NOT EDIT.

package worklink

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/worklink/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a list of website authorization providers associated with a specified
// fleet.
//
// Deprecated: Amazon WorkLink is no longer supported. This will be removed in a
// future version of the SDK.
func (c *Client) ListWebsiteAuthorizationProviders(ctx context.Context, params *ListWebsiteAuthorizationProvidersInput, optFns ...func(*Options)) (*ListWebsiteAuthorizationProvidersOutput, error) {
	if params == nil {
		params = &ListWebsiteAuthorizationProvidersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListWebsiteAuthorizationProviders", params, optFns, c.addOperationListWebsiteAuthorizationProvidersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListWebsiteAuthorizationProvidersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListWebsiteAuthorizationProvidersInput struct {

	// The ARN of the fleet.
	//
	// This member is required.
	FleetArn *string

	// The maximum number of results to be included in the next page.
	MaxResults *int32

	// The pagination token to use to retrieve the next page of results for this
	// operation. If this value is null, it retrieves the first page.
	NextToken *string

	noSmithyDocumentSerde
}

type ListWebsiteAuthorizationProvidersOutput struct {

	// The pagination token to use to retrieve the next page of results for this
	// operation. If this value is null, it retrieves the first page.
	NextToken *string

	// The website authorization providers.
	WebsiteAuthorizationProviders []types.WebsiteAuthorizationProviderSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListWebsiteAuthorizationProvidersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListWebsiteAuthorizationProviders{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListWebsiteAuthorizationProviders{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListWebsiteAuthorizationProviders"); err != nil {
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
	if err = addOpListWebsiteAuthorizationProvidersValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListWebsiteAuthorizationProviders(options.Region), middleware.Before); err != nil {
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

// ListWebsiteAuthorizationProvidersAPIClient is a client that implements the
// ListWebsiteAuthorizationProviders operation.
type ListWebsiteAuthorizationProvidersAPIClient interface {
	ListWebsiteAuthorizationProviders(context.Context, *ListWebsiteAuthorizationProvidersInput, ...func(*Options)) (*ListWebsiteAuthorizationProvidersOutput, error)
}

var _ ListWebsiteAuthorizationProvidersAPIClient = (*Client)(nil)

// ListWebsiteAuthorizationProvidersPaginatorOptions is the paginator options for
// ListWebsiteAuthorizationProviders
type ListWebsiteAuthorizationProvidersPaginatorOptions struct {
	// The maximum number of results to be included in the next page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListWebsiteAuthorizationProvidersPaginator is a paginator for
// ListWebsiteAuthorizationProviders
type ListWebsiteAuthorizationProvidersPaginator struct {
	options   ListWebsiteAuthorizationProvidersPaginatorOptions
	client    ListWebsiteAuthorizationProvidersAPIClient
	params    *ListWebsiteAuthorizationProvidersInput
	nextToken *string
	firstPage bool
}

// NewListWebsiteAuthorizationProvidersPaginator returns a new
// ListWebsiteAuthorizationProvidersPaginator
func NewListWebsiteAuthorizationProvidersPaginator(client ListWebsiteAuthorizationProvidersAPIClient, params *ListWebsiteAuthorizationProvidersInput, optFns ...func(*ListWebsiteAuthorizationProvidersPaginatorOptions)) *ListWebsiteAuthorizationProvidersPaginator {
	if params == nil {
		params = &ListWebsiteAuthorizationProvidersInput{}
	}

	options := ListWebsiteAuthorizationProvidersPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListWebsiteAuthorizationProvidersPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListWebsiteAuthorizationProvidersPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListWebsiteAuthorizationProviders page.
func (p *ListWebsiteAuthorizationProvidersPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListWebsiteAuthorizationProvidersOutput, error) {
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

	result, err := p.client.ListWebsiteAuthorizationProviders(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListWebsiteAuthorizationProviders(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListWebsiteAuthorizationProviders",
	}
}
