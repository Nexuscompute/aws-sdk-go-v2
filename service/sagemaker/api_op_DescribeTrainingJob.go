// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	"errors"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	smithywaiter "github.com/aws/smithy-go/waiter"
	jmespath "github.com/jmespath/go-jmespath"
	"time"
)

// Returns information about a training job.
//
// Some of the attributes below only appear if the training job successfully
// starts. If the training job fails, TrainingJobStatus is Failed and, depending
// on the FailureReason , attributes like TrainingStartTime , TrainingTimeInSeconds
// , TrainingEndTime , and BillableTimeInSeconds may not be present in the
// response.
func (c *Client) DescribeTrainingJob(ctx context.Context, params *DescribeTrainingJobInput, optFns ...func(*Options)) (*DescribeTrainingJobOutput, error) {
	if params == nil {
		params = &DescribeTrainingJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeTrainingJob", params, optFns, c.addOperationDescribeTrainingJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeTrainingJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeTrainingJobInput struct {

	// The name of the training job.
	//
	// This member is required.
	TrainingJobName *string

	noSmithyDocumentSerde
}

type DescribeTrainingJobOutput struct {

	// Information about the algorithm used for training, and algorithm metadata.
	//
	// This member is required.
	AlgorithmSpecification *types.AlgorithmSpecification

	// A timestamp that indicates when the training job was created.
	//
	// This member is required.
	CreationTime *time.Time

	// Information about the Amazon S3 location that is configured for storing model
	// artifacts.
	//
	// This member is required.
	ModelArtifacts *types.ModelArtifacts

	// Resources, including ML compute instances and ML storage volumes, that are
	// configured for model training.
	//
	// This member is required.
	ResourceConfig *types.ResourceConfig

	//  Provides detailed information about the state of the training job. For
	// detailed information on the secondary status of the training job, see
	// StatusMessage under [SecondaryStatusTransition].
	//
	// SageMaker provides primary statuses and secondary statuses that apply to each
	// of them:
	//
	// InProgress
	//   - Starting - Starting the training job.
	//
	//   - Downloading - An optional stage for algorithms that support File training
	//   input mode. It indicates that data is being downloaded to the ML storage
	//   volumes.
	//
	//   - Training - Training is in progress.
	//
	//   - Interrupted - The job stopped because the managed spot training instances
	//   were interrupted.
	//
	//   - Uploading - Training is complete and the model artifacts are being uploaded
	//   to the S3 location.
	//
	// Completed
	//   - Completed - The training job has completed.
	//
	// Failed
	//   - Failed - The training job has failed. The reason for the failure is returned
	//   in the FailureReason field of DescribeTrainingJobResponse .
	//
	// Stopped
	//   - MaxRuntimeExceeded - The job stopped because it exceeded the maximum allowed
	//   runtime.
	//
	//   - MaxWaitTimeExceeded - The job stopped because it exceeded the maximum
	//   allowed wait time.
	//
	//   - Stopped - The training job has stopped.
	//
	// Stopping
	//   - Stopping - Stopping the training job.
	//
	// Valid values for SecondaryStatus are subject to change.
	//
	// We no longer support the following secondary statuses:
	//
	//   - LaunchingMLInstances
	//
	//   - PreparingTraining
	//
	//   - DownloadingTrainingImage
	//
	// [SecondaryStatusTransition]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_SecondaryStatusTransition.html
	//
	// This member is required.
	SecondaryStatus types.SecondaryStatus

	// Specifies a limit to how long a model training job can run. It also specifies
	// how long a managed Spot training job has to complete. When the job reaches the
	// time limit, SageMaker ends the training job. Use this API to cap model training
	// costs.
	//
	// To stop a job, SageMaker sends the algorithm the SIGTERM signal, which delays
	// job termination for 120 seconds. Algorithms can use this 120-second window to
	// save the model artifacts, so the results of training are not lost.
	//
	// This member is required.
	StoppingCondition *types.StoppingCondition

	// The Amazon Resource Name (ARN) of the training job.
	//
	// This member is required.
	TrainingJobArn *string

	//  Name of the model training job.
	//
	// This member is required.
	TrainingJobName *string

	// The status of the training job.
	//
	// SageMaker provides the following training job statuses:
	//
	//   - InProgress - The training is in progress.
	//
	//   - Completed - The training job has completed.
	//
	//   - Failed - The training job has failed. To see the reason for the failure, see
	//   the FailureReason field in the response to a DescribeTrainingJobResponse call.
	//
	//   - Stopping - The training job is stopping.
	//
	//   - Stopped - The training job has stopped.
	//
	// For more detailed information, see SecondaryStatus .
	//
	// This member is required.
	TrainingJobStatus types.TrainingJobStatus

	// The Amazon Resource Name (ARN) of an AutoML job.
	AutoMLJobArn *string

	// The billable time in seconds. Billable time refers to the absolute wall-clock
	// time.
	//
	// Multiply BillableTimeInSeconds by the number of instances ( InstanceCount ) in
	// your training cluster to get the total compute time SageMaker bills you if you
	// run distributed training. The formula is as follows: BillableTimeInSeconds *
	// InstanceCount .
	//
	// You can calculate the savings from using managed spot training using the
	// formula (1 - BillableTimeInSeconds / TrainingTimeInSeconds) * 100 . For example,
	// if BillableTimeInSeconds is 100 and TrainingTimeInSeconds is 500, the savings
	// is 80%.
	BillableTimeInSeconds *int32

	// Contains information about the output location for managed spot training
	// checkpoint data.
	CheckpointConfig *types.CheckpointConfig

	// Configuration information for the Amazon SageMaker Debugger hook parameters,
	// metric and tensor collections, and storage paths. To learn more about how to
	// configure the DebugHookConfig parameter, see [Use the SageMaker and Debugger Configuration API Operations to Create, Update, and Debug Your Training Job].
	//
	// [Use the SageMaker and Debugger Configuration API Operations to Create, Update, and Debug Your Training Job]: https://docs.aws.amazon.com/sagemaker/latest/dg/debugger-createtrainingjob-api.html
	DebugHookConfig *types.DebugHookConfig

	// Configuration information for Amazon SageMaker Debugger rules for debugging
	// output tensors.
	DebugRuleConfigurations []types.DebugRuleConfiguration

	// Evaluation status of Amazon SageMaker Debugger rules for debugging on a
	// training job.
	DebugRuleEvaluationStatuses []types.DebugRuleEvaluationStatus

	// To encrypt all communications between ML compute instances in distributed
	// training, choose True . Encryption provides greater security for distributed
	// training, but training might take longer. How long it takes depends on the
	// amount of communication between compute instances, especially if you use a deep
	// learning algorithms in distributed training.
	EnableInterContainerTrafficEncryption *bool

	// A Boolean indicating whether managed spot training is enabled ( True ) or not (
	// False ).
	EnableManagedSpotTraining *bool

	// If you want to allow inbound or outbound network calls, except for calls
	// between peers within a training cluster for distributed training, choose True .
	// If you enable network isolation for training jobs that are configured to use a
	// VPC, SageMaker downloads and uploads customer data and model artifacts through
	// the specified VPC, but the training container does not have network access.
	EnableNetworkIsolation *bool

	// The environment variables to set in the Docker container.
	Environment map[string]string

	// Associates a SageMaker job as a trial component with an experiment and trial.
	// Specified when you call the following APIs:
	//
	// [CreateProcessingJob]
	//
	// [CreateTrainingJob]
	//
	// [CreateTransformJob]
	//
	// [CreateTransformJob]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateTransformJob.html
	// [CreateTrainingJob]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateTrainingJob.html
	// [CreateProcessingJob]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateProcessingJob.html
	ExperimentConfig *types.ExperimentConfig

	// If the training job failed, the reason it failed.
	FailureReason *string

	// A collection of MetricData objects that specify the names, values, and dates
	// and times that the training algorithm emitted to Amazon CloudWatch.
	FinalMetricDataList []types.MetricData

	// Algorithm-specific parameters.
	HyperParameters map[string]string

	// Contains information about the infrastructure health check configuration for
	// the training job.
	InfraCheckConfig *types.InfraCheckConfig

	// An array of Channel objects that describes each data input channel.
	InputDataConfig []types.Channel

	// The Amazon Resource Name (ARN) of the SageMaker Ground Truth labeling job that
	// created the transform or training job.
	LabelingJobArn *string

	// A timestamp that indicates when the status of the training job was last
	// modified.
	LastModifiedTime *time.Time

	// The S3 path where model artifacts that you configured when creating the job are
	// stored. SageMaker creates subfolders for model artifacts.
	OutputDataConfig *types.OutputDataConfig

	// Configuration information for Amazon SageMaker Debugger system monitoring,
	// framework profiling, and storage paths.
	ProfilerConfig *types.ProfilerConfig

	// Configuration information for Amazon SageMaker Debugger rules for profiling
	// system and framework metrics.
	ProfilerRuleConfigurations []types.ProfilerRuleConfiguration

	// Evaluation status of Amazon SageMaker Debugger rules for profiling on a
	// training job.
	ProfilerRuleEvaluationStatuses []types.ProfilerRuleEvaluationStatus

	// Profiling status of a training job.
	ProfilingStatus types.ProfilingStatus

	// Configuration for remote debugging. To learn more about the remote debugging
	// functionality of SageMaker, see [Access a training container through Amazon Web Services Systems Manager (SSM) for remote debugging].
	//
	// [Access a training container through Amazon Web Services Systems Manager (SSM) for remote debugging]: https://docs.aws.amazon.com/sagemaker/latest/dg/train-remote-debugging.html
	RemoteDebugConfig *types.RemoteDebugConfig

	// The number of times to retry the job when the job fails due to an
	// InternalServerError .
	RetryStrategy *types.RetryStrategy

	// The Amazon Web Services Identity and Access Management (IAM) role configured
	// for the training job.
	RoleArn *string

	// A history of all of the secondary statuses that the training job has
	// transitioned through.
	SecondaryStatusTransitions []types.SecondaryStatusTransition

	// Configuration of storage locations for the Amazon SageMaker Debugger
	// TensorBoard output data.
	TensorBoardOutputConfig *types.TensorBoardOutputConfig

	// Indicates the time when the training job ends on training instances. You are
	// billed for the time interval between the value of TrainingStartTime and this
	// time. For successful jobs and stopped jobs, this is the time after model
	// artifacts are uploaded. For failed jobs, this is the time when SageMaker detects
	// a job failure.
	TrainingEndTime *time.Time

	// Indicates the time when the training job starts on training instances. You are
	// billed for the time interval between this time and the value of TrainingEndTime
	// . The start time in CloudWatch Logs might be later than this time. The
	// difference is due to the time it takes to download the training data and to the
	// size of the training container.
	TrainingStartTime *time.Time

	// The training time in seconds.
	TrainingTimeInSeconds *int32

	// The Amazon Resource Name (ARN) of the associated hyperparameter tuning job if
	// the training job was launched by a hyperparameter tuning job.
	TuningJobArn *string

	// A [VpcConfig] object that specifies the VPC that this training job has access to. For more
	// information, see [Protect Training Jobs by Using an Amazon Virtual Private Cloud].
	//
	// [VpcConfig]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_VpcConfig.html
	// [Protect Training Jobs by Using an Amazon Virtual Private Cloud]: https://docs.aws.amazon.com/sagemaker/latest/dg/train-vpc.html
	VpcConfig *types.VpcConfig

	// The status of the warm pool associated with the training job.
	WarmPoolStatus *types.WarmPoolStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeTrainingJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeTrainingJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeTrainingJob{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeTrainingJob"); err != nil {
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
	if err = addOpDescribeTrainingJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeTrainingJob(options.Region), middleware.Before); err != nil {
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

// TrainingJobCompletedOrStoppedWaiterOptions are waiter options for
// TrainingJobCompletedOrStoppedWaiter
type TrainingJobCompletedOrStoppedWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	//
	// Passing options here is functionally equivalent to passing values to this
	// config's ClientOptions field that extend the inner client's APIOptions directly.
	APIOptions []func(*middleware.Stack) error

	// Functional options to be passed to all operations invoked by this client.
	//
	// Function values that modify the inner APIOptions are applied after the waiter
	// config's own APIOptions modifiers.
	ClientOptions []func(*Options)

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// TrainingJobCompletedOrStoppedWaiter will use default minimum delay of 120
	// seconds. Note that MinDelay must resolve to a value lesser than or equal to the
	// MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or
	// set to zero, TrainingJobCompletedOrStoppedWaiter will use default max delay of
	// 120 seconds. Note that MaxDelay must resolve to value greater than or equal to
	// the MinDelay.
	MaxDelay time.Duration

	// LogWaitAttempts is used to enable logging for waiter retry attempts
	LogWaitAttempts bool

	// Retryable is function that can be used to override the service defined
	// waiter-behavior based on operation output, or returned error. This function is
	// used by the waiter to decide if a state is retryable or a terminal state.
	//
	// By default service-modeled logic will populate this option. This option can
	// thus be used to define a custom waiter state with fall-back to service-modeled
	// waiter state mutators.The function returns an error in case of a failure state.
	// In case of retry state, this function returns a bool value of true and nil
	// error, while in case of success it returns a bool value of false and nil error.
	Retryable func(context.Context, *DescribeTrainingJobInput, *DescribeTrainingJobOutput, error) (bool, error)
}

// TrainingJobCompletedOrStoppedWaiter defines the waiters for
// TrainingJobCompletedOrStopped
type TrainingJobCompletedOrStoppedWaiter struct {
	client DescribeTrainingJobAPIClient

	options TrainingJobCompletedOrStoppedWaiterOptions
}

// NewTrainingJobCompletedOrStoppedWaiter constructs a
// TrainingJobCompletedOrStoppedWaiter.
func NewTrainingJobCompletedOrStoppedWaiter(client DescribeTrainingJobAPIClient, optFns ...func(*TrainingJobCompletedOrStoppedWaiterOptions)) *TrainingJobCompletedOrStoppedWaiter {
	options := TrainingJobCompletedOrStoppedWaiterOptions{}
	options.MinDelay = 120 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = trainingJobCompletedOrStoppedStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &TrainingJobCompletedOrStoppedWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for TrainingJobCompletedOrStopped waiter. The
// maxWaitDur is the maximum wait duration the waiter will wait. The maxWaitDur is
// required and must be greater than zero.
func (w *TrainingJobCompletedOrStoppedWaiter) Wait(ctx context.Context, params *DescribeTrainingJobInput, maxWaitDur time.Duration, optFns ...func(*TrainingJobCompletedOrStoppedWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for TrainingJobCompletedOrStopped
// waiter and returns the output of the successful operation. The maxWaitDur is the
// maximum wait duration the waiter will wait. The maxWaitDur is required and must
// be greater than zero.
func (w *TrainingJobCompletedOrStoppedWaiter) WaitForOutput(ctx context.Context, params *DescribeTrainingJobInput, maxWaitDur time.Duration, optFns ...func(*TrainingJobCompletedOrStoppedWaiterOptions)) (*DescribeTrainingJobOutput, error) {
	if maxWaitDur <= 0 {
		return nil, fmt.Errorf("maximum wait time for waiter must be greater than zero")
	}

	options := w.options
	for _, fn := range optFns {
		fn(&options)
	}

	if options.MaxDelay <= 0 {
		options.MaxDelay = 120 * time.Second
	}

	if options.MinDelay > options.MaxDelay {
		return nil, fmt.Errorf("minimum waiter delay %v must be lesser than or equal to maximum waiter delay of %v.", options.MinDelay, options.MaxDelay)
	}

	ctx, cancelFn := context.WithTimeout(ctx, maxWaitDur)
	defer cancelFn()

	logger := smithywaiter.Logger{}
	remainingTime := maxWaitDur

	var attempt int64
	for {

		attempt++
		apiOptions := options.APIOptions
		start := time.Now()

		if options.LogWaitAttempts {
			logger.Attempt = attempt
			apiOptions = append([]func(*middleware.Stack) error{}, options.APIOptions...)
			apiOptions = append(apiOptions, logger.AddLogger)
		}

		out, err := w.client.DescribeTrainingJob(ctx, params, func(o *Options) {
			baseOpts := []func(*Options){
				addIsWaiterUserAgent,
			}
			o.APIOptions = append(o.APIOptions, apiOptions...)
			for _, opt := range baseOpts {
				opt(o)
			}
			for _, opt := range options.ClientOptions {
				opt(o)
			}
		})

		retryable, err := options.Retryable(ctx, params, out, err)
		if err != nil {
			return nil, err
		}
		if !retryable {
			return out, nil
		}

		remainingTime -= time.Since(start)
		if remainingTime < options.MinDelay || remainingTime <= 0 {
			break
		}

		// compute exponential backoff between waiter retries
		delay, err := smithywaiter.ComputeDelay(
			attempt, options.MinDelay, options.MaxDelay, remainingTime,
		)
		if err != nil {
			return nil, fmt.Errorf("error computing waiter delay, %w", err)
		}

		remainingTime -= delay
		// sleep for the delay amount before invoking a request
		if err := smithytime.SleepWithContext(ctx, delay); err != nil {
			return nil, fmt.Errorf("request cancelled while waiting, %w", err)
		}
	}
	return nil, fmt.Errorf("exceeded max wait time for TrainingJobCompletedOrStopped waiter")
}

func trainingJobCompletedOrStoppedStateRetryable(ctx context.Context, input *DescribeTrainingJobInput, output *DescribeTrainingJobOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("TrainingJobStatus", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "Completed"
		value, ok := pathValue.(types.TrainingJobStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.TrainingJobStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("TrainingJobStatus", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "Stopped"
		value, ok := pathValue.(types.TrainingJobStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.TrainingJobStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("TrainingJobStatus", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "Failed"
		value, ok := pathValue.(types.TrainingJobStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.TrainingJobStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}

	if err != nil {
		var apiErr smithy.APIError
		ok := errors.As(err, &apiErr)
		if !ok {
			return false, fmt.Errorf("expected err to be of type smithy.APIError, got %w", err)
		}

		if "ValidationException" == apiErr.ErrorCode() {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}

	if err != nil {
		return false, err
	}
	return true, nil
}

// DescribeTrainingJobAPIClient is a client that implements the
// DescribeTrainingJob operation.
type DescribeTrainingJobAPIClient interface {
	DescribeTrainingJob(context.Context, *DescribeTrainingJobInput, ...func(*Options)) (*DescribeTrainingJobOutput, error)
}

var _ DescribeTrainingJobAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opDescribeTrainingJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeTrainingJob",
	}
}
