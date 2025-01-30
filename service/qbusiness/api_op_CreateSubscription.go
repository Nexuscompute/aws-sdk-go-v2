// Code generated by smithy-go-codegen DO NOT EDIT.

package qbusiness

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/qbusiness/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Subscribes an IAM Identity Center user or a group to a pricing tier for an
// Amazon Q Business application.
//
// Amazon Q Business offers two subscription tiers: Q_LITE and Q_BUSINESS .
// Subscription tier determines feature access for the user. For more information
// on subscriptions and pricing tiers, see [Amazon Q Business pricing].
//
// [Amazon Q Business pricing]: https://aws.amazon.com/q/business/pricing/
func (c *Client) CreateSubscription(ctx context.Context, params *CreateSubscriptionInput, optFns ...func(*Options)) (*CreateSubscriptionOutput, error) {
	if params == nil {
		params = &CreateSubscriptionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateSubscription", params, optFns, c.addOperationCreateSubscriptionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateSubscriptionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateSubscriptionInput struct {

	// The identifier of the Amazon Q Business application the subscription should be
	// added to.
	//
	// This member is required.
	ApplicationId *string

	// The IAM Identity Center UserId or GroupId of a user or group in the IAM
	// Identity Center instance connected to the Amazon Q Business application.
	//
	// This member is required.
	Principal types.SubscriptionPrincipal

	// The type of Amazon Q Business subscription you want to create.
	//
	// This member is required.
	Type types.SubscriptionType

	// A token that you provide to identify the request to create a subscription for
	// your Amazon Q Business application.
	ClientToken *string

	noSmithyDocumentSerde
}

type CreateSubscriptionOutput struct {

	// The type of your current Amazon Q Business subscription.
	CurrentSubscription *types.SubscriptionDetails

	// The type of the Amazon Q Business subscription for the next month.
	NextSubscription *types.SubscriptionDetails

	// The Amazon Resource Name (ARN) of the Amazon Q Business subscription created.
	SubscriptionArn *string

	// The identifier of the Amazon Q Business subscription created.
	SubscriptionId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateSubscriptionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateSubscription{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateSubscription{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateSubscription"); err != nil {
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
	if err = addIdempotencyToken_opCreateSubscriptionMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateSubscriptionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateSubscription(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateSubscription struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateSubscription) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateSubscription) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateSubscriptionInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateSubscriptionInput ")
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
func addIdempotencyToken_opCreateSubscriptionMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateSubscription{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateSubscription(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateSubscription",
	}
}
