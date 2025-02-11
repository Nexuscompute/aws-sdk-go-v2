// Code generated by smithy-go-codegen DO NOT EDIT.

package backupgateway

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Set the maintenance start time for a gateway.
func (c *Client) PutMaintenanceStartTime(ctx context.Context, params *PutMaintenanceStartTimeInput, optFns ...func(*Options)) (*PutMaintenanceStartTimeOutput, error) {
	if params == nil {
		params = &PutMaintenanceStartTimeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutMaintenanceStartTime", params, optFns, c.addOperationPutMaintenanceStartTimeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutMaintenanceStartTimeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutMaintenanceStartTimeInput struct {

	// The Amazon Resource Name (ARN) for the gateway, used to specify its maintenance
	// start time.
	//
	// This member is required.
	GatewayArn *string

	// The hour of the day to start maintenance on a gateway.
	//
	// This member is required.
	HourOfDay *int32

	// The minute of the hour to start maintenance on a gateway.
	//
	// This member is required.
	MinuteOfHour *int32

	// The day of the month start maintenance on a gateway.
	//
	// Valid values range from Sunday to Saturday .
	DayOfMonth *int32

	// The day of the week to start maintenance on a gateway.
	DayOfWeek *int32

	noSmithyDocumentSerde
}

type PutMaintenanceStartTimeOutput struct {

	// The Amazon Resource Name (ARN) of a gateway for which you set the maintenance
	// start time.
	GatewayArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutMaintenanceStartTimeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpPutMaintenanceStartTime{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpPutMaintenanceStartTime{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PutMaintenanceStartTime"); err != nil {
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
	if err = addOpPutMaintenanceStartTimeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutMaintenanceStartTime(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutMaintenanceStartTime(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PutMaintenanceStartTime",
	}
}
