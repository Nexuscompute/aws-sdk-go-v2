// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// Describes the current status of an account within an Amazon Web Services
// Organization, including service-linked roles (SLRs).
type AccountStatus struct {

	// The ID of an account within the Amazon Web Services Organization.
	AccountId *string

	// The status of SLR deployment for the account.
	SLRDeploymentStatus *string

	noSmithyDocumentSerde
}

// Describes a core network attachment.
type Attachment struct {

	// The ID of the attachment.
	AttachmentId *string

	// The policy rule number associated with the attachment.
	AttachmentPolicyRuleNumber *int32

	// The type of attachment.
	AttachmentType AttachmentType

	// The ARN of a core network.
	CoreNetworkArn *string

	// The ID of a core network.
	CoreNetworkId *string

	// The timestamp when the attachment was created.
	CreatedAt *time.Time

	// The Region where the edge is located. This is returned for all attachment types
	// except a Direct Connect gateway attachment, which instead returns EdgeLocations .
	EdgeLocation *string

	// The edge locations that the Direct Connect gateway is associated with. This is
	// returned only for Direct Connect gateway attachments. All other attachment types
	// retrun EdgeLocation .
	EdgeLocations []string

	// Describes the error associated with the attachment request.
	LastModificationErrors []AttachmentError

	// The name of the network function group.
	NetworkFunctionGroupName *string

	// The ID of the attachment account owner.
	OwnerAccountId *string

	// Describes a proposed change to a network function group associated with the
	// attachment.
	ProposedNetworkFunctionGroupChange *ProposedNetworkFunctionGroupChange

	// The attachment to move from one segment to another.
	ProposedSegmentChange *ProposedSegmentChange

	// The attachment resource ARN.
	ResourceArn *string

	// The name of the segment attachment.
	SegmentName *string

	// The state of the attachment.
	State AttachmentState

	// The tags associated with the attachment.
	Tags []Tag

	// The timestamp when the attachment was last updated.
	UpdatedAt *time.Time

	noSmithyDocumentSerde
}

// Describes the error associated with an attachment request.
type AttachmentError struct {

	// The error code for the attachment request.
	Code AttachmentErrorCode

	// The message associated with the error code .
	Message *string

	// The ID of the attachment request.
	RequestId *string

	// The ARN of the requested attachment resource.
	ResourceArn *string

	noSmithyDocumentSerde
}

// Specifies a location in Amazon Web Services.
type AWSLocation struct {

	// The Amazon Resource Name (ARN) of the subnet that the device is located in.
	SubnetArn *string

	// The Zone that the device is located in. Specify the ID of an Availability Zone,
	// Local Zone, Wavelength Zone, or an Outpost.
	Zone *string

	noSmithyDocumentSerde
}

// Describes bandwidth information.
type Bandwidth struct {

	// Download speed in Mbps.
	DownloadSpeed *int32

	// Upload speed in Mbps.
	UploadSpeed *int32

	noSmithyDocumentSerde
}

// Describes the BGP options.
type BgpOptions struct {

	// The Peer ASN of the BGP.
	PeerAsn *int64

	noSmithyDocumentSerde
}

// Describes a core network Connect attachment.
type ConnectAttachment struct {

	// The attachment details.
	Attachment *Attachment

	// Options for connecting an attachment.
	Options *ConnectAttachmentOptions

	// The ID of the transport attachment.
	TransportAttachmentId *string

	noSmithyDocumentSerde
}

// Describes a core network Connect attachment options.
type ConnectAttachmentOptions struct {

	// The protocol used for the attachment connection.
	Protocol TunnelProtocol

	noSmithyDocumentSerde
}

// Describes a connection.
type Connection struct {

	// The ID of the second device in the connection.
	ConnectedDeviceId *string

	// The ID of the link for the second device in the connection.
	ConnectedLinkId *string

	// The Amazon Resource Name (ARN) of the connection.
	ConnectionArn *string

	// The ID of the connection.
	ConnectionId *string

	// The date and time that the connection was created.
	CreatedAt *time.Time

	// The description of the connection.
	Description *string

	// The ID of the first device in the connection.
	DeviceId *string

	// The ID of the global network.
	GlobalNetworkId *string

	// The ID of the link for the first device in the connection.
	LinkId *string

	// The state of the connection.
	State ConnectionState

	// The tags for the connection.
	Tags []Tag

	noSmithyDocumentSerde
}

