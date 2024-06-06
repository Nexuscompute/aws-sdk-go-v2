// Code generated by smithy-go-codegen DO NOT EDIT.

package iotsitewise

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/iotsitewise/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a gateway, which is a virtual or edge device that delivers industrial
// data streams from local servers to IoT SiteWise. For more information, see [Ingesting data using a gateway]in
// the IoT SiteWise User Guide.
//
// [Ingesting data using a gateway]: https://docs.aws.amazon.com/iot-sitewise/latest/userguide/gateway-connector.html
func (c *Client) CreateGateway(ctx context.Context, params *CreateGatewayInput, optFns ...func(*Options)) (*CreateGatewayOutput, error) {
	if params == nil {
		params = &CreateGatewayInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateGateway", params, optFns, c.addOperationCreateGatewayMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateGatewayOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateGatewayInput struct {

	// A unique, friendly name for the gateway.
	//
	// This member is required.
	GatewayName *string

	// The gateway's platform. You can only specify one platform in a gateway.
	//
	// This member is required.
	GatewayPlatform *types.GatewayPlatform

	// A list of key-value pairs that contain metadata for the gateway. For more
	// information, see [Tagging your IoT SiteWise resources]in the IoT SiteWise User Guide.
	//
	// [Tagging your IoT SiteWise resources]: https://docs.aws.amazon.com/iot-sitewise/latest/userguide/tag-resources.html
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateGatewayOutput struct {

	// The [ARN] of the gateway, which has the following format.
	//
	//     arn:${Partition}:iotsitewise:${Region}:${Account}:gateway/${GatewayId}
	//
	// [ARN]: https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html
	//
	// This member is required.
	GatewayArn *string

	// The ID of the gateway device. You can use this ID when you call other IoT
	// SiteWise API operations.
	//
	// This member is required.
	GatewayId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateGatewayMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateGateway{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateGateway{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateGateway"); err != nil {
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
	if err = addEndpointPrefix_opCreateGatewayMiddleware(stack); err != nil {
		return err
	}
	if err = addOpCreateGatewayValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateGateway(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opCreateGatewayMiddleware struct {
}

func (*endpointPrefix_opCreateGatewayMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opCreateGatewayMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleFinalize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "api." + req.URL.Host

	return next.HandleFinalize(ctx, in)
}
func addEndpointPrefix_opCreateGatewayMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opCreateGatewayMiddleware{}, "ResolveEndpointV2", middleware.After)
}

func newServiceMetadataMiddleware_opCreateGateway(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateGateway",
	}
}
