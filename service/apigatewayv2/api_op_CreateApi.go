// Code generated by smithy-go-codegen DO NOT EDIT.

package apigatewayv2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates an Api resource.
func (c *Client) CreateApi(ctx context.Context, params *CreateApiInput, optFns ...func(*Options)) (*CreateApiOutput, error) {
	if params == nil {
		params = &CreateApiInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateApi", params, optFns, c.addOperationCreateApiMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateApiOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Creates a new Api resource to represent an API.
type CreateApiInput struct {

	// The name of the API.
	//
	// This member is required.
	Name *string

	// The API protocol.
	//
	// This member is required.
	ProtocolType types.ProtocolType

	// An API key selection expression. Supported only for WebSocket APIs. See [API Key Selection Expressions].
	//
	// [API Key Selection Expressions]: https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-apikey-selection-expressions
	ApiKeySelectionExpression *string

	// A CORS configuration. Supported only for HTTP APIs. See [Configuring CORS] for more information.
	//
	// [Configuring CORS]: https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-cors.html
	CorsConfiguration *types.Cors

	// This property is part of quick create. It specifies the credentials required
	// for the integration, if any. For a Lambda integration, three options are
	// available. To specify an IAM Role for API Gateway to assume, use the role's
	// Amazon Resource Name (ARN). To require that the caller's identity be passed
	// through from the request, specify arn:aws:iam::*:user/*. To use resource-based
	// permissions on supported AWS services, specify null. Currently, this property is
	// not used for HTTP integrations. Supported only for HTTP APIs.
	CredentialsArn *string

	// The description of the API.
	Description *string

	// Specifies whether clients can invoke your API by using the default execute-api
	// endpoint. By default, clients can invoke your API with the default
	// https://{api_id}.execute-api.{region}.amazonaws.com endpoint. To require that
	// clients use a custom domain name to invoke your API, disable the default
	// endpoint.
	DisableExecuteApiEndpoint *bool

	// Avoid validating models when creating a deployment. Supported only for
	// WebSocket APIs.
	DisableSchemaValidation *bool

	// This property is part of quick create. If you don't specify a routeKey, a
	// default route of $default is created. The $default route acts as a catch-all
	// for any request made to your API, for a particular stage. The $default route
	// key can't be modified. You can add routes after creating the API, and you can
	// update the route keys of additional routes. Supported only for HTTP APIs.
	RouteKey *string

	// The route selection expression for the API. For HTTP APIs, the
	// routeSelectionExpression must be ${request.method} ${request.path}. If not
	// provided, this will be the default for HTTP APIs. This property is required for
	// WebSocket APIs.
	RouteSelectionExpression *string

	// The collection of tags. Each tag element is associated with a given resource.
	Tags map[string]string

	// This property is part of quick create. Quick create produces an API with an
	// integration, a default catch-all route, and a default stage which is configured
	// to automatically deploy changes. For HTTP integrations, specify a fully
	// qualified URL. For Lambda integrations, specify a function ARN. The type of the
	// integration will be HTTP_PROXY or AWS_PROXY, respectively. Supported only for
	// HTTP APIs.
	Target *string

	// A version identifier for the API.
	Version *string

	noSmithyDocumentSerde
}

type CreateApiOutput struct {

	// The URI of the API, of the form {api-id}.execute-api.{region}.amazonaws.com.
	// The stage name is typically appended to this URI to form a complete path to a
	// deployed API stage.
	ApiEndpoint *string

	// Specifies whether an API is managed by API Gateway. You can't update or delete
	// a managed API by using API Gateway. A managed API can be deleted only through
	// the tooling or service that created it.
	ApiGatewayManaged *bool

	// The API ID.
	ApiId *string

	// An API key selection expression. Supported only for WebSocket APIs. See [API Key Selection Expressions].
	//
	// [API Key Selection Expressions]: https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-apikey-selection-expressions
	ApiKeySelectionExpression *string

	// A CORS configuration. Supported only for HTTP APIs.
	CorsConfiguration *types.Cors

	// The timestamp when the API was created.
	CreatedDate *time.Time

	// The description of the API.
	Description *string

	// Specifies whether clients can invoke your API by using the default execute-api
	// endpoint. By default, clients can invoke your API with the default
	// https://{api_id}.execute-api.{region}.amazonaws.com endpoint. To require that
	// clients use a custom domain name to invoke your API, disable the default
	// endpoint.
	DisableExecuteApiEndpoint *bool

	// Avoid validating models when creating a deployment. Supported only for
	// WebSocket APIs.
	DisableSchemaValidation *bool

	// The validation information during API import. This may include particular
	// properties of your OpenAPI definition which are ignored during import. Supported
	// only for HTTP APIs.
	ImportInfo []string

	// The name of the API.
	Name *string

	// The API protocol.
	ProtocolType types.ProtocolType

	// The route selection expression for the API. For HTTP APIs, the
	// routeSelectionExpression must be ${request.method} ${request.path}. If not
	// provided, this will be the default for HTTP APIs. This property is required for
	// WebSocket APIs.
	RouteSelectionExpression *string

	// A collection of tags associated with the API.
	Tags map[string]string

	// A version identifier for the API.
	Version *string

	// The warning messages reported when failonwarnings is turned on during API
	// import.
	Warnings []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateApiMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateApi{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateApi{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateApi"); err != nil {
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
	if err = addOpCreateApiValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateApi(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateApi(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateApi",
	}
}
