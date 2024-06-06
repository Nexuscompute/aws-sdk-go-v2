// Code generated by smithy-go-codegen DO NOT EDIT.

package servicecatalogappregistry

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all attribute groups that are associated with specified application.
// Results are paginated.
func (c *Client) ListAssociatedAttributeGroups(ctx context.Context, params *ListAssociatedAttributeGroupsInput, optFns ...func(*Options)) (*ListAssociatedAttributeGroupsOutput, error) {
	if params == nil {
		params = &ListAssociatedAttributeGroupsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAssociatedAttributeGroups", params, optFns, c.addOperationListAssociatedAttributeGroupsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAssociatedAttributeGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAssociatedAttributeGroupsInput struct {

	// The name or ID of the application.
	//
	// This member is required.
	Application *string

	// The upper bound of the number of results to return (cannot exceed 25). If this
	// parameter is omitted, it defaults to 25. This value is optional.
	MaxResults *int32

	// The token to use to get the next page of results after a previous API call.
	NextToken *string

	noSmithyDocumentSerde
}

type ListAssociatedAttributeGroupsOutput struct {

	// A list of attribute group IDs.
	AttributeGroups []string

	// The token to use to get the next page of results after a previous API call.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAssociatedAttributeGroupsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListAssociatedAttributeGroups{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListAssociatedAttributeGroups{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListAssociatedAttributeGroups"); err != nil {
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
	if err = addOpListAssociatedAttributeGroupsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAssociatedAttributeGroups(options.Region), middleware.Before); err != nil {
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

// ListAssociatedAttributeGroupsAPIClient is a client that implements the
// ListAssociatedAttributeGroups operation.
type ListAssociatedAttributeGroupsAPIClient interface {
	ListAssociatedAttributeGroups(context.Context, *ListAssociatedAttributeGroupsInput, ...func(*Options)) (*ListAssociatedAttributeGroupsOutput, error)
}

var _ ListAssociatedAttributeGroupsAPIClient = (*Client)(nil)

// ListAssociatedAttributeGroupsPaginatorOptions is the paginator options for
// ListAssociatedAttributeGroups
type ListAssociatedAttributeGroupsPaginatorOptions struct {
	// The upper bound of the number of results to return (cannot exceed 25). If this
	// parameter is omitted, it defaults to 25. This value is optional.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAssociatedAttributeGroupsPaginator is a paginator for
// ListAssociatedAttributeGroups
type ListAssociatedAttributeGroupsPaginator struct {
	options   ListAssociatedAttributeGroupsPaginatorOptions
	client    ListAssociatedAttributeGroupsAPIClient
	params    *ListAssociatedAttributeGroupsInput
	nextToken *string
	firstPage bool
}

// NewListAssociatedAttributeGroupsPaginator returns a new
// ListAssociatedAttributeGroupsPaginator
func NewListAssociatedAttributeGroupsPaginator(client ListAssociatedAttributeGroupsAPIClient, params *ListAssociatedAttributeGroupsInput, optFns ...func(*ListAssociatedAttributeGroupsPaginatorOptions)) *ListAssociatedAttributeGroupsPaginator {
	if params == nil {
		params = &ListAssociatedAttributeGroupsInput{}
	}

	options := ListAssociatedAttributeGroupsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAssociatedAttributeGroupsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAssociatedAttributeGroupsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListAssociatedAttributeGroups page.
func (p *ListAssociatedAttributeGroupsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAssociatedAttributeGroupsOutput, error) {
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

	result, err := p.client.ListAssociatedAttributeGroups(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListAssociatedAttributeGroups(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListAssociatedAttributeGroups",
	}
}
