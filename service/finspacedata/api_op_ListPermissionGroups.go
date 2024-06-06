// Code generated by smithy-go-codegen DO NOT EDIT.

package finspacedata

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/finspacedata/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all available permission groups in FinSpace.
//
// Deprecated: This method will be discontinued.
func (c *Client) ListPermissionGroups(ctx context.Context, params *ListPermissionGroupsInput, optFns ...func(*Options)) (*ListPermissionGroupsOutput, error) {
	if params == nil {
		params = &ListPermissionGroupsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListPermissionGroups", params, optFns, c.addOperationListPermissionGroupsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListPermissionGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPermissionGroupsInput struct {

	// The maximum number of results per page.
	//
	// This member is required.
	MaxResults *int32

	// A token that indicates where a results page should begin.
	NextToken *string

	noSmithyDocumentSerde
}

type ListPermissionGroupsOutput struct {

	// A token that indicates where a results page should begin.
	NextToken *string

	// A list of all the permission groups.
	PermissionGroups []types.PermissionGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListPermissionGroupsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListPermissionGroups{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListPermissionGroups{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListPermissionGroups"); err != nil {
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
	if err = addOpListPermissionGroupsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListPermissionGroups(options.Region), middleware.Before); err != nil {
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

// ListPermissionGroupsAPIClient is a client that implements the
// ListPermissionGroups operation.
type ListPermissionGroupsAPIClient interface {
	ListPermissionGroups(context.Context, *ListPermissionGroupsInput, ...func(*Options)) (*ListPermissionGroupsOutput, error)
}

var _ ListPermissionGroupsAPIClient = (*Client)(nil)

// ListPermissionGroupsPaginatorOptions is the paginator options for
// ListPermissionGroups
type ListPermissionGroupsPaginatorOptions struct {
	// The maximum number of results per page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListPermissionGroupsPaginator is a paginator for ListPermissionGroups
type ListPermissionGroupsPaginator struct {
	options   ListPermissionGroupsPaginatorOptions
	client    ListPermissionGroupsAPIClient
	params    *ListPermissionGroupsInput
	nextToken *string
	firstPage bool
}

// NewListPermissionGroupsPaginator returns a new ListPermissionGroupsPaginator
func NewListPermissionGroupsPaginator(client ListPermissionGroupsAPIClient, params *ListPermissionGroupsInput, optFns ...func(*ListPermissionGroupsPaginatorOptions)) *ListPermissionGroupsPaginator {
	if params == nil {
		params = &ListPermissionGroupsInput{}
	}

	options := ListPermissionGroupsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListPermissionGroupsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListPermissionGroupsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListPermissionGroups page.
func (p *ListPermissionGroupsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListPermissionGroupsOutput, error) {
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

	result, err := p.client.ListPermissionGroups(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListPermissionGroups(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListPermissionGroups",
	}
}
