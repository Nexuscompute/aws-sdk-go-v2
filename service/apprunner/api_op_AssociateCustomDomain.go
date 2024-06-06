// Code generated by smithy-go-codegen DO NOT EDIT.

package apprunner

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/apprunner/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Associate your own domain name with the App Runner subdomain URL of your App
// Runner service.
//
// After you call AssociateCustomDomain and receive a successful response, use the
// information in the CustomDomainrecord that's returned to add CNAME records to your Domain
// Name System (DNS). For each mapped domain name, add a mapping to the target App
// Runner subdomain and one or more certificate validation records. App Runner then
// performs DNS validation to verify that you own or control the domain name that
// you associated. App Runner tracks domain validity in a certificate stored in [AWS Certificate Manager (ACM)].
//
// [AWS Certificate Manager (ACM)]: https://docs.aws.amazon.com/acm/latest/userguide
func (c *Client) AssociateCustomDomain(ctx context.Context, params *AssociateCustomDomainInput, optFns ...func(*Options)) (*AssociateCustomDomainOutput, error) {
	if params == nil {
		params = &AssociateCustomDomainInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AssociateCustomDomain", params, optFns, c.addOperationAssociateCustomDomainMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AssociateCustomDomainOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateCustomDomainInput struct {

	// A custom domain endpoint to associate. Specify a root domain (for example,
	// example.com ), a subdomain (for example, login.example.com or
	// admin.login.example.com ), or a wildcard (for example, *.example.com ).
	//
	// This member is required.
	DomainName *string

	// The Amazon Resource Name (ARN) of the App Runner service that you want to
	// associate a custom domain name with.
	//
	// This member is required.
	ServiceArn *string

	// Set to true to associate the subdomain www.DomainName  with the App Runner
	// service in addition to the base domain.
	//
	// Default: true
	EnableWWWSubdomain *bool

	noSmithyDocumentSerde
}

type AssociateCustomDomainOutput struct {

	// A description of the domain name that's being associated.
	//
	// This member is required.
	CustomDomain *types.CustomDomain

	// The App Runner subdomain of the App Runner service. The custom domain name is
	// mapped to this target name.
	//
	// This member is required.
	DNSTarget *string

	// The Amazon Resource Name (ARN) of the App Runner service with which a custom
	// domain name is associated.
	//
	// This member is required.
	ServiceArn *string

	// DNS Target records for the custom domains of this Amazon VPC.
	//
	// This member is required.
	VpcDNSTargets []types.VpcDNSTarget

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAssociateCustomDomainMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpAssociateCustomDomain{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpAssociateCustomDomain{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "AssociateCustomDomain"); err != nil {
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
	if err = addOpAssociateCustomDomainValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateCustomDomain(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAssociateCustomDomain(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "AssociateCustomDomain",
	}
}
