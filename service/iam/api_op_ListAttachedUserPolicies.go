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

// Lists all managed policies that are attached to the specified IAM user.
//
// An IAM user can also have inline policies embedded with it. To list the inline
// policies for a user, use ListUserPolicies. For information about policies, see [Managed policies and inline policies] in the IAM User
// Guide.
//
// You can paginate the results using the MaxItems and Marker parameters. You can
// use the PathPrefix parameter to limit the list of policies to only those
// matching the specified path prefix. If there are no policies attached to the
// specified group (or none that match the specified path prefix), the operation
// returns an empty list.
//
// [Managed policies and inline policies]: https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-vs-inline.html
func (c *Client) ListAttachedUserPolicies(ctx context.Context, params *ListAttachedUserPoliciesInput, optFns ...func(*Options)) (*ListAttachedUserPoliciesOutput, error) {
	if params == nil {
		params = &ListAttachedUserPoliciesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAttachedUserPolicies", params, optFns, c.addOperationListAttachedUserPoliciesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAttachedUserPoliciesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAttachedUserPoliciesInput struct {

	// The name (friendly name, not ARN) of the user to list attached policies for.
	//
	// This parameter allows (through its [regex pattern]) a string of characters consisting of upper
	// and lowercase alphanumeric characters with no spaces. You can also include any
	// of the following characters: _+=,.@-
	//
	// [regex pattern]: http://wikipedia.org/wiki/regex
	//
	// This member is required.
	UserName *string

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

	// The path prefix for filtering the results. This parameter is optional. If it is
	// not included, it defaults to a slash (/), listing all policies.
	//
	// This parameter allows (through its [regex pattern]) a string of characters consisting of
	// either a forward slash (/) by itself or a string that must begin and end with
	// forward slashes. In addition, it can contain any ASCII character from the ! (
	// \u0021 ) through the DEL character ( \u007F ), including most punctuation
	// characters, digits, and upper and lowercased letters.
	//
	// [regex pattern]: http://wikipedia.org/wiki/regex
	PathPrefix *string

	noSmithyDocumentSerde
}

// Contains the response to a successful ListAttachedUserPolicies request.
type ListAttachedUserPoliciesOutput struct {

	// A list of the attached policies.
	AttachedPolicies []types.AttachedPolicy

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

func (c *Client) addOperationListAttachedUserPoliciesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpListAttachedUserPolicies{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpListAttachedUserPolicies{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListAttachedUserPolicies"); err != nil {
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
	if err = addOpListAttachedUserPoliciesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAttachedUserPolicies(options.Region), middleware.Before); err != nil {
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

// ListAttachedUserPoliciesAPIClient is a client that implements the
// ListAttachedUserPolicies operation.
type ListAttachedUserPoliciesAPIClient interface {
	ListAttachedUserPolicies(context.Context, *ListAttachedUserPoliciesInput, ...func(*Options)) (*ListAttachedUserPoliciesOutput, error)
}

var _ ListAttachedUserPoliciesAPIClient = (*Client)(nil)

// ListAttachedUserPoliciesPaginatorOptions is the paginator options for
// ListAttachedUserPolicies
type ListAttachedUserPoliciesPaginatorOptions struct {
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

// ListAttachedUserPoliciesPaginator is a paginator for ListAttachedUserPolicies
type ListAttachedUserPoliciesPaginator struct {
	options   ListAttachedUserPoliciesPaginatorOptions
	client    ListAttachedUserPoliciesAPIClient
	params    *ListAttachedUserPoliciesInput
	nextToken *string
	firstPage bool
}

// NewListAttachedUserPoliciesPaginator returns a new
// ListAttachedUserPoliciesPaginator
func NewListAttachedUserPoliciesPaginator(client ListAttachedUserPoliciesAPIClient, params *ListAttachedUserPoliciesInput, optFns ...func(*ListAttachedUserPoliciesPaginatorOptions)) *ListAttachedUserPoliciesPaginator {
	if params == nil {
		params = &ListAttachedUserPoliciesInput{}
	}

	options := ListAttachedUserPoliciesPaginatorOptions{}
	if params.MaxItems != nil {
		options.Limit = *params.MaxItems
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAttachedUserPoliciesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAttachedUserPoliciesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListAttachedUserPolicies page.
func (p *ListAttachedUserPoliciesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAttachedUserPoliciesOutput, error) {
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

	result, err := p.client.ListAttachedUserPolicies(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListAttachedUserPolicies(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListAttachedUserPolicies",
	}
}
