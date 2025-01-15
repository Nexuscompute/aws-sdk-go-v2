// Code generated by smithy-go-codegen DO NOT EDIT.

package partnercentralselling

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/partnercentralselling/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Use this action to retrieve a specific snapshot record.
func (c *Client) GetResourceSnapshot(ctx context.Context, params *GetResourceSnapshotInput, optFns ...func(*Options)) (*GetResourceSnapshotOutput, error) {
	if params == nil {
		params = &GetResourceSnapshotInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetResourceSnapshot", params, optFns, c.addOperationGetResourceSnapshotMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetResourceSnapshotOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetResourceSnapshotInput struct {

	// Specifies the catalog related to the request. Valid values are:
	//
	//   - AWS: Retrieves the snapshot from the production AWS environment.
	//
	//   - Sandbox: Retrieves the snapshot from a sandbox environment used for testing
	//   or development purposes.
	//
	// This member is required.
	Catalog *string

	// The unique identifier of the engagement associated with the snapshot. This
	// field links the snapshot to a specific engagement context.
	//
	// This member is required.
	EngagementIdentifier *string

	// The unique identifier of the specific resource that was snapshotted. The format
	// and constraints of this identifier depend on the ResourceType specified. For
	// Opportunity type, it will be an opportunity ID
	//
	// This member is required.
	ResourceIdentifier *string

	// he name of the template that defines the schema for the snapshot. This template
	// determines which subset of the resource data is included in the snapshot and
	// must correspond to an existing and valid template for the specified ResourceType
	// .
	//
	// This member is required.
	ResourceSnapshotTemplateIdentifier *string

	// Specifies the type of resource that was snapshotted. This field determines the
	// structure and content of the snapshot payload. Valid value includes: Opportunity
	// : For opportunity-related data.
	//
	// This member is required.
	ResourceType types.ResourceType

	// Specifies which revision of the snapshot to retrieve. If omitted returns the
	// latest revision.
	Revision *int32

	noSmithyDocumentSerde
}

type GetResourceSnapshotOutput struct {

	// The catalog in which the snapshot was created. Matches the Catalog specified in
	// the request.
	//
	// This member is required.
	Catalog *string

	// The Amazon Resource Name (ARN) that uniquely identifies the resource snapshot.
	Arn *string

	// The timestamp when the snapshot was created, in ISO 8601 format (e.g.,
	// "2023-06-01T14:30:00Z"). This allows for precise tracking of when the snapshot
	// was taken.
	CreatedAt *time.Time

	// The AWS account ID of the principal (user or role) who created the snapshot.
	// This helps in tracking the origin of the snapshot.
	CreatedBy *string

	// The identifier of the engagement associated with this snapshot. Matches the
	// EngagementIdentifier specified in the request.
	EngagementId *string

	//  Represents the payload of a resource snapshot. This structure is designed to
	// accommodate different types of resource snapshots, currently supporting
	// opportunity summaries.
	Payload types.ResourceSnapshotPayload

	// The identifier of the specific resource that was snapshotted. Matches the
	// ResourceIdentifier specified in the request.
	ResourceId *string

	// The name of the view used for this snapshot. This is the same as the template
	// name.
	ResourceSnapshotTemplateName *string

	// The type of the resource that was snapshotted. Matches the ResourceType
	// specified in the request.
	ResourceType types.ResourceType

	// The revision number of this snapshot. This is a positive integer that is
	// sequential and unique within the context of a resource view.
	Revision *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetResourceSnapshotMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpGetResourceSnapshot{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpGetResourceSnapshot{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetResourceSnapshot"); err != nil {
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
	if err = addOpGetResourceSnapshotValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetResourceSnapshot(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetResourceSnapshot(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetResourceSnapshot",
	}
}
