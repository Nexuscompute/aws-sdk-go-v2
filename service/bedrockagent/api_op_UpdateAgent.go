// Code generated by smithy-go-codegen DO NOT EDIT.

package bedrockagent

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/bedrockagent/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the configuration of an agent.
func (c *Client) UpdateAgent(ctx context.Context, params *UpdateAgentInput, optFns ...func(*Options)) (*UpdateAgentOutput, error) {
	if params == nil {
		params = &UpdateAgentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateAgent", params, optFns, c.addOperationUpdateAgentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateAgentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateAgentInput struct {

	// The unique identifier of the agent.
	//
	// This member is required.
	AgentId *string

	// Specifies a new name for the agent.
	//
	// This member is required.
	AgentName *string

	// The Amazon Resource Name (ARN) of the IAM role with permissions to invoke API
	// operations on the agent.
	//
	// This member is required.
	AgentResourceRoleArn *string

	// The identifier for the model that you want to be used for orchestration by the
	// agent you create.
	//
	// The modelId to provide depends on the type of model or throughput that you use:
	//
	//   - If you use a base model, specify the model ID or its ARN. For a list of
	//   model IDs for base models, see [Amazon Bedrock base model IDs (on-demand throughput)]in the Amazon Bedrock User Guide.
	//
	//   - If you use an inference profile, specify the inference profile ID or its
	//   ARN. For a list of inference profile IDs, see [Supported Regions and models for cross-region inference]in the Amazon Bedrock User
	//   Guide.
	//
	//   - If you use a provisioned model, specify the ARN of the Provisioned
	//   Throughput. For more information, see [Run inference using a Provisioned Throughput]in the Amazon Bedrock User Guide.
	//
	//   - If you use a custom model, first purchase Provisioned Throughput for it.
	//   Then specify the ARN of the resulting provisioned model. For more information,
	//   see [Use a custom model in Amazon Bedrock]in the Amazon Bedrock User Guide.
	//
	//   - If you use an [imported model], specify the ARN of the imported model. You can get the
	//   model ARN from a successful call to [CreateModelImportJob]or from the Imported models page in the
	//   Amazon Bedrock console.
	//
	// [Run inference using a Provisioned Throughput]: https://docs.aws.amazon.com/bedrock/latest/userguide/prov-thru-use.html
	// [Use a custom model in Amazon Bedrock]: https://docs.aws.amazon.com/bedrock/latest/userguide/model-customization-use.html
	// [imported model]: https://docs.aws.amazon.com/bedrock/latest/userguide/model-customization-import-model.html
	// [CreateModelImportJob]: https://docs.aws.amazon.com/bedrock/latest/APIReference/API_CreateModelImportJob.html
	// [Supported Regions and models for cross-region inference]: https://docs.aws.amazon.com/bedrock/latest/userguide/cross-region-inference-support.html
	// [Amazon Bedrock base model IDs (on-demand throughput)]: https://docs.aws.amazon.com/bedrock/latest/userguide/model-ids.html#model-ids-arns
	//
	// This member is required.
	FoundationModel *string

	// The agent's collaboration role.
	AgentCollaboration types.AgentCollaboration

	//  Contains details of the custom orchestration configured for the agent.
	CustomOrchestration *types.CustomOrchestration

	// The Amazon Resource Name (ARN) of the KMS key with which to encrypt the agent.
	CustomerEncryptionKeyArn *string

	// Specifies a new description of the agent.
	Description *string

	// The unique Guardrail configuration assigned to the agent when it is updated.
	GuardrailConfiguration *types.GuardrailConfiguration

	// The number of seconds for which Amazon Bedrock keeps information about a user's
	// conversation with the agent.
	//
	// A user interaction remains active for the amount of time specified. If no
	// conversation occurs during this time, the session expires and Amazon Bedrock
	// deletes any data provided before the timeout.
	IdleSessionTTLInSeconds *int32

	// Specifies new instructions that tell the agent what it should do and how it
	// should interact with users.
	Instruction *string

	// Specifies the new memory configuration for the agent.
	MemoryConfiguration *types.MemoryConfiguration

	//  Specifies the type of orchestration strategy for the agent. This is set to
	// DEFAULT orchestration type, by default.
	OrchestrationType types.OrchestrationType

	// Contains configurations to override prompts in different parts of an agent
	// sequence. For more information, see [Advanced prompts].
	//
	// [Advanced prompts]: https://docs.aws.amazon.com/bedrock/latest/userguide/advanced-prompts.html
	PromptOverrideConfiguration *types.PromptOverrideConfiguration

	noSmithyDocumentSerde
}

type UpdateAgentOutput struct {

	// Contains details about the agent that was updated.
	//
	// This member is required.
	Agent *types.Agent

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateAgentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateAgent{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateAgent{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateAgent"); err != nil {
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
	if err = addOpUpdateAgentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateAgent(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateAgent(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateAgent",
	}
}
