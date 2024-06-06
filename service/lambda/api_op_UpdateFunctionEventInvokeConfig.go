// Code generated by smithy-go-codegen DO NOT EDIT.

package lambda

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Updates the configuration for asynchronous invocation for a function, version,
// or alias.
//
// To configure options for asynchronous invocation, use PutFunctionEventInvokeConfig.
func (c *Client) UpdateFunctionEventInvokeConfig(ctx context.Context, params *UpdateFunctionEventInvokeConfigInput, optFns ...func(*Options)) (*UpdateFunctionEventInvokeConfigOutput, error) {
	if params == nil {
		params = &UpdateFunctionEventInvokeConfigInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateFunctionEventInvokeConfig", params, optFns, c.addOperationUpdateFunctionEventInvokeConfigMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateFunctionEventInvokeConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateFunctionEventInvokeConfigInput struct {

	// The name or ARN of the Lambda function, version, or alias.
	//
	// Name formats
	//
	//   - Function name - my-function (name-only), my-function:v1 (with alias).
	//
	//   - Function ARN - arn:aws:lambda:us-west-2:123456789012:function:my-function .
	//
	//   - Partial ARN - 123456789012:function:my-function .
	//
	// You can append a version number or alias to any of the formats. The length
	// constraint applies only to the full ARN. If you specify only the function name,
	// it is limited to 64 characters in length.
	//
	// This member is required.
	FunctionName *string

	// A destination for events after they have been sent to a function for processing.
	//
	// Destinations
	//
	//   - Function - The Amazon Resource Name (ARN) of a Lambda function.
	//
	//   - Queue - The ARN of a standard SQS queue.
	//
	//   - Topic - The ARN of a standard SNS topic.
	//
	//   - Event Bus - The ARN of an Amazon EventBridge event bus.
	DestinationConfig *types.DestinationConfig

	// The maximum age of a request that Lambda sends to a function for processing.
	MaximumEventAgeInSeconds *int32

	// The maximum number of times to retry when the function returns an error.
	MaximumRetryAttempts *int32

	// A version number or alias name.
	Qualifier *string

	noSmithyDocumentSerde
}

type UpdateFunctionEventInvokeConfigOutput struct {

	// A destination for events after they have been sent to a function for processing.
	//
	// Destinations
	//
	//   - Function - The Amazon Resource Name (ARN) of a Lambda function.
	//
	//   - Queue - The ARN of a standard SQS queue.
	//
	//   - Topic - The ARN of a standard SNS topic.
	//
	//   - Event Bus - The ARN of an Amazon EventBridge event bus.
	DestinationConfig *types.DestinationConfig

	// The Amazon Resource Name (ARN) of the function.
	FunctionArn *string

	// The date and time that the configuration was last updated.
	LastModified *time.Time

	// The maximum age of a request that Lambda sends to a function for processing.
	MaximumEventAgeInSeconds *int32

	// The maximum number of times to retry when the function returns an error.
	MaximumRetryAttempts *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateFunctionEventInvokeConfigMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateFunctionEventInvokeConfig{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateFunctionEventInvokeConfig{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateFunctionEventInvokeConfig"); err != nil {
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
	if err = addOpUpdateFunctionEventInvokeConfigValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateFunctionEventInvokeConfig(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateFunctionEventInvokeConfig(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateFunctionEventInvokeConfig",
	}
}
