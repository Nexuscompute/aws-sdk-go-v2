// Code generated by smithy-go-codegen DO NOT EDIT.

package backup

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/backup/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This operation updates the settings of a recovery point index.
//
// Required: BackupVaultName, RecoveryPointArn, and IAMRoleArn
func (c *Client) UpdateRecoveryPointIndexSettings(ctx context.Context, params *UpdateRecoveryPointIndexSettingsInput, optFns ...func(*Options)) (*UpdateRecoveryPointIndexSettingsOutput, error) {
	if params == nil {
		params = &UpdateRecoveryPointIndexSettingsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateRecoveryPointIndexSettings", params, optFns, c.addOperationUpdateRecoveryPointIndexSettingsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateRecoveryPointIndexSettingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateRecoveryPointIndexSettingsInput struct {

	// The name of a logical container where backups are stored. Backup vaults are
	// identified by names that are unique to the account used to create them and the
	// Region where they are created.
	//
	// Accepted characters include lowercase letters, numbers, and hyphens.
	//
	// This member is required.
	BackupVaultName *string

	// Index can have 1 of 2 possible values, either ENABLED or DISABLED .
	//
	// To create a backup index for an eligible ACTIVE recovery point that does not
	// yet have a backup index, set value to ENABLED .
	//
	// To delete a backup index, set value to DISABLED .
	//
	// This member is required.
	Index types.Index

	// An ARN that uniquely identifies a recovery point; for example,
	// arn:aws:backup:us-east-1:123456789012:recovery-point:1EB3B5E7-9EB0-435A-A80B-108B488B0D45
	// .
	//
	// This member is required.
	RecoveryPointArn *string

	// This specifies the IAM role ARN used for this operation.
	//
	// For example, arn:aws:iam::123456789012:role/S3Access
	IamRoleArn *string

	noSmithyDocumentSerde
}

type UpdateRecoveryPointIndexSettingsOutput struct {

	// The name of a logical container where backups are stored. Backup vaults are
	// identified by names that are unique to the account used to create them and the
	// Region where they are created.
	BackupVaultName *string

	// Index can have 1 of 2 possible values, either ENABLED or DISABLED .
	//
	// A value of ENABLED means a backup index for an eligible ACTIVE recovery point
	// has been created.
	//
	// A value of DISABLED means a backup index was deleted.
	Index types.Index

	// This is the current status for the backup index associated with the specified
	// recovery point.
	//
	// Statuses are: PENDING | ACTIVE | FAILED | DELETING
	//
	// A recovery point with an index that has the status of ACTIVE can be included in
	// a search.
	IndexStatus types.IndexStatus

	// An ARN that uniquely identifies a recovery point; for example,
	// arn:aws:backup:us-east-1:123456789012:recovery-point:1EB3B5E7-9EB0-435A-A80B-108B488B0D45
	// .
	RecoveryPointArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateRecoveryPointIndexSettingsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateRecoveryPointIndexSettings{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateRecoveryPointIndexSettings{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateRecoveryPointIndexSettings"); err != nil {
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
	if err = addOpUpdateRecoveryPointIndexSettingsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateRecoveryPointIndexSettings(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateRecoveryPointIndexSettings(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateRecoveryPointIndexSettings",
	}
}
