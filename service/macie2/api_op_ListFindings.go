// Code generated by smithy-go-codegen DO NOT EDIT.

package macie2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/macie2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a subset of information about one or more findings.
func (c *Client) ListFindings(ctx context.Context, params *ListFindingsInput, optFns ...func(*Options)) (*ListFindingsOutput, error) {
	if params == nil {
		params = &ListFindingsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListFindings", params, optFns, c.addOperationListFindingsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListFindingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListFindingsInput struct {

	// The criteria to use to filter the results.
	FindingCriteria *types.FindingCriteria

	// The maximum number of items to include in each page of the response.
	MaxResults *int32

	// The nextToken string that specifies which page of results to return in a
	// paginated response.
	NextToken *string

	// The criteria to use to sort the results.
	SortCriteria *types.SortCriteria

	noSmithyDocumentSerde
}

type ListFindingsOutput struct {

	// An array of strings, where each string is the unique identifier for a finding
	// that matches the filter criteria specified in the request.
	FindingIds []string

	// The string to use in a subsequent request to get the next page of results in a
	// paginated response. This value is null if there are no additional pages.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListFindingsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListFindings{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListFindings{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListFindings"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListFindings(options.Region), middleware.Before); err != nil {
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

// ListFindingsAPIClient is a client that implements the ListFindings operation.
type ListFindingsAPIClient interface {
	ListFindings(context.Context, *ListFindingsInput, ...func(*Options)) (*ListFindingsOutput, error)
}

var _ ListFindingsAPIClient = (*Client)(nil)

// ListFindingsPaginatorOptions is the paginator options for ListFindings
type ListFindingsPaginatorOptions struct {
	// The maximum number of items to include in each page of the response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListFindingsPaginator is a paginator for ListFindings
type ListFindingsPaginator struct {
	options   ListFindingsPaginatorOptions
	client    ListFindingsAPIClient
	params    *ListFindingsInput
	nextToken *string
	firstPage bool
}

// NewListFindingsPaginator returns a new ListFindingsPaginator
func NewListFindingsPaginator(client ListFindingsAPIClient, params *ListFindingsInput, optFns ...func(*ListFindingsPaginatorOptions)) *ListFindingsPaginator {
	if params == nil {
		params = &ListFindingsInput{}
	}

	options := ListFindingsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListFindingsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListFindingsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListFindings page.
func (p *ListFindingsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListFindingsOutput, error) {
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

	result, err := p.client.ListFindings(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListFindings(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListFindings",
	}
}
