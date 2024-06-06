// Code generated by smithy-go-codegen DO NOT EDIT.

package athena

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/athena/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Requests the cancellation of a calculation. A StopCalculationExecution call on
// a calculation that is already in a terminal state (for example, STOPPED , FAILED
// , or COMPLETED ) succeeds but has no effect.
//
// Cancelling a calculation is done on a best effort basis. If a calculation
// cannot be cancelled, you can be charged for its completion. If you are concerned
// about being charged for a calculation that cannot be cancelled, consider
// terminating the session in which the calculation is running.
func (c *Client) StopCalculationExecution(ctx context.Context, params *StopCalculationExecutionInput, optFns ...func(*Options)) (*StopCalculationExecutionOutput, error) {
	if params == nil {
		params = &StopCalculationExecutionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StopCalculationExecution", params, optFns, c.addOperationStopCalculationExecutionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StopCalculationExecutionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StopCalculationExecutionInput struct {

	// The calculation execution UUID.
	//
	// This member is required.
	CalculationExecutionId *string

	noSmithyDocumentSerde
}

type StopCalculationExecutionOutput struct {

	// CREATING - The calculation is in the process of being created.
	//
	// CREATED - The calculation has been created and is ready to run.
	//
	// QUEUED - The calculation has been queued for processing.
	//
	// RUNNING - The calculation is running.
	//
	// CANCELING - A request to cancel the calculation has been received and the
	// system is working to stop it.
	//
	// CANCELED - The calculation is no longer running as the result of a cancel
	// request.
	//
	// COMPLETED - The calculation has completed without error.
	//
	// FAILED - The calculation failed and is no longer running.
	State types.CalculationExecutionState

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStopCalculationExecutionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStopCalculationExecution{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStopCalculationExecution{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "StopCalculationExecution"); err != nil {
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
	if err = addOpStopCalculationExecutionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStopCalculationExecution(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStopCalculationExecution(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "StopCalculationExecution",
	}
}
