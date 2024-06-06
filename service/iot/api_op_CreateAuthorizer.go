// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an authorizer.
//
// Requires permission to access the [CreateAuthorizer] action.
//
// [CreateAuthorizer]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func (c *Client) CreateAuthorizer(ctx context.Context, params *CreateAuthorizerInput, optFns ...func(*Options)) (*CreateAuthorizerOutput, error) {
	if params == nil {
		params = &CreateAuthorizerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateAuthorizer", params, optFns, c.addOperationCreateAuthorizerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateAuthorizerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateAuthorizerInput struct {

	// The ARN of the authorizer's Lambda function.
	//
	// This member is required.
	AuthorizerFunctionArn *string

	// The authorizer name.
	//
	// This member is required.
	AuthorizerName *string

	// When true , the result from the authorizer’s Lambda function is cached for
	// clients that use persistent HTTP connections. The results are cached for the
	// time specified by the Lambda function in refreshAfterInSeconds . This value does
	// not affect authorization of clients that use MQTT connections.
	//
	// The default value is false .
	EnableCachingForHttp *bool

	// Specifies whether IoT validates the token signature in an authorization request.
	SigningDisabled *bool

	// The status of the create authorizer request.
	Status types.AuthorizerStatus

	// Metadata which can be used to manage the custom authorizer.
	//
	// For URI Request parameters use format: ...key1=value1&key2=value2...
	//
	// For the CLI command-line parameter use format: &&tags
	// "key1=value1&key2=value2..."
	//
	// For the cli-input-json file use format: "tags": "key1=value1&key2=value2..."
	Tags []types.Tag

	// The name of the token key used to extract the token from the HTTP headers.
	TokenKeyName *string

	// The public keys used to verify the digital signature returned by your custom
	// authentication service.
	TokenSigningPublicKeys map[string]string

	noSmithyDocumentSerde
}

type CreateAuthorizerOutput struct {

	// The authorizer ARN.
	AuthorizerArn *string

	// The authorizer's name.
	AuthorizerName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateAuthorizerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateAuthorizer{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateAuthorizer{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateAuthorizer"); err != nil {
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
	if err = addOpCreateAuthorizerValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateAuthorizer(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateAuthorizer(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateAuthorizer",
	}
}
