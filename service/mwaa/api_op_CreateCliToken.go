// Code generated by smithy-go-codegen DO NOT EDIT.

package mwaa

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a CLI token for the Airflow CLI. To learn more, see [Creating an Apache Airflow CLI token].
//
// [Creating an Apache Airflow CLI token]: https://docs.aws.amazon.com/mwaa/latest/userguide/call-mwaa-apis-cli.html
func (c *Client) CreateCliToken(ctx context.Context, params *CreateCliTokenInput, optFns ...func(*Options)) (*CreateCliTokenOutput, error) {
	if params == nil {
		params = &CreateCliTokenInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateCliToken", params, optFns, c.addOperationCreateCliTokenMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateCliTokenOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateCliTokenInput struct {

	// The name of the Amazon MWAA environment. For example, MyMWAAEnvironment .
	//
	// This member is required.
	Name *string

	noSmithyDocumentSerde
}

type CreateCliTokenOutput struct {

	// An Airflow CLI login token.
	CliToken *string

	// The Airflow web server hostname for the environment.
	WebServerHostname *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateCliTokenMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateCliToken{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateCliToken{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateCliToken"); err != nil {
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
	if err = addEndpointPrefix_opCreateCliTokenMiddleware(stack); err != nil {
		return err
	}
	if err = addOpCreateCliTokenValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateCliToken(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opCreateCliTokenMiddleware struct {
}

func (*endpointPrefix_opCreateCliTokenMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opCreateCliTokenMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleFinalize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "env." + req.URL.Host

	return next.HandleFinalize(ctx, in)
}
func addEndpointPrefix_opCreateCliTokenMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opCreateCliTokenMiddleware{}, "ResolveEndpointV2", middleware.After)
}

func newServiceMetadataMiddleware_opCreateCliToken(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateCliToken",
	}
}
