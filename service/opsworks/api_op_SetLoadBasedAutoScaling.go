// Code generated by smithy-go-codegen DO NOT EDIT.

package opsworks

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/opsworks/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Specify the load-based auto scaling configuration for a specified layer. For
// more information, see [Managing Load with Time-based and Load-based Instances].
//
// To use load-based auto scaling, you must create a set of load-based auto
// scaling instances. Load-based auto scaling operates only on the instances from
// that set, so you must ensure that you have created enough instances to handle
// the maximum anticipated load.
//
// Required Permissions: To use this action, an IAM user must have a Manage
// permissions level for the stack, or an attached policy that explicitly grants
// permissions. For more information on user permissions, see [Managing User Permissions].
//
// [Managing Load with Time-based and Load-based Instances]: https://docs.aws.amazon.com/opsworks/latest/userguide/workinginstances-autoscaling.html
// [Managing User Permissions]: https://docs.aws.amazon.com/opsworks/latest/userguide/opsworks-security-users.html
func (c *Client) SetLoadBasedAutoScaling(ctx context.Context, params *SetLoadBasedAutoScalingInput, optFns ...func(*Options)) (*SetLoadBasedAutoScalingOutput, error) {
	if params == nil {
		params = &SetLoadBasedAutoScalingInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SetLoadBasedAutoScaling", params, optFns, c.addOperationSetLoadBasedAutoScalingMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SetLoadBasedAutoScalingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SetLoadBasedAutoScalingInput struct {

	// The layer ID.
	//
	// This member is required.
	LayerId *string

	// An AutoScalingThresholds object with the downscaling threshold configuration.
	// If the load falls below these thresholds for a specified amount of time,
	// OpsWorks Stacks stops a specified number of instances.
	DownScaling *types.AutoScalingThresholds

	// Enables load-based auto scaling for the layer.
	Enable *bool

	// An AutoScalingThresholds object with the upscaling threshold configuration. If
	// the load exceeds these thresholds for a specified amount of time, OpsWorks
	// Stacks starts a specified number of instances.
	UpScaling *types.AutoScalingThresholds

	noSmithyDocumentSerde
}

type SetLoadBasedAutoScalingOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSetLoadBasedAutoScalingMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpSetLoadBasedAutoScaling{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpSetLoadBasedAutoScaling{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "SetLoadBasedAutoScaling"); err != nil {
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
	if err = addOpSetLoadBasedAutoScalingValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSetLoadBasedAutoScaling(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opSetLoadBasedAutoScaling(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "SetLoadBasedAutoScaling",
	}
}
