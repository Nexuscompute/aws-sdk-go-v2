// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Describes an existing import job.
//
// Poll job descriptions after starting a job to know when it has succeeded or
// failed. Job descriptions are available for 14 days after job starts.
func (c *Client) DescribeAssetBundleImportJob(ctx context.Context, params *DescribeAssetBundleImportJobInput, optFns ...func(*Options)) (*DescribeAssetBundleImportJobOutput, error) {
	if params == nil {
		params = &DescribeAssetBundleImportJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeAssetBundleImportJob", params, optFns, c.addOperationDescribeAssetBundleImportJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeAssetBundleImportJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAssetBundleImportJobInput struct {

	// The ID of the job. The job ID is set when you start a new job with a
	// StartAssetBundleImportJob API call.
	//
	// This member is required.
	AssetBundleImportJobId *string

	// The ID of the Amazon Web Services account the import job was executed in.
	//
	// This member is required.
	AwsAccountId *string

	noSmithyDocumentSerde
}

type DescribeAssetBundleImportJobOutput struct {

	// The Amazon Resource Name (ARN) for the import job.
	Arn *string

	// The ID of the job. The job ID is set when you start a new job with a
	// StartAssetBundleImportJob API call.
	AssetBundleImportJobId *string

	// The source of the asset bundle zip file that contains the data that is imported
	// by the job.
	AssetBundleImportSource *types.AssetBundleImportSourceDescription

	// The ID of the Amazon Web Services account the import job was executed in.
	AwsAccountId *string

	// The time that the import job was created.
	CreatedTime *time.Time

	// An array of error records that describes any failures that occurred during the
	// export job processing.
	//
	// Error records accumulate while the job is still running. The complete set of
	// error records is available after the job has completed and failed.
	Errors []types.AssetBundleImportJobError

	// The failure action for the import job.
	FailureAction types.AssetBundleImportFailureAction

	// Indicates the status of a job through its queuing and execution.
	//
	// Poll the DescribeAssetBundleImport API until JobStatus returns one of the
	// following values:
	//
	//   - SUCCESSFUL
	//
	//   - FAILED
	//
	//   - FAILED_ROLLBACK_COMPLETED
	//
	//   - FAILED_ROLLBACK_ERROR
	JobStatus types.AssetBundleImportJobStatus

	// Optional overrides that are applied to the resource configuration before import.
	OverrideParameters *types.AssetBundleImportJobOverrideParameters

	// Optional permission overrides that are applied to the resource configuration
	// before import.
	OverridePermissions *types.AssetBundleImportJobOverridePermissions

	// Optional tag overrides that are applied to the resource configuration before
	// import.
	OverrideTags *types.AssetBundleImportJobOverrideTags

	// An optional validation strategy override for all analyses and dashboards to be
	// applied to the resource configuration before import.
	OverrideValidationStrategy *types.AssetBundleImportJobOverrideValidationStrategy

	// The Amazon Web Services request ID for this operation.
	RequestId *string

	// An array of error records that describes any failures that occurred while an
	// import job was attempting a rollback.
	//
	// Error records accumulate while the job is still running. The complete set of
	// error records is available after the job has completed and failed.
	RollbackErrors []types.AssetBundleImportJobError

	// The HTTP status of the response.
	Status int32

	// An array of warning records that describe all permitted errors that are
	// encountered during the import job.
	Warnings []types.AssetBundleImportJobWarning

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeAssetBundleImportJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeAssetBundleImportJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeAssetBundleImportJob{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeAssetBundleImportJob"); err != nil {
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
	if err = addOpDescribeAssetBundleImportJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAssetBundleImportJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeAssetBundleImportJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeAssetBundleImportJob",
	}
}
