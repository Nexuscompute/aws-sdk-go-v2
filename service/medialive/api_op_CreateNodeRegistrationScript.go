// Code generated by smithy-go-codegen DO NOT EDIT.

package medialive

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/medialive/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Create the Register Node script for all the nodes intended for a specific
// Cluster. You will then run the script on each hardware unit that is intended for
// that Cluster. The script creates a Node in the specified Cluster. It then binds
// the Node to this hardware unit, and activates the node hardware for use with
// MediaLive Anywhere.
func (c *Client) CreateNodeRegistrationScript(ctx context.Context, params *CreateNodeRegistrationScriptInput, optFns ...func(*Options)) (*CreateNodeRegistrationScriptOutput, error) {
	if params == nil {
		params = &CreateNodeRegistrationScriptInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateNodeRegistrationScript", params, optFns, c.addOperationCreateNodeRegistrationScriptMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateNodeRegistrationScriptOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to create a new node registration script.
type CreateNodeRegistrationScriptInput struct {

	// The ID of the cluster
	//
	// This member is required.
	ClusterId *string

	// If you're generating a re-registration script for an already existing node,
	// this is where you provide the id.
	Id *string

	// Specify a pattern for MediaLive Anywhere to use to assign a name to each Node
	// in the Cluster. The pattern can include the variables $hn (hostname of the node
	// hardware) and $ts for the date and time that the Node is created, in UTC (for
	// example, 2024-08-20T23:35:12Z).
	Name *string

	// Documentation update needed
	NodeInterfaceMappings []types.NodeInterfaceMapping

	// An ID that you assign to a create request. This ID ensures idempotency when
	// creating resources.
	RequestId *string

	// The initial role of the Node in the Cluster. ACTIVE means the Node is available
	// for encoding. BACKUP means the Node is a redundant Node and might get used if an
	// ACTIVE Node fails.
	Role types.NodeRole

	noSmithyDocumentSerde
}

// Placeholder documentation for CreateNodeRegistrationScriptResponse
type CreateNodeRegistrationScriptOutput struct {

	// A script that can be run on a Bring Your Own Device Elemental Anywhere system
	// to create a node in a cluster.
	NodeRegistrationScript *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateNodeRegistrationScriptMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateNodeRegistrationScript{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateNodeRegistrationScript{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateNodeRegistrationScript"); err != nil {
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
	if err = addIdempotencyToken_opCreateNodeRegistrationScriptMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateNodeRegistrationScriptValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateNodeRegistrationScript(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateNodeRegistrationScript struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateNodeRegistrationScript) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateNodeRegistrationScript) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateNodeRegistrationScriptInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateNodeRegistrationScriptInput ")
	}

	if input.RequestId == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.RequestId = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateNodeRegistrationScriptMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateNodeRegistrationScript{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateNodeRegistrationScript(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateNodeRegistrationScript",
	}
}
