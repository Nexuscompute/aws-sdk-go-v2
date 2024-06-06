// Code generated by smithy-go-codegen DO NOT EDIT.

package networkmanager

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/networkmanager/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets the network telemetry of the specified global network.
func (c *Client) GetNetworkTelemetry(ctx context.Context, params *GetNetworkTelemetryInput, optFns ...func(*Options)) (*GetNetworkTelemetryOutput, error) {
	if params == nil {
		params = &GetNetworkTelemetryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetNetworkTelemetry", params, optFns, c.addOperationGetNetworkTelemetryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetNetworkTelemetryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetNetworkTelemetryInput struct {

	// The ID of the global network.
	//
	// This member is required.
	GlobalNetworkId *string

	// The Amazon Web Services account ID.
	AccountId *string

	// The Amazon Web Services Region.
	AwsRegion *string

	// The ID of a core network.
	CoreNetworkId *string

	// The maximum number of results to return.
	MaxResults *int32

	// The token for the next page of results.
	NextToken *string

	// The ARN of the gateway.
	RegisteredGatewayArn *string

	// The ARN of the resource.
	ResourceArn *string

	// The resource type.
	//
	// The following are the supported resource types for Direct Connect:
	//
	//   - dxcon
	//
	//   - dx-gateway
	//
	//   - dx-vif
	//
	// The following are the supported resource types for Network Manager:
	//
	//   - connection
	//
	//   - device
	//
	//   - link
	//
	//   - site
	//
	// The following are the supported resource types for Amazon VPC:
	//
	//   - customer-gateway
	//
	//   - transit-gateway
	//
	//   - transit-gateway-attachment
	//
	//   - transit-gateway-connect-peer
	//
	//   - transit-gateway-route-table
	//
	//   - vpn-connection
	ResourceType *string

	noSmithyDocumentSerde
}

type GetNetworkTelemetryOutput struct {

	// The network telemetry.
	NetworkTelemetry []types.NetworkTelemetry

	// The token for the next page of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetNetworkTelemetryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetNetworkTelemetry{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetNetworkTelemetry{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetNetworkTelemetry"); err != nil {
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
	if err = addOpGetNetworkTelemetryValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetNetworkTelemetry(options.Region), middleware.Before); err != nil {
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

// GetNetworkTelemetryAPIClient is a client that implements the
// GetNetworkTelemetry operation.
type GetNetworkTelemetryAPIClient interface {
	GetNetworkTelemetry(context.Context, *GetNetworkTelemetryInput, ...func(*Options)) (*GetNetworkTelemetryOutput, error)
}

var _ GetNetworkTelemetryAPIClient = (*Client)(nil)

// GetNetworkTelemetryPaginatorOptions is the paginator options for
// GetNetworkTelemetry
type GetNetworkTelemetryPaginatorOptions struct {
	// The maximum number of results to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetNetworkTelemetryPaginator is a paginator for GetNetworkTelemetry
type GetNetworkTelemetryPaginator struct {
	options   GetNetworkTelemetryPaginatorOptions
	client    GetNetworkTelemetryAPIClient
	params    *GetNetworkTelemetryInput
	nextToken *string
	firstPage bool
}

// NewGetNetworkTelemetryPaginator returns a new GetNetworkTelemetryPaginator
func NewGetNetworkTelemetryPaginator(client GetNetworkTelemetryAPIClient, params *GetNetworkTelemetryInput, optFns ...func(*GetNetworkTelemetryPaginatorOptions)) *GetNetworkTelemetryPaginator {
	if params == nil {
		params = &GetNetworkTelemetryInput{}
	}

	options := GetNetworkTelemetryPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetNetworkTelemetryPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetNetworkTelemetryPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetNetworkTelemetry page.
func (p *GetNetworkTelemetryPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetNetworkTelemetryOutput, error) {
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

	result, err := p.client.GetNetworkTelemetry(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetNetworkTelemetry(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetNetworkTelemetry",
	}
}
