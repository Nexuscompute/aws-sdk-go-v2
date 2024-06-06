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

// Retrieves a list of the loadIds for all active loader jobs.
//
// When invoking this operation in a Neptune cluster that has IAM authentication
// enabled, the IAM user or role making the request must have a policy attached
// that allows the [neptune-db:ListLoaderJobs]IAM action in that cluster..
//
// [neptune-db:ListLoaderJobs]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-dp-actions.html#listloaderjobs
func (c *Client) ListLoaderJobs(ctx context.Context, params *ListLoaderJobsInput, optFns ...func(*Options)) (*ListLoaderJobsOutput, error) {
	if params == nil {
		params = &ListLoaderJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListLoaderJobs", params, optFns, c.addOperationListLoaderJobsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListLoaderJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListLoaderJobsInput struct {

	// An optional parameter that can be used to exclude the load IDs of queued load
	// requests when requesting a list of load IDs by setting the parameter to FALSE .
	// The default value is TRUE .
	IncludeQueuedLoads *bool

	// The number of load IDs to list. Must be a positive integer greater than zero
	// and not more than 100 (which is the default).
	Limit *int32

	noSmithyDocumentSerde
}

type ListLoaderJobsOutput struct {

	// The requested list of job IDs.
	//
	// This member is required.
	Payload *types.LoaderIdResult

	// Returns the status of the job list request.
	//
	// This member is required.
	Status *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListLoaderJobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListLoaderJobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListLoaderJobs{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListLoaderJobs"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListLoaderJobs(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListLoaderJobs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListLoaderJobs",
	}
}
