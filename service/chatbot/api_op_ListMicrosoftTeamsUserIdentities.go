// Code generated by smithy-go-codegen DO NOT EDIT.

package chatbot

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/chatbot/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all Microsoft Teams user identities with a mapped role.
func (c *Client) ListMicrosoftTeamsUserIdentities(ctx context.Context, params *ListMicrosoftTeamsUserIdentitiesInput, optFns ...func(*Options)) (*ListMicrosoftTeamsUserIdentitiesOutput, error) {
	if params == nil {
		params = &ListMicrosoftTeamsUserIdentitiesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListMicrosoftTeamsUserIdentities", params, optFns, c.addOperationListMicrosoftTeamsUserIdentitiesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListMicrosoftTeamsUserIdentitiesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListMicrosoftTeamsUserIdentitiesInput struct {

	// The ARN of the MicrosoftTeamsChannelConfiguration associated with the user
	// identities to list.
	ChatConfigurationArn *string

	// The maximum number of results to include in the response. If more results exist
	// than the specified MaxResults value, a token is included in the response so that
	// the remaining results can be retrieved.
	MaxResults *int32

	// An optional token returned from a prior request. Use this token for pagination
	// of results from this action. If this parameter is specified, the response
	// includes only results beyond the token, up to the value specified by MaxResults.
	NextToken *string

	noSmithyDocumentSerde
}

type ListMicrosoftTeamsUserIdentitiesOutput struct {

	// An optional token returned from a prior request. Use this token for pagination
	// of results from this action. If this parameter is specified, the response
	// includes only results beyond the token, up to the value specified by MaxResults.
	NextToken *string

	// User level permissions associated to a channel configuration.
	TeamsUserIdentities []types.TeamsUserIdentity

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListMicrosoftTeamsUserIdentitiesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListMicrosoftTeamsUserIdentities{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListMicrosoftTeamsUserIdentities{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListMicrosoftTeamsUserIdentities"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListMicrosoftTeamsUserIdentities(options.Region), middleware.Before); err != nil {
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

// ListMicrosoftTeamsUserIdentitiesAPIClient is a client that implements the
// ListMicrosoftTeamsUserIdentities operation.
type ListMicrosoftTeamsUserIdentitiesAPIClient interface {
	ListMicrosoftTeamsUserIdentities(context.Context, *ListMicrosoftTeamsUserIdentitiesInput, ...func(*Options)) (*ListMicrosoftTeamsUserIdentitiesOutput, error)
}

var _ ListMicrosoftTeamsUserIdentitiesAPIClient = (*Client)(nil)

// ListMicrosoftTeamsUserIdentitiesPaginatorOptions is the paginator options for
// ListMicrosoftTeamsUserIdentities
type ListMicrosoftTeamsUserIdentitiesPaginatorOptions struct {
	// The maximum number of results to include in the response. If more results exist
	// than the specified MaxResults value, a token is included in the response so that
	// the remaining results can be retrieved.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListMicrosoftTeamsUserIdentitiesPaginator is a paginator for
// ListMicrosoftTeamsUserIdentities
type ListMicrosoftTeamsUserIdentitiesPaginator struct {
	options   ListMicrosoftTeamsUserIdentitiesPaginatorOptions
	client    ListMicrosoftTeamsUserIdentitiesAPIClient
	params    *ListMicrosoftTeamsUserIdentitiesInput
	nextToken *string
	firstPage bool
}

// NewListMicrosoftTeamsUserIdentitiesPaginator returns a new
// ListMicrosoftTeamsUserIdentitiesPaginator
func NewListMicrosoftTeamsUserIdentitiesPaginator(client ListMicrosoftTeamsUserIdentitiesAPIClient, params *ListMicrosoftTeamsUserIdentitiesInput, optFns ...func(*ListMicrosoftTeamsUserIdentitiesPaginatorOptions)) *ListMicrosoftTeamsUserIdentitiesPaginator {
	if params == nil {
		params = &ListMicrosoftTeamsUserIdentitiesInput{}
	}

	options := ListMicrosoftTeamsUserIdentitiesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListMicrosoftTeamsUserIdentitiesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListMicrosoftTeamsUserIdentitiesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListMicrosoftTeamsUserIdentities page.
func (p *ListMicrosoftTeamsUserIdentitiesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListMicrosoftTeamsUserIdentitiesOutput, error) {
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

	result, err := p.client.ListMicrosoftTeamsUserIdentities(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListMicrosoftTeamsUserIdentities(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListMicrosoftTeamsUserIdentities",
	}
}
