// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the current status of the configuration recorder you specify as well as
// the status of the last recording event for the configuration recorders.
//
// For a detailed status of recording events over time, add your Config events to
// Amazon CloudWatch metrics and use CloudWatch metrics.
//
// If a configuration recorder is not specified, this operation returns the status
// for the customer managed configuration recorder configured for the account, if
// applicable.
//
// When making a request to this operation, you can only specify one configuration
// recorder.
func (c *Client) DescribeConfigurationRecorderStatus(ctx context.Context, params *DescribeConfigurationRecorderStatusInput, optFns ...func(*Options)) (*DescribeConfigurationRecorderStatusOutput, error) {
	if params == nil {
		params = &DescribeConfigurationRecorderStatusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeConfigurationRecorderStatus", params, optFns, c.addOperationDescribeConfigurationRecorderStatusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeConfigurationRecorderStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for the DescribeConfigurationRecorderStatus action.
type DescribeConfigurationRecorderStatusInput struct {

	// The Amazon Resource Name (ARN) of the configuration recorder that you want to
	// specify.
	Arn *string

	// The name of the configuration recorder. If the name is not specified, the
	// opertation returns the status for the customer managed configuration recorder
	// configured for the account, if applicable.
	//
	// When making a request to this operation, you can only specify one configuration
	// recorder.
	ConfigurationRecorderNames []string

	// For service-linked configuration recorders, you can use the service principal
	// of the linked Amazon Web Services service to specify the configuration recorder.
	ServicePrincipal *string

	noSmithyDocumentSerde
}

// The output for the DescribeConfigurationRecorderStatus action, in JSON format.
type DescribeConfigurationRecorderStatusOutput struct {

	// A list that contains status of the specified recorders.
	ConfigurationRecordersStatus []types.ConfigurationRecorderStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeConfigurationRecorderStatusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeConfigurationRecorderStatus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeConfigurationRecorderStatus{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeConfigurationRecorderStatus"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeConfigurationRecorderStatus(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeConfigurationRecorderStatus(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeConfigurationRecorderStatus",
	}
}
