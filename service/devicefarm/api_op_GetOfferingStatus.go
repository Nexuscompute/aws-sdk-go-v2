// Code generated by smithy-go-codegen DO NOT EDIT.

package devicefarm

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/devicefarm/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets the current status and future status of all offerings purchased by an AWS
// account. The response indicates how many offerings are currently available and
// the offerings that will be available in the next period. The API returns a
// NotEligible error if the user is not permitted to invoke the operation. If you
// must be able to invoke this operation, contact aws-devicefarm-support@amazon.com.
func (c *Client) GetOfferingStatus(ctx context.Context, params *GetOfferingStatusInput, optFns ...func(*Options)) (*GetOfferingStatusOutput, error) {
	if params == nil {
		params = &GetOfferingStatusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetOfferingStatus", params, optFns, c.addOperationGetOfferingStatusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetOfferingStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the request to retrieve the offering status for the specified
// customer or account.
type GetOfferingStatusInput struct {

	// An identifier that was returned from the previous call to this operation, which
	// can be used to return the next set of items in the list.
	NextToken *string

	noSmithyDocumentSerde
}

// Returns the status result for a device offering.
type GetOfferingStatusOutput struct {

	// When specified, gets the offering status for the current period.
	Current map[string]types.OfferingStatus

	// When specified, gets the offering status for the next period.
	NextPeriod map[string]types.OfferingStatus

	// An identifier that was returned from the previous call to this operation, which
	// can be used to return the next set of items in the list.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetOfferingStatusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetOfferingStatus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetOfferingStatus{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetOfferingStatus"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetOfferingStatus(options.Region), middleware.Before); err != nil {
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

// GetOfferingStatusPaginatorOptions is the paginator options for GetOfferingStatus
type GetOfferingStatusPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetOfferingStatusPaginator is a paginator for GetOfferingStatus
type GetOfferingStatusPaginator struct {
	options   GetOfferingStatusPaginatorOptions
	client    GetOfferingStatusAPIClient
	params    *GetOfferingStatusInput
	nextToken *string
	firstPage bool
}

// NewGetOfferingStatusPaginator returns a new GetOfferingStatusPaginator
func NewGetOfferingStatusPaginator(client GetOfferingStatusAPIClient, params *GetOfferingStatusInput, optFns ...func(*GetOfferingStatusPaginatorOptions)) *GetOfferingStatusPaginator {
	if params == nil {
		params = &GetOfferingStatusInput{}
	}

	options := GetOfferingStatusPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetOfferingStatusPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetOfferingStatusPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetOfferingStatus page.
func (p *GetOfferingStatusPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetOfferingStatusOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	optFns = append([]func(*Options){
		addIsPaginatorUserAgent,
	}, optFns...)
	result, err := p.client.GetOfferingStatus(ctx, &params, optFns...)
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

// GetOfferingStatusAPIClient is a client that implements the GetOfferingStatus
// operation.
type GetOfferingStatusAPIClient interface {
	GetOfferingStatus(context.Context, *GetOfferingStatusInput, ...func(*Options)) (*GetOfferingStatusOutput, error)
}

var _ GetOfferingStatusAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opGetOfferingStatus(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetOfferingStatus",
	}
}
