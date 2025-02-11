// Code generated by smithy-go-codegen DO NOT EDIT.

package opsworks

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Associates one of the stack's registered Elastic IP addresses with a specified
// instance. The address must first be registered with the stack by calling RegisterElasticIp. For
// more information, see [Resource Management].
//
// Required Permissions: To use this action, an IAM user must have a Manage
// permissions level for the stack, or an attached policy that explicitly grants
// permissions. For more information on user permissions, see [Managing User Permissions].
//
// [Resource Management]: https://docs.aws.amazon.com/opsworks/latest/userguide/resources.html
// [Managing User Permissions]: https://docs.aws.amazon.com/opsworks/latest/userguide/opsworks-security-users.html
func (c *Client) AssociateElasticIp(ctx context.Context, params *AssociateElasticIpInput, optFns ...func(*Options)) (*AssociateElasticIpOutput, error) {
	if params == nil {
		params = &AssociateElasticIpInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AssociateElasticIp", params, optFns, c.addOperationAssociateElasticIpMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AssociateElasticIpOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateElasticIpInput struct {

	// The Elastic IP address.
	//
	// This member is required.
	ElasticIp *string

	// The instance ID.
	InstanceId *string

	noSmithyDocumentSerde
}

type AssociateElasticIpOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAssociateElasticIpMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAssociateElasticIp{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAssociateElasticIp{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "AssociateElasticIp"); err != nil {
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
	if err = addOpAssociateElasticIpValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateElasticIp(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAssociateElasticIp(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "AssociateElasticIp",
	}
}
