// Code generated by smithy-go-codegen DO NOT EDIT.

package healthlake

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/healthlake/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a data store that can ingest and export FHIR formatted data.
func (c *Client) CreateFHIRDatastore(ctx context.Context, params *CreateFHIRDatastoreInput, optFns ...func(*Options)) (*CreateFHIRDatastoreOutput, error) {
	if params == nil {
		params = &CreateFHIRDatastoreInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateFHIRDatastore", params, optFns, c.addOperationCreateFHIRDatastoreMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateFHIRDatastoreOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateFHIRDatastoreInput struct {

	// The FHIR version of the data store. The only supported version is R4.
	//
	// This member is required.
	DatastoreTypeVersion types.FHIRVersion

	// Optional user provided token used for ensuring idempotency.
	ClientToken *string

	// The user generated name for the data store.
	DatastoreName *string

	// The configuration of the identity provider that you want to use for your data
	// store.
	IdentityProviderConfiguration *types.IdentityProviderConfiguration

	// Optional parameter to preload data upon creation of the data store. Currently,
	// the only supported preloaded data is synthetic data generated from Synthea.
	PreloadDataConfig *types.PreloadDataConfig

	//  The server-side encryption key configuration for a customer provided
	// encryption key specified for creating a data store.
	SseConfiguration *types.SseConfiguration

	//  Resource tags that are applied to a data store when it is created.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateFHIRDatastoreOutput struct {

	// The data store ARN is generated during the creation of the data store and can
	// be found in the output from the initial data store creation call.
	//
	// This member is required.
	DatastoreArn *string

	// The AWS endpoint for the created data store.
	//
	// This member is required.
	DatastoreEndpoint *string

	// The AWS-generated data store id. This id is in the output from the initial data
	// store creation call.
	//
	// This member is required.
	DatastoreId *string

	// The status of the FHIR data store.
	//
	// This member is required.
	DatastoreStatus types.DatastoreStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateFHIRDatastoreMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpCreateFHIRDatastore{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpCreateFHIRDatastore{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateFHIRDatastore"); err != nil {
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
	if err = addIdempotencyToken_opCreateFHIRDatastoreMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateFHIRDatastoreValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateFHIRDatastore(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateFHIRDatastore struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateFHIRDatastore) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateFHIRDatastore) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateFHIRDatastoreInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateFHIRDatastoreInput ")
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
func addIdempotencyToken_opCreateFHIRDatastoreMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateFHIRDatastore{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateFHIRDatastore(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateFHIRDatastore",
	}
}
