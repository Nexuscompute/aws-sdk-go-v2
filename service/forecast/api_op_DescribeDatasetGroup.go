// Code generated by smithy-go-codegen DO NOT EDIT.

package forecast

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/forecast/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Describes a dataset group created using the [CreateDatasetGroup] operation.
//
// In addition to listing the parameters provided in the CreateDatasetGroup
// request, this operation includes the following properties:
//
//   - DatasetArns - The datasets belonging to the group.
//
//   - CreationTime
//
//   - LastModificationTime
//
//   - Status
//
// [CreateDatasetGroup]: https://docs.aws.amazon.com/forecast/latest/dg/API_CreateDatasetGroup.html
func (c *Client) DescribeDatasetGroup(ctx context.Context, params *DescribeDatasetGroupInput, optFns ...func(*Options)) (*DescribeDatasetGroupOutput, error) {
	if params == nil {
		params = &DescribeDatasetGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeDatasetGroup", params, optFns, c.addOperationDescribeDatasetGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeDatasetGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeDatasetGroupInput struct {

	// The Amazon Resource Name (ARN) of the dataset group.
	//
	// This member is required.
	DatasetGroupArn *string

	noSmithyDocumentSerde
}

type DescribeDatasetGroupOutput struct {

	// When the dataset group was created.
	CreationTime *time.Time

	// An array of Amazon Resource Names (ARNs) of the datasets contained in the
	// dataset group.
	DatasetArns []string

	// The ARN of the dataset group.
	DatasetGroupArn *string

	// The name of the dataset group.
	DatasetGroupName *string

	// The domain associated with the dataset group.
	Domain types.Domain

	// When the dataset group was created or last updated from a call to the [UpdateDatasetGroup]
	// operation. While the dataset group is being updated, LastModificationTime is
	// the current time of the DescribeDatasetGroup call.
	//
	// [UpdateDatasetGroup]: https://docs.aws.amazon.com/forecast/latest/dg/API_UpdateDatasetGroup.html
	LastModificationTime *time.Time

	// The status of the dataset group. States include:
	//
	//   - ACTIVE
	//
	//   - CREATE_PENDING , CREATE_IN_PROGRESS , CREATE_FAILED
	//
	//   - DELETE_PENDING , DELETE_IN_PROGRESS , DELETE_FAILED
	//
	//   - UPDATE_PENDING , UPDATE_IN_PROGRESS , UPDATE_FAILED
	//
	// The UPDATE states apply when you call the [UpdateDatasetGroup] operation.
	//
	// The Status of the dataset group must be ACTIVE before you can use the dataset
	// group to create a predictor.
	//
	// [UpdateDatasetGroup]: https://docs.aws.amazon.com/forecast/latest/dg/API_UpdateDatasetGroup.html
	Status *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeDatasetGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeDatasetGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeDatasetGroup{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeDatasetGroup"); err != nil {
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
	if err = addOpDescribeDatasetGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDatasetGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeDatasetGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeDatasetGroup",
	}
}
