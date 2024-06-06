// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new work team for labeling your data. A work team is defined by one
// or more Amazon Cognito user pools. You must first create the user pools before
// you can create a work team.
//
// You cannot create more than 25 work teams in an account and region.
func (c *Client) CreateWorkteam(ctx context.Context, params *CreateWorkteamInput, optFns ...func(*Options)) (*CreateWorkteamOutput, error) {
	if params == nil {
		params = &CreateWorkteamInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateWorkteam", params, optFns, c.addOperationCreateWorkteamMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateWorkteamOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateWorkteamInput struct {

	// A description of the work team.
	//
	// This member is required.
	Description *string

	// A list of MemberDefinition objects that contains objects that identify the
	// workers that make up the work team.
	//
	// Workforces can be created using Amazon Cognito or your own OIDC Identity
	// Provider (IdP). For private workforces created using Amazon Cognito use
	// CognitoMemberDefinition . For workforces created using your own OIDC identity
	// provider (IdP) use OidcMemberDefinition . Do not provide input for both of these
	// parameters in a single request.
	//
	// For workforces created using Amazon Cognito, private work teams correspond to
	// Amazon Cognito user groups within the user pool used to create a workforce. All
	// of the CognitoMemberDefinition objects that make up the member definition must
	// have the same ClientId and UserPool values. To add a Amazon Cognito user group
	// to an existing worker pool, see Adding groups to a User Pool. For more information about user pools, see [Amazon Cognito User Pools].
	//
	// For workforces created using your own OIDC IdP, specify the user groups that
	// you want to include in your private work team in OidcMemberDefinition by
	// listing those groups in Groups .
	//
	// [Amazon Cognito User Pools]: https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-identity-pools.html
	//
	// This member is required.
	MemberDefinitions []types.MemberDefinition

	// The name of the work team. Use this name to identify the work team.
	//
	// This member is required.
	WorkteamName *string

	// Configures notification of workers regarding available or expiring work items.
	NotificationConfiguration *types.NotificationConfiguration

	// An array of key-value pairs.
	//
	// For more information, see [Resource Tag] and [Using Cost Allocation Tags] in the Amazon Web Services Billing and Cost
	// Management User Guide.
	//
	// [Resource Tag]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html
	// [Using Cost Allocation Tags]: https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html#allocation-what
	Tags []types.Tag

	// Use this optional parameter to constrain access to an Amazon S3 resource based
	// on the IP address using supported IAM global condition keys. The Amazon S3
	// resource is accessed in the worker portal using a Amazon S3 presigned URL.
	WorkerAccessConfiguration *types.WorkerAccessConfiguration

	// The name of the workforce.
	WorkforceName *string

	noSmithyDocumentSerde
}

type CreateWorkteamOutput struct {

	// The Amazon Resource Name (ARN) of the work team. You can use this ARN to
	// identify the work team.
	WorkteamArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateWorkteamMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateWorkteam{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateWorkteam{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateWorkteam"); err != nil {
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
	if err = addOpCreateWorkteamValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateWorkteam(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateWorkteam(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateWorkteam",
	}
}
