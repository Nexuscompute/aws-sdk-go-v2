// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the users in the Amazon Cognito user pool.
func (c *Client) ListUsers(ctx context.Context, params *ListUsersInput, optFns ...func(*Options)) (*ListUsersOutput, error) {
	if params == nil {
		params = &ListUsersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListUsers", params, optFns, c.addOperationListUsersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListUsersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the request to list users.
type ListUsersInput struct {

	// The user pool ID for the user pool on which the search should be performed.
	//
	// This member is required.
	UserPoolId *string

	// An array of strings, where each string is the name of a user attribute to be
	// returned for each user in the search results. If the array is null, all
	// attributes are returned.
	AttributesToGet []string

	// A filter string of the form "AttributeName Filter-Type "AttributeValue"".
	// Quotation marks within the filter string must be escaped using the backslash (\)
	// character. For example, " family_name = \"Reddy\"".
	//   - AttributeName: The name of the attribute to search for. You can only search
	//   for one attribute at a time.
	//   - Filter-Type: For an exact match, use =, for example, " given_name =
	//   \"Jon\"". For a prefix ("starts with") match, use ^=, for example, " given_name
	//   ^= \"Jon\"".
	//   - AttributeValue: The attribute value that must be matched for each user.
	// If the filter string is empty, ListUsers returns all users in the user pool.
	// You can only search for the following standard attributes:
	//   - username (case-sensitive)
	//   - email
	//   - phone_number
	//   - name
	//   - given_name
	//   - family_name
	//   - preferred_username
	//   - cognito:user_status (called Status in the Console) (case-insensitive)
	//   - status (called Enabled in the Console) (case-sensitive)
	//   - sub
	// Custom attributes aren't searchable. You can also list users with a client-side
	// filter. The server-side filter matches no more than one attribute. For an
	// advanced search, use a client-side filter with the --query parameter of the
	// list-users action in the CLI. When you use a client-side filter, ListUsers
	// returns a paginated list of zero or more users. You can receive multiple pages
	// in a row with zero results. Repeat the query with each pagination token that is
	// returned until you receive a null pagination token value, and then review the
	// combined result. For more information about server-side and client-side
	// filtering, see FilteringCLI output (https://docs.aws.amazon.com/cli/latest/userguide/cli-usage-filter.html)
	// in the Command Line Interface User Guide (https://docs.aws.amazon.com/cli/latest/userguide/cli-usage-filter.html)
	// . For more information, see Searching for Users Using the ListUsers API (https://docs.aws.amazon.com/cognito/latest/developerguide/how-to-manage-user-accounts.html#cognito-user-pools-searching-for-users-using-listusers-api)
	// and Examples of Using the ListUsers API (https://docs.aws.amazon.com/cognito/latest/developerguide/how-to-manage-user-accounts.html#cognito-user-pools-searching-for-users-listusers-api-examples)
	// in the Amazon Cognito Developer Guide.
	Filter *string

	// Maximum number of users to be returned.
	Limit *int32

	// An identifier that was returned from the previous call to this operation, which
	// can be used to return the next set of items in the list.
	PaginationToken *string

	noSmithyDocumentSerde
}

// The response from the request to list users.
type ListUsersOutput struct {

	// An identifier that was returned from the previous call to this operation, which
	// can be used to return the next set of items in the list.
	PaginationToken *string

	// The users returned in the request to list users.
	Users []types.UserType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListUsersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListUsers{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListUsers{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
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
	if err = addListUsersResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpListUsersValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListUsers(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
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
	if err = addendpointDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

// ListUsersAPIClient is a client that implements the ListUsers operation.
type ListUsersAPIClient interface {
	ListUsers(context.Context, *ListUsersInput, ...func(*Options)) (*ListUsersOutput, error)
}

var _ ListUsersAPIClient = (*Client)(nil)

// ListUsersPaginatorOptions is the paginator options for ListUsers
type ListUsersPaginatorOptions struct {
	// Maximum number of users to be returned.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListUsersPaginator is a paginator for ListUsers
type ListUsersPaginator struct {
	options   ListUsersPaginatorOptions
	client    ListUsersAPIClient
	params    *ListUsersInput
	nextToken *string
	firstPage bool
}

// NewListUsersPaginator returns a new ListUsersPaginator
func NewListUsersPaginator(client ListUsersAPIClient, params *ListUsersInput, optFns ...func(*ListUsersPaginatorOptions)) *ListUsersPaginator {
	if params == nil {
		params = &ListUsersInput{}
	}

	options := ListUsersPaginatorOptions{}
	if params.Limit != nil {
		options.Limit = *params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListUsersPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.PaginationToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListUsersPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListUsers page.
func (p *ListUsersPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListUsersOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.PaginationToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.Limit = limit

	result, err := p.client.ListUsers(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.PaginationToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListUsers(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-idp",
		OperationName: "ListUsers",
	}
}

type opListUsersResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opListUsersResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opListUsersResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if awsmiddleware.GetRequiresLegacyEndpoints(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	if m.EndpointResolver == nil {
		return out, metadata, fmt.Errorf("expected endpoint resolver to not be nil")
	}

	params := EndpointParameters{}

	m.BuiltInResolver.ResolveBuiltIns(&params)

	var resolvedEndpoint smithyendpoints.Endpoint
	resolvedEndpoint, err = m.EndpointResolver.ResolveEndpoint(ctx, params)
	if err != nil {
		return out, metadata, fmt.Errorf("failed to resolve service endpoint, %w", err)
	}

	req.URL = &resolvedEndpoint.URI

	for k := range resolvedEndpoint.Headers {
		req.Header.Set(
			k,
			resolvedEndpoint.Headers.Get(k),
		)
	}

	authSchemes, err := internalauth.GetAuthenticationSchemes(&resolvedEndpoint.Properties)
	if err != nil {
		var nfe *internalauth.NoAuthenticationSchemesFoundError
		if errors.As(err, &nfe) {
			// if no auth scheme is found, default to sigv4
			signingName := "cognito-idp"
			signingRegion := m.BuiltInResolver.(*builtInResolver).Region
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)

		}
		var ue *internalauth.UnSupportedAuthenticationSchemeSpecifiedError
		if errors.As(err, &ue) {
			return out, metadata, fmt.Errorf(
				"This operation requests signer version(s) %v but the client only supports %v",
				ue.UnsupportedSchemes,
				internalauth.SupportedSchemes,
			)
		}
	}

	for _, authScheme := range authSchemes {
		switch authScheme.(type) {
		case *internalauth.AuthenticationSchemeV4:
			v4Scheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4)
			var signingName, signingRegion string
			if v4Scheme.SigningName == nil {
				signingName = "cognito-idp"
			} else {
				signingName = *v4Scheme.SigningName
			}
			if v4Scheme.SigningRegion == nil {
				signingRegion = m.BuiltInResolver.(*builtInResolver).Region
			} else {
				signingRegion = *v4Scheme.SigningRegion
			}
			if v4Scheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4Scheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)
			break
		case *internalauth.AuthenticationSchemeV4A:
			v4aScheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4A)
			if v4aScheme.SigningName == nil {
				v4aScheme.SigningName = aws.String("cognito-idp")
			}
			if v4aScheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4aScheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, *v4aScheme.SigningName)
			ctx = awsmiddleware.SetSigningRegion(ctx, v4aScheme.SigningRegionSet[0])
			break
		case *internalauth.AuthenticationSchemeNone:
			break
		}
	}

	return next.HandleSerialize(ctx, in)
}

func addListUsersResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opListUsersResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
