// Code generated by smithy-go-codegen DO NOT EDIT.

package mediapackagevod

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/mediapackagevod/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates a specific packaging group. You can't change the id attribute or any
// other system-generated attributes.
func (c *Client) UpdatePackagingGroup(ctx context.Context, params *UpdatePackagingGroupInput, optFns ...func(*Options)) (*UpdatePackagingGroupOutput, error) {
	if params == nil {
		params = &UpdatePackagingGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdatePackagingGroup", params, optFns, c.addOperationUpdatePackagingGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdatePackagingGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A MediaPackage VOD PackagingGroup resource configuration.
type UpdatePackagingGroupInput struct {

	// The ID of a MediaPackage VOD PackagingGroup resource.
	//
	// This member is required.
	Id *string

	// CDN Authorization credentials
	Authorization *types.Authorization

	noSmithyDocumentSerde
}

type UpdatePackagingGroupOutput struct {

	// The approximate asset count of the PackagingGroup.
	ApproximateAssetCount *int32

	// The ARN of the PackagingGroup.
	Arn *string

	// CDN Authorization credentials
	Authorization *types.Authorization

	// The time the PackagingGroup was created.
	CreatedAt *string

	// The fully qualified domain name for Assets in the PackagingGroup.
	DomainName *string

	// Configure egress access logging.
	EgressAccessLogs *types.EgressAccessLogs

	// The ID of the PackagingGroup.
	Id *string

	// A collection of tags associated with a resource
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdatePackagingGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdatePackagingGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdatePackagingGroup{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdatePackagingGroup"); err != nil {
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
	if err = addOpUpdatePackagingGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdatePackagingGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdatePackagingGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdatePackagingGroup",
	}
}
