// Code generated by smithy-go-codegen DO NOT EDIT.

package backup

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the Amazon Web Services resource types supported by Backup.
func (c *Client) GetSupportedResourceTypes(ctx context.Context, params *GetSupportedResourceTypesInput, optFns ...func(*Options)) (*GetSupportedResourceTypesOutput, error) {
	if params == nil {
		params = &GetSupportedResourceTypesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetSupportedResourceTypes", params, optFns, c.addOperationGetSupportedResourceTypesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetSupportedResourceTypesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetSupportedResourceTypesInput struct {
	noSmithyDocumentSerde
}

type GetSupportedResourceTypesOutput struct {

	// Contains a string with the supported Amazon Web Services resource types:
	//
	//   - Aurora for Amazon Aurora
	//
	//   - DynamoDB for Amazon DynamoDB
	//
	//   - EBS for Amazon Elastic Block Store
	//
	//   - EC2 for Amazon Elastic Compute Cloud
	//
	//   - EFS for Amazon Elastic File System
	//
	//   - FSX for Amazon FSx
	//
	//   - RDS for Amazon Relational Database Service
	//
	//   - Storage Gateway for Storage Gateway
	//
	//   - DocDB for Amazon DocumentDB (with MongoDB compatibility)
	//
	//   - Neptune for Amazon Neptune
	ResourceTypes []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetSupportedResourceTypesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetSupportedResourceTypes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetSupportedResourceTypes{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetSupportedResourceTypes"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetSupportedResourceTypes(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetSupportedResourceTypes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetSupportedResourceTypes",
	}
}
