// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudformation

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the estimated monthly cost of a template. The return value is an Amazon
// Web Services Simple Monthly Calculator URL with a query string that describes
// the resources required to run the template.
func (c *Client) EstimateTemplateCost(ctx context.Context, params *EstimateTemplateCostInput, optFns ...func(*Options)) (*EstimateTemplateCostOutput, error) {
	if params == nil {
		params = &EstimateTemplateCostInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "EstimateTemplateCost", params, optFns, c.addOperationEstimateTemplateCostMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*EstimateTemplateCostOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for an EstimateTemplateCost action.
type EstimateTemplateCostInput struct {

	// A list of Parameter structures that specify input parameters.
	Parameters []types.Parameter

	// Structure containing the template body with a minimum length of 1 byte and a
	// maximum length of 51,200 bytes. (For more information, go to [Template Anatomy]in the
	// CloudFormation User Guide.)
	//
	// Conditional: You must pass TemplateBody or TemplateURL . If both are passed,
	// only TemplateBody is used.
	//
	// [Template Anatomy]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/template-anatomy.html
	TemplateBody *string

	// Location of file containing the template body. The URL must point to a template
	// that's located in an Amazon S3 bucket or a Systems Manager document. For more
	// information, go to [Template Anatomy]in the CloudFormation User Guide. The location for an Amazon
	// S3 bucket must start with https:// .
	//
	// Conditional: You must pass TemplateURL or TemplateBody . If both are passed,
	// only TemplateBody is used.
	//
	// [Template Anatomy]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/template-anatomy.html
	TemplateURL *string

	noSmithyDocumentSerde
}

// The output for a EstimateTemplateCost action.
type EstimateTemplateCostOutput struct {

	// An Amazon Web Services Simple Monthly Calculator URL with a query string that
	// describes the resources required to run the template.
	Url *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationEstimateTemplateCostMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpEstimateTemplateCost{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpEstimateTemplateCost{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "EstimateTemplateCost"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opEstimateTemplateCost(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opEstimateTemplateCost(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "EstimateTemplateCost",
	}
}
