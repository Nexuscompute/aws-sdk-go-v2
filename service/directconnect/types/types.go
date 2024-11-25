// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// The Amazon Web Services Cloud WAN core network that the Direct Connect
// attachment is associated with.
type AssociatedCoreNetwork struct {

	// the ID of the Direct Connect attachment
	AttachmentId *string

	// The ID of the Cloud WAN core network.
	Id *string

	// The account owner of the Cloud WAN core network.
	OwnerAccount *string

	noSmithyDocumentSerde
}

// Information about the associated gateway.
type AssociatedGateway struct {

	// The ID of the associated gateway.
	Id *string

	// The ID of the Amazon Web Services account that owns the associated virtual
	// private gateway or transit gateway.
	OwnerAccount *string

	// The Region where the associated gateway is located.
	Region *string

	// The type of associated gateway.
	Type GatewayType

	noSmithyDocumentSerde
}

// Information about a BGP peer.
type BGPPeer struct {

	// The address family for the BGP peer.
	AddressFamily AddressFamily

	// The IP address assigned to the Amazon interface.
	AmazonAddress *string

	// The autonomous system (AS) number for Border Gateway Protocol (BGP)
	// configuration.
	Asn int32

	// The authentication key for BGP configuration. This string has a minimum length
	// of 6 characters and and a maximun lenth of 80 characters.
	AuthKey *string

	// The Direct Connect endpoint that terminates the BGP peer.
	AwsDeviceV2 *string

	// The Direct Connect endpoint that terminates the logical connection. This device
	// might be different than the device that terminates the physical connection.
	AwsLogicalDeviceId *string

	// The ID of the BGP peer.
	BgpPeerId *string

	// The state of the BGP peer. The following are the possible values:
	//
	//   - verifying : The BGP peering addresses or ASN require validation before the
	//   BGP peer can be created. This state applies only to public virtual interfaces.
	//
	//   - pending : The BGP peer is created, and remains in this state until it is
	//   ready to be established.
	//
	//   - available : The BGP peer is ready to be established.
	//
	//   - deleting : The BGP peer is being deleted.
	//
	//   - deleted : The BGP peer is deleted and cannot be established.
	BgpPeerState BGPPeerState

	// The status of the BGP peer. The following are the possible values:
	//
	//   - up : The BGP peer is established. This state does not indicate the state of
	//   the routing function. Ensure that you are receiving routes over the BGP session.
	//
	//   - down : The BGP peer is down.
	//
	//   - unknown : The BGP peer status is not available.
	BgpStatus BGPStatus

	// The IP address assigned to the customer interface.
	CustomerAddress *string

	noSmithyDocumentSerde
}

