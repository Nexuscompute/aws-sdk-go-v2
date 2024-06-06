// Code generated by smithy-go-codegen DO NOT EDIT.

package controltower

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/controltower/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of summaries describing EnabledBaseline resources. You can
// filter the list by the corresponding Baseline or Target of the EnabledBaseline
// resources. For usage examples, see [the Amazon Web Services Control Tower User Guide].
//
// [the Amazon Web Services Control Tower User Guide]: https://docs.aws.amazon.com/controltower/latest/userguide/baseline-api-examples.html
func (c *Client) ListEnabledBaselines(ctx context.Context, params *ListEnabledBaselinesInput, optFns ...func(*Options)) (*ListEnabledBaselinesOutput, error) {
	if params == nil {
		params = &ListEnabledBaselinesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListEnabledBaselines", params, optFns, c.addOperationListEnabledBaselinesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListEnabledBaselinesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListEnabledBaselinesInput struct {

	// A filter applied on the ListEnabledBaseline operation. Allowed filters are
	// baselineIdentifiers and targetIdentifiers . The filter can be applied for
	// either, or both.
	Filter *types.EnabledBaselineFilter

	// The maximum number of results to be shown.
	MaxResults *int32

	// A pagination token.
	NextToken *string

	noSmithyDocumentSerde
}

type ListEnabledBaselinesOutput struct {

	// Retuens a list of summaries of EnabledBaseline resources.
	//
	// This member is required.
	EnabledBaselines []types.EnabledBaselineSummary

	// A pagination token.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListEnabledBaselinesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListEnabledBaselines{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListEnabledBaselines{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListEnabledBaselines"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListEnabledBaselines(options.Region), middleware.Before); err != nil {
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

// ListEnabledBaselinesAPIClient is a client that implements the
// ListEnabledBaselines operation.
type ListEnabledBaselinesAPIClient interface {
	ListEnabledBaselines(context.Context, *ListEnabledBaselinesInput, ...func(*Options)) (*ListEnabledBaselinesOutput, error)
}

var _ ListEnabledBaselinesAPIClient = (*Client)(nil)

// ListEnabledBaselinesPaginatorOptions is the paginator options for
// ListEnabledBaselines
type ListEnabledBaselinesPaginatorOptions struct {
	// The maximum number of results to be shown.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListEnabledBaselinesPaginator is a paginator for ListEnabledBaselines
type ListEnabledBaselinesPaginator struct {
	options   ListEnabledBaselinesPaginatorOptions
	client    ListEnabledBaselinesAPIClient
	params    *ListEnabledBaselinesInput
	nextToken *string
	firstPage bool
}

// NewListEnabledBaselinesPaginator returns a new ListEnabledBaselinesPaginator
func NewListEnabledBaselinesPaginator(client ListEnabledBaselinesAPIClient, params *ListEnabledBaselinesInput, optFns ...func(*ListEnabledBaselinesPaginatorOptions)) *ListEnabledBaselinesPaginator {
	if params == nil {
		params = &ListEnabledBaselinesInput{}
	}

	options := ListEnabledBaselinesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListEnabledBaselinesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListEnabledBaselinesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListEnabledBaselines page.
func (p *ListEnabledBaselinesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListEnabledBaselinesOutput, error) {
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

	result, err := p.client.ListEnabledBaselines(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListEnabledBaselines(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListEnabledBaselines",
	}
}
