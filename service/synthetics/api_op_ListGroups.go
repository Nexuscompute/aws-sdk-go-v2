// Code generated by smithy-go-codegen DO NOT EDIT.

package synthetics

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/synthetics/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of all groups in the account, displaying their names, unique
// IDs, and ARNs. The groups from all Regions are returned.
func (c *Client) ListGroups(ctx context.Context, params *ListGroupsInput, optFns ...func(*Options)) (*ListGroupsOutput, error) {
	if params == nil {
		params = &ListGroupsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListGroups", params, optFns, c.addOperationListGroupsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListGroupsInput struct {

	// Specify this parameter to limit how many groups are returned each time you use
	// the ListGroups operation. If you omit this parameter, the default of 20 is used.
	MaxResults *int32

	// A token that indicates that there is more data available. You can use this
	// token in a subsequent operation to retrieve the next set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListGroupsOutput struct {

	// An array of structures that each contain information about one group.
	Groups []types.GroupSummary

	// A token that indicates that there is more data available. You can use this
	// token in a subsequent ListGroups operation to retrieve the next set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListGroupsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListGroups{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListGroups{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListGroups"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListGroups(options.Region), middleware.Before); err != nil {
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

// ListGroupsAPIClient is a client that implements the ListGroups operation.
type ListGroupsAPIClient interface {
	ListGroups(context.Context, *ListGroupsInput, ...func(*Options)) (*ListGroupsOutput, error)
}

var _ ListGroupsAPIClient = (*Client)(nil)

// ListGroupsPaginatorOptions is the paginator options for ListGroups
type ListGroupsPaginatorOptions struct {
	// Specify this parameter to limit how many groups are returned each time you use
	// the ListGroups operation. If you omit this parameter, the default of 20 is used.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListGroupsPaginator is a paginator for ListGroups
type ListGroupsPaginator struct {
	options   ListGroupsPaginatorOptions
	client    ListGroupsAPIClient
	params    *ListGroupsInput
	nextToken *string
	firstPage bool
}

// NewListGroupsPaginator returns a new ListGroupsPaginator
func NewListGroupsPaginator(client ListGroupsAPIClient, params *ListGroupsInput, optFns ...func(*ListGroupsPaginatorOptions)) *ListGroupsPaginator {
	if params == nil {
		params = &ListGroupsInput{}
	}

	options := ListGroupsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListGroupsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListGroupsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListGroups page.
func (p *ListGroupsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListGroupsOutput, error) {
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

	result, err := p.client.ListGroups(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListGroups(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListGroups",
	}
}
