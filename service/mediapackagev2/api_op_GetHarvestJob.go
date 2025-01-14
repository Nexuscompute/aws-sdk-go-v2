// Code generated by smithy-go-codegen DO NOT EDIT.

package mediapackagev2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/mediapackagev2/types"
	"github.com/aws/smithy-go/middleware"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	smithywaiter "github.com/aws/smithy-go/waiter"
	jmespath "github.com/jmespath/go-jmespath"
	"time"
)

// Retrieves the details of a specific harvest job.
func (c *Client) GetHarvestJob(ctx context.Context, params *GetHarvestJobInput, optFns ...func(*Options)) (*GetHarvestJobOutput, error) {
	if params == nil {
		params = &GetHarvestJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetHarvestJob", params, optFns, c.addOperationGetHarvestJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetHarvestJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request object for retrieving a specific harvest job.
type GetHarvestJobInput struct {

	// The name of the channel group containing the channel associated with the
	// harvest job.
	//
	// This member is required.
	ChannelGroupName *string

	// The name of the channel associated with the harvest job.
	//
	// This member is required.
	ChannelName *string

	// The name of the harvest job to retrieve.
	//
	// This member is required.
	HarvestJobName *string

	// The name of the origin endpoint associated with the harvest job.
	//
	// This member is required.
	OriginEndpointName *string

	noSmithyDocumentSerde
}

// The response object containing the details of the requested harvest job.
type GetHarvestJobOutput struct {

	// The Amazon Resource Name (ARN) of the harvest job.
	//
	// This member is required.
	Arn *string

	// The name of the channel group containing the channel associated with the
	// harvest job.
	//
	// This member is required.
	ChannelGroupName *string

	// The name of the channel associated with the harvest job.
	//
	// This member is required.
	ChannelName *string

	// The date and time when the harvest job was created.
	//
	// This member is required.
	CreatedAt *time.Time

	// The S3 destination where the harvested content is being placed.
	//
	// This member is required.
	Destination *types.Destination

	// The name of the harvest job.
	//
	// This member is required.
	HarvestJobName *string

	// A list of manifests that are being or have been harvested.
	//
	// This member is required.
	HarvestedManifests *types.HarvestedManifests

	// The date and time when the harvest job was last modified.
	//
	// This member is required.
	ModifiedAt *time.Time

	// The name of the origin endpoint associated with the harvest job.
	//
	// This member is required.
	OriginEndpointName *string

	// The configuration for when the harvest job is scheduled to run, including start
	// and end times.
	//
	// This member is required.
	ScheduleConfiguration *types.HarvesterScheduleConfiguration

	// The current status of the harvest job (e.g., QUEUED, IN_PROGRESS, CANCELLED,
	// COMPLETED, FAILED).
	//
	// This member is required.
	Status types.HarvestJobStatus

	// The description of the harvest job, if provided.
	Description *string

	// The current version of the harvest job. Used for concurrency control.
	ETag *string

	// An error message if the harvest job encountered any issues.
	ErrorMessage *string

	// A collection of tags associated with the harvest job.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetHarvestJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetHarvestJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetHarvestJob{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetHarvestJob"); err != nil {
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
	if err = addOpGetHarvestJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetHarvestJob(options.Region), middleware.Before); err != nil {
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

// HarvestJobFinishedWaiterOptions are waiter options for HarvestJobFinishedWaiter
type HarvestJobFinishedWaiterOptions struct {

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
	// HarvestJobFinishedWaiter will use default minimum delay of 2 seconds. Note that
	// MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or
	// set to zero, HarvestJobFinishedWaiter will use default max delay of 120 seconds.
	// Note that MaxDelay must resolve to value greater than or equal to the MinDelay.
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
	Retryable func(context.Context, *GetHarvestJobInput, *GetHarvestJobOutput, error) (bool, error)
}

// HarvestJobFinishedWaiter defines the waiters for HarvestJobFinished
type HarvestJobFinishedWaiter struct {
	client GetHarvestJobAPIClient

	options HarvestJobFinishedWaiterOptions
}

// NewHarvestJobFinishedWaiter constructs a HarvestJobFinishedWaiter.
func NewHarvestJobFinishedWaiter(client GetHarvestJobAPIClient, optFns ...func(*HarvestJobFinishedWaiterOptions)) *HarvestJobFinishedWaiter {
	options := HarvestJobFinishedWaiterOptions{}
	options.MinDelay = 2 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = harvestJobFinishedStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &HarvestJobFinishedWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for HarvestJobFinished waiter. The maxWaitDur is
// the maximum wait duration the waiter will wait. The maxWaitDur is required and
// must be greater than zero.
func (w *HarvestJobFinishedWaiter) Wait(ctx context.Context, params *GetHarvestJobInput, maxWaitDur time.Duration, optFns ...func(*HarvestJobFinishedWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for HarvestJobFinished waiter and
// returns the output of the successful operation. The maxWaitDur is the maximum
// wait duration the waiter will wait. The maxWaitDur is required and must be
// greater than zero.
func (w *HarvestJobFinishedWaiter) WaitForOutput(ctx context.Context, params *GetHarvestJobInput, maxWaitDur time.Duration, optFns ...func(*HarvestJobFinishedWaiterOptions)) (*GetHarvestJobOutput, error) {
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

		out, err := w.client.GetHarvestJob(ctx, params, func(o *Options) {
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
	return nil, fmt.Errorf("exceeded max wait time for HarvestJobFinished waiter")
}

func harvestJobFinishedStateRetryable(ctx context.Context, input *GetHarvestJobInput, output *GetHarvestJobOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("Status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "COMPLETED"
		value, ok := pathValue.(types.HarvestJobStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.HarvestJobStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("Status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "CANCELLED"
		value, ok := pathValue.(types.HarvestJobStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.HarvestJobStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("Status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "FAILED"
		value, ok := pathValue.(types.HarvestJobStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.HarvestJobStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("Status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "QUEUED"
		value, ok := pathValue.(types.HarvestJobStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.HarvestJobStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return true, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("Status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "IN_PROGRESS"
		value, ok := pathValue.(types.HarvestJobStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.HarvestJobStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return true, nil
		}
	}

	if err != nil {
		return false, err
	}
	return true, nil
}

// GetHarvestJobAPIClient is a client that implements the GetHarvestJob operation.
type GetHarvestJobAPIClient interface {
	GetHarvestJob(context.Context, *GetHarvestJobInput, ...func(*Options)) (*GetHarvestJobOutput, error)
}

var _ GetHarvestJobAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opGetHarvestJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetHarvestJob",
	}
}
