// Code generated by smithy-go-codegen DO NOT EDIT.

package backupsearch

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/backupsearch/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// This operation creates a search job which returns recovery points filtered by
// SearchScope and items filtered by ItemFilters.
//
// You can optionally include ClientToken, EncryptionKeyArn, Name, and/or Tags.
func (c *Client) StartSearchJob(ctx context.Context, params *StartSearchJobInput, optFns ...func(*Options)) (*StartSearchJobOutput, error) {
	if params == nil {
		params = &StartSearchJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartSearchJob", params, optFns, c.addOperationStartSearchJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartSearchJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartSearchJobInput struct {

	// This object can contain BackupResourceTypes, BackupResourceArns,
	// BackupResourceCreationTime, BackupResourceTags, and SourceResourceArns to filter
	// the recovery points returned by the search job.
	//
	// This member is required.
	SearchScope *types.SearchScope

	// Include this parameter to allow multiple identical calls for idempotency.
	//
	// A client token is valid for 8 hours after the first request that uses it is
	// completed. After this time, any request with the same token is treated as a new
	// request.
	ClientToken *string

	// The encryption key for the specified search job.
	EncryptionKeyArn *string

	// Item Filters represent all input item properties specified when the search was
	// created.
	//
	// Contains either EBSItemFilters or S3ItemFilters
	ItemFilters *types.ItemFilters

	// Include alphanumeric characters to create a name for this search job.
	Name *string

	// List of tags returned by the operation.
	Tags map[string]*string

	noSmithyDocumentSerde
}

type StartSearchJobOutput struct {

	// The date and time that a job was created, in Unix format and Coordinated
	// Universal Time (UTC). The value of CompletionTime is accurate to milliseconds.
	// For example, the value 1516925490.087 represents Friday, January 26, 2018
	// 12:11:30.087 AM.
	CreationTime *time.Time

	// The unique string that identifies the Amazon Resource Name (ARN) of the
	// specified search job.
	SearchJobArn *string

	// The unique string that specifies the search job.
	SearchJobIdentifier *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartSearchJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartSearchJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartSearchJob{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "StartSearchJob"); err != nil {
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
	if err = addOpStartSearchJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartSearchJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartSearchJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "StartSearchJob",
	}
}
