// Code generated by smithy-go-codegen DO NOT EDIT.

package outposts

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/outposts/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets the instance types that an Outpost can support in InstanceTypeCapacity .
// This will generally include instance types that are not currently configured and
// therefore cannot be launched with the current Outpost capacity configuration.
func (c *Client) GetOutpostSupportedInstanceTypes(ctx context.Context, params *GetOutpostSupportedInstanceTypesInput, optFns ...func(*Options)) (*GetOutpostSupportedInstanceTypesOutput, error) {
	if params == nil {
		params = &GetOutpostSupportedInstanceTypesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetOutpostSupportedInstanceTypes", params, optFns, c.addOperationGetOutpostSupportedInstanceTypesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetOutpostSupportedInstanceTypesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetOutpostSupportedInstanceTypesInput struct {

	// The ID or ARN of the Outpost.
	//
	// This member is required.
	OutpostIdentifier *string

	// The maximum page size.
	MaxResults *int32

	// The pagination token.
	NextToken *string

	// The ID for the Amazon Web Services Outposts order.
	OrderId *string

	noSmithyDocumentSerde
}

type GetOutpostSupportedInstanceTypesOutput struct {

	// Information about the instance types.
	InstanceTypes []types.InstanceTypeItem

	// The pagination token.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetOutpostSupportedInstanceTypesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetOutpostSupportedInstanceTypes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetOutpostSupportedInstanceTypes{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetOutpostSupportedInstanceTypes"); err != nil {
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
	if err = addOpGetOutpostSupportedInstanceTypesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetOutpostSupportedInstanceTypes(options.Region), middleware.Before); err != nil {
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

// GetOutpostSupportedInstanceTypesPaginatorOptions is the paginator options for
// GetOutpostSupportedInstanceTypes
type GetOutpostSupportedInstanceTypesPaginatorOptions struct {
	// The maximum page size.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetOutpostSupportedInstanceTypesPaginator is a paginator for
// GetOutpostSupportedInstanceTypes
type GetOutpostSupportedInstanceTypesPaginator struct {
	options   GetOutpostSupportedInstanceTypesPaginatorOptions
	client    GetOutpostSupportedInstanceTypesAPIClient
	params    *GetOutpostSupportedInstanceTypesInput
	nextToken *string
	firstPage bool
}

// NewGetOutpostSupportedInstanceTypesPaginator returns a new
// GetOutpostSupportedInstanceTypesPaginator
func NewGetOutpostSupportedInstanceTypesPaginator(client GetOutpostSupportedInstanceTypesAPIClient, params *GetOutpostSupportedInstanceTypesInput, optFns ...func(*GetOutpostSupportedInstanceTypesPaginatorOptions)) *GetOutpostSupportedInstanceTypesPaginator {
	if params == nil {
		params = &GetOutpostSupportedInstanceTypesInput{}
	}

	options := GetOutpostSupportedInstanceTypesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetOutpostSupportedInstanceTypesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetOutpostSupportedInstanceTypesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetOutpostSupportedInstanceTypes page.
func (p *GetOutpostSupportedInstanceTypesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetOutpostSupportedInstanceTypesOutput, error) {
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
	result, err := p.client.GetOutpostSupportedInstanceTypes(ctx, &params, optFns...)
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

// GetOutpostSupportedInstanceTypesAPIClient is a client that implements the
// GetOutpostSupportedInstanceTypes operation.
type GetOutpostSupportedInstanceTypesAPIClient interface {
	GetOutpostSupportedInstanceTypes(context.Context, *GetOutpostSupportedInstanceTypesInput, ...func(*Options)) (*GetOutpostSupportedInstanceTypesOutput, error)
}

var _ GetOutpostSupportedInstanceTypesAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opGetOutpostSupportedInstanceTypes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetOutpostSupportedInstanceTypes",
	}
}
