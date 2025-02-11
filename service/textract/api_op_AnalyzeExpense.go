// Code generated by smithy-go-codegen DO NOT EDIT.

package textract

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/textract/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// AnalyzeExpense synchronously analyzes an input document for financially related
// relationships between text.
//
// Information is returned as ExpenseDocuments and seperated as follows:
//
//   - LineItemGroups - A data set containing LineItems which store information
//     about the lines of text, such as an item purchased and its price on a receipt.
//
//   - SummaryFields - Contains all other information a receipt, such as header
//     information or the vendors name.
func (c *Client) AnalyzeExpense(ctx context.Context, params *AnalyzeExpenseInput, optFns ...func(*Options)) (*AnalyzeExpenseOutput, error) {
	if params == nil {
		params = &AnalyzeExpenseInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AnalyzeExpense", params, optFns, c.addOperationAnalyzeExpenseMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AnalyzeExpenseOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AnalyzeExpenseInput struct {

	// The input document, either as bytes or as an S3 object.
	//
	// You pass image bytes to an Amazon Textract API operation by using the Bytes
	// property. For example, you would use the Bytes property to pass a document
	// loaded from a local file system. Image bytes passed by using the Bytes property
	// must be base64 encoded. Your code might not need to encode document file bytes
	// if you're using an AWS SDK to call Amazon Textract API operations.
	//
	// You pass images stored in an S3 bucket to an Amazon Textract API operation by
	// using the S3Object property. Documents stored in an S3 bucket don't need to be
	// base64 encoded.
	//
	// The AWS Region for the S3 bucket that contains the S3 object must match the AWS
	// Region that you use for Amazon Textract operations.
	//
	// If you use the AWS CLI to call Amazon Textract operations, passing image bytes
	// using the Bytes property isn't supported. You must first upload the document to
	// an Amazon S3 bucket, and then call the operation using the S3Object property.
	//
	// For Amazon Textract to process an S3 object, the user must have permission to
	// access the S3 object.
	//
	// This member is required.
	Document *types.Document

	noSmithyDocumentSerde
}

type AnalyzeExpenseOutput struct {

	// Information about the input document.
	DocumentMetadata *types.DocumentMetadata

	// The expenses detected by Amazon Textract.
	ExpenseDocuments []types.ExpenseDocument

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAnalyzeExpenseMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAnalyzeExpense{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAnalyzeExpense{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "AnalyzeExpense"); err != nil {
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
	if err = addOpAnalyzeExpenseValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAnalyzeExpense(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAnalyzeExpense(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "AnalyzeExpense",
	}
}
