// Code generated by smithy-go-codegen DO NOT EDIT.

package chimesdkmessaging

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/chimesdkmessaging/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all the SubChannels in an elastic channel when given a channel ID.
// Available only to the app instance admins and channel moderators of elastic
// channels.
func (c *Client) ListSubChannels(ctx context.Context, params *ListSubChannelsInput, optFns ...func(*Options)) (*ListSubChannelsOutput, error) {
	if params == nil {
		params = &ListSubChannelsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListSubChannels", params, optFns, c.addOperationListSubChannelsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListSubChannelsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSubChannelsInput struct {

	// The ARN of elastic channel.
	//
	// This member is required.
	ChannelArn *string

	// The AppInstanceUserArn of the user making the API call.
	//
	// This member is required.
	ChimeBearer *string

	// The maximum number of sub-channels that you want to return.
	MaxResults *int32

	// The token passed by previous API calls until all requested sub-channels are
	// returned.
	NextToken *string

	noSmithyDocumentSerde
}

type ListSubChannelsOutput struct {

	// The ARN of elastic channel.
	ChannelArn *string

	// The token passed by previous API calls until all requested sub-channels are
	// returned.
	NextToken *string

	// The information about each sub-channel.
	SubChannels []types.SubChannelSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListSubChannelsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListSubChannels{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListSubChannels{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListSubChannels"); err != nil {
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
	if err = addSpanRetryLoop(stack, options); err != nil {
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
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = addOpListSubChannelsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListSubChannels(options.Region), middleware.Before); err != nil {
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
	if err = addSpanInitializeStart(stack); err != nil {
		return err
	}
	if err = addSpanInitializeEnd(stack); err != nil {
		return err
	}
	if err = addSpanBuildRequestStart(stack); err != nil {
		return err
	}
	if err = addSpanBuildRequestEnd(stack); err != nil {
		return err
	}
	return nil
}

// ListSubChannelsPaginatorOptions is the paginator options for ListSubChannels
type ListSubChannelsPaginatorOptions struct {
	// The maximum number of sub-channels that you want to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListSubChannelsPaginator is a paginator for ListSubChannels
type ListSubChannelsPaginator struct {
	options   ListSubChannelsPaginatorOptions
	client    ListSubChannelsAPIClient
	params    *ListSubChannelsInput
	nextToken *string
	firstPage bool
}

// NewListSubChannelsPaginator returns a new ListSubChannelsPaginator
func NewListSubChannelsPaginator(client ListSubChannelsAPIClient, params *ListSubChannelsInput, optFns ...func(*ListSubChannelsPaginatorOptions)) *ListSubChannelsPaginator {
	if params == nil {
		params = &ListSubChannelsInput{}
	}

	options := ListSubChannelsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListSubChannelsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListSubChannelsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListSubChannels page.
func (p *ListSubChannelsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListSubChannelsOutput, error) {
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

	optFns = append([]func(*Options){
		addIsPaginatorUserAgent,
	}, optFns...)
	result, err := p.client.ListSubChannels(ctx, &params, optFns...)
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

// ListSubChannelsAPIClient is a client that implements the ListSubChannels
// operation.
type ListSubChannelsAPIClient interface {
	ListSubChannels(context.Context, *ListSubChannelsInput, ...func(*Options)) (*ListSubChannelsOutput, error)
}

var _ ListSubChannelsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListSubChannels(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListSubChannels",
	}
}
