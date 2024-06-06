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
)

// Deletes the resource policy of the S3 Access Grants instance. The resource
// policy is used to manage cross-account access to your S3 Access Grants instance.
// By deleting the resource policy, you delete any cross-account permissions to
// your S3 Access Grants instance.
//
// Permissions You must have the s3:DeleteAccessGrantsInstanceResourcePolicy
// permission to use this operation.
func (c *Client) DeleteAccessGrantsInstanceResourcePolicy(ctx context.Context, params *DeleteAccessGrantsInstanceResourcePolicyInput, optFns ...func(*Options)) (*DeleteAccessGrantsInstanceResourcePolicyOutput, error) {
	if params == nil {
		params = &DeleteAccessGrantsInstanceResourcePolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteAccessGrantsInstanceResourcePolicy", params, optFns, c.addOperationDeleteAccessGrantsInstanceResourcePolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteAccessGrantsInstanceResourcePolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteAccessGrantsInstanceResourcePolicyInput struct {

	// The ID of the Amazon Web Services account that is making this request.
	//
	// This member is required.
	AccountId *string

	noSmithyDocumentSerde
}

func (in *DeleteAccessGrantsInstanceResourcePolicyInput) bindEndpointParams(p *EndpointParameters) {
	p.AccountId = in.AccountId
	p.RequiresAccountId = ptr.Bool(true)
}

type DeleteAccessGrantsInstanceResourcePolicyOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteAccessGrantsInstanceResourcePolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestxml_serializeOpDeleteAccessGrantsInstanceResourcePolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpDeleteAccessGrantsInstanceResourcePolicy{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeleteAccessGrantsInstanceResourcePolicy"); err != nil {
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
	if err = smithyhttp.AddContentChecksumMiddleware(stack); err != nil {
		return err
	}
	if err = addEndpointPrefix_opDeleteAccessGrantsInstanceResourcePolicyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpDeleteAccessGrantsInstanceResourcePolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteAccessGrantsInstanceResourcePolicy(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addDeleteAccessGrantsInstanceResourcePolicyUpdateEndpoint(stack, options); err != nil {
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

type endpointPrefix_opDeleteAccessGrantsInstanceResourcePolicyMiddleware struct {
}

func (*endpointPrefix_opDeleteAccessGrantsInstanceResourcePolicyMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opDeleteAccessGrantsInstanceResourcePolicyMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
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
	input, ok := opaqueInput.(*DeleteAccessGrantsInstanceResourcePolicyInput)
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
func addEndpointPrefix_opDeleteAccessGrantsInstanceResourcePolicyMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opDeleteAccessGrantsInstanceResourcePolicyMiddleware{}, "ResolveEndpointV2", middleware.After)
}

func newServiceMetadataMiddleware_opDeleteAccessGrantsInstanceResourcePolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeleteAccessGrantsInstanceResourcePolicy",
	}
}

func copyDeleteAccessGrantsInstanceResourcePolicyInputForUpdateEndpoint(params interface{}) (interface{}, error) {
	input, ok := params.(*DeleteAccessGrantsInstanceResourcePolicyInput)
	if !ok {
		return nil, fmt.Errorf("expect *DeleteAccessGrantsInstanceResourcePolicyInput type, got %T", params)
	}
	cpy := *input
	return &cpy, nil
}
func (in *DeleteAccessGrantsInstanceResourcePolicyInput) copy() interface{} {
	v := *in
	return &v
}
func backFillDeleteAccessGrantsInstanceResourcePolicyAccountID(input interface{}, v string) error {
	in := input.(*DeleteAccessGrantsInstanceResourcePolicyInput)
	if in.AccountId != nil {
		if !strings.EqualFold(*in.AccountId, v) {
			return fmt.Errorf("error backfilling account id")
		}
		return nil
	}
	in.AccountId = &v
	return nil
}
func addDeleteAccessGrantsInstanceResourcePolicyUpdateEndpoint(stack *middleware.Stack, options Options) error {
	return s3controlcust.UpdateEndpoint(stack, s3controlcust.UpdateEndpointOptions{
		Accessor: s3controlcust.UpdateEndpointParameterAccessor{GetARNInput: nopGetARNAccessor,
			BackfillAccountID: nopBackfillAccountIDAccessor,
			GetOutpostIDInput: nopGetOutpostIDFromInput,
			UpdateARNField:    nopSetARNAccessor,
			CopyInput:         copyDeleteAccessGrantsInstanceResourcePolicyInputForUpdateEndpoint,
		},
		EndpointResolver:        options.EndpointResolver,
		EndpointResolverOptions: options.EndpointOptions,
		UseARNRegion:            options.UseARNRegion,
	})
}
