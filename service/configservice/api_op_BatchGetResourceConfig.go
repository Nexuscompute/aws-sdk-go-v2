// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the BaseConfigurationItem for one or more requested resources. The
// operation also returns a list of resources that are not processed in the current
// request. If there are no unprocessed resources, the operation returns an empty
// unprocessedResourceKeys list.
//
//   - The API does not return results for deleted resources.
//
//   - The API does not return any tags for the requested resources. This
//     information is filtered out of the supplementaryConfiguration section of the API
//     response.
func (c *Client) BatchGetResourceConfig(ctx context.Context, params *BatchGetResourceConfigInput, optFns ...func(*Options)) (*BatchGetResourceConfigOutput, error) {
	if params == nil {
		params = &BatchGetResourceConfigInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchGetResourceConfig", params, optFns, c.addOperationBatchGetResourceConfigMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchGetResourceConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchGetResourceConfigInput struct {

	// A list of resource keys to be processed with the current request. Each element
	// in the list consists of the resource type and resource ID.
	//
	// This member is required.
	ResourceKeys []types.ResourceKey

	noSmithyDocumentSerde
}

type BatchGetResourceConfigOutput struct {

	// A list that contains the current configuration of one or more resources.
	BaseConfigurationItems []types.BaseConfigurationItem

	// A list of resource keys that were not processed with the current response. The
	// unprocessesResourceKeys value is in the same form as ResourceKeys, so the value
	// can be directly provided to a subsequent BatchGetResourceConfig operation.
	//
	// If there are no unprocessed resource keys, the response contains an empty
	// unprocessedResourceKeys list.
	UnprocessedResourceKeys []types.ResourceKey

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationBatchGetResourceConfigMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpBatchGetResourceConfig{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpBatchGetResourceConfig{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "BatchGetResourceConfig"); err != nil {
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
	if err = addOpBatchGetResourceConfigValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchGetResourceConfig(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opBatchGetResourceConfig(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "BatchGetResourceConfig",
	}
}
