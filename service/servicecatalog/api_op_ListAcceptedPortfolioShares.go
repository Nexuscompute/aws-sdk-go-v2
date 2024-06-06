// Code generated by smithy-go-codegen DO NOT EDIT.

package servicecatalog

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all imported portfolios for which account-to-account shares were accepted
// by this account. By specifying the PortfolioShareType , you can list portfolios
// for which organizational shares were accepted by this account.
func (c *Client) ListAcceptedPortfolioShares(ctx context.Context, params *ListAcceptedPortfolioSharesInput, optFns ...func(*Options)) (*ListAcceptedPortfolioSharesOutput, error) {
	if params == nil {
		params = &ListAcceptedPortfolioSharesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAcceptedPortfolioShares", params, optFns, c.addOperationListAcceptedPortfolioSharesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAcceptedPortfolioSharesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAcceptedPortfolioSharesInput struct {

	// The language code.
	//
	//   - jp - Japanese
	//
	//   - zh - Chinese
	AcceptLanguage *string

	// The maximum number of items to return with this call.
	PageSize int32

	// The page token for the next set of results. To retrieve the first set of
	// results, use null.
	PageToken *string

	// The type of shared portfolios to list. The default is to list imported
	// portfolios.
	//
	//   - AWS_ORGANIZATIONS - List portfolios accepted and shared via organizational
	//   sharing by the management account or delegated administrator of your
	//   organization.
	//
	//   - AWS_SERVICECATALOG - Deprecated type.
	//
	//   - IMPORTED - List imported portfolios that have been accepted and shared
	//   through account-to-account sharing.
	PortfolioShareType types.PortfolioShareType

	noSmithyDocumentSerde
}

type ListAcceptedPortfolioSharesOutput struct {

	// The page token to use to retrieve the next set of results. If there are no
	// additional results, this value is null.
	NextPageToken *string

	// Information about the portfolios.
	PortfolioDetails []types.PortfolioDetail

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAcceptedPortfolioSharesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListAcceptedPortfolioShares{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListAcceptedPortfolioShares{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListAcceptedPortfolioShares"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAcceptedPortfolioShares(options.Region), middleware.Before); err != nil {
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

// ListAcceptedPortfolioSharesAPIClient is a client that implements the
// ListAcceptedPortfolioShares operation.
type ListAcceptedPortfolioSharesAPIClient interface {
	ListAcceptedPortfolioShares(context.Context, *ListAcceptedPortfolioSharesInput, ...func(*Options)) (*ListAcceptedPortfolioSharesOutput, error)
}

var _ ListAcceptedPortfolioSharesAPIClient = (*Client)(nil)

// ListAcceptedPortfolioSharesPaginatorOptions is the paginator options for
// ListAcceptedPortfolioShares
type ListAcceptedPortfolioSharesPaginatorOptions struct {
	// The maximum number of items to return with this call.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAcceptedPortfolioSharesPaginator is a paginator for
// ListAcceptedPortfolioShares
type ListAcceptedPortfolioSharesPaginator struct {
	options   ListAcceptedPortfolioSharesPaginatorOptions
	client    ListAcceptedPortfolioSharesAPIClient
	params    *ListAcceptedPortfolioSharesInput
	nextToken *string
	firstPage bool
}

// NewListAcceptedPortfolioSharesPaginator returns a new
// ListAcceptedPortfolioSharesPaginator
func NewListAcceptedPortfolioSharesPaginator(client ListAcceptedPortfolioSharesAPIClient, params *ListAcceptedPortfolioSharesInput, optFns ...func(*ListAcceptedPortfolioSharesPaginatorOptions)) *ListAcceptedPortfolioSharesPaginator {
	if params == nil {
		params = &ListAcceptedPortfolioSharesInput{}
	}

	options := ListAcceptedPortfolioSharesPaginatorOptions{}
	if params.PageSize != 0 {
		options.Limit = params.PageSize
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAcceptedPortfolioSharesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.PageToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAcceptedPortfolioSharesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListAcceptedPortfolioShares page.
func (p *ListAcceptedPortfolioSharesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAcceptedPortfolioSharesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.PageToken = p.nextToken

	params.PageSize = p.options.Limit

	result, err := p.client.ListAcceptedPortfolioShares(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextPageToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListAcceptedPortfolioShares(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListAcceptedPortfolioShares",
	}
}
