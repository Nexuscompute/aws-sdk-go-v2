// Code generated by smithy-go-codegen DO NOT EDIT.

package batch

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/batch/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes a list of job definitions. You can specify a status (such as ACTIVE )
// to only return job definitions that match that status.
func (c *Client) DescribeJobDefinitions(ctx context.Context, params *DescribeJobDefinitionsInput, optFns ...func(*Options)) (*DescribeJobDefinitionsOutput, error) {
	if params == nil {
		params = &DescribeJobDefinitionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeJobDefinitions", params, optFns, c.addOperationDescribeJobDefinitionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeJobDefinitionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for DescribeJobDefinitions .
type DescribeJobDefinitionsInput struct {

	// The name of the job definition to describe.
	JobDefinitionName *string

	// A list of up to 100 job definitions. Each entry in the list can either be an
	// ARN in the format
	// arn:aws:batch:${Region}:${Account}:job-definition/${JobDefinitionName}:${Revision}
	// or a short version using the form ${JobDefinitionName}:${Revision} . This
	// parameter can't be used with other parameters.
	JobDefinitions []string

	// The maximum number of results returned by DescribeJobDefinitions in paginated
	// output. When this parameter is used, DescribeJobDefinitions only returns
	// maxResults results in a single page and a nextToken response element. The
	// remaining results of the initial request can be seen by sending another
	// DescribeJobDefinitions request with the returned nextToken value. This value
	// can be between 1 and 100. If this parameter isn't used, then
	// DescribeJobDefinitions returns up to 100 results and a nextToken value if
	// applicable.
	MaxResults *int32

	// The nextToken value returned from a previous paginated DescribeJobDefinitions
	// request where maxResults was used and the results exceeded the value of that
	// parameter. Pagination continues from the end of the previous results that
	// returned the nextToken value. This value is null when there are no more results
	// to return.
	//
	// Treat this token as an opaque identifier that's only used to retrieve the next
	// items in a list and not for other programmatic purposes.
	NextToken *string

	// The status used to filter job definitions.
	Status *string

	noSmithyDocumentSerde
}

type DescribeJobDefinitionsOutput struct {

	// The list of job definitions.
	JobDefinitions []types.JobDefinition

	// The nextToken value to include in a future DescribeJobDefinitions request. When
	// the results of a DescribeJobDefinitions request exceed maxResults , this value
	// can be used to retrieve the next page of results. This value is null when there
	// are no more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeJobDefinitionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeJobDefinitions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeJobDefinitions{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeJobDefinitions"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeJobDefinitions(options.Region), middleware.Before); err != nil {
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

// DescribeJobDefinitionsAPIClient is a client that implements the
// DescribeJobDefinitions operation.
type DescribeJobDefinitionsAPIClient interface {
	DescribeJobDefinitions(context.Context, *DescribeJobDefinitionsInput, ...func(*Options)) (*DescribeJobDefinitionsOutput, error)
}

var _ DescribeJobDefinitionsAPIClient = (*Client)(nil)

// DescribeJobDefinitionsPaginatorOptions is the paginator options for
// DescribeJobDefinitions
type DescribeJobDefinitionsPaginatorOptions struct {
	// The maximum number of results returned by DescribeJobDefinitions in paginated
	// output. When this parameter is used, DescribeJobDefinitions only returns
	// maxResults results in a single page and a nextToken response element. The
	// remaining results of the initial request can be seen by sending another
	// DescribeJobDefinitions request with the returned nextToken value. This value
	// can be between 1 and 100. If this parameter isn't used, then
	// DescribeJobDefinitions returns up to 100 results and a nextToken value if
	// applicable.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeJobDefinitionsPaginator is a paginator for DescribeJobDefinitions
type DescribeJobDefinitionsPaginator struct {
	options   DescribeJobDefinitionsPaginatorOptions
	client    DescribeJobDefinitionsAPIClient
	params    *DescribeJobDefinitionsInput
	nextToken *string
	firstPage bool
}

// NewDescribeJobDefinitionsPaginator returns a new DescribeJobDefinitionsPaginator
func NewDescribeJobDefinitionsPaginator(client DescribeJobDefinitionsAPIClient, params *DescribeJobDefinitionsInput, optFns ...func(*DescribeJobDefinitionsPaginatorOptions)) *DescribeJobDefinitionsPaginator {
	if params == nil {
		params = &DescribeJobDefinitionsInput{}
	}

	options := DescribeJobDefinitionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeJobDefinitionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeJobDefinitionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeJobDefinitions page.
func (p *DescribeJobDefinitionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeJobDefinitionsOutput, error) {
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

	result, err := p.client.DescribeJobDefinitions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeJobDefinitions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeJobDefinitions",
	}
}
