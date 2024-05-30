// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudtrail

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates a new event data store.
func (c *Client) CreateEventDataStore(ctx context.Context, params *CreateEventDataStoreInput, optFns ...func(*Options)) (*CreateEventDataStoreOutput, error) {
	if params == nil {
		params = &CreateEventDataStoreInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateEventDataStore", params, optFns, c.addOperationCreateEventDataStoreMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateEventDataStoreOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateEventDataStoreInput struct {

	// The name of the event data store.
	//
	// This member is required.
	Name *string

	// The advanced event selectors to use to select the events for the data store.
	// You can configure up to five advanced event selectors for each event data store.
	//
	// For more information about how to use advanced event selectors to log
	// CloudTrail events, see [Log events by using advanced event selectors]in the CloudTrail User Guide.
	//
	// For more information about how to use advanced event selectors to include
	// Config configuration items in your event data store, see [Create an event data store for Config configuration items]in the CloudTrail User
	// Guide.
	//
	// For more information about how to use advanced event selectors to include
	// events outside of Amazon Web Services events in your event data store, see [Create an integration to log events from outside Amazon Web Services]in
	// the CloudTrail User Guide.
	//
	// [Create an event data store for Config configuration items]: https://docs.aws.amazon.com/awscloudtrail/latest/userguide/lake-eds-cli.html#lake-cli-create-eds-config
	// [Log events by using advanced event selectors]: https://docs.aws.amazon.com/awscloudtrail/latest/userguide/logging-data-events-with-cloudtrail.html#creating-data-event-selectors-advanced
	// [Create an integration to log events from outside Amazon Web Services]: https://docs.aws.amazon.com/awscloudtrail/latest/userguide/lake-integrations-cli.html#lake-cli-create-integration
	AdvancedEventSelectors []types.AdvancedEventSelector

	// The billing mode for the event data store determines the cost for ingesting
	// events and the default and maximum retention period for the event data store.
	//
	// The following are the possible values:
	//
	//   - EXTENDABLE_RETENTION_PRICING - This billing mode is generally recommended if
	//   you want a flexible retention period of up to 3653 days (about 10 years). The
	//   default retention period for this billing mode is 366 days.
	//
	//   - FIXED_RETENTION_PRICING - This billing mode is recommended if you expect to
	//   ingest more than 25 TB of event data per month and need a retention period of up
	//   to 2557 days (about 7 years). The default retention period for this billing mode
	//   is 2557 days.
	//
	// The default value is EXTENDABLE_RETENTION_PRICING .
	//
	// For more information about CloudTrail pricing, see [CloudTrail Pricing] and [Managing CloudTrail Lake costs].
	//
	// [CloudTrail Pricing]: http://aws.amazon.com/cloudtrail/pricing/
	// [Managing CloudTrail Lake costs]: https://docs.aws.amazon.com/awscloudtrail/latest/userguide/cloudtrail-lake-manage-costs.html
	BillingMode types.BillingMode

	// Specifies the KMS key ID to use to encrypt the events delivered by CloudTrail.
	// The value can be an alias name prefixed by alias/ , a fully specified ARN to an
	// alias, a fully specified ARN to a key, or a globally unique identifier.
	//
	// Disabling or deleting the KMS key, or removing CloudTrail permissions on the
	// key, prevents CloudTrail from logging events to the event data store, and
	// prevents users from querying the data in the event data store that was encrypted
	// with the key. After you associate an event data store with a KMS key, the KMS
	// key cannot be removed or changed. Before you disable or delete a KMS key that
	// you are using with an event data store, delete or back up your event data store.
	//
	// CloudTrail also supports KMS multi-Region keys. For more information about
	// multi-Region keys, see [Using multi-Region keys]in the Key Management Service Developer Guide.
	//
	// Examples:
	//
	//   - alias/MyAliasName
	//
	//   - arn:aws:kms:us-east-2:123456789012:alias/MyAliasName
	//
	//   - arn:aws:kms:us-east-2:123456789012:key/12345678-1234-1234-1234-123456789012
	//
	//   - 12345678-1234-1234-1234-123456789012
	//
	// [Using multi-Region keys]: https://docs.aws.amazon.com/kms/latest/developerguide/multi-region-keys-overview.html
	KmsKeyId *string

	// Specifies whether the event data store includes events from all Regions, or
	// only from the Region in which the event data store is created.
	MultiRegionEnabled *bool

	// Specifies whether an event data store collects events logged for an
	// organization in Organizations.
	OrganizationEnabled *bool

	// The retention period of the event data store, in days. If BillingMode is set to
	// EXTENDABLE_RETENTION_PRICING , you can set a retention period of up to 3653
	// days, the equivalent of 10 years. If BillingMode is set to
	// FIXED_RETENTION_PRICING , you can set a retention period of up to 2557 days, the
	// equivalent of seven years.
	//
	// CloudTrail Lake determines whether to retain an event by checking if the
	// eventTime of the event is within the specified retention period. For example, if
	// you set a retention period of 90 days, CloudTrail will remove events when the
	// eventTime is older than 90 days.
	//
	// If you plan to copy trail events to this event data store, we recommend that
	// you consider both the age of the events that you want to copy as well as how
	// long you want to keep the copied events in your event data store. For example,
	// if you copy trail events that are 5 years old and specify a retention period of
	// 7 years, the event data store will retain those events for two years.
	RetentionPeriod *int32

	// Specifies whether the event data store should start ingesting live events. The
	// default is true.
	StartIngestion *bool

	// A list of tags.
	TagsList []types.Tag

	// Specifies whether termination protection is enabled for the event data store.
	// If termination protection is enabled, you cannot delete the event data store
	// until termination protection is disabled.
	TerminationProtectionEnabled *bool

	noSmithyDocumentSerde
}

type CreateEventDataStoreOutput struct {

	// The advanced event selectors that were used to select the events for the data
	// store.
	AdvancedEventSelectors []types.AdvancedEventSelector

	// The billing mode for the event data store.
	BillingMode types.BillingMode

	// The timestamp that shows when the event data store was created.
	CreatedTimestamp *time.Time

	// The ARN of the event data store.
	EventDataStoreArn *string

	// Specifies the KMS key ID that encrypts the events delivered by CloudTrail. The
	// value is a fully specified ARN to a KMS key in the following format.
	//
	//     arn:aws:kms:us-east-2:123456789012:key/12345678-1234-1234-1234-123456789012
	KmsKeyId *string

	// Indicates whether the event data store collects events from all Regions, or
	// only from the Region in which it was created.
	MultiRegionEnabled *bool

	// The name of the event data store.
	Name *string

	// Indicates whether an event data store is collecting logged events for an
	// organization in Organizations.
	OrganizationEnabled *bool

	// The retention period of an event data store, in days.
	RetentionPeriod *int32

	// The status of event data store creation.
	Status types.EventDataStoreStatus

	// A list of tags.
	TagsList []types.Tag

	// Indicates whether termination protection is enabled for the event data store.
	TerminationProtectionEnabled *bool

	// The timestamp that shows when an event data store was updated, if applicable.
	// UpdatedTimestamp is always either the same or newer than the time shown in
	// CreatedTimestamp .
	UpdatedTimestamp *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateEventDataStoreMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateEventDataStore{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateEventDataStore{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateEventDataStore"); err != nil {
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
	if err = addOpCreateEventDataStoreValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateEventDataStore(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateEventDataStore(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateEventDataStore",
	}
}
