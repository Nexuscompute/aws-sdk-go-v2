// Code generated by smithy-go-codegen DO NOT EDIT.

package servicediscovery

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/servicediscovery/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a private namespace based on DNS, which is visible only inside a
// specified Amazon VPC. The namespace defines your service naming scheme. For
// example, if you name your namespace example.com and name your service backend ,
// the resulting DNS name for the service is backend.example.com . Service
// instances that are registered using a private DNS namespace can be discovered
// using either a DiscoverInstances request or using DNS. For the current quota on
// the number of namespaces that you can create using the same Amazon Web Services
// account, see [Cloud Map quotas]in the Cloud Map Developer Guide.
//
// [Cloud Map quotas]: https://docs.aws.amazon.com/cloud-map/latest/dg/cloud-map-limits.html
func (c *Client) CreatePrivateDnsNamespace(ctx context.Context, params *CreatePrivateDnsNamespaceInput, optFns ...func(*Options)) (*CreatePrivateDnsNamespaceOutput, error) {
	if params == nil {
		params = &CreatePrivateDnsNamespaceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreatePrivateDnsNamespace", params, optFns, c.addOperationCreatePrivateDnsNamespaceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreatePrivateDnsNamespaceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreatePrivateDnsNamespaceInput struct {

	// The name that you want to assign to this namespace. When you create a private
	// DNS namespace, Cloud Map automatically creates an Amazon Route 53 private hosted
	// zone that has the same name as the namespace.
	//
	// This member is required.
	Name *string

	// The ID of the Amazon VPC that you want to associate the namespace with.
	//
	// This member is required.
	Vpc *string

	// A unique string that identifies the request and that allows failed
	// CreatePrivateDnsNamespace requests to be retried without the risk of running the
	// operation twice. CreatorRequestId can be any unique string (for example, a
	// date/timestamp).
	CreatorRequestId *string

	// A description for the namespace.
	Description *string

	// Properties for the private DNS namespace.
	Properties *types.PrivateDnsNamespaceProperties

	// The tags to add to the namespace. Each tag consists of a key and an optional
	// value that you define. Tags keys can be up to 128 characters in length, and tag
	// values can be up to 256 characters in length.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreatePrivateDnsNamespaceOutput struct {

	// A value that you can use to determine whether the request completed
	// successfully. To get the status of the operation, see [GetOperation].
	//
	// [GetOperation]: https://docs.aws.amazon.com/cloud-map/latest/api/API_GetOperation.html
	OperationId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreatePrivateDnsNamespaceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreatePrivateDnsNamespace{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreatePrivateDnsNamespace{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreatePrivateDnsNamespace"); err != nil {
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
	if err = addIdempotencyToken_opCreatePrivateDnsNamespaceMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreatePrivateDnsNamespaceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreatePrivateDnsNamespace(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreatePrivateDnsNamespace struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreatePrivateDnsNamespace) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreatePrivateDnsNamespace) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreatePrivateDnsNamespaceInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreatePrivateDnsNamespaceInput ")
	}

	if input.CreatorRequestId == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.CreatorRequestId = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreatePrivateDnsNamespaceMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreatePrivateDnsNamespace{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreatePrivateDnsNamespace(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreatePrivateDnsNamespace",
	}
}
