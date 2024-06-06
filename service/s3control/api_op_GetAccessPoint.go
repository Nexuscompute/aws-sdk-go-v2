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
	"time"
)

// This operation is not supported by directory buckets.
//
// Returns configuration information about the specified access point.
//
// All Amazon S3 on Outposts REST API requests for this action require an
// additional parameter of x-amz-outpost-id to be passed with the request. In
// addition, you must use an S3 on Outposts endpoint hostname prefix instead of
// s3-control . For an example of the request syntax for Amazon S3 on Outposts that
// uses the S3 on Outposts endpoint hostname prefix and the x-amz-outpost-id
// derived by using the access point ARN, see the [Examples]section.
//
// The following actions are related to GetAccessPoint :
//
// [CreateAccessPoint]
//
// [DeleteAccessPoint]
//
// [ListAccessPoints]
//
// [ListAccessPoints]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_ListAccessPoints.html
// [CreateAccessPoint]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_CreateAccessPoint.html
// [DeleteAccessPoint]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteAccessPoint.html
// [Examples]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetAccessPoint.html#API_control_GetAccessPoint_Examples
func (c *Client) GetAccessPoint(ctx context.Context, params *GetAccessPointInput, optFns ...func(*Options)) (*GetAccessPointOutput, error) {
	if params == nil {
		params = &GetAccessPointInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetAccessPoint", params, optFns, c.addOperationGetAccessPointMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetAccessPointOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetAccessPointInput struct {

	// The Amazon Web Services account ID for the account that owns the specified
	// access point.
	//
	// This member is required.
	AccountId *string

	// The name of the access point whose configuration information you want to
	// retrieve.
	//
	// For using this parameter with Amazon S3 on Outposts with the REST API, you must
	// specify the name and the x-amz-outpost-id as well.
	//
	// For using this parameter with S3 on Outposts with the Amazon Web Services SDK
	// and CLI, you must specify the ARN of the access point accessed in the format
	// arn:aws:s3-outposts:::outpost//accesspoint/ . For example, to access the access
	// point reports-ap through Outpost my-outpost owned by account 123456789012 in
	// Region us-west-2 , use the URL encoding of
	// arn:aws:s3-outposts:us-west-2:123456789012:outpost/my-outpost/accesspoint/reports-ap
	// . The value must be URL encoded.
	//
	// This member is required.
	Name *string

	noSmithyDocumentSerde
}

func (in *GetAccessPointInput) bindEndpointParams(p *EndpointParameters) {
	p.AccountId = in.AccountId
	p.AccessPointName = in.Name
	p.RequiresAccountId = ptr.Bool(true)
}

type GetAccessPointOutput struct {

	// The ARN of the access point.
	AccessPointArn *string

	// The name or alias of the access point.
	Alias *string

	// The name of the bucket associated with the specified access point.
	Bucket *string

	// The Amazon Web Services account ID associated with the S3 bucket associated
	// with this access point.
	BucketAccountId *string

	// The date and time when the specified access point was created.
	CreationDate *time.Time

	// The VPC endpoint for the access point.
	Endpoints map[string]string

	// The name of the specified access point.
	Name *string

	// Indicates whether this access point allows access from the public internet. If
	// VpcConfiguration is specified for this access point, then NetworkOrigin is VPC ,
	// and the access point doesn't allow access from the public internet. Otherwise,
	// NetworkOrigin is Internet , and the access point allows access from the public
	// internet, subject to the access point and bucket access policies.
	//
	// This will always be true for an Amazon S3 on Outposts access point
	NetworkOrigin types.NetworkOrigin

	// The PublicAccessBlock configuration that you want to apply to this Amazon S3
	// account. You can enable the configuration options in any combination. For more
	// information about when Amazon S3 considers a bucket or object public, see [The Meaning of "Public"]in
	// the Amazon S3 User Guide.
	//
	// This data type is not supported for Amazon S3 on Outposts.
	//
	// [The Meaning of "Public"]: https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html#access-control-block-public-access-policy-status
	PublicAccessBlockConfiguration *types.PublicAccessBlockConfiguration

	// Contains the virtual private cloud (VPC) configuration for the specified access
	// point.
	//
	// This element is empty if this access point is an Amazon S3 on Outposts access
	// point that is used by other Amazon Web Services.
	VpcConfiguration *types.VpcConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetAccessPointMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestxml_serializeOpGetAccessPoint{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpGetAccessPoint{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetAccessPoint"); err != nil {
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
	if err = addEndpointPrefix_opGetAccessPointMiddleware(stack); err != nil {
		return err
	}
	if err = addOpGetAccessPointValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetAccessPoint(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addGetAccessPointUpdateEndpoint(stack, options); err != nil {
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

func (m *GetAccessPointInput) GetARNMember() (*string, bool) {
	if m.Name == nil {
		return nil, false
	}
	return m.Name, true
}

func (m *GetAccessPointInput) SetARNMember(v string) error {
	m.Name = &v
	return nil
}

type endpointPrefix_opGetAccessPointMiddleware struct {
}

func (*endpointPrefix_opGetAccessPointMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opGetAccessPointMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
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
	input, ok := opaqueInput.(*GetAccessPointInput)
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
func addEndpointPrefix_opGetAccessPointMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opGetAccessPointMiddleware{}, "ResolveEndpointV2", middleware.After)
}

func newServiceMetadataMiddleware_opGetAccessPoint(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetAccessPoint",
	}
}

func copyGetAccessPointInputForUpdateEndpoint(params interface{}) (interface{}, error) {
	input, ok := params.(*GetAccessPointInput)
	if !ok {
		return nil, fmt.Errorf("expect *GetAccessPointInput type, got %T", params)
	}
	cpy := *input
	return &cpy, nil
}
func (in *GetAccessPointInput) copy() interface{} {
	v := *in
	return &v
}
func getGetAccessPointARNMember(input interface{}) (*string, bool) {
	in := input.(*GetAccessPointInput)
	if in.Name == nil {
		return nil, false
	}
	return in.Name, true
}
func setGetAccessPointARNMember(input interface{}, v string) error {
	in := input.(*GetAccessPointInput)
	in.Name = &v
	return nil
}
func backFillGetAccessPointAccountID(input interface{}, v string) error {
	in := input.(*GetAccessPointInput)
	if in.AccountId != nil {
		if !strings.EqualFold(*in.AccountId, v) {
			return fmt.Errorf("error backfilling account id")
		}
		return nil
	}
	in.AccountId = &v
	return nil
}
func addGetAccessPointUpdateEndpoint(stack *middleware.Stack, options Options) error {
	return s3controlcust.UpdateEndpoint(stack, s3controlcust.UpdateEndpointOptions{
		Accessor: s3controlcust.UpdateEndpointParameterAccessor{GetARNInput: getGetAccessPointARNMember,
			BackfillAccountID: backFillGetAccessPointAccountID,
			GetOutpostIDInput: nopGetOutpostIDFromInput,
			UpdateARNField:    setGetAccessPointARNMember,
			CopyInput:         copyGetAccessPointInputForUpdateEndpoint,
		},
		EndpointResolver:        options.EndpointResolver,
		EndpointResolverOptions: options.EndpointOptions,
		UseARNRegion:            options.UseARNRegion,
	})
}
