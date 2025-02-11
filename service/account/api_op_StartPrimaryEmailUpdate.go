// Code generated by smithy-go-codegen DO NOT EDIT.

package account

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/account/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Starts the process to update the primary email address for the specified
// account.
func (c *Client) StartPrimaryEmailUpdate(ctx context.Context, params *StartPrimaryEmailUpdateInput, optFns ...func(*Options)) (*StartPrimaryEmailUpdateOutput, error) {
	if params == nil {
		params = &StartPrimaryEmailUpdateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartPrimaryEmailUpdate", params, optFns, c.addOperationStartPrimaryEmailUpdateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartPrimaryEmailUpdateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartPrimaryEmailUpdateInput struct {

	// Specifies the 12-digit account ID number of the Amazon Web Services account
	// that you want to access or modify with this operation. To use this parameter,
	// the caller must be an identity in the [organization's management account]or a delegated administrator account. The
	// specified account ID must be a member account in the same organization. The
	// organization must have [all features enabled], and the organization must have [trusted access] enabled for the
	// Account Management service, and optionally a [delegated admin]account assigned.
	//
	// This operation can only be called from the management account or the delegated
	// administrator account of an organization for a member account.
	//
	// The management account can't specify its own AccountId .
	//
	// [organization's management account]: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_getting-started_concepts.html#account
	// [all features enabled]: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_org_support-all-features.html
	// [delegated admin]: https://docs.aws.amazon.com/organizations/latest/userguide/using-orgs-delegated-admin.html
	// [trusted access]: https://docs.aws.amazon.com/organizations/latest/userguide/using-orgs-trusted-access.html
	//
	// This member is required.
	AccountId *string

	// The new primary email address (also known as the root user email address) to
	// use in the specified account.
	//
	// This member is required.
	PrimaryEmail *string

	noSmithyDocumentSerde
}

type StartPrimaryEmailUpdateOutput struct {

	// The status of the primary email update request.
	Status types.PrimaryEmailUpdateStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartPrimaryEmailUpdateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartPrimaryEmailUpdate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartPrimaryEmailUpdate{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "StartPrimaryEmailUpdate"); err != nil {
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
	if err = addOpStartPrimaryEmailUpdateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartPrimaryEmailUpdate(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartPrimaryEmailUpdate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "StartPrimaryEmailUpdate",
	}
}
