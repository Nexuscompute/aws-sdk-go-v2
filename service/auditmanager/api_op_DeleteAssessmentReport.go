// Code generated by smithy-go-codegen DO NOT EDIT.

package auditmanager

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes an assessment report in Audit Manager.
//
// When you run the DeleteAssessmentReport operation, Audit Manager attempts to
// delete the following data:
//
//   - The specified assessment report that’s stored in your S3 bucket
//
//   - The associated metadata that’s stored in Audit Manager
//
// If Audit Manager can’t access the assessment report in your S3 bucket, the
// report isn’t deleted. In this event, the DeleteAssessmentReport operation
// doesn’t fail. Instead, it proceeds to delete the associated metadata only. You
// must then delete the assessment report from the S3 bucket yourself.
//
// This scenario happens when Audit Manager receives a 403 (Forbidden) or 404 (Not
// Found) error from Amazon S3. To avoid this, make sure that your S3 bucket is
// available, and that you configured the correct permissions for Audit Manager to
// delete resources in your S3 bucket. For an example permissions policy that you
// can use, see [Assessment report destination permissions]in the Audit Manager User Guide. For information about the issues
// that could cause a 403 (Forbidden) or 404 (Not Found ) error from Amazon S3, see [List of Error Codes]
// in the Amazon Simple Storage Service API Reference.
//
// [List of Error Codes]: https://docs.aws.amazon.com/AmazonS3/latest/API/ErrorResponses.html#ErrorCodeList
// [Assessment report destination permissions]: https://docs.aws.amazon.com/audit-manager/latest/userguide/security_iam_id-based-policy-examples.html#full-administrator-access-assessment-report-destination
func (c *Client) DeleteAssessmentReport(ctx context.Context, params *DeleteAssessmentReportInput, optFns ...func(*Options)) (*DeleteAssessmentReportOutput, error) {
	if params == nil {
		params = &DeleteAssessmentReportInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteAssessmentReport", params, optFns, c.addOperationDeleteAssessmentReportMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteAssessmentReportOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteAssessmentReportInput struct {

	//  The unique identifier for the assessment.
	//
	// This member is required.
	AssessmentId *string

	//  The unique identifier for the assessment report.
	//
	// This member is required.
	AssessmentReportId *string

	noSmithyDocumentSerde
}

type DeleteAssessmentReportOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteAssessmentReportMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteAssessmentReport{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteAssessmentReport{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeleteAssessmentReport"); err != nil {
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
	if err = addOpDeleteAssessmentReportValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteAssessmentReport(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteAssessmentReport(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeleteAssessmentReport",
	}
}
