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

// Deletes the specified alternate contact from an Amazon Web Services account.
//
// For complete details about how to use the alternate contact operations, see [Access or updating the alternate contacts].
//
// Before you can update the alternate contact information for an Amazon Web
// Services account that is managed by Organizations, you must first enable
// integration between Amazon Web Services Account Management and Organizations.
// For more information, see [Enabling trusted access for Amazon Web Services Account Management].
//
// [Access or updating the alternate contacts]: https://docs.aws.amazon.com/accounts/latest/reference/manage-acct-update-contact.html
// [Enabling trusted access for Amazon Web Services Account Management]: https://docs.aws.amazon.com/accounts/latest/reference/using-orgs-trusted-access.html
func (c *Client) DeleteAlternateContact(ctx context.Context, params *DeleteAlternateContactInput, optFns ...func(*Options)) (*DeleteAlternateContactOutput, error) {
	if params == nil {
		params = &DeleteAlternateContactInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteAlternateContact", params, optFns, c.addOperationDeleteAlternateContactMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteAlternateContactOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteAlternateContactInput struct {

	// Specifies which of the alternate contacts to delete.
	//
	// This member is required.
	AlternateContactType types.AlternateContactType

	// Specifies the 12 digit account ID number of the Amazon Web Services account
	// that you want to access or modify with this operation.
	//
	// If you do not specify this parameter, it defaults to the Amazon Web Services
	// account of the identity used to call the operation.
	//
	// To use this parameter, the caller must be an identity in the [organization's management account] or a delegated
	// administrator account, and the specified account ID must be a member account in
	// the same organization. The organization must have [all features enabled], and the organization must
	// have [trusted access]enabled for the Account Management service, and optionally a [delegated admin] account
	// assigned.
	//
	// The management account can't specify its own AccountId ; it must call the
	// operation in standalone context by not including the AccountId parameter.
	//
	// To call this operation on an account that is not a member of an organization,
	// then don't specify this parameter, and call the operation using an identity
	// belonging to the account whose contacts you wish to retrieve or modify.
	//
	// [organization's management account]: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_getting-started_concepts.html#account
	// [all features enabled]: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_org_support-all-features.html
	// [delegated admin]: https://docs.aws.amazon.com/organizations/latest/userguide/using-orgs-delegated-admin.html
	// [trusted access]: https://docs.aws.amazon.com/organizations/latest/userguide/using-orgs-trusted-access.html
	AccountId *string

	noSmithyDocumentSerde
}

type DeleteAlternateContactOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteAlternateContactMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteAlternateContact{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteAlternateContact{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeleteAlternateContact"); err != nil {
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
	if err = addOpDeleteAlternateContactValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteAlternateContact(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteAlternateContact(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeleteAlternateContact",
	}
}
