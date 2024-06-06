// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Copies the specified option group.
func (c *Client) CopyOptionGroup(ctx context.Context, params *CopyOptionGroupInput, optFns ...func(*Options)) (*CopyOptionGroupOutput, error) {
	if params == nil {
		params = &CopyOptionGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CopyOptionGroup", params, optFns, c.addOperationCopyOptionGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CopyOptionGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CopyOptionGroupInput struct {

	// The identifier for the source option group.
	//
	// Constraints:
	//
	//   - Must specify a valid option group.
	//
	// This member is required.
	SourceOptionGroupIdentifier *string

	// The description for the copied option group.
	//
	// This member is required.
	TargetOptionGroupDescription *string

	// The identifier for the copied option group.
	//
	// Constraints:
	//
	//   - Can't be null, empty, or blank
	//
	//   - Must contain from 1 to 255 letters, numbers, or hyphens
	//
	//   - First character must be a letter
	//
	//   - Can't end with a hyphen or contain two consecutive hyphens
	//
	// Example: my-option-group
	//
	// This member is required.
	TargetOptionGroupIdentifier *string

	// A list of tags. For more information, see [Tagging Amazon RDS Resources] in the Amazon RDS User Guide.
	//
	// [Tagging Amazon RDS Resources]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Tagging.html
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CopyOptionGroupOutput struct {

	//
	OptionGroup *types.OptionGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCopyOptionGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCopyOptionGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCopyOptionGroup{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CopyOptionGroup"); err != nil {
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
	if err = addOpCopyOptionGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCopyOptionGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCopyOptionGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CopyOptionGroup",
	}
}
