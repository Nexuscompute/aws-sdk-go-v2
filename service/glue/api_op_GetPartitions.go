// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Retrieves information about the partitions in a table.
func (c *Client) GetPartitions(ctx context.Context, params *GetPartitionsInput, optFns ...func(*Options)) (*GetPartitionsOutput, error) {
	if params == nil {
		params = &GetPartitionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetPartitions", params, optFns, c.addOperationGetPartitionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetPartitionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetPartitionsInput struct {

	// The name of the catalog database where the partitions reside.
	//
	// This member is required.
	DatabaseName *string

	// The name of the partitions' table.
	//
	// This member is required.
	TableName *string

	// The ID of the Data Catalog where the partitions in question reside. If none is
	// provided, the Amazon Web Services account ID is used by default.
	CatalogId *string

	// When true, specifies not returning the partition column schema. Useful when you
	// are interested only in other partition attributes such as partition values or
	// location. This approach avoids the problem of a large response by not returning
	// duplicate data.
	ExcludeColumnSchema *bool

	// An expression that filters the partitions to be returned.
	//
	// The expression uses SQL syntax similar to the SQL WHERE filter clause. The SQL
	// statement parser [JSQLParser]parses the expression.
	//
	// Operators: The following are the operators that you can use in the Expression
	// API call:
	//
	// = Checks whether the values of the two operands are equal; if yes, then the
	// condition becomes true.
	//
	// Example: Assume 'variable a' holds 10 and 'variable b' holds 20.
	//
	// (a = b) is not true.
	//
	// < > Checks whether the values of two operands are equal; if the values are not
	// equal, then the condition becomes true.
	//
	// Example: (a < > b) is true.
	//
	// > Checks whether the value of the left operand is greater than the value of the
	// right operand; if yes, then the condition becomes true.
	//
	// Example: (a > b) is not true.
	//
	// < Checks whether the value of the left operand is less than the value of the
	// right operand; if yes, then the condition becomes true.
	//
	// Example: (a < b) is true.
	//
	// >= Checks whether the value of the left operand is greater than or equal to the
	// value of the right operand; if yes, then the condition becomes true.
	//
	// Example: (a >= b) is not true.
	//
	// <= Checks whether the value of the left operand is less than or equal to the
	// value of the right operand; if yes, then the condition becomes true.
	//
	// Example: (a <= b) is true.
	//
	// AND, OR, IN, BETWEEN, LIKE, NOT, IS NULL Logical operators.
	//
	// Supported Partition Key Types: The following are the supported partition keys.
	//
	//   - string
	//
	//   - date
	//
	//   - timestamp
	//
	//   - int
	//
	//   - bigint
	//
	//   - long
	//
	//   - tinyint
	//
	//   - smallint
	//
	//   - decimal
	//
	// If an type is encountered that is not valid, an exception is thrown.
	//
	// The following list shows the valid operators on each type. When you define a
	// crawler, the partitionKey type is created as a STRING , to be compatible with
	// the catalog partitions.
	//
	// Sample API Call:
	//
	// [JSQLParser]: http://jsqlparser.sourceforge.net/home.php
	Expression *string

	// The maximum number of partitions to return in a single response.
	MaxResults *int32

	// A continuation token, if this is not the first call to retrieve these
	// partitions.
	NextToken *string

	// The time as of when to read the partition contents. If not set, the most recent
	// transaction commit time will be used. Cannot be specified along with
	// TransactionId .
	QueryAsOfTime *time.Time

	// The segment of the table's partitions to scan in this request.
	Segment *types.Segment

	// The transaction ID at which to read the partition contents.
	TransactionId *string

	noSmithyDocumentSerde
}

type GetPartitionsOutput struct {

	// A continuation token, if the returned list of partitions does not include the
	// last one.
	NextToken *string

	// A list of requested partitions.
	Partitions []types.Partition

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetPartitionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetPartitions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetPartitions{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetPartitions"); err != nil {
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
	if err = addOpGetPartitionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetPartitions(options.Region), middleware.Before); err != nil {
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

// GetPartitionsAPIClient is a client that implements the GetPartitions operation.
type GetPartitionsAPIClient interface {
	GetPartitions(context.Context, *GetPartitionsInput, ...func(*Options)) (*GetPartitionsOutput, error)
}

var _ GetPartitionsAPIClient = (*Client)(nil)

// GetPartitionsPaginatorOptions is the paginator options for GetPartitions
type GetPartitionsPaginatorOptions struct {
	// The maximum number of partitions to return in a single response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetPartitionsPaginator is a paginator for GetPartitions
type GetPartitionsPaginator struct {
	options   GetPartitionsPaginatorOptions
	client    GetPartitionsAPIClient
	params    *GetPartitionsInput
	nextToken *string
	firstPage bool
}

// NewGetPartitionsPaginator returns a new GetPartitionsPaginator
func NewGetPartitionsPaginator(client GetPartitionsAPIClient, params *GetPartitionsInput, optFns ...func(*GetPartitionsPaginatorOptions)) *GetPartitionsPaginator {
	if params == nil {
		params = &GetPartitionsInput{}
	}

	options := GetPartitionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetPartitionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetPartitionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetPartitions page.
func (p *GetPartitionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetPartitionsOutput, error) {
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

	result, err := p.client.GetPartitions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetPartitions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetPartitions",
	}
}
