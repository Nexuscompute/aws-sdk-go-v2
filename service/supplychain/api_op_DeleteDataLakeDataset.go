// Code generated by smithy-go-codegen DO NOT EDIT.

package supplychain

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Enables you to programmatically delete an Amazon Web Services Supply Chain data
// lake dataset. Developers can delete the existing datasets for a given instance
// ID, namespace, and instance name.
func (c *Client) DeleteDataLakeDataset(ctx context.Context, params *DeleteDataLakeDatasetInput, optFns ...func(*Options)) (*DeleteDataLakeDatasetOutput, error) {
	if params == nil {
		params = &DeleteDataLakeDatasetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteDataLakeDataset", params, optFns, c.addOperationDeleteDataLakeDatasetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteDataLakeDatasetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request parameters of DeleteDataLakeDataset.
type DeleteDataLakeDatasetInput struct {

	// The AWS Supply Chain instance identifier.
	//
	// This member is required.
	InstanceId *string

	// The name of the dataset. For asc name space, the name must be one of the
	// supported data entities under [https://docs.aws.amazon.com/aws-supply-chain/latest/userguide/data-model-asc.html].
	//
	// [https://docs.aws.amazon.com/aws-supply-chain/latest/userguide/data-model-asc.html]: https://docs.aws.amazon.com/aws-supply-chain/latest/userguide/data-model-asc.html
	//
	// This member is required.
	Name *string

	// The name space of the dataset. The available values are:
	//
	//   - asc - For information on the Amazon Web Services Supply Chain supported
	//   datasets see [https://docs.aws.amazon.com/aws-supply-chain/latest/userguide/data-model-asc.html].
	//
	//   - default - For datasets with custom user-defined schemas.
	//
	// [https://docs.aws.amazon.com/aws-supply-chain/latest/userguide/data-model-asc.html]: https://docs.aws.amazon.com/aws-supply-chain/latest/userguide/data-model-asc.html
	//
	// This member is required.
	Namespace *string

	noSmithyDocumentSerde
}

// The response parameters of DeleteDataLakeDataset.
type DeleteDataLakeDatasetOutput struct {

	// The AWS Supply Chain instance identifier.
	//
	// This member is required.
	InstanceId *string

	// The name of deleted dataset.
	//
	// This member is required.
	Name *string

	// The name space of deleted dataset.
	//
	// This member is required.
	Namespace *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteDataLakeDatasetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteDataLakeDataset{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteDataLakeDataset{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeleteDataLakeDataset"); err != nil {
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
	if err = addOpDeleteDataLakeDatasetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteDataLakeDataset(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteDataLakeDataset(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeleteDataLakeDataset",
	}
}
