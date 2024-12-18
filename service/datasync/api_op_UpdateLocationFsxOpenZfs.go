// Code generated by smithy-go-codegen DO NOT EDIT.

package datasync

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/datasync/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Modifies the following configuration parameters of the Amazon FSx for OpenZFS
// transfer location that you're using with DataSync.
//
// For more information, see [Configuring DataSync transfers with FSx for OpenZFS].
//
// Request parameters related to SMB aren't supported with the
// UpdateLocationFsxOpenZfs operation.
//
// [Configuring DataSync transfers with FSx for OpenZFS]: https://docs.aws.amazon.com/datasync/latest/userguide/create-openzfs-location.html
func (c *Client) UpdateLocationFsxOpenZfs(ctx context.Context, params *UpdateLocationFsxOpenZfsInput, optFns ...func(*Options)) (*UpdateLocationFsxOpenZfsOutput, error) {
	if params == nil {
		params = &UpdateLocationFsxOpenZfsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateLocationFsxOpenZfs", params, optFns, c.addOperationUpdateLocationFsxOpenZfsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateLocationFsxOpenZfsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateLocationFsxOpenZfsInput struct {

	// Specifies the Amazon Resource Name (ARN) of the FSx for OpenZFS transfer
	// location that you're updating.
	//
	// This member is required.
	LocationArn *string

	// Specifies the data transfer protocol that DataSync uses to access your Amazon
	// FSx file system.
	Protocol *types.FsxProtocol

	// Specifies a subdirectory in the location's path that must begin with /fsx .
	// DataSync uses this subdirectory to read or write data (depending on whether the
	// file system is a source or destination location).
	Subdirectory *string

	noSmithyDocumentSerde
}

type UpdateLocationFsxOpenZfsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateLocationFsxOpenZfsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateLocationFsxOpenZfs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateLocationFsxOpenZfs{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateLocationFsxOpenZfs"); err != nil {
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
	if err = addOpUpdateLocationFsxOpenZfsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateLocationFsxOpenZfs(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateLocationFsxOpenZfs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateLocationFsxOpenZfs",
	}
}
