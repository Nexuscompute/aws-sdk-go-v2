// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/gamelift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

//	This operation has been expanded to use with the Amazon GameLift containers
//
// feature, which is currently in public preview.
//
// Retrieves the resource capacity settings for one or more fleets. For a
// container fleet, this operation also returns counts for replica container
// groups.
//
// With multi-location fleets, this operation retrieves data for the fleet's home
// Region only. To retrieve capacity for remote locations, see DescribeFleetLocationCapacity.
//
// This operation can be used in the following ways:
//
//   - To get capacity data for one or more specific fleets, provide a list of
//     fleet IDs or fleet ARNs.
//
//   - To get capacity data for all fleets, do not provide a fleet identifier.
//
// When requesting multiple fleets, use the pagination parameters to retrieve
// results as a set of sequential pages.
//
// If successful, a FleetCapacity object is returned for each requested fleet ID.
// Each FleetCapacity object includes a Location property, which is set to the
// fleet's home Region. Capacity values are returned only for fleets that currently
// exist.
//
// Some API operations may limit the number of fleet IDs that are allowed in one
// request. If a request exceeds this limit, the request fails and the error
// message includes the maximum allowed.
//
// # Learn more
//
// [Setting up Amazon GameLift fleets]
//
// [GameLift metrics for fleets]
//
// [Setting up Amazon GameLift fleets]: https://docs.aws.amazon.com/gamelift/latest/developerguide/fleets-intro.html
// [GameLift metrics for fleets]: https://docs.aws.amazon.com/gamelift/latest/developerguide/monitoring-cloudwatch.html#gamelift-metrics-fleet
func (c *Client) DescribeFleetCapacity(ctx context.Context, params *DescribeFleetCapacityInput, optFns ...func(*Options)) (*DescribeFleetCapacityOutput, error) {
	if params == nil {
		params = &DescribeFleetCapacityInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeFleetCapacity", params, optFns, c.addOperationDescribeFleetCapacityMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeFleetCapacityOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeFleetCapacityInput struct {

	// A unique identifier for the fleet to retrieve capacity information for. You can
	// use either the fleet ID or ARN value. Leave this parameter empty to retrieve
	// capacity information for all fleets.
	FleetIds []string

	// The maximum number of results to return. Use this parameter with NextToken to
	// get results as a set of sequential pages. This parameter is ignored when the
	// request specifies one or a list of fleet IDs.
	Limit *int32

	// A token that indicates the start of the next sequential page of results. Use
	// the token that is returned with a previous call to this operation. To start at
	// the beginning of the result set, do not specify a value. This parameter is
	// ignored when the request specifies one or a list of fleet IDs.
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeFleetCapacityOutput struct {

	// A collection of objects that contains capacity information for each requested
	// fleet ID. Capacity objects are returned only for fleets that currently exist.
	// Changes in desired instance value can take up to 1 minute to be reflected.
	FleetCapacity []types.FleetCapacity

	// A token that indicates where to resume retrieving results on the next call to
	// this operation. If no token is returned, these results represent the end of the
	// list.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeFleetCapacityMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeFleetCapacity{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeFleetCapacity{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeFleetCapacity"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeFleetCapacity(options.Region), middleware.Before); err != nil {
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

// DescribeFleetCapacityAPIClient is a client that implements the
// DescribeFleetCapacity operation.
type DescribeFleetCapacityAPIClient interface {
	DescribeFleetCapacity(context.Context, *DescribeFleetCapacityInput, ...func(*Options)) (*DescribeFleetCapacityOutput, error)
}

var _ DescribeFleetCapacityAPIClient = (*Client)(nil)

// DescribeFleetCapacityPaginatorOptions is the paginator options for
// DescribeFleetCapacity
type DescribeFleetCapacityPaginatorOptions struct {
	// The maximum number of results to return. Use this parameter with NextToken to
	// get results as a set of sequential pages. This parameter is ignored when the
	// request specifies one or a list of fleet IDs.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeFleetCapacityPaginator is a paginator for DescribeFleetCapacity
type DescribeFleetCapacityPaginator struct {
	options   DescribeFleetCapacityPaginatorOptions
	client    DescribeFleetCapacityAPIClient
	params    *DescribeFleetCapacityInput
	nextToken *string
	firstPage bool
}

// NewDescribeFleetCapacityPaginator returns a new DescribeFleetCapacityPaginator
func NewDescribeFleetCapacityPaginator(client DescribeFleetCapacityAPIClient, params *DescribeFleetCapacityInput, optFns ...func(*DescribeFleetCapacityPaginatorOptions)) *DescribeFleetCapacityPaginator {
	if params == nil {
		params = &DescribeFleetCapacityInput{}
	}

	options := DescribeFleetCapacityPaginatorOptions{}
	if params.Limit != nil {
		options.Limit = *params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeFleetCapacityPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeFleetCapacityPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeFleetCapacity page.
func (p *DescribeFleetCapacityPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeFleetCapacityOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.Limit = limit

	result, err := p.client.DescribeFleetCapacity(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeFleetCapacity(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeFleetCapacity",
	}
}
