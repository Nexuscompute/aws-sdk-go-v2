// Code generated by smithy-go-codegen DO NOT EDIT.

package datasync

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/datasync/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Provides details about how an DataSync transfer location for an Amazon EFS file
// system is configured.
func (c *Client) DescribeLocationEfs(ctx context.Context, params *DescribeLocationEfsInput, optFns ...func(*Options)) (*DescribeLocationEfsOutput, error) {
	if params == nil {
		params = &DescribeLocationEfsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeLocationEfs", params, optFns, c.addOperationDescribeLocationEfsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeLocationEfsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// DescribeLocationEfsRequest
type DescribeLocationEfsInput struct {

	// The Amazon Resource Name (ARN) of the Amazon EFS file system location that you
	// want information about.
	//
	// This member is required.
	LocationArn *string

	noSmithyDocumentSerde
}

// DescribeLocationEfsResponse
type DescribeLocationEfsOutput struct {

	// The ARN of the access point that DataSync uses to access the Amazon EFS file
	// system.
	AccessPointArn *string

	// The time that the location was created.
	CreationTime *time.Time

	// The subnet and security groups that DataSync uses to access your Amazon EFS
	// file system.
	Ec2Config *types.Ec2Config

	// The Identity and Access Management (IAM) role that DataSync assumes when
	// mounting the Amazon EFS file system.
	FileSystemAccessRoleArn *string

	// Describes whether DataSync uses Transport Layer Security (TLS) encryption when
	// copying data to or from the Amazon EFS file system.
	InTransitEncryption types.EfsInTransitEncryption

	// The ARN of the Amazon EFS file system location.
	LocationArn *string

	// The URL of the Amazon EFS file system location.
	LocationUri *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeLocationEfsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeLocationEfs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeLocationEfs{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeLocationEfs"); err != nil {
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
	if err = addOpDescribeLocationEfsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeLocationEfs(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeLocationEfs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeLocationEfs",
	}
}
