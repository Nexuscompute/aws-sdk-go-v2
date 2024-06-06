// Code generated by smithy-go-codegen DO NOT EDIT.

package codedeploy

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/codedeploy/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

//	The newer BatchGetDeploymentTargets should be used instead because it works
//
// with all compute types. ListDeploymentInstances throws an exception if it is
// used with a compute platform other than EC2/On-premises or Lambda.
//
// Lists the instance for a deployment associated with the user or Amazon Web
// Services account.
//
// Deprecated: This operation is deprecated, use ListDeploymentTargets instead.
func (c *Client) ListDeploymentInstances(ctx context.Context, params *ListDeploymentInstancesInput, optFns ...func(*Options)) (*ListDeploymentInstancesOutput, error) {
	if params == nil {
		params = &ListDeploymentInstancesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDeploymentInstances", params, optFns, c.addOperationListDeploymentInstancesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDeploymentInstancesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a ListDeploymentInstances operation.
type ListDeploymentInstancesInput struct {

	//  The unique ID of a deployment.
	//
	// This member is required.
	DeploymentId *string

	// A subset of instances to list by status:
	//
	//   - Pending : Include those instances with pending deployments.
	//
	//   - InProgress : Include those instances where deployments are still in progress.
	//
	//   - Succeeded : Include those instances with successful deployments.
	//
	//   - Failed : Include those instances with failed deployments.
	//
	//   - Skipped : Include those instances with skipped deployments.
	//
	//   - Unknown : Include those instances with deployments in an unknown state.
	InstanceStatusFilter []types.InstanceStatus

	// The set of instances in a blue/green deployment, either those in the original
	// environment ("BLUE") or those in the replacement environment ("GREEN"), for
	// which you want to view instance information.
	InstanceTypeFilter []types.InstanceType

	// An identifier returned from the previous list deployment instances call. It can
	// be used to return the next set of deployment instances in the list.
	NextToken *string

	noSmithyDocumentSerde
}

// Represents the output of a ListDeploymentInstances operation.
type ListDeploymentInstancesOutput struct {

	// A list of instance IDs.
	InstancesList []string

	// If a large amount of information is returned, an identifier is also returned.
	// It can be used in a subsequent list deployment instances call to return the next
	// set of deployment instances in the list.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListDeploymentInstancesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListDeploymentInstances{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListDeploymentInstances{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListDeploymentInstances"); err != nil {
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
	if err = addOpListDeploymentInstancesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListDeploymentInstances(options.Region), middleware.Before); err != nil {
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

// ListDeploymentInstancesAPIClient is a client that implements the
// ListDeploymentInstances operation.
type ListDeploymentInstancesAPIClient interface {
	ListDeploymentInstances(context.Context, *ListDeploymentInstancesInput, ...func(*Options)) (*ListDeploymentInstancesOutput, error)
}

var _ ListDeploymentInstancesAPIClient = (*Client)(nil)

// ListDeploymentInstancesPaginatorOptions is the paginator options for
// ListDeploymentInstances
type ListDeploymentInstancesPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListDeploymentInstancesPaginator is a paginator for ListDeploymentInstances
type ListDeploymentInstancesPaginator struct {
	options   ListDeploymentInstancesPaginatorOptions
	client    ListDeploymentInstancesAPIClient
	params    *ListDeploymentInstancesInput
	nextToken *string
	firstPage bool
}

// NewListDeploymentInstancesPaginator returns a new
// ListDeploymentInstancesPaginator
func NewListDeploymentInstancesPaginator(client ListDeploymentInstancesAPIClient, params *ListDeploymentInstancesInput, optFns ...func(*ListDeploymentInstancesPaginatorOptions)) *ListDeploymentInstancesPaginator {
	if params == nil {
		params = &ListDeploymentInstancesInput{}
	}

	options := ListDeploymentInstancesPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListDeploymentInstancesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListDeploymentInstancesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListDeploymentInstances page.
func (p *ListDeploymentInstancesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListDeploymentInstancesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	result, err := p.client.ListDeploymentInstances(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListDeploymentInstances(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListDeploymentInstances",
	}
}
