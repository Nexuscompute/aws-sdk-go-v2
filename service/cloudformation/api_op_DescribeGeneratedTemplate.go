// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudformation

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Describes a generated template. The output includes details about the progress
// of the creation of a generated template started by a CreateGeneratedTemplate
// API action or the update of a generated template started with an
// UpdateGeneratedTemplate API action.
func (c *Client) DescribeGeneratedTemplate(ctx context.Context, params *DescribeGeneratedTemplateInput, optFns ...func(*Options)) (*DescribeGeneratedTemplateOutput, error) {
	if params == nil {
		params = &DescribeGeneratedTemplateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeGeneratedTemplate", params, optFns, c.addOperationDescribeGeneratedTemplateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeGeneratedTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeGeneratedTemplateInput struct {

	// The name or Amazon Resource Name (ARN) of a generated template.
	//
	// This member is required.
	GeneratedTemplateName *string

	noSmithyDocumentSerde
}

type DescribeGeneratedTemplateOutput struct {

	// The time the generated template was created.
	CreationTime *time.Time

	// The Amazon Resource Name (ARN) of the generated template. The format is
	// arn:${Partition}:cloudformation:${Region}:${Account}:generatedtemplate/${Id} .
	// For example,
	// arn:aws:cloudformation:us-east-1:123456789012:generatedtemplate/2e8465c1-9a80-43ea-a3a3-4f2d692fe6dc
	// .
	GeneratedTemplateId *string

	// The name of the generated template.
	GeneratedTemplateName *string

	// The time the generated template was last updated.
	LastUpdatedTime *time.Time

	// An object describing the progress of the template generation.
	Progress *types.TemplateProgress

	// A list of objects describing the details of the resources in the template
	// generation.
	Resources []types.ResourceDetail

	// The stack ARN of the base stack if a base stack was provided when generating
	// the template.
	StackId *string

	// The status of the template generation. Supported values are:
	//
	//   - CreatePending - the creation of the template is pending.
	//
	//   - CreateInProgress - the creation of the template is in progress.
	//
	//   - DeletePending - the deletion of the template is pending.
	//
	//   - DeleteInProgress - the deletion of the template is in progress.
	//
	//   - UpdatePending - the update of the template is pending.
	//
	//   - UpdateInProgress - the update of the template is in progress.
	//
	//   - Failed - the template operation failed.
	//
	//   - Complete - the template operation is complete.
	Status types.GeneratedTemplateStatus

	// The reason for the current template generation status. This will provide more
	// details if a failure happened.
	StatusReason *string

	// The configuration details of the generated template, including the
	// DeletionPolicy and UpdateReplacePolicy .
	TemplateConfiguration *types.TemplateConfiguration

	// The number of warnings generated for this template. The warnings are found in
	// the details of each of the resources in the template.
	TotalWarnings *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeGeneratedTemplateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeGeneratedTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeGeneratedTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeGeneratedTemplate"); err != nil {
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
	if err = addOpDescribeGeneratedTemplateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeGeneratedTemplate(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeGeneratedTemplate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeGeneratedTemplate",
	}
}
