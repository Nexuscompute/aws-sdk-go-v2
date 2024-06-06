// Code generated by smithy-go-codegen DO NOT EDIT.

package sns

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Allows a topic owner to set an attribute of the topic to a new value.
//
// To remove the ability to change topic permissions, you must deny permissions to
// the AddPermission , RemovePermission , and SetTopicAttributes actions in your
// IAM policy.
func (c *Client) SetTopicAttributes(ctx context.Context, params *SetTopicAttributesInput, optFns ...func(*Options)) (*SetTopicAttributesOutput, error) {
	if params == nil {
		params = &SetTopicAttributesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SetTopicAttributes", params, optFns, c.addOperationSetTopicAttributesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SetTopicAttributesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Input for SetTopicAttributes action.
type SetTopicAttributesInput struct {

	// A map of attributes with their corresponding values.
	//
	// The following lists the names, descriptions, and values of the special request
	// parameters that the SetTopicAttributes action uses:
	//
	//   - ApplicationSuccessFeedbackRoleArn – Indicates failed message delivery status
	//   for an Amazon SNS topic that is subscribed to a platform application endpoint.
	//
	//   - DeliveryPolicy – The policy that defines how Amazon SNS retries failed
	//   deliveries to HTTP/S endpoints.
	//
	//   - DisplayName – The display name to use for a topic with SMS subscriptions.
	//
	//   - Policy – The policy that defines who can access your topic. By default, only
	//   the topic owner can publish or subscribe to the topic.
	//
	//   - TracingConfig – Tracing mode of an Amazon SNS topic. By default
	//   TracingConfig is set to PassThrough , and the topic passes through the tracing
	//   header it receives from an Amazon SNS publisher to its subscriptions. If set to
	//   Active , Amazon SNS will vend X-Ray segment data to topic owner account if the
	//   sampled flag in the tracing header is true. This is only supported on standard
	//   topics.
	//
	//   - HTTP
	//
	//   - HTTPSuccessFeedbackRoleArn – Indicates successful message delivery status
	//   for an Amazon SNS topic that is subscribed to an HTTP endpoint.
	//
	//   - HTTPSuccessFeedbackSampleRate – Indicates percentage of successful messages
	//   to sample for an Amazon SNS topic that is subscribed to an HTTP endpoint.
	//
	//   - HTTPFailureFeedbackRoleArn – Indicates failed message delivery status for an
	//   Amazon SNS topic that is subscribed to an HTTP endpoint.
	//
	//   - Amazon Kinesis Data Firehose
	//
	//   - FirehoseSuccessFeedbackRoleArn – Indicates successful message delivery
	//   status for an Amazon SNS topic that is subscribed to an Amazon Kinesis Data
	//   Firehose endpoint.
	//
	//   - FirehoseSuccessFeedbackSampleRate – Indicates percentage of successful
	//   messages to sample for an Amazon SNS topic that is subscribed to an Amazon
	//   Kinesis Data Firehose endpoint.
	//
	//   - FirehoseFailureFeedbackRoleArn – Indicates failed message delivery status
	//   for an Amazon SNS topic that is subscribed to an Amazon Kinesis Data Firehose
	//   endpoint.
	//
	//   - Lambda
	//
	//   - LambdaSuccessFeedbackRoleArn – Indicates successful message delivery status
	//   for an Amazon SNS topic that is subscribed to an Lambda endpoint.
	//
	//   - LambdaSuccessFeedbackSampleRate – Indicates percentage of successful
	//   messages to sample for an Amazon SNS topic that is subscribed to an Lambda
	//   endpoint.
	//
	//   - LambdaFailureFeedbackRoleArn – Indicates failed message delivery status for
	//   an Amazon SNS topic that is subscribed to an Lambda endpoint.
	//
	//   - Platform application endpoint
	//
	//   - ApplicationSuccessFeedbackRoleArn – Indicates successful message delivery
	//   status for an Amazon SNS topic that is subscribed to an Amazon Web Services
	//   application endpoint.
	//
	//   - ApplicationSuccessFeedbackSampleRate – Indicates percentage of successful
	//   messages to sample for an Amazon SNS topic that is subscribed to an Amazon Web
	//   Services application endpoint.
	//
	//   - ApplicationFailureFeedbackRoleArn – Indicates failed message delivery status
	//   for an Amazon SNS topic that is subscribed to an Amazon Web Services application
	//   endpoint.
	//
	// In addition to being able to configure topic attributes for message delivery
	//   status of notification messages sent to Amazon SNS application endpoints, you
	//   can also configure application attributes for the delivery status of push
	//   notification messages sent to push notification services.
	//
	// For example, For more information, see [Using Amazon SNS Application Attributes for Message Delivery Status].
	//
	//   - Amazon SQS
	//
	//   - SQSSuccessFeedbackRoleArn – Indicates successful message delivery status for
	//   an Amazon SNS topic that is subscribed to an Amazon SQS endpoint.
	//
	//   - SQSSuccessFeedbackSampleRate – Indicates percentage of successful messages
	//   to sample for an Amazon SNS topic that is subscribed to an Amazon SQS endpoint.
	//
	//   - SQSFailureFeedbackRoleArn – Indicates failed message delivery status for an
	//   Amazon SNS topic that is subscribed to an Amazon SQS endpoint.
	//
	// The SuccessFeedbackRoleArn and FailureFeedbackRoleArn attributes are used to
	// give Amazon SNS write access to use CloudWatch Logs on your behalf. The
	// SuccessFeedbackSampleRate attribute is for specifying the sample rate percentage
	// (0-100) of successfully delivered messages. After you configure the
	// FailureFeedbackRoleArn attribute, then all failed message deliveries generate
	// CloudWatch Logs.
	//
	// The following attribute applies only to [server-side-encryption]:
	//
	//   - KmsMasterKeyId – The ID of an Amazon Web Services managed customer master
	//   key (CMK) for Amazon SNS or a custom CMK. For more information, see [Key Terms]. For
	//   more examples, see [KeyId]in the Key Management Service API Reference.
	//
	//   - SignatureVersion – The signature version corresponds to the hashing
	//   algorithm used while creating the signature of the notifications, subscription
	//   confirmations, or unsubscribe confirmation messages sent by Amazon SNS. By
	//   default, SignatureVersion is set to 1 .
	//
	// The following attribute applies only to [FIFO topics]:
	//
	//   - ContentBasedDeduplication – Enables content-based deduplication for FIFO
	//   topics.
	//
	//   - By default, ContentBasedDeduplication is set to false . If you create a FIFO
	//   topic and this attribute is false , you must specify a value for the
	//   MessageDeduplicationId parameter for the [Publish]action.
	//
	//   - When you set ContentBasedDeduplication to true , Amazon SNS uses a SHA-256
	//   hash to generate the MessageDeduplicationId using the body of the message (but
	//   not the attributes of the message).
	//
	// (Optional) To override the generated value, you can specify a value for the
	//   MessageDeduplicationId parameter for the Publish action.
	//
	// [Using Amazon SNS Application Attributes for Message Delivery Status]: https://docs.aws.amazon.com/sns/latest/dg/sns-msg-status.html
	// [Key Terms]: https://docs.aws.amazon.com/sns/latest/dg/sns-server-side-encryption.html#sse-key-terms
	// [KeyId]: https://docs.aws.amazon.com/kms/latest/APIReference/API_DescribeKey.html#API_DescribeKey_RequestParameters
	// [server-side-encryption]: https://docs.aws.amazon.com/sns/latest/dg/sns-server-side-encryption.html
	// [Publish]: https://docs.aws.amazon.com/sns/latest/api/API_Publish.html
	// [FIFO topics]: https://docs.aws.amazon.com/sns/latest/dg/sns-fifo-topics.html
	//
	// This member is required.
	AttributeName *string

	// The ARN of the topic to modify.
	//
	// This member is required.
	TopicArn *string

	// The new value for the attribute.
	AttributeValue *string

	noSmithyDocumentSerde
}

type SetTopicAttributesOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSetTopicAttributesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpSetTopicAttributes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpSetTopicAttributes{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "SetTopicAttributes"); err != nil {
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
	if err = addOpSetTopicAttributesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSetTopicAttributes(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opSetTopicAttributes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "SetTopicAttributes",
	}
}
