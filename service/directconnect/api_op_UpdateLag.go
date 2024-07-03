// Code generated by smithy-go-codegen DO NOT EDIT.

package directconnect

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/directconnect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the attributes of the specified link aggregation group (LAG).
//
// You can update the following LAG attributes:
//
//   - The name of the LAG.
//
//   - The value for the minimum number of connections that must be operational
//     for the LAG itself to be operational.
//
//   - The LAG's MACsec encryption mode.
//
// Amazon Web Services assigns this value to each connection which is part of the
//
//	LAG.
//
//	- The tags
//
// If you adjust the threshold value for the minimum number of operational
// connections, ensure that the new value does not cause the LAG to fall below the
// threshold and become non-operational.
func (c *Client) UpdateLag(ctx context.Context, params *UpdateLagInput, optFns ...func(*Options)) (*UpdateLagOutput, error) {
	if params == nil {
		params = &UpdateLagInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateLag", params, optFns, c.addOperationUpdateLagMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateLagOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateLagInput struct {

	// The ID of the LAG.
	//
	// This member is required.
	LagId *string

	// The LAG MAC Security (MACsec) encryption mode.
	//
	// Amazon Web Services applies the value to all connections which are part of the
	// LAG.
	EncryptionMode *string

	// The name of the LAG.
	LagName *string

	// The minimum number of physical connections that must be operational for the LAG
	// itself to be operational.
	MinimumLinks int32

	noSmithyDocumentSerde
}

// Information about a link aggregation group (LAG).
type UpdateLagOutput struct {

	// Indicates whether the LAG can host other connections.
	AllowsHostedConnections bool

	// The Direct Connect endpoint that hosts the LAG.
	//
	// Deprecated: This member has been deprecated.
	AwsDevice *string

	// The Direct Connect endpoint that hosts the LAG.
	AwsDeviceV2 *string

	// The Direct Connect endpoint that terminates the logical connection. This device
	// might be different than the device that terminates the physical connection.
	AwsLogicalDeviceId *string

	// The connections bundled by the LAG.
	Connections []types.Connection

	// The individual bandwidth of the physical connections bundled by the LAG. The
	// possible values are 1Gbps, 10Gbps, 100Gbps, or 400 Gbps..
	ConnectionsBandwidth *string

	// The LAG MAC Security (MACsec) encryption mode.
	//
	// The valid values are no_encrypt , should_encrypt , and must_encrypt .
	EncryptionMode *string

	// Indicates whether the LAG supports a secondary BGP peer in the same address
	// family (IPv4/IPv6).
	HasLogicalRedundancy types.HasLogicalRedundancy

	// Indicates whether jumbo frames are supported.
	JumboFrameCapable *bool

	// The ID of the LAG.
	LagId *string

	// The name of the LAG.
	LagName *string

	// The state of the LAG. The following are the possible values:
	//
	//   - requested : The initial state of a LAG. The LAG stays in the requested state
	//   until the Letter of Authorization (LOA) is available.
	//
	//   - pending : The LAG has been approved and is being initialized.
	//
	//   - available : The network link is established and the LAG is ready for use.
	//
	//   - down : The network link is down.
	//
	//   - deleting : The LAG is being deleted.
	//
	//   - deleted : The LAG is deleted.
	//
	//   - unknown : The state of the LAG is not available.
	LagState types.LagState

	// The location of the LAG.
	Location *string

	// Indicates whether the LAG supports MAC Security (MACsec).
	MacSecCapable *bool

	// The MAC Security (MACsec) security keys associated with the LAG.
	MacSecKeys []types.MacSecKey

	// The minimum number of physical dedicated connections that must be operational
	// for the LAG itself to be operational.
	MinimumLinks int32

	// The number of physical dedicated connections initially provisioned and bundled
	// by the LAG. You can have a maximum of four connections when the port speed is 1
	// Gbps or 10 Gbps, or two when the port speed is 100 Gbps or 400 Gbps.
	NumberOfConnections int32

	// The ID of the Amazon Web Services account that owns the LAG.
	OwnerAccount *string

	// The name of the service provider associated with the LAG.
	ProviderName *string

	// The Amazon Web Services Region where the connection is located.
	Region *string

	// The tags associated with the LAG.
	Tags []types.Tag

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateLagMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateLag{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateLag{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateLag"); err != nil {
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
	if err = addOpUpdateLagValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateLag(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateLag(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateLag",
	}
}