// Information about an Direct Connect connection.
type Connection struct {

	// The Direct Connect endpoint on which the physical connection terminates.
	//
	// Deprecated: This member has been deprecated.
	AwsDevice *string

	// The Direct Connect endpoint that terminates the physical connection.
	AwsDeviceV2 *string

	// The Direct Connect endpoint that terminates the logical connection. This device
	// might be different than the device that terminates the physical connection.
	AwsLogicalDeviceId *string

	// The bandwidth of the connection.
	Bandwidth *string

	// The ID of the connection.
	ConnectionId *string

	// The name of the connection.
	ConnectionName *string

	// The state of the connection. The following are the possible values:
	//
	//   - ordering : The initial state of a hosted connection provisioned on an
	//   interconnect. The connection stays in the ordering state until the owner of the
	//   hosted connection confirms or declines the connection order.
	//
	//   - requested : The initial state of a standard connection. The connection stays
	//   in the requested state until the Letter of Authorization (LOA) is sent to the
	//   customer.
	//
	//   - pending : The connection has been approved and is being initialized.
	//
	//   - available : The network link is up and the connection is ready for use.
	//
	//   - down : The network link is down.
	//
	//   - deleting : The connection is being deleted.
	//
	//   - deleted : The connection has been deleted.
	//
	//   - rejected : A hosted connection in the ordering state enters the rejected
	//   state if it is deleted by the customer.
	//
	//   - unknown : The state of the connection is not available.
	ConnectionState ConnectionState

	// The MAC Security (MACsec) connection encryption mode.
	//
	// The valid values are no_encrypt , should_encrypt , and must_encrypt .
	EncryptionMode *string

	// Indicates whether the connection supports a secondary BGP peer in the same
	// address family (IPv4/IPv6).
	HasLogicalRedundancy HasLogicalRedundancy

	// Indicates whether jumbo frames are supported.
	JumboFrameCapable *bool

	// The ID of the LAG.
	LagId *string

	// The time of the most recent call to DescribeLoa for this connection.
	LoaIssueTime *time.Time

	// The location of the connection.
	Location *string

	// Indicates whether the connection supports MAC Security (MACsec).
	MacSecCapable *bool

	// The MAC Security (MACsec) security keys associated with the connection.
	MacSecKeys []MacSecKey

	// The ID of the Amazon Web Services account that owns the connection.
	OwnerAccount *string

	// The name of the Direct Connect service provider associated with the connection.
	PartnerName *string

	// The MAC Security (MACsec) port link status of the connection.
	//
	// The valid values are Encryption Up , which means that there is an active
	// Connection Key Name, or Encryption Down .
	PortEncryptionStatus *string

	// The name of the service provider associated with the connection.
	ProviderName *string

	// The Amazon Web Services Region where the connection is located.
	Region *string

	// The tags associated with the connection.
	Tags []Tag

	// The ID of the VLAN.
	Vlan int32

	noSmithyDocumentSerde
}

// The name and status of a customer agreement.
type CustomerAgreement struct {

	// The name of the agreement.
	AgreementName *string

	// The status of the customer agreement. This will be either signed or unsigned
	Status *string

	noSmithyDocumentSerde
}

// Information about a Direct Connect gateway, which enables you to connect
// virtual interfaces and virtual private gateway or transit gateways.
type DirectConnectGateway struct {

	// The autonomous system number (ASN) for the Amazon side of the connection.
	AmazonSideAsn *int64

	// The ID of the Direct Connect gateway.
	DirectConnectGatewayId *string

	// The name of the Direct Connect gateway.
	DirectConnectGatewayName *string

	// The state of the Direct Connect gateway. The following are the possible values:
	//
	//   - pending : The initial state after calling CreateDirectConnectGateway.
	//
	//   - available : The Direct Connect gateway is ready for use.
	//
	//   - deleting : The initial state after calling DeleteDirectConnectGateway.
	//
	//   - deleted : The Direct Connect gateway is deleted and cannot pass traffic.
	DirectConnectGatewayState DirectConnectGatewayState

	// The ID of the Amazon Web Services account that owns the Direct Connect gateway.
	OwnerAccount *string

	// The error message if the state of an object failed to advance.
	StateChangeError *string

	noSmithyDocumentSerde
}

