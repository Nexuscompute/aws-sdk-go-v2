// Code generated by smithy-go-codegen DO NOT EDIT.

package fms

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/fms/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns an array of PolicySummary objects.
func (c *Client) ListPolicies(ctx context.Context, params *ListPoliciesInput, optFns ...func(*Options)) (*ListPoliciesOutput, error) {
	if params == nil {
		params = &ListPoliciesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListPolicies", params, optFns, c.addOperationListPoliciesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListPoliciesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPoliciesInput struct {

	// Specifies the number of PolicySummary objects that you want Firewall Manager to
	// return for this request. If you have more PolicySummary objects than the number
	// that you specify for MaxResults , the response includes a NextToken value that
	// you can use to get another batch of PolicySummary objects.
	MaxResults *int32

	// If you specify a value for MaxResults and you have more PolicySummary objects
	// than the number that you specify for MaxResults , Firewall Manager returns a
	// NextToken value in the response that allows you to list another group of
	// PolicySummary objects. For the second and subsequent ListPolicies requests,
	// specify the value of NextToken from the previous response to get information
	// about another batch of PolicySummary objects.
	NextToken *string

	noSmithyDocumentSerde
}

type ListPoliciesOutput struct {

	// If you have more PolicySummary objects than the number that you specified for
	// MaxResults in the request, the response includes a NextToken value. To list
	// more PolicySummary objects, submit another ListPolicies request, and specify
	// the NextToken value from the response in the NextToken value in the next
	// request.
	NextToken *string

	// An array of PolicySummary objects.
	PolicyList []types.PolicySummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListPoliciesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListPolicies{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListPolicies{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListPolicies"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListPolicies(options.Region), middleware.Before); err != nil {
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

// ListPoliciesAPIClient is a client that implements the ListPolicies operation.
type ListPoliciesAPIClient interface {
	ListPolicies(context.Context, *ListPoliciesInput, ...func(*Options)) (*ListPoliciesOutput, error)
}

var _ ListPoliciesAPIClient = (*Client)(nil)

// ListPoliciesPaginatorOptions is the paginator options for ListPolicies
type ListPoliciesPaginatorOptions struct {
	// Specifies the number of PolicySummary objects that you want Firewall Manager to
	// return for this request. If you have more PolicySummary objects than the number
	// that you specify for MaxResults , the response includes a NextToken value that
	// you can use to get another batch of PolicySummary objects.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListPoliciesPaginator is a paginator for ListPolicies
type ListPoliciesPaginator struct {
	options   ListPoliciesPaginatorOptions
	client    ListPoliciesAPIClient
	params    *ListPoliciesInput
	nextToken *string
	firstPage bool
}

// NewListPoliciesPaginator returns a new ListPoliciesPaginator
func NewListPoliciesPaginator(client ListPoliciesAPIClient, params *ListPoliciesInput, optFns ...func(*ListPoliciesPaginatorOptions)) *ListPoliciesPaginator {
	if params == nil {
		params = &ListPoliciesInput{}
	}

	options := ListPoliciesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListPoliciesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListPoliciesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListPolicies page.
func (p *ListPoliciesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListPoliciesOutput, error) {
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

	result, err := p.client.ListPolicies(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListPolicies(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListPolicies",
	}
}
