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

//	Creates a new connector profile associated with your Amazon Web Services
//
// account. There is a soft quota of 100 connector profiles per Amazon Web Services
// account. If you need more connector profiles than this quota allows, you can
// submit a request to the Amazon AppFlow team through the Amazon AppFlow support
// channel. In each connector profile that you create, you can provide the
// credentials and properties for only one connector.
func (c *Client) CreateConnectorProfile(ctx context.Context, params *CreateConnectorProfileInput, optFns ...func(*Options)) (*CreateConnectorProfileOutput, error) {
	if params == nil {
		params = &CreateConnectorProfileInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateConnectorProfile", params, optFns, c.addOperationCreateConnectorProfileMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateConnectorProfileOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateConnectorProfileInput struct {

	//  Indicates the connection mode and specifies whether it is public or private.
	// Private flows use Amazon Web Services PrivateLink to route data over Amazon Web
	// Services infrastructure without exposing it to the public internet.
	//
	// This member is required.
	ConnectionMode types.ConnectionMode

	//  Defines the connector-specific configuration and credentials.
	//
	// This member is required.
	ConnectorProfileConfig *types.ConnectorProfileConfig

	//  The name of the connector profile. The name is unique for each ConnectorProfile
	// in your Amazon Web Services account.
	//
	// This member is required.
	ConnectorProfileName *string

	//  The type of connector, such as Salesforce, Amplitude, and so on.
	//
	// This member is required.
	ConnectorType types.ConnectorType

	// The clientToken parameter is an idempotency token. It ensures that your
	// CreateConnectorProfile request completes only once. You choose the value to
	// pass. For example, if you don't receive a response from your request, you can
	// safely retry the request with the same clientToken parameter value.
	//
	// If you omit a clientToken value, the Amazon Web Services SDK that you are using
	// inserts a value for you. This way, the SDK can safely retry requests multiple
	// times after a network error. You must provide your own value for other use
	// cases.
	//
	// If you specify input parameters that differ from your first request, an error
	// occurs. If you use a different value for clientToken , Amazon AppFlow considers
	// it a new call to CreateConnectorProfile . The token is active for 8 hours.
	ClientToken *string

	// The label of the connector. The label is unique for each ConnectorRegistration
	// in your Amazon Web Services account. Only needed if calling for CUSTOMCONNECTOR
	// connector type/.
	ConnectorLabel *string

	//  The ARN (Amazon Resource Name) of the Key Management Service (KMS) key you
	// provide for encryption. This is required if you do not want to use the Amazon
	// AppFlow-managed KMS key. If you don't provide anything here, Amazon AppFlow uses
	// the Amazon AppFlow-managed KMS key.
	KmsArn *string

	noSmithyDocumentSerde
}

type CreateConnectorProfileOutput struct {

	//  The Amazon Resource Name (ARN) of the connector profile.
	ConnectorProfileArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateConnectorProfileMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateConnectorProfile{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateConnectorProfile{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateConnectorProfile"); err != nil {
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
	if err = addIdempotencyToken_opCreateConnectorProfileMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateConnectorProfileValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateConnectorProfile(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateConnectorProfile struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateConnectorProfile) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateConnectorProfile) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateConnectorProfileInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateConnectorProfileInput ")
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
func addIdempotencyToken_opCreateConnectorProfileMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateConnectorProfile{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateConnectorProfile(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateConnectorProfile",
	}
}
