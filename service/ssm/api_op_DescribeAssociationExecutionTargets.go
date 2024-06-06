// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Views information about a specific execution of a specific association.
func (c *Client) DescribeAssociationExecutionTargets(ctx context.Context, params *DescribeAssociationExecutionTargetsInput, optFns ...func(*Options)) (*DescribeAssociationExecutionTargetsOutput, error) {
	if params == nil {
		params = &DescribeAssociationExecutionTargetsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeAssociationExecutionTargets", params, optFns, c.addOperationDescribeAssociationExecutionTargetsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeAssociationExecutionTargetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAssociationExecutionTargetsInput struct {

	// The association ID that includes the execution for which you want to view
	// details.
	//
	// This member is required.
	AssociationId *string

	// The execution ID for which you want to view details.
	//
	// This member is required.
	ExecutionId *string

	// Filters for the request. You can specify the following filters and values.
	//
	// Status (EQUAL)
	//
	// ResourceId (EQUAL)
	//
	// ResourceType (EQUAL)
	Filters []types.AssociationExecutionTargetsFilter

	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	MaxResults *int32

	// A token to start the list. Use this token to get the next set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeAssociationExecutionTargetsOutput struct {

	// Information about the execution.
	AssociationExecutionTargets []types.AssociationExecutionTarget

	// The token for the next set of items to return. Use this token to get the next
	// set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeAssociationExecutionTargetsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeAssociationExecutionTargets{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeAssociationExecutionTargets{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeAssociationExecutionTargets"); err != nil {
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
	if err = addOpDescribeAssociationExecutionTargetsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAssociationExecutionTargets(options.Region), middleware.Before); err != nil {
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

// DescribeAssociationExecutionTargetsAPIClient is a client that implements the
// DescribeAssociationExecutionTargets operation.
type DescribeAssociationExecutionTargetsAPIClient interface {
	DescribeAssociationExecutionTargets(context.Context, *DescribeAssociationExecutionTargetsInput, ...func(*Options)) (*DescribeAssociationExecutionTargetsOutput, error)
}

var _ DescribeAssociationExecutionTargetsAPIClient = (*Client)(nil)

// DescribeAssociationExecutionTargetsPaginatorOptions is the paginator options
// for DescribeAssociationExecutionTargets
type DescribeAssociationExecutionTargetsPaginatorOptions struct {
	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeAssociationExecutionTargetsPaginator is a paginator for
// DescribeAssociationExecutionTargets
type DescribeAssociationExecutionTargetsPaginator struct {
	options   DescribeAssociationExecutionTargetsPaginatorOptions
	client    DescribeAssociationExecutionTargetsAPIClient
	params    *DescribeAssociationExecutionTargetsInput
	nextToken *string
	firstPage bool
}

// NewDescribeAssociationExecutionTargetsPaginator returns a new
// DescribeAssociationExecutionTargetsPaginator
func NewDescribeAssociationExecutionTargetsPaginator(client DescribeAssociationExecutionTargetsAPIClient, params *DescribeAssociationExecutionTargetsInput, optFns ...func(*DescribeAssociationExecutionTargetsPaginatorOptions)) *DescribeAssociationExecutionTargetsPaginator {
	if params == nil {
		params = &DescribeAssociationExecutionTargetsInput{}
	}

	options := DescribeAssociationExecutionTargetsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeAssociationExecutionTargetsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeAssociationExecutionTargetsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeAssociationExecutionTargets page.
func (p *DescribeAssociationExecutionTargetsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeAssociationExecutionTargetsOutput, error) {
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

	result, err := p.client.DescribeAssociationExecutionTargets(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeAssociationExecutionTargets(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeAssociationExecutionTargets",
	}
}
