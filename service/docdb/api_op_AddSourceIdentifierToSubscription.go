// Code generated by smithy-go-codegen DO NOT EDIT.

package docdb

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/docdb/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Adds a source identifier to an existing event notification subscription.
func (c *Client) AddSourceIdentifierToSubscription(ctx context.Context, params *AddSourceIdentifierToSubscriptionInput, optFns ...func(*Options)) (*AddSourceIdentifierToSubscriptionOutput, error) {
	if params == nil {
		params = &AddSourceIdentifierToSubscriptionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AddSourceIdentifierToSubscription", params, optFns, c.addOperationAddSourceIdentifierToSubscriptionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AddSourceIdentifierToSubscriptionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input to AddSourceIdentifierToSubscription.
type AddSourceIdentifierToSubscriptionInput struct {

	// The identifier of the event source to be added:
	//
	//   - If the source type is an instance, a DBInstanceIdentifier must be provided.
	//
	//   - If the source type is a security group, a DBSecurityGroupName must be
	//   provided.
	//
	//   - If the source type is a parameter group, a DBParameterGroupName must be
	//   provided.
	//
	//   - If the source type is a snapshot, a DBSnapshotIdentifier must be provided.
	//
	// This member is required.
	SourceIdentifier *string

	// The name of the Amazon DocumentDB event notification subscription that you want
	// to add a source identifier to.
	//
	// This member is required.
	SubscriptionName *string

	noSmithyDocumentSerde
}

type AddSourceIdentifierToSubscriptionOutput struct {

	// Detailed information about an event to which you have subscribed.
	EventSubscription *types.EventSubscription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAddSourceIdentifierToSubscriptionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpAddSourceIdentifierToSubscription{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpAddSourceIdentifierToSubscription{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "AddSourceIdentifierToSubscription"); err != nil {
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
	if err = addOpAddSourceIdentifierToSubscriptionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAddSourceIdentifierToSubscription(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAddSourceIdentifierToSubscription(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "AddSourceIdentifierToSubscription",
	}
}
