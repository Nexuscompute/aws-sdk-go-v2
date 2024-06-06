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

// Creates a transfer location for an object storage system. DataSync can use this
// location as a source or destination for transferring data.
//
// Before you begin, make sure that you understand the [prerequisites] for DataSync to work with
// object storage systems.
//
// [prerequisites]: https://docs.aws.amazon.com/datasync/latest/userguide/create-object-location.html#create-object-location-prerequisites
func (c *Client) CreateLocationObjectStorage(ctx context.Context, params *CreateLocationObjectStorageInput, optFns ...func(*Options)) (*CreateLocationObjectStorageOutput, error) {
	if params == nil {
		params = &CreateLocationObjectStorageInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateLocationObjectStorage", params, optFns, c.addOperationCreateLocationObjectStorageMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateLocationObjectStorageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// CreateLocationObjectStorageRequest
type CreateLocationObjectStorageInput struct {

	// Specifies the Amazon Resource Names (ARNs) of the DataSync agents that can
	// securely connect with your location.
	//
	// This member is required.
	AgentArns []string

	// Specifies the name of the object storage bucket involved in the transfer.
	//
	// This member is required.
	BucketName *string

	// Specifies the domain name or IP address of the object storage server. A
	// DataSync agent uses this hostname to mount the object storage server in a
	// network.
	//
	// This member is required.
	ServerHostname *string

	// Specifies the access key (for example, a user name) if credentials are required
	// to authenticate with the object storage server.
	AccessKey *string

	// Specifies the secret key (for example, a password) if credentials are required
	// to authenticate with the object storage server.
	SecretKey *string

	// Specifies a certificate chain for DataSync to authenticate with your object
	// storage system if the system uses a private or self-signed certificate authority
	// (CA). You must specify a single .pem file with a full certificate chain (for
	// example, file:///home/user/.ssh/object_storage_certificates.pem ).
	//
	// The certificate chain might include:
	//
	//   - The object storage system's certificate
	//
	//   - All intermediate certificates (if there are any)
	//
	//   - The root certificate of the signing CA
	//
	// You can concatenate your certificates into a .pem file (which can be up to
	// 32768 bytes before base64 encoding). The following example cat command creates
	// an object_storage_certificates.pem file that includes three certificates:
	//
	//     cat object_server_certificate.pem intermediate_certificate.pem
	//     ca_root_certificate.pem > object_storage_certificates.pem
	//
	// To use this parameter, configure ServerProtocol to HTTPS .
	ServerCertificate []byte

	// Specifies the port that your object storage server accepts inbound network
	// traffic on (for example, port 443).
	ServerPort *int32

	// Specifies the protocol that your object storage server uses to communicate.
	ServerProtocol types.ObjectStorageServerProtocol

	// Specifies the object prefix for your object storage server. If this is a source
	// location, DataSync only copies objects with this prefix. If this is a
	// destination location, DataSync writes all objects with this prefix.
	Subdirectory *string

	// Specifies the key-value pair that represents a tag that you want to add to the
	// resource. Tags can help you manage, filter, and search for your resources. We
	// recommend creating a name tag for your location.
	Tags []types.TagListEntry

	noSmithyDocumentSerde
}

// CreateLocationObjectStorageResponse
type CreateLocationObjectStorageOutput struct {

	// Specifies the ARN of the object storage system location that you create.
	LocationArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateLocationObjectStorageMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateLocationObjectStorage{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateLocationObjectStorage{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateLocationObjectStorage"); err != nil {
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
	if err = addOpCreateLocationObjectStorageValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateLocationObjectStorage(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateLocationObjectStorage(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateLocationObjectStorage",
	}
}
