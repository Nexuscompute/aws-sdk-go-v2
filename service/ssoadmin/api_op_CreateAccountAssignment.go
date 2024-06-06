// Code generated by smithy-go-codegen DO NOT EDIT.

package ssoadmin

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/ssoadmin/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Assigns access to a principal for a specified Amazon Web Services account using
// a specified permission set.
//
// The term principal here refers to a user or group that is defined in IAM
// Identity Center.
//
// As part of a successful CreateAccountAssignment call, the specified permission
// set will automatically be provisioned to the account in the form of an IAM
// policy. That policy is attached to the IAM role created in IAM Identity Center.
// If the permission set is subsequently updated, the corresponding IAM policies
// attached to roles in your accounts will not be updated automatically. In this
// case, you must call ProvisionPermissionSetto make these updates.
//
// After a successful response, call DescribeAccountAssignmentCreationStatus to
// describe the status of an assignment creation request.
func (c *Client) CreateAccountAssignment(ctx context.Context, params *CreateAccountAssignmentInput, optFns ...func(*Options)) (*CreateAccountAssignmentOutput, error) {
	if params == nil {
		params = &CreateAccountAssignmentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateAccountAssignment", params, optFns, c.addOperationCreateAccountAssignmentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateAccountAssignmentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateAccountAssignmentInput struct {

	// The ARN of the IAM Identity Center instance under which the operation will be
	// executed. For more information about ARNs, see Amazon Resource Names (ARNs) and Amazon Web Services Service Namespacesin the Amazon Web Services
	// General Reference.
	//
	// This member is required.
	InstanceArn *string

	// The ARN of the permission set that the admin wants to grant the principal
	// access to.
	//
	// This member is required.
	PermissionSetArn *string

	// An identifier for an object in IAM Identity Center, such as a user or group.
	// PrincipalIds are GUIDs (For example, f81d4fae-7dec-11d0-a765-00a0c91e6bf6). For
	// more information about PrincipalIds in IAM Identity Center, see the IAM Identity Center Identity Store API Reference.
	//
	// This member is required.
	PrincipalId *string

	// The entity type for which the assignment will be created.
	//
	// This member is required.
	PrincipalType types.PrincipalType

	// TargetID is an Amazon Web Services account identifier, (For example,
	// 123456789012).
	//
	// This member is required.
	TargetId *string

	// The entity type for which the assignment will be created.
	//
	// This member is required.
	TargetType types.TargetType

	noSmithyDocumentSerde
}

type CreateAccountAssignmentOutput struct {

	// The status object for the account assignment creation operation.
	AccountAssignmentCreationStatus *types.AccountAssignmentOperationStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateAccountAssignmentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateAccountAssignment{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateAccountAssignment{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateAccountAssignment"); err != nil {
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
	if err = addOpCreateAccountAssignmentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateAccountAssignment(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateAccountAssignment(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateAccountAssignment",
	}
}
