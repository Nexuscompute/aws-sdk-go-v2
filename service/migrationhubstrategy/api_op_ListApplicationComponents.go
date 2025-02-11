// Code generated by smithy-go-codegen DO NOT EDIT.

package migrationhubstrategy

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/migrationhubstrategy/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a list of all the application components (processes).
func (c *Client) ListApplicationComponents(ctx context.Context, params *ListApplicationComponentsInput, optFns ...func(*Options)) (*ListApplicationComponentsOutput, error) {
	if params == nil {
		params = &ListApplicationComponentsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListApplicationComponents", params, optFns, c.addOperationListApplicationComponentsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListApplicationComponentsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListApplicationComponentsInput struct {

	//  Criteria for filtering the list of application components.
	ApplicationComponentCriteria types.ApplicationComponentCriteria

	//  Specify the value based on the application component criteria type. For
	// example, if applicationComponentCriteria is set to SERVER_ID and filterValue is
	// set to server1 , then ListApplicationComponents returns all the application components running on
	// server1.
	FilterValue *string

	//  The group ID specified in to filter on.
	GroupIdFilter []types.Group

	//  The maximum number of items to include in the response. The maximum value is
	// 100.
	MaxResults *int32

	//  The token from a previous call that you use to retrieve the next set of
	// results. For example, if a previous call to this action returned 100 items, but
	// you set maxResults to 10. You'll receive a set of 10 results along with a
	// token. You then use the returned token to retrieve the next set of 10.
	NextToken *string

	//  Specifies whether to sort by ascending ( ASC ) or descending ( DESC ) order.
	Sort types.SortOrder

	noSmithyDocumentSerde
}

type ListApplicationComponentsOutput struct {

	//  The list of application components with detailed information about each
	// component.
	ApplicationComponentInfos []types.ApplicationComponentDetail

	//  The token you use to retrieve the next set of results, or null if there are no
	// more results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListApplicationComponentsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListApplicationComponents{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListApplicationComponents{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListApplicationComponents"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListApplicationComponents(options.Region), middleware.Before); err != nil {
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

// ListApplicationComponentsPaginatorOptions is the paginator options for
// ListApplicationComponents
type ListApplicationComponentsPaginatorOptions struct {
	//  The maximum number of items to include in the response. The maximum value is
	// 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListApplicationComponentsPaginator is a paginator for ListApplicationComponents
type ListApplicationComponentsPaginator struct {
	options   ListApplicationComponentsPaginatorOptions
	client    ListApplicationComponentsAPIClient
	params    *ListApplicationComponentsInput
	nextToken *string
	firstPage bool
}

// NewListApplicationComponentsPaginator returns a new
// ListApplicationComponentsPaginator
func NewListApplicationComponentsPaginator(client ListApplicationComponentsAPIClient, params *ListApplicationComponentsInput, optFns ...func(*ListApplicationComponentsPaginatorOptions)) *ListApplicationComponentsPaginator {
	if params == nil {
		params = &ListApplicationComponentsInput{}
	}

	options := ListApplicationComponentsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListApplicationComponentsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListApplicationComponentsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListApplicationComponents page.
func (p *ListApplicationComponentsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListApplicationComponentsOutput, error) {
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
	result, err := p.client.ListApplicationComponents(ctx, &params, optFns...)
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

// ListApplicationComponentsAPIClient is a client that implements the
// ListApplicationComponents operation.
type ListApplicationComponentsAPIClient interface {
	ListApplicationComponents(context.Context, *ListApplicationComponentsInput, ...func(*Options)) (*ListApplicationComponentsOutput, error)
}

var _ ListApplicationComponentsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListApplicationComponents(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListApplicationComponents",
	}
}
