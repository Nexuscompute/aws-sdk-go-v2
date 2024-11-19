// Code generated by smithy-go-codegen DO NOT EDIT.

package keyspaces

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/keyspaces/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

//	The CreateType operation creates a new user-defined type in the specified
//
// keyspace.
//
// To configure the required permissions, see [Permissions to create a UDT] in the Amazon Keyspaces Developer
// Guide.
//
// For more information, see [User-defined types (UDTs)] in the Amazon Keyspaces Developer Guide.
//
// [User-defined types (UDTs)]: https://docs.aws.amazon.com/keyspaces/latest/devguide/udts.html
// [Permissions to create a UDT]: https://docs.aws.amazon.com/keyspaces/latest/devguide/configure-udt-permissions.html#udt-permissions-create
func (c *Client) CreateType(ctx context.Context, params *CreateTypeInput, optFns ...func(*Options)) (*CreateTypeOutput, error) {
	if params == nil {
		params = &CreateTypeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateType", params, optFns, c.addOperationCreateTypeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateTypeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateTypeInput struct {

	//  The field definitions, consisting of names and types, that define this type.
	//
	// This member is required.
	FieldDefinitions []types.FieldDefinition

	//  The name of the keyspace.
	//
	// This member is required.
	KeyspaceName *string

	//  The name of the user-defined type.
	//
	// UDT names must contain 48 characters or less, must begin with an alphabetic
	// character, and can only contain alpha-numeric characters and underscores. Amazon
	// Keyspaces converts upper case characters automatically into lower case
	// characters.
	//
	// Alternatively, you can declare a UDT name in double quotes. When declaring a
	// UDT name inside double quotes, Amazon Keyspaces preserves upper casing and
	// allows special characters.
	//
	// You can also use double quotes as part of the name when you create the UDT, but
	// you must escape each double quote character with an additional double quote
	// character.
	//
	// This member is required.
	TypeName *string

	noSmithyDocumentSerde
}

type CreateTypeOutput struct {

	//  The unique identifier of the keyspace that contains the new type in the format
	// of an Amazon Resource Name (ARN).
	//
	// This member is required.
	KeyspaceArn *string

	//  The formatted name of the user-defined type that was created. Note that Amazon
	// Keyspaces requires the formatted name of the type for other operations, for
	// example GetType .
	//
	// This member is required.
	TypeName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateTypeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpCreateType{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpCreateType{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateType"); err != nil {
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
	if err = addOpCreateTypeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateType(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateType(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateType",
	}
}
