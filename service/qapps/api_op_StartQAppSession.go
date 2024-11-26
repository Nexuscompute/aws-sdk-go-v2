// Code generated by smithy-go-codegen DO NOT EDIT.

package qapps

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/qapps/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Starts a new session for an Amazon Q App, allowing inputs to be provided and
// the app to be run.
//
// Each Q App session will be condensed into a single conversation in the web
// experience.
func (c *Client) StartQAppSession(ctx context.Context, params *StartQAppSessionInput, optFns ...func(*Options)) (*StartQAppSessionOutput, error) {
	if params == nil {
		params = &StartQAppSessionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartQAppSession", params, optFns, c.addOperationStartQAppSessionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartQAppSessionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartQAppSessionInput struct {

	// The unique identifier of the Q App to start a session for.
	//
	// This member is required.
	AppId *string

	// The version of the Q App to use for the session.
	//
	// This member is required.
	AppVersion *int32

	// The unique identifier of the Amazon Q Business application environment instance.
	//
	// This member is required.
	InstanceId *string

	// Optional initial input values to provide for the Q App session.
	InitialValues []types.CardValue

	// The unique identifier of the a Q App session.
	SessionId *string

	// Optional tags to associate with the new Q App session.
	Tags map[string]string

	noSmithyDocumentSerde
}

type StartQAppSessionOutput struct {

	// The Amazon Resource Name (ARN) of the new Q App session.
	//
	// This member is required.
	SessionArn *string

	// The unique identifier of the new or retrieved Q App session.
	//
	// This member is required.
	SessionId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartQAppSessionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartQAppSession{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartQAppSession{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "StartQAppSession"); err != nil {
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
	if err = addOpStartQAppSessionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartQAppSession(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartQAppSession(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "StartQAppSession",
	}
}
