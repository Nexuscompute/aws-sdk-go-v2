// Code generated by smithy-go-codegen DO NOT EDIT.

package ssmsap

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/ssmsap/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of operations events.
//
// Available parameters include OperationID , as well as optional parameters
// MaxResults , NextToken , and Filters .
func (c *Client) ListOperationEvents(ctx context.Context, params *ListOperationEventsInput, optFns ...func(*Options)) (*ListOperationEventsOutput, error) {
	if params == nil {
		params = &ListOperationEventsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListOperationEvents", params, optFns, c.addOperationListOperationEventsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListOperationEventsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListOperationEventsInput struct {

	// The ID of the operation.
	//
	// This member is required.
	OperationId *string

	// Optionally specify filters to narrow the returned operation event items.
	//
	// Valid filter names include status , resourceID , and resourceType . The valid
	// operator for all three filters is Equals .
	Filters []types.Filter

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	//
	// If you do not specify a value for MaxResults , the request returns 50 items per
	// page by default.
	MaxResults *int32

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	noSmithyDocumentSerde
}

type ListOperationEventsOutput struct {

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// A returned list of operation events that meet the filter criteria.
	OperationEvents []types.OperationEvent

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListOperationEventsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListOperationEvents{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListOperationEvents{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListOperationEvents"); err != nil {
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
	if err = addOpListOperationEventsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListOperationEvents(options.Region), middleware.Before); err != nil {
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

// ListOperationEventsAPIClient is a client that implements the
// ListOperationEvents operation.
type ListOperationEventsAPIClient interface {
	ListOperationEvents(context.Context, *ListOperationEventsInput, ...func(*Options)) (*ListOperationEventsOutput, error)
}

var _ ListOperationEventsAPIClient = (*Client)(nil)

// ListOperationEventsPaginatorOptions is the paginator options for
// ListOperationEvents
type ListOperationEventsPaginatorOptions struct {
	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	//
	// If you do not specify a value for MaxResults , the request returns 50 items per
	// page by default.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListOperationEventsPaginator is a paginator for ListOperationEvents
type ListOperationEventsPaginator struct {
	options   ListOperationEventsPaginatorOptions
	client    ListOperationEventsAPIClient
	params    *ListOperationEventsInput
	nextToken *string
	firstPage bool
}

// NewListOperationEventsPaginator returns a new ListOperationEventsPaginator
func NewListOperationEventsPaginator(client ListOperationEventsAPIClient, params *ListOperationEventsInput, optFns ...func(*ListOperationEventsPaginatorOptions)) *ListOperationEventsPaginator {
	if params == nil {
		params = &ListOperationEventsInput{}
	}

	options := ListOperationEventsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListOperationEventsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListOperationEventsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListOperationEvents page.
func (p *ListOperationEventsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListOperationEventsOutput, error) {
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

	result, err := p.client.ListOperationEvents(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListOperationEvents(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListOperationEvents",
	}
}
