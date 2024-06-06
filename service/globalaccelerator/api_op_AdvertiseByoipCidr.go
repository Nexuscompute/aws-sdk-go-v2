// Code generated by smithy-go-codegen DO NOT EDIT.

package globalaccelerator

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/globalaccelerator/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Advertises an IPv4 address range that is provisioned for use with your Amazon
// Web Services resources through bring your own IP addresses (BYOIP). It can take
// a few minutes before traffic to the specified addresses starts routing to Amazon
// Web Services because of propagation delays.
//
// To stop advertising the BYOIP address range, use [WithdrawByoipCidr].
//
// For more information, see [Bring your own IP addresses (BYOIP)] in the Global Accelerator Developer Guide.
//
// [Bring your own IP addresses (BYOIP)]: https://docs.aws.amazon.com/global-accelerator/latest/dg/using-byoip.html
// [WithdrawByoipCidr]: https://docs.aws.amazon.com/global-accelerator/latest/api/WithdrawByoipCidr.html
func (c *Client) AdvertiseByoipCidr(ctx context.Context, params *AdvertiseByoipCidrInput, optFns ...func(*Options)) (*AdvertiseByoipCidrOutput, error) {
	if params == nil {
		params = &AdvertiseByoipCidrInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AdvertiseByoipCidr", params, optFns, c.addOperationAdvertiseByoipCidrMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AdvertiseByoipCidrOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AdvertiseByoipCidrInput struct {

	// The address range, in CIDR notation. This must be the exact range that you
	// provisioned. You can't advertise only a portion of the provisioned range.
	//
	// For more information, see [Bring your own IP addresses (BYOIP)] in the Global Accelerator Developer Guide.
	//
	// [Bring your own IP addresses (BYOIP)]: https://docs.aws.amazon.com/global-accelerator/latest/dg/using-byoip.html
	//
	// This member is required.
	Cidr *string

	noSmithyDocumentSerde
}

type AdvertiseByoipCidrOutput struct {

	// Information about the address range.
	ByoipCidr *types.ByoipCidr

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAdvertiseByoipCidrMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAdvertiseByoipCidr{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAdvertiseByoipCidr{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "AdvertiseByoipCidr"); err != nil {
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
	if err = addOpAdvertiseByoipCidrValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAdvertiseByoipCidr(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAdvertiseByoipCidr(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "AdvertiseByoipCidr",
	}
}
