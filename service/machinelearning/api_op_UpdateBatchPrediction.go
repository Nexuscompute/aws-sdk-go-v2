// Code generated by smithy-go-codegen DO NOT EDIT.

package machinelearning

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the BatchPredictionName of a BatchPrediction .
//
// You can use the GetBatchPrediction operation to view the contents of the
// updated data element.
func (c *Client) UpdateBatchPrediction(ctx context.Context, params *UpdateBatchPredictionInput, optFns ...func(*Options)) (*UpdateBatchPredictionOutput, error) {
	if params == nil {
		params = &UpdateBatchPredictionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateBatchPrediction", params, optFns, c.addOperationUpdateBatchPredictionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateBatchPredictionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateBatchPredictionInput struct {

	// The ID assigned to the BatchPrediction during creation.
	//
	// This member is required.
	BatchPredictionId *string

	// A new user-supplied name or description of the BatchPrediction .
	//
	// This member is required.
	BatchPredictionName *string

	noSmithyDocumentSerde
}

// Represents the output of an UpdateBatchPrediction operation.
//
// You can see the updated content by using the GetBatchPrediction operation.
type UpdateBatchPredictionOutput struct {

	// The ID assigned to the BatchPrediction during creation. This value should be
	// identical to the value of the BatchPredictionId in the request.
	BatchPredictionId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateBatchPredictionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateBatchPrediction{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateBatchPrediction{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateBatchPrediction"); err != nil {
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
	if err = addOpUpdateBatchPredictionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateBatchPrediction(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateBatchPrediction(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateBatchPrediction",
	}
}
