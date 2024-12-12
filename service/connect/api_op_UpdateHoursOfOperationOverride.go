// Code generated by smithy-go-codegen DO NOT EDIT.

package connect

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/connect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Update the hours of operation override.
func (c *Client) UpdateHoursOfOperationOverride(ctx context.Context, params *UpdateHoursOfOperationOverrideInput, optFns ...func(*Options)) (*UpdateHoursOfOperationOverrideOutput, error) {
	if params == nil {
		params = &UpdateHoursOfOperationOverrideInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateHoursOfOperationOverride", params, optFns, c.addOperationUpdateHoursOfOperationOverrideMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateHoursOfOperationOverrideOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateHoursOfOperationOverrideInput struct {

	// The identifier for the hours of operation.
	//
	// This member is required.
	HoursOfOperationId *string

	// The identifier for the hours of operation override.
	//
	// This member is required.
	HoursOfOperationOverrideId *string

	// The identifier of the Amazon Connect instance.
	//
	// This member is required.
	InstanceId *string

	// Configuration information for the hours of operation override: day, start time,
	// and end time.
	Config []types.HoursOfOperationOverrideConfig

	// The description of the hours of operation override.
	Description *string

	// The date from when the hours of operation override would be effective.
	EffectiveFrom *string

	// The date till when the hours of operation override would be effective.
	EffectiveTill *string

	// The name of the hours of operation override.
	Name *string

	noSmithyDocumentSerde
}

type UpdateHoursOfOperationOverrideOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateHoursOfOperationOverrideMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateHoursOfOperationOverride{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateHoursOfOperationOverride{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateHoursOfOperationOverride"); err != nil {
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
	if err = addOpUpdateHoursOfOperationOverrideValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateHoursOfOperationOverride(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateHoursOfOperationOverride(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateHoursOfOperationOverride",
	}
}
