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

// Creates a BGP peer on the specified virtual interface.
//
// You must create a BGP peer for the corresponding address family (IPv4/IPv6) in
// order to access Amazon Web Services resources that also use that address family.
//
// If logical redundancy is not supported by the connection, interconnect, or LAG,
// the BGP peer cannot be in the same address family as an existing BGP peer on the
// virtual interface.
//
// When creating a IPv6 BGP peer, omit the Amazon address and customer address.
// IPv6 addresses are automatically assigned from the Amazon pool of IPv6
// addresses; you cannot specify custom IPv6 addresses.
//
// If you let Amazon Web Services auto-assign IPv4 addresses, a /30 CIDR will be
// allocated from 169.254.0.0/16. Amazon Web Services does not recommend this
// option if you intend to use the customer router peer IP address as the source
// and destination for traffic. Instead you should use RFC 1918 or other
// addressing, and specify the address yourself. For more information about RFC
// 1918 see [Address Allocation for Private Internets].
//
// For a public virtual interface, the Autonomous System Number (ASN) must be
// private or already on the allow list for the virtual interface.
//
// [Address Allocation for Private Internets]: https://datatracker.ietf.org/doc/html/rfc1918
func (c *Client) CreateBGPPeer(ctx context.Context, params *CreateBGPPeerInput, optFns ...func(*Options)) (*CreateBGPPeerOutput, error) {
	if params == nil {
		params = &CreateBGPPeerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateBGPPeer", params, optFns, c.addOperationCreateBGPPeerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateBGPPeerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateBGPPeerInput struct {

	// Information about the BGP peer.
	NewBGPPeer *types.NewBGPPeer

	// The ID of the virtual interface.
	VirtualInterfaceId *string

	noSmithyDocumentSerde
}

type CreateBGPPeerOutput struct {

	// The virtual interface.
	VirtualInterface *types.VirtualInterface

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateBGPPeerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateBGPPeer{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateBGPPeer{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateBGPPeer"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateBGPPeer(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateBGPPeer(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateBGPPeer",
	}
}
