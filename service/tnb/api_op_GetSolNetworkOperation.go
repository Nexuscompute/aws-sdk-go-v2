// Code generated by smithy-go-codegen DO NOT EDIT.

package tnb

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/tnb/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets the details of a network operation, including the tasks involved in the
// network operation and the status of the tasks.
//
// A network operation is any operation that is done to your network, such as
// network instance instantiation or termination.
func (c *Client) GetSolNetworkOperation(ctx context.Context, params *GetSolNetworkOperationInput, optFns ...func(*Options)) (*GetSolNetworkOperationOutput, error) {
	if params == nil {
		params = &GetSolNetworkOperationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetSolNetworkOperation", params, optFns, c.addOperationGetSolNetworkOperationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetSolNetworkOperationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetSolNetworkOperationInput struct {

	// The identifier of the network operation.
	//
	// This member is required.
	NsLcmOpOccId *string

	noSmithyDocumentSerde
}

type GetSolNetworkOperationOutput struct {

	// Network operation ARN.
	//
	// This member is required.
	Arn *string

	// Error related to this specific network operation occurrence.
	Error *types.ProblemDetails

	// ID of this network operation occurrence.
	Id *string

	// Type of the operation represented by this occurrence.
	LcmOperationType types.LcmOperationType

	// Metadata of this network operation occurrence.
	Metadata *types.GetSolNetworkOperationMetadata

	// ID of the network operation instance.
	NsInstanceId *string

	// The state of the network operation.
	OperationState types.NsLcmOperationState

	// A tag is a label that you assign to an Amazon Web Services resource. Each tag
	// consists of a key and an optional value. You can use tags to search and filter
	// your resources or track your Amazon Web Services costs.
	Tags map[string]string

	// All tasks associated with this operation occurrence.
	Tasks []types.GetSolNetworkOperationTaskDetails

	// Type of the update. Only present if the network operation lcmOperationType is
	// UPDATE .
	UpdateType types.UpdateSolNetworkType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetSolNetworkOperationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetSolNetworkOperation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetSolNetworkOperation{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetSolNetworkOperation"); err != nil {
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
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = addOpGetSolNetworkOperationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetSolNetworkOperation(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetSolNetworkOperation(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetSolNetworkOperation",
	}
}
