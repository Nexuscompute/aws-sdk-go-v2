// Code generated by smithy-go-codegen DO NOT EDIT.

package s3control

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	s3controlcust "github.com/aws/aws-sdk-go-v2/service/s3control/internal/customizations"
	"github.com/aws/aws-sdk-go-v2/service/s3control/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/ptr"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"strings"
)

// This operation is not supported by directory buckets.
//
// Retrieves the PublicAccessBlock configuration for an Amazon Web Services
// account. For more information, see [Using Amazon S3 block public access].
//
// Related actions include:
//
// [DeletePublicAccessBlock]
//
// [PutPublicAccessBlock]
//
// [PutPublicAccessBlock]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_PutPublicAccessBlock.html
// [DeletePublicAccessBlock]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeletePublicAccessBlock.html
// [Using Amazon S3 block public access]: https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html
func (c *Client) GetPublicAccessBlock(ctx context.Context, params *GetPublicAccessBlockInput, optFns ...func(*Options)) (*GetPublicAccessBlockOutput, error) {
	if params == nil {
		params = &GetPublicAccessBlockInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetPublicAccessBlock", params, optFns, c.addOperationGetPublicAccessBlockMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetPublicAccessBlockOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetPublicAccessBlockInput struct {

	// The account ID for the Amazon Web Services account whose PublicAccessBlock
	// configuration you want to retrieve.
	//
	// This member is required.
	AccountId *string

	noSmithyDocumentSerde
}

func (in *GetPublicAccessBlockInput) bindEndpointParams(p *EndpointParameters) {
	p.AccountId = in.AccountId
	p.RequiresAccountId = ptr.Bool(true)
}

type GetPublicAccessBlockOutput struct {

	// The PublicAccessBlock configuration currently in effect for this Amazon Web
	// Services account.
	PublicAccessBlockConfiguration *types.PublicAccessBlockConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetPublicAccessBlockMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestxml_serializeOpGetPublicAccessBlock{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpGetPublicAccessBlock{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetPublicAccessBlock"); err != nil {
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
	if err = s3controlcust.AddUpdateOutpostARN(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addEndpointPrefix_opGetPublicAccessBlockMiddleware(stack); err != nil {
		return err
	}
	if err = addOpGetPublicAccessBlockValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetPublicAccessBlock(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addGetPublicAccessBlockUpdateEndpoint(stack, options); err != nil {
		return err
	}
	if err = addStashOperationInput(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = v4.AddContentSHA256HeaderMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	if err = s3controlcust.AddDisableHostPrefixMiddleware(stack); err != nil {
		return err
	}
	return nil
}

type endpointPrefix_opGetPublicAccessBlockMiddleware struct {
}

func (*endpointPrefix_opGetPublicAccessBlockMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opGetPublicAccessBlockMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleFinalize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	opaqueInput := getOperationInput(ctx)
	input, ok := opaqueInput.(*GetPublicAccessBlockInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input type %T", opaqueInput)
	}

	var prefix strings.Builder
	if input.AccountId == nil {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("AccountId forms part of the endpoint host and so may not be nil")}
	} else if !smithyhttp.ValidHostLabel(*input.AccountId) {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("AccountId forms part of the endpoint host and so must match \"[a-zA-Z0-9-]{1,63}\", but was \"%s\"", *input.AccountId)}
	} else {
		prefix.WriteString(*input.AccountId)
	}
	prefix.WriteString(".")
	req.URL.Host = prefix.String() + req.URL.Host

	return next.HandleFinalize(ctx, in)
}
func addEndpointPrefix_opGetPublicAccessBlockMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opGetPublicAccessBlockMiddleware{}, "ResolveEndpointV2", middleware.After)
}

func newServiceMetadataMiddleware_opGetPublicAccessBlock(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetPublicAccessBlock",
	}
}

func copyGetPublicAccessBlockInputForUpdateEndpoint(params interface{}) (interface{}, error) {
	input, ok := params.(*GetPublicAccessBlockInput)
	if !ok {
		return nil, fmt.Errorf("expect *GetPublicAccessBlockInput type, got %T", params)
	}
	cpy := *input
	return &cpy, nil
}
func (in *GetPublicAccessBlockInput) copy() interface{} {
	v := *in
	return &v
}
func backFillGetPublicAccessBlockAccountID(input interface{}, v string) error {
	in := input.(*GetPublicAccessBlockInput)
	if in.AccountId != nil {
		if !strings.EqualFold(*in.AccountId, v) {
			return fmt.Errorf("error backfilling account id")
		}
		return nil
	}
	in.AccountId = &v
	return nil
}
func addGetPublicAccessBlockUpdateEndpoint(stack *middleware.Stack, options Options) error {
	return s3controlcust.UpdateEndpoint(stack, s3controlcust.UpdateEndpointOptions{
		Accessor: s3controlcust.UpdateEndpointParameterAccessor{GetARNInput: nopGetARNAccessor,
			BackfillAccountID: nopBackfillAccountIDAccessor,
			GetOutpostIDInput: nopGetOutpostIDFromInput,
			UpdateARNField:    nopSetARNAccessor,
			CopyInput:         copyGetPublicAccessBlockInputForUpdateEndpoint,
		},
		EndpointResolver:        options.EndpointResolver,
		EndpointResolverOptions: options.EndpointOptions,
		UseARNRegion:            options.UseARNRegion,
	})
}
