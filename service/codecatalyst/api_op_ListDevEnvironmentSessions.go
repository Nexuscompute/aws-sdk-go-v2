// Code generated by smithy-go-codegen DO NOT EDIT.

package codecatalyst

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/codecatalyst/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a list of active sessions for a Dev Environment in a project.
func (c *Client) ListDevEnvironmentSessions(ctx context.Context, params *ListDevEnvironmentSessionsInput, optFns ...func(*Options)) (*ListDevEnvironmentSessionsOutput, error) {
	if params == nil {
		params = &ListDevEnvironmentSessionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDevEnvironmentSessions", params, optFns, c.addOperationListDevEnvironmentSessionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDevEnvironmentSessionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDevEnvironmentSessionsInput struct {

	// The system-generated unique ID of the Dev Environment.
	//
	// This member is required.
	DevEnvironmentId *string

	// The name of the project in the space.
	//
	// This member is required.
	ProjectName *string

	// The name of the space.
	//
	// This member is required.
	SpaceName *string

	// The maximum number of results to show in a single call to this API. If the
	// number of results is larger than the number you specified, the response will
	// include a NextToken element, which you can use to obtain additional results.
	MaxResults *int32

	// A token returned from a call to this API to indicate the next batch of results
	// to return, if any.
	NextToken *string

	noSmithyDocumentSerde
}

type ListDevEnvironmentSessionsOutput struct {

	// Information about each session retrieved in the list.
	//
	// This member is required.
	Items []types.DevEnvironmentSessionSummary

	// A token returned from a call to this API to indicate the next batch of results
	// to return, if any.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListDevEnvironmentSessionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListDevEnvironmentSessions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListDevEnvironmentSessions{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListDevEnvironmentSessions"); err != nil {
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
	if err = addOpListDevEnvironmentSessionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListDevEnvironmentSessions(options.Region), middleware.Before); err != nil {
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

// ListDevEnvironmentSessionsAPIClient is a client that implements the
// ListDevEnvironmentSessions operation.
type ListDevEnvironmentSessionsAPIClient interface {
	ListDevEnvironmentSessions(context.Context, *ListDevEnvironmentSessionsInput, ...func(*Options)) (*ListDevEnvironmentSessionsOutput, error)
}

var _ ListDevEnvironmentSessionsAPIClient = (*Client)(nil)

// ListDevEnvironmentSessionsPaginatorOptions is the paginator options for
// ListDevEnvironmentSessions
type ListDevEnvironmentSessionsPaginatorOptions struct {
	// The maximum number of results to show in a single call to this API. If the
	// number of results is larger than the number you specified, the response will
	// include a NextToken element, which you can use to obtain additional results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListDevEnvironmentSessionsPaginator is a paginator for
// ListDevEnvironmentSessions
type ListDevEnvironmentSessionsPaginator struct {
	options   ListDevEnvironmentSessionsPaginatorOptions
	client    ListDevEnvironmentSessionsAPIClient
	params    *ListDevEnvironmentSessionsInput
	nextToken *string
	firstPage bool
}

// NewListDevEnvironmentSessionsPaginator returns a new
// ListDevEnvironmentSessionsPaginator
func NewListDevEnvironmentSessionsPaginator(client ListDevEnvironmentSessionsAPIClient, params *ListDevEnvironmentSessionsInput, optFns ...func(*ListDevEnvironmentSessionsPaginatorOptions)) *ListDevEnvironmentSessionsPaginator {
	if params == nil {
		params = &ListDevEnvironmentSessionsInput{}
	}

	options := ListDevEnvironmentSessionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListDevEnvironmentSessionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListDevEnvironmentSessionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListDevEnvironmentSessions page.
func (p *ListDevEnvironmentSessionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListDevEnvironmentSessionsOutput, error) {
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

	result, err := p.client.ListDevEnvironmentSessions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListDevEnvironmentSessions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListDevEnvironmentSessions",
	}
}
