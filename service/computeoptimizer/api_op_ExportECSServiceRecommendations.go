// Code generated by smithy-go-codegen DO NOT EDIT.

package computeoptimizer

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/computeoptimizer/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

//	Exports optimization recommendations for Amazon ECS services on Fargate.
//
// Recommendations are exported in a CSV file, and its metadata in a JSON file, to
// an existing Amazon Simple Storage Service (Amazon S3) bucket that you specify.
// For more information, see [Exporting Recommendations]in the Compute Optimizer User Guide.
//
// You can only have one Amazon ECS service export job in progress per Amazon Web
// Services Region.
//
// [Exporting Recommendations]: https://docs.aws.amazon.com/compute-optimizer/latest/ug/exporting-recommendations.html
func (c *Client) ExportECSServiceRecommendations(ctx context.Context, params *ExportECSServiceRecommendationsInput, optFns ...func(*Options)) (*ExportECSServiceRecommendationsOutput, error) {
	if params == nil {
		params = &ExportECSServiceRecommendationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ExportECSServiceRecommendations", params, optFns, c.addOperationExportECSServiceRecommendationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ExportECSServiceRecommendationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ExportECSServiceRecommendationsInput struct {

	// Describes the destination Amazon Simple Storage Service (Amazon S3) bucket name
	// and key prefix for a recommendations export job.
	//
	// You must create the destination Amazon S3 bucket for your recommendations
	// export before you create the export job. Compute Optimizer does not create the
	// S3 bucket for you. After you create the S3 bucket, ensure that it has the
	// required permission policy to allow Compute Optimizer to write the export file
	// to it. If you plan to specify an object prefix when you create the export job,
	// you must include the object prefix in the policy that you add to the S3 bucket.
	// For more information, see [Amazon S3 Bucket Policy for Compute Optimizer]in the Compute Optimizer User Guide.
	//
	// [Amazon S3 Bucket Policy for Compute Optimizer]: https://docs.aws.amazon.com/compute-optimizer/latest/ug/create-s3-bucket-policy-for-compute-optimizer.html
	//
	// This member is required.
	S3DestinationConfig *types.S3DestinationConfig

	//  The Amazon Web Services account IDs for the export Amazon ECS service
	// recommendations.
	//
	// If your account is the management account or the delegated administrator of an
	// organization, use this parameter to specify the member account you want to
	// export recommendations to.
	//
	// This parameter can't be specified together with the include member accounts
	// parameter. The parameters are mutually exclusive.
	//
	// If this parameter or the include member accounts parameter is omitted, the
	// recommendations for member accounts aren't included in the export.
	//
	// You can specify multiple account IDs per request.
	AccountIds []string

	// The recommendations data to include in the export file. For more information
	// about the fields that can be exported, see [Exported files]in the Compute Optimizer User Guide.
	//
	// [Exported files]: https://docs.aws.amazon.com/compute-optimizer/latest/ug/exporting-recommendations.html#exported-files
	FieldsToExport []types.ExportableECSServiceField

	//  The format of the export file.
	//
	// The CSV file is the only export file format currently supported.
	FileFormat types.FileFormat

	//  An array of objects to specify a filter that exports a more specific set of
	// Amazon ECS service recommendations.
	Filters []types.ECSServiceRecommendationFilter

	// If your account is the management account or the delegated administrator of an
	// organization, this parameter indicates whether to include recommendations for
	// resources in all member accounts of the organization.
	//
	// The member accounts must also be opted in to Compute Optimizer, and trusted
	// access for Compute Optimizer must be enabled in the organization account. For
	// more information, see [Compute Optimizer and Amazon Web Services Organizations trusted access]in the Compute Optimizer User Guide.
	//
	// If this parameter is omitted, recommendations for member accounts of the
	// organization aren't included in the export file.
	//
	// If this parameter or the account ID parameter is omitted, recommendations for
	// member accounts aren't included in the export.
	//
	// [Compute Optimizer and Amazon Web Services Organizations trusted access]: https://docs.aws.amazon.com/compute-optimizer/latest/ug/security-iam.html#trusted-service-access
	IncludeMemberAccounts bool

	noSmithyDocumentSerde
}

type ExportECSServiceRecommendationsOutput struct {

	//  The identification number of the export job.
	//
	// To view the status of an export job, use the DescribeRecommendationExportJobs action and specify the job ID.
	JobId *string

	// Describes the destination Amazon Simple Storage Service (Amazon S3) bucket name
	// and object keys of a recommendations export file, and its associated metadata
	// file.
	S3Destination *types.S3Destination

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationExportECSServiceRecommendationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpExportECSServiceRecommendations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpExportECSServiceRecommendations{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ExportECSServiceRecommendations"); err != nil {
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
	if err = addOpExportECSServiceRecommendationsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opExportECSServiceRecommendations(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opExportECSServiceRecommendations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ExportECSServiceRecommendations",
	}
}
