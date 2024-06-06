// Code generated by smithy-go-codegen DO NOT EDIT.

package clouddirectory

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Publishes a development schema with a major version and a recommended minor
// version.
func (c *Client) PublishSchema(ctx context.Context, params *PublishSchemaInput, optFns ...func(*Options)) (*PublishSchemaOutput, error) {
	if params == nil {
		params = &PublishSchemaInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PublishSchema", params, optFns, c.addOperationPublishSchemaMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PublishSchemaOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PublishSchemaInput struct {

	// The Amazon Resource Name (ARN) that is associated with the development schema.
	// For more information, see arns.
	//
	// This member is required.
	DevelopmentSchemaArn *string

	// The major version under which the schema will be published. Schemas have both a
	// major and minor version associated with them.
	//
	// This member is required.
	Version *string

	// The minor version under which the schema will be published. This parameter is
	// recommended. Schemas have both a major and minor version associated with them.
	MinorVersion *string

	// The new name under which the schema will be published. If this is not provided,
	// the development schema is considered.
	Name *string

	noSmithyDocumentSerde
}

type PublishSchemaOutput struct {

	// The ARN that is associated with the published schema. For more information, see arns
	// .
	PublishedSchemaArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPublishSchemaMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPublishSchema{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPublishSchema{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PublishSchema"); err != nil {
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
	if err = addOpPublishSchemaValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPublishSchema(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPublishSchema(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PublishSchema",
	}
}
