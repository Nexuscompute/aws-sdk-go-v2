// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticloadbalancingv2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes the specified rules or the rules for the specified listener. You must
// specify either a listener or one or more rules.
func (c *Client) DescribeRules(ctx context.Context, params *DescribeRulesInput, optFns ...func(*Options)) (*DescribeRulesOutput, error) {
	if params == nil {
		params = &DescribeRulesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeRules", params, optFns, c.addOperationDescribeRulesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeRulesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeRulesInput struct {

	// The Amazon Resource Name (ARN) of the listener.
	ListenerArn *string

	// The marker for the next set of results. (You received this marker from a
	// previous call.)
	Marker *string

	// The maximum number of results to return with this call.
	PageSize *int32

	// The Amazon Resource Names (ARN) of the rules.
	RuleArns []string

	noSmithyDocumentSerde
}

type DescribeRulesOutput struct {

	// If there are additional results, this is the marker for the next set of
	// results. Otherwise, this is null.
	NextMarker *string

	// Information about the rules.
	Rules []types.Rule

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeRulesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeRules{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeRules{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeRules"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeRules(options.Region), middleware.Before); err != nil {
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

// DescribeRulesPaginatorOptions is the paginator options for DescribeRules
type DescribeRulesPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeRulesPaginator is a paginator for DescribeRules
type DescribeRulesPaginator struct {
	options   DescribeRulesPaginatorOptions
	client    DescribeRulesAPIClient
	params    *DescribeRulesInput
	nextToken *string
	firstPage bool
}

// NewDescribeRulesPaginator returns a new DescribeRulesPaginator
func NewDescribeRulesPaginator(client DescribeRulesAPIClient, params *DescribeRulesInput, optFns ...func(*DescribeRulesPaginatorOptions)) *DescribeRulesPaginator {
	if params == nil {
		params = &DescribeRulesInput{}
	}

	options := DescribeRulesPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeRulesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeRulesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeRules page.
func (p *DescribeRulesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeRulesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	optFns = append([]func(*Options){
		addIsPaginatorUserAgent,
	}, optFns...)
	result, err := p.client.DescribeRules(ctx, &params, optFns...)
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

// DescribeRulesAPIClient is a client that implements the DescribeRules operation.
type DescribeRulesAPIClient interface {
	DescribeRules(context.Context, *DescribeRulesInput, ...func(*Options)) (*DescribeRulesOutput, error)
}

var _ DescribeRulesAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opDescribeRules(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeRules",
	}
}
