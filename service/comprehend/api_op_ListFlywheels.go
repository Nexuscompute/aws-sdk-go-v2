// Code generated by smithy-go-codegen DO NOT EDIT.

package comprehend

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/comprehend/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets a list of the flywheels that you have created.
func (c *Client) ListFlywheels(ctx context.Context, params *ListFlywheelsInput, optFns ...func(*Options)) (*ListFlywheelsOutput, error) {
	if params == nil {
		params = &ListFlywheelsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListFlywheels", params, optFns, c.addOperationListFlywheelsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListFlywheelsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListFlywheelsInput struct {

	// Filters the flywheels that are returned. You can filter flywheels on their
	// status, or the date and time that they were submitted. You can only set one
	// filter at a time.
	Filter *types.FlywheelFilter

	// Maximum number of results to return in a response. The default is 100.
	MaxResults *int32

	// Identifies the next page of results to return.
	NextToken *string

	noSmithyDocumentSerde
}

type ListFlywheelsOutput struct {

	// A list of flywheel properties retrieved by the service in response to the
	// request.
	FlywheelSummaryList []types.FlywheelSummary

	// Identifies the next page of results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListFlywheelsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListFlywheels{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListFlywheels{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListFlywheels"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListFlywheels(options.Region), middleware.Before); err != nil {
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

// ListFlywheelsAPIClient is a client that implements the ListFlywheels operation.
type ListFlywheelsAPIClient interface {
	ListFlywheels(context.Context, *ListFlywheelsInput, ...func(*Options)) (*ListFlywheelsOutput, error)
}

var _ ListFlywheelsAPIClient = (*Client)(nil)

// ListFlywheelsPaginatorOptions is the paginator options for ListFlywheels
type ListFlywheelsPaginatorOptions struct {
	// Maximum number of results to return in a response. The default is 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListFlywheelsPaginator is a paginator for ListFlywheels
type ListFlywheelsPaginator struct {
	options   ListFlywheelsPaginatorOptions
	client    ListFlywheelsAPIClient
	params    *ListFlywheelsInput
	nextToken *string
	firstPage bool
}

// NewListFlywheelsPaginator returns a new ListFlywheelsPaginator
func NewListFlywheelsPaginator(client ListFlywheelsAPIClient, params *ListFlywheelsInput, optFns ...func(*ListFlywheelsPaginatorOptions)) *ListFlywheelsPaginator {
	if params == nil {
		params = &ListFlywheelsInput{}
	}

	options := ListFlywheelsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListFlywheelsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListFlywheelsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListFlywheels page.
func (p *ListFlywheelsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListFlywheelsOutput, error) {
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

	result, err := p.client.ListFlywheels(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListFlywheels(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListFlywheels",
	}
}
