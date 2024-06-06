// Code generated by smithy-go-codegen DO NOT EDIT.

package qconnect

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/qconnect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates information about the content.
func (c *Client) UpdateContent(ctx context.Context, params *UpdateContentInput, optFns ...func(*Options)) (*UpdateContentOutput, error) {
	if params == nil {
		params = &UpdateContentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateContent", params, optFns, c.addOperationUpdateContentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateContentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateContentInput struct {

	// The identifier of the content. Can be either the ID or the ARN. URLs cannot
	// contain the ARN.
	//
	// This member is required.
	ContentId *string

	// The identifier of the knowledge base. This should not be a QUICK_RESPONSES type
	// knowledge base. Can be either the ID or the ARN
	//
	// This member is required.
	KnowledgeBaseId *string

	// A key/value map to store attributes without affecting tagging or
	// recommendations. For example, when synchronizing data between an external system
	// and Amazon Q in Connect, you can store an external version identifier as
	// metadata to utilize for determining drift.
	Metadata map[string]string

	// The URI for the article. If the knowledge base has a templateUri, setting this
	// argument overrides it for this piece of content. To remove an existing
	// overrideLinkOurUri , exclude this argument and set removeOverrideLinkOutUri to
	// true.
	OverrideLinkOutUri *string

	// Unset the existing overrideLinkOutUri if it exists.
	RemoveOverrideLinkOutUri *bool

	// The revisionId of the content resource to update, taken from an earlier call to
	// GetContent , GetContentSummary , SearchContent , or ListContents . If included,
	// this argument acts as an optimistic lock to ensure content was not modified
	// since it was last read. If it has been modified, this API throws a
	// PreconditionFailedException .
	RevisionId *string

	// The title of the content.
	Title *string

	// A pointer to the uploaded asset. This value is returned by [StartContentUpload].
	//
	// [StartContentUpload]: https://docs.aws.amazon.com/amazon-q-connect/latest/APIReference/API_StartContentUpload.html
	UploadId *string

	noSmithyDocumentSerde
}

type UpdateContentOutput struct {

	// The content.
	Content *types.ContentData

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateContentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateContent{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateContent{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateContent"); err != nil {
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
	if err = addOpUpdateContentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateContent(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateContent(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateContent",
	}
}
