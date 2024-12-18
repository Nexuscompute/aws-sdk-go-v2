// Code generated by smithy-go-codegen DO NOT EDIT.

package datasync

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Modifies the following configuration parameters of the Amazon FSx for Lustre
// transfer location that you're using with DataSync.
//
// For more information, see [Configuring DataSync transfers with FSx for Lustre].
//
// [Configuring DataSync transfers with FSx for Lustre]: https://docs.aws.amazon.com/datasync/latest/userguide/create-lustre-location.html
func (c *Client) UpdateLocationFsxLustre(ctx context.Context, params *UpdateLocationFsxLustreInput, optFns ...func(*Options)) (*UpdateLocationFsxLustreOutput, error) {
	if params == nil {
		params = &UpdateLocationFsxLustreInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateLocationFsxLustre", params, optFns, c.addOperationUpdateLocationFsxLustreMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateLocationFsxLustreOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateLocationFsxLustreInput struct {

	// Specifies the Amazon Resource Name (ARN) of the FSx for Lustre transfer
	// location that you're updating.
	//
	// This member is required.
	LocationArn *string

	// Specifies a mount path for your FSx for Lustre file system. The path can
	// include subdirectories.
	//
	// When the location is used as a source, DataSync reads data from the mount path.
	// When the location is used as a destination, DataSync writes data to the mount
	// path. If you don't include this parameter, DataSync uses the file system's root
	// directory ( / ).
	Subdirectory *string

	noSmithyDocumentSerde
}

type UpdateLocationFsxLustreOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateLocationFsxLustreMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateLocationFsxLustre{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateLocationFsxLustre{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateLocationFsxLustre"); err != nil {
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
	if err = addOpUpdateLocationFsxLustreValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateLocationFsxLustre(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateLocationFsxLustre(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateLocationFsxLustre",
	}
}
