// Code generated by smithy-go-codegen DO NOT EDIT.

package databasemigrationservice

import (
	"context"
	"errors"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice/types"
	"github.com/aws/smithy-go/middleware"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	smithywaiter "github.com/aws/smithy-go/waiter"
	"github.com/jmespath/go-jmespath"
	"time"
)

// Returns information about the endpoints for your account in the current region.
func (c *Client) DescribeEndpoints(ctx context.Context, params *DescribeEndpointsInput, optFns ...func(*Options)) (*DescribeEndpointsOutput, error) {
	if params == nil {
		params = &DescribeEndpointsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeEndpoints", params, optFns, c.addOperationDescribeEndpointsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeEndpointsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeEndpointsInput struct {

	// Filters applied to the endpoints.
	//
	// Valid filter names: endpoint-arn | endpoint-type | endpoint-id | engine-name
	Filters []types.Filter

	//  An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to the
	// value specified by MaxRecords .
	Marker *string

	//  The maximum number of records to include in the response. If more records
	// exist than the specified MaxRecords value, a pagination token called a marker
	// is included in the response so that the remaining results can be retrieved.
	//
	// Default: 100
	//
	// Constraints: Minimum 20, maximum 100.
	MaxRecords *int32

	noSmithyDocumentSerde
}

type DescribeEndpointsOutput struct {

	// Endpoint description.
	Endpoints []types.Endpoint

	//  An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to the
	// value specified by MaxRecords .
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeEndpointsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeEndpoints{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeEndpoints{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeEndpoints"); err != nil {
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
	if err = addOpDescribeEndpointsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeEndpoints(options.Region), middleware.Before); err != nil {
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

// DescribeEndpointsAPIClient is a client that implements the DescribeEndpoints
// operation.
type DescribeEndpointsAPIClient interface {
	DescribeEndpoints(context.Context, *DescribeEndpointsInput, ...func(*Options)) (*DescribeEndpointsOutput, error)
}

var _ DescribeEndpointsAPIClient = (*Client)(nil)

// DescribeEndpointsPaginatorOptions is the paginator options for DescribeEndpoints
type DescribeEndpointsPaginatorOptions struct {
	//  The maximum number of records to include in the response. If more records
	// exist than the specified MaxRecords value, a pagination token called a marker
	// is included in the response so that the remaining results can be retrieved.
	//
	// Default: 100
	//
	// Constraints: Minimum 20, maximum 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeEndpointsPaginator is a paginator for DescribeEndpoints
type DescribeEndpointsPaginator struct {
	options   DescribeEndpointsPaginatorOptions
	client    DescribeEndpointsAPIClient
	params    *DescribeEndpointsInput
	nextToken *string
	firstPage bool
}

// NewDescribeEndpointsPaginator returns a new DescribeEndpointsPaginator
func NewDescribeEndpointsPaginator(client DescribeEndpointsAPIClient, params *DescribeEndpointsInput, optFns ...func(*DescribeEndpointsPaginatorOptions)) *DescribeEndpointsPaginator {
	if params == nil {
		params = &DescribeEndpointsInput{}
	}

	options := DescribeEndpointsPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeEndpointsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeEndpointsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeEndpoints page.
func (p *DescribeEndpointsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeEndpointsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxRecords = limit

	result, err := p.client.DescribeEndpoints(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.Marker

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

// EndpointDeletedWaiterOptions are waiter options for EndpointDeletedWaiter
type EndpointDeletedWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	//
	// Passing options here is functionally equivalent to passing values to this
	// config's ClientOptions field that extend the inner client's APIOptions directly.
	APIOptions []func(*middleware.Stack) error

	// Functional options to be passed to all operations invoked by this client.
	//
	// Function values that modify the inner APIOptions are applied after the waiter
	// config's own APIOptions modifiers.
	ClientOptions []func(*Options)

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// EndpointDeletedWaiter will use default minimum delay of 5 seconds. Note that
	// MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or
	// set to zero, EndpointDeletedWaiter will use default max delay of 120 seconds.
	// Note that MaxDelay must resolve to value greater than or equal to the MinDelay.
	MaxDelay time.Duration

	// LogWaitAttempts is used to enable logging for waiter retry attempts
	LogWaitAttempts bool

	// Retryable is function that can be used to override the service defined
	// waiter-behavior based on operation output, or returned error. This function is
	// used by the waiter to decide if a state is retryable or a terminal state.
	//
	// By default service-modeled logic will populate this option. This option can
	// thus be used to define a custom waiter state with fall-back to service-modeled
	// waiter state mutators.The function returns an error in case of a failure state.
	// In case of retry state, this function returns a bool value of true and nil
	// error, while in case of success it returns a bool value of false and nil error.
	Retryable func(context.Context, *DescribeEndpointsInput, *DescribeEndpointsOutput, error) (bool, error)
}

// EndpointDeletedWaiter defines the waiters for EndpointDeleted
type EndpointDeletedWaiter struct {
	client DescribeEndpointsAPIClient

	options EndpointDeletedWaiterOptions
}

// NewEndpointDeletedWaiter constructs a EndpointDeletedWaiter.
func NewEndpointDeletedWaiter(client DescribeEndpointsAPIClient, optFns ...func(*EndpointDeletedWaiterOptions)) *EndpointDeletedWaiter {
	options := EndpointDeletedWaiterOptions{}
	options.MinDelay = 5 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = endpointDeletedStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &EndpointDeletedWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for EndpointDeleted waiter. The maxWaitDur is
// the maximum wait duration the waiter will wait. The maxWaitDur is required and
// must be greater than zero.
func (w *EndpointDeletedWaiter) Wait(ctx context.Context, params *DescribeEndpointsInput, maxWaitDur time.Duration, optFns ...func(*EndpointDeletedWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for EndpointDeleted waiter and returns
// the output of the successful operation. The maxWaitDur is the maximum wait
// duration the waiter will wait. The maxWaitDur is required and must be greater
// than zero.
func (w *EndpointDeletedWaiter) WaitForOutput(ctx context.Context, params *DescribeEndpointsInput, maxWaitDur time.Duration, optFns ...func(*EndpointDeletedWaiterOptions)) (*DescribeEndpointsOutput, error) {
	if maxWaitDur <= 0 {
		return nil, fmt.Errorf("maximum wait time for waiter must be greater than zero")
	}

	options := w.options
	for _, fn := range optFns {
		fn(&options)
	}

	if options.MaxDelay <= 0 {
		options.MaxDelay = 120 * time.Second
	}

	if options.MinDelay > options.MaxDelay {
		return nil, fmt.Errorf("minimum waiter delay %v must be lesser than or equal to maximum waiter delay of %v.", options.MinDelay, options.MaxDelay)
	}

	ctx, cancelFn := context.WithTimeout(ctx, maxWaitDur)
	defer cancelFn()

	logger := smithywaiter.Logger{}
	remainingTime := maxWaitDur

	var attempt int64
	for {

		attempt++
		apiOptions := options.APIOptions
		start := time.Now()

		if options.LogWaitAttempts {
			logger.Attempt = attempt
			apiOptions = append([]func(*middleware.Stack) error{}, options.APIOptions...)
			apiOptions = append(apiOptions, logger.AddLogger)
		}

		out, err := w.client.DescribeEndpoints(ctx, params, func(o *Options) {
			o.APIOptions = append(o.APIOptions, apiOptions...)
			for _, opt := range options.ClientOptions {
				opt(o)
			}
		})

		retryable, err := options.Retryable(ctx, params, out, err)
		if err != nil {
			return nil, err
		}
		if !retryable {
			return out, nil
		}

		remainingTime -= time.Since(start)
		if remainingTime < options.MinDelay || remainingTime <= 0 {
			break
		}

		// compute exponential backoff between waiter retries
		delay, err := smithywaiter.ComputeDelay(
			attempt, options.MinDelay, options.MaxDelay, remainingTime,
		)
		if err != nil {
			return nil, fmt.Errorf("error computing waiter delay, %w", err)
		}

		remainingTime -= delay
		// sleep for the delay amount before invoking a request
		if err := smithytime.SleepWithContext(ctx, delay); err != nil {
			return nil, fmt.Errorf("request cancelled while waiting, %w", err)
		}
	}
	return nil, fmt.Errorf("exceeded max wait time for EndpointDeleted waiter")
}

func endpointDeletedStateRetryable(ctx context.Context, input *DescribeEndpointsInput, output *DescribeEndpointsOutput, err error) (bool, error) {

	if err != nil {
		var errorType *types.ResourceNotFoundFault
		if errors.As(err, &errorType) {
			return false, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("Endpoints[].Status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "active"
		listOfValues, ok := pathValue.([]interface{})
		if !ok {
			return false, fmt.Errorf("waiter comparator expected list got %T", pathValue)
		}

		for _, v := range listOfValues {
			value, ok := v.(*string)
			if !ok {
				return false, fmt.Errorf("waiter comparator expected *string value, got %T", pathValue)
			}

			if string(*value) == expectedValue {
				return false, fmt.Errorf("waiter state transitioned to Failure")
			}
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("Endpoints[].Status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "creating"
		listOfValues, ok := pathValue.([]interface{})
		if !ok {
			return false, fmt.Errorf("waiter comparator expected list got %T", pathValue)
		}

		for _, v := range listOfValues {
			value, ok := v.(*string)
			if !ok {
				return false, fmt.Errorf("waiter comparator expected *string value, got %T", pathValue)
			}

			if string(*value) == expectedValue {
				return false, fmt.Errorf("waiter state transitioned to Failure")
			}
		}
	}

	return true, nil
}

func newServiceMetadataMiddleware_opDescribeEndpoints(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeEndpoints",
	}
}
