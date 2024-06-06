// Code generated by smithy-go-codegen DO NOT EDIT.

package migrationhubstrategy

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/migrationhubstrategy/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Retrieves the details about a specific import task.
func (c *Client) GetImportFileTask(ctx context.Context, params *GetImportFileTaskInput, optFns ...func(*Options)) (*GetImportFileTaskOutput, error) {
	if params == nil {
		params = &GetImportFileTaskInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetImportFileTask", params, optFns, c.addOperationGetImportFileTaskMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetImportFileTaskOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetImportFileTaskInput struct {

	//  The ID of the import file task. This ID is returned in the response of StartImportFileTask.
	//
	// This member is required.
	Id *string

	noSmithyDocumentSerde
}

type GetImportFileTaskOutput struct {

	//  The time that the import task completed.
	CompletionTime *time.Time

	//  The import file task id returned in the response of StartImportFileTask.
	Id *string

	//  The name of the import task given in StartImportFileTask.
	ImportName *string

	//  The S3 bucket where import file is located.
	InputS3Bucket *string

	//  The Amazon S3 key name of the import file.
	InputS3Key *string

	//  The number of records that failed to be imported.
	NumberOfRecordsFailed *int32

	//  The number of records successfully imported.
	NumberOfRecordsSuccess *int32

	//  Start time of the import task.
	StartTime *time.Time

	//  Status of import file task.
	Status types.ImportFileTaskStatus

	//  The S3 bucket name for status report of import task.
	StatusReportS3Bucket *string

	//  The Amazon S3 key name for status report of import task. The report contains
	// details about whether each record imported successfully or why it did not.
	StatusReportS3Key *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetImportFileTaskMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetImportFileTask{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetImportFileTask{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetImportFileTask"); err != nil {
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
	if err = addOpGetImportFileTaskValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetImportFileTask(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetImportFileTask(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetImportFileTask",
	}
}
