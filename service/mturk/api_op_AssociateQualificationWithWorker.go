// Code generated by smithy-go-codegen DO NOT EDIT.

package mturk

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

//	The AssociateQualificationWithWorker operation gives a Worker a Qualification.
//
// AssociateQualificationWithWorker does not require that the Worker submit a
// Qualification request. It gives the Qualification directly to the Worker.
//
// You can only assign a Qualification of a Qualification type that you created
// (using the CreateQualificationType operation).
//
// Note: AssociateQualificationWithWorker does not affect any pending
// Qualification requests for the Qualification by the Worker. If you assign a
// Qualification to a Worker, then later grant a Qualification request made by the
// Worker, the granting of the request may modify the Qualification score. To
// resolve a pending Qualification request without affecting the Qualification the
// Worker already has, reject the request with the RejectQualificationRequest
// operation.
func (c *Client) AssociateQualificationWithWorker(ctx context.Context, params *AssociateQualificationWithWorkerInput, optFns ...func(*Options)) (*AssociateQualificationWithWorkerOutput, error) {
	if params == nil {
		params = &AssociateQualificationWithWorkerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AssociateQualificationWithWorker", params, optFns, c.addOperationAssociateQualificationWithWorkerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AssociateQualificationWithWorkerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateQualificationWithWorkerInput struct {

	// The ID of the Qualification type to use for the assigned Qualification.
	//
	// This member is required.
	QualificationTypeId *string

	//  The ID of the Worker to whom the Qualification is being assigned. Worker IDs
	// are included with submitted HIT assignments and Qualification requests.
	//
	// This member is required.
	WorkerId *string

	// The value of the Qualification to assign.
	IntegerValue *int32

	//  Specifies whether to send a notification email message to the Worker saying
	// that the qualification was assigned to the Worker. Note: this is true by
	// default.
	SendNotification *bool

	noSmithyDocumentSerde
}

type AssociateQualificationWithWorkerOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAssociateQualificationWithWorkerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAssociateQualificationWithWorker{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAssociateQualificationWithWorker{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "AssociateQualificationWithWorker"); err != nil {
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
	if err = addOpAssociateQualificationWithWorkerValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateQualificationWithWorker(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAssociateQualificationWithWorker(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "AssociateQualificationWithWorker",
	}
}
