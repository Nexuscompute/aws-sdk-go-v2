// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the content and status of IP rules. Traffic from a source is allowed
// when the source satisfies either the IpRestrictionRule , VpcIdRestrictionRule ,
// or VpcEndpointIdRestrictionRule . To use this operation, you must provide the
// entire map of rules. You can use the DescribeIpRestriction operation to get the
// current rule map.
func (c *Client) UpdateIpRestriction(ctx context.Context, params *UpdateIpRestrictionInput, optFns ...func(*Options)) (*UpdateIpRestrictionOutput, error) {
	if params == nil {
		params = &UpdateIpRestrictionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateIpRestriction", params, optFns, c.addOperationUpdateIpRestrictionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateIpRestrictionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateIpRestrictionInput struct {

	// The ID of the Amazon Web Services account that contains the IP rules.
	//
	// This member is required.
	AwsAccountId *string

	// A value that specifies whether IP rules are turned on.
	Enabled *bool

	// A map that describes the updated IP rules with CIDR ranges and descriptions.
	IpRestrictionRuleMap map[string]string

	// A map of allowed VPC endpoint IDs and their corresponding rule descriptions.
	VpcEndpointIdRestrictionRuleMap map[string]string

	// A map of VPC IDs and their corresponding rules. When you configure this
	// parameter, traffic from all VPC endpoints that are present in the specified VPC
	// is allowed.
	VpcIdRestrictionRuleMap map[string]string

	noSmithyDocumentSerde
}

type UpdateIpRestrictionOutput struct {

	// The ID of the Amazon Web Services account that contains the IP rules.
	AwsAccountId *string

	// The Amazon Web Services request ID for this operation.
	RequestId *string

	// The HTTP status of the request.
	Status int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateIpRestrictionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateIpRestriction{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateIpRestriction{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateIpRestriction"); err != nil {
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
	if err = addOpUpdateIpRestrictionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateIpRestriction(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateIpRestriction(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateIpRestriction",
	}
}
