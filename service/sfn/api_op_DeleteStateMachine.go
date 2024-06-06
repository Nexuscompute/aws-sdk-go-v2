// Code generated by smithy-go-codegen DO NOT EDIT.

package sfn

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes a state machine. This is an asynchronous operation. It sets the state
// machine's status to DELETING and begins the deletion process. A state machine
// is deleted only when all its executions are completed. On the next state
// transition, the state machine's executions are terminated.
//
// A qualified state machine ARN can either refer to a Distributed Map state
// defined within a state machine, a version ARN, or an alias ARN.
//
// The following are some examples of qualified and unqualified state machine ARNs:
//
//   - The following qualified state machine ARN refers to a Distributed Map state
//     with a label mapStateLabel in a state machine named myStateMachine .
//
// arn:partition:states:region:account-id:stateMachine:myStateMachine/mapStateLabel
//
// If you provide a qualified state machine ARN that refers to a Distributed Map
//
//	state, the request fails with ValidationException .
//
//	- The following unqualified state machine ARN refers to a state machine named
//	myStateMachine .
//
// arn:partition:states:region:account-id:stateMachine:myStateMachine
//
// This API action also deletes all [versions] and [aliases] associated with a state machine.
//
// For EXPRESS state machines, the deletion happens eventually (usually in less
// than a minute). Running executions may emit logs after DeleteStateMachine API
// is called.
//
// [aliases]: https://docs.aws.amazon.com/step-functions/latest/dg/concepts-state-machine-alias.html
// [versions]: https://docs.aws.amazon.com/step-functions/latest/dg/concepts-state-machine-version.html
func (c *Client) DeleteStateMachine(ctx context.Context, params *DeleteStateMachineInput, optFns ...func(*Options)) (*DeleteStateMachineOutput, error) {
	if params == nil {
		params = &DeleteStateMachineInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteStateMachine", params, optFns, c.addOperationDeleteStateMachineMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteStateMachineOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteStateMachineInput struct {

	// The Amazon Resource Name (ARN) of the state machine to delete.
	//
	// This member is required.
	StateMachineArn *string

	noSmithyDocumentSerde
}

type DeleteStateMachineOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteStateMachineMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpDeleteStateMachine{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpDeleteStateMachine{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeleteStateMachine"); err != nil {
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
	if err = addOpDeleteStateMachineValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteStateMachine(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteStateMachine(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeleteStateMachine",
	}
}
