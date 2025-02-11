// Code generated by smithy-go-codegen DO NOT EDIT.

package efs

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes a file system, permanently severing access to its contents. Upon
// return, the file system no longer exists and you can't access any contents of
// the deleted file system.
//
// You need to manually delete mount targets attached to a file system before you
// can delete an EFS file system. This step is performed for you when you use the
// Amazon Web Services console to delete a file system.
//
// You cannot delete a file system that is part of an EFS replication
// configuration. You need to delete the replication configuration first.
//
// You can't delete a file system that is in use. That is, if the file system has
// any mount targets, you must first delete them. For more information, see DescribeMountTargetsand DeleteMountTarget.
//
// The DeleteFileSystem call returns while the file system state is still deleting
// . You can check the file system deletion status by calling the DescribeFileSystemsoperation, which
// returns a list of file systems in your account. If you pass file system ID or
// creation token for the deleted file system, the DescribeFileSystemsreturns a 404 FileSystemNotFound
// error.
//
// This operation requires permissions for the elasticfilesystem:DeleteFileSystem
// action.
func (c *Client) DeleteFileSystem(ctx context.Context, params *DeleteFileSystemInput, optFns ...func(*Options)) (*DeleteFileSystemOutput, error) {
	if params == nil {
		params = &DeleteFileSystemInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteFileSystem", params, optFns, c.addOperationDeleteFileSystemMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteFileSystemOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteFileSystemInput struct {

	// The ID of the file system you want to delete.
	//
	// This member is required.
	FileSystemId *string

	noSmithyDocumentSerde
}

type DeleteFileSystemOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteFileSystemMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteFileSystem{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteFileSystem{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeleteFileSystem"); err != nil {
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
	if err = addOpDeleteFileSystemValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteFileSystem(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteFileSystem(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeleteFileSystem",
	}
}
