// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/gamelift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

//	This operation has been expanded to use with the Amazon GameLift containers
//
// feature, which is currently in public preview.
//
// Creates a fleet of compute resources to host your game servers. Use this
// operation to set up the following types of fleets based on compute type:
//
// # Managed EC2 fleet
//
// An EC2 fleet is a set of Amazon Elastic Compute Cloud (Amazon EC2) instances.
// Your game server build is deployed to each fleet instance. Amazon GameLift
// manages the fleet's instances and controls the lifecycle of game server
// processes, which host game sessions for players. EC2 fleets can have instances
// in multiple locations. Each instance in the fleet is designated a Compute .
//
// To create an EC2 fleet, provide these required parameters:
//
//   - Either BuildId or ScriptId
//
//   - ComputeType set to EC2 (the default value)
//
//   - EC2InboundPermissions
//
//   - EC2InstanceType
//
//   - FleetType
//
//   - Name
//
//   - RuntimeConfiguration with at least one ServerProcesses configuration
//
// If successful, this operation creates a new fleet resource and places it in NEW
// status while Amazon GameLift initiates the [fleet creation workflow]. To debug your fleet, fetch logs,
// view performance metrics or other actions on the fleet, create a development
// fleet with port 22/3389 open. As a best practice, we recommend opening ports for
// remote access only when you need them and closing them when you're finished.
//
// When the fleet status is ACTIVE, you can adjust capacity settings and turn
// autoscaling on/off for each location.
//
// # Managed container fleet
//
// A container fleet is a set of Amazon Elastic Compute Cloud (Amazon EC2)
// instances. Your container architecture is deployed to each fleet instance based
// on the fleet configuration. Amazon GameLift manages the containers on each fleet
// instance and controls the lifecycle of game server processes, which host game
// sessions for players. Container fleets can have instances in multiple locations.
// Each container on an instance that runs game server processes is registered as a
// Compute .
//
// To create a container fleet, provide these required parameters:
//
//   - ComputeType set to CONTAINER
//
//   - ContainerGroupsConfiguration
//
//   - EC2InboundPermissions
//
//   - EC2InstanceType
//
//   - FleetType set to ON_DEMAND
//
//   - Name
//
//   - RuntimeConfiguration with at least one ServerProcesses configuration
//
// If successful, this operation creates a new fleet resource and places it in NEW
// status while Amazon GameLift initiates the [fleet creation workflow].
//
// When the fleet status is ACTIVE, you can adjust capacity settings and turn
// autoscaling on/off for each location.
//
// # Anywhere fleet
//
// An Anywhere fleet represents compute resources that are not owned or managed by
// Amazon GameLift. You might create an Anywhere fleet with your local machine for
// testing, or use one to host game servers with on-premises hardware or other game
// hosting solutions.
//
// To create an Anywhere fleet, provide these required parameters:
//
//   - ComputeType set to ANYWHERE
//
//   - Locations specifying a custom location
//
//   - Name
//
// If successful, this operation creates a new fleet resource and places it in
// ACTIVE status. You can register computes with a fleet in ACTIVE status.
//
// # Learn more
//
// [Setting up fleets]
//
// [Setting up a container fleet]
//
// [Debug fleet creation issues]
//
// [Multi-location fleets]
//
// [fleet creation workflow]: https://docs.aws.amazon.com/gamelift/latest/developerguide/fleets-creating-all.html#fleets-creation-workflow
// [Setting up a container fleet]: https://docs.aws.amazon.com/gamelift/latest/developerguide/containers-build-fleet.html
// [Multi-location fleets]: https://docs.aws.amazon.com/gamelift/latest/developerguide/fleets-intro.html
// [Debug fleet creation issues]: https://docs.aws.amazon.com/gamelift/latest/developerguide/fleets-creating-debug.html#fleets-creating-debug-creation
// [Setting up fleets]: https://docs.aws.amazon.com/gamelift/latest/developerguide/fleets-intro.html
func (c *Client) CreateFleet(ctx context.Context, params *CreateFleetInput, optFns ...func(*Options)) (*CreateFleetOutput, error) {
	if params == nil {
		params = &CreateFleetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateFleet", params, optFns, c.addOperationCreateFleetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateFleetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateFleetInput struct {

	// A descriptive label that is associated with a fleet. Fleet names do not need to
	// be unique.
	//
	// This member is required.
	Name *string

	// Amazon GameLift Anywhere configuration options.
	AnywhereConfiguration *types.AnywhereConfiguration

	// The unique identifier for a custom game server build to be deployed to a fleet
	// with compute type EC2 . You can use either the build ID or ARN. The build must
	// be uploaded to Amazon GameLift and in READY status. This fleet property can't
	// be changed after the fleet is created.
	BuildId *string

	// Prompts Amazon GameLift to generate a TLS/SSL certificate for the fleet. Amazon
	// GameLift uses the certificates to encrypt traffic between game clients and the
	// game servers running on Amazon GameLift. By default, the
	// CertificateConfiguration is DISABLED . You can't change this property after you
	// create the fleet.
	//
	// Certificate Manager (ACM) certificates expire after 13 months. Certificate
	// expiration can cause fleets to fail, preventing players from connecting to
	// instances in the fleet. We recommend you replace fleets before 13 months,
	// consider using fleet aliases for a smooth transition.
	//
	// ACM isn't available in all Amazon Web Services regions. A fleet creation
	// request with certificate generation enabled in an unsupported Region, fails with
	// a 4xx error. For more information about the supported Regions, see [Supported Regions]in the
	// Certificate Manager User Guide.
	//
	// [Supported Regions]: https://docs.aws.amazon.com/acm/latest/userguide/acm-regions.html
	CertificateConfiguration *types.CertificateConfiguration

	// The type of compute resource used to host your game servers.
	//
	//   - EC2 – The game server build is deployed to Amazon EC2 instances for cloud
	//   hosting. This is the default setting.
	//
	//   - CONTAINER – Container images with your game server build and supporting
	//   software are deployed to Amazon EC2 instances for cloud hosting. With this
	//   compute type, you must specify the ContainerGroupsConfiguration parameter.
	//
	//   - ANYWHERE – Game servers or container images with your game server and
	//   supporting software are deployed to compute resources that are provided and
	//   managed by you. With this compute type, you can also set the
	//   AnywhereConfiguration parameter.
	ComputeType types.ComputeType

	// The container groups to deploy to instances in the container fleet and other
	// fleet-level configuration settings. Use the CreateContainerGroupDefinitionaction to create container groups.
	// A container fleet must have exactly one replica container group, and can
	// optionally have one daemon container group. You can't change this property after
	// you create the fleet.
	ContainerGroupsConfiguration *types.ContainerGroupsConfiguration

	// A description for the fleet.
	Description *string

	// The IP address ranges and port settings that allow inbound traffic to access
	// game server processes and other processes on this fleet. Set this parameter for
	// EC2 and container fleets. You can leave this parameter empty when creating the
	// fleet, but you must call UpdateFleetPortSettingsto set it before players can connect to game sessions.
	// As a best practice, we recommend opening ports for remote access only when you
	// need them and closing them when you're finished. For Realtime Servers fleets,
	// Amazon GameLift automatically sets TCP and UDP ranges.
	//
	// To manage inbound access for a container fleet, set this parameter to the same
	// port numbers that you set for the fleet's connection port range. During the life
	// of the fleet, update this parameter to control which connection ports are open
	// to inbound traffic.
	EC2InboundPermissions []types.IpPermission

	// The Amazon GameLift-supported Amazon EC2 instance type to use with EC2 and
	// container fleets. Instance type determines the computing resources that will be
	// used to host your game servers, including CPU, memory, storage, and networking
	// capacity. See [Amazon Elastic Compute Cloud Instance Types]for detailed descriptions of Amazon EC2 instance types.
	//
	// [Amazon Elastic Compute Cloud Instance Types]: http://aws.amazon.com/ec2/instance-types/
	EC2InstanceType types.EC2InstanceType

	// Indicates whether to use On-Demand or Spot instances for this fleet. By
	// default, this property is set to ON_DEMAND . Learn more about when to use [On-Demand versus Spot Instances].
	// This fleet property can't be changed after the fleet is created.
	//
	// [On-Demand versus Spot Instances]: https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-ec2-instances.html#gamelift-ec2-instances-spot
	FleetType types.FleetType

	// A unique identifier for an IAM role with access permissions to other Amazon Web
	// Services services. Any application that runs on an instance in the
	// fleet--including install scripts, server processes, and other processes--can use
	// these permissions to interact with Amazon Web Services resources that you own or
	// have access to. For more information about using the role with your game server
	// builds, see [Communicate with other Amazon Web Services resources from your fleets]. This fleet property can't be changed after the fleet is created.
	//
	// [Communicate with other Amazon Web Services resources from your fleets]: https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-sdk-server-resources.html
	InstanceRoleArn *string

	// Prompts Amazon GameLift to generate a shared credentials file for the IAM role
	// that's defined in InstanceRoleArn . The shared credentials file is stored on
	// each fleet instance and refreshed as needed. Use shared credentials for
	// applications that are deployed along with the game server executable, if the
	// game server is integrated with server SDK version 5.x. For more information
	// about using shared credentials, see [Communicate with other Amazon Web Services resources from your fleets].
	//
	// [Communicate with other Amazon Web Services resources from your fleets]: https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-sdk-server-resources.html
	InstanceRoleCredentialsProvider types.InstanceRoleCredentialsProvider

	// A set of remote locations to deploy additional instances to and manage as a
	// multi-location fleet. Use this parameter when creating a fleet in Amazon Web
	// Services Regions that support multiple locations. You can add any Amazon Web
	// Services Region or Local Zone that's supported by Amazon GameLift. Provide a
	// list of one or more Amazon Web Services Region codes, such as us-west-2 , or
	// Local Zone names. When using this parameter, Amazon GameLift requires you to
	// include your home location in the request. For a list of supported Regions and
	// Local Zones, see [Amazon GameLift service locations]for managed hosting.
	//
	// [Amazon GameLift service locations]: https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-regions.html
	Locations []types.LocationConfiguration

	//  This parameter is no longer used. To specify where Amazon GameLift should
	// store log files once a server process shuts down, use the Amazon GameLift server
	// API ProcessReady() and specify one or more directory paths in logParameters .
	// For more information, see [Initialize the server process]in the Amazon GameLift Developer Guide.
	//
	// [Initialize the server process]: https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-sdk-server-api.html#gamelift-sdk-server-initialize
	LogPaths []string

	// The name of an Amazon Web Services CloudWatch metric group to add this fleet
	// to. A metric group is used to aggregate the metrics for multiple fleets. You can
	// specify an existing metric group name or set a new name to create a new metric
	// group. A fleet can be included in only one metric group at a time.
	MetricGroups []string

	// The status of termination protection for active game sessions on the fleet. By
	// default, this property is set to NoProtection . You can also set game session
	// protection for an individual game session by calling UpdateGameSession.
	//
	//   - NoProtection - Game sessions can be terminated during active gameplay as a
	//   result of a scale-down event.
	//
	//   - FullProtection - Game sessions in ACTIVE status cannot be terminated during
	//   a scale-down event.
	NewGameSessionProtectionPolicy types.ProtectionPolicy

	// Used when peering your Amazon GameLift fleet with a VPC, the unique identifier
	// for the Amazon Web Services account that owns the VPC. You can find your account
	// ID in the Amazon Web Services Management Console under account settings.
	PeerVpcAwsAccountId *string

	// A unique identifier for a VPC with resources to be accessed by your Amazon
	// GameLift fleet. The VPC must be in the same Region as your fleet. To look up a
	// VPC ID, use the [VPC Dashboard]in the Amazon Web Services Management Console. Learn more about
	// VPC peering in [VPC Peering with Amazon GameLift Fleets].
	//
	// [VPC Dashboard]: https://console.aws.amazon.com/vpc/
	// [VPC Peering with Amazon GameLift Fleets]: https://docs.aws.amazon.com/gamelift/latest/developerguide/vpc-peering.html
	PeerVpcId *string

	// A policy that limits the number of game sessions that an individual player can
	// create on instances in this fleet within a specified span of time.
	ResourceCreationLimitPolicy *types.ResourceCreationLimitPolicy

	// Instructions for how to launch and run server processes on the fleet. Set
	// runtime configuration for EC2 fleets and container fleets. For an Anywhere
	// fleets, set this parameter only if the fleet is running the Amazon GameLift
	// Agent. The runtime configuration defines one or more server process
	// configurations. Each server process identifies a game executable or Realtime
	// script file and the number of processes to run concurrently.
	//
	// This parameter replaces the parameters ServerLaunchPath and
	// ServerLaunchParameters , which are still supported for backward compatibility.
	RuntimeConfiguration *types.RuntimeConfiguration

	// The unique identifier for a Realtime configuration script to be deployed to a
	// fleet with compute type EC2 . You can use either the script ID or ARN. Scripts
	// must be uploaded to Amazon GameLift prior to creating the fleet. This fleet
	// property can't be changed after the fleet is created.
	ScriptId *string

	//  This parameter is no longer used. Specify server launch parameters using the
	// RuntimeConfiguration parameter. Requests that use this parameter instead
	// continue to be valid.
	ServerLaunchParameters *string

	//  This parameter is no longer used. Specify a server launch path using the
	// RuntimeConfiguration parameter. Requests that use this parameter instead
	// continue to be valid.
	ServerLaunchPath *string

	// A list of labels to assign to the new fleet resource. Tags are
	// developer-defined key-value pairs. Tagging Amazon Web Services resources are
	// useful for resource management, access management and cost allocation. For more
	// information, see [Tagging Amazon Web Services Resources]in the Amazon Web Services General Reference.
	//
	// [Tagging Amazon Web Services Resources]: https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateFleetOutput struct {

	// The properties for the new fleet, including the current status. All fleets are
	// placed in NEW status on creation.
	FleetAttributes *types.FleetAttributes

	// The fleet's locations and life-cycle status of each location. For new fleets,
	// the status of all locations is set to NEW . During fleet creation, Amazon
	// GameLift updates each location status as instances are deployed there and
	// prepared for game hosting. This list includes an entry for the fleet's home
	// Region. For fleets with no remote locations, only one entry, representing the
	// home Region, is returned.
	LocationStates []types.LocationState

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateFleetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateFleet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateFleet{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateFleet"); err != nil {
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
	if err = addOpCreateFleetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateFleet(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateFleet(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateFleet",
	}
}
