// Code generated by smithy-go-codegen DO NOT EDIT.

package deadline

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/deadline/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Gets a worker.
func (c *Client) GetWorker(ctx context.Context, params *GetWorkerInput, optFns ...func(*Options)) (*GetWorkerOutput, error) {
	if params == nil {
		params = &GetWorkerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetWorker", params, optFns, c.addOperationGetWorkerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetWorkerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetWorkerInput struct {

	// The farm ID for the worker.
	//
	// This member is required.
	FarmId *string

	// The fleet ID of the worker.
	//
	// This member is required.
	FleetId *string

	// The worker ID.
	//
	// This member is required.
	WorkerId *string

	noSmithyDocumentSerde
}

type GetWorkerOutput struct {

	// The date and time the resource was created.
	//
	// This member is required.
	CreatedAt *time.Time

	// The user or system that created this resource.
	//
	// This member is required.
	CreatedBy *string

	// The farm ID.
	//
	// This member is required.
	FarmId *string

	// The fleet ID.
	//
	// This member is required.
	FleetId *string

	// The status of the worker.
	//
	// This member is required.
	Status types.WorkerStatus

	// The worker ID.
	//
	// This member is required.
	WorkerId *string

	// The host properties for the worker.
	HostProperties *types.HostPropertiesResponse

	// The logs for the associated worker.
	Log *types.LogConfiguration

	// The date and time the resource was updated.
	UpdatedAt *time.Time

	// The user or system that updated this resource.
	UpdatedBy *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetWorkerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetWorker{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetWorker{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetWorker"); err != nil {
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
	if err = addEndpointPrefix_opGetWorkerMiddleware(stack); err != nil {
		return err
	}
	if err = addOpGetWorkerValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetWorker(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opGetWorkerMiddleware struct {
}

func (*endpointPrefix_opGetWorkerMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opGetWorkerMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleFinalize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "management." + req.URL.Host

	return next.HandleFinalize(ctx, in)
}
func addEndpointPrefix_opGetWorkerMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opGetWorkerMiddleware{}, "ResolveEndpointV2", middleware.After)
}

func newServiceMetadataMiddleware_opGetWorker(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetWorker",
	}
}
