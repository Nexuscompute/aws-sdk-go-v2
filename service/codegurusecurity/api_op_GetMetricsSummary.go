// Code generated by smithy-go-codegen DO NOT EDIT.

package codegurusecurity

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codegurusecurity/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns top level metrics about an account from a specified date, including
// number of open findings, the categories with most findings, the scans with most
// open findings, and scans with most open critical findings.
func (c *Client) GetMetricsSummary(ctx context.Context, params *GetMetricsSummaryInput, optFns ...func(*Options)) (*GetMetricsSummaryOutput, error) {
	if params == nil {
		params = &GetMetricsSummaryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetMetricsSummary", params, optFns, c.addOperationGetMetricsSummaryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetMetricsSummaryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetMetricsSummaryInput struct {

	// The date you want to retrieve summary metrics from, rounded to the nearest day.
	// The date must be within the past two years since metrics data is only stored for
	// two years. If a date outside of this range is passed, the response will be
	// empty.
	//
	// This member is required.
	Date *time.Time

	noSmithyDocumentSerde
}

type GetMetricsSummaryOutput struct {

	// The summary metrics from the specified date.
	MetricsSummary *types.MetricsSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetMetricsSummaryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetMetricsSummary{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetMetricsSummary{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpGetMetricsSummaryValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetMetricsSummary(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
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
	return nil
}

func newServiceMetadataMiddleware_opGetMetricsSummary(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codeguru-security",
		OperationName: "GetMetricsSummary",
	}
}