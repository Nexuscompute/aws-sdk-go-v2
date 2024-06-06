// Code generated by smithy-go-codegen DO NOT EDIT.

package ivschat

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/ivschat/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets summary information about all your rooms in the AWS region where the API
// request is processed. Results are sorted in descending order of updateTime .
func (c *Client) ListRooms(ctx context.Context, params *ListRoomsInput, optFns ...func(*Options)) (*ListRoomsOutput, error) {
	if params == nil {
		params = &ListRoomsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListRooms", params, optFns, c.addOperationListRoomsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListRoomsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListRoomsInput struct {

	// Logging-configuration identifier.
	LoggingConfigurationIdentifier *string

	// Maximum number of rooms to return. Default: 50.
	MaxResults *int32

	// Filters the list to match the specified message review handler URI.
	MessageReviewHandlerUri *string

	// Filters the list to match the specified room name.
	Name *string

	// The first room to retrieve. This is used for pagination; see the nextToken
	// response field.
	NextToken *string

	noSmithyDocumentSerde
}

type ListRoomsOutput struct {

	// List of the matching rooms (summary information only).
	//
	// This member is required.
	Rooms []types.RoomSummary

	// If there are more rooms than maxResults , use nextToken in the request to get
	// the next set.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListRoomsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListRooms{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListRooms{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListRooms"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListRooms(options.Region), middleware.Before); err != nil {
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

// ListRoomsAPIClient is a client that implements the ListRooms operation.
type ListRoomsAPIClient interface {
	ListRooms(context.Context, *ListRoomsInput, ...func(*Options)) (*ListRoomsOutput, error)
}

var _ ListRoomsAPIClient = (*Client)(nil)

// ListRoomsPaginatorOptions is the paginator options for ListRooms
type ListRoomsPaginatorOptions struct {
	// Maximum number of rooms to return. Default: 50.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListRoomsPaginator is a paginator for ListRooms
type ListRoomsPaginator struct {
	options   ListRoomsPaginatorOptions
	client    ListRoomsAPIClient
	params    *ListRoomsInput
	nextToken *string
	firstPage bool
}

// NewListRoomsPaginator returns a new ListRoomsPaginator
func NewListRoomsPaginator(client ListRoomsAPIClient, params *ListRoomsInput, optFns ...func(*ListRoomsPaginatorOptions)) *ListRoomsPaginator {
	if params == nil {
		params = &ListRoomsInput{}
	}

	options := ListRoomsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListRoomsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListRoomsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListRooms page.
func (p *ListRoomsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListRoomsOutput, error) {
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

	result, err := p.client.ListRooms(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListRooms(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListRooms",
	}
}
