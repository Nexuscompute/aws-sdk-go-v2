// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the detailed parameter list for a particular DB cluster parameter group.
//
// For more information on Amazon Aurora, see [What is Amazon Aurora?] in the Amazon Aurora User Guide.
//
// For more information on Multi-AZ DB clusters, see [Multi-AZ DB cluster deployments] in the Amazon RDS User Guide.
//
// [What is Amazon Aurora?]: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/CHAP_AuroraOverview.html
// [Multi-AZ DB cluster deployments]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/multi-az-db-clusters-concepts.html
func (c *Client) DescribeDBClusterParameters(ctx context.Context, params *DescribeDBClusterParametersInput, optFns ...func(*Options)) (*DescribeDBClusterParametersOutput, error) {
	if params == nil {
		params = &DescribeDBClusterParametersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeDBClusterParameters", params, optFns, c.addOperationDescribeDBClusterParametersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeDBClusterParametersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeDBClusterParametersInput struct {

	// The name of a specific DB cluster parameter group to return parameter details
	// for.
	//
	// Constraints:
	//
	//   - If supplied, must match the name of an existing DBClusterParameterGroup.
	//
	// This member is required.
	DBClusterParameterGroupName *string

	// A filter that specifies one or more DB cluster parameters to describe.
	//
	// The only supported filter is parameter-name . The results list only includes
	// information about the DB cluster parameters with these names.
	Filters []types.Filter

	// An optional pagination token provided by a previous DescribeDBClusterParameters
	// request. If this parameter is specified, the response includes only records
	// beyond the marker, up to the value specified by MaxRecords .
	Marker *string

	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so you can retrieve the remaining results.
	//
	// Default: 100
	//
	// Constraints: Minimum 20, maximum 100.
	MaxRecords *int32

	// A specific source to return parameters for.
	//
	// Valid Values:
	//
	//   - engine-default
	//
	//   - system
	//
	//   - user
	Source *string

	noSmithyDocumentSerde
}

// Provides details about a DB cluster parameter group including the parameters in
// the DB cluster parameter group.
type DescribeDBClusterParametersOutput struct {

	// An optional pagination token provided by a previous DescribeDBClusterParameters
	// request. If this parameter is specified, the response includes only records
	// beyond the marker, up to the value specified by MaxRecords .
	Marker *string

	// Provides a list of parameters for the DB cluster parameter group.
	Parameters []types.Parameter

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeDBClusterParametersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeDBClusterParameters{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeDBClusterParameters{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeDBClusterParameters"); err != nil {
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
	if err = addOpDescribeDBClusterParametersValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDBClusterParameters(options.Region), middleware.Before); err != nil {
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

// DescribeDBClusterParametersPaginatorOptions is the paginator options for
// DescribeDBClusterParameters
type DescribeDBClusterParametersPaginatorOptions struct {
	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so you can retrieve the remaining results.
	//
	// Default: 100
	//
	// Constraints: Minimum 20, maximum 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeDBClusterParametersPaginator is a paginator for
// DescribeDBClusterParameters
type DescribeDBClusterParametersPaginator struct {
	options   DescribeDBClusterParametersPaginatorOptions
	client    DescribeDBClusterParametersAPIClient
	params    *DescribeDBClusterParametersInput
	nextToken *string
	firstPage bool
}

// NewDescribeDBClusterParametersPaginator returns a new
// DescribeDBClusterParametersPaginator
func NewDescribeDBClusterParametersPaginator(client DescribeDBClusterParametersAPIClient, params *DescribeDBClusterParametersInput, optFns ...func(*DescribeDBClusterParametersPaginatorOptions)) *DescribeDBClusterParametersPaginator {
	if params == nil {
		params = &DescribeDBClusterParametersInput{}
	}

	options := DescribeDBClusterParametersPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeDBClusterParametersPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeDBClusterParametersPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeDBClusterParameters page.
func (p *DescribeDBClusterParametersPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeDBClusterParametersOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxRecords = limit

	optFns = append([]func(*Options){
		addIsPaginatorUserAgent,
	}, optFns...)
	result, err := p.client.DescribeDBClusterParameters(ctx, &params, optFns...)
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

// DescribeDBClusterParametersAPIClient is a client that implements the
// DescribeDBClusterParameters operation.
type DescribeDBClusterParametersAPIClient interface {
	DescribeDBClusterParameters(context.Context, *DescribeDBClusterParametersInput, ...func(*Options)) (*DescribeDBClusterParametersOutput, error)
}

var _ DescribeDBClusterParametersAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opDescribeDBClusterParameters(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeDBClusterParameters",
	}
}
