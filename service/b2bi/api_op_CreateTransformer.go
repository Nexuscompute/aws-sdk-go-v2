// Code generated by smithy-go-codegen DO NOT EDIT.

package b2bi

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/b2bi/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates a transformer. A transformer describes how to process the incoming EDI
// documents and extract the necessary information to the output file.
func (c *Client) CreateTransformer(ctx context.Context, params *CreateTransformerInput, optFns ...func(*Options)) (*CreateTransformerOutput, error) {
	if params == nil {
		params = &CreateTransformerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateTransformer", params, optFns, c.addOperationCreateTransformerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateTransformerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateTransformerInput struct {

	// Specifies the details for the EDI standard that is being used for the
	// transformer. Currently, only X12 is supported. X12 is a set of standards and
	// corresponding messages that define specific business documents.
	//
	// This member is required.
	EdiType types.EdiType

	// Specifies that the currently supported file formats for EDI transformations are
	// JSON and XML .
	//
	// This member is required.
	FileFormat types.FileFormat

	// Specifies the mapping template for the transformer. This template is used to
	// map the parsed EDI file using JSONata or XSLT.
	//
	// This member is required.
	MappingTemplate *string

	// Specifies the name of the transformer, used to identify it.
	//
	// This member is required.
	Name *string

	// Reserved for future use.
	ClientToken *string

	// Specifies a sample EDI document that is used by a transformer as a guide for
	// processing the EDI data.
	SampleDocument *string

	// Specifies the key-value pairs assigned to ARNs that you can use to group and
	// search for resources by type. You can attach this metadata to resources
	// (capabilities, partnerships, and so on) for any purpose.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateTransformerOutput struct {

	// Returns a timestamp for creation date and time of the transformer.
	//
	// This member is required.
	CreatedAt *time.Time

	// Returns the details for the EDI standard that is being used for the
	// transformer. Currently, only X12 is supported. X12 is a set of standards and
	// corresponding messages that define specific business documents.
	//
	// This member is required.
	EdiType types.EdiType

	// Returns that the currently supported file formats for EDI transformations are
	// JSON and XML .
	//
	// This member is required.
	FileFormat types.FileFormat

	// Returns the mapping template for the transformer. This template is used to map
	// the parsed EDI file using JSONata or XSLT.
	//
	// This member is required.
	MappingTemplate *string

	// Returns the name of the transformer, used to identify it.
	//
	// This member is required.
	Name *string

	// Returns the state of the newly created transformer. The transformer can be
	// either active or inactive . For the transformer to be used in a capability, its
	// status must active .
	//
	// This member is required.
	Status types.TransformerStatus

	// Returns an Amazon Resource Name (ARN) for a specific Amazon Web Services
	// resource, such as a capability, partnership, profile, or transformer.
	//
	// This member is required.
	TransformerArn *string

	// Returns the system-assigned unique identifier for the transformer.
	//
	// This member is required.
	TransformerId *string

	// Returns a sample EDI document that is used by a transformer as a guide for
	// processing the EDI data.
	SampleDocument *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateTransformerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpCreateTransformer{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpCreateTransformer{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateTransformer"); err != nil {
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
	if err = addIdempotencyToken_opCreateTransformerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateTransformerValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateTransformer(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateTransformer struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateTransformer) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateTransformer) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateTransformerInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateTransformerInput ")
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
func addIdempotencyToken_opCreateTransformerMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateTransformer{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateTransformer(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateTransformer",
	}
}
