// Code generated by smithy-go-codegen DO NOT EDIT.

package deadline

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/deadline/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Assigns a farm membership level to a member.
func (c *Client) AssociateMemberToFarm(ctx context.Context, params *AssociateMemberToFarmInput, optFns ...func(*Options)) (*AssociateMemberToFarmOutput, error) {
	if params == nil {
		params = &AssociateMemberToFarmInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AssociateMemberToFarm", params, optFns, c.addOperationAssociateMemberToFarmMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AssociateMemberToFarmOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateMemberToFarmInput struct {

	// The ID of the farm to associate with the member.
	//
	// This member is required.
	FarmId *string

	// The identity store ID of the member to associate with the farm.
	//
	// This member is required.
	IdentityStoreId *string

	// The principal's membership level for the associated farm.
	//
	// This member is required.
	MembershipLevel types.MembershipLevel

	// The member's principal ID to associate with the farm.
	//
	// This member is required.
	PrincipalId *string

	// The principal type of the member to associate with the farm.
	//
	// This member is required.
	PrincipalType types.DeadlinePrincipalType

	noSmithyDocumentSerde
}

type AssociateMemberToFarmOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAssociateMemberToFarmMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpAssociateMemberToFarm{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpAssociateMemberToFarm{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "AssociateMemberToFarm"); err != nil {
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
	if err = addEndpointPrefix_opAssociateMemberToFarmMiddleware(stack); err != nil {
		return err
	}
	if err = addOpAssociateMemberToFarmValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateMemberToFarm(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opAssociateMemberToFarmMiddleware struct {
}

func (*endpointPrefix_opAssociateMemberToFarmMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opAssociateMemberToFarmMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleFinalize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "management." + req.URL.Host

	return next.HandleFinalize(ctx, in)
}
func addEndpointPrefix_opAssociateMemberToFarmMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opAssociateMemberToFarmMiddleware{}, "ResolveEndpointV2", middleware.After)
}

func newServiceMetadataMiddleware_opAssociateMemberToFarm(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "AssociateMemberToFarm",
	}
}
