// Code generated by smithy-go-codegen DO NOT EDIT.

package connect

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Releases a phone number previously claimed to an Amazon Connect instance or
// traffic distribution group. You can call this API only in the Amazon Web
// Services Region where the number was claimed.
//
// To release phone numbers from a traffic distribution group, use the
// ReleasePhoneNumber API, not the Amazon Connect admin website.
//
// After releasing a phone number, the phone number enters into a cooldown period
// for up to 180 days. It cannot be searched for or claimed again until the period
// has ended. If you accidentally release a phone number, contact Amazon Web
// Services Support.
//
// If you plan to claim and release numbers frequently, contact us for a service
// quota exception. Otherwise, it is possible you will be blocked from claiming and
// releasing any more numbers until up to 180 days past the oldest number released
// has expired.
//
// By default you can claim and release up to 200% of your maximum number of
// active phone numbers. If you claim and release phone numbers using the UI or API
// during a rolling 180 day cycle that exceeds 200% of your phone number service
// level quota, you will be blocked from claiming any more numbers until 180 days
// past the oldest number released has expired.
//
// For example, if you already have 99 claimed numbers and a service level quota
// of 99 phone numbers, and in any 180 day period you release 99, claim 99, and
// then release 99, you will have exceeded the 200% limit. At that point you are
// blocked from claiming any more numbers until you open an Amazon Web Services
// support ticket.
func (c *Client) ReleasePhoneNumber(ctx context.Context, params *ReleasePhoneNumberInput, optFns ...func(*Options)) (*ReleasePhoneNumberOutput, error) {
	if params == nil {
		params = &ReleasePhoneNumberInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ReleasePhoneNumber", params, optFns, c.addOperationReleasePhoneNumberMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ReleasePhoneNumberOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ReleasePhoneNumberInput struct {

	// A unique identifier for the phone number.
	//
	// This member is required.
	PhoneNumberId *string

	// A unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request. If not provided, the Amazon Web Services SDK populates this
	// field. For more information about idempotency, see [Making retries safe with idempotent APIs].
	//
	// [Making retries safe with idempotent APIs]: https://aws.amazon.com/builders-library/making-retries-safe-with-idempotent-APIs/
	ClientToken *string

	noSmithyDocumentSerde
}

type ReleasePhoneNumberOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationReleasePhoneNumberMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpReleasePhoneNumber{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpReleasePhoneNumber{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ReleasePhoneNumber"); err != nil {
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
	if err = addIdempotencyToken_opReleasePhoneNumberMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpReleasePhoneNumberValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opReleasePhoneNumber(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpReleasePhoneNumber struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpReleasePhoneNumber) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpReleasePhoneNumber) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*ReleasePhoneNumberInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *ReleasePhoneNumberInput ")
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
func addIdempotencyToken_opReleasePhoneNumberMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpReleasePhoneNumber{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opReleasePhoneNumber(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ReleasePhoneNumber",
	}
}
