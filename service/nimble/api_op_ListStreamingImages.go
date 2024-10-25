// Code generated by smithy-go-codegen DO NOT EDIT.

package nimble

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/nimble/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// List the streaming image resources available to this studio.
//
// This list will contain both images provided by Amazon Web Services, as well as
// streaming images that you have created in your studio.
//
// Deprecated: AWS has deprecated this service. It is no longer available for use.
func (c *Client) ListStreamingImages(ctx context.Context, params *ListStreamingImagesInput, optFns ...func(*Options)) (*ListStreamingImagesOutput, error) {
	if params == nil {
		params = &ListStreamingImagesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListStreamingImages", params, optFns, c.addOperationListStreamingImagesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListStreamingImagesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListStreamingImagesInput struct {

	// The studio ID.
	//
	// This member is required.
	//
	// Deprecated: AWS has deprecated this service. It is no longer available for use.
	StudioId *string

	// The token for the next set of results, or null if there are no more results.
	//
	// Deprecated: AWS has deprecated this service. It is no longer available for use.
	NextToken *string

	// Filter this request to streaming images with the given owner
	//
	// Deprecated: AWS has deprecated this service. It is no longer available for use.
	Owner *string

	noSmithyDocumentSerde
}

type ListStreamingImagesOutput struct {

	// The token for the next set of results, or null if there are no more results.
	//
	// Deprecated: AWS has deprecated this service. It is no longer available for use.
	NextToken *string

	// A collection of streaming images.
	//
	// Deprecated: AWS has deprecated this service. It is no longer available for use.
	StreamingImages []types.StreamingImage

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListStreamingImagesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListStreamingImages{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListStreamingImages{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListStreamingImages"); err != nil {
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
	if err = addOpListStreamingImagesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListStreamingImages(options.Region), middleware.Before); err != nil {
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

// ListStreamingImagesPaginatorOptions is the paginator options for
// ListStreamingImages
type ListStreamingImagesPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListStreamingImagesPaginator is a paginator for ListStreamingImages
type ListStreamingImagesPaginator struct {
	options   ListStreamingImagesPaginatorOptions
	client    ListStreamingImagesAPIClient
	params    *ListStreamingImagesInput
	nextToken *string
	firstPage bool
}

// NewListStreamingImagesPaginator returns a new ListStreamingImagesPaginator
func NewListStreamingImagesPaginator(client ListStreamingImagesAPIClient, params *ListStreamingImagesInput, optFns ...func(*ListStreamingImagesPaginatorOptions)) *ListStreamingImagesPaginator {
	if params == nil {
		params = &ListStreamingImagesInput{}
	}

	options := ListStreamingImagesPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListStreamingImagesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListStreamingImagesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListStreamingImages page.
func (p *ListStreamingImagesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListStreamingImagesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	optFns = append([]func(*Options){
		addIsPaginatorUserAgent,
	}, optFns...)
	result, err := p.client.ListStreamingImages(ctx, &params, optFns...)
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

// ListStreamingImagesAPIClient is a client that implements the
// ListStreamingImages operation.
type ListStreamingImagesAPIClient interface {
	ListStreamingImages(context.Context, *ListStreamingImagesInput, ...func(*Options)) (*ListStreamingImagesOutput, error)
}

var _ ListStreamingImagesAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListStreamingImages(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListStreamingImages",
	}
}
