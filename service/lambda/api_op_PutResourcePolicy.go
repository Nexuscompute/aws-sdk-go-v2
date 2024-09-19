// Code generated by smithy-go-codegen DO NOT EDIT.

package lambda

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// The option to create and modify full JSON resource-based policies, and to use
// the PutResourcePolicy, GetResourcePolicy, and DeleteResourcePolicy APIs, won't
// be available in all Amazon Web Services Regions until September 30, 2024.
//
// Adds a [resource-based policy] to a function. You can use resource-based policies to grant access to
// other [Amazon Web Services accounts], [organizations], or [services]. Resource-based policies apply to a single function, version, or
// alias.
//
// Adding a resource-based policy using this API action replaces any existing
// policy you've previously created. This means that if you've previously added
// resource-based permissions to a function using the AddPermissionaction, those permissions
// will be overwritten by your new policy.
//
// [resource-based policy]: https://docs.aws.amazon.com/lambda/latest/dg/access-control-resource-based.html
// [Amazon Web Services accounts]: https://docs.aws.amazon.com/lambda/latest/dg/permissions-function-cross-account.html
// [organizations]: https://docs.aws.amazon.com/lambda/latest/dg/permissions-function-organization.html
// [services]: https://docs.aws.amazon.com/lambda/latest/dg/permissions-function-services.html
func (c *Client) PutResourcePolicy(ctx context.Context, params *PutResourcePolicyInput, optFns ...func(*Options)) (*PutResourcePolicyOutput, error) {
	if params == nil {
		params = &PutResourcePolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutResourcePolicy", params, optFns, c.addOperationPutResourcePolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutResourcePolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutResourcePolicyInput struct {

	// The JSON resource-based policy you want to add to your function.
	//
	// To learn more about creating resource-based policies for controlling access to
	// Lambda, see [Working with resource-based IAM policies in Lambda]in the Lambda Developer Guide.
	//
	// [Working with resource-based IAM policies in Lambda]: https://docs.aws.amazon.com/
	//
	// This member is required.
	Policy *string

	// The Amazon Resource Name (ARN) of the function you want to add the policy to.
	// You can use either a qualified or an unqualified ARN, but the value you specify
	// must be a complete ARN and wildcard characters are not accepted.
	//
	// This member is required.
	ResourceArn *string

	// Replace the existing policy only if its revision ID matches the string you
	// specify. To find the revision ID of the policy currently attached to your
	// function, use the GetResourcePolicyaction.
	RevisionId *string

	noSmithyDocumentSerde
}

type PutResourcePolicyOutput struct {

	// The policy Lambda added to your function.
	Policy *string

	// The revision ID of the policy Lambda added to your function.
	RevisionId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutResourcePolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutResourcePolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutResourcePolicy{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PutResourcePolicy"); err != nil {
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
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = addOpPutResourcePolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutResourcePolicy(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutResourcePolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PutResourcePolicy",
	}
}
