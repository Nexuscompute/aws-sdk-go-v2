// Code generated by smithy-go-codegen DO NOT EDIT.

package lambda

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Adds a provisioned concurrency configuration to a function's alias or version.
func (c *Client) PutProvisionedConcurrencyConfig(ctx context.Context, params *PutProvisionedConcurrencyConfigInput, optFns ...func(*Options)) (*PutProvisionedConcurrencyConfigOutput, error) {
	if params == nil {
		params = &PutProvisionedConcurrencyConfigInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutProvisionedConcurrencyConfig", params, optFns, c.addOperationPutProvisionedConcurrencyConfigMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutProvisionedConcurrencyConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutProvisionedConcurrencyConfigInput struct {

	// The name or ARN of the Lambda function.
	//
	// Name formats
	//
	//   - Function name – my-function .
	//
	//   - Function ARN – arn:aws:lambda:us-west-2:123456789012:function:my-function .
	//
	//   - Partial ARN – 123456789012:function:my-function .
	//
	// The length constraint applies only to the full ARN. If you specify only the
	// function name, it is limited to 64 characters in length.
	//
	// This member is required.
	FunctionName *string

	// The amount of provisioned concurrency to allocate for the version or alias.
	//
	// This member is required.
	ProvisionedConcurrentExecutions *int32

	// The version number or alias name.
	//
	// This member is required.
	Qualifier *string

	noSmithyDocumentSerde
}

type PutProvisionedConcurrencyConfigOutput struct {

	// The amount of provisioned concurrency allocated. When a weighted alias is used
	// during linear and canary deployments, this value fluctuates depending on the
	// amount of concurrency that is provisioned for the function versions.
	AllocatedProvisionedConcurrentExecutions *int32

	// The amount of provisioned concurrency available.
	AvailableProvisionedConcurrentExecutions *int32

	// The date and time that a user last updated the configuration, in [ISO 8601 format].
	//
	// [ISO 8601 format]: https://www.iso.org/iso-8601-date-and-time-format.html
	LastModified *string

	// The amount of provisioned concurrency requested.
	RequestedProvisionedConcurrentExecutions *int32

	// The status of the allocation process.
	Status types.ProvisionedConcurrencyStatusEnum

	// For failed allocations, the reason that provisioned concurrency could not be
	// allocated.
	StatusReason *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutProvisionedConcurrencyConfigMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutProvisionedConcurrencyConfig{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutProvisionedConcurrencyConfig{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PutProvisionedConcurrencyConfig"); err != nil {
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
	if err = addOpPutProvisionedConcurrencyConfigValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutProvisionedConcurrencyConfig(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutProvisionedConcurrencyConfig(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PutProvisionedConcurrencyConfig",
	}
}
