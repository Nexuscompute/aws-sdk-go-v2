// Code generated by smithy-go-codegen DO NOT EDIT.

package lightsail

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates two URLs that are used to access a virtual computer’s graphical user
// interface (GUI) session. The primary URL initiates a web-based NICE DCV session
// to the virtual computer's application. The secondary URL initiates a web-based
// NICE DCV session to the virtual computer's operating session.
//
// Use StartGUISession to open the session.
func (c *Client) CreateGUISessionAccessDetails(ctx context.Context, params *CreateGUISessionAccessDetailsInput, optFns ...func(*Options)) (*CreateGUISessionAccessDetailsOutput, error) {
	if params == nil {
		params = &CreateGUISessionAccessDetailsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateGUISessionAccessDetails", params, optFns, c.addOperationCreateGUISessionAccessDetailsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateGUISessionAccessDetailsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateGUISessionAccessDetailsInput struct {

	// The resource name.
	//
	// This member is required.
	ResourceName *string

	noSmithyDocumentSerde
}

type CreateGUISessionAccessDetailsOutput struct {

	// The reason the operation failed.
	FailureReason *string

	// The percentage of completion for the operation.
	PercentageComplete *int32

	// The resource name.
	ResourceName *string

	// Returns information about the specified NICE DCV GUI session.
	Sessions []types.Session

	// The status of the operation.
	Status types.Status

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateGUISessionAccessDetailsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateGUISessionAccessDetails{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateGUISessionAccessDetails{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateGUISessionAccessDetails"); err != nil {
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
	if err = addOpCreateGUISessionAccessDetailsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateGUISessionAccessDetails(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateGUISessionAccessDetails(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateGUISessionAccessDetails",
	}
}
