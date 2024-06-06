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

// Updates an access entry.
func (c *Client) UpdateAccessEntry(ctx context.Context, params *UpdateAccessEntryInput, optFns ...func(*Options)) (*UpdateAccessEntryOutput, error) {
	if params == nil {
		params = &UpdateAccessEntryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateAccessEntry", params, optFns, c.addOperationUpdateAccessEntryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateAccessEntryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateAccessEntryInput struct {

	// The name of your cluster.
	//
	// This member is required.
	ClusterName *string

	// The ARN of the IAM principal for the AccessEntry .
	//
	// This member is required.
	PrincipalArn *string

	// A unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request.
	ClientRequestToken *string

	// The value for name that you've specified for kind: Group as a subject in a
	// Kubernetes RoleBinding or ClusterRoleBinding object. Amazon EKS doesn't confirm
	// that the value for name exists in any bindings on your cluster. You can specify
	// one or more names.
	//
	// Kubernetes authorizes the principalArn of the access entry to access any
	// cluster objects that you've specified in a Kubernetes Role or ClusterRole
	// object that is also specified in a binding's roleRef . For more information
	// about creating Kubernetes RoleBinding , ClusterRoleBinding , Role , or
	// ClusterRole objects, see [Using RBAC Authorization in the Kubernetes documentation].
	//
	// If you want Amazon EKS to authorize the principalArn (instead of, or in
	// addition to Kubernetes authorizing the principalArn ), you can associate one or
	// more access policies to the access entry using AssociateAccessPolicy . If you
	// associate any access policies, the principalARN has all permissions assigned in
	// the associated access policies and all permissions in any Kubernetes Role or
	// ClusterRole objects that the group names are bound to.
	//
	// [Using RBAC Authorization in the Kubernetes documentation]: https://kubernetes.io/docs/reference/access-authn-authz/rbac/
	KubernetesGroups []string

	// The username to authenticate to Kubernetes with. We recommend not specifying a
	// username and letting Amazon EKS specify it for you. For more information about
	// the value Amazon EKS specifies for you, or constraints before specifying your
	// own username, see [Creating access entries]in the Amazon EKS User Guide.
	//
	// [Creating access entries]: https://docs.aws.amazon.com/eks/latest/userguide/access-entries.html#creating-access-entries
	Username *string

	noSmithyDocumentSerde
}

type UpdateAccessEntryOutput struct {

	// The ARN of the IAM principal for the AccessEntry .
	AccessEntry *types.AccessEntry

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateAccessEntryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateAccessEntry{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateAccessEntry{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateAccessEntry"); err != nil {
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
	if err = addIdempotencyToken_opUpdateAccessEntryMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpUpdateAccessEntryValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateAccessEntry(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpUpdateAccessEntry struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpUpdateAccessEntry) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpUpdateAccessEntry) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*UpdateAccessEntryInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *UpdateAccessEntryInput ")
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
func addIdempotencyToken_opUpdateAccessEntryMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpUpdateAccessEntry{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opUpdateAccessEntry(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateAccessEntry",
	}
}
