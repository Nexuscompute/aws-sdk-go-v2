// Code generated by smithy-go-codegen DO NOT EDIT.

package ses

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Enables or disables email sending across your entire Amazon SES account in the
// current Amazon Web Services Region. You can use this operation in conjunction
// with Amazon CloudWatch alarms to temporarily pause email sending across your
// Amazon SES account in a given Amazon Web Services Region when reputation metrics
// (such as your bounce or complaint rates) reach certain thresholds.
//
// You can execute this operation no more than once per second.
func (c *Client) UpdateAccountSendingEnabled(ctx context.Context, params *UpdateAccountSendingEnabledInput, optFns ...func(*Options)) (*UpdateAccountSendingEnabledOutput, error) {
	if params == nil {
		params = &UpdateAccountSendingEnabledInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateAccountSendingEnabled", params, optFns, c.addOperationUpdateAccountSendingEnabledMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateAccountSendingEnabledOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a request to enable or disable the email sending capabilities for
// your entire Amazon SES account.
type UpdateAccountSendingEnabledInput struct {

	// Describes whether email sending is enabled or disabled for your Amazon SES
	// account in the current Amazon Web Services Region.
	Enabled bool

	noSmithyDocumentSerde
}

type UpdateAccountSendingEnabledOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateAccountSendingEnabledMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpUpdateAccountSendingEnabled{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpUpdateAccountSendingEnabled{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateAccountSendingEnabled"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateAccountSendingEnabled(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateAccountSendingEnabled(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateAccountSendingEnabled",
	}
}
