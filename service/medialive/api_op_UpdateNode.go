// Code generated by smithy-go-codegen DO NOT EDIT.

package medialive

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/medialive/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Change the settings for a Node.
func (c *Client) UpdateNode(ctx context.Context, params *UpdateNodeInput, optFns ...func(*Options)) (*UpdateNodeOutput, error) {
	if params == nil {
		params = &UpdateNodeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateNode", params, optFns, c.addOperationUpdateNodeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateNodeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to update the node.
type UpdateNodeInput struct {

	// The ID of the cluster
	//
	// This member is required.
	ClusterId *string

	// The ID of the node.
	//
	// This member is required.
	NodeId *string

	// Include this parameter only if you want to change the current name of the Node.
	// Specify a name that is unique in the Cluster. You can't change the name. Names
	// are case-sensitive.
	Name *string

	// The initial role of the Node in the Cluster. ACTIVE means the Node is available
	// for encoding. BACKUP means the Node is a redundant Node and might get used if an
	// ACTIVE Node fails.
	Role types.NodeRole

	noSmithyDocumentSerde
}

// Placeholder documentation for UpdateNodeResponse
type UpdateNodeOutput struct {

	// The ARN of the Node. It is automatically assigned when the Node is created.
	Arn *string

	// An array of IDs. Each ID is one ChannelPlacementGroup that is associated with
	// this Node. Empty if the Node is not yet associated with any groups.
	ChannelPlacementGroups []string

	// The ID of the Cluster that the Node belongs to.
	ClusterId *string

	// The current connection state of the Node.
	ConnectionState types.NodeConnectionState

	// The unique ID of the Node. Unique in the Cluster. The ID is the resource-id
	// portion of the ARN.
	Id *string

	// The ARN of the EC2 instance hosting the Node.
	InstanceArn *string

	// The name that you specified for the Node.
	Name *string

	// Documentation update needed
	NodeInterfaceMappings []types.NodeInterfaceMapping

	// The initial role current role of the Node in the Cluster. ACTIVE means the Node
	// is available for encoding. BACKUP means the Node is a redundant Node and might
	// get used if an ACTIVE Node fails.
	Role types.NodeRole

	// The current state of the Node.
	State types.NodeState

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateNodeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateNode{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateNode{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateNode"); err != nil {
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
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = addOpUpdateNodeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateNode(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateNode(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateNode",
	}
}
