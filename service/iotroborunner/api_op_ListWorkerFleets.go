// Code generated by smithy-go-codegen DO NOT EDIT.

package iotroborunner

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotroborunner/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Grants permission to list worker fleets
func (c *Client) ListWorkerFleets(ctx context.Context, params *ListWorkerFleetsInput, optFns ...func(*Options)) (*ListWorkerFleetsOutput, error) {
	if params == nil {
		params = &ListWorkerFleetsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListWorkerFleets", params, optFns, c.addOperationListWorkerFleetsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListWorkerFleetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListWorkerFleetsInput struct {

	// Site ARN.
	//
	// This member is required.
	Site *string

	// Maximum number of results to retrieve in a single ListWorkerFleets call.
	MaxResults *int32

	// Pagination token returned when another page of data exists. Provide it in your
	// next call to the API to receive the next page.
	NextToken *string

	noSmithyDocumentSerde
}

type ListWorkerFleetsOutput struct {

	// Pagination token returned when another page of data exists. Provide it in your
	// next call to the API to receive the next page.
	NextToken *string

	// List of worker fleets.
	WorkerFleets []types.WorkerFleet

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListWorkerFleetsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListWorkerFleets{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListWorkerFleets{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpListWorkerFleetsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListWorkerFleets(options.Region), middleware.Before); err != nil {
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
	return nil
}

// ListWorkerFleetsAPIClient is a client that implements the ListWorkerFleets
// operation.
type ListWorkerFleetsAPIClient interface {
	ListWorkerFleets(context.Context, *ListWorkerFleetsInput, ...func(*Options)) (*ListWorkerFleetsOutput, error)
}

var _ ListWorkerFleetsAPIClient = (*Client)(nil)

// ListWorkerFleetsPaginatorOptions is the paginator options for ListWorkerFleets
type ListWorkerFleetsPaginatorOptions struct {
	// Maximum number of results to retrieve in a single ListWorkerFleets call.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListWorkerFleetsPaginator is a paginator for ListWorkerFleets
type ListWorkerFleetsPaginator struct {
	options   ListWorkerFleetsPaginatorOptions
	client    ListWorkerFleetsAPIClient
	params    *ListWorkerFleetsInput
	nextToken *string
	firstPage bool
}

// NewListWorkerFleetsPaginator returns a new ListWorkerFleetsPaginator
func NewListWorkerFleetsPaginator(client ListWorkerFleetsAPIClient, params *ListWorkerFleetsInput, optFns ...func(*ListWorkerFleetsPaginatorOptions)) *ListWorkerFleetsPaginator {
	if params == nil {
		params = &ListWorkerFleetsInput{}
	}

	options := ListWorkerFleetsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListWorkerFleetsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListWorkerFleetsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListWorkerFleets page.
func (p *ListWorkerFleetsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListWorkerFleetsOutput, error) {
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

	result, err := p.client.ListWorkerFleets(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListWorkerFleets(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotroborunner",
		OperationName: "ListWorkerFleets",
	}
}