// Describes connection health.
type ConnectionHealth struct {

	// The connection status.
	Status ConnectionStatus

	// The time the status was last updated.
	Timestamp *time.Time

	// The connection type.
	Type ConnectionType

	noSmithyDocumentSerde
}

// Describes a core network Connect peer.
type ConnectPeer struct {

	// The configuration of the Connect peer.
	Configuration *ConnectPeerConfiguration

	// The ID of the attachment to connect.
	ConnectAttachmentId *string

	// The ID of the Connect peer.
	ConnectPeerId *string

	// The ID of a core network.
	CoreNetworkId *string

	// The timestamp when the Connect peer was created.
	CreatedAt *time.Time

	// The Connect peer Regions where edges are located.
	EdgeLocation *string

	// Describes the error associated with the attachment request.
	LastModificationErrors []ConnectPeerError

	// The state of the Connect peer.
	State ConnectPeerState

	// The subnet ARN for the Connect peer. This only applies only when the protocol
	// is NO_ENCAP.
	SubnetArn *string

	// The list of key-value tags associated with the Connect peer.
	Tags []Tag

	noSmithyDocumentSerde
}

// Describes a core network Connect peer association.
type ConnectPeerAssociation struct {

	// The ID of the Connect peer.
	ConnectPeerId *string

	// The ID of the device to connect to.
	DeviceId *string

	// The ID of the global network.
	GlobalNetworkId *string

	// The ID of the link.
	LinkId *string

	// The state of the Connect peer association.
	State ConnectPeerAssociationState

	noSmithyDocumentSerde
}

// Describes a core network BGP configuration.
type ConnectPeerBgpConfiguration struct {

	// The address of a core network.
	CoreNetworkAddress *string

	// The ASN of the Coret Network.
	CoreNetworkAsn *int64

	// The address of a core network Connect peer.
	PeerAddress *string

	// The ASN of the Connect peer.
	PeerAsn *int64

	noSmithyDocumentSerde
}

// Describes a core network Connect peer configuration.
type ConnectPeerConfiguration struct {

	// The Connect peer BGP configurations.
	BgpConfigurations []ConnectPeerBgpConfiguration

	// The IP address of a core network.
	CoreNetworkAddress *string

	// The inside IP addresses used for a Connect peer configuration.
	InsideCidrBlocks []string

	// The IP address of the Connect peer.
	PeerAddress *string

	// The protocol used for a Connect peer configuration.
	Protocol TunnelProtocol

	noSmithyDocumentSerde
}

// Describes an error associated with a Connect peer request
type ConnectPeerError struct {

	// The error code for the Connect peer request.
	Code ConnectPeerErrorCode

	// The message associated with the error code .
	Message *string

	// The ID of the Connect peer request.
	RequestId *string

	// The ARN of the requested Connect peer resource.
	ResourceArn *string

	noSmithyDocumentSerde
}

// Summary description of a Connect peer.
type ConnectPeerSummary struct {

	// The ID of a Connect peer attachment.
	ConnectAttachmentId *string

	// The ID of a Connect peer.
	ConnectPeerId *string

	// The state of a Connect peer.
	ConnectPeerState ConnectPeerState

	// The ID of a core network.
	CoreNetworkId *string

	// The timestamp when a Connect peer was created.
	CreatedAt *time.Time

	// The Region where the edge is located.
	EdgeLocation *string

	// The subnet ARN for the Connect peer summary.
	SubnetArn *string

	// The list of key-value tags associated with the Connect peer summary.
	Tags []Tag

	noSmithyDocumentSerde
}

// Describes a core network.
type CoreNetwork struct {

	// The ARN of a core network.
	CoreNetworkArn *string

	// The ID of a core network.
	CoreNetworkId *string

	// The timestamp when a core network was created.
	CreatedAt *time.Time

	// The description of a core network.
	Description *string

	// The edges within a core network.
	Edges []CoreNetworkEdge

	// The ID of the global network that your core network is a part of.
	GlobalNetworkId *string

	// The network function groups associated with a core network.
	NetworkFunctionGroups []CoreNetworkNetworkFunctionGroup

	// The segments within a core network.
	Segments []CoreNetworkSegment

	// The current state of a core network.
	State CoreNetworkState

	// The list of key-value tags associated with a core network.
	Tags []Tag

	noSmithyDocumentSerde
}

