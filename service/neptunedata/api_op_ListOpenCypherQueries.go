// Code generated by smithy-go-codegen DO NOT EDIT.

package neptunedata

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/neptunedata/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists active openCypher queries. See [Neptune openCypher status endpoint] for more information.
//
// When invoking this operation in a Neptune cluster that has IAM authentication
// enabled, the IAM user or role making the request must have a policy attached
// that allows the [neptune-db:GetQueryStatus]IAM action in that cluster.
//
// Note that the [neptune-db:QueryLanguage:OpenCypher] IAM condition key can be used in the policy document to restrict
// the use of openCypher queries (see [Condition keys available in Neptune IAM data-access policy statements]).
//
// [Condition keys available in Neptune IAM data-access policy statements]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-data-condition-keys.html
// [Neptune openCypher status endpoint]: https://docs.aws.amazon.com/neptune/latest/userguide/access-graph-opencypher-status.html
// [neptune-db:GetQueryStatus]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-dp-actions.html#getquerystatus
// [neptune-db:QueryLanguage:OpenCypher]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-data-condition-keys.html#iam-neptune-condition-keys
func (c *Client) ListOpenCypherQueries(ctx context.Context, params *ListOpenCypherQueriesInput, optFns ...func(*Options)) (*ListOpenCypherQueriesOutput, error) {
	if params == nil {
		params = &ListOpenCypherQueriesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListOpenCypherQueries", params, optFns, c.addOperationListOpenCypherQueriesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListOpenCypherQueriesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListOpenCypherQueriesInput struct {

	//  When set to TRUE and other parameters are not present, causes status
	// information to be returned for waiting queries as well as for running queries.
	IncludeWaiting *bool

	noSmithyDocumentSerde
}

type ListOpenCypherQueriesOutput struct {

	// The number of queries that have been accepted but not yet completed, including
	// queries in the queue.
	AcceptedQueryCount *int32

	// A list of current openCypher queries.
	Queries []types.GremlinQueryStatus

	// The number of currently running openCypher queries.
	RunningQueryCount *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListOpenCypherQueriesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListOpenCypherQueries{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListOpenCypherQueries{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListOpenCypherQueries"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListOpenCypherQueries(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListOpenCypherQueries(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListOpenCypherQueries",
	}
}
