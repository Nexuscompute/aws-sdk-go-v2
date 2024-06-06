// Code generated by smithy-go-codegen DO NOT EDIT.

package transfer

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Allows you to delete the access specified in the ServerID and ExternalID
// parameters.
func (c *Client) DeleteAccess(ctx context.Context, params *DeleteAccessInput, optFns ...func(*Options)) (*DeleteAccessOutput, error) {
	if params == nil {
		params = &DeleteAccessInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteAccess", params, optFns, c.addOperationDeleteAccessMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteAccessOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteAccessInput struct {

	// A unique identifier that is required to identify specific groups within your
	// directory. The users of the group that you associate have access to your Amazon
	// S3 or Amazon EFS resources over the enabled protocols using Transfer Family. If
	// you know the group name, you can view the SID values by running the following
	// command using Windows PowerShell.
	//
	//     Get-ADGroup -Filter {samAccountName -like "YourGroupName*"} -Properties * |
	//     Select SamAccountName,ObjectSid
	//
	// In that command, replace YourGroupName with the name of your Active Directory
	// group.
	//
	// The regular expression used to validate this parameter is a string of
	// characters consisting of uppercase and lowercase alphanumeric characters with no
	// spaces. You can also include underscores or any of the following characters:
	// =,.@:/-
	//
	// This member is required.
	ExternalId *string

	// A system-assigned unique identifier for a server that has this user assigned.
	//
	// This member is required.
	ServerId *string

	noSmithyDocumentSerde
}

type DeleteAccessOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteAccessMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteAccess{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteAccess{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeleteAccess"); err != nil {
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
	if err = addOpDeleteAccessValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteAccess(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteAccess(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeleteAccess",
	}
}
