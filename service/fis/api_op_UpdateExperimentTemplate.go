// Code generated by smithy-go-codegen DO NOT EDIT.

package fis

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/fis/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the specified experiment template.
func (c *Client) UpdateExperimentTemplate(ctx context.Context, params *UpdateExperimentTemplateInput, optFns ...func(*Options)) (*UpdateExperimentTemplateOutput, error) {
	if params == nil {
		params = &UpdateExperimentTemplateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateExperimentTemplate", params, optFns, c.addOperationUpdateExperimentTemplateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateExperimentTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateExperimentTemplateInput struct {

	// The ID of the experiment template.
	//
	// This member is required.
	Id *string

	// The actions for the experiment.
	Actions map[string]types.UpdateExperimentTemplateActionInputItem

	// A description for the template.
	Description *string

	// The experiment options for the experiment template.
	ExperimentOptions *types.UpdateExperimentTemplateExperimentOptionsInput

	// The configuration for experiment logging.
	LogConfiguration *types.UpdateExperimentTemplateLogConfigurationInput

	// The Amazon Resource Name (ARN) of an IAM role that grants the FIS service
	// permission to perform service actions on your behalf.
	RoleArn *string

	// The stop conditions for the experiment.
	StopConditions []types.UpdateExperimentTemplateStopConditionInput

	// The targets for the experiment.
	Targets map[string]types.UpdateExperimentTemplateTargetInput

	noSmithyDocumentSerde
}

type UpdateExperimentTemplateOutput struct {

	// Information about the experiment template.
	ExperimentTemplate *types.ExperimentTemplate

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateExperimentTemplateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateExperimentTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateExperimentTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateExperimentTemplate"); err != nil {
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
	if err = addOpUpdateExperimentTemplateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateExperimentTemplate(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateExperimentTemplate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateExperimentTemplate",
	}
}
