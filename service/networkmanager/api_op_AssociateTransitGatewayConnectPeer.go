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

// Associates a transit gateway Connect peer with a device, and optionally, with a
// link. If you specify a link, it must be associated with the specified device.
//
// You can only associate transit gateway Connect peers that have been created on
// a transit gateway that's registered in your global network.
//
// You cannot associate a transit gateway Connect peer with more than one device
// and link.
func (c *Client) AssociateTransitGatewayConnectPeer(ctx context.Context, params *AssociateTransitGatewayConnectPeerInput, optFns ...func(*Options)) (*AssociateTransitGatewayConnectPeerOutput, error) {
	if params == nil {
		params = &AssociateTransitGatewayConnectPeerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AssociateTransitGatewayConnectPeer", params, optFns, c.addOperationAssociateTransitGatewayConnectPeerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AssociateTransitGatewayConnectPeerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateTransitGatewayConnectPeerInput struct {

	// The ID of the device.
	//
	// This member is required.
	DeviceId *string

	// The ID of the global network.
	//
	// This member is required.
	GlobalNetworkId *string

	// The Amazon Resource Name (ARN) of the Connect peer.
	//
	// This member is required.
	TransitGatewayConnectPeerArn *string

	// The ID of the link.
	LinkId *string

	noSmithyDocumentSerde
}

type AssociateTransitGatewayConnectPeerOutput struct {

	// The transit gateway Connect peer association.
	TransitGatewayConnectPeerAssociation *types.TransitGatewayConnectPeerAssociation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAssociateTransitGatewayConnectPeerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpAssociateTransitGatewayConnectPeer{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpAssociateTransitGatewayConnectPeer{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "AssociateTransitGatewayConnectPeer"); err != nil {
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
	if err = addOpAssociateTransitGatewayConnectPeerValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateTransitGatewayConnectPeer(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAssociateTransitGatewayConnectPeer(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "AssociateTransitGatewayConnectPeer",
	}
}
