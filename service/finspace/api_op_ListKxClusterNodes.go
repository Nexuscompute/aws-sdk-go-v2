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

// Lists all the nodes in a kdb cluster.
func (c *Client) ListKxClusterNodes(ctx context.Context, params *ListKxClusterNodesInput, optFns ...func(*Options)) (*ListKxClusterNodesOutput, error) {
	if params == nil {
		params = &ListKxClusterNodesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListKxClusterNodes", params, optFns, c.addOperationListKxClusterNodesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListKxClusterNodesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListKxClusterNodesInput struct {

	// A unique name for the cluster.
	//
	// This member is required.
	ClusterName *string

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

type ListKxClusterNodesOutput struct {

	// A token that indicates where a results page should begin.
	NextToken *string

	// A list of nodes associated with the cluster.
	Nodes []types.KxNode

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListKxClusterNodesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListKxClusterNodes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListKxClusterNodes{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListKxClusterNodes"); err != nil {
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
	if err = addOpListKxClusterNodesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListKxClusterNodes(options.Region), middleware.Before); err != nil {
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

// ListKxClusterNodesAPIClient is a client that implements the ListKxClusterNodes
// operation.
type ListKxClusterNodesAPIClient interface {
	ListKxClusterNodes(context.Context, *ListKxClusterNodesInput, ...func(*Options)) (*ListKxClusterNodesOutput, error)
}

var _ ListKxClusterNodesAPIClient = (*Client)(nil)

// ListKxClusterNodesPaginatorOptions is the paginator options for
// ListKxClusterNodes
type ListKxClusterNodesPaginatorOptions struct {
	// The maximum number of results to return in this request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListKxClusterNodesPaginator is a paginator for ListKxClusterNodes
type ListKxClusterNodesPaginator struct {
	options   ListKxClusterNodesPaginatorOptions
	client    ListKxClusterNodesAPIClient
	params    *ListKxClusterNodesInput
	nextToken *string
	firstPage bool
}

// NewListKxClusterNodesPaginator returns a new ListKxClusterNodesPaginator
func NewListKxClusterNodesPaginator(client ListKxClusterNodesAPIClient, params *ListKxClusterNodesInput, optFns ...func(*ListKxClusterNodesPaginatorOptions)) *ListKxClusterNodesPaginator {
	if params == nil {
		params = &ListKxClusterNodesInput{}
	}

	options := ListKxClusterNodesPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListKxClusterNodesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListKxClusterNodesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListKxClusterNodes page.
func (p *ListKxClusterNodesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListKxClusterNodesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListKxClusterNodes(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListKxClusterNodes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListKxClusterNodes",
	}
}
