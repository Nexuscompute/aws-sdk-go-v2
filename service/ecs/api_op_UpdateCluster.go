// Code generated by smithy-go-codegen DO NOT EDIT.

package ecs

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the cluster.
func (c *Client) UpdateCluster(ctx context.Context, params *UpdateClusterInput, optFns ...func(*Options)) (*UpdateClusterOutput, error) {
	if params == nil {
		params = &UpdateClusterInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateCluster", params, optFns, c.addOperationUpdateClusterMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateClusterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateClusterInput struct {

	// The name of the cluster to modify the settings for.
	//
	// This member is required.
	Cluster *string

	// The execute command configuration for the cluster.
	Configuration *types.ClusterConfiguration

	// Use this parameter to set a default Service Connect namespace. After you set a
	// default Service Connect namespace, any new services with Service Connect turned
	// on that are created in the cluster are added as client services in the
	// namespace. This setting only applies to new services that set the enabled
	// parameter to true in the ServiceConnectConfiguration . You can set the namespace
	// of each service individually in the ServiceConnectConfiguration to override
	// this default parameter.
	//
	// Tasks that run in a namespace can use short names to connect to services in the
	// namespace. Tasks can connect to services across all of the clusters in the
	// namespace. Tasks connect through a managed proxy container that collects logs
	// and metrics for increased visibility. Only the tasks that Amazon ECS services
	// create are supported with Service Connect. For more information, see [Service Connect]in the
	// Amazon Elastic Container Service Developer Guide.
	//
	// [Service Connect]: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/service-connect.html
	ServiceConnectDefaults *types.ClusterServiceConnectDefaultsRequest

	// The cluster settings for your cluster.
	Settings []types.ClusterSetting

	noSmithyDocumentSerde
}

type UpdateClusterOutput struct {

	// Details about the cluster.
	Cluster *types.Cluster

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateClusterMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateCluster{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateCluster{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateCluster"); err != nil {
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
	if err = addOpUpdateClusterValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateCluster(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateCluster(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateCluster",
	}
}
