// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the read and write permissions for an analysis.
func (c *Client) UpdateAnalysisPermissions(ctx context.Context, params *UpdateAnalysisPermissionsInput, optFns ...func(*Options)) (*UpdateAnalysisPermissionsOutput, error) {
	if params == nil {
		params = &UpdateAnalysisPermissionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateAnalysisPermissions", params, optFns, c.addOperationUpdateAnalysisPermissionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateAnalysisPermissionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateAnalysisPermissionsInput struct {

	// The ID of the analysis whose permissions you're updating. The ID is part of the
	// analysis URL.
	//
	// This member is required.
	AnalysisId *string

	// The ID of the Amazon Web Services account that contains the analysis whose
	// permissions you're updating. You must be using the Amazon Web Services account
	// that the analysis is in.
	//
	// This member is required.
	AwsAccountId *string

	// A structure that describes the permissions to add and the principal to add them
	// to.
	GrantPermissions []types.ResourcePermission

	// A structure that describes the permissions to remove and the principal to
	// remove them from.
	RevokePermissions []types.ResourcePermission

	noSmithyDocumentSerde
}

type UpdateAnalysisPermissionsOutput struct {

	// The Amazon Resource Name (ARN) of the analysis that you updated.
	AnalysisArn *string

	// The ID of the analysis that you updated permissions for.
	AnalysisId *string

	// A structure that describes the principals and the resource-level permissions on
	// an analysis.
	Permissions []types.ResourcePermission

	// The Amazon Web Services request ID for this operation.
	RequestId *string

	// The HTTP status of the request.
	Status int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateAnalysisPermissionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateAnalysisPermissions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateAnalysisPermissions{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateAnalysisPermissions"); err != nil {
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
	if err = addOpUpdateAnalysisPermissionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateAnalysisPermissions(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateAnalysisPermissions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateAnalysisPermissions",
	}
}
