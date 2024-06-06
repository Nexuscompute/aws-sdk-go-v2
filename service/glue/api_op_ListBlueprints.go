// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all the blueprint names in an account.
func (c *Client) ListBlueprints(ctx context.Context, params *ListBlueprintsInput, optFns ...func(*Options)) (*ListBlueprintsOutput, error) {
	if params == nil {
		params = &ListBlueprintsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListBlueprints", params, optFns, c.addOperationListBlueprintsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListBlueprintsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListBlueprintsInput struct {

	// The maximum size of a list to return.
	MaxResults *int32

	// A continuation token, if this is a continuation request.
	NextToken *string

	// Filters the list by an Amazon Web Services resource tag.
	Tags map[string]string

	noSmithyDocumentSerde
}

type ListBlueprintsOutput struct {

	// List of names of blueprints in the account.
	Blueprints []string

	// A continuation token, if not all blueprint names have been returned.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListBlueprintsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListBlueprints{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListBlueprints{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListBlueprints"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListBlueprints(options.Region), middleware.Before); err != nil {
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

// ListBlueprintsAPIClient is a client that implements the ListBlueprints
// operation.
type ListBlueprintsAPIClient interface {
	ListBlueprints(context.Context, *ListBlueprintsInput, ...func(*Options)) (*ListBlueprintsOutput, error)
}

var _ ListBlueprintsAPIClient = (*Client)(nil)

// ListBlueprintsPaginatorOptions is the paginator options for ListBlueprints
type ListBlueprintsPaginatorOptions struct {
	// The maximum size of a list to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListBlueprintsPaginator is a paginator for ListBlueprints
type ListBlueprintsPaginator struct {
	options   ListBlueprintsPaginatorOptions
	client    ListBlueprintsAPIClient
	params    *ListBlueprintsInput
	nextToken *string
	firstPage bool
}

// NewListBlueprintsPaginator returns a new ListBlueprintsPaginator
func NewListBlueprintsPaginator(client ListBlueprintsAPIClient, params *ListBlueprintsInput, optFns ...func(*ListBlueprintsPaginatorOptions)) *ListBlueprintsPaginator {
	if params == nil {
		params = &ListBlueprintsInput{}
	}

	options := ListBlueprintsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListBlueprintsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListBlueprintsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListBlueprints page.
func (p *ListBlueprintsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListBlueprintsOutput, error) {
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

	result, err := p.client.ListBlueprints(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListBlueprints(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListBlueprints",
	}
}
