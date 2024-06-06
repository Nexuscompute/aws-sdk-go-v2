// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudtrail

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Registers an organization’s member account as the CloudTrail [delegated administrator].
//
// [delegated administrator]: https://docs.aws.amazon.com/awscloudtrail/latest/userguide/cloudtrail-delegated-administrator.html
func (c *Client) RegisterOrganizationDelegatedAdmin(ctx context.Context, params *RegisterOrganizationDelegatedAdminInput, optFns ...func(*Options)) (*RegisterOrganizationDelegatedAdminOutput, error) {
	if params == nil {
		params = &RegisterOrganizationDelegatedAdminInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RegisterOrganizationDelegatedAdmin", params, optFns, c.addOperationRegisterOrganizationDelegatedAdminMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RegisterOrganizationDelegatedAdminOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Specifies an organization member account ID as a CloudTrail delegated
// administrator.
type RegisterOrganizationDelegatedAdminInput struct {

	// An organization member account ID that you want to designate as a delegated
	// administrator.
	//
	// This member is required.
	MemberAccountId *string

	noSmithyDocumentSerde
}

// Returns the following response if successful. Otherwise, returns an error.
type RegisterOrganizationDelegatedAdminOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRegisterOrganizationDelegatedAdminMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpRegisterOrganizationDelegatedAdmin{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpRegisterOrganizationDelegatedAdmin{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "RegisterOrganizationDelegatedAdmin"); err != nil {
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
	if err = addOpRegisterOrganizationDelegatedAdminValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRegisterOrganizationDelegatedAdmin(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRegisterOrganizationDelegatedAdmin(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "RegisterOrganizationDelegatedAdmin",
	}
}
