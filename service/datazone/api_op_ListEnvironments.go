// Code generated by smithy-go-codegen DO NOT EDIT.

package datazone

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/datazone/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists Amazon DataZone environments.
func (c *Client) ListEnvironments(ctx context.Context, params *ListEnvironmentsInput, optFns ...func(*Options)) (*ListEnvironmentsOutput, error) {
	if params == nil {
		params = &ListEnvironmentsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListEnvironments", params, optFns, c.addOperationListEnvironmentsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListEnvironmentsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListEnvironmentsInput struct {

	// The identifier of the Amazon DataZone domain.
	//
	// This member is required.
	DomainIdentifier *string

	// The identifier of the Amazon DataZone project.
	//
	// This member is required.
	ProjectIdentifier *string

	// The identifier of the Amazon Web Services account where you want to list
	// environments.
	AwsAccountId *string

	// The Amazon Web Services region where you want to list environments.
	AwsAccountRegion *string

	// The identifier of the Amazon DataZone blueprint.
	EnvironmentBlueprintIdentifier *string

	// The identifier of the environment profile.
	EnvironmentProfileIdentifier *string

	// The maximum number of environments to return in a single call to
	// ListEnvironments . When the number of environments to be listed is greater than
	// the value of MaxResults , the response contains a NextToken value that you can
	// use in a subsequent call to ListEnvironments to list the next set of
	// environments.
	MaxResults *int32

	// The name of the environment.
	Name *string

	// When the number of environments is greater than the default value for the
	// MaxResults parameter, or if you explicitly specify a value for MaxResults that
	// is less than the number of environments, the response includes a pagination
	// token named NextToken . You can specify this NextToken value in a subsequent
	// call to ListEnvironments to list the next set of environments.
	NextToken *string

	// The provider of the environment.
	Provider *string

	// The status of the environments that you want to list.
	Status types.EnvironmentStatus

	noSmithyDocumentSerde
}

type ListEnvironmentsOutput struct {

	// The results of the ListEnvironments action.
	//
	// This member is required.
	Items []types.EnvironmentSummary

	// When the number of environments is greater than the default value for the
	// MaxResults parameter, or if you explicitly specify a value for MaxResults that
	// is less than the number of environments, the response includes a pagination
	// token named NextToken . You can specify this NextToken value in a subsequent
	// call to ListEnvironments to list the next set of environments.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListEnvironmentsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListEnvironments{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListEnvironments{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListEnvironments"); err != nil {
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
	if err = addOpListEnvironmentsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListEnvironments(options.Region), middleware.Before); err != nil {
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

// ListEnvironmentsAPIClient is a client that implements the ListEnvironments
// operation.
type ListEnvironmentsAPIClient interface {
	ListEnvironments(context.Context, *ListEnvironmentsInput, ...func(*Options)) (*ListEnvironmentsOutput, error)
}

var _ ListEnvironmentsAPIClient = (*Client)(nil)

// ListEnvironmentsPaginatorOptions is the paginator options for ListEnvironments
type ListEnvironmentsPaginatorOptions struct {
	// The maximum number of environments to return in a single call to
	// ListEnvironments . When the number of environments to be listed is greater than
	// the value of MaxResults , the response contains a NextToken value that you can
	// use in a subsequent call to ListEnvironments to list the next set of
	// environments.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListEnvironmentsPaginator is a paginator for ListEnvironments
type ListEnvironmentsPaginator struct {
	options   ListEnvironmentsPaginatorOptions
	client    ListEnvironmentsAPIClient
	params    *ListEnvironmentsInput
	nextToken *string
	firstPage bool
}

// NewListEnvironmentsPaginator returns a new ListEnvironmentsPaginator
func NewListEnvironmentsPaginator(client ListEnvironmentsAPIClient, params *ListEnvironmentsInput, optFns ...func(*ListEnvironmentsPaginatorOptions)) *ListEnvironmentsPaginator {
	if params == nil {
		params = &ListEnvironmentsInput{}
	}

	options := ListEnvironmentsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListEnvironmentsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListEnvironmentsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListEnvironments page.
func (p *ListEnvironmentsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListEnvironmentsOutput, error) {
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

	result, err := p.client.ListEnvironments(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListEnvironments(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListEnvironments",
	}
}
