// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a collection of fleet resources in an Amazon Web Services Region. You
// can filter the result set to find only those fleets that are deployed with a
// specific build or script. For fleets that have multiple locations, this
// operation retrieves fleets based on their home Region only.
//
// You can use operation in the following ways:
//
//   - To get a list of all fleets in a Region, don't provide a build or script
//     identifier.
//
//   - To get a list of all fleets where a specific game build is deployed,
//     provide the build ID.
//
//   - To get a list of all Realtime Servers fleets with a specific configuration
//     script, provide the script ID.
//
// Use the pagination parameters to retrieve results as a set of sequential pages.
//
// If successful, this operation returns a list of fleet IDs that match the
// request parameters. A NextToken value is also returned if there are more result
// pages to retrieve.
//
// Fleet IDs are returned in no particular order.
func (c *Client) ListFleets(ctx context.Context, params *ListFleetsInput, optFns ...func(*Options)) (*ListFleetsOutput, error) {
	if params == nil {
		params = &ListFleetsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListFleets", params, optFns, c.addOperationListFleetsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListFleetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListFleetsInput struct {

	// A unique identifier for the build to request fleets for. Use this parameter to
	// return only fleets using a specified build. Use either the build ID or ARN
	// value.
	BuildId *string

	// The maximum number of results to return. Use this parameter with NextToken to
	// get results as a set of sequential pages.
	Limit *int32

	// A token that indicates the start of the next sequential page of results. Use
	// the token that is returned with a previous call to this operation. To start at
	// the beginning of the result set, do not specify a value.
	NextToken *string

	// A unique identifier for the Realtime script to request fleets for. Use this
	// parameter to return only fleets using a specified script. Use either the script
	// ID or ARN value.
	ScriptId *string

	noSmithyDocumentSerde
}

type ListFleetsOutput struct {

	// A set of fleet IDs that match the list request.
	FleetIds []string

	// A token that indicates where to resume retrieving results on the next call to
	// this operation. If no token is returned, these results represent the end of the
	// list.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListFleetsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListFleets{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListFleets{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListFleets"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListFleets(options.Region), middleware.Before); err != nil {
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

// ListFleetsPaginatorOptions is the paginator options for ListFleets
type ListFleetsPaginatorOptions struct {
	// The maximum number of results to return. Use this parameter with NextToken to
	// get results as a set of sequential pages.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListFleetsPaginator is a paginator for ListFleets
type ListFleetsPaginator struct {
	options   ListFleetsPaginatorOptions
	client    ListFleetsAPIClient
	params    *ListFleetsInput
	nextToken *string
	firstPage bool
}

// NewListFleetsPaginator returns a new ListFleetsPaginator
func NewListFleetsPaginator(client ListFleetsAPIClient, params *ListFleetsInput, optFns ...func(*ListFleetsPaginatorOptions)) *ListFleetsPaginator {
	if params == nil {
		params = &ListFleetsInput{}
	}

	options := ListFleetsPaginatorOptions{}
	if params.Limit != nil {
		options.Limit = *params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListFleetsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListFleetsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListFleets page.
func (p *ListFleetsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListFleetsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.Limit = limit

	optFns = append([]func(*Options){
		addIsPaginatorUserAgent,
	}, optFns...)
	result, err := p.client.ListFleets(ctx, &params, optFns...)
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

// ListFleetsAPIClient is a client that implements the ListFleets operation.
type ListFleetsAPIClient interface {
	ListFleets(context.Context, *ListFleetsInput, ...func(*Options)) (*ListFleetsOutput, error)
}

var _ ListFleetsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListFleets(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListFleets",
	}
}
