// Code generated by smithy-go-codegen DO NOT EDIT.

package customerprofiles

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/customerprofiles/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns information about the specified event stream in a specific domain.
func (c *Client) GetEventStream(ctx context.Context, params *GetEventStreamInput, optFns ...func(*Options)) (*GetEventStreamOutput, error) {
	if params == nil {
		params = &GetEventStreamInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetEventStream", params, optFns, c.addOperationGetEventStreamMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetEventStreamOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetEventStreamInput struct {

	// The unique name of the domain.
	//
	// This member is required.
	DomainName *string

	// The name of the event stream provided during create operations.
	//
	// This member is required.
	EventStreamName *string

	noSmithyDocumentSerde
}

type GetEventStreamOutput struct {

	// The timestamp of when the export was created.
	//
	// This member is required.
	CreatedAt *time.Time

	// Details regarding the Kinesis stream.
	//
	// This member is required.
	DestinationDetails *types.EventStreamDestinationDetails

	// The unique name of the domain.
	//
	// This member is required.
	DomainName *string

	// A unique identifier for the event stream.
	//
	// This member is required.
	EventStreamArn *string

	// The operational state of destination stream for export.
	//
	// This member is required.
	State types.EventStreamState

	// The timestamp when the State changed to STOPPED .
	StoppedSince *time.Time

	// The tags used to organize, track, or control access for this resource.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetEventStreamMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetEventStream{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetEventStream{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetEventStream"); err != nil {
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
	if err = addOpGetEventStreamValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetEventStream(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetEventStream(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetEventStream",
	}
}
