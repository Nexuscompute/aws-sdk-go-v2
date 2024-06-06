// Code generated by smithy-go-codegen DO NOT EDIT.

package batch

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/batch/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates an Batch compute environment.
func (c *Client) UpdateComputeEnvironment(ctx context.Context, params *UpdateComputeEnvironmentInput, optFns ...func(*Options)) (*UpdateComputeEnvironmentOutput, error) {
	if params == nil {
		params = &UpdateComputeEnvironmentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateComputeEnvironment", params, optFns, c.addOperationUpdateComputeEnvironmentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateComputeEnvironmentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for UpdateComputeEnvironment .
type UpdateComputeEnvironmentInput struct {

	// The name or full Amazon Resource Name (ARN) of the compute environment to
	// update.
	//
	// This member is required.
	ComputeEnvironment *string

	// Details of the compute resources managed by the compute environment. Required
	// for a managed compute environment. For more information, see [Compute Environments]in the Batch User
	// Guide.
	//
	// [Compute Environments]: https://docs.aws.amazon.com/batch/latest/userguide/compute_environments.html
	ComputeResources *types.ComputeResourceUpdate

	// The full Amazon Resource Name (ARN) of the IAM role that allows Batch to make
	// calls to other Amazon Web Services services on your behalf. For more
	// information, see [Batch service IAM role]in the Batch User Guide.
	//
	// If the compute environment has a service-linked role, it can't be changed to
	// use a regular IAM role. Likewise, if the compute environment has a regular IAM
	// role, it can't be changed to use a service-linked role. To update the parameters
	// for the compute environment that require an infrastructure update to change, the
	// AWSServiceRoleForBatch service-linked role must be used. For more information,
	// see [Updating compute environments]in the Batch User Guide.
	//
	// If your specified role has a path other than / , then you must either specify
	// the full role ARN (recommended) or prefix the role name with the path.
	//
	// Depending on how you created your Batch service role, its ARN might contain the
	// service-role path prefix. When you only specify the name of the service role,
	// Batch assumes that your ARN doesn't use the service-role path prefix. Because
	// of this, we recommend that you specify the full ARN of your service role when
	// you create compute environments.
	//
	// [Updating compute environments]: https://docs.aws.amazon.com/batch/latest/userguide/updating-compute-environments.html
	// [Batch service IAM role]: https://docs.aws.amazon.com/batch/latest/userguide/service_IAM_role.html
	ServiceRole *string

	// The state of the compute environment. Compute environments in the ENABLED state
	// can accept jobs from a queue and scale in or out automatically based on the
	// workload demand of its associated queues.
	//
	// If the state is ENABLED , then the Batch scheduler can attempt to place jobs
	// from an associated job queue on the compute resources within the environment. If
	// the compute environment is managed, then it can scale its instances out or in
	// automatically, based on the job queue demand.
	//
	// If the state is DISABLED , then the Batch scheduler doesn't attempt to place
	// jobs within the environment. Jobs in a STARTING or RUNNING state continue to
	// progress normally. Managed compute environments in the DISABLED state don't
	// scale out.
	//
	// Compute environments in a DISABLED state may continue to incur billing charges.
	// To prevent additional charges, turn off and then delete the compute environment.
	// For more information, see [State]in the Batch User Guide.
	//
	// When an instance is idle, the instance scales down to the minvCpus value.
	// However, the instance size doesn't change. For example, consider a c5.8xlarge
	// instance with a minvCpus value of 4 and a desiredvCpus value of 36 . This
	// instance doesn't scale down to a c5.large instance.
	//
	// [State]: https://docs.aws.amazon.com/batch/latest/userguide/compute_environment_parameters.html#compute_environment_state
	State types.CEState

	// The maximum number of vCPUs expected to be used for an unmanaged compute
	// environment. Don't specify this parameter for a managed compute environment.
	// This parameter is only used for fair share scheduling to reserve vCPU capacity
	// for new share identifiers. If this parameter isn't provided for a fair share job
	// queue, no vCPU capacity is reserved.
	UnmanagedvCpus *int32

	// Specifies the updated infrastructure update policy for the compute environment.
	// For more information about infrastructure updates, see [Updating compute environments]in the Batch User Guide.
	//
	// [Updating compute environments]: https://docs.aws.amazon.com/batch/latest/userguide/updating-compute-environments.html
	UpdatePolicy *types.UpdatePolicy

	noSmithyDocumentSerde
}

type UpdateComputeEnvironmentOutput struct {

	// The Amazon Resource Name (ARN) of the compute environment.
	ComputeEnvironmentArn *string

	// The name of the compute environment. It can be up to 128 characters long. It
	// can contain uppercase and lowercase letters, numbers, hyphens (-), and
	// underscores (_).
	ComputeEnvironmentName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateComputeEnvironmentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateComputeEnvironment{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateComputeEnvironment{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateComputeEnvironment"); err != nil {
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
	if err = addOpUpdateComputeEnvironmentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateComputeEnvironment(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateComputeEnvironment(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateComputeEnvironment",
	}
}