// Details describing a core network change.
type CoreNetworkChange struct {

	// The action to take for a core network.
	Action ChangeAction

	// The resource identifier.
	Identifier *string

	// Uniquely identifies the path for a change within the changeset. For example,
	// the IdentifierPath for a core network segment change might be
	// "CORE_NETWORK_SEGMENT/us-east-1/devsegment" .
	IdentifierPath *string

	// The new value for a core network
	NewValues *CoreNetworkChangeValues

	// The previous values for a core network.
	PreviousValues *CoreNetworkChangeValues

	// The type of change.
	Type ChangeType

	noSmithyDocumentSerde
}

// Describes a core network change event. This can be a change to a segment,
// attachment, route, etc.
type CoreNetworkChangeEvent struct {

	// The action taken for the change event.
	Action ChangeAction

	// The timestamp for an event change in status.
	EventTime *time.Time

	// Uniquely identifies the path for a change within the changeset. For example,
	// the IdentifierPath for a core network segment change might be
	// "CORE_NETWORK_SEGMENT/us-east-1/devsegment" .
	IdentifierPath *string

	// The status of the core network change event.
	Status ChangeStatus

	// Describes the type of change event.
	Type ChangeType

	// Details of the change event.
	Values *CoreNetworkChangeEventValues

	noSmithyDocumentSerde
}

// Describes a core network change event.
type CoreNetworkChangeEventValues struct {

	// The ID of the attachment if the change event is associated with an attachment.
	AttachmentId *string

	// For a STATIC_ROUTE event, this is the IP address.
	Cidr *string

	// The edge location for the core network change event.
	EdgeLocation *string

	// The changed network function group name.
	NetworkFunctionGroupName *string

	// The segment name if the change event is associated with a segment.
	SegmentName *string

	noSmithyDocumentSerde
}

// Describes a core network change.
type CoreNetworkChangeValues struct {

	// The ASN of a core network.
	Asn *int64

	// The IP addresses used for a core network.
	Cidr *string

	// The ID of the destination.
	DestinationIdentifier *string

	// The Regions where edges are located in a core network.
	EdgeLocations []string

	// The inside IP addresses used for core network change values.
	InsideCidrBlocks []string

	// The network function group name if the change event is associated with a
	// network function group.
	NetworkFunctionGroupName *string

	// The names of the segments in a core network.
	SegmentName *string

	// Describes the service insertion action.
	ServiceInsertionActions []ServiceInsertionAction

	// The shared segments for a core network change value.
	SharedSegments []string

	noSmithyDocumentSerde
}

// Describes a core network edge.
type CoreNetworkEdge struct {

	// The ASN of a core network edge.
	Asn *int64

	// The Region where a core network edge is located.
	EdgeLocation *string

	// The inside IP addresses used for core network edges.
	InsideCidrBlocks []string

	noSmithyDocumentSerde
}

// Describes a network function group.
type CoreNetworkNetworkFunctionGroup struct {

	// The core network edge locations.
	EdgeLocations []string

	// The name of the network function group.
	Name *string

	// The segments associated with the network function group.
	Segments *ServiceInsertionSegments

	noSmithyDocumentSerde
}

// Describes a core network
type CoreNetworkNetworkFunctionGroupIdentifier struct {

	// The ID of the core network.
	CoreNetworkId *string

	// The location for the core network edge.
	EdgeLocation *string

	// The network function group name.
	NetworkFunctionGroupName *string

	noSmithyDocumentSerde
}

// Describes a core network policy. You can have only one LIVE Core Policy.
type CoreNetworkPolicy struct {

	// Whether a core network policy is the current LIVE policy or the most recently
	// submitted policy.
	Alias CoreNetworkPolicyAlias

	// The state of a core network policy.
	ChangeSetState ChangeSetState

	// The ID of a core network.
	CoreNetworkId *string

	// The timestamp when a core network policy was created.
	CreatedAt *time.Time

	// The description of a core network policy.
	Description *string

	// Describes a core network policy.
	//
	// This value conforms to the media type: application/json
	PolicyDocument *string

	// Describes any errors in a core network policy.
	PolicyErrors []CoreNetworkPolicyError

	// The ID of the policy version.
	PolicyVersionId *int32

	noSmithyDocumentSerde
}

// Provides details about an error in a core network policy.
type CoreNetworkPolicyError struct {

	// The error code associated with a core network policy error.
	//
	// This member is required.
	ErrorCode *string

	// The message associated with a core network policy error code.
	//
	// This member is required.
	Message *string

	// The JSON path where the error was discovered in the policy document.
	Path *string

	noSmithyDocumentSerde
}

