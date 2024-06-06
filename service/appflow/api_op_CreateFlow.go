// Code generated by smithy-go-codegen DO NOT EDIT.

package appflow

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/appflow/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

//	Enables your application to create a new flow using Amazon AppFlow. You must
//
// create a connector profile before calling this API. Please note that the Request
// Syntax below shows syntax for multiple destinations, however, you can only
// transfer data to one item in this list at a time. Amazon AppFlow does not
// currently support flows to multiple destinations at once.
func (c *Client) CreateFlow(ctx context.Context, params *CreateFlowInput, optFns ...func(*Options)) (*CreateFlowOutput, error) {
	if params == nil {
		params = &CreateFlowInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateFlow", params, optFns, c.addOperationCreateFlowMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateFlowOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateFlowInput struct {

	//  The configuration that controls how Amazon AppFlow places data in the
	// destination connector.
	//
	// This member is required.
	DestinationFlowConfigList []types.DestinationFlowConfig

	//  The specified name of the flow. Spaces are not allowed. Use underscores (_) or
	// hyphens (-) only.
	//
	// This member is required.
	FlowName *string

	//  The configuration that controls how Amazon AppFlow retrieves data from the
	// source connector.
	//
	// This member is required.
	SourceFlowConfig *types.SourceFlowConfig

	//  A list of tasks that Amazon AppFlow performs while transferring the data in
	// the flow run.
	//
	// This member is required.
	Tasks []types.Task

	//  The trigger settings that determine how and when the flow runs.
	//
	// This member is required.
	TriggerConfig *types.TriggerConfig

	// The clientToken parameter is an idempotency token. It ensures that your
	// CreateFlow request completes only once. You choose the value to pass. For
	// example, if you don't receive a response from your request, you can safely retry
	// the request with the same clientToken parameter value.
	//
	// If you omit a clientToken value, the Amazon Web Services SDK that you are using
	// inserts a value for you. This way, the SDK can safely retry requests multiple
	// times after a network error. You must provide your own value for other use
	// cases.
	//
	// If you specify input parameters that differ from your first request, an error
	// occurs. If you use a different value for clientToken , Amazon AppFlow considers
	// it a new call to CreateFlow . The token is active for 8 hours.
	ClientToken *string

	//  A description of the flow you want to create.
	Description *string

	//  The ARN (Amazon Resource Name) of the Key Management Service (KMS) key you
	// provide for encryption. This is required if you do not want to use the Amazon
	// AppFlow-managed KMS key. If you don't provide anything here, Amazon AppFlow uses
	// the Amazon AppFlow-managed KMS key.
	KmsArn *string

	// Specifies the configuration that Amazon AppFlow uses when it catalogs the data
	// that's transferred by the associated flow. When Amazon AppFlow catalogs the data
	// from a flow, it stores metadata in a data catalog.
	MetadataCatalogConfig *types.MetadataCatalogConfig

	//  The tags used to organize, track, or control access for your flow.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateFlowOutput struct {

	//  The flow's Amazon Resource Name (ARN).
	FlowArn *string

	//  Indicates the current status of the flow.
	FlowStatus types.FlowStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateFlowMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateFlow{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateFlow{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateFlow"); err != nil {
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
	if err = addIdempotencyToken_opCreateFlowMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateFlowValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateFlow(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateFlow struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateFlow) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateFlow) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateFlowInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateFlowInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateFlowMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateFlow{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateFlow(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateFlow",
	}
}
