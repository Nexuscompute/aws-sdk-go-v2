// Code generated by smithy-go-codegen DO NOT EDIT.

package tnb

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/tnb/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Uploads the contents of a network package.
//
// A network package is a .zip file in CSAR (Cloud Service Archive) format defines
// the function packages you want to deploy and the Amazon Web Services
// infrastructure you want to deploy them on.
func (c *Client) PutSolNetworkPackageContent(ctx context.Context, params *PutSolNetworkPackageContentInput, optFns ...func(*Options)) (*PutSolNetworkPackageContentOutput, error) {
	if params == nil {
		params = &PutSolNetworkPackageContentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutSolNetworkPackageContent", params, optFns, c.addOperationPutSolNetworkPackageContentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutSolNetworkPackageContentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutSolNetworkPackageContentInput struct {

	// Network package file.
	//
	// This member is required.
	File []byte

	// Network service descriptor info ID.
	//
	// This member is required.
	NsdInfoId *string

	// Network package content type.
	ContentType types.PackageContentType

	noSmithyDocumentSerde
}

type PutSolNetworkPackageContentOutput struct {

	// Network package ARN.
	//
	// This member is required.
	Arn *string

	// Network package ID.
	//
	// This member is required.
	Id *string

	// Network package metadata.
	//
	// This member is required.
	Metadata *types.PutSolNetworkPackageContentMetadata

	// Network service descriptor ID.
	//
	// This member is required.
	NsdId *string

	// Network service descriptor name.
	//
	// This member is required.
	NsdName *string

	// Network service descriptor version.
	//
	// This member is required.
	NsdVersion *string

	// Function package IDs.
	//
	// This member is required.
	VnfPkgIds []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutSolNetworkPackageContentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutSolNetworkPackageContent{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutSolNetworkPackageContent{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PutSolNetworkPackageContent"); err != nil {
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
	if err = addOpPutSolNetworkPackageContentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutSolNetworkPackageContent(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutSolNetworkPackageContent(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PutSolNetworkPackageContent",
	}
}
