// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Disables the HTTP endpoint for the specified DB cluster. Disabling this
// endpoint disables RDS Data API.
//
// For more information, see [Using RDS Data API] in the Amazon Aurora User Guide.
//
// This operation applies only to Aurora PostgreSQL Serverless v2 and provisioned
// DB clusters. To disable the HTTP endpoint for Aurora Serverless v1 DB clusters,
// use the EnableHttpEndpoint parameter of the ModifyDBCluster operation.
//
// [Using RDS Data API]: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/data-api.html
func (c *Client) DisableHttpEndpoint(ctx context.Context, params *DisableHttpEndpointInput, optFns ...func(*Options)) (*DisableHttpEndpointOutput, error) {
	if params == nil {
		params = &DisableHttpEndpointInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DisableHttpEndpoint", params, optFns, c.addOperationDisableHttpEndpointMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DisableHttpEndpointOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisableHttpEndpointInput struct {

	// The Amazon Resource Name (ARN) of the DB cluster.
	//
	// This member is required.
	ResourceArn *string

	noSmithyDocumentSerde
}

type DisableHttpEndpointOutput struct {

	// Indicates whether the HTTP endpoint is enabled or disabled for the DB cluster.
	HttpEndpointEnabled *bool

	// The ARN of the DB cluster.
	ResourceArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDisableHttpEndpointMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDisableHttpEndpoint{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDisableHttpEndpoint{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DisableHttpEndpoint"); err != nil {
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
	if err = addOpDisableHttpEndpointValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDisableHttpEndpoint(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDisableHttpEndpoint(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DisableHttpEndpoint",
	}
}
