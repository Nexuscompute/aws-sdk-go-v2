// Code generated by smithy-go-codegen DO NOT EDIT.

package s3control

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	s3controlcust "github.com/aws/aws-sdk-go-v2/service/s3control/internal/customizations"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/ptr"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"strings"
	"time"
)

// Retrieves the S3 Access Grants instance for a Region in your account.
//
// Permissions You must have the s3:GetAccessGrantsInstance permission to use this
// operation.
//
// GetAccessGrantsInstance is not supported for cross-account access. You can only
// call the API from the account that owns the S3 Access Grants instance.
func (c *Client) GetAccessGrantsInstance(ctx context.Context, params *GetAccessGrantsInstanceInput, optFns ...func(*Options)) (*GetAccessGrantsInstanceOutput, error) {
	if params == nil {
		params = &GetAccessGrantsInstanceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetAccessGrantsInstance", params, optFns, c.addOperationGetAccessGrantsInstanceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetAccessGrantsInstanceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetAccessGrantsInstanceInput struct {

	// The Amazon Web Services account ID of the S3 Access Grants instance.
	//
	// This member is required.
	AccountId *string

	noSmithyDocumentSerde
}

func (in *GetAccessGrantsInstanceInput) bindEndpointParams(p *EndpointParameters) {

	p.AccountId = in.AccountId
	p.RequiresAccountId = ptr.Bool(true)
}

type GetAccessGrantsInstanceOutput struct {

	// The Amazon Resource Name (ARN) of the S3 Access Grants instance.
	AccessGrantsInstanceArn *string

	// The ID of the S3 Access Grants instance. The ID is default . You can have one S3
	// Access Grants instance per Region per account.
	AccessGrantsInstanceId *string

	// The date and time when you created the S3 Access Grants instance.
	CreatedAt *time.Time

	// If you associated your S3 Access Grants instance with an Amazon Web Services
	// IAM Identity Center instance, this field returns the Amazon Resource Name (ARN)
	// of the IAM Identity Center instance application; a subresource of the original
	// Identity Center instance. S3 Access Grants creates this Identity Center
	// application for the specific S3 Access Grants instance.
	IdentityCenterApplicationArn *string

	// If you associated your S3 Access Grants instance with an Amazon Web Services
	// IAM Identity Center instance, this field returns the Amazon Resource Name (ARN)
	// of the IAM Identity Center instance application; a subresource of the original
	// Identity Center instance. S3 Access Grants creates this Identity Center
	// application for the specific S3 Access Grants instance.
	//
	// Deprecated: IdentityCenterArn has been deprecated. Use
	// IdentityCenterInstanceArn or IdentityCenterApplicationArn.
	IdentityCenterArn *string

	// The Amazon Resource Name (ARN) of the Amazon Web Services IAM Identity Center
	// instance that you are associating with your S3 Access Grants instance. An IAM
	// Identity Center instance is your corporate identity directory that you added to
	// the IAM Identity Center. You can use the [ListInstances]API operation to retrieve a list of
	// your Identity Center instances and their ARNs.
	//
	// [ListInstances]: https://docs.aws.amazon.com/singlesignon/latest/APIReference/API_ListInstances.html
	IdentityCenterInstanceArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetAccessGrantsInstanceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestxml_serializeOpGetAccessGrantsInstance{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpGetAccessGrantsInstance{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetAccessGrantsInstance"); err != nil {
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
	if err = s3controlcust.AddUpdateOutpostARN(stack); err != nil {
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
	if err = smithyhttp.AddContentChecksumMiddleware(stack); err != nil {
		return err
	}
	if err = addEndpointPrefix_opGetAccessGrantsInstanceMiddleware(stack); err != nil {
		return err
	}
	if err = addOpGetAccessGrantsInstanceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetAccessGrantsInstance(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addGetAccessGrantsInstanceUpdateEndpoint(stack, options); err != nil {
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

type endpointPrefix_opGetAccessGrantsInstanceMiddleware struct {
}

func (*endpointPrefix_opGetAccessGrantsInstanceMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opGetAccessGrantsInstanceMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
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
	input, ok := opaqueInput.(*GetAccessGrantsInstanceInput)
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
func addEndpointPrefix_opGetAccessGrantsInstanceMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opGetAccessGrantsInstanceMiddleware{}, "ResolveEndpointV2", middleware.After)
}

func newServiceMetadataMiddleware_opGetAccessGrantsInstance(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetAccessGrantsInstance",
	}
}

func copyGetAccessGrantsInstanceInputForUpdateEndpoint(params interface{}) (interface{}, error) {
	input, ok := params.(*GetAccessGrantsInstanceInput)
	if !ok {
		return nil, fmt.Errorf("expect *GetAccessGrantsInstanceInput type, got %T", params)
	}
	cpy := *input
	return &cpy, nil
}
func (in *GetAccessGrantsInstanceInput) copy() interface{} {
	v := *in
	return &v
}
func backFillGetAccessGrantsInstanceAccountID(input interface{}, v string) error {
	in := input.(*GetAccessGrantsInstanceInput)
	if in.AccountId != nil {
		if !strings.EqualFold(*in.AccountId, v) {
			return fmt.Errorf("error backfilling account id")
		}
		return nil
	}
	in.AccountId = &v
	return nil
}
func addGetAccessGrantsInstanceUpdateEndpoint(stack *middleware.Stack, options Options) error {
	return s3controlcust.UpdateEndpoint(stack, s3controlcust.UpdateEndpointOptions{
		Accessor: s3controlcust.UpdateEndpointParameterAccessor{GetARNInput: nopGetARNAccessor,
			BackfillAccountID: nopBackfillAccountIDAccessor,
			GetOutpostIDInput: nopGetOutpostIDFromInput,
			UpdateARNField:    nopSetARNAccessor,
			CopyInput:         copyGetAccessGrantsInstanceInputForUpdateEndpoint,
		},
		EndpointResolver:        options.EndpointResolver,
		EndpointResolverOptions: options.EndpointOptions,
		UseARNRegion:            options.UseARNRegion,
	})
}