// Describes a core network policy version.
type CoreNetworkPolicyVersion struct {

	// Whether a core network policy is the current policy or the most recently
	// submitted policy.
	Alias CoreNetworkPolicyAlias

	// The status of the policy version change set.
	ChangeSetState ChangeSetState

	// The ID of a core network.
	CoreNetworkId *string

	// The timestamp when a core network policy version was created.
	CreatedAt *time.Time

	// The description of a core network policy version.
	Description *string

	// The ID of the policy version.
	PolicyVersionId *int32

	noSmithyDocumentSerde
}

// Describes a core network segment, which are dedicated routes. Only attachments
// within this segment can communicate with each other.
type CoreNetworkSegment struct {

	// The Regions where the edges are located.
	EdgeLocations []string

	// The name of a core network segment.
	Name *string

	// The shared segments of a core network.
	SharedSegments []string

	noSmithyDocumentSerde
}

// Returns details about a core network edge.
type CoreNetworkSegmentEdgeIdentifier struct {

	// The ID of a core network.
	CoreNetworkId *string

	// The Region where the segment edge is located.
	EdgeLocation *string

	// The name of the segment edge.
	SegmentName *string

	noSmithyDocumentSerde
}

// Returns summary information about a core network.
type CoreNetworkSummary struct {

	// a core network ARN.
	CoreNetworkArn *string

	// The ID of a core network.
	CoreNetworkId *string

	// The description of a core network.
	Description *string

	// The global network ID.
	GlobalNetworkId *string

	// The ID of the account owner.
	OwnerAccountId *string

	// The state of a core network.
	State CoreNetworkState

	// The key-value tags associated with a core network summary.
	Tags []Tag

	noSmithyDocumentSerde
}

// Describes the association between a customer gateway, a device, and a link.
type CustomerGatewayAssociation struct {

	// The Amazon Resource Name (ARN) of the customer gateway.
	CustomerGatewayArn *string

	// The ID of the device.
	DeviceId *string

	// The ID of the global network.
	GlobalNetworkId *string

	// The ID of the link.
	LinkId *string

	// The association state.
	State CustomerGatewayAssociationState

	noSmithyDocumentSerde
}

// Describes a device.
type Device struct {

	// The Amazon Web Services location of the device.
	AWSLocation *AWSLocation

	// The date and time that the site was created.
	CreatedAt *time.Time

	// The description of the device.
	Description *string

	// The Amazon Resource Name (ARN) of the device.
	DeviceArn *string

	// The ID of the device.
	DeviceId *string

	// The ID of the global network.
	GlobalNetworkId *string

	// The site location.
	Location *Location

	// The device model.
	Model *string

	// The device serial number.
	SerialNumber *string

	// The site ID.
	SiteId *string

	// The device state.
	State DeviceState

	// The tags for the device.
	Tags []Tag

	// The device type.
	Type *string

	// The device vendor.
	Vendor *string

	noSmithyDocumentSerde
}

// Describes a Direct Connect gateway attachment.
type DirectConnectGatewayAttachment struct {

	// Describes a core network attachment.
	Attachment *Attachment

	// The Direct Connect gateway attachment ARN.
	DirectConnectGatewayArn *string

	noSmithyDocumentSerde
}

// Describes the edge that's used for the override.
type EdgeOverride struct {

	// The list of edge locations.
	EdgeSets [][]string

	// The edge that should be used when overriding the current edge order.
	UseEdge *string

	noSmithyDocumentSerde
}

// Describes a global network. This is a single private network acting as a
// high-level container for your network objects, including an Amazon Web
// Services-managed Core Network.
type GlobalNetwork struct {

	// The date and time that the global network was created.
	CreatedAt *time.Time

	// The description of the global network.
	Description *string

	// The Amazon Resource Name (ARN) of the global network.
	GlobalNetworkArn *string

	// The ID of the global network.
	GlobalNetworkId *string

	// The state of the global network.
	State GlobalNetworkState

	// The tags for the global network.
	Tags []Tag

	noSmithyDocumentSerde
}

