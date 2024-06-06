// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Modifies a snapshot schedule for a cluster.
func (c *Client) ModifyClusterSnapshotSchedule(ctx context.Context, params *ModifyClusterSnapshotScheduleInput, optFns ...func(*Options)) (*ModifyClusterSnapshotScheduleOutput, error) {
	if params == nil {
		params = &ModifyClusterSnapshotScheduleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyClusterSnapshotSchedule", params, optFns, c.addOperationModifyClusterSnapshotScheduleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyClusterSnapshotScheduleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyClusterSnapshotScheduleInput struct {

	// A unique identifier for the cluster whose snapshot schedule you want to modify.
	//
	// This member is required.
	ClusterIdentifier *string

	// A boolean to indicate whether to remove the assoiciation between the cluster
	// and the schedule.
	DisassociateSchedule *bool

	// A unique alphanumeric identifier for the schedule that you want to associate
	// with the cluster.
	ScheduleIdentifier *string

	noSmithyDocumentSerde
}

type ModifyClusterSnapshotScheduleOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationModifyClusterSnapshotScheduleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpModifyClusterSnapshotSchedule{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpModifyClusterSnapshotSchedule{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ModifyClusterSnapshotSchedule"); err != nil {
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
	if err = addOpModifyClusterSnapshotScheduleValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyClusterSnapshotSchedule(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opModifyClusterSnapshotSchedule(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ModifyClusterSnapshotSchedule",
	}
}
