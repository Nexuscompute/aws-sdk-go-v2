// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Set the capacity of an Aurora Serverless v1 DB cluster to a specific value.
//
// Aurora Serverless v1 scales seamlessly based on the workload on the DB cluster.
// In some cases, the capacity might not scale fast enough to meet a sudden change
// in workload, such as a large number of new transactions. Call
// ModifyCurrentDBClusterCapacity to set the capacity explicitly.
//
// After this call sets the DB cluster capacity, Aurora Serverless v1 can
// automatically scale the DB cluster based on the cooldown period for scaling up
// and the cooldown period for scaling down.
//
// For more information about Aurora Serverless v1, see [Using Amazon Aurora Serverless v1] in the Amazon Aurora User
// Guide.
//
// If you call ModifyCurrentDBClusterCapacity with the default TimeoutAction ,
// connections that prevent Aurora Serverless v1 from finding a scaling point might
// be dropped. For more information about scaling points, see [Autoscaling for Aurora Serverless v1]in the Amazon Aurora
// User Guide.
//
// This operation only applies to Aurora Serverless v1 DB clusters.
//
// [Autoscaling for Aurora Serverless v1]: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-serverless.how-it-works.html#aurora-serverless.how-it-works.auto-scaling
// [Using Amazon Aurora Serverless v1]: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-serverless.html
func (c *Client) ModifyCurrentDBClusterCapacity(ctx context.Context, params *ModifyCurrentDBClusterCapacityInput, optFns ...func(*Options)) (*ModifyCurrentDBClusterCapacityOutput, error) {
	if params == nil {
		params = &ModifyCurrentDBClusterCapacityInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyCurrentDBClusterCapacity", params, optFns, c.addOperationModifyCurrentDBClusterCapacityMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyCurrentDBClusterCapacityOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyCurrentDBClusterCapacityInput struct {

	// The DB cluster identifier for the cluster being modified. This parameter isn't
	// case-sensitive.
	//
	// Constraints:
	//
	//   - Must match the identifier of an existing DB cluster.
	//
	// This member is required.
	DBClusterIdentifier *string

	// The DB cluster capacity.
	//
	// When you change the capacity of a paused Aurora Serverless v1 DB cluster, it
	// automatically resumes.
	//
	// Constraints:
	//
	//   - For Aurora MySQL, valid capacity values are 1 , 2 , 4 , 8 , 16 , 32 , 64 ,
	//   128 , and 256 .
	//
	//   - For Aurora PostgreSQL, valid capacity values are 2 , 4 , 8 , 16 , 32 , 64 ,
	//   192 , and 384 .
	Capacity *int32

	// The amount of time, in seconds, that Aurora Serverless v1 tries to find a
	// scaling point to perform seamless scaling before enforcing the timeout action.
	// The default is 300.
	//
	// Specify a value between 10 and 600 seconds.
	SecondsBeforeTimeout *int32

	// The action to take when the timeout is reached, either ForceApplyCapacityChange
	// or RollbackCapacityChange .
	//
	// ForceApplyCapacityChange , the default, sets the capacity to the specified value
	// as soon as possible.
	//
	// RollbackCapacityChange ignores the capacity change if a scaling point isn't
	// found in the timeout period.
	TimeoutAction *string

	noSmithyDocumentSerde
}

type ModifyCurrentDBClusterCapacityOutput struct {

	// The current capacity of the DB cluster.
	CurrentCapacity *int32

	// A user-supplied DB cluster identifier. This identifier is the unique key that
	// identifies a DB cluster.
	DBClusterIdentifier *string

	// A value that specifies the capacity that the DB cluster scales to next.
	PendingCapacity *int32

	// The number of seconds before a call to ModifyCurrentDBClusterCapacity times out.
	SecondsBeforeTimeout *int32

	// The timeout action of a call to ModifyCurrentDBClusterCapacity , either
	// ForceApplyCapacityChange or RollbackCapacityChange .
	TimeoutAction *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationModifyCurrentDBClusterCapacityMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpModifyCurrentDBClusterCapacity{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpModifyCurrentDBClusterCapacity{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ModifyCurrentDBClusterCapacity"); err != nil {
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
	if err = addOpModifyCurrentDBClusterCapacityValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyCurrentDBClusterCapacity(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opModifyCurrentDBClusterCapacity(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ModifyCurrentDBClusterCapacity",
	}
}
