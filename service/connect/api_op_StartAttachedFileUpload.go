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

// Provides a pre-signed Amazon S3 URL in response for uploading your content.
//
// You may only use this API to upload attachments to an [Amazon Connect Case] or [Amazon Connect Email].
//
// [Amazon Connect Case]: https://docs.aws.amazon.com/connect/latest/APIReference/API_connect-cases_CreateCase.html
// [Amazon Connect Email]: https://docs.aws.amazon.com/connect/latest/adminguide/setup-email-channel.html
func (c *Client) StartAttachedFileUpload(ctx context.Context, params *StartAttachedFileUploadInput, optFns ...func(*Options)) (*StartAttachedFileUploadOutput, error) {
	if params == nil {
		params = &StartAttachedFileUploadInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartAttachedFileUpload", params, optFns, c.addOperationStartAttachedFileUploadMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartAttachedFileUploadOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartAttachedFileUploadInput struct {

	// The resource to which the attached file is (being) uploaded to. The supported
	// resources are [Cases]and [Email].
	//
	// This value must be a valid ARN.
	//
	// [Email]: https://docs.aws.amazon.com/connect/latest/adminguide/setup-email-channel.html
	// [Cases]: https://docs.aws.amazon.com/connect/latest/adminguide/cases.html
	//
	// This member is required.
	AssociatedResourceArn *string

	// A case-sensitive name of the attached file being uploaded.
	//
	// This member is required.
	FileName *string

	// The size of the attached file in bytes.
	//
	// This member is required.
	FileSizeInBytes *int64

	// The use case for the file.
	//
	// Only ATTACHMENTS are supported.
	//
	// This member is required.
	FileUseCaseType types.FileUseCaseType

	// The unique identifier of the Amazon Connect instance.
	//
	// This member is required.
	InstanceId *string

	// A unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request. If not provided, the Amazon Web Services SDK populates this
	// field. For more information about idempotency, see [Making retries safe with idempotent APIs].
	//
	// [Making retries safe with idempotent APIs]: https://aws.amazon.com/builders-library/making-retries-safe-with-idempotent-APIs/
	ClientToken *string

	// Represents the identity that created the file.
	CreatedBy types.CreatedByInfo

	// The tags used to organize, track, or control access for this resource. For
	// example, { "Tags": {"key1":"value1", "key2":"value2"} } .
	Tags map[string]string

	// Optional override for the expiry of the pre-signed S3 URL in seconds. The
	// default value is 300.
	UrlExpiryInSeconds *int32

	noSmithyDocumentSerde
}

// Response from StartAttachedFileUpload API.
type StartAttachedFileUploadOutput struct {

	// Represents the identity that created the file.
	CreatedBy types.CreatedByInfo

	// The time of Creation of the file resource as an ISO timestamp. It's specified
	// in ISO 8601 format: yyyy-MM-ddThh:mm:ss.SSSZ . For example,
	// 2024-05-03T02:41:28.172Z .
	CreationTime *string

	// The unique identifier of the attached file resource (ARN).
	FileArn *string

	// The unique identifier of the attached file resource.
	FileId *string

	// The current status of the attached file.
	FileStatus types.FileStatusType

	// Information to be used while uploading the attached file.
	UploadUrlMetadata *types.UploadUrlMetadata

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartAttachedFileUploadMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartAttachedFileUpload{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartAttachedFileUpload{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "StartAttachedFileUpload"); err != nil {
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
	if err = addIdempotencyToken_opStartAttachedFileUploadMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpStartAttachedFileUploadValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartAttachedFileUpload(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpStartAttachedFileUpload struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpStartAttachedFileUpload) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpStartAttachedFileUpload) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*StartAttachedFileUploadInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *StartAttachedFileUploadInput ")
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
func addIdempotencyToken_opStartAttachedFileUploadMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpStartAttachedFileUpload{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opStartAttachedFileUpload(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "StartAttachedFileUpload",
	}
}
