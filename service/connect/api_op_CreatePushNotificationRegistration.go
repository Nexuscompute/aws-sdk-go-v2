// Code generated by smithy-go-codegen DO NOT EDIT.

package connect

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/connect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates registration for a device token and a chat contact to receive real-time
// push notifications. For more information about push notifications, see [Set up push notifications in Amazon Connect for mobile chat]in the
// Amazon Connect Administrator Guide.
//
// [Set up push notifications in Amazon Connect for mobile chat]: https://docs.aws.amazon.com/connect/latest/adminguide/enable-push-notifications-for-mobile-chat.html
func (c *Client) CreatePushNotificationRegistration(ctx context.Context, params *CreatePushNotificationRegistrationInput, optFns ...func(*Options)) (*CreatePushNotificationRegistrationOutput, error) {
	if params == nil {
		params = &CreatePushNotificationRegistrationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreatePushNotificationRegistration", params, optFns, c.addOperationCreatePushNotificationRegistrationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreatePushNotificationRegistrationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreatePushNotificationRegistrationInput struct {

	// The contact configuration for push notification registration.
	//
	// This member is required.
	ContactConfiguration *types.ContactConfiguration

	// The push notification token issued by the Apple or Google gateways.
	//
	// This member is required.
	DeviceToken *string

	// The device type to use when sending the message.
	//
	// This member is required.
	DeviceType types.DeviceType

	// The identifier of the Amazon Connect instance. You can [find the instance ID] in the Amazon Resource
	// Name (ARN) of the instance.
	//
	// [find the instance ID]: https://docs.aws.amazon.com/connect/latest/adminguide/find-instance-arn.html
	//
	// This member is required.
	InstanceId *string

	// The Amazon Resource Name (ARN) of the Pinpoint application.
	//
	// This member is required.
	PinpointAppArn *string

	// A unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request. If not provided, the Amazon Web Services SDK populates this
	// field. For more information about idempotency, see [Making retries safe with idempotent APIs].
	//
	// [Making retries safe with idempotent APIs]: https://aws.amazon.com/builders-library/making-retries-safe-with-idempotent-APIs/
	ClientToken *string

	noSmithyDocumentSerde
}

type CreatePushNotificationRegistrationOutput struct {

	// The identifier for the registration.
	//
	// This member is required.
	RegistrationId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreatePushNotificationRegistrationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreatePushNotificationRegistration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreatePushNotificationRegistration{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreatePushNotificationRegistration"); err != nil {
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
	if err = addIdempotencyToken_opCreatePushNotificationRegistrationMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreatePushNotificationRegistrationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreatePushNotificationRegistration(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreatePushNotificationRegistration struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreatePushNotificationRegistration) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreatePushNotificationRegistration) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreatePushNotificationRegistrationInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreatePushNotificationRegistrationInput ")
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
func addIdempotencyToken_opCreatePushNotificationRegistrationMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreatePushNotificationRegistration{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreatePushNotificationRegistration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreatePushNotificationRegistration",
	}
}
