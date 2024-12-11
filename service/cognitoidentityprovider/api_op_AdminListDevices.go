// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists a user's registered devices. Remembered devices are used in
// authentication services where you offer a "Remember me" option for users who you
// want to permit to sign in without MFA from a trusted device. Users can bypass
// MFA while your application performs device SRP authentication on the back end.
// For more information, see [Working with devices].
//
// Amazon Cognito evaluates Identity and Access Management (IAM) policies in
// requests for this API operation. For this operation, you must use IAM
// credentials to authorize requests, and you must grant yourself the corresponding
// IAM permission in a policy.
//
// # Learn more
//
// [Signing Amazon Web Services API Requests]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]
//
// [Working with devices]: https://docs.aws.amazon.com/cognito/latest/developerguide/amazon-cognito-user-pools-device-tracking.html
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
// [Signing Amazon Web Services API Requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-signing.html
func (c *Client) AdminListDevices(ctx context.Context, params *AdminListDevicesInput, optFns ...func(*Options)) (*AdminListDevicesOutput, error) {
	if params == nil {
		params = &AdminListDevicesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AdminListDevices", params, optFns, c.addOperationAdminListDevicesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AdminListDevicesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the request to list devices, as an administrator.
type AdminListDevicesInput struct {

	// The ID of the user pool where the device owner is a user.
	//
	// This member is required.
	UserPoolId *string

	// The username of the user that you want to query or modify. The value of this
	// parameter is typically your user's username, but it can be any of their alias
	// attributes. If username isn't an alias attribute in your user pool, this value
	// must be the sub of a local user or the username of a user from a third-party
	// IdP.
	//
	// This member is required.
	Username *string

	// The maximum number of devices that you want Amazon Cognito to return in the
	// response.
	Limit *int32

	// This API operation returns a limited number of results. The pagination token is
	// an identifier that you can present in an additional API request with the same
	// parameters. When you include the pagination token, Amazon Cognito returns the
	// next set of items after the current list. Subsequent requests return a new
	// pagination token. By use of this token, you can paginate through the full list
	// of items.
	PaginationToken *string

	noSmithyDocumentSerde
}

// Lists the device's response, as an administrator.
type AdminListDevicesOutput struct {

	// An array of devices and their information. Each entry that's returned includes
	// device information, last-accessed and created dates, and the device key.
	Devices []types.DeviceType

	// The identifier that Amazon Cognito returned with the previous request to this
	// operation. When you include a pagination token in your request, Amazon Cognito
	// returns the next set of items in the list. By use of this token, you can
	// paginate through the full list of items.
	PaginationToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAdminListDevicesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAdminListDevices{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAdminListDevices{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "AdminListDevices"); err != nil {
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
	if err = addOpAdminListDevicesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAdminListDevices(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAdminListDevices(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "AdminListDevices",
	}
}
