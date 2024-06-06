// Code generated by smithy-go-codegen DO NOT EDIT.

package wafv2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Associates a web ACL with a regional application resource, to protect the
// resource. A regional application can be an Application Load Balancer (ALB), an
// Amazon API Gateway REST API, an AppSync GraphQL API, an Amazon Cognito user
// pool, an App Runner service, or an Amazon Web Services Verified Access instance.
//
// For Amazon CloudFront, don't use this call. Instead, use your CloudFront
// distribution configuration. To associate a web ACL, in the CloudFront call
// UpdateDistribution , set the web ACL ID to the Amazon Resource Name (ARN) of the
// web ACL. For information, see [UpdateDistribution]in the Amazon CloudFront Developer Guide.
//
// # Required permissions for customer-managed IAM policies
//
// This call requires permissions that are specific to the protected resource
// type. For details, see [Permissions for AssociateWebACL]in the WAF Developer Guide.
//
// # Temporary inconsistencies during updates
//
// When you create or change a web ACL or other WAF resources, the changes take a
// small amount of time to propagate to all areas where the resources are stored.
// The propagation time can be from a few seconds to a number of minutes.
//
// The following are examples of the temporary inconsistencies that you might
// notice during change propagation:
//
//   - After you create a web ACL, if you try to associate it with a resource, you
//     might get an exception indicating that the web ACL is unavailable.
//
//   - After you add a rule group to a web ACL, the new rule group rules might be
//     in effect in one area where the web ACL is used and not in another.
//
//   - After you change a rule action setting, you might see the old action in
//     some places and the new action in others.
//
//   - After you add an IP address to an IP set that is in use in a blocking rule,
//     the new address might be blocked in one area while still allowed in another.
//
// [UpdateDistribution]: https://docs.aws.amazon.com/cloudfront/latest/APIReference/API_UpdateDistribution.html
// [Permissions for AssociateWebACL]: https://docs.aws.amazon.com/waf/latest/developerguide/security_iam_service-with-iam.html#security_iam_action-AssociateWebACL
func (c *Client) AssociateWebACL(ctx context.Context, params *AssociateWebACLInput, optFns ...func(*Options)) (*AssociateWebACLOutput, error) {
	if params == nil {
		params = &AssociateWebACLInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AssociateWebACL", params, optFns, c.addOperationAssociateWebACLMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AssociateWebACLOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateWebACLInput struct {

	// The Amazon Resource Name (ARN) of the resource to associate with the web ACL.
	//
	// The ARN must be in one of the following formats:
	//
	//   - For an Application Load Balancer:
	//   arn:partition:elasticloadbalancing:region:account-id:loadbalancer/app/load-balancer-name/load-balancer-id
	//
	//   - For an Amazon API Gateway REST API:
	//   arn:partition:apigateway:region::/restapis/api-id/stages/stage-name
	//
	//   - For an AppSync GraphQL API:
	//   arn:partition:appsync:region:account-id:apis/GraphQLApiId
	//
	//   - For an Amazon Cognito user pool:
	//   arn:partition:cognito-idp:region:account-id:userpool/user-pool-id
	//
	//   - For an App Runner service:
	//   arn:partition:apprunner:region:account-id:service/apprunner-service-name/apprunner-service-id
	//
	//   - For an Amazon Web Services Verified Access instance:
	//   arn:partition:ec2:region:account-id:verified-access-instance/instance-id
	//
	// This member is required.
	ResourceArn *string

	// The Amazon Resource Name (ARN) of the web ACL that you want to associate with
	// the resource.
	//
	// This member is required.
	WebACLArn *string

	noSmithyDocumentSerde
}

type AssociateWebACLOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAssociateWebACLMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAssociateWebACL{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAssociateWebACL{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "AssociateWebACL"); err != nil {
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
	if err = addOpAssociateWebACLValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateWebACL(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAssociateWebACL(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "AssociateWebACL",
	}
}