// Describes a link.
type Link struct {

	// The bandwidth for the link.
	Bandwidth *Bandwidth

	// The date and time that the link was created.
	CreatedAt *time.Time

	// The description of the link.
	Description *string

	// The ID of the global network.
	GlobalNetworkId *string

	// The Amazon Resource Name (ARN) of the link.
	LinkArn *string

	// The ID of the link.
	LinkId *string

	// The provider of the link.
	Provider *string

	// The ID of the site.
	SiteId *string

	// The state of the link.
	State LinkState

	// The tags for the link.
	Tags []Tag

	// The type of the link.
	Type *string

	noSmithyDocumentSerde
}

// Describes the association between a device and a link.
type LinkAssociation struct {

	// The device ID for the link association.
	DeviceId *string

	// The ID of the global network.
	GlobalNetworkId *string

	// The state of the association.
	LinkAssociationState LinkAssociationState

	// The ID of the link.
	LinkId *string

	noSmithyDocumentSerde
}

// Describes a location.
type Location struct {

	// The physical address.
	Address *string

	// The latitude.
	Latitude *string

	// The longitude.
	Longitude *string

	noSmithyDocumentSerde
}

// Describes a network function group for service insertion.
type NetworkFunctionGroup struct {

	// The name of the network function group.
	Name *string

	noSmithyDocumentSerde
}

// Describes a network resource.
type NetworkResource struct {

	// The Amazon Web Services account ID.
	AccountId *string

	// The Amazon Web Services Region.
	AwsRegion *string

	// The ID of a core network.
	CoreNetworkId *string

	// Information about the resource, in JSON format. Network Manager gets this
	// information by describing the resource using its Describe API call.
	Definition *string

	// The time that the resource definition was retrieved.
	DefinitionTimestamp *time.Time

	// The resource metadata.
	Metadata map[string]string

	// The ARN of the gateway.
	RegisteredGatewayArn *string

	// The ARN of the resource.
	ResourceArn *string

	// The ID of the resource.
	ResourceId *string

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
	//   - attachment
	//
	//   - connect-peer
	//
	//   - connection
	//
	//   - core-network
	//
	//   - device
	//
	//   - link
	//
	//   - peering
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

	// The tags.
	Tags []Tag

	noSmithyDocumentSerde
}

// Describes a resource count.
type NetworkResourceCount struct {

	// The resource count.
	Count *int32

	// The resource type.
	ResourceType *string

	noSmithyDocumentSerde
}

// Describes a network resource.
type NetworkResourceSummary struct {

	// Information about the resource, in JSON format. Network Manager gets this
	// information by describing the resource using its Describe API call.
	Definition *string

	// Indicates whether this is a middlebox appliance.
	IsMiddlebox bool

	// The value for the Name tag.
	NameTag *string

	// The ARN of the gateway.
	RegisteredGatewayArn *string

	// The ARN of the resource.
	ResourceArn *string

	// The resource type.
	ResourceType *string

	noSmithyDocumentSerde
}

// Describes a network route.
type NetworkRoute struct {

	// A unique identifier for the route, such as a CIDR block.
	DestinationCidrBlock *string

	// The destinations.
	Destinations []NetworkRouteDestination

	// The ID of the prefix list.
	PrefixListId *string

	// The route state. The possible values are active and blackhole .
	State RouteState

	// The route type. The possible values are propagated and static .
	Type RouteType

	noSmithyDocumentSerde
}

// Describes the destination of a network route.
type NetworkRouteDestination struct {

	// The ID of a core network attachment.
	CoreNetworkAttachmentId *string

	// The edge location for the network destination.
	EdgeLocation *string

	// The network function group name associated with the destination.
	NetworkFunctionGroupName *string

	// The ID of the resource.
	ResourceId *string

	// The resource type.
	ResourceType *string

	// The name of the segment.
	SegmentName *string

	// The ID of the transit gateway attachment.
	TransitGatewayAttachmentId *string

	noSmithyDocumentSerde
}

// Describes the telemetry information for a resource.
type NetworkTelemetry struct {

	// The Amazon Web Services account ID.
	AccountId *string

	// The address.
	Address *string

	// The Amazon Web Services Region.
	AwsRegion *string

	// The ID of a core network.
	CoreNetworkId *string

	// The connection health.
	Health *ConnectionHealth

	// The ARN of the gateway.
	RegisteredGatewayArn *string

	// The ARN of the resource.
	ResourceArn *string

	// The ID of the resource.
	ResourceId *string

	// The resource type.
	ResourceType *string

	noSmithyDocumentSerde
}

