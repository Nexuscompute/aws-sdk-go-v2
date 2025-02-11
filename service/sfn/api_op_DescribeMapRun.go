// Code generated by smithy-go-codegen DO NOT EDIT.

package sfn

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/sfn/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Provides information about a Map Run's configuration, progress, and results. If
// you've [redriven]a Map Run, this API action also returns information about the redrives
// of that Map Run. For more information, see [Examining Map Run]in the Step Functions Developer
// Guide.
//
// [redriven]: https://docs.aws.amazon.com/step-functions/latest/dg/redrive-map-run.html
// [Examining Map Run]: https://docs.aws.amazon.com/step-functions/latest/dg/concepts-examine-map-run.html
func (c *Client) DescribeMapRun(ctx context.Context, params *DescribeMapRunInput, optFns ...func(*Options)) (*DescribeMapRunOutput, error) {
	if params == nil {
		params = &DescribeMapRunInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeMapRun", params, optFns, c.addOperationDescribeMapRunMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeMapRunOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeMapRunInput struct {

	// The Amazon Resource Name (ARN) that identifies a Map Run.
	//
	// This member is required.
	MapRunArn *string

	noSmithyDocumentSerde
}

type DescribeMapRunOutput struct {

	// The Amazon Resource Name (ARN) that identifies the execution in which the Map
	// Run was started.
	//
	// This member is required.
	ExecutionArn *string

	// A JSON object that contains information about the total number of child
	// workflow executions for the Map Run, and the count of child workflow executions
	// for each status, such as failed and succeeded .
	//
	// This member is required.
	ExecutionCounts *types.MapRunExecutionCounts

	// A JSON object that contains information about the total number of items, and
	// the item count for each processing status, such as pending and failed .
	//
	// This member is required.
	ItemCounts *types.MapRunItemCounts

	// The Amazon Resource Name (ARN) that identifies a Map Run.
	//
	// This member is required.
	MapRunArn *string

	// The maximum number of child workflow executions configured to run in parallel
	// for the Map Run at the same time.
	//
	// This member is required.
	MaxConcurrency int32

	// The date when the Map Run was started.
	//
	// This member is required.
	StartDate *time.Time

	// The current status of the Map Run.
	//
	// This member is required.
	Status types.MapRunStatus

	// The maximum number of failed child workflow executions before the Map Run fails.
	//
	// This member is required.
	ToleratedFailureCount int64

	// The maximum percentage of failed child workflow executions before the Map Run
	// fails.
	//
	// This member is required.
	ToleratedFailurePercentage float32

	// The number of times you've redriven a Map Run. If you have not yet redriven a
	// Map Run, the redriveCount is 0. This count is only updated if you successfully
	// redrive a Map Run.
	RedriveCount *int32

	// The date a Map Run was last redriven. If you have not yet redriven a Map Run,
	// the redriveDate is null.
	RedriveDate *time.Time

	// The date when the Map Run was stopped.
	StopDate *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeMapRunMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpDescribeMapRun{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpDescribeMapRun{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeMapRun"); err != nil {
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
	if err = addOpDescribeMapRunValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeMapRun(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeMapRun(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeMapRun",
	}
}
