// Code generated by smithy-go-codegen DO NOT EDIT.

package cleanroomsml

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/cleanroomsml/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of ML input channels.
func (c *Client) ListMLInputChannels(ctx context.Context, params *ListMLInputChannelsInput, optFns ...func(*Options)) (*ListMLInputChannelsOutput, error) {
	if params == nil {
		params = &ListMLInputChannelsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListMLInputChannels", params, optFns, c.addOperationListMLInputChannelsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListMLInputChannelsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListMLInputChannelsInput struct {

	// The membership ID of the membership that contains the ML input channels that
	// you want to list.
	//
	// This member is required.
	MembershipIdentifier *string

	// The maximum number of ML input channels to return.
	MaxResults *int32

	// The token value retrieved from a previous call to access the next page of
	// results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListMLInputChannelsOutput struct {

	// The list of ML input channels that you wanted.
	//
	// This member is required.
	MlInputChannelsList []types.MLInputChannelSummary

	// The token value used to access the next page of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListMLInputChannelsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListMLInputChannels{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListMLInputChannels{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListMLInputChannels"); err != nil {
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
	if err = addOpListMLInputChannelsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListMLInputChannels(options.Region), middleware.Before); err != nil {
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

// ListMLInputChannelsPaginatorOptions is the paginator options for
// ListMLInputChannels
type ListMLInputChannelsPaginatorOptions struct {
	// The maximum number of ML input channels to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListMLInputChannelsPaginator is a paginator for ListMLInputChannels
type ListMLInputChannelsPaginator struct {
	options   ListMLInputChannelsPaginatorOptions
	client    ListMLInputChannelsAPIClient
	params    *ListMLInputChannelsInput
	nextToken *string
	firstPage bool
}

// NewListMLInputChannelsPaginator returns a new ListMLInputChannelsPaginator
func NewListMLInputChannelsPaginator(client ListMLInputChannelsAPIClient, params *ListMLInputChannelsInput, optFns ...func(*ListMLInputChannelsPaginatorOptions)) *ListMLInputChannelsPaginator {
	if params == nil {
		params = &ListMLInputChannelsInput{}
	}

	options := ListMLInputChannelsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListMLInputChannelsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListMLInputChannelsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListMLInputChannels page.
func (p *ListMLInputChannelsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListMLInputChannelsOutput, error) {
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
	result, err := p.client.ListMLInputChannels(ctx, &params, optFns...)
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

// ListMLInputChannelsAPIClient is a client that implements the
// ListMLInputChannels operation.
type ListMLInputChannelsAPIClient interface {
	ListMLInputChannels(context.Context, *ListMLInputChannelsInput, ...func(*Options)) (*ListMLInputChannelsOutput, error)
}

var _ ListMLInputChannelsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListMLInputChannels(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListMLInputChannels",
	}
}
