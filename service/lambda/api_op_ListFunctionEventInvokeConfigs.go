// Code generated by smithy-go-codegen DO NOT EDIT.

package lambda

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a list of configurations for asynchronous invocation for a function.
//
// To configure options for asynchronous invocation, use PutFunctionEventInvokeConfig.
func (c *Client) ListFunctionEventInvokeConfigs(ctx context.Context, params *ListFunctionEventInvokeConfigsInput, optFns ...func(*Options)) (*ListFunctionEventInvokeConfigsOutput, error) {
	if params == nil {
		params = &ListFunctionEventInvokeConfigsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListFunctionEventInvokeConfigs", params, optFns, c.addOperationListFunctionEventInvokeConfigsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListFunctionEventInvokeConfigsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListFunctionEventInvokeConfigsInput struct {

	// The name or ARN of the Lambda function.
	//
	// Name formats
	//
	//   - Function name - my-function .
	//
	//   - Function ARN - arn:aws:lambda:us-west-2:123456789012:function:my-function .
	//
	//   - Partial ARN - 123456789012:function:my-function .
	//
	// The length constraint applies only to the full ARN. If you specify only the
	// function name, it is limited to 64 characters in length.
	//
	// This member is required.
	FunctionName *string

	// Specify the pagination token that's returned by a previous request to retrieve
	// the next page of results.
	Marker *string

	// The maximum number of configurations to return.
	MaxItems *int32

	noSmithyDocumentSerde
}

type ListFunctionEventInvokeConfigsOutput struct {

	// A list of configurations.
	FunctionEventInvokeConfigs []types.FunctionEventInvokeConfig

	// The pagination token that's included if more results are available.
	NextMarker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListFunctionEventInvokeConfigsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListFunctionEventInvokeConfigs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListFunctionEventInvokeConfigs{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListFunctionEventInvokeConfigs"); err != nil {
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
	if err = addOpListFunctionEventInvokeConfigsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListFunctionEventInvokeConfigs(options.Region), middleware.Before); err != nil {
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

// ListFunctionEventInvokeConfigsAPIClient is a client that implements the
// ListFunctionEventInvokeConfigs operation.
type ListFunctionEventInvokeConfigsAPIClient interface {
	ListFunctionEventInvokeConfigs(context.Context, *ListFunctionEventInvokeConfigsInput, ...func(*Options)) (*ListFunctionEventInvokeConfigsOutput, error)
}

var _ ListFunctionEventInvokeConfigsAPIClient = (*Client)(nil)

// ListFunctionEventInvokeConfigsPaginatorOptions is the paginator options for
// ListFunctionEventInvokeConfigs
type ListFunctionEventInvokeConfigsPaginatorOptions struct {
	// The maximum number of configurations to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListFunctionEventInvokeConfigsPaginator is a paginator for
// ListFunctionEventInvokeConfigs
type ListFunctionEventInvokeConfigsPaginator struct {
	options   ListFunctionEventInvokeConfigsPaginatorOptions
	client    ListFunctionEventInvokeConfigsAPIClient
	params    *ListFunctionEventInvokeConfigsInput
	nextToken *string
	firstPage bool
}

// NewListFunctionEventInvokeConfigsPaginator returns a new
// ListFunctionEventInvokeConfigsPaginator
func NewListFunctionEventInvokeConfigsPaginator(client ListFunctionEventInvokeConfigsAPIClient, params *ListFunctionEventInvokeConfigsInput, optFns ...func(*ListFunctionEventInvokeConfigsPaginatorOptions)) *ListFunctionEventInvokeConfigsPaginator {
	if params == nil {
		params = &ListFunctionEventInvokeConfigsInput{}
	}

	options := ListFunctionEventInvokeConfigsPaginatorOptions{}
	if params.MaxItems != nil {
		options.Limit = *params.MaxItems
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListFunctionEventInvokeConfigsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListFunctionEventInvokeConfigsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListFunctionEventInvokeConfigs page.
func (p *ListFunctionEventInvokeConfigsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListFunctionEventInvokeConfigsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxItems = limit

	result, err := p.client.ListFunctionEventInvokeConfigs(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextMarker

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListFunctionEventInvokeConfigs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListFunctionEventInvokeConfigs",
	}
}
