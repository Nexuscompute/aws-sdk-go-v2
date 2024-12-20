// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Describes the Amazon SageMaker AI Studio Lifecycle Configuration.
func (c *Client) DescribeStudioLifecycleConfig(ctx context.Context, params *DescribeStudioLifecycleConfigInput, optFns ...func(*Options)) (*DescribeStudioLifecycleConfigOutput, error) {
	if params == nil {
		params = &DescribeStudioLifecycleConfigInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeStudioLifecycleConfig", params, optFns, c.addOperationDescribeStudioLifecycleConfigMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeStudioLifecycleConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeStudioLifecycleConfigInput struct {

	// The name of the Amazon SageMaker AI Studio Lifecycle Configuration to describe.
	//
	// This member is required.
	StudioLifecycleConfigName *string

	noSmithyDocumentSerde
}

type DescribeStudioLifecycleConfigOutput struct {

	// The creation time of the Amazon SageMaker AI Studio Lifecycle Configuration.
	CreationTime *time.Time

	// This value is equivalent to CreationTime because Amazon SageMaker AI Studio
	// Lifecycle Configurations are immutable.
	LastModifiedTime *time.Time

	// The App type that the Lifecycle Configuration is attached to.
	StudioLifecycleConfigAppType types.StudioLifecycleConfigAppType

	// The ARN of the Lifecycle Configuration to describe.
	StudioLifecycleConfigArn *string

	// The content of your Amazon SageMaker AI Studio Lifecycle Configuration script.
	StudioLifecycleConfigContent *string

	// The name of the Amazon SageMaker AI Studio Lifecycle Configuration that is
	// described.
	StudioLifecycleConfigName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeStudioLifecycleConfigMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeStudioLifecycleConfig{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeStudioLifecycleConfig{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeStudioLifecycleConfig"); err != nil {
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
	if err = addOpDescribeStudioLifecycleConfigValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeStudioLifecycleConfig(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeStudioLifecycleConfig(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeStudioLifecycleConfig",
	}
}
