// Code generated by smithy-go-codegen DO NOT EDIT.

package shield

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/shield/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves ProtectionGroup objects for the account. You can retrieve all protection groups or
// you can provide filtering criteria and retrieve just the subset of protection
// groups that match the criteria.
func (c *Client) ListProtectionGroups(ctx context.Context, params *ListProtectionGroupsInput, optFns ...func(*Options)) (*ListProtectionGroupsOutput, error) {
	if params == nil {
		params = &ListProtectionGroupsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListProtectionGroups", params, optFns, c.addOperationListProtectionGroupsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListProtectionGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListProtectionGroupsInput struct {

	// Narrows the set of protection groups that the call retrieves. You can retrieve
	// a single protection group by its name and you can retrieve all protection groups
	// that are configured with specific pattern or aggregation settings. You can
	// provide up to one criteria per filter type. Shield Advanced returns the
	// protection groups that exactly match all of the search criteria that you
	// provide.
	InclusionFilters *types.InclusionProtectionGroupFilters

	// The greatest number of objects that you want Shield Advanced to return to the
	// list request. Shield Advanced might return fewer objects than you indicate in
	// this setting, even if more objects are available. If there are more objects
	// remaining, Shield Advanced will always also return a NextToken value in the
	// response.
	//
	// The default setting is 20.
	MaxResults *int32

	// When you request a list of objects from Shield Advanced, if the response does
	// not include all of the remaining available objects, Shield Advanced includes a
	// NextToken value in the response. You can retrieve the next batch of objects by
	// requesting the list again and providing the token that was returned by the prior
	// call in your request.
	//
	// You can indicate the maximum number of objects that you want Shield Advanced to
	// return for a single call with the MaxResults setting. Shield Advanced will not
	// return more than MaxResults objects, but may return fewer, even if more objects
	// are still available.
	//
	// Whenever more objects remain that Shield Advanced has not yet returned to you,
	// the response will include a NextToken value.
	//
	// On your first call to a list operation, leave this setting empty.
	NextToken *string

	noSmithyDocumentSerde
}

type ListProtectionGroupsOutput struct {

	//
	//
	// This member is required.
	ProtectionGroups []types.ProtectionGroup

	// When you request a list of objects from Shield Advanced, if the response does
	// not include all of the remaining available objects, Shield Advanced includes a
	// NextToken value in the response. You can retrieve the next batch of objects by
	// requesting the list again and providing the token that was returned by the prior
	// call in your request.
	//
	// You can indicate the maximum number of objects that you want Shield Advanced to
	// return for a single call with the MaxResults setting. Shield Advanced will not
	// return more than MaxResults objects, but may return fewer, even if more objects
	// are still available.
	//
	// Whenever more objects remain that Shield Advanced has not yet returned to you,
	// the response will include a NextToken value.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListProtectionGroupsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListProtectionGroups{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListProtectionGroups{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListProtectionGroups"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListProtectionGroups(options.Region), middleware.Before); err != nil {
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

// ListProtectionGroupsAPIClient is a client that implements the
// ListProtectionGroups operation.
type ListProtectionGroupsAPIClient interface {
	ListProtectionGroups(context.Context, *ListProtectionGroupsInput, ...func(*Options)) (*ListProtectionGroupsOutput, error)
}

var _ ListProtectionGroupsAPIClient = (*Client)(nil)

// ListProtectionGroupsPaginatorOptions is the paginator options for
// ListProtectionGroups
type ListProtectionGroupsPaginatorOptions struct {
	// The greatest number of objects that you want Shield Advanced to return to the
	// list request. Shield Advanced might return fewer objects than you indicate in
	// this setting, even if more objects are available. If there are more objects
	// remaining, Shield Advanced will always also return a NextToken value in the
	// response.
	//
	// The default setting is 20.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListProtectionGroupsPaginator is a paginator for ListProtectionGroups
type ListProtectionGroupsPaginator struct {
	options   ListProtectionGroupsPaginatorOptions
	client    ListProtectionGroupsAPIClient
	params    *ListProtectionGroupsInput
	nextToken *string
	firstPage bool
}

// NewListProtectionGroupsPaginator returns a new ListProtectionGroupsPaginator
func NewListProtectionGroupsPaginator(client ListProtectionGroupsAPIClient, params *ListProtectionGroupsInput, optFns ...func(*ListProtectionGroupsPaginatorOptions)) *ListProtectionGroupsPaginator {
	if params == nil {
		params = &ListProtectionGroupsInput{}
	}

	options := ListProtectionGroupsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListProtectionGroupsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListProtectionGroupsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListProtectionGroups page.
func (p *ListProtectionGroupsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListProtectionGroupsOutput, error) {
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

	result, err := p.client.ListProtectionGroups(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListProtectionGroups(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListProtectionGroups",
	}
}
