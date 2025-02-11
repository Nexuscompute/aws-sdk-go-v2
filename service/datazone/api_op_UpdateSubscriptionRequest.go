// Code generated by smithy-go-codegen DO NOT EDIT.

package datazone

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/datazone/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Updates a specified subscription request in Amazon DataZone.
func (c *Client) UpdateSubscriptionRequest(ctx context.Context, params *UpdateSubscriptionRequestInput, optFns ...func(*Options)) (*UpdateSubscriptionRequestOutput, error) {
	if params == nil {
		params = &UpdateSubscriptionRequestInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateSubscriptionRequest", params, optFns, c.addOperationUpdateSubscriptionRequestMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateSubscriptionRequestOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateSubscriptionRequestInput struct {

	// The identifier of the Amazon DataZone domain in which a subscription request is
	// to be updated.
	//
	// This member is required.
	DomainIdentifier *string

	// The identifier of the subscription request that is to be updated.
	//
	// This member is required.
	Identifier *string

	// The reason for the UpdateSubscriptionRequest action.
	//
	// This member is required.
	RequestReason *string

	noSmithyDocumentSerde
}

type UpdateSubscriptionRequestOutput struct {

	// The timestamp of when the subscription request was created.
	//
	// This member is required.
	CreatedAt *time.Time

	// The Amazon DataZone user who created the subscription request.
	//
	// This member is required.
	CreatedBy *string

	// The identifier of the Amazon DataZone domain in which a subscription request is
	// to be updated.
	//
	// This member is required.
	DomainId *string

	// The identifier of the subscription request that is to be updated.
	//
	// This member is required.
	Id *string

	// The reason for the UpdateSubscriptionRequest action.
	//
	// This member is required.
	RequestReason *string

	// The status of the subscription request.
	//
	// This member is required.
	Status types.SubscriptionRequestStatus

	// The subscribed listings of the subscription request.
	//
	// This member is required.
	SubscribedListings []types.SubscribedListing

	// The subscribed principals of the subscription request.
	//
	// This member is required.
	SubscribedPrincipals []types.SubscribedPrincipal

	// The timestamp of when the subscription request was updated.
	//
	// This member is required.
	UpdatedAt *time.Time

	// The decision comment of the UpdateSubscriptionRequest action.
	DecisionComment *string

	// The ID of the existing subscription.
	ExistingSubscriptionId *string

	// Metadata forms included in the subscription request.
	MetadataForms []types.FormOutput

	// The identifier of the Amazon DataZone user who reviews the subscription request.
	ReviewerId *string

	// The Amazon DataZone user who updated the subscription request.
	UpdatedBy *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateSubscriptionRequestMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateSubscriptionRequest{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateSubscriptionRequest{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateSubscriptionRequest"); err != nil {
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
	if err = addOpUpdateSubscriptionRequestValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateSubscriptionRequest(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateSubscriptionRequest(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateSubscriptionRequest",
	}
}
