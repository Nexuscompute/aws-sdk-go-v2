// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatch

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

//	Returns a list that contains the number of managed Contributor Insights rules
//
// in your account.
func (c *Client) ListManagedInsightRules(ctx context.Context, params *ListManagedInsightRulesInput, optFns ...func(*Options)) (*ListManagedInsightRulesOutput, error) {
	if params == nil {
		params = &ListManagedInsightRulesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListManagedInsightRules", params, optFns, c.addOperationListManagedInsightRulesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListManagedInsightRulesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListManagedInsightRulesInput struct {

	//  The ARN of an Amazon Web Services resource that has managed Contributor
	// Insights rules.
	//
	// This member is required.
	ResourceARN *string

	//  The maximum number of results to return in one operation. If you omit this
	// parameter, the default number is used. The default number is 100 .
	MaxResults *int32

	//  Include this value to get the next set of rules if the value was returned by
	// the previous operation.
	NextToken *string

	noSmithyDocumentSerde
}

type ListManagedInsightRulesOutput struct {

	//  The managed rules that are available for the specified Amazon Web Services
	// resource.
	ManagedRules []types.ManagedRuleDescription

	//  Include this value to get the next set of rules if the value was returned by
	// the previous operation.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListManagedInsightRulesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpListManagedInsightRules{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpListManagedInsightRules{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListManagedInsightRules"); err != nil {
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
	if err = addOpListManagedInsightRulesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListManagedInsightRules(options.Region), middleware.Before); err != nil {
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

// ListManagedInsightRulesAPIClient is a client that implements the
// ListManagedInsightRules operation.
type ListManagedInsightRulesAPIClient interface {
	ListManagedInsightRules(context.Context, *ListManagedInsightRulesInput, ...func(*Options)) (*ListManagedInsightRulesOutput, error)
}

var _ ListManagedInsightRulesAPIClient = (*Client)(nil)

// ListManagedInsightRulesPaginatorOptions is the paginator options for
// ListManagedInsightRules
type ListManagedInsightRulesPaginatorOptions struct {
	//  The maximum number of results to return in one operation. If you omit this
	// parameter, the default number is used. The default number is 100 .
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListManagedInsightRulesPaginator is a paginator for ListManagedInsightRules
type ListManagedInsightRulesPaginator struct {
	options   ListManagedInsightRulesPaginatorOptions
	client    ListManagedInsightRulesAPIClient
	params    *ListManagedInsightRulesInput
	nextToken *string
	firstPage bool
}

// NewListManagedInsightRulesPaginator returns a new
// ListManagedInsightRulesPaginator
func NewListManagedInsightRulesPaginator(client ListManagedInsightRulesAPIClient, params *ListManagedInsightRulesInput, optFns ...func(*ListManagedInsightRulesPaginatorOptions)) *ListManagedInsightRulesPaginator {
	if params == nil {
		params = &ListManagedInsightRulesInput{}
	}

	options := ListManagedInsightRulesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListManagedInsightRulesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListManagedInsightRulesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListManagedInsightRules page.
func (p *ListManagedInsightRulesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListManagedInsightRulesOutput, error) {
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

	result, err := p.client.ListManagedInsightRules(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListManagedInsightRules(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListManagedInsightRules",
	}
}
