// Code generated by smithy-go-codegen DO NOT EDIT.

package transfer

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Tests whether your SFTP connector is set up successfully. We highly recommend
// that you call this operation to test your ability to transfer files between
// local Amazon Web Services storage and a trading partner's SFTP server.
func (c *Client) TestConnection(ctx context.Context, params *TestConnectionInput, optFns ...func(*Options)) (*TestConnectionOutput, error) {
	if params == nil {
		params = &TestConnectionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "TestConnection", params, optFns, c.addOperationTestConnectionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*TestConnectionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type TestConnectionInput struct {

	// The unique identifier for the connector.
	//
	// This member is required.
	ConnectorId *string

	noSmithyDocumentSerde
}

type TestConnectionOutput struct {

	// Returns the identifier of the connector object that you are testing.
	ConnectorId *string

	// Returns OK for successful test, or ERROR if the test fails.
	Status *string

	// Returns Connection succeeded if the test is successful. Or, returns a
	// descriptive error message if the test fails. The following list provides
	// troubleshooting details, depending on the error message that you receive.
	//
	//   - Verify that your secret name aligns with the one in Transfer Role
	//   permissions.
	//
	//   - Verify the server URL in the connector configuration , and verify that the
	//   login credentials work successfully outside of the connector.
	//
	//   - Verify that the secret exists and is formatted correctly.
	//
	//   - Verify that the trusted host key in the connector configuration matches the
	//   ssh-keyscan output.
	StatusMessage *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationTestConnectionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpTestConnection{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpTestConnection{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "TestConnection"); err != nil {
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
	if err = addOpTestConnectionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opTestConnection(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opTestConnection(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "TestConnection",
	}
}
