// Code generated by smithy-go-codegen DO NOT EDIT.

package backup

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/backup/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// This operation returns the metadata and details specific to the backup index
// associated with the specified recovery point.
func (c *Client) GetRecoveryPointIndexDetails(ctx context.Context, params *GetRecoveryPointIndexDetailsInput, optFns ...func(*Options)) (*GetRecoveryPointIndexDetailsOutput, error) {
	if params == nil {
		params = &GetRecoveryPointIndexDetailsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetRecoveryPointIndexDetails", params, optFns, c.addOperationGetRecoveryPointIndexDetailsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetRecoveryPointIndexDetailsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetRecoveryPointIndexDetailsInput struct {

	// The name of a logical container where backups are stored. Backup vaults are
	// identified by names that are unique to the account used to create them and the
	// Region where they are created.
	//
	// Accepted characters include lowercase letters, numbers, and hyphens.
	//
	// This member is required.
	BackupVaultName *string

	// An ARN that uniquely identifies a recovery point; for example,
	// arn:aws:backup:us-east-1:123456789012:recovery-point:1EB3B5E7-9EB0-435A-A80B-108B488B0D45
	// .
	//
	// This member is required.
	RecoveryPointArn *string

	noSmithyDocumentSerde
}

type GetRecoveryPointIndexDetailsOutput struct {

	// An ARN that uniquely identifies the backup vault where the recovery point index
	// is stored.
	//
	// For example, arn:aws:backup:us-east-1:123456789012:backup-vault:aBackupVault .
	BackupVaultArn *string

	// The date and time that a backup index finished creation, in Unix format and
	// Coordinated Universal Time (UTC). The value of CreationDate is accurate to
	// milliseconds. For example, the value 1516925490.087 represents Friday, January
	// 26, 2018 12:11:30.087 AM.
	IndexCompletionDate *time.Time

	// The date and time that a backup index was created, in Unix format and
	// Coordinated Universal Time (UTC). The value of CreationDate is accurate to
	// milliseconds. For example, the value 1516925490.087 represents Friday, January
	// 26, 2018 12:11:30.087 AM.
	IndexCreationDate *time.Time

	// The date and time that a backup index was deleted, in Unix format and
	// Coordinated Universal Time (UTC). The value of CreationDate is accurate to
	// milliseconds. For example, the value 1516925490.087 represents Friday, January
	// 26, 2018 12:11:30.087 AM.
	IndexDeletionDate *time.Time

	// This is the current status for the backup index associated with the specified
	// recovery point.
	//
	// Statuses are: PENDING | ACTIVE | FAILED | DELETING
	//
	// A recovery point with an index that has the status of ACTIVE can be included in
	// a search.
	IndexStatus types.IndexStatus

	// A detailed message explaining the status of a backup index associated with the
	// recovery point.
	IndexStatusMessage *string

	// An ARN that uniquely identifies a recovery point; for example,
	// arn:aws:backup:us-east-1:123456789012:recovery-point:1EB3B5E7-9EB0-435A-A80B-108B488B0D45
	// .
	RecoveryPointArn *string

	// A string of the Amazon Resource Name (ARN) that uniquely identifies the source
	// resource.
	SourceResourceArn *string

	// Count of items within the backup index associated with the recovery point.
	TotalItemsIndexed *int64

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetRecoveryPointIndexDetailsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetRecoveryPointIndexDetails{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetRecoveryPointIndexDetails{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetRecoveryPointIndexDetails"); err != nil {
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
	if err = addOpGetRecoveryPointIndexDetailsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetRecoveryPointIndexDetails(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetRecoveryPointIndexDetails(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetRecoveryPointIndexDetails",
	}
}
