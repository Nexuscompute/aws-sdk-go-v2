// Code generated by smithy-go-codegen DO NOT EDIT.

package timestreaminfluxdb

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/timestreaminfluxdb/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new Timestream for InfluxDB DB instance.
func (c *Client) CreateDbInstance(ctx context.Context, params *CreateDbInstanceInput, optFns ...func(*Options)) (*CreateDbInstanceOutput, error) {
	if params == nil {
		params = &CreateDbInstanceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDbInstance", params, optFns, c.addOperationCreateDbInstanceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDbInstanceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDbInstanceInput struct {

	// The amount of storage to allocate for your DB storage type in GiB (gibibytes).
	//
	// This member is required.
	AllocatedStorage *int32

	// The Timestream for InfluxDB DB instance type to run InfluxDB on.
	//
	// This member is required.
	DbInstanceType types.DbInstanceType

	// The name that uniquely identifies the DB instance when interacting with the
	// Amazon Timestream for InfluxDB API and CLI commands. This name will also be a
	// prefix included in the endpoint. DB instance names must be unique per customer
	// and per region.
	//
	// This member is required.
	Name *string

	// The password of the initial admin user created in InfluxDB. This password will
	// allow you to access the InfluxDB UI to perform various administrative tasks and
	// also use the InfluxDB CLI to create an operator token. These attributes will be
	// stored in a Secret created in AWS SecretManager in your account.
	//
	// This member is required.
	Password *string

	// A list of VPC security group IDs to associate with the DB instance.
	//
	// This member is required.
	VpcSecurityGroupIds []string

	// A list of VPC subnet IDs to associate with the DB instance. Provide at least
	// two VPC subnet IDs in different availability zones when deploying with a
	// Multi-AZ standby.
	//
	// This member is required.
	VpcSubnetIds []string

	// The name of the initial InfluxDB bucket. All InfluxDB data is stored in a
	// bucket. A bucket combines the concept of a database and a retention period (the
	// duration of time that each data point persists). A bucket belongs to an
	// organization.
	Bucket *string

	// The id of the DB parameter group to assign to your DB instance. DB parameter
	// groups specify how the database is configured. For example, DB parameter groups
	// can specify the limit for query concurrency.
	DbParameterGroupIdentifier *string

	// The Timestream for InfluxDB DB storage type to read and write InfluxDB data.
	//
	// You can choose between 3 different types of provisioned Influx IOPS included
	// storage according to your workloads requirements:
	//
	//   - Influx IO Included 3000 IOPS
	//
	//   - Influx IO Included 12000 IOPS
	//
	//   - Influx IO Included 16000 IOPS
	DbStorageType types.DbStorageType

	// Specifies whether the DB instance will be deployed as a standalone instance or
	// with a Multi-AZ standby for high availability.
	DeploymentType types.DeploymentType

	// Configuration for sending InfluxDB engine logs to a specified S3 bucket.
	LogDeliveryConfiguration *types.LogDeliveryConfiguration

	// The name of the initial organization for the initial admin user in InfluxDB. An
	// InfluxDB organization is a workspace for a group of users.
	Organization *string

	// The port number on which InfluxDB accepts connections.
	//
	// Valid Values: 1024-65535
	//
	// Default: 8086
	//
	// Constraints: The value can't be 2375-2376, 7788-7799, 8090, or 51678-51680
	Port *int32

	// Configures the DB instance with a public IP to facilitate access.
	PubliclyAccessible *bool

	// A list of key-value pairs to associate with the DB instance.
	Tags map[string]string

	// The username of the initial admin user created in InfluxDB. Must start with a
	// letter and can't end with a hyphen or contain two consecutive hyphens. For
	// example, my-user1. This username will allow you to access the InfluxDB UI to
	// perform various administrative tasks and also use the InfluxDB CLI to create an
	// operator token. These attributes will be stored in a Secret created in Amazon
	// Secrets Manager in your account.
	Username *string

	noSmithyDocumentSerde
}

type CreateDbInstanceOutput struct {

	// The Amazon Resource Name (ARN) of the DB instance.
	//
	// This member is required.
	Arn *string

	// A service-generated unique identifier.
	//
	// This member is required.
	Id *string

	// The customer-supplied name that uniquely identifies the DB instance when
	// interacting with the Amazon Timestream for InfluxDB API and CLI commands.
	//
	// This member is required.
	Name *string

	// A list of VPC subnet IDs associated with the DB instance.
	//
	// This member is required.
	VpcSubnetIds []string

	// The amount of storage allocated for your DB storage type (in gibibytes).
	AllocatedStorage *int32

	// The Availability Zone in which the DB instance resides.
	AvailabilityZone *string

	// The Timestream for InfluxDB instance type that InfluxDB runs on.
	DbInstanceType types.DbInstanceType

	// The id of the DB parameter group assigned to your DB instance.
	DbParameterGroupIdentifier *string

	// The Timestream for InfluxDB DB storage type that InfluxDB stores data on.
	DbStorageType types.DbStorageType

	// Specifies whether the Timestream for InfluxDB is deployed as Single-AZ or with
	// a MultiAZ Standby for High availability.
	DeploymentType types.DeploymentType

	// The endpoint used to connect to InfluxDB. The default InfluxDB port is 8086.
	Endpoint *string

	// The Amazon Resource Name (ARN) of the AWS Secrets Manager secret containing the
	// initial InfluxDB authorization parameters. The secret value is a JSON formatted
	// key-value pair holding InfluxDB authorization values: organization, bucket,
	// username, and password.
	InfluxAuthParametersSecretArn *string

	// Configuration for sending InfluxDB engine logs to send to specified S3 bucket.
	LogDeliveryConfiguration *types.LogDeliveryConfiguration

	// The port number on which InfluxDB accepts connections. The default value is
	// 8086.
	Port *int32

	// Indicates if the DB instance has a public IP to facilitate access.
	PubliclyAccessible *bool

	// The Availability Zone in which the standby instance is located when deploying
	// with a MultiAZ standby instance.
	SecondaryAvailabilityZone *string

	// The status of the DB instance.
	Status types.Status

	// A list of VPC security group IDs associated with the DB instance.
	VpcSecurityGroupIds []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateDbInstanceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpCreateDbInstance{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpCreateDbInstance{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateDbInstance"); err != nil {
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
	if err = addOpCreateDbInstanceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDbInstance(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateDbInstance(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateDbInstance",
	}
}
