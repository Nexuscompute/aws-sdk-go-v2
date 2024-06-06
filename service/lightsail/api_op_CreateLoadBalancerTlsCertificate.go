// Code generated by smithy-go-codegen DO NOT EDIT.

package lightsail

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an SSL/TLS certificate for an Amazon Lightsail load balancer.
//
// TLS is just an updated, more secure version of Secure Socket Layer (SSL).
//
// The CreateLoadBalancerTlsCertificate operation supports tag-based access
// control via resource tags applied to the resource identified by load balancer
// name . For more information, see the [Amazon Lightsail Developer Guide].
//
// [Amazon Lightsail Developer Guide]: https://lightsail.aws.amazon.com/ls/docs/en_us/articles/amazon-lightsail-controlling-access-using-tags
func (c *Client) CreateLoadBalancerTlsCertificate(ctx context.Context, params *CreateLoadBalancerTlsCertificateInput, optFns ...func(*Options)) (*CreateLoadBalancerTlsCertificateOutput, error) {
	if params == nil {
		params = &CreateLoadBalancerTlsCertificateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateLoadBalancerTlsCertificate", params, optFns, c.addOperationCreateLoadBalancerTlsCertificateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateLoadBalancerTlsCertificateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateLoadBalancerTlsCertificateInput struct {

	// The domain name ( example.com ) for your SSL/TLS certificate.
	//
	// This member is required.
	CertificateDomainName *string

	// The SSL/TLS certificate name.
	//
	// You can have up to 10 certificates in your account at one time. Each Lightsail
	// load balancer can have up to 2 certificates associated with it at one time.
	// There is also an overall limit to the number of certificates that can be issue
	// in a 365-day period. For more information, see [Limits].
	//
	// [Limits]: http://docs.aws.amazon.com/acm/latest/userguide/acm-limits.html
	//
	// This member is required.
	CertificateName *string

	// The load balancer name where you want to create the SSL/TLS certificate.
	//
	// This member is required.
	LoadBalancerName *string

	// An array of strings listing alternative domains and subdomains for your SSL/TLS
	// certificate. Lightsail will de-dupe the names for you. You can have a maximum of
	// 9 alternative names (in addition to the 1 primary domain). We do not support
	// wildcards ( *.example.com ).
	CertificateAlternativeNames []string

	// The tag keys and optional values to add to the resource during create.
	//
	// Use the TagResource action to tag a resource after it's created.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateLoadBalancerTlsCertificateOutput struct {

	// An array of objects that describe the result of the action, such as the status
	// of the request, the timestamp of the request, and the resources affected by the
	// request.
	Operations []types.Operation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateLoadBalancerTlsCertificateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateLoadBalancerTlsCertificate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateLoadBalancerTlsCertificate{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateLoadBalancerTlsCertificate"); err != nil {
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
	if err = addOpCreateLoadBalancerTlsCertificateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateLoadBalancerTlsCertificate(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateLoadBalancerTlsCertificate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateLoadBalancerTlsCertificate",
	}
}
