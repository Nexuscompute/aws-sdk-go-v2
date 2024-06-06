// Code generated by smithy-go-codegen DO NOT EDIT.

package ssmcontacts

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/ssmcontacts/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Used to acknowledge an engagement to a contact channel during an incident.
func (c *Client) AcceptPage(ctx context.Context, params *AcceptPageInput, optFns ...func(*Options)) (*AcceptPageOutput, error) {
	if params == nil {
		params = &AcceptPageInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AcceptPage", params, optFns, c.addOperationAcceptPageMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AcceptPageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AcceptPageInput struct {

	// A 6-digit code used to acknowledge the page.
	//
	// This member is required.
	AcceptCode *string

	// The type indicates if the page was DELIVERED or READ .
	//
	// This member is required.
	AcceptType types.AcceptType

	// The Amazon Resource Name (ARN) of the engagement to a contact channel.
	//
	// This member is required.
	PageId *string

	// An optional field that Incident Manager uses to ENFORCE AcceptCode validation
	// when acknowledging an page. Acknowledgement can occur by replying to a page, or
	// when entering the AcceptCode in the console. Enforcing AcceptCode validation
	// causes Incident Manager to verify that the code entered by the user matches the
	// code sent by Incident Manager with the page.
	//
	// Incident Manager can also IGNORE AcceptCode validation. Ignoring AcceptCode
	// validation causes Incident Manager to accept any value entered for the
	// AcceptCode .
	AcceptCodeValidation types.AcceptCodeValidation

	// The ARN of the contact channel.
	ContactChannelId *string

	// Information provided by the user when the user acknowledges the page.
	Note *string

	noSmithyDocumentSerde
}

type AcceptPageOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAcceptPageMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAcceptPage{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAcceptPage{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "AcceptPage"); err != nil {
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
	if err = addOpAcceptPageValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAcceptPage(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAcceptPage(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "AcceptPage",
	}
}
