// Code generated by smithy-go-codegen DO NOT EDIT.

package s3tables

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the metadata location for a table.
func (c *Client) UpdateTableMetadataLocation(ctx context.Context, params *UpdateTableMetadataLocationInput, optFns ...func(*Options)) (*UpdateTableMetadataLocationOutput, error) {
	if params == nil {
		params = &UpdateTableMetadataLocationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateTableMetadataLocation", params, optFns, c.addOperationUpdateTableMetadataLocationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateTableMetadataLocationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateTableMetadataLocationInput struct {

	// The new metadata location for the table.
	//
	// This member is required.
	MetadataLocation *string

	// The name of the table.
	//
	// This member is required.
	Name *string

	// The namespace of the table.
	//
	// This member is required.
	Namespace *string

	// The Amazon Resource Name (ARN) of the table bucket.
	//
	// This member is required.
	TableBucketARN *string

	// The version token of the table.
	//
	// This member is required.
	VersionToken *string

	noSmithyDocumentSerde
}

type UpdateTableMetadataLocationOutput struct {

	// The metadata location of the table.
	//
	// This member is required.
	MetadataLocation *string

	// The name of the table.
	//
	// This member is required.
	Name *string

	// The namespace the table is associated with.
	//
	// This member is required.
	Namespace []string

	// The Amazon Resource Number (ARN) of the table.
	//
	// This member is required.
	TableARN *string

	// The version token of the table.
	//
	// This member is required.
	VersionToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateTableMetadataLocationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateTableMetadataLocation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateTableMetadataLocation{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateTableMetadataLocation"); err != nil {
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
	if err = addOpUpdateTableMetadataLocationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateTableMetadataLocation(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateTableMetadataLocation(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateTableMetadataLocation",
	}
}
