// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new virtual MFA device for the Amazon Web Services account. After
// creating the virtual MFA, use EnableMFADeviceto attach the MFA device to an IAM user. For more
// information about creating and working with virtual MFA devices, see [Using a virtual MFA device]in the IAM
// User Guide.
//
// For information about the maximum number of MFA devices you can create, see [IAM and STS quotas] in
// the IAM User Guide.
//
// The seed information contained in the QR code and the Base32 string should be
// treated like any other secret access information. In other words, protect the
// seed information as you would your Amazon Web Services access keys or your
// passwords. After you provision your virtual device, you should ensure that the
// information is destroyed following secure procedures.
//
// [Using a virtual MFA device]: https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_VirtualMFA.html
// [IAM and STS quotas]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_iam-quotas.html
func (c *Client) CreateVirtualMFADevice(ctx context.Context, params *CreateVirtualMFADeviceInput, optFns ...func(*Options)) (*CreateVirtualMFADeviceOutput, error) {
	if params == nil {
		params = &CreateVirtualMFADeviceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateVirtualMFADevice", params, optFns, c.addOperationCreateVirtualMFADeviceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateVirtualMFADeviceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateVirtualMFADeviceInput struct {

	// The name of the virtual MFA device, which must be unique. Use with path to
	// uniquely identify a virtual MFA device.
	//
	// This parameter allows (through its [regex pattern]) a string of characters consisting of upper
	// and lowercase alphanumeric characters with no spaces. You can also include any
	// of the following characters: _+=,.@-
	//
	// [regex pattern]: http://wikipedia.org/wiki/regex
	//
	// This member is required.
	VirtualMFADeviceName *string

	//  The path for the virtual MFA device. For more information about paths, see [IAM identifiers] in
	// the IAM User Guide.
	//
	// This parameter is optional. If it is not included, it defaults to a slash (/).
	//
	// This parameter allows (through its [regex pattern]) a string of characters consisting of
	// either a forward slash (/) by itself or a string that must begin and end with
	// forward slashes. In addition, it can contain any ASCII character from the ! (
	// \u0021 ) through the DEL character ( \u007F ), including most punctuation
	// characters, digits, and upper and lowercased letters.
	//
	// [IAM identifiers]: https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html
	// [regex pattern]: http://wikipedia.org/wiki/regex
	Path *string

	// A list of tags that you want to attach to the new IAM virtual MFA device. Each
	// tag consists of a key name and an associated value. For more information about
	// tagging, see [Tagging IAM resources]in the IAM User Guide.
	//
	// If any one of the tags is invalid or if you exceed the allowed maximum number
	// of tags, then the entire request fails and the resource is not created.
	//
	// [Tagging IAM resources]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_tags.html
	Tags []types.Tag

	noSmithyDocumentSerde
}

// Contains the response to a successful CreateVirtualMFADevice request.
type CreateVirtualMFADeviceOutput struct {

	// A structure containing details about the new virtual MFA device.
	//
	// This member is required.
	VirtualMFADevice *types.VirtualMFADevice

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateVirtualMFADeviceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCreateVirtualMFADevice{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateVirtualMFADevice{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateVirtualMFADevice"); err != nil {
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
	if err = addOpCreateVirtualMFADeviceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateVirtualMFADevice(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateVirtualMFADevice(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateVirtualMFADevice",
	}
}
