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

// Lists all ReplicationConfigurationTemplates, filtered by Source Server IDs.
func (c *Client) DescribeReplicationConfigurationTemplates(ctx context.Context, params *DescribeReplicationConfigurationTemplatesInput, optFns ...func(*Options)) (*DescribeReplicationConfigurationTemplatesOutput, error) {
	if params == nil {
		params = &DescribeReplicationConfigurationTemplatesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeReplicationConfigurationTemplates", params, optFns, c.addOperationDescribeReplicationConfigurationTemplatesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeReplicationConfigurationTemplatesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeReplicationConfigurationTemplatesInput struct {

	// Request to describe Replication Configuration template by max results.
	MaxResults *int32

	// Request to describe Replication Configuration template by next token.
	NextToken *string

	// Request to describe Replication Configuration template by template IDs.
	ReplicationConfigurationTemplateIDs []string

	noSmithyDocumentSerde
}

type DescribeReplicationConfigurationTemplatesOutput struct {

	// Request to describe Replication Configuration template by items.
	Items []types.ReplicationConfigurationTemplate

	// Request to describe Replication Configuration template by next token.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeReplicationConfigurationTemplatesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeReplicationConfigurationTemplates{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeReplicationConfigurationTemplates{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeReplicationConfigurationTemplates"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeReplicationConfigurationTemplates(options.Region), middleware.Before); err != nil {
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

// DescribeReplicationConfigurationTemplatesAPIClient is a client that implements
// the DescribeReplicationConfigurationTemplates operation.
type DescribeReplicationConfigurationTemplatesAPIClient interface {
	DescribeReplicationConfigurationTemplates(context.Context, *DescribeReplicationConfigurationTemplatesInput, ...func(*Options)) (*DescribeReplicationConfigurationTemplatesOutput, error)
}

var _ DescribeReplicationConfigurationTemplatesAPIClient = (*Client)(nil)

// DescribeReplicationConfigurationTemplatesPaginatorOptions is the paginator
// options for DescribeReplicationConfigurationTemplates
type DescribeReplicationConfigurationTemplatesPaginatorOptions struct {
	// Request to describe Replication Configuration template by max results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeReplicationConfigurationTemplatesPaginator is a paginator for
// DescribeReplicationConfigurationTemplates
type DescribeReplicationConfigurationTemplatesPaginator struct {
	options   DescribeReplicationConfigurationTemplatesPaginatorOptions
	client    DescribeReplicationConfigurationTemplatesAPIClient
	params    *DescribeReplicationConfigurationTemplatesInput
	nextToken *string
	firstPage bool
}

// NewDescribeReplicationConfigurationTemplatesPaginator returns a new
// DescribeReplicationConfigurationTemplatesPaginator
func NewDescribeReplicationConfigurationTemplatesPaginator(client DescribeReplicationConfigurationTemplatesAPIClient, params *DescribeReplicationConfigurationTemplatesInput, optFns ...func(*DescribeReplicationConfigurationTemplatesPaginatorOptions)) *DescribeReplicationConfigurationTemplatesPaginator {
	if params == nil {
		params = &DescribeReplicationConfigurationTemplatesInput{}
	}

	options := DescribeReplicationConfigurationTemplatesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeReplicationConfigurationTemplatesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeReplicationConfigurationTemplatesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeReplicationConfigurationTemplates page.
func (p *DescribeReplicationConfigurationTemplatesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeReplicationConfigurationTemplatesOutput, error) {
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

	result, err := p.client.DescribeReplicationConfigurationTemplates(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeReplicationConfigurationTemplates(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeReplicationConfigurationTemplates",
	}
}
