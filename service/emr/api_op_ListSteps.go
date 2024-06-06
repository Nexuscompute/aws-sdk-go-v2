// Code generated by smithy-go-codegen DO NOT EDIT.

package emr

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/emr/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Provides a list of steps for the cluster in reverse order unless you specify
// stepIds with the request or filter by StepStates . You can specify a maximum of
// 10 stepIDs . The CLI automatically paginates results to return a list greater
// than 50 steps. To return more than 50 steps using the CLI, specify a Marker ,
// which is a pagination token that indicates the next set of steps to retrieve.
func (c *Client) ListSteps(ctx context.Context, params *ListStepsInput, optFns ...func(*Options)) (*ListStepsOutput, error) {
	if params == nil {
		params = &ListStepsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListSteps", params, optFns, c.addOperationListStepsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListStepsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// This input determines which steps to list.
type ListStepsInput struct {

	// The identifier of the cluster for which to list the steps.
	//
	// This member is required.
	ClusterId *string

	// The maximum number of steps that a single ListSteps action returns is 50. To
	// return a longer list of steps, use multiple ListSteps actions along with the
	// Marker parameter, which is a pagination token that indicates the next set of
	// results to retrieve.
	Marker *string

	// The filter to limit the step list based on the identifier of the steps. You can
	// specify a maximum of ten Step IDs. The character constraint applies to the
	// overall length of the array.
	StepIds []string

	// The filter to limit the step list based on certain states.
	StepStates []types.StepState

	noSmithyDocumentSerde
}

// This output contains the list of steps returned in reverse order. This means
// that the last step is the first element in the list.
type ListStepsOutput struct {

	// The maximum number of steps that a single ListSteps action returns is 50. To
	// return a longer list of steps, use multiple ListSteps actions along with the
	// Marker parameter, which is a pagination token that indicates the next set of
	// results to retrieve.
	Marker *string

	// The filtered list of steps for the cluster.
	Steps []types.StepSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListStepsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListSteps{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListSteps{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListSteps"); err != nil {
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
	if err = addOpListStepsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListSteps(options.Region), middleware.Before); err != nil {
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

// ListStepsAPIClient is a client that implements the ListSteps operation.
type ListStepsAPIClient interface {
	ListSteps(context.Context, *ListStepsInput, ...func(*Options)) (*ListStepsOutput, error)
}

var _ ListStepsAPIClient = (*Client)(nil)

// ListStepsPaginatorOptions is the paginator options for ListSteps
type ListStepsPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListStepsPaginator is a paginator for ListSteps
type ListStepsPaginator struct {
	options   ListStepsPaginatorOptions
	client    ListStepsAPIClient
	params    *ListStepsInput
	nextToken *string
	firstPage bool
}

// NewListStepsPaginator returns a new ListStepsPaginator
func NewListStepsPaginator(client ListStepsAPIClient, params *ListStepsInput, optFns ...func(*ListStepsPaginatorOptions)) *ListStepsPaginator {
	if params == nil {
		params = &ListStepsInput{}
	}

	options := ListStepsPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListStepsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListStepsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListSteps page.
func (p *ListStepsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListStepsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	result, err := p.client.ListSteps(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.Marker

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListSteps(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListSteps",
	}
}
