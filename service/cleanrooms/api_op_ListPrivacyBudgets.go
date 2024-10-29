// Code generated by smithy-go-codegen DO NOT EDIT.

package cleanrooms

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/cleanrooms/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns detailed information about the privacy budgets in a specified
// membership.
func (c *Client) ListPrivacyBudgets(ctx context.Context, params *ListPrivacyBudgetsInput, optFns ...func(*Options)) (*ListPrivacyBudgetsOutput, error) {
	if params == nil {
		params = &ListPrivacyBudgetsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListPrivacyBudgets", params, optFns, c.addOperationListPrivacyBudgetsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListPrivacyBudgetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPrivacyBudgetsInput struct {

	// A unique identifier for one of your memberships for a collaboration. The
	// privacy budget is retrieved from the collaboration that this membership belongs
	// to. Accepts a membership ID.
	//
	// This member is required.
	MembershipIdentifier *string

	// The privacy budget type.
	//
	// This member is required.
	PrivacyBudgetType types.PrivacyBudgetType

	// The maximum number of results that are returned for an API request call. The
	// service chooses a default number if you don't set one. The service might return
	// a `nextToken` even if the `maxResults` value has not been met.
	MaxResults *int32

	// The pagination token that's used to fetch the next set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListPrivacyBudgetsOutput struct {

	// An array that summarizes the privacy budgets. The summary includes
	// collaboration information, membership information, privacy budget template
	// information, and privacy budget details.
	//
	// This member is required.
	PrivacyBudgetSummaries []types.PrivacyBudgetSummary

	// The pagination token that's used to fetch the next set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListPrivacyBudgetsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListPrivacyBudgets{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListPrivacyBudgets{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListPrivacyBudgets"); err != nil {
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
	if err = addOpListPrivacyBudgetsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListPrivacyBudgets(options.Region), middleware.Before); err != nil {
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

// ListPrivacyBudgetsPaginatorOptions is the paginator options for
// ListPrivacyBudgets
type ListPrivacyBudgetsPaginatorOptions struct {
	// The maximum number of results that are returned for an API request call. The
	// service chooses a default number if you don't set one. The service might return
	// a `nextToken` even if the `maxResults` value has not been met.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListPrivacyBudgetsPaginator is a paginator for ListPrivacyBudgets
type ListPrivacyBudgetsPaginator struct {
	options   ListPrivacyBudgetsPaginatorOptions
	client    ListPrivacyBudgetsAPIClient
	params    *ListPrivacyBudgetsInput
	nextToken *string
	firstPage bool
}

// NewListPrivacyBudgetsPaginator returns a new ListPrivacyBudgetsPaginator
func NewListPrivacyBudgetsPaginator(client ListPrivacyBudgetsAPIClient, params *ListPrivacyBudgetsInput, optFns ...func(*ListPrivacyBudgetsPaginatorOptions)) *ListPrivacyBudgetsPaginator {
	if params == nil {
		params = &ListPrivacyBudgetsInput{}
	}

	options := ListPrivacyBudgetsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListPrivacyBudgetsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListPrivacyBudgetsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListPrivacyBudgets page.
func (p *ListPrivacyBudgetsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListPrivacyBudgetsOutput, error) {
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
	result, err := p.client.ListPrivacyBudgets(ctx, &params, optFns...)
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

// ListPrivacyBudgetsAPIClient is a client that implements the ListPrivacyBudgets
// operation.
type ListPrivacyBudgetsAPIClient interface {
	ListPrivacyBudgets(context.Context, *ListPrivacyBudgetsInput, ...func(*Options)) (*ListPrivacyBudgetsOutput, error)
}

var _ ListPrivacyBudgetsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListPrivacyBudgets(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListPrivacyBudgets",
	}
}