// Information about an association between a Direct Connect gateway and a virtual
// private gateway or transit gateway.
type DirectConnectGatewayAssociation struct {

	// The Amazon VPC prefixes to advertise to the Direct Connect gateway.
	AllowedPrefixesToDirectConnectGateway []RouteFilterPrefix

	// The ID of the Cloud WAN core network associated with the Direct Connect
	// attachment.
	AssociatedCoreNetwork *AssociatedCoreNetwork

	// Information about the associated gateway.
	AssociatedGateway *AssociatedGateway

	// The ID of the Direct Connect gateway association.
	AssociationId *string

	// The state of the association. The following are the possible values:
	//
	//   - associating : The initial state after calling CreateDirectConnectGatewayAssociation.
	//
	//   - associated : The Direct Connect gateway and virtual private gateway or
	//   transit gateway are successfully associated and ready to pass traffic.
	//
	//   - disassociating : The initial state after calling DeleteDirectConnectGatewayAssociation.
	//
	//   - disassociated : The virtual private gateway or transit gateway is
	//   disassociated from the Direct Connect gateway. Traffic flow between the Direct
	//   Connect gateway and virtual private gateway or transit gateway is stopped.
	//
	//   - updating : The CIDR blocks for the virtual private gateway or transit
	//   gateway are currently being updated. This could be new CIDR blocks added or
	//   current CIDR blocks removed.
	AssociationState DirectConnectGatewayAssociationState

	// The ID of the Direct Connect gateway.
	DirectConnectGatewayId *string

	// The ID of the Amazon Web Services account that owns the associated gateway.
	DirectConnectGatewayOwnerAccount *string

	// The error message if the state of an object failed to advance.
	StateChangeError *string

	// The ID of the virtual private gateway. Applies only to private virtual
	// interfaces.
	VirtualGatewayId *string

	// The ID of the Amazon Web Services account that owns the virtual private gateway.
	VirtualGatewayOwnerAccount *string

	// The Amazon Web Services Region where the virtual private gateway is located.
	//
	// Deprecated: This member has been deprecated.
	VirtualGatewayRegion *string

	noSmithyDocumentSerde
}

// Information about the proposal request to attach a virtual private gateway to a
// Direct Connect gateway.
type DirectConnectGatewayAssociationProposal struct {

	// Information about the associated gateway.
	AssociatedGateway *AssociatedGateway

	// The ID of the Direct Connect gateway.
	DirectConnectGatewayId *string

	// The ID of the Amazon Web Services account that owns the Direct Connect gateway.
	DirectConnectGatewayOwnerAccount *string

	// The existing Amazon VPC prefixes advertised to the Direct Connect gateway.
	ExistingAllowedPrefixesToDirectConnectGateway []RouteFilterPrefix

	// The ID of the association proposal.
	ProposalId *string

	// The state of the proposal. The following are possible values:
	//
	//   - accepted : The proposal has been accepted. The Direct Connect gateway
	//   association is available to use in this state.
	//
	//   - deleted : The proposal has been deleted by the owner that made the proposal.
	//   The Direct Connect gateway association cannot be used in this state.
	//
	//   - requested : The proposal has been requested. The Direct Connect gateway
	//   association cannot be used in this state.
	ProposalState DirectConnectGatewayAssociationProposalState

	// The Amazon VPC prefixes to advertise to the Direct Connect gateway.
	RequestedAllowedPrefixesToDirectConnectGateway []RouteFilterPrefix

	noSmithyDocumentSerde
}

// Information about an attachment between a Direct Connect gateway and a virtual
// interface.
type DirectConnectGatewayAttachment struct {

	// The state of the attachment. The following are the possible values:
	//
	//   - attaching : The initial state after a virtual interface is created using the
	//   Direct Connect gateway.
	//
	//   - attached : The Direct Connect gateway and virtual interface are attached and
	//   ready to pass traffic.
	//
	//   - detaching : The initial state after calling DeleteVirtualInterface.
	//
	//   - detached : The virtual interface is detached from the Direct Connect
	//   gateway. Traffic flow between the Direct Connect gateway and virtual interface
	//   is stopped.
	AttachmentState DirectConnectGatewayAttachmentState

	// The type of attachment.
	AttachmentType DirectConnectGatewayAttachmentType

	// The ID of the Direct Connect gateway.
	DirectConnectGatewayId *string

	// The error message if the state of an object failed to advance.
	StateChangeError *string

	// The ID of the virtual interface.
	VirtualInterfaceId *string

	// The ID of the Amazon Web Services account that owns the virtual interface.
	VirtualInterfaceOwnerAccount *string

	// The Amazon Web Services Region where the virtual interface is located.
	VirtualInterfaceRegion *string

	noSmithyDocumentSerde
}

