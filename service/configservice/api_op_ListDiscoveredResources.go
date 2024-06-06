// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Accepts a resource type and returns a list of resource identifiers for the
// resources of that type. A resource identifier includes the resource type, ID,
// and (if available) the custom resource name. The results consist of resources
// that Config has discovered, including those that Config is not currently
// recording. You can narrow the results to include only resources that have
// specific resource IDs or a resource name.
//
// You can specify either resource IDs or a resource name, but not both, in the
// same request.
//
// The response is paginated. By default, Config lists 100 resource identifiers on
// each page. You can customize this number with the limit parameter. The response
// includes a nextToken string. To get the next page of results, run the request
// again and specify the string for the nextToken parameter.
func (c *Client) ListDiscoveredResources(ctx context.Context, params *ListDiscoveredResourcesInput, optFns ...func(*Options)) (*ListDiscoveredResourcesOutput, error) {
	if params == nil {
		params = &ListDiscoveredResourcesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDiscoveredResources", params, optFns, c.addOperationListDiscoveredResourcesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDiscoveredResourcesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDiscoveredResourcesInput struct {

	// The type of resources that you want Config to list in the response.
	//
	// This member is required.
	ResourceType types.ResourceType

	// Specifies whether Config includes deleted resources in the results. By default,
	// deleted resources are not included.
	IncludeDeletedResources bool

	// The maximum number of resource identifiers returned on each page. The default
	// is 100. You cannot specify a number greater than 100. If you specify 0, Config
	// uses the default.
	Limit int32

	// The nextToken string returned on a previous page that you use to get the next
	// page of results in a paginated response.
	NextToken *string

	// The IDs of only those resources that you want Config to list in the response.
	// If you do not specify this parameter, Config lists all resources of the
	// specified type that it has discovered. You can list a minimum of 1 resourceID
	// and a maximum of 20 resourceIds.
	ResourceIds []string

	// The custom name of only those resources that you want Config to list in the
	// response. If you do not specify this parameter, Config lists all resources of
	// the specified type that it has discovered.
	ResourceName *string

	noSmithyDocumentSerde
}

type ListDiscoveredResourcesOutput struct {

	// The string that you use in a subsequent request to get the next page of results
	// in a paginated response.
	NextToken *string

	// The details that identify a resource that is discovered by Config, including
	// the resource type, ID, and (if available) the custom resource name.
	ResourceIdentifiers []types.ResourceIdentifier

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListDiscoveredResourcesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListDiscoveredResources{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListDiscoveredResources{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListDiscoveredResources"); err != nil {
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
	if err = addOpListDiscoveredResourcesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListDiscoveredResources(options.Region), middleware.Before); err != nil {
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

// ListDiscoveredResourcesAPIClient is a client that implements the
// ListDiscoveredResources operation.
type ListDiscoveredResourcesAPIClient interface {
	ListDiscoveredResources(context.Context, *ListDiscoveredResourcesInput, ...func(*Options)) (*ListDiscoveredResourcesOutput, error)
}

var _ ListDiscoveredResourcesAPIClient = (*Client)(nil)

// ListDiscoveredResourcesPaginatorOptions is the paginator options for
// ListDiscoveredResources
type ListDiscoveredResourcesPaginatorOptions struct {
	// The maximum number of resource identifiers returned on each page. The default
	// is 100. You cannot specify a number greater than 100. If you specify 0, Config
	// uses the default.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListDiscoveredResourcesPaginator is a paginator for ListDiscoveredResources
type ListDiscoveredResourcesPaginator struct {
	options   ListDiscoveredResourcesPaginatorOptions
	client    ListDiscoveredResourcesAPIClient
	params    *ListDiscoveredResourcesInput
	nextToken *string
	firstPage bool
}

// NewListDiscoveredResourcesPaginator returns a new
// ListDiscoveredResourcesPaginator
func NewListDiscoveredResourcesPaginator(client ListDiscoveredResourcesAPIClient, params *ListDiscoveredResourcesInput, optFns ...func(*ListDiscoveredResourcesPaginatorOptions)) *ListDiscoveredResourcesPaginator {
	if params == nil {
		params = &ListDiscoveredResourcesInput{}
	}

	options := ListDiscoveredResourcesPaginatorOptions{}
	if params.Limit != 0 {
		options.Limit = params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListDiscoveredResourcesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListDiscoveredResourcesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListDiscoveredResources page.
func (p *ListDiscoveredResourcesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListDiscoveredResourcesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.Limit = p.options.Limit

	result, err := p.client.ListDiscoveredResources(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListDiscoveredResources(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListDiscoveredResources",
	}
}
