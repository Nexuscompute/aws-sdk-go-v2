// Code generated by smithy-go-codegen DO NOT EDIT.

package securityir

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/securityir/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Grants permissions to view the aidt log for edits made to a designated case.
func (c *Client) ListCaseEdits(ctx context.Context, params *ListCaseEditsInput, optFns ...func(*Options)) (*ListCaseEditsOutput, error) {
	if params == nil {
		params = &ListCaseEditsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListCaseEdits", params, optFns, c.addOperationListCaseEditsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListCaseEditsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListCaseEditsInput struct {

	// Required element used with ListCaseEdits to identify the case to query.
	//
	// This member is required.
	CaseId *string

	// Optional element to identify how many results to obtain. There is a maximum
	// value of 25.
	MaxResults *int32

	// Optional element for a customer provided token.
	NextToken *string

	noSmithyDocumentSerde
}

type ListCaseEditsOutput struct {

	// Response element for ListCaseEdits that includes the action, eventtimestamp,
	// message, and principal for the response.
	Items []types.CaseEditItem

	// Optional element.
	NextToken *string

	// Response element for ListCaseEdits that identifies the total number of edits.
	Total *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListCaseEditsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListCaseEdits{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListCaseEdits{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListCaseEdits"); err != nil {
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
	if err = addOpListCaseEditsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListCaseEdits(options.Region), middleware.Before); err != nil {
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

// ListCaseEditsPaginatorOptions is the paginator options for ListCaseEdits
type ListCaseEditsPaginatorOptions struct {
	// Optional element to identify how many results to obtain. There is a maximum
	// value of 25.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListCaseEditsPaginator is a paginator for ListCaseEdits
type ListCaseEditsPaginator struct {
	options   ListCaseEditsPaginatorOptions
	client    ListCaseEditsAPIClient
	params    *ListCaseEditsInput
	nextToken *string
	firstPage bool
}

// NewListCaseEditsPaginator returns a new ListCaseEditsPaginator
func NewListCaseEditsPaginator(client ListCaseEditsAPIClient, params *ListCaseEditsInput, optFns ...func(*ListCaseEditsPaginatorOptions)) *ListCaseEditsPaginator {
	if params == nil {
		params = &ListCaseEditsInput{}
	}

	options := ListCaseEditsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListCaseEditsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListCaseEditsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListCaseEdits page.
func (p *ListCaseEditsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListCaseEditsOutput, error) {
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
	result, err := p.client.ListCaseEdits(ctx, &params, optFns...)
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

// ListCaseEditsAPIClient is a client that implements the ListCaseEdits operation.
type ListCaseEditsAPIClient interface {
	ListCaseEdits(context.Context, *ListCaseEditsInput, ...func(*Options)) (*ListCaseEditsOutput, error)
}

var _ ListCaseEditsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListCaseEdits(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListCaseEdits",
	}
}
