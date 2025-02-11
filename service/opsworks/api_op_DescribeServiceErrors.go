// Code generated by smithy-go-codegen DO NOT EDIT.

package opsworks

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/opsworks/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes OpsWorks Stacks service errors.
//
// Required Permissions: To use this action, an IAM user must have a Show, Deploy,
// or Manage permissions level for the stack, or an attached policy that explicitly
// grants permissions. For more information about user permissions, see [Managing User Permissions].
//
// This call accepts only one resource-identifying parameter.
//
// [Managing User Permissions]: https://docs.aws.amazon.com/opsworks/latest/userguide/opsworks-security-users.html
func (c *Client) DescribeServiceErrors(ctx context.Context, params *DescribeServiceErrorsInput, optFns ...func(*Options)) (*DescribeServiceErrorsOutput, error) {
	if params == nil {
		params = &DescribeServiceErrorsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeServiceErrors", params, optFns, c.addOperationDescribeServiceErrorsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeServiceErrorsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeServiceErrorsInput struct {

	// The instance ID. If you use this parameter, DescribeServiceErrors returns
	// descriptions of the errors associated with the specified instance.
	InstanceId *string

	// An array of service error IDs. If you use this parameter, DescribeServiceErrors
	// returns descriptions of the specified errors. Otherwise, it returns a
	// description of every error.
	ServiceErrorIds []string

	// The stack ID. If you use this parameter, DescribeServiceErrors returns
	// descriptions of the errors associated with the specified stack.
	StackId *string

	noSmithyDocumentSerde
}

// Contains the response to a DescribeServiceErrors request.
type DescribeServiceErrorsOutput struct {

	// An array of ServiceError objects that describe the specified service errors.
	ServiceErrors []types.ServiceError

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeServiceErrorsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeServiceErrors{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeServiceErrors{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeServiceErrors"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeServiceErrors(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeServiceErrors(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeServiceErrors",
	}
}
