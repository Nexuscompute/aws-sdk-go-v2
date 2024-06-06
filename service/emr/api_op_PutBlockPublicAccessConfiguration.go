// Code generated by smithy-go-codegen DO NOT EDIT.

package emr

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/emr/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates or updates an Amazon EMR block public access configuration for your
// Amazon Web Services account in the current Region. For more information see [Configure Block Public Access for Amazon EMR]in
// the Amazon EMR Management Guide.
//
// [Configure Block Public Access for Amazon EMR]: https://docs.aws.amazon.com/emr/latest/ManagementGuide/configure-block-public-access.html
func (c *Client) PutBlockPublicAccessConfiguration(ctx context.Context, params *PutBlockPublicAccessConfigurationInput, optFns ...func(*Options)) (*PutBlockPublicAccessConfigurationOutput, error) {
	if params == nil {
		params = &PutBlockPublicAccessConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutBlockPublicAccessConfiguration", params, optFns, c.addOperationPutBlockPublicAccessConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutBlockPublicAccessConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutBlockPublicAccessConfigurationInput struct {

	// A configuration for Amazon EMR block public access. The configuration applies
	// to all clusters created in your account for the current Region. The
	// configuration specifies whether block public access is enabled. If block public
	// access is enabled, security groups associated with the cluster cannot have rules
	// that allow inbound traffic from 0.0.0.0/0 or ::/0 on a port, unless the port is
	// specified as an exception using PermittedPublicSecurityGroupRuleRanges in the
	// BlockPublicAccessConfiguration . By default, Port 22 (SSH) is an exception, and
	// public access is allowed on this port. You can change this by updating
	// BlockPublicSecurityGroupRules to remove the exception.
	//
	// For accounts that created clusters in a Region before November 25, 2019, block
	// public access is disabled by default in that Region. To use this feature, you
	// must manually enable and configure it. For accounts that did not create an
	// Amazon EMR cluster in a Region before this date, block public access is enabled
	// by default in that Region.
	//
	// This member is required.
	BlockPublicAccessConfiguration *types.BlockPublicAccessConfiguration

	noSmithyDocumentSerde
}

type PutBlockPublicAccessConfigurationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutBlockPublicAccessConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutBlockPublicAccessConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutBlockPublicAccessConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PutBlockPublicAccessConfiguration"); err != nil {
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
	if err = addOpPutBlockPublicAccessConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutBlockPublicAccessConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutBlockPublicAccessConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PutBlockPublicAccessConfiguration",
	}
}
