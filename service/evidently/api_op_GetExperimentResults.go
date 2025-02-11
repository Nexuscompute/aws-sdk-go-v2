// Code generated by smithy-go-codegen DO NOT EDIT.

package evidently

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/evidently/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Retrieves the results of a running or completed experiment. No results are
// available until there have been 100 events for each variation and at least 10
// minutes have passed since the start of the experiment. To increase the
// statistical power, Evidently performs an additional offline p-value analysis at
// the end of the experiment. Offline p-value analysis can detect statistical
// significance in some cases where the anytime p-values used during the experiment
// do not find statistical significance.
//
// Experiment results are available up to 63 days after the start of the
// experiment. They are not available after that because of CloudWatch data
// retention policies.
func (c *Client) GetExperimentResults(ctx context.Context, params *GetExperimentResultsInput, optFns ...func(*Options)) (*GetExperimentResultsOutput, error) {
	if params == nil {
		params = &GetExperimentResultsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetExperimentResults", params, optFns, c.addOperationGetExperimentResultsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetExperimentResultsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetExperimentResultsInput struct {

	// The name of the experiment to retrieve the results of.
	//
	// This member is required.
	Experiment *string

	// The names of the experiment metrics that you want to see the results of.
	//
	// This member is required.
	MetricNames []string

	// The name or ARN of the project that contains the experiment that you want to
	// see the results of.
	//
	// This member is required.
	Project *string

	// The names of the experiment treatments that you want to see the results for.
	//
	// This member is required.
	TreatmentNames []string

	// The statistic used to calculate experiment results. Currently the only valid
	// value is mean , which uses the mean of the collected values as the statistic.
	BaseStat types.ExperimentBaseStat

	// The date and time that the experiment ended, if it is completed. This must be
	// no longer than 30 days after the experiment start time.
	EndTime *time.Time

	// In seconds, the amount of time to aggregate results together.
	Period *int64

	// The names of the report types that you want to see. Currently, BayesianInference
	// is the only valid value.
	ReportNames []types.ExperimentReportName

	// The statistics that you want to see in the returned results.
	//
	//   - PValue specifies to use p-values for the results. A p-value is used in
	//   hypothesis testing to measure how often you are willing to make a mistake in
	//   rejecting the null hypothesis. A general practice is to reject the null
	//   hypothesis and declare that the results are statistically significant when the
	//   p-value is less than 0.05.
	//
	//   - ConfidenceInterval specifies a confidence interval for the results. The
	//   confidence interval represents the range of values for the chosen metric that is
	//   likely to contain the true difference between the baseStat of a variation and
	//   the baseline. Evidently returns the 95% confidence interval.
	//
	//   - TreatmentEffect is the difference in the statistic specified by the baseStat
	//   parameter between each variation and the default variation.
	//
	//   - BaseStat returns the statistical values collected for the metric for each
	//   variation. The statistic uses the same statistic specified in the baseStat
	//   parameter. Therefore, if baseStat is mean , this returns the mean of the
	//   values collected for each variation.
	ResultStats []types.ExperimentResultRequestType

	// The date and time that the experiment started.
	StartTime *time.Time

	noSmithyDocumentSerde
}

type GetExperimentResultsOutput struct {

	// If the experiment doesn't yet have enough events to provide valid results, this
	// field is returned with the message Not enough events to generate results . If
	// there are enough events to provide valid results, this field is not returned.
	Details *string

	// An array of structures that include the reports that you requested.
	Reports []types.ExperimentReport

	// An array of structures that include experiment results including metric names
	// and values.
	ResultsData []types.ExperimentResultsData

	// The timestamps of each result returned.
	Timestamps []time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetExperimentResultsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetExperimentResults{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetExperimentResults{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetExperimentResults"); err != nil {
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
	if err = addOpGetExperimentResultsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetExperimentResults(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetExperimentResults(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetExperimentResults",
	}
}
