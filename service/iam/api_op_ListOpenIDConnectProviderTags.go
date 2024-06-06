// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the tags that are attached to the specified OpenID Connect
// (OIDC)-compatible identity provider. The returned list of tags is sorted by tag
// key. For more information, see [About web identity federation].
//
// For more information about tagging, see [Tagging IAM resources] in the IAM User Guide.
//
// [Tagging IAM resources]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_tags.html
// [About web identity federation]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_providers_oidc.html
func (c *Client) ListOpenIDConnectProviderTags(ctx context.Context, params *ListOpenIDConnectProviderTagsInput, optFns ...func(*Options)) (*ListOpenIDConnectProviderTagsOutput, error) {
	if params == nil {
		params = &ListOpenIDConnectProviderTagsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListOpenIDConnectProviderTags", params, optFns, c.addOperationListOpenIDConnectProviderTagsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListOpenIDConnectProviderTagsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListOpenIDConnectProviderTagsInput struct {

	// The ARN of the OpenID Connect (OIDC) identity provider whose tags you want to
	// see.
	//
	// This parameter allows (through its [regex pattern]) a string of characters consisting of upper
	// and lowercase alphanumeric characters with no spaces. You can also include any
	// of the following characters: _+=,.@-
	//
	// [regex pattern]: http://wikipedia.org/wiki/regex
	//
	// This member is required.
	OpenIDConnectProviderArn *string

	// Use this parameter only when paginating results and only after you receive a
	// response indicating that the results are truncated. Set it to the value of the
	// Marker element in the response that you received to indicate where the next call
	// should start.
	Marker *string

	// Use this only when paginating results to indicate the maximum number of items
	// you want in the response. If additional items exist beyond the maximum you
	// specify, the IsTruncated response element is true .
	//
	// If you do not include this parameter, the number of items defaults to 100. Note
	// that IAM might return fewer results, even when there are more results available.
	// In that case, the IsTruncated response element returns true , and Marker
	// contains a value to include in the subsequent call that tells the service where
	// to continue from.
	MaxItems *int32

	noSmithyDocumentSerde
}

type ListOpenIDConnectProviderTagsOutput struct {

	// The list of tags that are currently attached to the OpenID Connect (OIDC)
	// identity provider. Each tag consists of a key name and an associated value. If
	// no tags are attached to the specified resource, the response contains an empty
	// list.
	//
	// This member is required.
	Tags []types.Tag

	// A flag that indicates whether there are more items to return. If your results
	// were truncated, you can make a subsequent pagination request using the Marker
	// request parameter to retrieve more items. Note that IAM might return fewer than
	// the MaxItems number of results even when there are more results available. We
	// recommend that you check IsTruncated after every call to ensure that you
	// receive all your results.
	IsTruncated bool

	// When IsTruncated is true , this element is present and contains the value to use
	// for the Marker parameter in a subsequent pagination request.
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListOpenIDConnectProviderTagsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpListOpenIDConnectProviderTags{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpListOpenIDConnectProviderTags{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListOpenIDConnectProviderTags"); err != nil {
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
	if err = addOpListOpenIDConnectProviderTagsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListOpenIDConnectProviderTags(options.Region), middleware.Before); err != nil {
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

// ListOpenIDConnectProviderTagsAPIClient is a client that implements the
// ListOpenIDConnectProviderTags operation.
type ListOpenIDConnectProviderTagsAPIClient interface {
	ListOpenIDConnectProviderTags(context.Context, *ListOpenIDConnectProviderTagsInput, ...func(*Options)) (*ListOpenIDConnectProviderTagsOutput, error)
}

var _ ListOpenIDConnectProviderTagsAPIClient = (*Client)(nil)

// ListOpenIDConnectProviderTagsPaginatorOptions is the paginator options for
// ListOpenIDConnectProviderTags
type ListOpenIDConnectProviderTagsPaginatorOptions struct {
	// Use this only when paginating results to indicate the maximum number of items
	// you want in the response. If additional items exist beyond the maximum you
	// specify, the IsTruncated response element is true .
	//
	// If you do not include this parameter, the number of items defaults to 100. Note
	// that IAM might return fewer results, even when there are more results available.
	// In that case, the IsTruncated response element returns true , and Marker
	// contains a value to include in the subsequent call that tells the service where
	// to continue from.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListOpenIDConnectProviderTagsPaginator is a paginator for
// ListOpenIDConnectProviderTags
type ListOpenIDConnectProviderTagsPaginator struct {
	options   ListOpenIDConnectProviderTagsPaginatorOptions
	client    ListOpenIDConnectProviderTagsAPIClient
	params    *ListOpenIDConnectProviderTagsInput
	nextToken *string
	firstPage bool
}

// NewListOpenIDConnectProviderTagsPaginator returns a new
// ListOpenIDConnectProviderTagsPaginator
func NewListOpenIDConnectProviderTagsPaginator(client ListOpenIDConnectProviderTagsAPIClient, params *ListOpenIDConnectProviderTagsInput, optFns ...func(*ListOpenIDConnectProviderTagsPaginatorOptions)) *ListOpenIDConnectProviderTagsPaginator {
	if params == nil {
		params = &ListOpenIDConnectProviderTagsInput{}
	}

	options := ListOpenIDConnectProviderTagsPaginatorOptions{}
	if params.MaxItems != nil {
		options.Limit = *params.MaxItems
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListOpenIDConnectProviderTagsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListOpenIDConnectProviderTagsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListOpenIDConnectProviderTags page.
func (p *ListOpenIDConnectProviderTagsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListOpenIDConnectProviderTagsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxItems = limit

	result, err := p.client.ListOpenIDConnectProviderTags(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListOpenIDConnectProviderTags(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListOpenIDConnectProviderTags",
	}
}
