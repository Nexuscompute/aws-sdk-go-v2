// Code generated by smithy-go-codegen DO NOT EDIT.

package mgn

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/mgn/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Put template post migration custom action.
func (c *Client) PutTemplateAction(ctx context.Context, params *PutTemplateActionInput, optFns ...func(*Options)) (*PutTemplateActionOutput, error) {
	if params == nil {
		params = &PutTemplateActionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutTemplateAction", params, optFns, c.addOperationPutTemplateActionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutTemplateActionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutTemplateActionInput struct {

	// Template post migration custom action ID.
	//
	// This member is required.
	ActionID *string

	// Template post migration custom action name.
	//
	// This member is required.
	ActionName *string

	// Template post migration custom action document identifier.
	//
	// This member is required.
	DocumentIdentifier *string

	// Launch configuration template ID.
	//
	// This member is required.
	LaunchConfigurationTemplateID *string

	// Template post migration custom action order.
	//
	// This member is required.
	Order *int32

	// Template post migration custom action active status.
	Active *bool

	// Template post migration custom action category.
	Category types.ActionCategory

	// Template post migration custom action description.
	Description *string

	// Template post migration custom action document version.
	DocumentVersion *string

	// Template post migration custom action external parameters.
	ExternalParameters map[string]types.SsmExternalParameter

	// Template post migration custom action must succeed for cutover.
	MustSucceedForCutover *bool

	// Operating system eligible for this template post migration custom action.
	OperatingSystem *string

	// Template post migration custom action parameters.
	Parameters map[string][]types.SsmParameterStoreParameter

	// Template post migration custom action timeout in seconds.
	TimeoutSeconds *int32

	noSmithyDocumentSerde
}

type PutTemplateActionOutput struct {

	// Template post migration custom action ID.
	ActionID *string

	// Template post migration custom action name.
	ActionName *string

	// Template post migration custom action active status.
	Active *bool

	// Template post migration custom action category.
	Category types.ActionCategory

	// Template post migration custom action description.
	Description *string

	// Template post migration custom action document identifier.
	DocumentIdentifier *string

	// Template post migration custom action document version.
	DocumentVersion *string

	// Template post migration custom action external parameters.
	ExternalParameters map[string]types.SsmExternalParameter

	// Template post migration custom action must succeed for cutover.
	MustSucceedForCutover *bool

	// Operating system eligible for this template post migration custom action.
	OperatingSystem *string

	// Template post migration custom action order.
	Order *int32

	// Template post migration custom action parameters.
	Parameters map[string][]types.SsmParameterStoreParameter

	// Template post migration custom action timeout in seconds.
	TimeoutSeconds *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutTemplateActionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutTemplateAction{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutTemplateAction{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PutTemplateAction"); err != nil {
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
	if err = addOpPutTemplateActionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutTemplateAction(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutTemplateAction(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PutTemplateAction",
	}
}
