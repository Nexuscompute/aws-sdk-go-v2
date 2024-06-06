// Code generated by smithy-go-codegen DO NOT EDIT.

package mgn

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/mgn/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves all SourceServers or multiple SourceServers by ID.
func (c *Client) DescribeSourceServers(ctx context.Context, params *DescribeSourceServersInput, optFns ...func(*Options)) (*DescribeSourceServersOutput, error) {
	if params == nil {
		params = &DescribeSourceServersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeSourceServers", params, optFns, c.addOperationDescribeSourceServersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeSourceServersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeSourceServersInput struct {

	// Request to filter Source Servers list by Accoun ID.
	AccountID *string

	// Request to filter Source Servers list.
	Filters *types.DescribeSourceServersRequestFilters

	// Request to filter Source Servers list by maximum results.
	MaxResults *int32

	// Request to filter Source Servers list by next token.
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeSourceServersOutput struct {

	// Request to filter Source Servers list by item.
	Items []types.SourceServer

	// Request to filter Source Servers next token.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeSourceServersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeSourceServers{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeSourceServers{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeSourceServers"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeSourceServers(options.Region), middleware.Before); err != nil {
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

// DescribeSourceServersAPIClient is a client that implements the
// DescribeSourceServers operation.
type DescribeSourceServersAPIClient interface {
	DescribeSourceServers(context.Context, *DescribeSourceServersInput, ...func(*Options)) (*DescribeSourceServersOutput, error)
}

var _ DescribeSourceServersAPIClient = (*Client)(nil)

// DescribeSourceServersPaginatorOptions is the paginator options for
// DescribeSourceServers
type DescribeSourceServersPaginatorOptions struct {
	// Request to filter Source Servers list by maximum results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeSourceServersPaginator is a paginator for DescribeSourceServers
type DescribeSourceServersPaginator struct {
	options   DescribeSourceServersPaginatorOptions
	client    DescribeSourceServersAPIClient
	params    *DescribeSourceServersInput
	nextToken *string
	firstPage bool
}

// NewDescribeSourceServersPaginator returns a new DescribeSourceServersPaginator
func NewDescribeSourceServersPaginator(client DescribeSourceServersAPIClient, params *DescribeSourceServersInput, optFns ...func(*DescribeSourceServersPaginatorOptions)) *DescribeSourceServersPaginator {
	if params == nil {
		params = &DescribeSourceServersInput{}
	}

	options := DescribeSourceServersPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeSourceServersPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeSourceServersPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeSourceServers page.
func (p *DescribeSourceServersPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeSourceServersOutput, error) {
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

	result, err := p.client.DescribeSourceServers(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeSourceServers(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeSourceServers",
	}
}
