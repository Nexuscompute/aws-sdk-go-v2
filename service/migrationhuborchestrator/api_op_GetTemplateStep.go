// Code generated by smithy-go-codegen DO NOT EDIT.

package migrationhuborchestrator

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/migrationhuborchestrator/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Get a specific step in a template.
func (c *Client) GetTemplateStep(ctx context.Context, params *GetTemplateStepInput, optFns ...func(*Options)) (*GetTemplateStepOutput, error) {
	if params == nil {
		params = &GetTemplateStepInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetTemplateStep", params, optFns, c.addOperationGetTemplateStepMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetTemplateStepOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetTemplateStepInput struct {

	// The ID of the step.
	//
	// This member is required.
	Id *string

	// The ID of the step group.
	//
	// This member is required.
	StepGroupId *string

	// The ID of the template.
	//
	// This member is required.
	TemplateId *string

	noSmithyDocumentSerde
}

type GetTemplateStepOutput struct {

	// The time at which the step was created.
	CreationTime *string

	// The description of the step.
	Description *string

	// The ID of the step.
	Id *string

	// The name of the step.
	Name *string

	// The next step.
	Next []string

	// The outputs of the step.
	Outputs []types.StepOutput

	// The previous step.
	Previous []string

	// The action type of the step. You must run and update the status of a manual
	// step for the workflow to continue after the completion of the step.
	StepActionType types.StepActionType

	// The custom script to run tests on source or target environments.
	StepAutomationConfiguration *types.StepAutomationConfiguration

	// The ID of the step group.
	StepGroupId *string

	// The ID of the template.
	TemplateId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetTemplateStepMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetTemplateStep{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetTemplateStep{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetTemplateStep"); err != nil {
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
	if err = addOpGetTemplateStepValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetTemplateStep(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetTemplateStep(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetTemplateStep",
	}
}
