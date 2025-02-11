// Code generated by smithy-go-codegen DO NOT EDIT.

package datazone

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/datazone/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates a custom asset type.
func (c *Client) CreateAssetType(ctx context.Context, params *CreateAssetTypeInput, optFns ...func(*Options)) (*CreateAssetTypeOutput, error) {
	if params == nil {
		params = &CreateAssetTypeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateAssetType", params, optFns, c.addOperationCreateAssetTypeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateAssetTypeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateAssetTypeInput struct {

	// The unique identifier of the Amazon DataZone domain where the custom asset type
	// is being created.
	//
	// This member is required.
	DomainIdentifier *string

	// The metadata forms that are to be attached to the custom asset type.
	//
	// This member is required.
	FormsInput map[string]types.FormEntryInput

	// The name of the custom asset type.
	//
	// This member is required.
	Name *string

	// The identifier of the Amazon DataZone project that is to own the custom asset
	// type.
	//
	// This member is required.
	OwningProjectIdentifier *string

	// The descripton of the custom asset type.
	Description *string

	noSmithyDocumentSerde
}

type CreateAssetTypeOutput struct {

	// The ID of the Amazon DataZone domain in which the asset type was created.
	//
	// This member is required.
	DomainId *string

	// The metadata forms that are attached to the asset type.
	//
	// This member is required.
	FormsOutput map[string]types.FormEntryOutput

	// The name of the asset type.
	//
	// This member is required.
	Name *string

	// The revision of the custom asset type.
	//
	// This member is required.
	Revision *string

	// The timestamp of when the asset type is to be created.
	CreatedAt *time.Time

	// The Amazon DataZone user who creates this custom asset type.
	CreatedBy *string

	// The description of the custom asset type.
	Description *string

	// The ID of the Amazon DataZone domain where the asset type was originally
	// created.
	OriginDomainId *string

	// The ID of the Amazon DataZone project where the asset type was originally
	// created.
	OriginProjectId *string

	// The ID of the Amazon DataZone project that currently owns this asset type.
	OwningProjectId *string

	// The timestamp of when the custom type was created.
	UpdatedAt *time.Time

	// The Amazon DataZone user that created the custom asset type.
	UpdatedBy *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateAssetTypeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateAssetType{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateAssetType{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateAssetType"); err != nil {
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
	if err = addOpCreateAssetTypeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateAssetType(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateAssetType(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateAssetType",
	}
}
