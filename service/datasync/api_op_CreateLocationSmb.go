// Code generated by smithy-go-codegen DO NOT EDIT.

package datasync

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/datasync/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a transfer location for a Server Message Block (SMB) file server.
// DataSync can use this location as a source or destination for transferring data.
//
// Before you begin, make sure that you understand how DataSync [accesses SMB file servers].
//
// [accesses SMB file servers]: https://docs.aws.amazon.com/datasync/latest/userguide/create-smb-location.html#configuring-smb
func (c *Client) CreateLocationSmb(ctx context.Context, params *CreateLocationSmbInput, optFns ...func(*Options)) (*CreateLocationSmbOutput, error) {
	if params == nil {
		params = &CreateLocationSmbInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateLocationSmb", params, optFns, c.addOperationCreateLocationSmbMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateLocationSmbOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// CreateLocationSmbRequest
type CreateLocationSmbInput struct {

	// Specifies the DataSync agent (or agents) which you want to connect to your SMB
	// file server. You specify an agent by using its Amazon Resource Name (ARN).
	//
	// This member is required.
	AgentArns []string

	// Specifies the password of the user who can mount your SMB file server and has
	// permission to access the files and folders involved in your transfer.
	//
	// For more information, see [required permissions] for SMB locations.
	//
	// [required permissions]: https://docs.aws.amazon.com/datasync/latest/userguide/create-smb-location.html#configuring-smb-permissions
	//
	// This member is required.
	Password *string

	// Specifies the Domain Name Service (DNS) name or IP address of the SMB file
	// server that your DataSync agent will mount.
	//
	// You can't specify an IP version 6 (IPv6) address.
	//
	// This member is required.
	ServerHostname *string

	// Specifies the name of the share exported by your SMB file server where DataSync
	// will read or write data. You can include a subdirectory in the share path (for
	// example, /path/to/subdirectory ). Make sure that other SMB clients in your
	// network can also mount this path.
	//
	// To copy all data in the subdirectory, DataSync must be able to mount the SMB
	// share and access all of its data. For more information, see [required permissions]for SMB locations.
	//
	// [required permissions]: https://docs.aws.amazon.com/datasync/latest/userguide/create-smb-location.html#configuring-smb-permissions
	//
	// This member is required.
	Subdirectory *string

	// Specifies the user that can mount and access the files, folders, and file
	// metadata in your SMB file server.
	//
	// For information about choosing a user with the right level of access for your
	// transfer, see [required permissions]for SMB locations.
	//
	// [required permissions]: https://docs.aws.amazon.com/datasync/latest/userguide/create-smb-location.html#configuring-smb-permissions
	//
	// This member is required.
	User *string

	// Specifies the name of the Active Directory domain that your SMB file server
	// belongs to.
	//
	// If you have multiple Active Directory domains in your environment, configuring
	// this parameter makes sure that DataSync connects to the right file server.
	Domain *string

	// Specifies the version of the SMB protocol that DataSync uses to access your SMB
	// file server.
	MountOptions *types.SmbMountOptions

	// Specifies labels that help you categorize, filter, and search for your Amazon
	// Web Services resources. We recommend creating at least a name tag for your
	// location.
	Tags []types.TagListEntry

	noSmithyDocumentSerde
}

// CreateLocationSmbResponse
type CreateLocationSmbOutput struct {

	// The ARN of the SMB location that you created.
	LocationArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateLocationSmbMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateLocationSmb{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateLocationSmb{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateLocationSmb"); err != nil {
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
	if err = addOpCreateLocationSmbValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateLocationSmb(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateLocationSmb(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateLocationSmb",
	}
}