// The status of an Amazon Web Services Organization and the accounts within that
// organization.
type OrganizationStatus struct {

	// The current service-linked role (SLR) deployment status for an Amazon Web
	// Services Organization's accounts. This will be either SUCCEEDED or IN_PROGRESS .
	AccountStatusList []AccountStatus

	// The status of the organization's AWS service access. This will be ENABLED or
	// DISABLED .
	OrganizationAwsServiceAccessStatus *string

	// The ID of an Amazon Web Services Organization.
	OrganizationId *string

	// The status of the SLR deployment for the account. This will be either SUCCEEDED
	// or IN_PROGRESS .
	SLRDeploymentStatus *string

	noSmithyDocumentSerde
}

// Describes a path component.
type PathComponent struct {

	// The destination CIDR block in the route table.
	DestinationCidrBlock *string

	// The resource.
	Resource *NetworkResourceSummary

	// The sequence number in the path. The destination is 0.
	Sequence *int32

	noSmithyDocumentSerde
}

// Describes a peering connection.
type Peering struct {

	// The ARN of a core network.
	CoreNetworkArn *string

	// The ID of the core network for the peering request.
	CoreNetworkId *string

	// The timestamp when the attachment peer was created.
	CreatedAt *time.Time

	// The edge location for the peer.
	EdgeLocation *string

	// Describes the error associated with the Connect peer request.
	LastModificationErrors []PeeringError

	// The ID of the account owner.
	OwnerAccountId *string

	// The ID of the peering attachment.
	PeeringId *string

	// The type of peering. This will be TRANSIT_GATEWAY .
	PeeringType PeeringType

	// The resource ARN of the peer.
	ResourceArn *string

	// The current state of the peering connection.
	State PeeringState

	// The list of key-value tags associated with the peering.
	Tags []Tag

	noSmithyDocumentSerde
}

// Describes an error associated with a peering request.
type PeeringError struct {

	// The error code for the peering request.
	Code PeeringErrorCode

	// The message associated with the error code .
	Message *string

	// Provides additional information about missing permissions for the peering error.
	MissingPermissionsContext *PermissionsErrorContext

	// The ID of the Peering request.
	RequestId *string

	// The ARN of the requested peering resource.
	ResourceArn *string

	noSmithyDocumentSerde
}

// Describes additional information about missing permissions.
type PermissionsErrorContext struct {

	// The missing permissions.
	MissingPermission *string

	noSmithyDocumentSerde
}

// Describes proposed changes to a network function group.
type ProposedNetworkFunctionGroupChange struct {

	// The proposed new attachment policy rule number for the network function group.
	AttachmentPolicyRuleNumber *int32

	// The proposed name change for the network function group name.
	NetworkFunctionGroupName *string

	// The list of proposed changes to the key-value tags associated with the network
	// function group.
	Tags []Tag

	noSmithyDocumentSerde
}

// Describes a proposed segment change. In some cases, the segment change must
// first be evaluated and accepted.
type ProposedSegmentChange struct {

	// The rule number in the policy document that applies to this change.
	AttachmentPolicyRuleNumber *int32

	// The name of the segment to change.
	SegmentName *string

	// The list of key-value tags that changed for the segment.
	Tags []Tag

	noSmithyDocumentSerde
}

// Describes a resource relationship.
type Relationship struct {

	// The ARN of the resource.
	From *string

	// The ARN of the resource.
	To *string

	noSmithyDocumentSerde
}

// Describes a route analysis.
type RouteAnalysis struct {

	// The destination.
	Destination *RouteAnalysisEndpointOptions

	// The forward path.
	ForwardPath *RouteAnalysisPath

	// The ID of the global network.
	GlobalNetworkId *string

	// Indicates whether to analyze the return path. The return path is not analyzed
	// if the forward path analysis does not succeed.
	IncludeReturnPath bool

	// The ID of the AWS account that created the route analysis.
	OwnerAccountId *string

	// The return path.
	ReturnPath *RouteAnalysisPath

	// The ID of the route analysis.
	RouteAnalysisId *string

	// The source.
	Source *RouteAnalysisEndpointOptions

	// The time that the analysis started.
	StartTimestamp *time.Time

	// The status of the route analysis.
	Status RouteAnalysisStatus

	// Indicates whether to include the location of middlebox appliances in the route
	// analysis.
	UseMiddleboxes bool

	noSmithyDocumentSerde
}

