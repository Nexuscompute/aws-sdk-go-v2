// Code generated by smithy-go-codegen DO NOT EDIT.

package redshiftserverless

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/redshiftserverless/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates a namespace with the specified settings. Unless required, you can't
// update multiple parameters in one request. For example, you must specify both
// adminUsername and adminUserPassword to update either field, but you can't
// update both kmsKeyId and logExports in a single request.
func (c *Client) UpdateNamespace(ctx context.Context, params *UpdateNamespaceInput, optFns ...func(*Options)) (*UpdateNamespaceOutput, error) {
	if params == nil {
		params = &UpdateNamespaceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateNamespace", params, optFns, c.addOperationUpdateNamespaceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateNamespaceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateNamespaceInput struct {

	// The name of the namespace to update. You can't update the name of a namespace
	// once it is created.
	//
	// This member is required.
	NamespaceName *string

	// The ID of the Key Management Service (KMS) key used to encrypt and store the
	// namespace's admin credentials secret. You can only use this parameter if
	// manageAdminPassword is true.
	AdminPasswordSecretKmsKeyId *string

	// The password of the administrator for the first database created in the
	// namespace. This parameter must be updated together with adminUsername .
	//
	// You can't use adminUserPassword if manageAdminPassword is true.
	AdminUserPassword *string

	// The username of the administrator for the first database created in the
	// namespace. This parameter must be updated together with adminUserPassword .
	AdminUsername *string

	// The Amazon Resource Name (ARN) of the IAM role to set as a default in the
	// namespace. This parameter must be updated together with iamRoles .
	DefaultIamRoleArn *string

	// A list of IAM roles to associate with the namespace. This parameter must be
	// updated together with defaultIamRoleArn .
	IamRoles []string

	// The ID of the Amazon Web Services Key Management Service key used to encrypt
	// your data.
	KmsKeyId *string

	// The types of logs the namespace can export. The export types are userlog ,
	// connectionlog , and useractivitylog .
	LogExports []types.LogExport

	// If true , Amazon Redshift uses Secrets Manager to manage the namespace's admin
	// credentials. You can't use adminUserPassword if manageAdminPassword is true. If
	// manageAdminPassword is false or not set, Amazon Redshift uses adminUserPassword
	// for the admin user account's password.
	ManageAdminPassword *bool

	noSmithyDocumentSerde
}

type UpdateNamespaceOutput struct {

	// A list of tag instances.
	//
	// This member is required.
	Namespace *types.Namespace

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateNamespaceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateNamespace{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateNamespace{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateNamespace"); err != nil {
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
	if err = addOpUpdateNamespaceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateNamespace(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateNamespace(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateNamespace",
	}
}
