// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a thing record in the registry. If this call is made multiple times
// using the same thing name and configuration, the call will succeed. If this call
// is made with the same thing name but different configuration a
// ResourceAlreadyExistsException is thrown.
//
// This is a control plane operation. See [Authorization] for information about authorizing
// control plane actions.
//
// Requires permission to access the [CreateThing] action.
//
// [Authorization]: https://docs.aws.amazon.com/iot/latest/developerguide/iot-authorization.html
// [CreateThing]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func (c *Client) CreateThing(ctx context.Context, params *CreateThingInput, optFns ...func(*Options)) (*CreateThingOutput, error) {
	if params == nil {
		params = &CreateThingInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateThing", params, optFns, c.addOperationCreateThingMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateThingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for the CreateThing operation.
type CreateThingInput struct {

	// The name of the thing to create.
	//
	// You can't change a thing's name after you create it. To change a thing's name,
	// you must create a new thing, give it the new name, and then delete the old
	// thing.
	//
	// This member is required.
	ThingName *string

	// The attribute payload, which consists of up to three name/value pairs in a JSON
	// document. For example:
	//
	//     {\"attributes\":{\"string1\":\"string2\"}}
	AttributePayload *types.AttributePayload

	// The name of the billing group the thing will be added to.
	BillingGroupName *string

	// The name of the thing type associated with the new thing.
	ThingTypeName *string

	noSmithyDocumentSerde
}

// The output of the CreateThing operation.
type CreateThingOutput struct {

	// The ARN of the new thing.
	ThingArn *string

	// The thing ID.
	ThingId *string

	// The name of the new thing.
	ThingName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateThingMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateThing{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateThing{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateThing"); err != nil {
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
	if err = addOpCreateThingValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateThing(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateThing(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateThing",
	}
}
