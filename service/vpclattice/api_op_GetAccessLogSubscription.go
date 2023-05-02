// Code generated by smithy-go-codegen DO NOT EDIT.

package vpclattice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Retrieves information about the specified access log subscription.
func (c *Client) GetAccessLogSubscription(ctx context.Context, params *GetAccessLogSubscriptionInput, optFns ...func(*Options)) (*GetAccessLogSubscriptionOutput, error) {
	if params == nil {
		params = &GetAccessLogSubscriptionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetAccessLogSubscription", params, optFns, c.addOperationGetAccessLogSubscriptionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetAccessLogSubscriptionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetAccessLogSubscriptionInput struct {

	// The ID or Amazon Resource Name (ARN) of the access log subscription.
	//
	// This member is required.
	AccessLogSubscriptionIdentifier *string

	noSmithyDocumentSerde
}

type GetAccessLogSubscriptionOutput struct {

	// The Amazon Resource Name (ARN) of the access log subscription.
	//
	// This member is required.
	Arn *string

	// The date and time that the access log subscription was created, specified in
	// ISO-8601 format.
	//
	// This member is required.
	CreatedAt *time.Time

	// The Amazon Resource Name (ARN) of the access log destination.
	//
	// This member is required.
	DestinationArn *string

	// The ID of the access log subscription.
	//
	// This member is required.
	Id *string

	// The date and time that the access log subscription was last updated, specified
	// in ISO-8601 format.
	//
	// This member is required.
	LastUpdatedAt *time.Time

	// The Amazon Resource Name (ARN) of the service network or service.
	//
	// This member is required.
	ResourceArn *string

	// The ID of the service network or service.
	//
	// This member is required.
	ResourceId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetAccessLogSubscriptionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetAccessLogSubscription{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetAccessLogSubscription{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpGetAccessLogSubscriptionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetAccessLogSubscription(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
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
	return nil
}

func newServiceMetadataMiddleware_opGetAccessLogSubscription(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "vpc-lattice",
		OperationName: "GetAccessLogSubscription",
	}
}