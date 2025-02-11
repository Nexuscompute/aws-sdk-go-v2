// Code generated by smithy-go-codegen DO NOT EDIT.

package amplify

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/amplify/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the branches of an Amplify app.
func (c *Client) ListBranches(ctx context.Context, params *ListBranchesInput, optFns ...func(*Options)) (*ListBranchesOutput, error) {
	if params == nil {
		params = &ListBranchesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListBranches", params, optFns, c.addOperationListBranchesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListBranchesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request structure for the list branches request.
type ListBranchesInput struct {

	// The unique ID for an Amplify app.
	//
	// This member is required.
	AppId *string

	//  The maximum number of records to list in a single response.
	MaxResults int32

	// A pagination token. Set to null to start listing branches from the start. If a
	// non-null pagination token is returned in a result, pass its value in here to
	// list more branches.
	NextToken *string

	noSmithyDocumentSerde
}

// The result structure for the list branches request.
type ListBranchesOutput struct {

	//  A list of branches for an Amplify app.
	//
	// This member is required.
	Branches []types.Branch

	//  A pagination token. If a non-null pagination token is returned in a result,
	// pass its value in another request to retrieve more entries.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListBranchesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListBranches{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListBranches{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListBranches"); err != nil {
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
	if err = addOpListBranchesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListBranches(options.Region), middleware.Before); err != nil {
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

// ListBranchesPaginatorOptions is the paginator options for ListBranches
type ListBranchesPaginatorOptions struct {
	//  The maximum number of records to list in a single response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListBranchesPaginator is a paginator for ListBranches
type ListBranchesPaginator struct {
	options   ListBranchesPaginatorOptions
	client    ListBranchesAPIClient
	params    *ListBranchesInput
	nextToken *string
	firstPage bool
}

// NewListBranchesPaginator returns a new ListBranchesPaginator
func NewListBranchesPaginator(client ListBranchesAPIClient, params *ListBranchesInput, optFns ...func(*ListBranchesPaginatorOptions)) *ListBranchesPaginator {
	if params == nil {
		params = &ListBranchesInput{}
	}

	options := ListBranchesPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListBranchesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListBranchesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListBranches page.
func (p *ListBranchesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListBranchesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	optFns = append([]func(*Options){
		addIsPaginatorUserAgent,
	}, optFns...)
	result, err := p.client.ListBranches(ctx, &params, optFns...)
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

// ListBranchesAPIClient is a client that implements the ListBranches operation.
type ListBranchesAPIClient interface {
	ListBranches(context.Context, *ListBranchesInput, ...func(*Options)) (*ListBranchesOutput, error)
}

var _ ListBranchesAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListBranches(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListBranches",
	}
}
