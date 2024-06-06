// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesis

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/kinesis/types"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/ptr"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Provides a summarized description of the specified Kinesis data stream without
// the shard list.
//
// When invoking this API, you must use either the StreamARN or the StreamName
// parameter, or both. It is recommended that you use the StreamARN input
// parameter when you invoke this API.
//
// The information returned includes the stream name, Amazon Resource Name (ARN),
// status, record retention period, approximate creation time, monitoring,
// encryption details, and open shard count.
//
// DescribeStreamSummaryhas a limit of 20 transactions per second per account.
func (c *Client) DescribeStreamSummary(ctx context.Context, params *DescribeStreamSummaryInput, optFns ...func(*Options)) (*DescribeStreamSummaryOutput, error) {
	if params == nil {
		params = &DescribeStreamSummaryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeStreamSummary", params, optFns, c.addOperationDescribeStreamSummaryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeStreamSummaryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeStreamSummaryInput struct {

	// The ARN of the stream.
	StreamARN *string

	// The name of the stream to describe.
	StreamName *string

	noSmithyDocumentSerde
}

func (in *DescribeStreamSummaryInput) bindEndpointParams(p *EndpointParameters) {
	p.StreamARN = in.StreamARN
	p.OperationType = ptr.String("control")
}

type DescribeStreamSummaryOutput struct {

	// A StreamDescriptionSummary containing information about the stream.
	//
	// This member is required.
	StreamDescriptionSummary *types.StreamDescriptionSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeStreamSummaryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeStreamSummary{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeStreamSummary{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeStreamSummary"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeStreamSummary(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeStreamSummary(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeStreamSummary",
	}
}
