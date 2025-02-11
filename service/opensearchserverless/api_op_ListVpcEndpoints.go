// Code generated by smithy-go-codegen DO NOT EDIT.

package opensearchserverless

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/opensearchserverless/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the OpenSearch Serverless-managed interface VPC endpoints associated
// with the current account. For more information, see [Access Amazon OpenSearch Serverless using an interface endpoint].
//
// [Access Amazon OpenSearch Serverless using an interface endpoint]: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/serverless-vpc.html
func (c *Client) ListVpcEndpoints(ctx context.Context, params *ListVpcEndpointsInput, optFns ...func(*Options)) (*ListVpcEndpointsOutput, error) {
	if params == nil {
		params = &ListVpcEndpointsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListVpcEndpoints", params, optFns, c.addOperationListVpcEndpointsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListVpcEndpointsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListVpcEndpointsInput struct {

	// An optional parameter that specifies the maximum number of results to return.
	// You can use nextToken to get the next page of results. The default is 20.
	MaxResults *int32

	// If your initial ListVpcEndpoints operation returns a nextToken , you can include
	// the returned nextToken in subsequent ListVpcEndpoints operations, which returns
	// results in the next page.
	NextToken *string

	// Filter the results according to the current status of the VPC endpoint.
	// Possible statuses are CREATING , DELETING , UPDATING , ACTIVE , and FAILED .
	VpcEndpointFilters *types.VpcEndpointFilters

	noSmithyDocumentSerde
}

type ListVpcEndpointsOutput struct {

	// When nextToken is returned, there are more results available. The value of
	// nextToken is a unique pagination token for each page. Make the call again using
	// the returned token to retrieve the next page.
	NextToken *string

	// Details about each VPC endpoint, including the name and current status.
	VpcEndpointSummaries []types.VpcEndpointSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListVpcEndpointsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpListVpcEndpoints{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpListVpcEndpoints{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListVpcEndpoints"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListVpcEndpoints(options.Region), middleware.Before); err != nil {
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

// ListVpcEndpointsPaginatorOptions is the paginator options for ListVpcEndpoints
type ListVpcEndpointsPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListVpcEndpointsPaginator is a paginator for ListVpcEndpoints
type ListVpcEndpointsPaginator struct {
	options   ListVpcEndpointsPaginatorOptions
	client    ListVpcEndpointsAPIClient
	params    *ListVpcEndpointsInput
	nextToken *string
	firstPage bool
}

// NewListVpcEndpointsPaginator returns a new ListVpcEndpointsPaginator
func NewListVpcEndpointsPaginator(client ListVpcEndpointsAPIClient, params *ListVpcEndpointsInput, optFns ...func(*ListVpcEndpointsPaginatorOptions)) *ListVpcEndpointsPaginator {
	if params == nil {
		params = &ListVpcEndpointsInput{}
	}

	options := ListVpcEndpointsPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListVpcEndpointsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListVpcEndpointsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListVpcEndpoints page.
func (p *ListVpcEndpointsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListVpcEndpointsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	optFns = append([]func(*Options){
		addIsPaginatorUserAgent,
	}, optFns...)
	result, err := p.client.ListVpcEndpoints(ctx, &params, optFns...)
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

// ListVpcEndpointsAPIClient is a client that implements the ListVpcEndpoints
// operation.
type ListVpcEndpointsAPIClient interface {
	ListVpcEndpoints(context.Context, *ListVpcEndpointsInput, ...func(*Options)) (*ListVpcEndpointsOutput, error)
}

var _ ListVpcEndpointsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListVpcEndpoints(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListVpcEndpoints",
	}
}
