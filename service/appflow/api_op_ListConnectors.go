// Code generated by smithy-go-codegen DO NOT EDIT.

package appflow

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/appflow/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the list of all registered custom connectors in your Amazon Web
// Services account. This API lists only custom connectors registered in this
// account, not the Amazon Web Services authored connectors.
func (c *Client) ListConnectors(ctx context.Context, params *ListConnectorsInput, optFns ...func(*Options)) (*ListConnectorsOutput, error) {
	if params == nil {
		params = &ListConnectorsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListConnectors", params, optFns, c.addOperationListConnectorsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListConnectorsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListConnectorsInput struct {

	// Specifies the maximum number of items that should be returned in the result
	// set. The default for maxResults is 20 (for all paginated API operations).
	MaxResults *int32

	// The pagination token for the next page of data.
	NextToken *string

	noSmithyDocumentSerde
}

type ListConnectorsOutput struct {

	// Contains information about the connectors supported by Amazon AppFlow.
	Connectors []types.ConnectorDetail

	// The pagination token for the next page of data. If nextToken=null, this means
	// that all records have been fetched.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListConnectorsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListConnectors{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListConnectors{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListConnectors"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListConnectors(options.Region), middleware.Before); err != nil {
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

// ListConnectorsAPIClient is a client that implements the ListConnectors
// operation.
type ListConnectorsAPIClient interface {
	ListConnectors(context.Context, *ListConnectorsInput, ...func(*Options)) (*ListConnectorsOutput, error)
}

var _ ListConnectorsAPIClient = (*Client)(nil)

// ListConnectorsPaginatorOptions is the paginator options for ListConnectors
type ListConnectorsPaginatorOptions struct {
	// Specifies the maximum number of items that should be returned in the result
	// set. The default for maxResults is 20 (for all paginated API operations).
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListConnectorsPaginator is a paginator for ListConnectors
type ListConnectorsPaginator struct {
	options   ListConnectorsPaginatorOptions
	client    ListConnectorsAPIClient
	params    *ListConnectorsInput
	nextToken *string
	firstPage bool
}

// NewListConnectorsPaginator returns a new ListConnectorsPaginator
func NewListConnectorsPaginator(client ListConnectorsAPIClient, params *ListConnectorsInput, optFns ...func(*ListConnectorsPaginatorOptions)) *ListConnectorsPaginator {
	if params == nil {
		params = &ListConnectorsInput{}
	}

	options := ListConnectorsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListConnectorsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListConnectorsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListConnectors page.
func (p *ListConnectorsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListConnectorsOutput, error) {
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

	result, err := p.client.ListConnectors(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListConnectors(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListConnectors",
	}
}
