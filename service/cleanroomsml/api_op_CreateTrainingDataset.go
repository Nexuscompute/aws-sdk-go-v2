// Code generated by smithy-go-codegen DO NOT EDIT.

package cleanroomsml

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/cleanroomsml/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Defines the information necessary to create a training dataset. In Clean Rooms
// ML, the TrainingDataset is metadata that points to a Glue table, which is read
// only during AudienceModel creation.
func (c *Client) CreateTrainingDataset(ctx context.Context, params *CreateTrainingDatasetInput, optFns ...func(*Options)) (*CreateTrainingDatasetOutput, error) {
	if params == nil {
		params = &CreateTrainingDatasetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateTrainingDataset", params, optFns, c.addOperationCreateTrainingDatasetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateTrainingDatasetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateTrainingDatasetInput struct {

	// The name of the training dataset. This name must be unique in your account and
	// region.
	//
	// This member is required.
	Name *string

	// The ARN of the IAM role that Clean Rooms ML can assume to read the data
	// referred to in the dataSource field of each dataset.
	//
	// Passing a role across AWS accounts is not allowed. If you pass a role that
	// isn't in your account, you get an AccessDeniedException error.
	//
	// This member is required.
	RoleArn *string

	// An array of information that lists the Dataset objects, which specifies the
	// dataset type and details on its location and schema. You must provide a role
	// that has read access to these tables.
	//
	// This member is required.
	TrainingData []types.Dataset

	// The description of the training dataset.
	Description *string

	// The optional metadata that you apply to the resource to help you categorize and
	// organize them. Each tag consists of a key and an optional value, both of which
	// you define.
	//
	// The following basic restrictions apply to tags:
	//
	//   - Maximum number of tags per resource - 50.
	//
	//   - For each resource, each tag key must be unique, and each tag key can have
	//   only one value.
	//
	//   - Maximum key length - 128 Unicode characters in UTF-8.
	//
	//   - Maximum value length - 256 Unicode characters in UTF-8.
	//
	//   - If your tagging schema is used across multiple services and resources,
	//   remember that other services may have restrictions on allowed characters.
	//   Generally allowed characters are: letters, numbers, and spaces representable in
	//   UTF-8, and the following characters: + - = . _ : / @.
	//
	//   - Tag keys and values are case sensitive.
	//
	//   - Do not use aws:, AWS:, or any upper or lowercase combination of such as a
	//   prefix for keys as it is reserved for AWS use. You cannot edit or delete tag
	//   keys with this prefix. Values can have this prefix. If a tag value has aws as
	//   its prefix but the key does not, then Clean Rooms ML considers it to be a user
	//   tag and will count against the limit of 50 tags. Tags with only the key prefix
	//   of aws do not count against your tags per resource limit.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateTrainingDatasetOutput struct {

	// The Amazon Resource Name (ARN) of the training dataset resource.
	//
	// This member is required.
	TrainingDatasetArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateTrainingDatasetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateTrainingDataset{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateTrainingDataset{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateTrainingDataset"); err != nil {
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
	if err = addOpCreateTrainingDatasetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateTrainingDataset(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateTrainingDataset(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateTrainingDataset",
	}
}
