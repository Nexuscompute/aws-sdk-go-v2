// Code generated by smithy-go-codegen DO NOT EDIT.

package verifiedpermissions

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/verifiedpermissions/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates a reference to an Amazon Cognito user pool as an external identity
// provider (IdP).
//
// After you create an identity source, you can use the identities provided by the
// IdP as proxies for the principal in authorization queries that use the [IsAuthorizedWithToken]
// operation. These identities take the form of tokens that contain claims about
// the user, such as IDs, attributes and group memberships. Amazon Cognito provides
// both identity tokens and access tokens, and Verified Permissions can use either
// or both. Any combination of identity and access tokens results in the same Cedar
// principal. Verified Permissions automatically translates the information about
// the identities into the standard Cedar attributes that can be evaluated by your
// policies. Because the Amazon Cognito identity and access tokens can contain
// different information, the tokens you choose to use determine which principal
// attributes are available to access when evaluating Cedar policies.
//
// If you delete a Amazon Cognito user pool or user, tokens from that deleted pool
// or that deleted user continue to be usable until they expire.
//
// To reference a user from this identity source in your Cedar policies, use the
// following syntax.
//
// IdentityType::"<CognitoUserPoolIdentifier>|<CognitoClientId>
//
// Where IdentityType is the string that you provide to the PrincipalEntityType
// parameter for this operation. The CognitoUserPoolId and CognitoClientId are
// defined by the Amazon Cognito user pool.
//
// Verified Permissions is [eventually consistent] . It can take a few seconds for a new or changed
// element to propagate through the service and be visible in the results of other
// Verified Permissions operations.
//
// [IsAuthorizedWithToken]: https://docs.aws.amazon.com/verifiedpermissions/latest/apireference/API_IsAuthorizedWithToken.html
// [eventually consistent]: https://wikipedia.org/wiki/Eventual_consistency
func (c *Client) CreateIdentitySource(ctx context.Context, params *CreateIdentitySourceInput, optFns ...func(*Options)) (*CreateIdentitySourceOutput, error) {
	if params == nil {
		params = &CreateIdentitySourceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateIdentitySource", params, optFns, c.addOperationCreateIdentitySourceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateIdentitySourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateIdentitySourceInput struct {

	// Specifies the details required to communicate with the identity provider (IdP)
	// associated with this identity source.
	//
	// At this time, the only valid member of this structure is a Amazon Cognito user
	// pool configuration.
	//
	// You must specify a UserPoolArn , and optionally, a ClientId .
	//
	// This member is required.
	Configuration types.Configuration

	// Specifies the ID of the policy store in which you want to store this identity
	// source. Only policies and requests made using this policy store can reference
	// identities from the identity provider configured in the new identity source.
	//
	// This member is required.
	PolicyStoreId *string

	// Specifies a unique, case-sensitive ID that you provide to ensure the
	// idempotency of the request. This lets you safely retry the request without
	// accidentally performing the same operation a second time. Passing the same value
	// to a later call to an operation requires that you also pass the same value for
	// all other parameters. We recommend that you use a [UUID type of value.].
	//
	// If you don't provide this value, then Amazon Web Services generates a random
	// one for you.
	//
	// If you retry the operation with the same ClientToken , but with different
	// parameters, the retry fails with an ConflictException error.
	//
	// Verified Permissions recognizes a ClientToken for eight hours. After eight
	// hours, the next request with the same parameters performs the operation again
	// regardless of the value of ClientToken .
	//
	// [UUID type of value.]: https://wikipedia.org/wiki/Universally_unique_identifier
	ClientToken *string

	// Specifies the namespace and data type of the principals generated for
	// identities authenticated by the new identity source.
	PrincipalEntityType *string

	noSmithyDocumentSerde
}

type CreateIdentitySourceOutput struct {

	// The date and time the identity source was originally created.
	//
	// This member is required.
	CreatedDate *time.Time

	// The unique ID of the new identity source.
	//
	// This member is required.
	IdentitySourceId *string

	// The date and time the identity source was most recently updated.
	//
	// This member is required.
	LastUpdatedDate *time.Time

	// The ID of the policy store that contains the identity source.
	//
	// This member is required.
	PolicyStoreId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateIdentitySourceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpCreateIdentitySource{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpCreateIdentitySource{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateIdentitySource"); err != nil {
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
	if err = addIdempotencyToken_opCreateIdentitySourceMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateIdentitySourceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateIdentitySource(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateIdentitySource struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateIdentitySource) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateIdentitySource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateIdentitySourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateIdentitySourceInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateIdentitySourceMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateIdentitySource{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateIdentitySource(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateIdentitySource",
	}
}
