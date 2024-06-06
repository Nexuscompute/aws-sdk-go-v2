// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// The related resources of an Audit finding. The following resources can be
// returned from calling this API:
//
//   - DEVICE_CERTIFICATE
//
//   - CA_CERTIFICATE
//
//   - IOT_POLICY
//
//   - COGNITO_IDENTITY_POOL
//
//   - CLIENT_ID
//
//   - ACCOUNT_SETTINGS
//
//   - ROLE_ALIAS
//
//   - IAM_ROLE
//
//   - ISSUER_CERTIFICATE
//
// This API is similar to DescribeAuditFinding's [RelatedResources] but provides pagination and is
// not limited to 10 resources. When calling [DescribeAuditFinding]for the intermediate CA revoked for
// active device certificates check, RelatedResources will not be populated. You
// must use this API, ListRelatedResourcesForAuditFinding, to list the
// certificates.
//
// [RelatedResources]: https://docs.aws.amazon.com/iot/latest/apireference/API_DescribeAuditFinding.html
// [DescribeAuditFinding]: https://docs.aws.amazon.com/iot/latest/apireference/API_DescribeAuditFinding.html
func (c *Client) ListRelatedResourcesForAuditFinding(ctx context.Context, params *ListRelatedResourcesForAuditFindingInput, optFns ...func(*Options)) (*ListRelatedResourcesForAuditFindingOutput, error) {
	if params == nil {
		params = &ListRelatedResourcesForAuditFindingInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListRelatedResourcesForAuditFinding", params, optFns, c.addOperationListRelatedResourcesForAuditFindingMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListRelatedResourcesForAuditFindingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListRelatedResourcesForAuditFindingInput struct {

	// The finding Id.
	//
	// This member is required.
	FindingId *string

	// The maximum number of results to return at one time.
	MaxResults *int32

	// A token that can be used to retrieve the next set of results, or null if there
	// are no additional results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListRelatedResourcesForAuditFindingOutput struct {

	// A token that can be used to retrieve the next set of results, or null for the
	// first API call.
	NextToken *string

	// The related resources.
	RelatedResources []types.RelatedResource

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListRelatedResourcesForAuditFindingMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListRelatedResourcesForAuditFinding{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListRelatedResourcesForAuditFinding{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListRelatedResourcesForAuditFinding"); err != nil {
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
	if err = addOpListRelatedResourcesForAuditFindingValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListRelatedResourcesForAuditFinding(options.Region), middleware.Before); err != nil {
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

// ListRelatedResourcesForAuditFindingAPIClient is a client that implements the
// ListRelatedResourcesForAuditFinding operation.
type ListRelatedResourcesForAuditFindingAPIClient interface {
	ListRelatedResourcesForAuditFinding(context.Context, *ListRelatedResourcesForAuditFindingInput, ...func(*Options)) (*ListRelatedResourcesForAuditFindingOutput, error)
}

var _ ListRelatedResourcesForAuditFindingAPIClient = (*Client)(nil)

// ListRelatedResourcesForAuditFindingPaginatorOptions is the paginator options
// for ListRelatedResourcesForAuditFinding
type ListRelatedResourcesForAuditFindingPaginatorOptions struct {
	// The maximum number of results to return at one time.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListRelatedResourcesForAuditFindingPaginator is a paginator for
// ListRelatedResourcesForAuditFinding
type ListRelatedResourcesForAuditFindingPaginator struct {
	options   ListRelatedResourcesForAuditFindingPaginatorOptions
	client    ListRelatedResourcesForAuditFindingAPIClient
	params    *ListRelatedResourcesForAuditFindingInput
	nextToken *string
	firstPage bool
}

// NewListRelatedResourcesForAuditFindingPaginator returns a new
// ListRelatedResourcesForAuditFindingPaginator
func NewListRelatedResourcesForAuditFindingPaginator(client ListRelatedResourcesForAuditFindingAPIClient, params *ListRelatedResourcesForAuditFindingInput, optFns ...func(*ListRelatedResourcesForAuditFindingPaginatorOptions)) *ListRelatedResourcesForAuditFindingPaginator {
	if params == nil {
		params = &ListRelatedResourcesForAuditFindingInput{}
	}

	options := ListRelatedResourcesForAuditFindingPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListRelatedResourcesForAuditFindingPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListRelatedResourcesForAuditFindingPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListRelatedResourcesForAuditFinding page.
func (p *ListRelatedResourcesForAuditFindingPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListRelatedResourcesForAuditFindingOutput, error) {
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

	result, err := p.client.ListRelatedResourcesForAuditFinding(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListRelatedResourcesForAuditFinding(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListRelatedResourcesForAuditFinding",
	}
}