// Information about an interconnect.
type Interconnect struct {

	// The Direct Connect endpoint on which the physical connection terminates.
	//
	// Deprecated: This member has been deprecated.
	AwsDevice *string

	// The Direct Connect endpoint that terminates the physical connection.
	AwsDeviceV2 *string

	// The Direct Connect endpoint that terminates the logical connection. This device
	// might be different than the device that terminates the physical connection.
	AwsLogicalDeviceId *string

	// The bandwidth of the connection.
	Bandwidth *string

	// Indicates whether the interconnect supports a secondary BGP in the same address
	// family (IPv4/IPv6).
	HasLogicalRedundancy HasLogicalRedundancy

	// The ID of the interconnect.
	InterconnectId *string

	// The name of the interconnect.
	InterconnectName *string

	// The state of the interconnect. The following are the possible values:
	//
	//   - requested : The initial state of an interconnect. The interconnect stays in
	//   the requested state until the Letter of Authorization (LOA) is sent to the
	//   customer.
	//
	//   - pending : The interconnect is approved, and is being initialized.
	//
	//   - available : The network link is up, and the interconnect is ready for use.
	//
	//   - down : The network link is down.
	//
	//   - deleting : The interconnect is being deleted.
	//
	//   - deleted : The interconnect is deleted.
	//
	//   - unknown : The state of the interconnect is not available.
	InterconnectState InterconnectState

	// Indicates whether jumbo frames are supported.
	JumboFrameCapable *bool

	// The ID of the LAG.
	LagId *string

	// The time of the most recent call to DescribeLoa for this connection.
	LoaIssueTime *time.Time

	// The location of the connection.
	Location *string

	// The name of the service provider associated with the interconnect.
	ProviderName *string

	// The Amazon Web Services Region where the connection is located.
	Region *string

	// The tags associated with the interconnect.
	Tags []Tag

	noSmithyDocumentSerde
}

// Information about a link aggregation group (LAG).
type Lag struct {

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
	Connections []Connection

	// The individual bandwidth of the physical connections bundled by the LAG. The
	// possible values are 1Gbps, 10Gbps, 100Gbps, or 400 Gbps..
	ConnectionsBandwidth *string

	// The LAG MAC Security (MACsec) encryption mode.
	//
	// The valid values are no_encrypt , should_encrypt , and must_encrypt .
	EncryptionMode *string

	// Indicates whether the LAG supports a secondary BGP peer in the same address
	// family (IPv4/IPv6).
	HasLogicalRedundancy HasLogicalRedundancy

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
	LagState LagState

	// The location of the LAG.
	Location *string

	// Indicates whether the LAG supports MAC Security (MACsec).
	MacSecCapable *bool

	// The MAC Security (MACsec) security keys associated with the LAG.
	MacSecKeys []MacSecKey

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
	Tags []Tag

	noSmithyDocumentSerde
}

// Information about a Letter of Authorization - Connecting Facility Assignment
// (LOA-CFA) for a connection.
type Loa struct {

	// The binary contents of the LOA-CFA document.
	LoaContent []byte

	// The standard media type for the LOA-CFA document. The only supported value is
	// application/pdf.
	LoaContentType LoaContentType

	noSmithyDocumentSerde
}

// Information about an Direct Connect location.
type Location struct {

	// The available MAC Security (MACsec) port speeds for the location.
	AvailableMacSecPortSpeeds []string

	// The available port speeds for the location.
	AvailablePortSpeeds []string

	// The name of the service provider for the location.
	AvailableProviders []string

	// The code for the location.
	LocationCode *string

	// The name of the location. This includes the name of the colocation partner and
	// the physical site of the building.
	LocationName *string

	// The Amazon Web Services Region for the location.
	Region *string

	noSmithyDocumentSerde
}

