// Code generated by smithy-go-codegen DO NOT EDIT.

package secretsmanager

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Validates that a resource policy does not grant a wide range of principals
// access to your secret. A resource-based policy is optional for secrets.
//
// The API performs three checks when validating the policy:
//
//   - Sends a call to [Zelkova], an automated reasoning engine, to ensure your resource
//     policy does not allow broad access to your secret, for example policies that use
//     a wildcard for the principal.
//
//   - Checks for correct syntax in a policy.
//
//   - Verifies the policy does not lock out a caller.
//
// Secrets Manager generates a CloudTrail log entry when you call this action. Do
// not include sensitive information in request parameters because it might be
// logged. For more information, see [Logging Secrets Manager events with CloudTrail].
//
// Required permissions: secretsmanager:ValidateResourcePolicy and
// secretsmanager:PutResourcePolicy . For more information, see [IAM policy actions for Secrets Manager] and [Authentication and access control in Secrets Manager].
//
// [Authentication and access control in Secrets Manager]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access.html
// [Logging Secrets Manager events with CloudTrail]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/retrieve-ct-entries.html
// [IAM policy actions for Secrets Manager]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/reference_iam-permissions.html#reference_iam-permissions_actions
// [Zelkova]: https://aws.amazon.com/blogs/security/protect-sensitive-data-in-the-cloud-with-automated-reasoning-zelkova/
func (c *Client) ValidateResourcePolicy(ctx context.Context, params *ValidateResourcePolicyInput, optFns ...func(*Options)) (*ValidateResourcePolicyOutput, error) {
	if params == nil {
		params = &ValidateResourcePolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ValidateResourcePolicy", params, optFns, c.addOperationValidateResourcePolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ValidateResourcePolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ValidateResourcePolicyInput struct {

	// A JSON-formatted string that contains an Amazon Web Services resource-based
	// policy. The policy in the string identifies who can access or manage this secret
	// and its versions. For example policies, see [Permissions policy examples].
	//
	// [Permissions policy examples]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access_examples.html
	//
	// This member is required.
	ResourcePolicy *string

	// This field is reserved for internal use.
	SecretId *string

	noSmithyDocumentSerde
}

type ValidateResourcePolicyOutput struct {

	// True if your policy passes validation, otherwise false.
	PolicyValidationPassed bool

	// Validation errors if your policy didn't pass validation.
	ValidationErrors []types.ValidationErrorsEntry

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationValidateResourcePolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpValidateResourcePolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpValidateResourcePolicy{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ValidateResourcePolicy"); err != nil {
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
	if err = addOpValidateResourcePolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opValidateResourcePolicy(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opValidateResourcePolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ValidateResourcePolicy",
	}
}
