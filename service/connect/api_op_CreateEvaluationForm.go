// Code generated by smithy-go-codegen DO NOT EDIT.

package connect

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/connect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an evaluation form in the specified Amazon Connect instance. The form
// can be used to define questions related to agent performance, and create
// sections to organize such questions. An evaluation form must have a unique title
// within an instance. Question and section identifiers cannot be duplicated within
// the same evaluation form.
func (c *Client) CreateEvaluationForm(ctx context.Context, params *CreateEvaluationFormInput, optFns ...func(*Options)) (*CreateEvaluationFormOutput, error) {
	if params == nil {
		params = &CreateEvaluationFormInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateEvaluationForm", params, optFns, c.addOperationCreateEvaluationFormMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateEvaluationFormOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateEvaluationFormInput struct {

	// The identifier of the Amazon Connect instance. You can find the instance ID (https://docs.aws.amazon.com/connect/latest/adminguide/find-instance-arn.html)
	// in the Amazon Resource Name (ARN) of the instance.
	//
	// This member is required.
	InstanceId *string

	// Items that are part of the evaluation form. The total number of sections and
	// questions must not exceed 100 each. Questions must be contained in a section.
	//
	// This member is required.
	Items []types.EvaluationFormItem

	// A unique title of the evaluation form.
	//
	// This member is required.
	Title *string

	// A unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request. If not provided, the Amazon Web Services SDK populates this
	// field. For more information about idempotency, see Making retries safe with
	// idempotent APIs (https://aws.amazon.com/builders-library/making-retries-safe-with-idempotent-APIs/)
	// .
	ClientToken *string

	// The description of the evaluation form.
	Description *string

	// A scoring strategy of the evaluation form.
	ScoringStrategy *types.EvaluationFormScoringStrategy

	noSmithyDocumentSerde
}

type CreateEvaluationFormOutput struct {

	// The Amazon Resource Name (ARN) for the evaluation form resource.
	//
	// This member is required.
	EvaluationFormArn *string

	// The unique identifier for the evaluation form.
	//
	// This member is required.
	EvaluationFormId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateEvaluationFormMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateEvaluationForm{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateEvaluationForm{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addIdempotencyToken_opCreateEvaluationFormMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateEvaluationFormValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateEvaluationForm(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
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
	return nil
}

type idempotencyToken_initializeOpCreateEvaluationForm struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateEvaluationForm) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateEvaluationForm) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateEvaluationFormInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateEvaluationFormInput ")
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
func addIdempotencyToken_opCreateEvaluationFormMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateEvaluationForm{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateEvaluationForm(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "connect",
		OperationName: "CreateEvaluationForm",
	}
}