// Information about the MAC Security (MACsec) secret key.
type MacSecKey struct {

	// The Connection Key Name (CKN) for the MAC Security secret key.
	Ckn *string

	// The Amazon Resource Name (ARN) of the MAC Security (MACsec) secret key.
	SecretARN *string

	// The date that the MAC Security (MACsec) secret key takes effect. The value is
	// displayed in UTC format.
	StartOn *string

	// The state of the MAC Security (MACsec) secret key.
	//
	// The possible values are:
	//
	//   - associating : The MAC Security (MACsec) secret key is being validated and
	//   not yet associated with the connection or LAG.
	//
	//   - associated : The MAC Security (MACsec) secret key is validated and
	//   associated with the connection or LAG.
	//
	//   - disassociating : The MAC Security (MACsec) secret key is being disassociated
	//   from the connection or LAG
	//
	//   - disassociated : The MAC Security (MACsec) secret key is no longer associated
	//   with the connection or LAG.
	State *string

	noSmithyDocumentSerde
}

// Information about a new BGP peer.
type NewBGPPeer struct {

	// The address family for the BGP peer.
	AddressFamily AddressFamily

	// The IP address assigned to the Amazon interface.
	AmazonAddress *string

	// The autonomous system (AS) number for Border Gateway Protocol (BGP)
	// configuration.
	Asn int32

	// The authentication key for BGP configuration. This string has a minimum length
	// of 6 characters and and a maximun lenth of 80 characters.
	AuthKey *string

	// The IP address assigned to the customer interface.
	CustomerAddress *string

	noSmithyDocumentSerde
}

// Information about a private virtual interface.
type NewPrivateVirtualInterface struct {

	// The autonomous system (AS) number for Border Gateway Protocol (BGP)
	// configuration.
	//
	// The valid values are 1-2147483647.
	//
	// This member is required.
	Asn int32

	// The name of the virtual interface assigned by the customer network. The name
	// has a maximum of 100 characters. The following are valid characters: a-z, 0-9
	// and a hyphen (-).
	//
	// This member is required.
	VirtualInterfaceName *string

	// The ID of the VLAN.
	//
	// This member is required.
	Vlan int32

	// The address family for the BGP peer.
	AddressFamily AddressFamily

	// The IP address assigned to the Amazon interface.
	AmazonAddress *string

	// The authentication key for BGP configuration. This string has a minimum length
	// of 6 characters and and a maximun lenth of 80 characters.
	AuthKey *string

	// The IP address assigned to the customer interface.
	CustomerAddress *string

	// The ID of the Direct Connect gateway.
	DirectConnectGatewayId *string

	// Indicates whether to enable or disable SiteLink.
	EnableSiteLink *bool

	// The maximum transmission unit (MTU), in bytes. The supported values are 1500
	// and 8500. The default value is 1500.
	Mtu *int32

	// The tags associated with the private virtual interface.
	Tags []Tag

	// The ID of the virtual private gateway.
	VirtualGatewayId *string

	noSmithyDocumentSerde
}

// Information about a private virtual interface to be provisioned on a connection.
type NewPrivateVirtualInterfaceAllocation struct {

	// The autonomous system (AS) number for Border Gateway Protocol (BGP)
	// configuration.
	//
	// The valid values are 1-2147483647.
	//
	// This member is required.
	Asn int32

	// The name of the virtual interface assigned by the customer network. The name
	// has a maximum of 100 characters. The following are valid characters: a-z, 0-9
	// and a hyphen (-).
	//
	// This member is required.
	VirtualInterfaceName *string

	// The ID of the VLAN.
	//
	// This member is required.
	Vlan int32

	// The address family for the BGP peer.
	AddressFamily AddressFamily

	// The IP address assigned to the Amazon interface.
	AmazonAddress *string

	// The authentication key for BGP configuration. This string has a minimum length
	// of 6 characters and and a maximun lenth of 80 characters.
	AuthKey *string

	// The IP address assigned to the customer interface.
	CustomerAddress *string

	// The maximum transmission unit (MTU), in bytes. The supported values are 1500
	// and 8500. The default value is 1500.
	Mtu *int32

	// The tags associated with the private virtual interface.
	Tags []Tag

	noSmithyDocumentSerde
}

