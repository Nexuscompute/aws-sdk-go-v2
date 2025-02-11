// Code generated by smithy-go-codegen DO NOT EDIT.

package apigateway

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Generates a ClientCertificate resource.
func (c *Client) GenerateClientCertificate(ctx context.Context, params *GenerateClientCertificateInput, optFns ...func(*Options)) (*GenerateClientCertificateOutput, error) {
	if params == nil {
		params = &GenerateClientCertificateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GenerateClientCertificate", params, optFns, c.addOperationGenerateClientCertificateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GenerateClientCertificateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to generate a ClientCertificate resource.
type GenerateClientCertificateInput struct {

	// The description of the ClientCertificate.
	Description *string

	// The key-value map of strings. The valid character set is [a-zA-Z+-=._:/]. The
	// tag key can be up to 128 characters and must not start with aws: . The tag value
	// can be up to 256 characters.
	Tags map[string]string

	noSmithyDocumentSerde
}

// Represents a client certificate used to configure client-side SSL
// authentication while sending requests to the integration endpoint.
type GenerateClientCertificateOutput struct {

	// The identifier of the client certificate.
	ClientCertificateId *string

	// The timestamp when the client certificate was created.
	CreatedDate *time.Time

	// The description of the client certificate.
	Description *string

	// The timestamp when the client certificate will expire.
	ExpirationDate *time.Time

	// The PEM-encoded public key of the client certificate, which can be used to
	// configure certificate authentication in the integration endpoint .
	PemEncodedCertificate *string

	// The collection of tags. Each tag element is associated with a given resource.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGenerateClientCertificateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGenerateClientCertificate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGenerateClientCertificate{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GenerateClientCertificate"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGenerateClientCertificate(options.Region), middleware.Before); err != nil {
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
	if err = addAcceptHeader(stack); err != nil {
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

func newServiceMetadataMiddleware_opGenerateClientCertificate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GenerateClientCertificate",
	}
}
