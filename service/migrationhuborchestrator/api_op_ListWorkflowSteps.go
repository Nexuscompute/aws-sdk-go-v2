// Code generated by smithy-go-codegen DO NOT EDIT.

package migrationhuborchestrator

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/migrationhuborchestrator/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// List the steps in a workflow.
func (c *Client) ListWorkflowSteps(ctx context.Context, params *ListWorkflowStepsInput, optFns ...func(*Options)) (*ListWorkflowStepsOutput, error) {
	if params == nil {
		params = &ListWorkflowStepsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListWorkflowSteps", params, optFns, c.addOperationListWorkflowStepsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListWorkflowStepsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListWorkflowStepsInput struct {

	// The ID of the step group.
	//
	// This member is required.
	StepGroupId *string

	// The ID of the migration workflow.
	//
	// This member is required.
	WorkflowId *string

	// The maximum number of results that can be returned.
	MaxResults int32

	// The pagination token.
	NextToken *string

	noSmithyDocumentSerde
}

type ListWorkflowStepsOutput struct {

	// The summary of steps in a migration workflow.
	//
	// This member is required.
	WorkflowStepsSummary []types.WorkflowStepSummary

	// The pagination token.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListWorkflowStepsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListWorkflowSteps{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListWorkflowSteps{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListWorkflowSteps"); err != nil {
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
	if err = addOpListWorkflowStepsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListWorkflowSteps(options.Region), middleware.Before); err != nil {
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

// ListWorkflowStepsPaginatorOptions is the paginator options for ListWorkflowSteps
type ListWorkflowStepsPaginatorOptions struct {
	// The maximum number of results that can be returned.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListWorkflowStepsPaginator is a paginator for ListWorkflowSteps
type ListWorkflowStepsPaginator struct {
	options   ListWorkflowStepsPaginatorOptions
	client    ListWorkflowStepsAPIClient
	params    *ListWorkflowStepsInput
	nextToken *string
	firstPage bool
}

// NewListWorkflowStepsPaginator returns a new ListWorkflowStepsPaginator
func NewListWorkflowStepsPaginator(client ListWorkflowStepsAPIClient, params *ListWorkflowStepsInput, optFns ...func(*ListWorkflowStepsPaginatorOptions)) *ListWorkflowStepsPaginator {
	if params == nil {
		params = &ListWorkflowStepsInput{}
	}

	options := ListWorkflowStepsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListWorkflowStepsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListWorkflowStepsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListWorkflowSteps page.
func (p *ListWorkflowStepsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListWorkflowStepsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	optFns = append([]func(*Options){
		addIsPaginatorUserAgent,
	}, optFns...)
	result, err := p.client.ListWorkflowSteps(ctx, &params, optFns...)
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

// ListWorkflowStepsAPIClient is a client that implements the ListWorkflowSteps
// operation.
type ListWorkflowStepsAPIClient interface {
	ListWorkflowSteps(context.Context, *ListWorkflowStepsInput, ...func(*Options)) (*ListWorkflowStepsOutput, error)
}

var _ ListWorkflowStepsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListWorkflowSteps(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListWorkflowSteps",
	}
}
