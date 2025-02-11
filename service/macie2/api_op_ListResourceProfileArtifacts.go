// Code generated by smithy-go-codegen DO NOT EDIT.

package macie2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/macie2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves information about objects that Amazon Macie selected from an S3
// bucket for automated sensitive data discovery.
func (c *Client) ListResourceProfileArtifacts(ctx context.Context, params *ListResourceProfileArtifactsInput, optFns ...func(*Options)) (*ListResourceProfileArtifactsOutput, error) {
	if params == nil {
		params = &ListResourceProfileArtifactsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListResourceProfileArtifacts", params, optFns, c.addOperationListResourceProfileArtifactsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListResourceProfileArtifactsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListResourceProfileArtifactsInput struct {

	// The Amazon Resource Name (ARN) of the S3 bucket that the request applies to.
	//
	// This member is required.
	ResourceArn *string

	// The nextToken string that specifies which page of results to return in a
	// paginated response.
	NextToken *string

	noSmithyDocumentSerde
}

type ListResourceProfileArtifactsOutput struct {

	// An array of objects, one for each of 1-100 S3 objects that Amazon Macie
	// selected for analysis.
	//
	// If Macie has analyzed more than 100 objects in the bucket, Macie populates the
	// array based on the value for the ResourceProfileArtifact.sensitive field for an
	// object: true (sensitive), followed by false (not sensitive). Macie then
	// populates any remaining items in the array with information about objects where
	// the value for the ResourceProfileArtifact.classificationResultStatus field is
	// SKIPPED.
	Artifacts []types.ResourceProfileArtifact

	// The string to use in a subsequent request to get the next page of results in a
	// paginated response. This value is null if there are no additional pages.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListResourceProfileArtifactsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListResourceProfileArtifacts{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListResourceProfileArtifacts{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListResourceProfileArtifacts"); err != nil {
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
	if err = addOpListResourceProfileArtifactsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListResourceProfileArtifacts(options.Region), middleware.Before); err != nil {
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

// ListResourceProfileArtifactsPaginatorOptions is the paginator options for
// ListResourceProfileArtifacts
type ListResourceProfileArtifactsPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListResourceProfileArtifactsPaginator is a paginator for
// ListResourceProfileArtifacts
type ListResourceProfileArtifactsPaginator struct {
	options   ListResourceProfileArtifactsPaginatorOptions
	client    ListResourceProfileArtifactsAPIClient
	params    *ListResourceProfileArtifactsInput
	nextToken *string
	firstPage bool
}

// NewListResourceProfileArtifactsPaginator returns a new
// ListResourceProfileArtifactsPaginator
func NewListResourceProfileArtifactsPaginator(client ListResourceProfileArtifactsAPIClient, params *ListResourceProfileArtifactsInput, optFns ...func(*ListResourceProfileArtifactsPaginatorOptions)) *ListResourceProfileArtifactsPaginator {
	if params == nil {
		params = &ListResourceProfileArtifactsInput{}
	}

	options := ListResourceProfileArtifactsPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListResourceProfileArtifactsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListResourceProfileArtifactsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListResourceProfileArtifacts page.
func (p *ListResourceProfileArtifactsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListResourceProfileArtifactsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	optFns = append([]func(*Options){
		addIsPaginatorUserAgent,
	}, optFns...)
	result, err := p.client.ListResourceProfileArtifacts(ctx, &params, optFns...)
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

// ListResourceProfileArtifactsAPIClient is a client that implements the
// ListResourceProfileArtifacts operation.
type ListResourceProfileArtifactsAPIClient interface {
	ListResourceProfileArtifacts(context.Context, *ListResourceProfileArtifactsInput, ...func(*Options)) (*ListResourceProfileArtifactsOutput, error)
}

var _ ListResourceProfileArtifactsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListResourceProfileArtifacts(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListResourceProfileArtifacts",
	}
}