// Describes the status of an analysis at completion.
type RouteAnalysisCompletion struct {

	// The reason code. Available only if a connection is not found.
	//
	//   - BLACKHOLE_ROUTE_FOR_DESTINATION_FOUND - Found a black hole route with the
	//   destination CIDR block.
	//
	//   - CYCLIC_PATH_DETECTED - Found the same resource multiple times while
	//   traversing the path.
	//
	//   - INACTIVE_ROUTE_FOR_DESTINATION_FOUND - Found an inactive route with the
	//   destination CIDR block.
	//
	//   - MAX_HOPS_EXCEEDED - Analysis exceeded 64 hops without finding the
	//   destination.
	//
	//   - ROUTE_NOT_FOUND - Cannot find a route table with the destination CIDR block.
	//
	//   - TGW_ATTACH_ARN_NO_MATCH - Found an attachment, but not with the correct
	//   destination ARN.
	//
	//   - TGW_ATTACH_NOT_FOUND - Cannot find an attachment.
	//
	//   - TGW_ATTACH_NOT_IN_TGW - Found an attachment, but not to the correct transit
	//   gateway.
	//
	//   - TGW_ATTACH_STABLE_ROUTE_TABLE_NOT_FOUND - The state of the route table
	//   association is not associated.
	ReasonCode RouteAnalysisCompletionReasonCode

	// Additional information about the path. Available only if a connection is not
	// found.
	ReasonContext map[string]string

	// The result of the analysis. If the status is NOT_CONNECTED , check the reason
	// code.
	ResultCode RouteAnalysisCompletionResultCode

	noSmithyDocumentSerde
}

// Describes a source or a destination.
type RouteAnalysisEndpointOptions struct {

	// The IP address.
	IpAddress *string

	// The ARN of the transit gateway.
	TransitGatewayArn *string

	// The ARN of the transit gateway attachment.
	TransitGatewayAttachmentArn *string

	noSmithyDocumentSerde
}

// Describes a source or a destination.
type RouteAnalysisEndpointOptionsSpecification struct {

	// The IP address.
	IpAddress *string

	// The ARN of the transit gateway attachment.
	TransitGatewayAttachmentArn *string

	noSmithyDocumentSerde
}

// Describes a route analysis path.
type RouteAnalysisPath struct {

	// The status of the analysis at completion.
	CompletionStatus *RouteAnalysisCompletion

	// The route analysis path.
	Path []PathComponent

	noSmithyDocumentSerde
}

// Describes a route table.
type RouteTableIdentifier struct {

	// The route table identifier associated with the network function group.
	CoreNetworkNetworkFunctionGroup *CoreNetworkNetworkFunctionGroupIdentifier

	// The segment edge in a core network.
	CoreNetworkSegmentEdge *CoreNetworkSegmentEdgeIdentifier

	// The ARN of the transit gateway route table for the attachment request. For
	// example, "TransitGatewayRouteTableArn":
	// "arn:aws:ec2:us-west-2:123456789012:transit-gateway-route-table/tgw-rtb-9876543210123456"
	// .
	TransitGatewayRouteTableArn *string

	noSmithyDocumentSerde
}

// Describes the action that the service insertion will take for any segments
// associated with it.
type ServiceInsertionAction struct {

	// The action the service insertion takes for traffic. send-via sends east-west
	// traffic between attachments. send-to sends north-south traffic to the security
	// appliance, and then from that to either the Internet or to an on-premesis
	// location.
	Action SegmentActionServiceInsertion

	// Describes the mode packets take for the send-via action. This is not used when
	// the action is send-to . dual-hop packets traverse attachments in both the
	// source to the destination core network edges. This mode requires that an
	// inspection attachment must be present in all Regions of the service
	// insertion-enabled segments. For single-hop , packets traverse a single
	// intermediate inserted attachment. You can use EdgeOverride to specify a
	// specific edge to use.
	Mode SendViaMode

	// The list of network function groups and any edge overrides for the chosen
	// service insertion action. Used for both send-to or send-via .
	Via *Via

	// The list of destination segments if the service insertion action is send-via .
	WhenSentTo *WhenSentTo

	noSmithyDocumentSerde
}

// Describes the segments associated with the service insertion action.
type ServiceInsertionSegments struct {

	// The list of segments associated with the send-to action.
	SendTo []string

	// The list of segments associated with the send-via action.
	SendVia []string

	noSmithyDocumentSerde
}

