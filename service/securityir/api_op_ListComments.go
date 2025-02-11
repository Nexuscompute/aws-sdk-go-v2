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

// Grants permissions to list and view comments for a designated case.
func (c *Client) ListComments(ctx context.Context, params *ListCommentsInput, optFns ...func(*Options)) (*ListCommentsOutput, error) {
	if params == nil {
		params = &ListCommentsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListComments", params, optFns, c.addOperationListCommentsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListCommentsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListCommentsInput struct {

	// Required element for ListComments to designate the case to query.
	//
	// This member is required.
	CaseId *string

	// Optional element for ListComments to limit the number of responses.
	MaxResults *int32

	// Optional element.
	NextToken *string

	noSmithyDocumentSerde
}

type ListCommentsOutput struct {

	// Response element for ListComments providing the body, commentID, createDate,
	// creator, lastUpdatedBy and lastUpdatedDate for each response.
	Items []types.ListCommentsItem

	// Optional request elements.
	NextToken *string

	// Response element for ListComments identifying the number of responses.
	Total *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListCommentsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListComments{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListComments{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListComments"); err != nil {
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
	if err = addOpListCommentsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListComments(options.Region), middleware.Before); err != nil {
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

// ListCommentsPaginatorOptions is the paginator options for ListComments
type ListCommentsPaginatorOptions struct {
	// Optional element for ListComments to limit the number of responses.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListCommentsPaginator is a paginator for ListComments
type ListCommentsPaginator struct {
	options   ListCommentsPaginatorOptions
	client    ListCommentsAPIClient
	params    *ListCommentsInput
	nextToken *string
	firstPage bool
}

// NewListCommentsPaginator returns a new ListCommentsPaginator
func NewListCommentsPaginator(client ListCommentsAPIClient, params *ListCommentsInput, optFns ...func(*ListCommentsPaginatorOptions)) *ListCommentsPaginator {
	if params == nil {
		params = &ListCommentsInput{}
	}

	options := ListCommentsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListCommentsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListCommentsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListComments page.
func (p *ListCommentsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListCommentsOutput, error) {
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
	result, err := p.client.ListComments(ctx, &params, optFns...)
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

// ListCommentsAPIClient is a client that implements the ListComments operation.
type ListCommentsAPIClient interface {
	ListComments(context.Context, *ListCommentsInput, ...func(*Options)) (*ListCommentsOutput, error)
}

var _ ListCommentsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListComments(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListComments",
	}
}
