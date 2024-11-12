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

// Creates a ContainerGroupDefinition that describes a set of containers for
// hosting your game server with Amazon GameLift managed containers hosting. An
// Amazon GameLift container group is similar to a container task or pod. Use
// container group definitions when you create a container fleet with CreateContainerFleet.
//
// A container group definition determines how Amazon GameLift deploys your
// containers to each instance in a container fleet. You can maintain multiple
// versions of a container group definition.
//
// There are two types of container groups:
//
//   - A game server container group has the containers that run your game server
//     application and supporting software. A game server container group can have
//     these container types:
//
//   - Game server container. This container runs your game server. You can define
//     one game server container in a game server container group.
//
//   - Support container. This container runs software in parallel with your game
//     server. You can define up to 8 support containers in a game server group.
//
// When building a game server container group definition, you can choose to
//
//	bundle your game server executable and all dependent software into a single game
//	server container. Alternatively, you can separate the software into one game
//	server container and one or more support containers.
//
// On a container fleet instance, a game server container group can be deployed
//
//	multiple times (depending on the compute resources of the instance). This means
//	that all containers in the container group are replicated together.
//
//	- A per-instance container group has containers for processes that aren't
//	replicated on a container fleet instance. This might include background
//	services, logging, test processes, or processes that need to persist
//	independently of the game server container group. When building a per-instance
//	container group, you can define up to 10 support containers.
//
// This operation requires Identity and Access Management (IAM) permissions to
// access container images in Amazon ECR repositories. See [IAM permissions for Amazon GameLift]for help setting the
// appropriate permissions.
//
// # Request options
//
// Use this operation to make the following types of requests. You can specify
// values for the minimum required parameters and customize optional values later.
//
//   - Create a game server container group definition. Provide the following
//     required parameter values:
//
//   - Name
//
//   - ContainerGroupType ( GAME_SERVER )
//
//   - OperatingSystem (omit to use default value)
//
//   - TotalMemoryLimitMebibytes (omit to use default value)
//
//   - TotalVcpuLimit (omit to use default value)
//
//   - At least one GameServerContainerDefinition
//
//   - ContainerName
//
//   - ImageUrl
//
//   - PortConfiguration
//
//   - ServerSdkVersion (omit to use default value)
//
//   - Create a per-instance container group definition. Provide the following
//     required parameter values:
//
//   - Name
//
//   - ContainerGroupType ( PER_INSTANCE )
//
//   - OperatingSystem (omit to use default value)
//
//   - TotalMemoryLimitMebibytes (omit to use default value)
//
//   - TotalVcpuLimit (omit to use default value)
//
//   - At least one SupportContainerDefinition
//
//   - ContainerName
//
//   - ImageUrl
//
// # Results
//
// If successful, this request creates a ContainerGroupDefinition resource and
// assigns a unique ARN value. You can update most properties of a container group
// definition by calling UpdateContainerGroupDefinition, and optionally save the update as a new version.
//
// [IAM permissions for Amazon GameLift]: https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-iam-policy-examples.html
func (c *Client) CreateContainerGroupDefinition(ctx context.Context, params *CreateContainerGroupDefinitionInput, optFns ...func(*Options)) (*CreateContainerGroupDefinitionOutput, error) {
	if params == nil {
		params = &CreateContainerGroupDefinitionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateContainerGroupDefinition", params, optFns, c.addOperationCreateContainerGroupDefinitionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateContainerGroupDefinitionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateContainerGroupDefinitionInput struct {

	// A descriptive identifier for the container group definition. The name value
	// must be unique in an Amazon Web Services Region.
	//
	// This member is required.
	Name *string

	// The platform that all containers in the group use. Containers in a group must
	// run on the same operating system.
	//
	// Default value: AMAZON_LINUX_2023
	//
	// Amazon Linux 2 (AL2) will reach end of support on 6/30/2025. See more details
	// in the [Amazon Linux 2 FAQs]. For game servers that are hosted on AL2 and use Amazon GameLift server
	// SDK 4.x, first update the game server build to server SDK 5.x, and then deploy
	// to AL2023 instances. See [Migrate to Amazon GameLift server SDK version 5.]
	//
	// [Amazon Linux 2 FAQs]: https://aws.amazon.com/amazon-linux-2/faqs/
	// [Migrate to Amazon GameLift server SDK version 5.]: https://docs.aws.amazon.com/gamelift/latest/developerguide/reference-serversdk5-migration.html
	//
	// This member is required.
	OperatingSystem types.ContainerOperatingSystem

	// The maximum amount of memory (in MiB) to allocate to the container group. All
	// containers in the group share this memory. If you specify memory limits for an
	// individual container, the total value must be greater than any individual
	// container's memory limit.
	//
	// Default value: 1024
	//
	// This member is required.
	TotalMemoryLimitMebibytes *int32

	// The maximum amount of vCPU units to allocate to the container group (1 vCPU is
	// equal to 1024 CPU units). All containers in the group share this memory. If you
	// specify vCPU limits for individual containers, the total value must be equal to
	// or greater than the sum of the CPU limits for all containers in the group.
	//
	// Default value: 1
	//
	// This member is required.
	TotalVcpuLimit *float64

	// The type of container group being defined. Container group type determines how
	// Amazon GameLift deploys the container group on each fleet instance.
	//
	// Default value: GAME_SERVER
	ContainerGroupType types.ContainerGroupType

	// The definition for the game server container in this group. Define a game
	// server container only when the container group type is GAME_SERVER . Game server
	// containers specify a container image with your game server build. You can pass
	// in your container definitions as a JSON file.
	GameServerContainerDefinition *types.GameServerContainerDefinitionInput

	// One or more definition for support containers in this group. You can define a
	// support container in any type of container group. You can pass in your container
	// definitions as a JSON file.
	SupportContainerDefinitions []types.SupportContainerDefinitionInput

	// A list of labels to assign to the container group definition resource. Tags are
	// developer-defined key-value pairs. Tagging Amazon Web Services resources are
	// useful for resource management, access management and cost allocation. For more
	// information, see [Tagging Amazon Web Services Resources]in the Amazon Web Services General Reference.
	//
	// [Tagging Amazon Web Services Resources]: https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html
	Tags []types.Tag

	// A description for the initial version of this container group definition.
	VersionDescription *string

	noSmithyDocumentSerde
}

type CreateContainerGroupDefinitionOutput struct {

	// The properties of the new container group definition resource. You can use this
	// resource to create a container fleet.
	ContainerGroupDefinition *types.ContainerGroupDefinition

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateContainerGroupDefinitionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateContainerGroupDefinition{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateContainerGroupDefinition{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateContainerGroupDefinition"); err != nil {
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
	if err = addOpCreateContainerGroupDefinitionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateContainerGroupDefinition(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateContainerGroupDefinition(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateContainerGroupDefinition",
	}
}
