// Code generated by smithy-go-codegen DO NOT EDIT.

package eks

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/eks/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Associates an identity provider configuration to a cluster.
//
// If you want to authenticate identities using an identity provider, you can
// create an identity provider configuration and associate it to your cluster.
// After configuring authentication to your cluster you can create Kubernetes Role
// and ClusterRole objects, assign permissions to them, and then bind them to the
// identities using Kubernetes RoleBinding and ClusterRoleBinding objects. For
// more information see [Using RBAC Authorization]in the Kubernetes documentation.
//
// [Using RBAC Authorization]: https://kubernetes.io/docs/reference/access-authn-authz/rbac/
func (c *Client) AssociateIdentityProviderConfig(ctx context.Context, params *AssociateIdentityProviderConfigInput, optFns ...func(*Options)) (*AssociateIdentityProviderConfigOutput, error) {
	if params == nil {
		params = &AssociateIdentityProviderConfigInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AssociateIdentityProviderConfig", params, optFns, c.addOperationAssociateIdentityProviderConfigMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AssociateIdentityProviderConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateIdentityProviderConfigInput struct {

	// The name of your cluster.
	//
	// This member is required.
	ClusterName *string

	// An object representing an OpenID Connect (OIDC) identity provider configuration.
	//
	// This member is required.
	Oidc *types.OidcIdentityProviderConfigRequest

	// A unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request.
	ClientRequestToken *string

	// Metadata that assists with categorization and organization. Each tag consists
	// of a key and an optional value. You define both. Tags don't propagate to any
	// other cluster or Amazon Web Services resources.
	Tags map[string]string

	noSmithyDocumentSerde
}

type AssociateIdentityProviderConfigOutput struct {

	// The tags for the resource.
	Tags map[string]string

	// An object representing an asynchronous update.
	Update *types.Update

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAssociateIdentityProviderConfigMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpAssociateIdentityProviderConfig{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpAssociateIdentityProviderConfig{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "AssociateIdentityProviderConfig"); err != nil {
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
	if err = addIdempotencyToken_opAssociateIdentityProviderConfigMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpAssociateIdentityProviderConfigValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateIdentityProviderConfig(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpAssociateIdentityProviderConfig struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpAssociateIdentityProviderConfig) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpAssociateIdentityProviderConfig) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*AssociateIdentityProviderConfigInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *AssociateIdentityProviderConfigInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opAssociateIdentityProviderConfigMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpAssociateIdentityProviderConfig{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opAssociateIdentityProviderConfig(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "AssociateIdentityProviderConfig",
	}
}
