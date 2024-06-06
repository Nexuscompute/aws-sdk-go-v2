// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticache

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/elasticache/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Apply the service update. For more information on service updates and applying
// them, see [Applying Service Updates].
//
// [Applying Service Updates]: https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/applying-updates.html
func (c *Client) BatchApplyUpdateAction(ctx context.Context, params *BatchApplyUpdateActionInput, optFns ...func(*Options)) (*BatchApplyUpdateActionOutput, error) {
	if params == nil {
		params = &BatchApplyUpdateActionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchApplyUpdateAction", params, optFns, c.addOperationBatchApplyUpdateActionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchApplyUpdateActionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchApplyUpdateActionInput struct {

	// The unique ID of the service update
	//
	// This member is required.
	ServiceUpdateName *string

	// The cache cluster IDs
	CacheClusterIds []string

	// The replication group IDs
	ReplicationGroupIds []string

	noSmithyDocumentSerde
}

type BatchApplyUpdateActionOutput struct {

	// Update actions that have been processed successfully
	ProcessedUpdateActions []types.ProcessedUpdateAction

	// Update actions that haven't been processed successfully
	UnprocessedUpdateActions []types.UnprocessedUpdateAction

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationBatchApplyUpdateActionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpBatchApplyUpdateAction{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpBatchApplyUpdateAction{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "BatchApplyUpdateAction"); err != nil {
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
	if err = addOpBatchApplyUpdateActionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchApplyUpdateAction(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opBatchApplyUpdateAction(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "BatchApplyUpdateAction",
	}
}
