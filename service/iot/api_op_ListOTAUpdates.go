// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists OTA updates.
//
// Requires permission to access the [ListOTAUpdates] action.
//
// [ListOTAUpdates]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func (c *Client) ListOTAUpdates(ctx context.Context, params *ListOTAUpdatesInput, optFns ...func(*Options)) (*ListOTAUpdatesOutput, error) {
	if params == nil {
		params = &ListOTAUpdatesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListOTAUpdates", params, optFns, c.addOperationListOTAUpdatesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListOTAUpdatesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListOTAUpdatesInput struct {

	// The maximum number of results to return at one time.
	MaxResults *int32

	// A token used to retrieve the next set of results.
	NextToken *string

	// The OTA update job status.
	OtaUpdateStatus types.OTAUpdateStatus

	noSmithyDocumentSerde
}

type ListOTAUpdatesOutput struct {

	// A token to use to get the next set of results.
	NextToken *string

	// A list of OTA update jobs.
	OtaUpdates []types.OTAUpdateSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListOTAUpdatesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListOTAUpdates{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListOTAUpdates{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListOTAUpdates"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListOTAUpdates(options.Region), middleware.Before); err != nil {
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

// ListOTAUpdatesAPIClient is a client that implements the ListOTAUpdates
// operation.
type ListOTAUpdatesAPIClient interface {
	ListOTAUpdates(context.Context, *ListOTAUpdatesInput, ...func(*Options)) (*ListOTAUpdatesOutput, error)
}

var _ ListOTAUpdatesAPIClient = (*Client)(nil)

// ListOTAUpdatesPaginatorOptions is the paginator options for ListOTAUpdates
type ListOTAUpdatesPaginatorOptions struct {
	// The maximum number of results to return at one time.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListOTAUpdatesPaginator is a paginator for ListOTAUpdates
type ListOTAUpdatesPaginator struct {
	options   ListOTAUpdatesPaginatorOptions
	client    ListOTAUpdatesAPIClient
	params    *ListOTAUpdatesInput
	nextToken *string
	firstPage bool
}

// NewListOTAUpdatesPaginator returns a new ListOTAUpdatesPaginator
func NewListOTAUpdatesPaginator(client ListOTAUpdatesAPIClient, params *ListOTAUpdatesInput, optFns ...func(*ListOTAUpdatesPaginatorOptions)) *ListOTAUpdatesPaginator {
	if params == nil {
		params = &ListOTAUpdatesInput{}
	}

	options := ListOTAUpdatesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListOTAUpdatesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListOTAUpdatesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListOTAUpdates page.
func (p *ListOTAUpdatesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListOTAUpdatesOutput, error) {
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

	result, err := p.client.ListOTAUpdates(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListOTAUpdates(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListOTAUpdates",
	}
}
