// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Changes the size of the cluster. You can change the cluster's type, or change
// the number or type of nodes. The default behavior is to use the elastic resize
// method. With an elastic resize, your cluster is available for read and write
// operations more quickly than with the classic resize method.
//
// Elastic resize operations have the following restrictions:
//
//   - You can only resize clusters of the following types:
//
//   - dc1.large (if your cluster is in a VPC)
//
//   - dc1.8xlarge (if your cluster is in a VPC)
//
//   - dc2.large
//
//   - dc2.8xlarge
//
//   - ds2.xlarge
//
//   - ds2.8xlarge
//
//   - ra3.xlplus
//
//   - ra3.4xlarge
//
//   - ra3.16xlarge
//
//   - The type of nodes that you add must match the node type for the cluster.
func (c *Client) ResizeCluster(ctx context.Context, params *ResizeClusterInput, optFns ...func(*Options)) (*ResizeClusterOutput, error) {
	if params == nil {
		params = &ResizeClusterInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ResizeCluster", params, optFns, c.addOperationResizeClusterMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ResizeClusterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Describes a resize cluster operation. For example, a scheduled action to run
// the ResizeCluster API operation.
type ResizeClusterInput struct {

	// The unique identifier for the cluster to resize.
	//
	// This member is required.
	ClusterIdentifier *string

	// A boolean value indicating whether the resize operation is using the classic
	// resize process. If you don't provide this parameter or set the value to false ,
	// the resize type is elastic.
	Classic *bool

	// The new cluster type for the specified cluster.
	ClusterType *string

	// The new node type for the nodes you are adding. If not specified, the cluster's
	// current node type is used.
	NodeType *string

	// The new number of nodes for the cluster. If not specified, the cluster's
	// current number of nodes is used.
	NumberOfNodes *int32

	// The identifier of the reserved node.
	ReservedNodeId *string

	// The identifier of the target reserved node offering.
	TargetReservedNodeOfferingId *string

	noSmithyDocumentSerde
}

type ResizeClusterOutput struct {

	// Describes a cluster.
	Cluster *types.Cluster

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationResizeClusterMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpResizeCluster{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpResizeCluster{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ResizeCluster"); err != nil {
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
	if err = addOpResizeClusterValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opResizeCluster(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opResizeCluster(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ResizeCluster",
	}
}
