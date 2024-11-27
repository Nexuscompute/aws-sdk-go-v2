// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Authorizes the aggregator account and region to collect data from the source
// account and region.
//
// # Tags are added at creation and cannot be updated with this operation
//
// PutAggregationAuthorization is an idempotent API. Subsequent requests won’t
// create a duplicate resource if one was already created. If a following request
// has different tags values, Config will ignore these differences and treat it as
// an idempotent request of the previous. In this case, tags will not be updated,
// even if they are different.
//
// Use [TagResource] and [UntagResource] to update tags after creation.
//
// [TagResource]: https://docs.aws.amazon.com/config/latest/APIReference/API_TagResource.html
// [UntagResource]: https://docs.aws.amazon.com/config/latest/APIReference/API_UntagResource.html
func (c *Client) PutAggregationAuthorization(ctx context.Context, params *PutAggregationAuthorizationInput, optFns ...func(*Options)) (*PutAggregationAuthorizationOutput, error) {
	if params == nil {
		params = &PutAggregationAuthorizationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutAggregationAuthorization", params, optFns, c.addOperationPutAggregationAuthorizationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutAggregationAuthorizationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutAggregationAuthorizationInput struct {

	// The 12-digit account ID of the account authorized to aggregate data.
	//
	// This member is required.
	AuthorizedAccountId *string

	// The region authorized to collect aggregated data.
	//
	// This member is required.
	AuthorizedAwsRegion *string

	// An array of tag object.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type PutAggregationAuthorizationOutput struct {

	// Returns an AggregationAuthorization object.
	AggregationAuthorization *types.AggregationAuthorization

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutAggregationAuthorizationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutAggregationAuthorization{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutAggregationAuthorization{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PutAggregationAuthorization"); err != nil {
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
	if err = addOpPutAggregationAuthorizationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutAggregationAuthorization(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutAggregationAuthorization(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PutAggregationAuthorization",
	}
}
