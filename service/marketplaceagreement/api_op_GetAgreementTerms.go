// Code generated by smithy-go-codegen DO NOT EDIT.

package marketplaceagreement

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/marketplaceagreement/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Obtains details about the terms in an agreement that you participated in as
// proposer or acceptor.
//
// The details include:
//
//   - TermType – The type of term, such as LegalTerm , RenewalTerm , or
//     ConfigurableUpfrontPricingTerm .
//
//   - TermID – The ID of the particular term, which is common between offer and
//     agreement.
//
//   - TermPayload – The key information contained in the term, such as the EULA
//     for LegalTerm or pricing and dimensions for various pricing terms, such as
//     ConfigurableUpfrontPricingTerm or UsageBasedPricingTerm .
//
//   - Configuration – The buyer/acceptor's selection at the time of agreement
//     creation, such as the number of units purchased for a dimension or setting the
//     EnableAutoRenew flag.
func (c *Client) GetAgreementTerms(ctx context.Context, params *GetAgreementTermsInput, optFns ...func(*Options)) (*GetAgreementTermsOutput, error) {
	if params == nil {
		params = &GetAgreementTermsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetAgreementTerms", params, optFns, c.addOperationGetAgreementTermsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetAgreementTermsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetAgreementTermsInput struct {

	// The unique identifier of the agreement.
	//
	// This member is required.
	AgreementId *string

	// The maximum number of agreements to return in the response.
	MaxResults *int32

	// A token to specify where to start pagination
	NextToken *string

	noSmithyDocumentSerde
}

type GetAgreementTermsOutput struct {

	// A subset of terms proposed by the proposer that have been accepted by the
	// acceptor as part of the agreement creation.
	AcceptedTerms []types.AcceptedTerm

	// A token to specify where to start pagination
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetAgreementTermsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpGetAgreementTerms{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpGetAgreementTerms{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetAgreementTerms"); err != nil {
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
	if err = addOpGetAgreementTermsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetAgreementTerms(options.Region), middleware.Before); err != nil {
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

// GetAgreementTermsPaginatorOptions is the paginator options for GetAgreementTerms
type GetAgreementTermsPaginatorOptions struct {
	// The maximum number of agreements to return in the response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetAgreementTermsPaginator is a paginator for GetAgreementTerms
type GetAgreementTermsPaginator struct {
	options   GetAgreementTermsPaginatorOptions
	client    GetAgreementTermsAPIClient
	params    *GetAgreementTermsInput
	nextToken *string
	firstPage bool
}

// NewGetAgreementTermsPaginator returns a new GetAgreementTermsPaginator
func NewGetAgreementTermsPaginator(client GetAgreementTermsAPIClient, params *GetAgreementTermsInput, optFns ...func(*GetAgreementTermsPaginatorOptions)) *GetAgreementTermsPaginator {
	if params == nil {
		params = &GetAgreementTermsInput{}
	}

	options := GetAgreementTermsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetAgreementTermsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetAgreementTermsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetAgreementTerms page.
func (p *GetAgreementTermsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetAgreementTermsOutput, error) {
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

	optFns = append([]func(*Options){
		addIsPaginatorUserAgent,
	}, optFns...)
	result, err := p.client.GetAgreementTerms(ctx, &params, optFns...)
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

// GetAgreementTermsAPIClient is a client that implements the GetAgreementTerms
// operation.
type GetAgreementTermsAPIClient interface {
	GetAgreementTerms(context.Context, *GetAgreementTermsInput, ...func(*Options)) (*GetAgreementTermsOutput, error)
}

var _ GetAgreementTermsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opGetAgreementTerms(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetAgreementTerms",
	}
}