// Information about a public virtual interface.
type NewPublicVirtualInterface struct {

	// The autonomous system (AS) number for Border Gateway Protocol (BGP)
	// configuration.
	//
	// The valid values are 1-2147483647.
	//
	// This member is required.
	Asn int32

	// The name of the virtual interface assigned by the customer network. The name
	// has a maximum of 100 characters. The following are valid characters: a-z, 0-9
	// and a hyphen (-).
	//
	// This member is required.
	VirtualInterfaceName *string

	// The ID of the VLAN.
	//
	// This member is required.
	Vlan int32

	// The address family for the BGP peer.
	AddressFamily AddressFamily

	// The IP address assigned to the Amazon interface.
	AmazonAddress *string

	// The authentication key for BGP configuration. This string has a minimum length
	// of 6 characters and and a maximun lenth of 80 characters.
	AuthKey *string

	// The IP address assigned to the customer interface.
	CustomerAddress *string

	// The routes to be advertised to the Amazon Web Services network in this Region.
	// Applies to public virtual interfaces.
	RouteFilterPrefixes []RouteFilterPrefix

	// The tags associated with the public virtual interface.
	Tags []Tag

	noSmithyDocumentSerde
}

// Information about a public virtual interface to be provisioned on a connection.
type NewPublicVirtualInterfaceAllocation struct {

	// The autonomous system (AS) number for Border Gateway Protocol (BGP)
	// configuration.
	//
	// The valid values are 1-2147483647.
	//
	// This member is required.
	Asn int32

	// The name of the virtual interface assigned by the customer network. The name
	// has a maximum of 100 characters. The following are valid characters: a-z, 0-9
	// and a hyphen (-).
	//
	// This member is required.
	VirtualInterfaceName *string

	// The ID of the VLAN.
	//
	// This member is required.
	Vlan int32

	// The address family for the BGP peer.
	AddressFamily AddressFamily

	// The IP address assigned to the Amazon interface.
	AmazonAddress *string

	// The authentication key for BGP configuration. This string has a minimum length
	// of 6 characters and and a maximun lenth of 80 characters.
	AuthKey *string

	// The IP address assigned to the customer interface.
	CustomerAddress *string

	// The routes to be advertised to the Amazon Web Services network in this Region.
	// Applies to public virtual interfaces.
	RouteFilterPrefixes []RouteFilterPrefix

	// The tags associated with the public virtual interface.
	Tags []Tag

	noSmithyDocumentSerde
}

// Information about a transit virtual interface.
type NewTransitVirtualInterface struct {

	// The address family for the BGP peer.
	AddressFamily AddressFamily

	// The IP address assigned to the Amazon interface.
	AmazonAddress *string

	// The autonomous system (AS) number for Border Gateway Protocol (BGP)
	// configuration.
	//
	// The valid values are 1-2147483647.
	Asn int32

	// The authentication key for BGP configuration. This string has a minimum length
	// of 6 characters and and a maximun lenth of 80 characters.
	AuthKey *string

	// The IP address assigned to the customer interface.
	CustomerAddress *string

	// The ID of the Direct Connect gateway.
	DirectConnectGatewayId *string

	// Indicates whether to enable or disable SiteLink.
	EnableSiteLink *bool

	// The maximum transmission unit (MTU), in bytes. The supported values are 1500
	// and 8500. The default value is 1500.
	Mtu *int32

	// The tags associated with the transitive virtual interface.
	Tags []Tag

	// The name of the virtual interface assigned by the customer network. The name
	// has a maximum of 100 characters. The following are valid characters: a-z, 0-9
	// and a hyphen (-).
	VirtualInterfaceName *string

	// The ID of the VLAN.
	Vlan int32

	noSmithyDocumentSerde
}

