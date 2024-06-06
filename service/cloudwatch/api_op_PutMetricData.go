// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatch

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
	"github.com/aws/smithy-go/middleware"
	smithyrequestcompression "github.com/aws/smithy-go/private/requestcompression"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Publishes metric data points to Amazon CloudWatch. CloudWatch associates the
// data points with the specified metric. If the specified metric does not exist,
// CloudWatch creates the metric. When CloudWatch creates a metric, it can take up
// to fifteen minutes for the metric to appear in calls to [ListMetrics].
//
// You can publish either individual data points in the Value field, or arrays of
// values and the number of times each value occurred during the period by using
// the Values and Counts fields in the MetricData structure. Using the Values and
// Counts method enables you to publish up to 150 values per metric with one
// PutMetricData request, and supports retrieving percentile statistics on this
// data.
//
// Each PutMetricData request is limited to 1 MB in size for HTTP POST requests.
// You can send a payload compressed by gzip. Each request is also limited to no
// more than 1000 different metrics.
//
// Although the Value parameter accepts numbers of type Double , CloudWatch rejects
// values that are either too small or too large. Values must be in the range of
// -2^360 to 2^360. In addition, special values (for example, NaN, +Infinity,
// -Infinity) are not supported.
//
// You can use up to 30 dimensions per metric to further clarify what data the
// metric collects. Each dimension consists of a Name and Value pair. For more
// information about specifying dimensions, see [Publishing Metrics]in the Amazon CloudWatch User
// Guide.
//
// You specify the time stamp to be associated with each data point. You can
// specify time stamps that are as much as two weeks before the current date, and
// as much as 2 hours after the current day and time.
//
// Data points with time stamps from 24 hours ago or longer can take at least 48
// hours to become available for [GetMetricData]or [GetMetricStatistics] from the time they are submitted. Data points
// with time stamps between 3 and 24 hours ago can take as much as 2 hours to
// become available for for [GetMetricData]or [GetMetricStatistics].
//
// CloudWatch needs raw data points to calculate percentile statistics. If you
// publish data using a statistic set instead, you can only retrieve percentile
// statistics for this data if one of the following conditions is true:
//
//   - The SampleCount value of the statistic set is 1 and Min , Max , and Sum are
//     all equal.
//
//   - The Min and Max are equal, and Sum is equal to Min multiplied by SampleCount
//     .
//
// [GetMetricData]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_GetMetricData.html
// [GetMetricStatistics]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_GetMetricStatistics.html
// [ListMetrics]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_ListMetrics.html
// [Publishing Metrics]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/publishingMetrics.html
func (c *Client) PutMetricData(ctx context.Context, params *PutMetricDataInput, optFns ...func(*Options)) (*PutMetricDataOutput, error) {
	if params == nil {
		params = &PutMetricDataInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutMetricData", params, optFns, c.addOperationPutMetricDataMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutMetricDataOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutMetricDataInput struct {

	// The data for the metric. The array can include no more than 1000 metrics per
	// call.
	//
	// This member is required.
	MetricData []types.MetricDatum

	// The namespace for the metric data. You can use ASCII characters for the
	// namespace, except for control characters which are not supported.
	//
	// To avoid conflicts with Amazon Web Services service namespaces, you should not
	// specify a namespace that begins with AWS/
	//
	// This member is required.
	Namespace *string

	noSmithyDocumentSerde
}

type PutMetricDataOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutMetricDataMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpPutMetricData{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpPutMetricData{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PutMetricData"); err != nil {
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
	if err = addOperationPutMetricDataRequestCompressionMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpPutMetricDataValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutMetricData(options.Region), middleware.Before); err != nil {
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

func addOperationPutMetricDataRequestCompressionMiddleware(stack *middleware.Stack, options Options) error {
	return smithyrequestcompression.AddRequestCompression(stack, options.DisableRequestCompression, options.RequestMinCompressSizeBytes,
		[]string{
			"gzip",
		})
}

func newServiceMetadataMiddleware_opPutMetricData(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PutMetricData",
	}
}
