// Code generated by smithy-go-codegen DO NOT EDIT.

package bedrock

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/bedrock/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Gets details about a guardrail. If you don't specify a version, the response
// returns details for the DRAFT version.
func (c *Client) GetGuardrail(ctx context.Context, params *GetGuardrailInput, optFns ...func(*Options)) (*GetGuardrailOutput, error) {
	if params == nil {
		params = &GetGuardrailInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetGuardrail", params, optFns, c.addOperationGetGuardrailMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetGuardrailOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetGuardrailInput struct {

	// The unique identifier of the guardrail for which to get details.
	//
	// This member is required.
	GuardrailIdentifier *string

	// The version of the guardrail for which to get details. If you don't specify a
	// version, the response returns details for the DRAFT version.
	GuardrailVersion *string

	noSmithyDocumentSerde
}

type GetGuardrailOutput struct {

	// The message that the guardrail returns when it blocks a prompt.
	//
	// This member is required.
	BlockedInputMessaging *string

	// The message that the guardrail returns when it blocks a model response.
	//
	// This member is required.
	BlockedOutputsMessaging *string

	// The date and time at which the guardrail was created.
	//
	// This member is required.
	CreatedAt *time.Time

	// The ARN of the guardrail that was created.
	//
	// This member is required.
	GuardrailArn *string

	// The unique identifier of the guardrail.
	//
	// This member is required.
	GuardrailId *string

	// The name of the guardrail.
	//
	// This member is required.
	Name *string

	// The status of the guardrail.
	//
	// This member is required.
	Status types.GuardrailStatus

	// The date and time at which the guardrail was updated.
	//
	// This member is required.
	UpdatedAt *time.Time

	// The version of the guardrail.
	//
	// This member is required.
	Version *string

	// The content policy that was configured for the guardrail.
	ContentPolicy *types.GuardrailContentPolicy

	// The description of the guardrail.
	Description *string

	// Appears if the status of the guardrail is FAILED . A list of recommendations to
	// carry out before retrying the request.
	FailureRecommendations []string

	// The ARN of the KMS key that encrypts the guardrail.
	KmsKeyArn *string

	// The sensitive information policy that was configured for the guardrail.
	SensitiveInformationPolicy *types.GuardrailSensitiveInformationPolicy

	// Appears if the status is FAILED . A list of reasons for why the guardrail failed
	// to be created, updated, versioned, or deleted.
	StatusReasons []string

	// The topic policy that was configured for the guardrail.
	TopicPolicy *types.GuardrailTopicPolicy

	// The word policy that was configured for the guardrail.
	WordPolicy *types.GuardrailWordPolicy

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetGuardrailMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetGuardrail{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetGuardrail{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetGuardrail"); err != nil {
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
	if err = addOpGetGuardrailValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetGuardrail(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetGuardrail(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetGuardrail",
	}
}
