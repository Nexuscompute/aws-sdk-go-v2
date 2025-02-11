// Code generated by smithy-go-codegen DO NOT EDIT.

package qbusiness

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/qbusiness/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Adds one or more documents to an Amazon Q Business index.
//
// You use this API to:
//
//   - ingest your structured and unstructured documents and documents stored in
//     an Amazon S3 bucket into an Amazon Q Business index.
//
//   - add custom attributes to documents in an Amazon Q Business index.
//
//   - attach an access control list to the documents added to an Amazon Q
//     Business index.
//
// You can see the progress of the deletion, and any error messages related to the
// process, by using CloudWatch.
func (c *Client) BatchPutDocument(ctx context.Context, params *BatchPutDocumentInput, optFns ...func(*Options)) (*BatchPutDocumentOutput, error) {
	if params == nil {
		params = &BatchPutDocumentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchPutDocument", params, optFns, c.addOperationBatchPutDocumentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchPutDocumentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchPutDocumentInput struct {

	// The identifier of the Amazon Q Business application.
	//
	// This member is required.
	ApplicationId *string

	// One or more documents to add to the index.
	//
	// This member is required.
	Documents []types.Document

	// The identifier of the Amazon Q Business index to add the documents to.
	//
	// This member is required.
	IndexId *string

	// The identifier of the data source sync during which the documents were added.
	DataSourceSyncId *string

	// The Amazon Resource Name (ARN) of an IAM role with permission to access your S3
	// bucket.
	RoleArn *string

	noSmithyDocumentSerde
}

type BatchPutDocumentOutput struct {

	//  A list of documents that were not added to the Amazon Q Business index because
	// the document failed a validation check. Each document contains an error message
	// that indicates why the document couldn't be added to the index.
	FailedDocuments []types.FailedDocument

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationBatchPutDocumentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpBatchPutDocument{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpBatchPutDocument{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "BatchPutDocument"); err != nil {
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
	if err = addOpBatchPutDocumentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchPutDocument(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opBatchPutDocument(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "BatchPutDocument",
	}
}
