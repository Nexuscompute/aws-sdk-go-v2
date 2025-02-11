// Code generated by smithy-go-codegen DO NOT EDIT.

package customerprofiles

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/customerprofiles/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Updates an existing calculated attribute definition. When updating the
// Conditions, note that increasing the date range of a calculated attribute will
// not trigger inclusion of historical data greater than the current date range.
func (c *Client) UpdateCalculatedAttributeDefinition(ctx context.Context, params *UpdateCalculatedAttributeDefinitionInput, optFns ...func(*Options)) (*UpdateCalculatedAttributeDefinitionOutput, error) {
	if params == nil {
		params = &UpdateCalculatedAttributeDefinitionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateCalculatedAttributeDefinition", params, optFns, c.addOperationUpdateCalculatedAttributeDefinitionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateCalculatedAttributeDefinitionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateCalculatedAttributeDefinitionInput struct {

	// The unique name of the calculated attribute.
	//
	// This member is required.
	CalculatedAttributeName *string

	// The unique name of the domain.
	//
	// This member is required.
	DomainName *string

	// The conditions including range, object count, and threshold for the calculated
	// attribute.
	Conditions *types.Conditions

	// The description of the calculated attribute.
	Description *string

	// The display name of the calculated attribute.
	DisplayName *string

	noSmithyDocumentSerde
}

type UpdateCalculatedAttributeDefinitionOutput struct {

	// The mathematical expression and a list of attribute items specified in that
	// expression.
	AttributeDetails *types.AttributeDetails

	// The unique name of the calculated attribute.
	CalculatedAttributeName *string

	// The conditions including range, object count, and threshold for the calculated
	// attribute.
	Conditions *types.Conditions

	// The timestamp of when the calculated attribute definition was created.
	CreatedAt *time.Time

	// The description of the calculated attribute.
	Description *string

	// The display name of the calculated attribute.
	DisplayName *string

	// The timestamp of when the calculated attribute definition was most recently
	// edited.
	LastUpdatedAt *time.Time

	// The aggregation operation to perform for the calculated attribute.
	Statistic types.Statistic

	// The tags used to organize, track, or control access for this resource.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateCalculatedAttributeDefinitionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateCalculatedAttributeDefinition{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateCalculatedAttributeDefinition{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateCalculatedAttributeDefinition"); err != nil {
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
	if err = addOpUpdateCalculatedAttributeDefinitionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateCalculatedAttributeDefinition(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateCalculatedAttributeDefinition(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateCalculatedAttributeDefinition",
	}
}
