// Code generated by smithy-go-codegen DO NOT EDIT.

package globalaccelerator

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/globalaccelerator/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Update a cross-account attachment to add or remove principals or resources.
// When you update an attachment to remove a principal (account ID or accelerator)
// or a resource, Global Accelerator revokes the permission for specific resources.
//
// For more information, see [Working with cross-account attachments and resources in Global Accelerator] in the Global Accelerator Developer Guide.
//
// [Working with cross-account attachments and resources in Global Accelerator]: https://docs.aws.amazon.com/global-accelerator/latest/dg/cross-account-resources.html
func (c *Client) UpdateCrossAccountAttachment(ctx context.Context, params *UpdateCrossAccountAttachmentInput, optFns ...func(*Options)) (*UpdateCrossAccountAttachmentOutput, error) {
	if params == nil {
		params = &UpdateCrossAccountAttachmentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateCrossAccountAttachment", params, optFns, c.addOperationUpdateCrossAccountAttachmentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateCrossAccountAttachmentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateCrossAccountAttachmentInput struct {

	// The Amazon Resource Name (ARN) of the cross-account attachment to update.
	//
	// This member is required.
	AttachmentArn *string

	// The principals to add to the cross-account attachment. A principal is an
	// account or the Amazon Resource Name (ARN) of an accelerator that the attachment
	// gives permission to work with resources from another account. The resources are
	// also listed in the attachment.
	//
	// To add more than one principal, separate the account numbers or accelerator
	// ARNs, or both, with commas.
	AddPrincipals []string

	// The resources to add to the cross-account attachment. A resource listed in a
	// cross-account attachment can be used with an accelerator by the principals that
	// are listed in the attachment.
	//
	// To add more than one resource, separate the resource ARNs with commas.
	AddResources []types.Resource

	// The name of the cross-account attachment.
	Name *string

	// The principals to remove from the cross-account attachment. A principal is an
	// account or the Amazon Resource Name (ARN) of an accelerator that the attachment
	// gives permission to work with resources from another account. The resources are
	// also listed in the attachment.
	//
	// To remove more than one principal, separate the account numbers or accelerator
	// ARNs, or both, with commas.
	RemovePrincipals []string

	// The resources to remove from the cross-account attachment. A resource listed in
	// a cross-account attachment can be used with an accelerator by the principals
	// that are listed in the attachment.
	//
	// To remove more than one resource, separate the resource ARNs with commas.
	RemoveResources []types.Resource

	noSmithyDocumentSerde
}

type UpdateCrossAccountAttachmentOutput struct {

	// Information about the updated cross-account attachment.
	CrossAccountAttachment *types.Attachment

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateCrossAccountAttachmentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateCrossAccountAttachment{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateCrossAccountAttachment{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateCrossAccountAttachment"); err != nil {
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
	if err = addOpUpdateCrossAccountAttachmentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateCrossAccountAttachment(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateCrossAccountAttachment(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateCrossAccountAttachment",
	}
}
