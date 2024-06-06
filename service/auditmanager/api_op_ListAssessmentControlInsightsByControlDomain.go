// Code generated by smithy-go-codegen DO NOT EDIT.

package auditmanager

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/auditmanager/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the latest analytics data for controls within a specific control domain
// and a specific active assessment.
//
// Control insights are listed only if the control belongs to the control domain
// and assessment that was specified. Moreover, the control must have collected
// evidence on the lastUpdated date of controlInsightsByAssessment . If neither of
// these conditions are met, no data is listed for that control.
func (c *Client) ListAssessmentControlInsightsByControlDomain(ctx context.Context, params *ListAssessmentControlInsightsByControlDomainInput, optFns ...func(*Options)) (*ListAssessmentControlInsightsByControlDomainOutput, error) {
	if params == nil {
		params = &ListAssessmentControlInsightsByControlDomainInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAssessmentControlInsightsByControlDomain", params, optFns, c.addOperationListAssessmentControlInsightsByControlDomainMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAssessmentControlInsightsByControlDomainOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAssessmentControlInsightsByControlDomainInput struct {

	// The unique identifier for the active assessment.
	//
	// This member is required.
	AssessmentId *string

	// The unique identifier for the control domain.
	//
	// This member is required.
	ControlDomainId *string

	// Represents the maximum number of results on a page or for an API request call.
	MaxResults *int32

	// The pagination token that's used to fetch the next set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListAssessmentControlInsightsByControlDomainOutput struct {

	// The assessment control analytics data that the
	// ListAssessmentControlInsightsByControlDomain API returned.
	ControlInsightsByAssessment []types.ControlInsightsMetadataByAssessmentItem

	// The pagination token that's used to fetch the next set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAssessmentControlInsightsByControlDomainMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListAssessmentControlInsightsByControlDomain{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListAssessmentControlInsightsByControlDomain{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListAssessmentControlInsightsByControlDomain"); err != nil {
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
	if err = addOpListAssessmentControlInsightsByControlDomainValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAssessmentControlInsightsByControlDomain(options.Region), middleware.Before); err != nil {
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

// ListAssessmentControlInsightsByControlDomainAPIClient is a client that
// implements the ListAssessmentControlInsightsByControlDomain operation.
type ListAssessmentControlInsightsByControlDomainAPIClient interface {
	ListAssessmentControlInsightsByControlDomain(context.Context, *ListAssessmentControlInsightsByControlDomainInput, ...func(*Options)) (*ListAssessmentControlInsightsByControlDomainOutput, error)
}

var _ ListAssessmentControlInsightsByControlDomainAPIClient = (*Client)(nil)

// ListAssessmentControlInsightsByControlDomainPaginatorOptions is the paginator
// options for ListAssessmentControlInsightsByControlDomain
type ListAssessmentControlInsightsByControlDomainPaginatorOptions struct {
	// Represents the maximum number of results on a page or for an API request call.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAssessmentControlInsightsByControlDomainPaginator is a paginator for
// ListAssessmentControlInsightsByControlDomain
type ListAssessmentControlInsightsByControlDomainPaginator struct {
	options   ListAssessmentControlInsightsByControlDomainPaginatorOptions
	client    ListAssessmentControlInsightsByControlDomainAPIClient
	params    *ListAssessmentControlInsightsByControlDomainInput
	nextToken *string
	firstPage bool
}

// NewListAssessmentControlInsightsByControlDomainPaginator returns a new
// ListAssessmentControlInsightsByControlDomainPaginator
func NewListAssessmentControlInsightsByControlDomainPaginator(client ListAssessmentControlInsightsByControlDomainAPIClient, params *ListAssessmentControlInsightsByControlDomainInput, optFns ...func(*ListAssessmentControlInsightsByControlDomainPaginatorOptions)) *ListAssessmentControlInsightsByControlDomainPaginator {
	if params == nil {
		params = &ListAssessmentControlInsightsByControlDomainInput{}
	}

	options := ListAssessmentControlInsightsByControlDomainPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAssessmentControlInsightsByControlDomainPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAssessmentControlInsightsByControlDomainPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListAssessmentControlInsightsByControlDomain page.
func (p *ListAssessmentControlInsightsByControlDomainPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAssessmentControlInsightsByControlDomainOutput, error) {
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

	result, err := p.client.ListAssessmentControlInsightsByControlDomain(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListAssessmentControlInsightsByControlDomain(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListAssessmentControlInsightsByControlDomain",
	}
}
