// Code generated by smithy-go-codegen DO NOT EDIT.

package codecommit

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/codecommit/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Adds or updates a file in a branch in an CodeCommit repository, and generates a
// commit for the addition in the specified branch.
func (c *Client) PutFile(ctx context.Context, params *PutFileInput, optFns ...func(*Options)) (*PutFileOutput, error) {
	if params == nil {
		params = &PutFileInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutFile", params, optFns, c.addOperationPutFileMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutFileOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutFileInput struct {

	// The name of the branch where you want to add or update the file. If this is an
	// empty repository, this branch is created.
	//
	// This member is required.
	BranchName *string

	// The content of the file, in binary object format.
	//
	// This member is required.
	FileContent []byte

	// The name of the file you want to add or update, including the relative path to
	// the file in the repository.
	//
	// If the path does not currently exist in the repository, the path is created as
	// part of adding the file.
	//
	// This member is required.
	FilePath *string

	// The name of the repository where you want to add or update the file.
	//
	// This member is required.
	RepositoryName *string

	// A message about why this file was added or updated. Although it is optional, a
	// message makes the commit history for your repository more useful.
	CommitMessage *string

	// An email address for the person adding or updating the file.
	Email *string

	// The file mode permissions of the blob. Valid file mode permissions are listed
	// here.
	FileMode types.FileModeTypeEnum

	// The name of the person adding or updating the file. Although it is optional, a
	// name makes the commit history for your repository more useful.
	Name *string

	// The full commit ID of the head commit in the branch where you want to add or
	// update the file. If this is an empty repository, no commit ID is required. If
	// this is not an empty repository, a commit ID is required.
	//
	// The commit ID must match the ID of the head commit at the time of the
	// operation. Otherwise, an error occurs, and the file is not added or updated.
	ParentCommitId *string

	noSmithyDocumentSerde
}

type PutFileOutput struct {

	// The ID of the blob, which is its SHA-1 pointer.
	//
	// This member is required.
	BlobId *string

	// The full SHA ID of the commit that contains this file change.
	//
	// This member is required.
	CommitId *string

	// The full SHA-1 pointer of the tree information for the commit that contains
	// this file change.
	//
	// This member is required.
	TreeId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutFileMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutFile{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutFile{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PutFile"); err != nil {
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
	if err = addOpPutFileValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutFile(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutFile(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PutFile",
	}
}
