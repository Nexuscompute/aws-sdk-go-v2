// Code generated by smithy-go-codegen DO NOT EDIT.

package redshiftserverless

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/redshiftserverless/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists custom domain associations for Amazon Redshift Serverless.
func (c *Client) ListCustomDomainAssociations(ctx context.Context, params *ListCustomDomainAssociationsInput, optFns ...func(*Options)) (*ListCustomDomainAssociationsOutput, error) {
	if params == nil {
		params = &ListCustomDomainAssociationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListCustomDomainAssociations", params, optFns, c.addOperationListCustomDomainAssociationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListCustomDomainAssociationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListCustomDomainAssociationsInput struct {

	// The custom domain name’s certificate Amazon resource name (ARN).
	CustomDomainCertificateArn *string

	// The custom domain name associated with the workgroup.
	CustomDomainName *string

	// An optional parameter that specifies the maximum number of results to return.
	// You can use nextToken to display the next page of results.
	MaxResults *int32

	// When nextToken is returned, there are more results available. The value of
	// nextToken is a unique pagination token for each page. Make the call again using
	// the returned token to retrieve the next page.
	NextToken *string

	noSmithyDocumentSerde
}

type ListCustomDomainAssociationsOutput struct {

	// A list of Association objects.
	Associations []types.Association

	// When nextToken is returned, there are more results available. The value of
	// nextToken is a unique pagination token for each page. Make the call again using
	// the returned token to retrieve the next page.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListCustomDomainAssociationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListCustomDomainAssociations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListCustomDomainAssociations{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListCustomDomainAssociations"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListCustomDomainAssociations(options.Region), middleware.Before); err != nil {
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

// ListCustomDomainAssociationsAPIClient is a client that implements the
// ListCustomDomainAssociations operation.
type ListCustomDomainAssociationsAPIClient interface {
	ListCustomDomainAssociations(context.Context, *ListCustomDomainAssociationsInput, ...func(*Options)) (*ListCustomDomainAssociationsOutput, error)
}

var _ ListCustomDomainAssociationsAPIClient = (*Client)(nil)

// ListCustomDomainAssociationsPaginatorOptions is the paginator options for
// ListCustomDomainAssociations
type ListCustomDomainAssociationsPaginatorOptions struct {
	// An optional parameter that specifies the maximum number of results to return.
	// You can use nextToken to display the next page of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListCustomDomainAssociationsPaginator is a paginator for
// ListCustomDomainAssociations
type ListCustomDomainAssociationsPaginator struct {
	options   ListCustomDomainAssociationsPaginatorOptions
	client    ListCustomDomainAssociationsAPIClient
	params    *ListCustomDomainAssociationsInput
	nextToken *string
	firstPage bool
}

// NewListCustomDomainAssociationsPaginator returns a new
// ListCustomDomainAssociationsPaginator
func NewListCustomDomainAssociationsPaginator(client ListCustomDomainAssociationsAPIClient, params *ListCustomDomainAssociationsInput, optFns ...func(*ListCustomDomainAssociationsPaginatorOptions)) *ListCustomDomainAssociationsPaginator {
	if params == nil {
		params = &ListCustomDomainAssociationsInput{}
	}

	options := ListCustomDomainAssociationsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListCustomDomainAssociationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListCustomDomainAssociationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListCustomDomainAssociations page.
func (p *ListCustomDomainAssociationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListCustomDomainAssociationsOutput, error) {
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

	result, err := p.client.ListCustomDomainAssociations(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListCustomDomainAssociations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListCustomDomainAssociations",
	}
}
