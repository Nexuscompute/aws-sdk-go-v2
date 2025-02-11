// Code generated by smithy-go-codegen DO NOT EDIT.

package accessanalyzer

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/accessanalyzer/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Checks whether new access is allowed for an updated policy when compared to the
// existing policy.
//
// You can find examples for reference policies and learn how to set up and run a
// custom policy check for new access in the [IAM Access Analyzer custom policy checks samples]repository on GitHub. The reference
// policies in this repository are meant to be passed to the existingPolicyDocument
// request parameter.
//
// [IAM Access Analyzer custom policy checks samples]: https://github.com/aws-samples/iam-access-analyzer-custom-policy-check-samples
func (c *Client) CheckNoNewAccess(ctx context.Context, params *CheckNoNewAccessInput, optFns ...func(*Options)) (*CheckNoNewAccessOutput, error) {
	if params == nil {
		params = &CheckNoNewAccessInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CheckNoNewAccess", params, optFns, c.addOperationCheckNoNewAccessMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CheckNoNewAccessOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CheckNoNewAccessInput struct {

	// The JSON policy document to use as the content for the existing policy.
	//
	// This member is required.
	ExistingPolicyDocument *string

	// The JSON policy document to use as the content for the updated policy.
	//
	// This member is required.
	NewPolicyDocument *string

	// The type of policy to compare. Identity policies grant permissions to IAM
	// principals. Identity policies include managed and inline policies for IAM roles,
	// users, and groups.
	//
	// Resource policies grant permissions on Amazon Web Services resources. Resource
	// policies include trust policies for IAM roles and bucket policies for Amazon S3
	// buckets. You can provide a generic input such as identity policy or resource
	// policy or a specific input such as managed policy or Amazon S3 bucket policy.
	//
	// This member is required.
	PolicyType types.AccessCheckPolicyType

	noSmithyDocumentSerde
}

type CheckNoNewAccessOutput struct {

	// The message indicating whether the updated policy allows new access.
	Message *string

	// A description of the reasoning of the result.
	Reasons []types.ReasonSummary

	// The result of the check for new access. If the result is PASS , no new access is
	// allowed by the updated policy. If the result is FAIL , the updated policy might
	// allow new access.
	Result types.CheckNoNewAccessResult

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCheckNoNewAccessMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCheckNoNewAccess{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCheckNoNewAccess{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CheckNoNewAccess"); err != nil {
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
	if err = addOpCheckNoNewAccessValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCheckNoNewAccess(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCheckNoNewAccess(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CheckNoNewAccess",
	}
}
