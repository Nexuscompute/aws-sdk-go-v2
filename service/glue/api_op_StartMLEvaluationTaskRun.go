// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Starts a task to estimate the quality of the transform.
//
// When you provide label sets as examples of truth, Glue machine learning uses
// some of those examples to learn from them. The rest of the labels are used as a
// test to estimate quality.
//
// Returns a unique identifier for the run. You can call GetMLTaskRun to get more
// information about the stats of the EvaluationTaskRun .
func (c *Client) StartMLEvaluationTaskRun(ctx context.Context, params *StartMLEvaluationTaskRunInput, optFns ...func(*Options)) (*StartMLEvaluationTaskRunOutput, error) {
	if params == nil {
		params = &StartMLEvaluationTaskRunInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartMLEvaluationTaskRun", params, optFns, c.addOperationStartMLEvaluationTaskRunMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartMLEvaluationTaskRunOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartMLEvaluationTaskRunInput struct {

	// The unique identifier of the machine learning transform.
	//
	// This member is required.
	TransformId *string

	noSmithyDocumentSerde
}

type StartMLEvaluationTaskRunOutput struct {

	// The unique identifier associated with this run.
	TaskRunId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartMLEvaluationTaskRunMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStartMLEvaluationTaskRun{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartMLEvaluationTaskRun{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "StartMLEvaluationTaskRun"); err != nil {
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
	if err = addOpStartMLEvaluationTaskRunValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartMLEvaluationTaskRun(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartMLEvaluationTaskRun(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "StartMLEvaluationTaskRun",
	}
}
