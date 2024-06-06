// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpointsmsvoicev2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Permanently delete the protect configuration. The protect configuration must
// have deletion protection disabled and must not be associated as the account
// default protect configuration or associated with a configuration set.
func (c *Client) DeleteProtectConfiguration(ctx context.Context, params *DeleteProtectConfigurationInput, optFns ...func(*Options)) (*DeleteProtectConfigurationOutput, error) {
	if params == nil {
		params = &DeleteProtectConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteProtectConfiguration", params, optFns, c.addOperationDeleteProtectConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteProtectConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteProtectConfigurationInput struct {

	// The unique identifier for the protect configuration.
	//
	// This member is required.
	ProtectConfigurationId *string

	noSmithyDocumentSerde
}

type DeleteProtectConfigurationOutput struct {

	// This is true if the protect configuration is set as your account default
	// protect configuration.
	//
	// This member is required.
	AccountDefault bool

	// The time when the protect configuration was created, in [UNIX epoch time] format.
	//
	// [UNIX epoch time]: https://www.epochconverter.com/
	//
	// This member is required.
	CreatedTimestamp *time.Time

	// The status of deletion protection for the protect configuration. When set to
	// true deletion protection is enabled. By default this is set to false.
	//
	// This member is required.
	DeletionProtectionEnabled bool

	// The Amazon Resource Name (ARN) of the protect configuration.
	//
	// This member is required.
	ProtectConfigurationArn *string

	// The unique identifier for the protect configuration.
	//
	// This member is required.
	ProtectConfigurationId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteProtectConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpDeleteProtectConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpDeleteProtectConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeleteProtectConfiguration"); err != nil {
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
	if err = addOpDeleteProtectConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteProtectConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteProtectConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeleteProtectConfiguration",
	}
}
