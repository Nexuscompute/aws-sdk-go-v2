// Code generated by smithy-go-codegen DO NOT EDIT.

package b2bi

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/b2bi/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Updates some of the parameters for a capability, based on the specified
// parameters. A trading capability contains the information required to transform
// incoming EDI documents into JSON or XML outputs.
func (c *Client) UpdateCapability(ctx context.Context, params *UpdateCapabilityInput, optFns ...func(*Options)) (*UpdateCapabilityOutput, error) {
	if params == nil {
		params = &UpdateCapabilityInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateCapability", params, optFns, c.addOperationUpdateCapabilityMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateCapabilityOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateCapabilityInput struct {

	// Specifies a system-assigned unique identifier for the capability.
	//
	// This member is required.
	CapabilityId *string

	// Specifies a structure that contains the details for a capability.
	Configuration types.CapabilityConfiguration

	// Specifies one or more locations in Amazon S3, each specifying an EDI document
	// that can be used with this capability. Each item contains the name of the bucket
	// and the key, to identify the document's location.
	InstructionsDocuments []types.S3Location

	// Specifies a new name for the capability, to replace the existing name.
	Name *string

	noSmithyDocumentSerde
}

type UpdateCapabilityOutput struct {

	// Returns an Amazon Resource Name (ARN) for a specific Amazon Web Services
	// resource, such as a capability, partnership, profile, or transformer.
	//
	// This member is required.
	CapabilityArn *string

	// Returns a system-assigned unique identifier for the capability.
	//
	// This member is required.
	CapabilityId *string

	// Returns a structure that contains the details for a capability.
	//
	// This member is required.
	Configuration types.CapabilityConfiguration

	// Returns a timestamp for creation date and time of the capability.
	//
	// This member is required.
	CreatedAt *time.Time

	// Returns the name of the capability, used to identify it.
	//
	// This member is required.
	Name *string

	// Returns the type of the capability. Currently, only edi is supported.
	//
	// This member is required.
	Type types.CapabilityType

	// Returns one or more locations in Amazon S3, each specifying an EDI document
	// that can be used with this capability. Each item contains the name of the bucket
	// and the key, to identify the document's location.
	InstructionsDocuments []types.S3Location

	// Returns a timestamp for last time the capability was modified.
	ModifiedAt *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateCapabilityMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpUpdateCapability{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpUpdateCapability{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateCapability"); err != nil {
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
	if err = addOpUpdateCapabilityValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateCapability(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateCapability(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateCapability",
	}
}