// Information about a transit virtual interface to be provisioned on a connection.
type NewTransitVirtualInterfaceAllocation struct {

	// The address family for the BGP peer.
	AddressFamily AddressFamily

	// The IP address assigned to the Amazon interface.
	AmazonAddress *string

	// The autonomous system (AS) number for Border Gateway Protocol (BGP)
	// configuration.
	//
	// The valid values are 1-2147483647.
	Asn int32

	// The authentication key for BGP configuration. This string has a minimum length
	// of 6 characters and and a maximun lenth of 80 characters.
	AuthKey *string

	// The IP address assigned to the customer interface.
	CustomerAddress *string

	// The maximum transmission unit (MTU), in bytes. The supported values are 1500
	// and 8500. The default value is 1500
	Mtu *int32

	// The tags associated with the transitive virtual interface.
	Tags []Tag

	// The name of the virtual interface assigned by the customer network. The name
	// has a maximum of 100 characters. The following are valid characters: a-z, 0-9
	// and a hyphen (-).
	VirtualInterfaceName *string

	// The ID of the VLAN.
	Vlan int32

	noSmithyDocumentSerde
}

// Information about a tag associated with an Direct Connect resource.
type ResourceTag struct {

	// The Amazon Resource Name (ARN) of the resource.
	ResourceArn *string

	// The tags.
	Tags []Tag

	noSmithyDocumentSerde
}

// Information about a route filter prefix that a customer can advertise through
// Border Gateway Protocol (BGP) over a public virtual interface.
type RouteFilterPrefix struct {

	// The CIDR block for the advertised route. Separate multiple routes using commas.
	// An IPv6 CIDR must use /64 or shorter.
	Cidr *string

	noSmithyDocumentSerde
}

// Information about the virtual router.
type RouterType struct {

	// The virtual interface router platform.
	Platform *string

	// Identifies the router by a combination of vendor, platform, and software
	// version. For example, CiscoSystemsInc-2900SeriesRouters-IOS124 .
	RouterTypeIdentifier *string

	// The router software.
	Software *string

	// The vendor for the virtual interface's router.
	Vendor *string

	// The template for the virtual interface's router.
	XsltTemplateName *string

	// The MAC Security (MACsec) template for the virtual interface's router.
	XsltTemplateNameForMacSec *string

	noSmithyDocumentSerde
}

// Information about a tag.
type Tag struct {

	// The key.
	//
	// This member is required.
	Key *string

	// The value.
	Value *string

	noSmithyDocumentSerde
}

// Information about a virtual private gateway for a private virtual interface.
type VirtualGateway struct {

	// The ID of the virtual private gateway.
	VirtualGatewayId *string

	// The state of the virtual private gateway. The following are the possible values:
	//
	//   - pending : Initial state after creating the virtual private gateway.
	//
	//   - available : Ready for use by a private virtual interface.
	//
	//   - deleting : Initial state after deleting the virtual private gateway.
	//
	//   - deleted : The virtual private gateway is deleted. The private virtual
	//   interface is unable to send traffic over this gateway.
	VirtualGatewayState *string

	noSmithyDocumentSerde
}

