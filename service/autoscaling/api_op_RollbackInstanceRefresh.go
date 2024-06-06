// Code generated by smithy-go-codegen DO NOT EDIT.

package autoscaling

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Cancels an instance refresh that is in progress and rolls back any changes that
// it made. Amazon EC2 Auto Scaling replaces any instances that were replaced
// during the instance refresh. This restores your Auto Scaling group to the
// configuration that it was using before the start of the instance refresh.
//
// This operation is part of the [instance refresh feature] in Amazon EC2 Auto Scaling, which helps you
// update instances in your Auto Scaling group after you make configuration
// changes.
//
// A rollback is not supported in the following situations:
//
//   - There is no desired configuration specified for the instance refresh.
//
//   - The Auto Scaling group has a launch template that uses an Amazon Web
//     Services Systems Manager parameter instead of an AMI ID for the ImageId
//     property.
//
//   - The Auto Scaling group uses the launch template's $Latest or $Default
//     version.
//
// When you receive a successful response from this operation, Amazon EC2 Auto
// Scaling immediately begins replacing instances. You can check the status of this
// operation through the DescribeInstanceRefreshesAPI operation.
//
// [instance refresh feature]: https://docs.aws.amazon.com/autoscaling/ec2/userguide/asg-instance-refresh.html
func (c *Client) RollbackInstanceRefresh(ctx context.Context, params *RollbackInstanceRefreshInput, optFns ...func(*Options)) (*RollbackInstanceRefreshOutput, error) {
	if params == nil {
		params = &RollbackInstanceRefreshInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RollbackInstanceRefresh", params, optFns, c.addOperationRollbackInstanceRefreshMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RollbackInstanceRefreshOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RollbackInstanceRefreshInput struct {

	// The name of the Auto Scaling group.
	//
	// This member is required.
	AutoScalingGroupName *string

	noSmithyDocumentSerde
}

type RollbackInstanceRefreshOutput struct {

	// The instance refresh ID associated with the request. This is the unique ID
	// assigned to the instance refresh when it was started.
	InstanceRefreshId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRollbackInstanceRefreshMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpRollbackInstanceRefresh{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpRollbackInstanceRefresh{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "RollbackInstanceRefresh"); err != nil {
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
	if err = addOpRollbackInstanceRefreshValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRollbackInstanceRefresh(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRollbackInstanceRefresh(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "RollbackInstanceRefresh",
	}
}
