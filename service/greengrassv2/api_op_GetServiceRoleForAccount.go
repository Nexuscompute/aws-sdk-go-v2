// Code generated by smithy-go-codegen DO NOT EDIT.

package greengrassv2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets the service role associated with IoT Greengrass for your Amazon Web
// Services account in this Amazon Web Services Region. IoT Greengrass uses this
// role to verify the identity of client devices and manage core device
// connectivity information. For more information, see [Greengrass service role]in the IoT Greengrass
// Version 2 Developer Guide.
//
// [Greengrass service role]: https://docs.aws.amazon.com/greengrass/v2/developerguide/greengrass-service-role.html
func (c *Client) GetServiceRoleForAccount(ctx context.Context, params *GetServiceRoleForAccountInput, optFns ...func(*Options)) (*GetServiceRoleForAccountOutput, error) {
	if params == nil {
		params = &GetServiceRoleForAccountInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetServiceRoleForAccount", params, optFns, c.addOperationGetServiceRoleForAccountMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetServiceRoleForAccountOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetServiceRoleForAccountInput struct {
	noSmithyDocumentSerde
}

type GetServiceRoleForAccountOutput struct {

	// The time when the service role was associated with IoT Greengrass for your
	// Amazon Web Services account in this Amazon Web Services Region.
	AssociatedAt *string

	// The ARN of the service role that is associated with IoT Greengrass for your
	// Amazon Web Services account in this Amazon Web Services Region.
	RoleArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetServiceRoleForAccountMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetServiceRoleForAccount{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetServiceRoleForAccount{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetServiceRoleForAccount"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetServiceRoleForAccount(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetServiceRoleForAccount(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetServiceRoleForAccount",
	}
}
