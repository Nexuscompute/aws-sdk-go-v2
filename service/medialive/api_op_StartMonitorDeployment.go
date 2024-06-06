// Code generated by smithy-go-codegen DO NOT EDIT.

package medialive

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/medialive/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Initiates a deployment to deploy the latest monitor of the specified signal map.
func (c *Client) StartMonitorDeployment(ctx context.Context, params *StartMonitorDeploymentInput, optFns ...func(*Options)) (*StartMonitorDeploymentOutput, error) {
	if params == nil {
		params = &StartMonitorDeploymentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartMonitorDeployment", params, optFns, c.addOperationStartMonitorDeploymentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartMonitorDeploymentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Placeholder documentation for StartMonitorDeploymentRequest
type StartMonitorDeploymentInput struct {

	// A signal map's identifier. Can be either be its id or current name.
	//
	// This member is required.
	Identifier *string

	// Placeholder documentation for __boolean
	DryRun *bool

	noSmithyDocumentSerde
}

// Placeholder documentation for StartMonitorDeploymentResponse
type StartMonitorDeploymentOutput struct {

	// A signal map's ARN (Amazon Resource Name)
	Arn *string

	// Placeholder documentation for __listOf__stringMin7Max11PatternAws097
	CloudWatchAlarmTemplateGroupIds []string

	// Placeholder documentation for __timestampIso8601
	CreatedAt *time.Time

	// A resource's optional description.
	Description *string

	// A top-level supported AWS resource ARN to discovery a signal map from.
	DiscoveryEntryPointArn *string

	// Error message associated with a failed creation or failed update attempt of a
	// signal map.
	ErrorMessage *string

	// Placeholder documentation for __listOf__stringMin7Max11PatternAws097
	EventBridgeRuleTemplateGroupIds []string

	// A map representing an incomplete AWS media workflow as a graph.
	FailedMediaResourceMap map[string]types.MediaResource

	// A signal map's id.
	Id *string

	// Placeholder documentation for __timestampIso8601
	LastDiscoveredAt *time.Time

	// Represents the latest successful monitor deployment of a signal map.
	LastSuccessfulMonitorDeployment *types.SuccessfulMonitorDeployment

	// A map representing an AWS media workflow as a graph.
	MediaResourceMap map[string]types.MediaResource

	// Placeholder documentation for __timestampIso8601
	ModifiedAt *time.Time

	// If true, there are pending monitor changes for this signal map that can be
	// deployed.
	MonitorChangesPendingDeployment *bool

	// Represents the latest monitor deployment of a signal map.
	MonitorDeployment *types.MonitorDeployment

	// A resource's name. Names must be unique within the scope of a resource type in
	// a specific region.
	Name *string

	// A signal map's current status which is dependent on its lifecycle actions or
	// associated jobs.
	Status types.SignalMapStatus

	// Represents the tags associated with a resource.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartMonitorDeploymentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartMonitorDeployment{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartMonitorDeployment{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "StartMonitorDeployment"); err != nil {
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
	if err = addOpStartMonitorDeploymentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartMonitorDeployment(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartMonitorDeployment(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "StartMonitorDeployment",
	}
}
