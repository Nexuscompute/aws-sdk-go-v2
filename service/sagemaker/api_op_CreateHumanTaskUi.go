// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Defines the settings you will use for the human review workflow user interface.
// Reviewers will see a three-panel interface with an instruction area, the item to
// review, and an input area.
func (c *Client) CreateHumanTaskUi(ctx context.Context, params *CreateHumanTaskUiInput, optFns ...func(*Options)) (*CreateHumanTaskUiOutput, error) {
	if params == nil {
		params = &CreateHumanTaskUiInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateHumanTaskUi", params, optFns, c.addOperationCreateHumanTaskUiMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateHumanTaskUiOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateHumanTaskUiInput struct {

	// The name of the user interface you are creating.
	//
	// This member is required.
	HumanTaskUiName *string

	// The Liquid template for the worker user interface.
	//
	// This member is required.
	UiTemplate *types.UiTemplate

	// An array of key-value pairs that contain metadata to help you categorize and
	// organize a human review workflow user interface. Each tag consists of a key and
	// a value, both of which you define.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateHumanTaskUiOutput struct {

	// The Amazon Resource Name (ARN) of the human review workflow user interface you
	// create.
	//
	// This member is required.
	HumanTaskUiArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateHumanTaskUiMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateHumanTaskUi{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateHumanTaskUi{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateHumanTaskUi"); err != nil {
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
	if err = addOpCreateHumanTaskUiValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateHumanTaskUi(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateHumanTaskUi(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateHumanTaskUi",
	}
}
