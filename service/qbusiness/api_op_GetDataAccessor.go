// Code generated by smithy-go-codegen DO NOT EDIT.

package qbusiness

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/qbusiness/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Retrieves information about a specified data accessor. This operation returns
// details about the data accessor, including its display name, unique identifier,
// Amazon Resource Name (ARN), the associated Amazon Q Business application and IAM
// Identity Center application, the IAM role for the ISV, the action
// configurations, and the timestamps for when the data accessor was created and
// last updated.
func (c *Client) GetDataAccessor(ctx context.Context, params *GetDataAccessorInput, optFns ...func(*Options)) (*GetDataAccessorOutput, error) {
	if params == nil {
		params = &GetDataAccessorInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetDataAccessor", params, optFns, c.addOperationGetDataAccessorMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetDataAccessorOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetDataAccessorInput struct {

	// The unique identifier of the Amazon Q Business application.
	//
	// This member is required.
	ApplicationId *string

	// The unique identifier of the data accessor to retrieve.
	//
	// This member is required.
	DataAccessorId *string

	noSmithyDocumentSerde
}

type GetDataAccessorOutput struct {

	// The list of action configurations specifying the allowed actions and any
	// associated filters.
	ActionConfigurations []types.ActionConfiguration

	// The unique identifier of the Amazon Q Business application associated with this
	// data accessor.
	ApplicationId *string

	// The timestamp when the data accessor was created.
	CreatedAt *time.Time

	// The Amazon Resource Name (ARN) of the data accessor.
	DataAccessorArn *string

	// The unique identifier of the data accessor.
	DataAccessorId *string

	// The friendly name of the data accessor.
	DisplayName *string

	// The Amazon Resource Name (ARN) of the IAM Identity Center application
	// associated with this data accessor.
	IdcApplicationArn *string

	// The Amazon Resource Name (ARN) of the IAM role for the ISV associated with this
	// data accessor.
	Principal *string

	// The timestamp when the data accessor was last updated.
	UpdatedAt *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetDataAccessorMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetDataAccessor{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetDataAccessor{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetDataAccessor"); err != nil {
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
	if err = addOpGetDataAccessorValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetDataAccessor(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetDataAccessor(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetDataAccessor",
	}
}
