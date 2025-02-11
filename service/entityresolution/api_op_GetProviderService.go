// Code generated by smithy-go-codegen DO NOT EDIT.

package entityresolution

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/entityresolution/document"
	"github.com/aws/aws-sdk-go-v2/service/entityresolution/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the ProviderService of a given name.
func (c *Client) GetProviderService(ctx context.Context, params *GetProviderServiceInput, optFns ...func(*Options)) (*GetProviderServiceOutput, error) {
	if params == nil {
		params = &GetProviderServiceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetProviderService", params, optFns, c.addOperationGetProviderServiceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetProviderServiceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetProviderServiceInput struct {

	// The name of the provider. This name is typically the company name.
	//
	// This member is required.
	ProviderName *string

	// The ARN (Amazon Resource Name) of the product that the provider service
	// provides.
	//
	// This member is required.
	ProviderServiceName *string

	noSmithyDocumentSerde
}

type GetProviderServiceOutput struct {

	// Specifies whether output data from the provider is anonymized. A value of TRUE
	// means the output will be anonymized and you can't relate the data that comes
	// back from the provider to the identifying input. A value of FALSE means the
	// output won't be anonymized and you can relate the data that comes back from the
	// provider to your source data.
	//
	// This member is required.
	AnonymizedOutput *bool

	// The required configuration fields to use with the provider service.
	//
	// This member is required.
	ProviderEndpointConfiguration types.ProviderEndpointConfiguration

	// The definition of the provider entity output.
	//
	// This member is required.
	ProviderEntityOutputDefinition document.Interface

	// The name of the provider. This name is typically the company name.
	//
	// This member is required.
	ProviderName *string

	// The ARN (Amazon Resource Name) that Entity Resolution generated for the
	// provider service.
	//
	// This member is required.
	ProviderServiceArn *string

	// The display name of the provider service.
	//
	// This member is required.
	ProviderServiceDisplayName *string

	// The name of the product that the provider service provides.
	//
	// This member is required.
	ProviderServiceName *string

	// The type of provider service.
	//
	// This member is required.
	ProviderServiceType types.ServiceType

	// Input schema for the provider service.
	ProviderComponentSchema *types.ProviderComponentSchema

	// The definition of the provider configuration.
	ProviderConfigurationDefinition document.Interface

	// The provider configuration required for different ID namespace types.
	ProviderIdNameSpaceConfiguration *types.ProviderIdNameSpaceConfiguration

	// The Amazon Web Services accounts and the S3 permissions that are required by
	// some providers to create an S3 bucket for intermediate data storage.
	ProviderIntermediateDataAccessConfiguration *types.ProviderIntermediateDataAccessConfiguration

	// Provider service job configurations.
	ProviderJobConfiguration document.Interface

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetProviderServiceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetProviderService{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetProviderService{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetProviderService"); err != nil {
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
	if err = addOpGetProviderServiceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetProviderService(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetProviderService(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetProviderService",
	}
}
