// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudfront

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets a list of response headers policies.
//
// You can optionally apply a filter to get only the managed policies created by
// Amazon Web Services, or only the custom policies created in your Amazon Web
// Services account.
//
// You can optionally specify the maximum number of items to receive in the
// response. If the total number of items in the list exceeds the maximum that you
// specify, or the default maximum, the response is paginated. To get the next page
// of items, send a subsequent request that specifies the NextMarker value from
// the current response as the Marker value in the subsequent request.
func (c *Client) ListResponseHeadersPolicies(ctx context.Context, params *ListResponseHeadersPoliciesInput, optFns ...func(*Options)) (*ListResponseHeadersPoliciesOutput, error) {
	if params == nil {
		params = &ListResponseHeadersPoliciesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListResponseHeadersPolicies", params, optFns, c.addOperationListResponseHeadersPoliciesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListResponseHeadersPoliciesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListResponseHeadersPoliciesInput struct {

	// Use this field when paginating results to indicate where to begin in your list
	// of response headers policies. The response includes response headers policies in
	// the list that occur after the marker. To get the next page of the list, set this
	// field's value to the value of NextMarker from the current page's response.
	Marker *string

	// The maximum number of response headers policies that you want to get in the
	// response.
	MaxItems *int32

	// A filter to get only the specified kind of response headers policies. Valid
	// values are:
	//
	//   - managed – Gets only the managed policies created by Amazon Web Services.
	//
	//   - custom – Gets only the custom policies created in your Amazon Web Services
	//   account.
	Type types.ResponseHeadersPolicyType

	noSmithyDocumentSerde
}

type ListResponseHeadersPoliciesOutput struct {

	// A list of response headers policies.
	ResponseHeadersPolicyList *types.ResponseHeadersPolicyList

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListResponseHeadersPoliciesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestxml_serializeOpListResponseHeadersPolicies{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpListResponseHeadersPolicies{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListResponseHeadersPolicies"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListResponseHeadersPolicies(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListResponseHeadersPolicies(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListResponseHeadersPolicies",
	}
}
