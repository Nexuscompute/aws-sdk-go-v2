// Code generated by smithy-go-codegen DO NOT EDIT.

package managedblockchainquery

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/managedblockchainquery/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all the transaction events for a transaction.
func (c *Client) ListTransactions(ctx context.Context, params *ListTransactionsInput, optFns ...func(*Options)) (*ListTransactionsOutput, error) {
	if params == nil {
		params = &ListTransactionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListTransactions", params, optFns, c.addOperationListTransactionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListTransactionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListTransactionsInput struct {

	// The address (either a contract or wallet), whose transactions are being
	// requested.
	//
	// This member is required.
	Address *string

	// The blockchain network where the transactions occurred.
	//
	// This member is required.
	Network types.QueryNetwork

	// This filter is used to include transactions in the response that haven't
	// reached [finality]. Transactions that have reached finality are always part of the
	// response.
	//
	// [finality]: https://docs.aws.amazon.com/managed-blockchain/latest/ambq-dg/key-concepts.html#finality
	ConfirmationStatusFilter *types.ConfirmationStatusFilter

	// The container for time.
	FromBlockchainInstant *types.BlockchainInstant

	// The maximum number of transactions to list.
	//
	// Default: 100
	//
	// Even if additional results can be retrieved, the request can return less
	// results than maxResults or an empty array of results.
	//
	// To retrieve the next set of results, make another request with the returned
	// nextToken value. The value of nextToken is null when there are no more results
	// to return
	MaxResults *int32

	// The pagination token that indicates the next set of results to retrieve.
	NextToken *string

	// The order by which the results will be sorted.
	Sort *types.ListTransactionsSort

	// The container for time.
	ToBlockchainInstant *types.BlockchainInstant

	noSmithyDocumentSerde
}

type ListTransactionsOutput struct {

	// The array of transactions returned by the request.
	//
	// This member is required.
	Transactions []types.TransactionOutputItem

	// The pagination token that indicates the next set of results to retrieve.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListTransactionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListTransactions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListTransactions{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListTransactions"); err != nil {
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
	if err = addOpListTransactionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListTransactions(options.Region), middleware.Before); err != nil {
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

// ListTransactionsPaginatorOptions is the paginator options for ListTransactions
type ListTransactionsPaginatorOptions struct {
	// The maximum number of transactions to list.
	//
	// Default: 100
	//
	// Even if additional results can be retrieved, the request can return less
	// results than maxResults or an empty array of results.
	//
	// To retrieve the next set of results, make another request with the returned
	// nextToken value. The value of nextToken is null when there are no more results
	// to return
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListTransactionsPaginator is a paginator for ListTransactions
type ListTransactionsPaginator struct {
	options   ListTransactionsPaginatorOptions
	client    ListTransactionsAPIClient
	params    *ListTransactionsInput
	nextToken *string
	firstPage bool
}

// NewListTransactionsPaginator returns a new ListTransactionsPaginator
func NewListTransactionsPaginator(client ListTransactionsAPIClient, params *ListTransactionsInput, optFns ...func(*ListTransactionsPaginatorOptions)) *ListTransactionsPaginator {
	if params == nil {
		params = &ListTransactionsInput{}
	}

	options := ListTransactionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListTransactionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListTransactionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListTransactions page.
func (p *ListTransactionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListTransactionsOutput, error) {
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
	result, err := p.client.ListTransactions(ctx, &params, optFns...)
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

// ListTransactionsAPIClient is a client that implements the ListTransactions
// operation.
type ListTransactionsAPIClient interface {
	ListTransactions(context.Context, *ListTransactionsInput, ...func(*Options)) (*ListTransactionsOutput, error)
}

var _ ListTransactionsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListTransactions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListTransactions",
	}
}
