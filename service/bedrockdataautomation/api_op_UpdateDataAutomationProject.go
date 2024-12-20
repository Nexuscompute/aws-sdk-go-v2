// Code generated by smithy-go-codegen DO NOT EDIT.

package bedrockdataautomation

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/bedrockdataautomation/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates an existing Amazon Bedrock Data Automation Project
func (c *Client) UpdateDataAutomationProject(ctx context.Context, params *UpdateDataAutomationProjectInput, optFns ...func(*Options)) (*UpdateDataAutomationProjectOutput, error) {
	if params == nil {
		params = &UpdateDataAutomationProjectInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateDataAutomationProject", params, optFns, c.addOperationUpdateDataAutomationProjectMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateDataAutomationProjectOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Update DataAutomationProject Request
type UpdateDataAutomationProjectInput struct {

	// ARN generated at the server side when a DataAutomationProject is created
	//
	// This member is required.
	ProjectArn *string

	// Standard output configuration
	//
	// This member is required.
	StandardOutputConfiguration *types.StandardOutputConfiguration

	// Custom output configuration
	CustomOutputConfiguration *types.CustomOutputConfiguration

	// Override configuration
	OverrideConfiguration *types.OverrideConfiguration

	// Description of the DataAutomationProject
	ProjectDescription *string

	// Stage of the Project
	ProjectStage types.DataAutomationProjectStage

	noSmithyDocumentSerde
}

// Update DataAutomationProject Response
type UpdateDataAutomationProjectOutput struct {

	// ARN of a DataAutomationProject
	//
	// This member is required.
	ProjectArn *string

	// Stage of the Project
	ProjectStage types.DataAutomationProjectStage

	// Status of Data Automation Project
	Status types.DataAutomationProjectStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateDataAutomationProjectMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateDataAutomationProject{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateDataAutomationProject{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateDataAutomationProject"); err != nil {
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
	if err = addOpUpdateDataAutomationProjectValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateDataAutomationProject(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateDataAutomationProject(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateDataAutomationProject",
	}
}
