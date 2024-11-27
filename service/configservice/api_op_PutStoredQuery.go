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

// Saves a new query or updates an existing saved query. The QueryName must be
// unique for a single Amazon Web Services account and a single Amazon Web Services
// Region. You can create upto 300 queries in a single Amazon Web Services account
// and a single Amazon Web Services Region.
//
// # Tags are added at creation and cannot be updated
//
// PutStoredQuery is an idempotent API. Subsequent requests won’t create a
// duplicate resource if one was already created. If a following request has
// different tags values, Config will ignore these differences and treat it as an
// idempotent request of the previous. In this case, tags will not be updated,
// even if they are different.
func (c *Client) PutStoredQuery(ctx context.Context, params *PutStoredQueryInput, optFns ...func(*Options)) (*PutStoredQueryOutput, error) {
	if params == nil {
		params = &PutStoredQueryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutStoredQuery", params, optFns, c.addOperationPutStoredQueryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutStoredQueryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutStoredQueryInput struct {

	// A list of StoredQuery objects. The mandatory fields are QueryName and Expression
	// .
	//
	// When you are creating a query, you must provide a query name and an expression.
	// When you are updating a query, you must provide a query name but updating the
	// description is optional.
	//
	// This member is required.
	StoredQuery *types.StoredQuery

	// A list of Tags object.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type PutStoredQueryOutput struct {

	// Amazon Resource Name (ARN) of the query. For example,
	// arn:partition:service:region:account-id:resource-type/resource-name/resource-id.
	QueryArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutStoredQueryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutStoredQuery{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutStoredQuery{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PutStoredQuery"); err != nil {
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
	if err = addOpPutStoredQueryValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutStoredQuery(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutStoredQuery(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PutStoredQuery",
	}
}
