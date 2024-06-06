// Code generated by smithy-go-codegen DO NOT EDIT.

package datasync

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/datasync/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns information about an DataSync agent, such as its name, service endpoint
// type, and status.
func (c *Client) DescribeAgent(ctx context.Context, params *DescribeAgentInput, optFns ...func(*Options)) (*DescribeAgentOutput, error) {
	if params == nil {
		params = &DescribeAgentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeAgent", params, optFns, c.addOperationDescribeAgentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeAgentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// DescribeAgent
type DescribeAgentInput struct {

	// Specifies the Amazon Resource Name (ARN) of the DataSync agent that you want
	// information about.
	//
	// This member is required.
	AgentArn *string

	noSmithyDocumentSerde
}

// DescribeAgentResponse
type DescribeAgentOutput struct {

	// The ARN of the agent.
	AgentArn *string

	// The time that the agent was [activated].
	//
	// [activated]: https://docs.aws.amazon.com/datasync/latest/userguide/activate-agent.html
	CreationTime *time.Time

	// The type of [service endpoint] that your agent is connected to.
	//
	// [service endpoint]: https://docs.aws.amazon.com/datasync/latest/userguide/choose-service-endpoint.html
	EndpointType types.EndpointType

	// The last time that the agent was communicating with the DataSync service.
	LastConnectionTime *time.Time

	// The name of the agent.
	Name *string

	// The platform-related details about the agent, such as the version number.
	Platform *types.Platform

	// The network configuration that the agent uses when connecting to a [VPC service endpoint].
	//
	// [VPC service endpoint]: https://docs.aws.amazon.com/datasync/latest/userguide/choose-service-endpoint.html#choose-service-endpoint-vpc
	PrivateLinkConfig *types.PrivateLinkConfig

	// The status of the agent.
	//
	//   - If the status is ONLINE , the agent is configured properly and ready to use.
	//
	//   - If the status is OFFLINE , the agent has been out of contact with DataSync
	//   for five minutes or longer. This can happen for a few reasons. For more
	//   information, see [What do I do if my agent is offline?]
	//
	// [What do I do if my agent is offline?]: https://docs.aws.amazon.com/datasync/latest/userguide/troubleshooting-datasync-agents.html#troubleshoot-agent-offline
	Status types.AgentStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeAgentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeAgent{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeAgent{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeAgent"); err != nil {
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
	if err = addOpDescribeAgentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAgent(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeAgent(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeAgent",
	}
}
