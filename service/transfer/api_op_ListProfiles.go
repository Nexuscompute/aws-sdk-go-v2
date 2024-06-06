// Code generated by smithy-go-codegen DO NOT EDIT.

package transfer

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/transfer/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of the profiles for your system. If you want to limit the
// results to a certain number, supply a value for the MaxResults parameter. If
// you ran the command previously and received a value for NextToken , you can
// supply that value to continue listing profiles from where you left off.
func (c *Client) ListProfiles(ctx context.Context, params *ListProfilesInput, optFns ...func(*Options)) (*ListProfilesOutput, error) {
	if params == nil {
		params = &ListProfilesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListProfiles", params, optFns, c.addOperationListProfilesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListProfilesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListProfilesInput struct {

	// The maximum number of profiles to return.
	MaxResults *int32

	// When there are additional results that were not returned, a NextToken parameter
	// is returned. You can use that value for a subsequent call to ListProfiles to
	// continue listing results.
	NextToken *string

	// Indicates whether to list only LOCAL type profiles or only PARTNER type
	// profiles. If not supplied in the request, the command lists all types of
	// profiles.
	ProfileType types.ProfileType

	noSmithyDocumentSerde
}

type ListProfilesOutput struct {

	// Returns an array, where each item contains the details of a profile.
	//
	// This member is required.
	Profiles []types.ListedProfile

	// Returns a token that you can use to call ListProfiles again and receive
	// additional results, if there are any.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListProfilesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListProfiles{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListProfiles{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListProfiles"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListProfiles(options.Region), middleware.Before); err != nil {
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

// ListProfilesAPIClient is a client that implements the ListProfiles operation.
type ListProfilesAPIClient interface {
	ListProfiles(context.Context, *ListProfilesInput, ...func(*Options)) (*ListProfilesOutput, error)
}

var _ ListProfilesAPIClient = (*Client)(nil)

// ListProfilesPaginatorOptions is the paginator options for ListProfiles
type ListProfilesPaginatorOptions struct {
	// The maximum number of profiles to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListProfilesPaginator is a paginator for ListProfiles
type ListProfilesPaginator struct {
	options   ListProfilesPaginatorOptions
	client    ListProfilesAPIClient
	params    *ListProfilesInput
	nextToken *string
	firstPage bool
}

// NewListProfilesPaginator returns a new ListProfilesPaginator
func NewListProfilesPaginator(client ListProfilesAPIClient, params *ListProfilesInput, optFns ...func(*ListProfilesPaginatorOptions)) *ListProfilesPaginator {
	if params == nil {
		params = &ListProfilesInput{}
	}

	options := ListProfilesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListProfilesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListProfilesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListProfiles page.
func (p *ListProfilesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListProfilesOutput, error) {
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

	result, err := p.client.ListProfiles(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListProfiles(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListProfiles",
	}
}
