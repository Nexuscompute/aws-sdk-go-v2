// Code generated by smithy-go-codegen DO NOT EDIT.

package datasync

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/datasync/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Starts an DataSync transfer task. For each task, you can only run one task
// execution at a time.
//
// There are several steps to a task execution. For more information, see [Task execution statuses].
//
// If you're planning to transfer data to or from an Amazon S3 location, review [how DataSync can affect your S3 request charges]
// and the [DataSync pricing page]before you begin.
//
// [Task execution statuses]: https://docs.aws.amazon.com/datasync/latest/userguide/working-with-task-executions.html#understand-task-execution-statuses
// [how DataSync can affect your S3 request charges]: https://docs.aws.amazon.com/datasync/latest/userguide/create-s3-location.html#create-s3-location-s3-requests
// [DataSync pricing page]: http://aws.amazon.com/datasync/pricing/
func (c *Client) StartTaskExecution(ctx context.Context, params *StartTaskExecutionInput, optFns ...func(*Options)) (*StartTaskExecutionOutput, error) {
	if params == nil {
		params = &StartTaskExecutionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartTaskExecution", params, optFns, c.addOperationStartTaskExecutionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartTaskExecutionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// StartTaskExecutionRequest
type StartTaskExecutionInput struct {

	// Specifies the Amazon Resource Name (ARN) of the task that you want to start.
	//
	// This member is required.
	TaskArn *string

	// Specifies a list of filter rules that determines which files to exclude from a
	// task. The list contains a single filter string that consists of the patterns to
	// exclude. The patterns are delimited by "|" (that is, a pipe), for example,
	// "/folder1|/folder2" .
	Excludes []types.FilterRule

	// Specifies a list of filter rules that determines which files to include when
	// running a task. The pattern should contain a single filter string that consists
	// of the patterns to include. The patterns are delimited by "|" (that is, a pipe),
	// for example, "/folder1|/folder2" .
	Includes []types.FilterRule

	// Configures a manifest, which is a list of files or objects that you want
	// DataSync to transfer. For more information and configuration examples, see [Specifying what DataSync transfers by using a manifest].
	//
	// When using this parameter, your caller identity (the role that you're using
	// DataSync with) must have the iam:PassRole permission. The [AWSDataSyncFullAccess] policy includes this
	// permission.
	//
	// To remove a manifest configuration, specify this parameter with an empty value.
	//
	// [AWSDataSyncFullAccess]: https://docs.aws.amazon.com/datasync/latest/userguide/security-iam-awsmanpol.html#security-iam-awsmanpol-awsdatasyncfullaccess
	// [Specifying what DataSync transfers by using a manifest]: https://docs.aws.amazon.com/datasync/latest/userguide/transferring-with-manifest.html
	ManifestConfig *types.ManifestConfig

	// Indicates how your transfer task is configured. These options include how
	// DataSync handles files, objects, and their associated metadata during your
	// transfer. You also can specify how to verify data integrity, set bandwidth
	// limits for your task, among other options.
	//
	// Each option has a default value. Unless you need to, you don't have to
	// configure any option before calling [StartTaskExecution].
	//
	// You also can override your task options for each task execution. For example,
	// you might want to adjust the LogLevel for an individual execution.
	//
	// [StartTaskExecution]: https://docs.aws.amazon.com/datasync/latest/userguide/API_StartTaskExecution.html
	OverrideOptions *types.Options

	// Specifies the tags that you want to apply to the Amazon Resource Name (ARN)
	// representing the task execution.
	//
	// Tags are key-value pairs that help you manage, filter, and search for your
	// DataSync resources.
	Tags []types.TagListEntry

	// Specifies how you want to configure a task report, which provides detailed
	// information about your DataSync transfer. For more information, see [Monitoring your DataSync transfers with task reports].
	//
	// When using this parameter, your caller identity (the role that you're using
	// DataSync with) must have the iam:PassRole permission. The [AWSDataSyncFullAccess] policy includes this
	// permission.
	//
	// To remove a task report configuration, specify this parameter as empty.
	//
	// [AWSDataSyncFullAccess]: https://docs.aws.amazon.com/datasync/latest/userguide/security-iam-awsmanpol.html#security-iam-awsmanpol-awsdatasyncfullaccess
	// [Monitoring your DataSync transfers with task reports]: https://docs.aws.amazon.com/datasync/latest/userguide/task-reports.html
	TaskReportConfig *types.TaskReportConfig

	noSmithyDocumentSerde
}

// StartTaskExecutionResponse
type StartTaskExecutionOutput struct {

	// The ARN of the running task execution.
	TaskExecutionArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartTaskExecutionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStartTaskExecution{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartTaskExecution{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "StartTaskExecution"); err != nil {
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
	if err = addOpStartTaskExecutionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartTaskExecution(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartTaskExecution(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "StartTaskExecution",
	}
}