// Information about a virtual interface.
type VirtualInterface struct {

	// The address family for the BGP peer.
	AddressFamily AddressFamily

	// The IP address assigned to the Amazon interface.
	AmazonAddress *string

	// The autonomous system number (ASN) for the Amazon side of the connection.
	AmazonSideAsn *int64

	// The autonomous system (AS) number for Border Gateway Protocol (BGP)
	// configuration.
	//
	// The valid values are 1-2147483647.
	Asn int32

	// The authentication key for BGP configuration. This string has a minimum length
	// of 6 characters and and a maximun lenth of 80 characters.
	AuthKey *string

	// The Direct Connect endpoint that terminates the physical connection.
	AwsDeviceV2 *string

	// The Direct Connect endpoint that terminates the logical connection. This device
	// might be different than the device that terminates the physical connection.
	AwsLogicalDeviceId *string

	// The BGP peers configured on this virtual interface.
	BgpPeers []BGPPeer

	// The ID of the connection.
	ConnectionId *string

	// The IP address assigned to the customer interface.
	CustomerAddress *string

	// The customer router configuration.
	CustomerRouterConfig *string

	// The ID of the Direct Connect gateway.
	DirectConnectGatewayId *string

	// Indicates whether jumbo frames are supported.
	JumboFrameCapable *bool

	// The location of the connection.
	Location *string

	// The maximum transmission unit (MTU), in bytes. The supported values are 1500
	// and 8500. The default value is 1500
	Mtu *int32

	// The ID of the Amazon Web Services account that owns the virtual interface.
	OwnerAccount *string

	// The Amazon Web Services Region where the virtual interface is located.
	Region *string

	// The routes to be advertised to the Amazon Web Services network in this Region.
	// Applies to public virtual interfaces.
	RouteFilterPrefixes []RouteFilterPrefix

	// Indicates whether SiteLink is enabled.
	SiteLinkEnabled *bool

	// The tags associated with the virtual interface.
	Tags []Tag

	// The ID of the virtual private gateway. Applies only to private virtual
	// interfaces.
	VirtualGatewayId *string

	// The ID of the virtual interface.
	VirtualInterfaceId *string

	// The name of the virtual interface assigned by the customer network. The name
	// has a maximum of 100 characters. The following are valid characters: a-z, 0-9
	// and a hyphen (-).
	VirtualInterfaceName *string

	// The state of the virtual interface. The following are the possible values:
	//
	//   - confirming : The creation of the virtual interface is pending confirmation
	//   from the virtual interface owner. If the owner of the virtual interface is
	//   different from the owner of the connection on which it is provisioned, then the
	//   virtual interface will remain in this state until it is confirmed by the virtual
	//   interface owner.
	//
	//   - verifying : This state only applies to public virtual interfaces. Each
	//   public virtual interface needs validation before the virtual interface can be
	//   created.
	//
	//   - pending : A virtual interface is in this state from the time that it is
	//   created until the virtual interface is ready to forward traffic.
	//
	//   - available : A virtual interface that is able to forward traffic.
	//
	//   - down : A virtual interface that is BGP down.
	//
	//   - deleting : A virtual interface is in this state immediately after calling DeleteVirtualInterface
	//   until it can no longer forward traffic.
	//
	//   - deleted : A virtual interface that cannot forward traffic.
	//
	//   - rejected : The virtual interface owner has declined creation of the virtual
	//   interface. If a virtual interface in the Confirming state is deleted by the
	//   virtual interface owner, the virtual interface enters the Rejected state.
	//
	//   - unknown : The state of the virtual interface is not available.
	VirtualInterfaceState VirtualInterfaceState

	// The type of virtual interface. The possible values are private , public and
	// transit .
	VirtualInterfaceType *string

	// The ID of the VLAN.
	Vlan int32

	noSmithyDocumentSerde
}

// Information about the virtual interface failover test.
type VirtualInterfaceTestHistory struct {

	// The BGP peers that were put in the DOWN state as part of the virtual interface
	// failover test.
	BgpPeers []string

	// The time that the virtual interface moves out of the DOWN state.
	EndTime *time.Time

	// The owner ID of the tested virtual interface.
	OwnerAccount *string

	// The time that the virtual interface moves to the DOWN state.
	StartTime *time.Time

	// The status of the virtual interface failover test.
	Status *string

	// The time that the virtual interface failover test ran in minutes.
	TestDurationInMinutes *int32

	// The ID of the virtual interface failover test.
	TestId *string

	// The ID of the tested virtual interface.
	VirtualInterfaceId *string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
