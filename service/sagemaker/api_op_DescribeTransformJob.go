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

// Returns information about a transform job.
func (c *Client) DescribeTransformJob(ctx context.Context, params *DescribeTransformJobInput, optFns ...func(*Options)) (*DescribeTransformJobOutput, error) {
	if params == nil {
		params = &DescribeTransformJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeTransformJob", params, optFns, c.addOperationDescribeTransformJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeTransformJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeTransformJobInput struct {

	// The name of the transform job that you want to view details of.
	//
	// This member is required.
	TransformJobName *string

	noSmithyDocumentSerde
}

type DescribeTransformJobOutput struct {

	// A timestamp that shows when the transform Job was created.
	//
	// This member is required.
	CreationTime *time.Time

	// The name of the model used in the transform job.
	//
	// This member is required.
	ModelName *string

	// Describes the dataset to be transformed and the Amazon S3 location where it is
	// stored.
	//
	// This member is required.
	TransformInput *types.TransformInput

	// The Amazon Resource Name (ARN) of the transform job.
	//
	// This member is required.
	TransformJobArn *string

	// The name of the transform job.
	//
	// This member is required.
	TransformJobName *string

	// The status of the transform job. If the transform job failed, the reason is
	// returned in the FailureReason field.
	//
	// This member is required.
	TransformJobStatus types.TransformJobStatus

	// Describes the resources, including ML instance types and ML instance count, to
	// use for the transform job.
	//
	// This member is required.
	TransformResources *types.TransformResources

	// The Amazon Resource Name (ARN) of the AutoML transform job.
	AutoMLJobArn *string

	// Specifies the number of records to include in a mini-batch for an HTTP
	// inference request. A record is a single unit of input data that inference can be
	// made on. For example, a single line in a CSV file is a record.
	//
	// To enable the batch strategy, you must set SplitType to Line , RecordIO , or
	// TFRecord .
	BatchStrategy types.BatchStrategy

	// Configuration to control how SageMaker captures inference data.
	DataCaptureConfig *types.BatchDataCaptureConfig

	// The data structure used to specify the data to be used for inference in a batch
	// transform job and to associate the data that is relevant to the prediction
	// results in the output. The input filter provided allows you to exclude input
	// data that is not needed for inference in a batch transform job. The output
	// filter provided allows you to include input data relevant to interpreting the
	// predictions in the output from the job. For more information, see [Associate Prediction Results with their Corresponding Input Records].
	//
	// [Associate Prediction Results with their Corresponding Input Records]: https://docs.aws.amazon.com/sagemaker/latest/dg/batch-transform-data-processing.html
	DataProcessing *types.DataProcessing

	// The environment variables to set in the Docker container. We support up to 16
	// key and values entries in the map.
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

	// If the transform job failed, FailureReason describes why it failed. A transform
	// job creates a log file, which includes error messages, and stores it as an
	// Amazon S3 object. For more information, see [Log Amazon SageMaker Events with Amazon CloudWatch].
	//
	// [Log Amazon SageMaker Events with Amazon CloudWatch]: https://docs.aws.amazon.com/sagemaker/latest/dg/logging-cloudwatch.html
	FailureReason *string

	// The Amazon Resource Name (ARN) of the Amazon SageMaker Ground Truth labeling
	// job that created the transform or training job.
	LabelingJobArn *string

	// The maximum number of parallel requests on each instance node that can be
	// launched in a transform job. The default value is 1.
	MaxConcurrentTransforms *int32

	// The maximum payload size, in MB, used in the transform job.
	MaxPayloadInMB *int32

	// The timeout and maximum number of retries for processing a transform job
	// invocation.
	ModelClientConfig *types.ModelClientConfig

	// Indicates when the transform job has been
	//
	// completed, or has stopped or failed. You are billed for the time interval
	// between this time and the value of TransformStartTime .
	TransformEndTime *time.Time

	// Identifies the Amazon S3 location where you want Amazon SageMaker to save the
	// results from the transform job.
	TransformOutput *types.TransformOutput

	// Indicates when the transform job starts on ML instances. You are billed for the
	// time interval between this time and the value of TransformEndTime .
	TransformStartTime *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeTransformJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeTransformJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeTransformJob{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeTransformJob"); err != nil {
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
	if err = addOpDescribeTransformJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeTransformJob(options.Region), middleware.Before); err != nil {
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

// TransformJobCompletedOrStoppedWaiterOptions are waiter options for
// TransformJobCompletedOrStoppedWaiter
type TransformJobCompletedOrStoppedWaiterOptions struct {

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
	// TransformJobCompletedOrStoppedWaiter will use default minimum delay of 60
	// seconds. Note that MinDelay must resolve to a value lesser than or equal to the
	// MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or
	// set to zero, TransformJobCompletedOrStoppedWaiter will use default max delay of
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
	Retryable func(context.Context, *DescribeTransformJobInput, *DescribeTransformJobOutput, error) (bool, error)
}

// TransformJobCompletedOrStoppedWaiter defines the waiters for
// TransformJobCompletedOrStopped
type TransformJobCompletedOrStoppedWaiter struct {
	client DescribeTransformJobAPIClient

	options TransformJobCompletedOrStoppedWaiterOptions
}

// NewTransformJobCompletedOrStoppedWaiter constructs a
// TransformJobCompletedOrStoppedWaiter.
func NewTransformJobCompletedOrStoppedWaiter(client DescribeTransformJobAPIClient, optFns ...func(*TransformJobCompletedOrStoppedWaiterOptions)) *TransformJobCompletedOrStoppedWaiter {
	options := TransformJobCompletedOrStoppedWaiterOptions{}
	options.MinDelay = 60 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = transformJobCompletedOrStoppedStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &TransformJobCompletedOrStoppedWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for TransformJobCompletedOrStopped waiter. The
// maxWaitDur is the maximum wait duration the waiter will wait. The maxWaitDur is
// required and must be greater than zero.
func (w *TransformJobCompletedOrStoppedWaiter) Wait(ctx context.Context, params *DescribeTransformJobInput, maxWaitDur time.Duration, optFns ...func(*TransformJobCompletedOrStoppedWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for TransformJobCompletedOrStopped
// waiter and returns the output of the successful operation. The maxWaitDur is the
// maximum wait duration the waiter will wait. The maxWaitDur is required and must
// be greater than zero.
func (w *TransformJobCompletedOrStoppedWaiter) WaitForOutput(ctx context.Context, params *DescribeTransformJobInput, maxWaitDur time.Duration, optFns ...func(*TransformJobCompletedOrStoppedWaiterOptions)) (*DescribeTransformJobOutput, error) {
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

		out, err := w.client.DescribeTransformJob(ctx, params, func(o *Options) {
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
	return nil, fmt.Errorf("exceeded max wait time for TransformJobCompletedOrStopped waiter")
}

func transformJobCompletedOrStoppedStateRetryable(ctx context.Context, input *DescribeTransformJobInput, output *DescribeTransformJobOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("TransformJobStatus", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "Completed"
		value, ok := pathValue.(types.TransformJobStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.TransformJobStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("TransformJobStatus", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "Stopped"
		value, ok := pathValue.(types.TransformJobStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.TransformJobStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("TransformJobStatus", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "Failed"
		value, ok := pathValue.(types.TransformJobStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.TransformJobStatus value, got %T", pathValue)
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

// DescribeTransformJobAPIClient is a client that implements the
// DescribeTransformJob operation.
type DescribeTransformJobAPIClient interface {
	DescribeTransformJob(context.Context, *DescribeTransformJobInput, ...func(*Options)) (*DescribeTransformJobOutput, error)
}

var _ DescribeTransformJobAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opDescribeTransformJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeTransformJob",
	}
}
