// Code generated by smithy-go-codegen DO NOT EDIT.

package kafka

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/kafka/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns a description of this MSK VPC connection.
func (c *Client) DescribeVpcConnection(ctx context.Context, params *DescribeVpcConnectionInput, optFns ...func(*Options)) (*DescribeVpcConnectionOutput, error) {
	if params == nil {
		params = &DescribeVpcConnectionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeVpcConnection", params, optFns, c.addOperationDescribeVpcConnectionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeVpcConnectionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeVpcConnectionInput struct {

	// The Amazon Resource Name (ARN) that uniquely identifies a MSK VPC connection.
	//
	// This member is required.
	Arn *string

	noSmithyDocumentSerde
}

type DescribeVpcConnectionOutput struct {

	// The authentication type of VPC connection.
	Authentication *string

	// The creation time of the VPC connection.
	CreationTime *time.Time

	// The list of security groups for the VPC connection.
	SecurityGroups []string

	// The state of VPC connection.
	State types.VpcConnectionState

	// The list of subnets for the VPC connection.
	Subnets []string

	// A map of tags for the VPC connection.
	Tags map[string]string

	// The Amazon Resource Name (ARN) that uniquely identifies an MSK cluster.
	TargetClusterArn *string

	// The Amazon Resource Name (ARN) that uniquely identifies a MSK VPC connection.
	VpcConnectionArn *string

	// The VPC Id for the VPC connection.
	VpcId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeVpcConnectionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeVpcConnection{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeVpcConnection{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeVpcConnection"); err != nil {
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
	if err = addOpDescribeVpcConnectionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeVpcConnection(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeVpcConnection(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeVpcConnection",
	}
}