// Describes a site.
type Site struct {

	// The date and time that the site was created.
	CreatedAt *time.Time

	// The description of the site.
	Description *string

	// The ID of the global network.
	GlobalNetworkId *string

	// The location of the site.
	Location *Location

	// The Amazon Resource Name (ARN) of the site.
	SiteArn *string

	// The ID of the site.
	SiteId *string

	// The state of the site.
	State SiteState

	// The tags for the site.
	Tags []Tag

	noSmithyDocumentSerde
}

// Creates a site-to-site VPN attachment.
type SiteToSiteVpnAttachment struct {

	// Provides details about a site-to-site VPN attachment.
	Attachment *Attachment

	// The ARN of the site-to-site VPN attachment.
	VpnConnectionArn *string

	noSmithyDocumentSerde
}

// Describes a tag.
type Tag struct {

	// The tag key.
	//
	// Constraints: Maximum length of 128 characters.
	Key *string

	// The tag value.
	//
	// Constraints: Maximum length of 256 characters.
	Value *string

	noSmithyDocumentSerde
}

// Describes a transit gateway Connect peer association.
type TransitGatewayConnectPeerAssociation struct {

	// The ID of the device.
	DeviceId *string

	// The ID of the global network.
	GlobalNetworkId *string

	// The ID of the link.
	LinkId *string

	// The state of the association.
	State TransitGatewayConnectPeerAssociationState

	// The Amazon Resource Name (ARN) of the transit gateway Connect peer.
	TransitGatewayConnectPeerArn *string

	noSmithyDocumentSerde
}

// Describes a transit gateway peering attachment.
type TransitGatewayPeering struct {

	// Describes a transit gateway peer connection.
	Peering *Peering

	// The ARN of the transit gateway.
	TransitGatewayArn *string

	// The ID of the transit gateway peering attachment.
	TransitGatewayPeeringAttachmentId *string

	noSmithyDocumentSerde
}

// Describes the registration of a transit gateway to a global network.
type TransitGatewayRegistration struct {

	// The ID of the global network.
	GlobalNetworkId *string

	// The state of the transit gateway registration.
	State *TransitGatewayRegistrationStateReason

	// The Amazon Resource Name (ARN) of the transit gateway.
	TransitGatewayArn *string

	noSmithyDocumentSerde
}

// Describes the status of a transit gateway registration.
type TransitGatewayRegistrationStateReason struct {

	// The code for the state reason.
	Code TransitGatewayRegistrationState

	// The message for the state reason.
	Message *string

	noSmithyDocumentSerde
}

// Describes a transit gateway route table attachment.
type TransitGatewayRouteTableAttachment struct {

	// Describes a core network attachment.
	Attachment *Attachment

	// The ID of the peering attachment.
	PeeringId *string

	// The ARN of the transit gateway attachment route table. For example,
	// "TransitGatewayRouteTableArn":
	// "arn:aws:ec2:us-west-2:123456789012:transit-gateway-route-table/tgw-rtb-9876543210123456"
	// .
	TransitGatewayRouteTableArn *string

	noSmithyDocumentSerde
}

// Describes a validation exception for a field.
type ValidationExceptionField struct {

	// The message for the field.
	//
	// This member is required.
	Message *string

	// The name of the field.
	//
	// This member is required.
	Name *string

	noSmithyDocumentSerde
}

// The list of network function groups and edge overrides for the service
// insertion action. Used for both the send-to and send-via actions.
type Via struct {

	// The list of network function groups associated with the service insertion
	// action.
	NetworkFunctionGroups []NetworkFunctionGroup

	// Describes any edge overrides. An edge override is a specific edge to be used
	// for traffic.
	WithEdgeOverrides []EdgeOverride

	noSmithyDocumentSerde
}

// Describes a VPC attachment.
type VpcAttachment struct {

	// Provides details about the VPC attachment.
	Attachment *Attachment

	// Provides details about the VPC attachment.
	Options *VpcOptions

	// The subnet ARNs.
	SubnetArns []string

	noSmithyDocumentSerde
}

// Describes the VPC options.
type VpcOptions struct {

	// Indicates whether appliance mode is supported. If enabled, traffic flow between
	// a source and destination use the same Availability Zone for the VPC attachment
	// for the lifetime of that flow. The default value is false .
	ApplianceModeSupport bool

	// Indicates whether IPv6 is supported.
	Ipv6Support bool

	noSmithyDocumentSerde
}

// Displays a list of the destination segments. Used only when the service
// insertion action is send-to .
type WhenSentTo struct {

	// The list of destination segments when the service insertion action is send-to .
	WhenSentToSegmentsList []string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
