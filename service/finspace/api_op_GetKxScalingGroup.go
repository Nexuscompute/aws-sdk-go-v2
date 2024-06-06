// Code generated by smithy-go-codegen DO NOT EDIT.

package finspace

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/finspace/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Retrieves details of a scaling group.
func (c *Client) GetKxScalingGroup(ctx context.Context, params *GetKxScalingGroupInput, optFns ...func(*Options)) (*GetKxScalingGroupOutput, error) {
	if params == nil {
		params = &GetKxScalingGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetKxScalingGroup", params, optFns, c.addOperationGetKxScalingGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetKxScalingGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetKxScalingGroupInput struct {

	// A unique identifier for the kdb environment.
	//
	// This member is required.
	EnvironmentId *string

	// A unique identifier for the kdb scaling group.
	//
	// This member is required.
	ScalingGroupName *string

	noSmithyDocumentSerde
}

type GetKxScalingGroupOutput struct {

	// The identifier of the availability zones.
	AvailabilityZoneId *string

	//  The list of Managed kdb clusters that are currently active in the given
	// scaling group.
	Clusters []string

	//  The timestamp at which the scaling group was created in FinSpace. The value is
	// determined as epoch time in milliseconds. For example, the value for Monday,
	// November 1, 2021 12:00:00 PM UTC is specified as 1635768000000.
	CreatedTimestamp *time.Time

	//  The memory and CPU capabilities of the scaling group host on which FinSpace
	// Managed kdb clusters will be placed.
	//
	// It can have one of the following values:
	//
	//   - kx.sg.4xlarge – The host type with a configuration of 108 GiB memory and 16
	//   vCPUs.
	//
	//   - kx.sg.8xlarge – The host type with a configuration of 216 GiB memory and 32
	//   vCPUs.
	//
	//   - kx.sg.16xlarge – The host type with a configuration of 432 GiB memory and 64
	//   vCPUs.
	//
	//   - kx.sg.32xlarge – The host type with a configuration of 864 GiB memory and
	//   128 vCPUs.
	//
	//   - kx.sg1.16xlarge – The host type with a configuration of 1949 GiB memory and
	//   64 vCPUs.
	//
	//   - kx.sg1.24xlarge – The host type with a configuration of 2948 GiB memory and
	//   96 vCPUs.
	HostType *string

	//  The last time that the scaling group was updated in FinSpace. The value is
	// determined as epoch time in milliseconds. For example, the value for Monday,
	// November 1, 2021 12:00:00 PM UTC is specified as 1635768000000.
	LastModifiedTimestamp *time.Time

	//  The ARN identifier for the scaling group.
	ScalingGroupArn *string

	// A unique identifier for the kdb scaling group.
	ScalingGroupName *string

	// The status of scaling group.
	//
	//   - CREATING – The scaling group creation is in progress.
	//
	//   - CREATE_FAILED – The scaling group creation has failed.
	//
	//   - ACTIVE – The scaling group is active.
	//
	//   - UPDATING – The scaling group is in the process of being updated.
	//
	//   - UPDATE_FAILED – The update action failed.
	//
	//   - DELETING – The scaling group is in the process of being deleted.
	//
	//   - DELETE_FAILED – The system failed to delete the scaling group.
	//
	//   - DELETED – The scaling group is successfully deleted.
	Status types.KxScalingGroupStatus

	//  The error message when a failed state occurs.
	StatusReason *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetKxScalingGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetKxScalingGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetKxScalingGroup{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetKxScalingGroup"); err != nil {
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
	if err = addRestJsonContentTypeCustomization(stack); err != nil {
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
	if err = addOpGetKxScalingGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetKxScalingGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetKxScalingGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetKxScalingGroup",
	}
}
