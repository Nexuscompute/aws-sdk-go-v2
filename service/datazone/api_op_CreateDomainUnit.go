// Code generated by smithy-go-codegen DO NOT EDIT.

package datazone

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/datazone/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates a domain unit in Amazon DataZone.
func (c *Client) CreateDomainUnit(ctx context.Context, params *CreateDomainUnitInput, optFns ...func(*Options)) (*CreateDomainUnitOutput, error) {
	if params == nil {
		params = &CreateDomainUnitInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDomainUnit", params, optFns, c.addOperationCreateDomainUnitMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDomainUnitOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDomainUnitInput struct {

	// The ID of the domain where you want to crate a domain unit.
	//
	// This member is required.
	DomainIdentifier *string

	// The name of the domain unit.
	//
	// This member is required.
	Name *string

	// The ID of the parent domain unit.
	//
	// This member is required.
	ParentDomainUnitIdentifier *string

	// A unique, case-sensitive identifier that is provided to ensure the idempotency
	// of the request.
	ClientToken *string

	// The description of the domain unit.
	Description *string

	noSmithyDocumentSerde
}

type CreateDomainUnitOutput struct {

	// The IDs of the ancestor domain units.
	//
	// This member is required.
	AncestorDomainUnitIds []string

	// The ID of the domain where the domain unit was created.
	//
	// This member is required.
	DomainId *string

	// The ID of the domain unit.
	//
	// This member is required.
	Id *string

	// The name of the domain unit.
	//
	// This member is required.
	Name *string

	// The owners of the domain unit.
	//
	// This member is required.
	Owners []types.DomainUnitOwnerProperties

	// The timestamp at which the domain unit was created.
	CreatedAt *time.Time

	// The user who created the domain unit.
	CreatedBy *string

	// The description of the domain unit.
	Description *string

	// The ID of the parent domain unit.
	ParentDomainUnitId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateDomainUnitMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateDomainUnit{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateDomainUnit{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateDomainUnit"); err != nil {
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
	if err = addIdempotencyToken_opCreateDomainUnitMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateDomainUnitValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDomainUnit(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateDomainUnit struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateDomainUnit) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateDomainUnit) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateDomainUnitInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateDomainUnitInput ")
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
func addIdempotencyToken_opCreateDomainUnitMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateDomainUnit{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateDomainUnit(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateDomainUnit",
	}
}
