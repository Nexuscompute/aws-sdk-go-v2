// Code generated by smithy-go-codegen DO NOT EDIT.

package scheduler

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/scheduler/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates the specified schedule.
func (c *Client) CreateSchedule(ctx context.Context, params *CreateScheduleInput, optFns ...func(*Options)) (*CreateScheduleOutput, error) {
	if params == nil {
		params = &CreateScheduleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateSchedule", params, optFns, c.addOperationCreateScheduleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateScheduleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateScheduleInput struct {

	// Allows you to configure a time window during which EventBridge Scheduler
	// invokes the schedule.
	//
	// This member is required.
	FlexibleTimeWindow *types.FlexibleTimeWindow

	// The name of the schedule that you are creating.
	//
	// This member is required.
	Name *string

	//  The expression that defines when the schedule runs. The following formats are
	// supported.
	//
	//   - at expression - at(yyyy-mm-ddThh:mm:ss)
	//
	//   - rate expression - rate(value unit)
	//
	//   - cron expression - cron(fields)
	//
	// You can use at expressions to create one-time schedules that invoke a target
	// once, at the time and in the time zone, that you specify. You can use rate and
	// cron expressions to create recurring schedules. Rate-based schedules are useful
	// when you want to invoke a target at regular intervals, such as every 15 minutes
	// or every five days. Cron-based schedules are useful when you want to invoke a
	// target periodically at a specific time, such as at 8:00 am (UTC+0) every 1st day
	// of the month.
	//
	// A cron expression consists of six fields separated by white spaces: (minutes
	// hours day_of_month month day_of_week year) .
	//
	// A rate expression consists of a value as a positive integer, and a unit with
	// the following options: minute | minutes | hour | hours | day | days
	//
	// For more information and examples, see [Schedule types on EventBridge Scheduler] in the EventBridge Scheduler User
	// Guide.
	//
	// [Schedule types on EventBridge Scheduler]: https://docs.aws.amazon.com/scheduler/latest/UserGuide/schedule-types.html
	//
	// This member is required.
	ScheduleExpression *string

	// The schedule's target.
	//
	// This member is required.
	Target *types.Target

	// Specifies the action that EventBridge Scheduler applies to the schedule after
	// the schedule completes invoking the target.
	ActionAfterCompletion types.ActionAfterCompletion

	//  Unique, case-sensitive identifier you provide to ensure the idempotency of the
	// request. If you do not specify a client token, EventBridge Scheduler uses a
	// randomly generated token for the request to ensure idempotency.
	ClientToken *string

	// The description you specify for the schedule.
	Description *string

	// The date, in UTC, before which the schedule can invoke its target. Depending on
	// the schedule's recurrence expression, invocations might stop on, or before, the
	// EndDate you specify. EventBridge Scheduler ignores EndDate for one-time
	// schedules.
	EndDate *time.Time

	// The name of the schedule group to associate with this schedule. If you omit
	// this, the default schedule group is used.
	GroupName *string

	// The Amazon Resource Name (ARN) for the customer managed KMS key that
	// EventBridge Scheduler will use to encrypt and decrypt your data.
	KmsKeyArn *string

	// The timezone in which the scheduling expression is evaluated.
	ScheduleExpressionTimezone *string

	// The date, in UTC, after which the schedule can begin invoking its target.
	// Depending on the schedule's recurrence expression, invocations might occur on,
	// or after, the StartDate you specify. EventBridge Scheduler ignores StartDate
	// for one-time schedules.
	StartDate *time.Time

	// Specifies whether the schedule is enabled or disabled.
	State types.ScheduleState

	noSmithyDocumentSerde
}

type CreateScheduleOutput struct {

	// The Amazon Resource Name (ARN) of the schedule.
	//
	// This member is required.
	ScheduleArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateScheduleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateSchedule{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateSchedule{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateSchedule"); err != nil {
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
	if err = addIdempotencyToken_opCreateScheduleMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateScheduleValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateSchedule(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateSchedule struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateSchedule) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateSchedule) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateScheduleInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateScheduleInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateScheduleMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateSchedule{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateSchedule(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateSchedule",
	}
}
