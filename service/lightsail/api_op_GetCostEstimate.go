// Code generated by smithy-go-codegen DO NOT EDIT.

package lightsail

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Retrieves information about the cost estimate for a specified resource. A cost
// estimate will not generate for a resource that has been deleted.
func (c *Client) GetCostEstimate(ctx context.Context, params *GetCostEstimateInput, optFns ...func(*Options)) (*GetCostEstimateOutput, error) {
	if params == nil {
		params = &GetCostEstimateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetCostEstimate", params, optFns, c.addOperationGetCostEstimateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetCostEstimateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetCostEstimateInput struct {

	// The cost estimate end time.
	//
	// Constraints:
	//
	//   - Specified in Coordinated Universal Time (UTC).
	//
	//   - Specified in the Unix time format.
	//
	// For example, if you want to use an end time of October 1, 2018, at 9 PM UTC,
	//   specify 1538427600 as the end time.
	//
	// You can convert a human-friendly time to Unix time format using a converter
	// like [Epoch converter].
	//
	// [Epoch converter]: https://www.epochconverter.com/
	//
	// This member is required.
	EndTime *time.Time

	// The resource name.
	//
	// This member is required.
	ResourceName *string

	// The cost estimate start time.
	//
	// Constraints:
	//
	//   - Specified in Coordinated Universal Time (UTC).
	//
	//   - Specified in the Unix time format.
	//
	// For example, if you want to use a start time of October 1, 2018, at 8 PM UTC,
	//   specify 1538424000 as the start time.
	//
	// You can convert a human-friendly time to Unix time format using a converter
	// like [Epoch converter].
	//
	// [Epoch converter]: https://www.epochconverter.com/
	//
	// This member is required.
	StartTime *time.Time

	noSmithyDocumentSerde
}

type GetCostEstimateOutput struct {

	// Returns the estimate's forecasted cost or usage.
	ResourcesBudgetEstimate []types.ResourceBudgetEstimate

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetCostEstimateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetCostEstimate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetCostEstimate{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetCostEstimate"); err != nil {
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
	if err = addOpGetCostEstimateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetCostEstimate(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetCostEstimate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetCostEstimate",
	}
}
