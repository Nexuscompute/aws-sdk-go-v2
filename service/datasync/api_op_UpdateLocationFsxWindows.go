// Code generated by smithy-go-codegen DO NOT EDIT.

package datasync

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Modifies the following configuration parameters of the Amazon FSx for Windows
// File Server transfer location that you're using with DataSync.
//
// For more information, see [Configuring DataSync transfers with FSx for Windows File Server].
//
// [Configuring DataSync transfers with FSx for Windows File Server]: https://docs.aws.amazon.com/datasync/latest/userguide/create-fsx-location.html
func (c *Client) UpdateLocationFsxWindows(ctx context.Context, params *UpdateLocationFsxWindowsInput, optFns ...func(*Options)) (*UpdateLocationFsxWindowsOutput, error) {
	if params == nil {
		params = &UpdateLocationFsxWindowsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateLocationFsxWindows", params, optFns, c.addOperationUpdateLocationFsxWindowsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateLocationFsxWindowsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateLocationFsxWindowsInput struct {

	// Specifies the ARN of the FSx for Windows File Server transfer location that
	// you're updating.
	//
	// This member is required.
	LocationArn *string

	// Specifies the name of the Windows domain that your FSx for Windows File Server
	// file system belongs to.
	//
	// If you have multiple Active Directory domains in your environment, configuring
	// this parameter makes sure that DataSync connects to the right file system.
	Domain *string

	// Specifies the password of the user with the permissions to mount and access the
	// files, folders, and file metadata in your FSx for Windows File Server file
	// system.
	Password *string

	// Specifies a mount path for your file system using forward slashes. DataSync
	// uses this subdirectory to read or write data (depending on whether the file
	// system is a source or destination location).
	Subdirectory *string

	// Specifies the user with the permissions to mount and access the files, folders,
	// and file metadata in your FSx for Windows File Server file system.
	//
	// For information about choosing a user with the right level of access for your
	// transfer, see [required permissions]for FSx for Windows File Server locations.
	//
	// [required permissions]: https://docs.aws.amazon.com/datasync/latest/userguide/create-fsx-location.html#create-fsx-windows-location-permissions
	User *string

	noSmithyDocumentSerde
}

type UpdateLocationFsxWindowsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateLocationFsxWindowsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateLocationFsxWindows{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateLocationFsxWindows{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateLocationFsxWindows"); err != nil {
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
	if err = addOpUpdateLocationFsxWindowsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateLocationFsxWindows(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateLocationFsxWindows(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateLocationFsxWindows",
	}
}
