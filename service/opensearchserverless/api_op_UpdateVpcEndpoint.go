// Code generated by smithy-go-codegen DO NOT EDIT.

package opensearchserverless

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/opensearchserverless/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates an OpenSearch Serverless-managed interface endpoint. For more
// information, see [Access Amazon OpenSearch Serverless using an interface endpoint].
//
// [Access Amazon OpenSearch Serverless using an interface endpoint]: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/serverless-vpc.html
func (c *Client) UpdateVpcEndpoint(ctx context.Context, params *UpdateVpcEndpointInput, optFns ...func(*Options)) (*UpdateVpcEndpointOutput, error) {
	if params == nil {
		params = &UpdateVpcEndpointInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateVpcEndpoint", params, optFns, c.addOperationUpdateVpcEndpointMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateVpcEndpointOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateVpcEndpointInput struct {

	// The unique identifier of the interface endpoint to update.
	//
	// This member is required.
	Id *string

	// The unique identifiers of the security groups to add to the endpoint. Security
	// groups define the ports, protocols, and sources for inbound traffic that you are
	// authorizing into your endpoint.
	AddSecurityGroupIds []string

	// The ID of one or more subnets to add to the endpoint.
	AddSubnetIds []string

	// Unique, case-sensitive identifier to ensure idempotency of the request.
	ClientToken *string

	// The unique identifiers of the security groups to remove from the endpoint.
	RemoveSecurityGroupIds []string

	// The unique identifiers of the subnets to remove from the endpoint.
	RemoveSubnetIds []string

	noSmithyDocumentSerde
}

type UpdateVpcEndpointOutput struct {

	// Details about the updated VPC endpoint.
	UpdateVpcEndpointDetail *types.UpdateVpcEndpointDetail

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateVpcEndpointMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpUpdateVpcEndpoint{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpUpdateVpcEndpoint{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateVpcEndpoint"); err != nil {
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
	if err = addIdempotencyToken_opUpdateVpcEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpUpdateVpcEndpointValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateVpcEndpoint(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpUpdateVpcEndpoint struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpUpdateVpcEndpoint) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpUpdateVpcEndpoint) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*UpdateVpcEndpointInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *UpdateVpcEndpointInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opUpdateVpcEndpointMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpUpdateVpcEndpoint{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opUpdateVpcEndpoint(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateVpcEndpoint",
	}
}
