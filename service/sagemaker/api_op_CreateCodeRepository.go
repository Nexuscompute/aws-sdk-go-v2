// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a Git repository as a resource in your SageMaker AI account. You can
// associate the repository with notebook instances so that you can use Git source
// control for the notebooks you create. The Git repository is a resource in your
// SageMaker AI account, so it can be associated with more than one notebook
// instance, and it persists independently from the lifecycle of any notebook
// instances it is associated with.
//
// The repository can be hosted either in [Amazon Web Services CodeCommit] or in any other Git repository.
//
// [Amazon Web Services CodeCommit]: https://docs.aws.amazon.com/codecommit/latest/userguide/welcome.html
func (c *Client) CreateCodeRepository(ctx context.Context, params *CreateCodeRepositoryInput, optFns ...func(*Options)) (*CreateCodeRepositoryOutput, error) {
	if params == nil {
		params = &CreateCodeRepositoryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateCodeRepository", params, optFns, c.addOperationCreateCodeRepositoryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateCodeRepositoryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateCodeRepositoryInput struct {

	// The name of the Git repository. The name must have 1 to 63 characters. Valid
	// characters are a-z, A-Z, 0-9, and - (hyphen).
	//
	// This member is required.
	CodeRepositoryName *string

	// Specifies details about the repository, including the URL where the repository
	// is located, the default branch, and credentials to use to access the repository.
	//
	// This member is required.
	GitConfig *types.GitConfig

	// An array of key-value pairs. You can use tags to categorize your Amazon Web
	// Services resources in different ways, for example, by purpose, owner, or
	// environment. For more information, see [Tagging Amazon Web Services Resources].
	//
	// [Tagging Amazon Web Services Resources]: https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateCodeRepositoryOutput struct {

	// The Amazon Resource Name (ARN) of the new repository.
	//
	// This member is required.
	CodeRepositoryArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateCodeRepositoryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateCodeRepository{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateCodeRepository{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateCodeRepository"); err != nil {
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
	if err = addOpCreateCodeRepositoryValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateCodeRepository(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateCodeRepository(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateCodeRepository",
	}
}
