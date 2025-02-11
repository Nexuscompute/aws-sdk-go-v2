// Code generated by smithy-go-codegen DO NOT EDIT.

package socialmessaging

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/socialmessaging/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Add an event destination to log event data from WhatsApp for a WhatsApp
// Business Account (WABA). A WABA can only have one event destination at a time.
// All resources associated with the WABA use the same event destination.
func (c *Client) PutWhatsAppBusinessAccountEventDestinations(ctx context.Context, params *PutWhatsAppBusinessAccountEventDestinationsInput, optFns ...func(*Options)) (*PutWhatsAppBusinessAccountEventDestinationsOutput, error) {
	if params == nil {
		params = &PutWhatsAppBusinessAccountEventDestinationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutWhatsAppBusinessAccountEventDestinations", params, optFns, c.addOperationPutWhatsAppBusinessAccountEventDestinationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutWhatsAppBusinessAccountEventDestinationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutWhatsAppBusinessAccountEventDestinationsInput struct {

	// An array of WhatsAppBusinessAccountEventDestination event destinations.
	//
	// This member is required.
	EventDestinations []types.WhatsAppBusinessAccountEventDestination

	// The unique identifier of your WhatsApp Business Account. WABA identifiers are
	// formatted as waba-01234567890123456789012345678901 . Use [ListLinkedWhatsAppBusinessAccounts] to list all WABAs and
	// their details.
	//
	// [ListLinkedWhatsAppBusinessAccounts]: https://docs.aws.amazon.com/social-messaging/latest/APIReference/API_ListLinkedWhatsAppBusinessAccounts.html
	//
	// This member is required.
	Id *string

	noSmithyDocumentSerde
}

type PutWhatsAppBusinessAccountEventDestinationsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutWhatsAppBusinessAccountEventDestinationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutWhatsAppBusinessAccountEventDestinations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutWhatsAppBusinessAccountEventDestinations{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PutWhatsAppBusinessAccountEventDestinations"); err != nil {
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
	if err = addOpPutWhatsAppBusinessAccountEventDestinationsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutWhatsAppBusinessAccountEventDestinations(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutWhatsAppBusinessAccountEventDestinations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PutWhatsAppBusinessAccountEventDestinations",
	}
}
