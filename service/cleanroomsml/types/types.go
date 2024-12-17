// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// Defines the Amazon S3 bucket where the configured audience is stored.
type AudienceDestination struct {

	// The Amazon S3 bucket and path for the configured audience.
	//
	// This member is required.
	S3Destination *S3ConfigMap

	noSmithyDocumentSerde
}

// Provides information about the audience export job.
type AudienceExportJobSummary struct {

	// The Amazon Resource Name (ARN) of the audience generation job that was exported.
	//
	// This member is required.
	AudienceGenerationJobArn *string

	// The size of the generated audience. Must match one of the sizes in the
	// configured audience model.
	//
	// This member is required.
	AudienceSize *AudienceSize

	// The time at which the audience export job was created.
	//
	// This member is required.
	CreateTime *time.Time

	// The name of the audience export job.
	//
	// This member is required.
	Name *string

	// The status of the audience export job.
	//
	// This member is required.
	Status AudienceExportJobStatus

	// The most recent time at which the audience export job was updated.
	//
	// This member is required.
	UpdateTime *time.Time

	// The description of the audience export job.
	Description *string

	// The Amazon S3 bucket where the audience export is stored.
	OutputLocation *string

	// Details about the status of a resource.
	StatusDetails *StatusDetails

	noSmithyDocumentSerde
}

// Defines the Amazon S3 bucket where the seed audience for the generating
// audience is stored.
type AudienceGenerationJobDataSource struct {

	// The ARN of the IAM role that can read the Amazon S3 bucket where the seed
	// audience is stored.
	//
	// This member is required.
	RoleArn *string

	// Defines the Amazon S3 bucket where the seed audience for the generating
	// audience is stored. A valid data source is a JSON line file in the following
	// format:
	//
	//     {"user_id": "111111"}
	//
	//     {"user_id": "222222"}
	//
	//     ...
	DataSource *S3ConfigMap

	// Provides configuration information for the instances that will perform the
	// compute work.
	SqlComputeConfiguration ComputeConfiguration

	// The protected SQL query parameters.
	SqlParameters *ProtectedQuerySQLParameters

	noSmithyDocumentSerde
}

// Provides information about the configured audience generation job.
type AudienceGenerationJobSummary struct {

	// The Amazon Resource Name (ARN) of the audience generation job.
	//
	// This member is required.
	AudienceGenerationJobArn *string

	// The Amazon Resource Name (ARN) of the configured audience model that was used
	// for this audience generation job.
	//
	// This member is required.
	ConfiguredAudienceModelArn *string

	// The time at which the audience generation job was created.
	//
	// This member is required.
	CreateTime *time.Time

	// The name of the audience generation job.
	//
	// This member is required.
	Name *string

	// The status of the audience generation job.
	//
	// This member is required.
	Status AudienceGenerationJobStatus

	// The most recent time at which the audience generation job was updated.
	//
	// This member is required.
	UpdateTime *time.Time

	// The identifier of the collaboration that contains this audience generation job.
	CollaborationId *string

	// The description of the audience generation job.
	Description *string

	// The AWS Account that submitted the job.
	StartedBy *string

	noSmithyDocumentSerde
}

// Information about the audience model.
type AudienceModelSummary struct {

	// The Amazon Resource Name (ARN) of the audience model.
	//
	// This member is required.
	AudienceModelArn *string

	// The time at which the audience model was created.
	//
	// This member is required.
	CreateTime *time.Time

	// The name of the audience model.
	//
	// This member is required.
	Name *string

	// The status of the audience model.
	//
	// This member is required.
	Status AudienceModelStatus

	// The Amazon Resource Name (ARN) of the training dataset that was used for the
	// audience model.
	//
	// This member is required.
	TrainingDatasetArn *string

	// The most recent time at which the audience model was updated.
	//
	// This member is required.
	UpdateTime *time.Time

	// The description of the audience model.
	Description *string

	noSmithyDocumentSerde
}

// Metrics that describe the quality of the generated audience.
type AudienceQualityMetrics struct {

	// The relevance scores of the generated audience.
	//
	// This member is required.
	RelevanceMetrics []RelevanceMetric

	// The recall score of the generated audience. Recall is the percentage of the
	// most similar users (by default, the most similar 20%) from a sample of the
	// training data that are included in the seed audience by the audience generation
	// job. Values range from 0-1, larger values indicate a better audience. A recall
	// value approximately equal to the maximum bin size indicates that the audience
	// model is equivalent to random selection.
	RecallMetric *float64

	noSmithyDocumentSerde
}

// The size of the generated audience. Must match one of the sizes in the
// configured audience model.
type AudienceSize struct {

	// Whether the audience size is defined in absolute terms or as a percentage. You
	// can use the ABSOLUTEAudienceSize to configure out audience sizes using the count of
	// identifiers in the output. You can use the PercentageAudienceSize to configure sizes in the
	// range 1-100 percent.
	//
	// This member is required.
	Type AudienceSizeType

	// Specify an audience size value.
	//
	// This member is required.
	Value *int32

	noSmithyDocumentSerde
}

// Returns the relevance scores at these audience sizes when used in the GetAudienceGenerationJob for a
// specified audience generation job and configured audience model.
//
// Specifies the list of allowed audienceSize values when used in the StartAudienceExportJob for an
// audience generation job. You can use the ABSOLUTEAudienceSize to configure out audience
// sizes using the count of identifiers in the output. You can use the PercentageAudienceSize
// to configure sizes in the range 1-100 percent.
type AudienceSizeConfig struct {

	// An array of the different audience output sizes.
	//
	// This member is required.
	AudienceSizeBins []int32

	// Whether the audience output sizes are defined as an absolute number or a
	// percentage.
	//
	// This member is required.
	AudienceSizeType AudienceSizeType

	noSmithyDocumentSerde
}

// Provides summary information about a configured model algorithm in a
// collaboration.
type CollaborationConfiguredModelAlgorithmAssociationSummary struct {

	// The collaboration ID of the collaboration that contains the configured model
	// algorithm association.
	//
	// This member is required.
	CollaborationIdentifier *string

	// The Amazon Resource Name (ARN) of the configured model algorithm that is
	// associated to the collaboration.
	//
	// This member is required.
	ConfiguredModelAlgorithmArn *string

	// The Amazon Resource Name (ARN) of the configured model algorithm association.
	//
	// This member is required.
	ConfiguredModelAlgorithmAssociationArn *string

	// The time at which the configured model algorithm association was created.
	//
	// This member is required.
	CreateTime *time.Time

	// The account ID of the member that created the configured model algorithm
	// association.
	//
	// This member is required.
	CreatorAccountId *string

	// The membership ID of the member that created the configured model algorithm
	// association.
	//
	// This member is required.
	MembershipIdentifier *string

	// The name of the configured model algorithm association.
	//
	// This member is required.
	Name *string

	// The most recent time at which the configured model algorithm association was
	// updated.
	//
	// This member is required.
	UpdateTime *time.Time

	// The description of the configured model algorithm association.
	Description *string

	noSmithyDocumentSerde
}

// Provides summary information about an ML input channel in a collaboration.
type CollaborationMLInputChannelSummary struct {

	// The collaboration ID of the collaboration that contains the ML input channel.
	//
	// This member is required.
	CollaborationIdentifier *string

	// The associated configured model algorithms used to create the ML input channel.
	//
	// This member is required.
	ConfiguredModelAlgorithmAssociations []string

	// The time at which the ML input channel was created.
	//
	// This member is required.
	CreateTime *time.Time

	// The account ID of the member who created the ML input channel.
	//
	// This member is required.
	CreatorAccountId *string

	// The membership ID of the membership that contains the ML input channel.
	//
	// This member is required.
	MembershipIdentifier *string

	// The Amazon Resource Name (ARN) of the ML input channel.
	//
	// This member is required.
	MlInputChannelArn *string

	// The name of the ML input channel.
	//
	// This member is required.
	Name *string

	// The status of the ML input channel.
	//
	// This member is required.
	Status MLInputChannelStatus

	// The most recent time at which the ML input channel was updated.
	//
	// This member is required.
	UpdateTime *time.Time

	// The description of the ML input channel.
	Description *string

	noSmithyDocumentSerde
}

// Provides summary information about a trained model export job in a
// collaboration.
type CollaborationTrainedModelExportJobSummary struct {

	// The collaboration ID of the collaboration that contains the trained model
	// export job.
	//
	// This member is required.
	CollaborationIdentifier *string

	// The time at which the trained model export job was created.
	//
	// This member is required.
	CreateTime *time.Time

	// The account ID of the member that created the trained model.
	//
	// This member is required.
	CreatorAccountId *string

	// The membership ID of the member that created the trained model export job.
	//
	// This member is required.
	MembershipIdentifier *string

	// The name of the trained model export job.
	//
	// This member is required.
	Name *string

	// Information about the output of the trained model export job.
	//
	// This member is required.
	OutputConfiguration *TrainedModelExportOutputConfiguration

	// The status of the trained model.
	//
	// This member is required.
	Status TrainedModelExportJobStatus

	// The Amazon Resource Name (ARN) of the trained model that is being exported.
	//
	// This member is required.
	TrainedModelArn *string

	// The most recent time at which the trained model export job was updated.
	//
	// This member is required.
	UpdateTime *time.Time

	// The description of the trained model.
	Description *string

	// Details about the status of a resource.
	StatusDetails *StatusDetails

	noSmithyDocumentSerde
}

// Provides summary information about a trained model inference job in a
// collaboration.
type CollaborationTrainedModelInferenceJobSummary struct {

	// The collaboration ID of the collaboration that contains the trained model
	// inference job.
	//
	// This member is required.
	CollaborationIdentifier *string

	// The time at which the trained model inference job was created.
	//
	// This member is required.
	CreateTime *time.Time

	// The account ID that created the trained model inference job.
	//
	// This member is required.
	CreatorAccountId *string

	// The membership ID of the membership that contains the trained model inference
	// job.
	//
	// This member is required.
	MembershipIdentifier *string

	// The name of the trained model inference job.
	//
	// This member is required.
	Name *string

	// Returns output configuration information for the trained model inference job.
	//
	// This member is required.
	OutputConfiguration *InferenceOutputConfiguration

	// The status of the trained model inference job.
	//
	// This member is required.
	Status TrainedModelInferenceJobStatus

	// The Amazon Resource Name (ARN) of the trained model that is used for the
	// trained model inference job.
	//
	// This member is required.
	TrainedModelArn *string

	// The Amazon Resource Name (ARN) of the trained model inference job.
	//
	// This member is required.
	TrainedModelInferenceJobArn *string

	// The most recent time at which the trained model inference job was updated.
	//
	// This member is required.
	UpdateTime *time.Time

	// The Amazon Resource Name (ARN) of the configured model algorithm association
	// that is used for the trained model inference job.
	ConfiguredModelAlgorithmAssociationArn *string

	// The description of the trained model inference job.
	Description *string

	// The trained model inference job logs status.
	LogsStatus LogsStatus

	// Details about the logs status for the trained model inference job.
	LogsStatusDetails *string

	// the trained model inference job metrics status.
	MetricsStatus MetricsStatus

	// Details about the metrics status for trained model inference job.
	MetricsStatusDetails *string

	noSmithyDocumentSerde
}

// Provides summary information about a trained model in a collaboration.
type CollaborationTrainedModelSummary struct {

	// The collaboration ID of the collaboration that contains the trained model.
	//
	// This member is required.
	CollaborationIdentifier *string

	// The Amazon Resource Name (ARN) of the configured model algorithm association
	// that is used for this trained model.
	//
	// This member is required.
	ConfiguredModelAlgorithmAssociationArn *string

	// The time at which the trained model was created.
	//
	// This member is required.
	CreateTime *time.Time

	// The account ID of the member that created the trained model.
	//
	// This member is required.
	CreatorAccountId *string

	// The membership ID of the member that created the trained model.
	//
	// This member is required.
	MembershipIdentifier *string

	// The name of the trained model.
	//
	// This member is required.
	Name *string

	// The status of the trained model.
	//
	// This member is required.
	Status TrainedModelStatus

	// The Amazon Resource Name (ARN) of the trained model.
	//
	// This member is required.
	TrainedModelArn *string

	// The most recent time at which the trained model was updated.
	//
	// This member is required.
	UpdateTime *time.Time

	// The description of the trained model.
	Description *string

	noSmithyDocumentSerde
}

// Metadata for a column.
type ColumnSchema struct {

	// The name of a column.
	//
	// This member is required.
	ColumnName *string

	// The data type of column.
	//
	// This member is required.
	ColumnTypes []ColumnType

	noSmithyDocumentSerde
}

// Provides configuration information for the instances that will perform the
// compute work.
//
// The following types satisfy this interface:
//
//	ComputeConfigurationMemberWorker
type ComputeConfiguration interface {
	isComputeConfiguration()
}

// The worker instances that will perform the compute work.
type ComputeConfigurationMemberWorker struct {
	Value WorkerComputeConfiguration

	noSmithyDocumentSerde
}

func (*ComputeConfigurationMemberWorker) isComputeConfiguration() {}

// Configuration information necessary for the configure audience model output.
type ConfiguredAudienceModelOutputConfig struct {

	// Defines the Amazon S3 bucket where the configured audience is stored.
	//
	// This member is required.
	Destination *AudienceDestination

	// The ARN of the IAM role that can write the Amazon S3 bucket.
	//
	// This member is required.
	RoleArn *string

	noSmithyDocumentSerde
}

// Information about the configured audience model.
type ConfiguredAudienceModelSummary struct {

	// The Amazon Resource Name (ARN) of the audience model that was used to create
	// the configured audience model.
	//
	// This member is required.
	AudienceModelArn *string

	// The Amazon Resource Name (ARN) of the configured audience model that you are
	// interested in.
	//
	// This member is required.
	ConfiguredAudienceModelArn *string

	// The time at which the configured audience model was created.
	//
	// This member is required.
	CreateTime *time.Time

	// The name of the configured audience model.
	//
	// This member is required.
	Name *string

	// The output configuration of the configured audience model.
	//
	// This member is required.
	OutputConfig *ConfiguredAudienceModelOutputConfig

	// The status of the configured audience model.
	//
	// This member is required.
	Status ConfiguredAudienceModelStatus

	// The most recent time at which the configured audience model was updated.
	//
	// This member is required.
	UpdateTime *time.Time

	// The description of the configured audience model.
	Description *string

	noSmithyDocumentSerde
}

// Provides summary information about the configured model algorithm association.
type ConfiguredModelAlgorithmAssociationSummary struct {

	// The collaboration ID of the collaboration that contains the configured model
	// algorithm association.
	//
	// This member is required.
	CollaborationIdentifier *string

	// The Amazon Resource Name (ARN) of the configured model algorithm that is being
	// associated.
	//
	// This member is required.
	ConfiguredModelAlgorithmArn *string

	// The Amazon Resource Name (ARN) of the configured model algorithm association.
	//
	// This member is required.
	ConfiguredModelAlgorithmAssociationArn *string

	// The time at which the configured model algorithm association was created.
	//
	// This member is required.
	CreateTime *time.Time

	// The membership ID of the member that created the configured model algorithm
	// association.
	//
	// This member is required.
	MembershipIdentifier *string

	// The name of the configured model algorithm association.
	//
	// This member is required.
	Name *string

	// The most recent time at which the configured model algorithm association was
	// updated.
	//
	// This member is required.
	UpdateTime *time.Time

	// The description of the configured model algorithm association.
	Description *string

	noSmithyDocumentSerde
}

// Provides summary information about a configured model algorithm.
type ConfiguredModelAlgorithmSummary struct {

	// The Amazon Resource Name (ARN) of the configured model algorithm.
	//
	// This member is required.
	ConfiguredModelAlgorithmArn *string

	// The time at which the configured model algorithm was created.
	//
	// This member is required.
	CreateTime *time.Time

	// The name of the configured model algorithm.
	//
	// This member is required.
	Name *string

	// The most recent time at which the configured model algorithm was updated.
	//
	// This member is required.
	UpdateTime *time.Time

	// The description of the configured model algorithm.
	Description *string

	noSmithyDocumentSerde
}

// Provides configuration information for the dockerized container where the model
// algorithm is stored.
type ContainerConfig struct {

	// The registry path of the docker image that contains the algorithm. Clean Rooms
	// ML supports both registry/repository[:tag] and registry/repositry[@digest]
	// image path formats. For more information about using images in Clean Rooms ML,
	// see the [Sagemaker API reference].
	//
	// [Sagemaker API reference]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_AlgorithmSpecification.html#sagemaker-Type-AlgorithmSpecification-TrainingImage
	//
	// This member is required.
	ImageUri *string

	// The arguments for a container used to run a training job. See How Amazon
	// SageMaker Runs Your Training Image for additional information. For more
	// information, see [How Sagemaker runs your training image].
	//
	// [How Sagemaker runs your training image]: https://docs.aws.amazon.com/sagemaker/latest/dg/your-algorithms-training-algo-dockerfile.html
	Arguments []string

	// The entrypoint script for a Docker container used to run a training job. This
	// script takes precedence over the default train processing instructions. See How
	// Amazon SageMaker Runs Your Training Image for additional information. For more
	// information, see [How Sagemaker runs your training image].
	//
	// [How Sagemaker runs your training image]: https://docs.aws.amazon.com/sagemaker/latest/dg/your-algorithms-training-algo-dockerfile.html
	Entrypoint []string

	// A list of metric definition objects. Each object specifies the metric name and
	// regular expressions used to parse algorithm logs. Amazon Web Services Clean
	// Rooms ML publishes each metric to all members' Amazon CloudWatch using IAM role
	// configured in PutMLConfiguration.
	MetricDefinitions []MetricDefinition

	noSmithyDocumentSerde
}

// Defines where the training dataset is located, what type of data it contains,
// and how to access the data.
type Dataset struct {

	// A DatasetInputConfig object that defines the data source and schema mapping.
	//
	// This member is required.
	InputConfig *DatasetInputConfig

	// What type of information is found in the dataset.
	//
	// This member is required.
	Type DatasetType

	noSmithyDocumentSerde
}

// Defines the Glue data source and schema mapping information.
type DatasetInputConfig struct {

	// A DataSource object that specifies the Glue data source for the training data.
	//
	// This member is required.
	DataSource *DataSource

	// The schema information for the training data.
	//
	// This member is required.
	Schema []ColumnSchema

	noSmithyDocumentSerde
}

// Defines information about the Glue data source that contains the training data.
type DataSource struct {

	// A GlueDataSource object that defines the catalog ID, database name, and table
	// name for the training data.
	//
	// This member is required.
	GlueDataSource *GlueDataSource

	noSmithyDocumentSerde
}

// The Amazon S3 location where the exported model artifacts are stored.
type Destination struct {

	// Provides information about an Amazon S3 bucket and path.
	//
	// This member is required.
	S3Destination *S3ConfigMap

	noSmithyDocumentSerde
}

// Defines the Glue data source that contains the training data.
type GlueDataSource struct {

	// The Glue database that contains the training data.
	//
	// This member is required.
	DatabaseName *string

	// The Glue table that contains the training data.
	//
	// This member is required.
	TableName *string

	// The Glue catalog that contains the training data.
	CatalogId *string

	noSmithyDocumentSerde
}

// Provides configuration information for the inference container.
type InferenceContainerConfig struct {

	// The registry path of the docker image that contains the inference algorithm.
	// Clean Rooms ML supports both registry/repository[:tag] and
	// registry/repositry[@digest] image path formats. For more information about using
	// images in Clean Rooms ML, see the [Sagemaker API reference].
	//
	// [Sagemaker API reference]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_AlgorithmSpecification.html#sagemaker-Type-AlgorithmSpecification-TrainingImage
	//
	// This member is required.
	ImageUri *string

	noSmithyDocumentSerde
}

// Provides execution parameters for the inference container.
type InferenceContainerExecutionParameters struct {

	// The maximum size of the inference container payload, specified in MB.
	MaxPayloadInMB *int32

	noSmithyDocumentSerde
}

// Configuration information about how the inference output is stored.
type InferenceOutputConfiguration struct {

	// Defines the members that can receive inference output.
	//
	// This member is required.
	Members []InferenceReceiverMember

	// The MIME type used to specify the output data.
	Accept *string

	noSmithyDocumentSerde
}

// Defines who will receive inference results.
type InferenceReceiverMember struct {

	// The account ID of the member that can receive inference results.
	//
	// This member is required.
	AccountId *string

	noSmithyDocumentSerde
}

// Defines the resources used to perform model inference.
type InferenceResourceConfig struct {

	// The type of instance that is used to perform model inference.
	//
	// This member is required.
	InstanceType InferenceInstanceType

	// The number of instances to use.
	InstanceCount *int32

	noSmithyDocumentSerde
}

// Provides information about the data source that is used to create an ML input
// channel.
type InputChannel struct {

	// The data source that is used to create the ML input channel.
	//
	// This member is required.
	DataSource InputChannelDataSource

	// The ARN of the IAM role that Clean Rooms ML can assume to read the data
	// referred to in the dataSource field the input channel.
	//
	// Passing a role across AWS accounts is not allowed. If you pass a role that
	// isn't in your account, you get an AccessDeniedException error.
	//
	// This member is required.
	RoleArn *string

	noSmithyDocumentSerde
}

// Provides the data source that is used to define an input channel.
//
// The following types satisfy this interface:
//
//	InputChannelDataSourceMemberProtectedQueryInputParameters
type InputChannelDataSource interface {
	isInputChannelDataSource()
}

// Provides information necessary to perform the protected query.
type InputChannelDataSourceMemberProtectedQueryInputParameters struct {
	Value ProtectedQueryInputParameters

	noSmithyDocumentSerde
}

func (*InputChannelDataSourceMemberProtectedQueryInputParameters) isInputChannelDataSource() {}

// Provides the information necessary for a user to access the logs.
type LogsConfigurationPolicy struct {

	// A list of account IDs that are allowed to access the logs.
	//
	// This member is required.
	AllowedAccountIds []string

	// A regular expression pattern that is used to parse the logs and return
	// information that matches the pattern.
	FilterPattern *string

	noSmithyDocumentSerde
}

// Information about the model metric that is reported for a trained model.
type MetricDefinition struct {

	// The name of the model metric.
	//
	// This member is required.
	Name *string

	// The regular expression statement that defines how the model metric is reported.
	//
	// This member is required.
	Regex *string

	noSmithyDocumentSerde
}

// Provides the configuration policy for metrics generation.
type MetricsConfigurationPolicy struct {

	// The noise level for the generated metrics.
	//
	// This member is required.
	NoiseLevel NoiseLevelType

	noSmithyDocumentSerde
}

// Provides summary information about the ML input channel.
type MLInputChannelSummary struct {

	// The collaboration ID of the collaboration that contains the ML input channel.
	//
	// This member is required.
	CollaborationIdentifier *string

	// The associated configured model algorithms used to create the ML input channel.
	//
	// This member is required.
	ConfiguredModelAlgorithmAssociations []string

	// The time at which the ML input channel was created.
	//
	// This member is required.
	CreateTime *time.Time

	// The membership ID of the membership that contains the ML input channel.
	//
	// This member is required.
	MembershipIdentifier *string

	// The Amazon Resource Name (ARN) of the ML input channel.
	//
	// This member is required.
	MlInputChannelArn *string

	// The name of the ML input channel.
	//
	// This member is required.
	Name *string

	// The status of the ML input channel.
	//
	// This member is required.
	Status MLInputChannelStatus

	// The most recent time at which the ML input channel was updated.
	//
	// This member is required.
	UpdateTime *time.Time

	// The description of the ML input channel.
	Description *string

	// The ID of the protected query that was used to create the ML input channel.
	ProtectedQueryIdentifier *string

	noSmithyDocumentSerde
}

// Configuration information about how the exported model artifacts are stored.
type MLOutputConfiguration struct {

	// The Amazon Resource Name (ARN) of the service access role that is used to store
	// the model artifacts.
	//
	// This member is required.
	RoleArn *string

	// The Amazon S3 location where exported model artifacts are stored.
	Destination *Destination

	noSmithyDocumentSerde
}

// Defines information about the data source used for model inference.
type ModelInferenceDataSource struct {

	// The Amazon Resource Name (ARN) of the ML input channel for this model inference
	// data source.
	//
	// This member is required.
	MlInputChannelArn *string

	noSmithyDocumentSerde
}

// Information about the model training data channel. A training data channel is a
// named data source that the training algorithms can consume.
type ModelTrainingDataChannel struct {

	// The name of the training data channel.
	//
	// This member is required.
	ChannelName *string

	// The Amazon Resource Name (ARN) of the ML input channel for this model training
	// data channel.
	//
	// This member is required.
	MlInputChannelArn *string

	noSmithyDocumentSerde
}

// Information about the privacy configuration for a configured model algorithm
// association.
type PrivacyConfiguration struct {

	// The privacy configuration policies for a configured model algorithm association.
	//
	// This member is required.
	Policies *PrivacyConfigurationPolicies

	noSmithyDocumentSerde
}

// Information about the privacy configuration policies for a configured model
// algorithm association.
type PrivacyConfigurationPolicies struct {

	// Specifies who will receive the trained model export.
	TrainedModelExports *TrainedModelExportsConfigurationPolicy

	// Specifies who will receive the trained model inference jobs.
	TrainedModelInferenceJobs *TrainedModelInferenceJobsConfigurationPolicy

	// Specifies who will receive the trained models.
	TrainedModels *TrainedModelsConfigurationPolicy

	noSmithyDocumentSerde
}

// Provides information necessary to perform the protected query.
type ProtectedQueryInputParameters struct {

	// The parameters for the SQL type Protected Query.
	//
	// This member is required.
	SqlParameters *ProtectedQuerySQLParameters

	// Provides configuration information for the workers that will perform the
	// protected query.
	ComputeConfiguration ComputeConfiguration

	noSmithyDocumentSerde
}

// The parameters for the SQL type Protected Query.
type ProtectedQuerySQLParameters struct {

	// The Amazon Resource Name (ARN) associated with the analysis template within a
	// collaboration.
	AnalysisTemplateArn *string

	// The protected query SQL parameters.
	Parameters map[string]string

	// The query string to be submitted.
	QueryString *string

	noSmithyDocumentSerde
}

// The relevance score of a generated audience.
type RelevanceMetric struct {

	// The size of the generated audience. Must match one of the sizes in the
	// configured audience model.
	//
	// This member is required.
	AudienceSize *AudienceSize

	// The relevance score of the generated audience.
	Score *float64

	noSmithyDocumentSerde
}

// Information about the EC2 resources that are used to train the model.
type ResourceConfig struct {

	// The instance type that is used to train the model.
	//
	// This member is required.
	InstanceType InstanceType

	// The maximum size of the instance that is used to train the model.
	//
	// This member is required.
	VolumeSizeInGB *int32

	// The number of resources that are used to train the model.
	InstanceCount *int32

	noSmithyDocumentSerde
}

// Provides information about an Amazon S3 bucket and path.
type S3ConfigMap struct {

	// The Amazon S3 location URI.
	//
	// This member is required.
	S3Uri *string

	noSmithyDocumentSerde
}

// Details about the status of a resource.
type StatusDetails struct {

	// The error message that was returned. The message is intended for human
	// consumption and can change at any time. Use the statusCode for programmatic
	// error handling.
	Message *string

	// The status code that was returned. The status code is intended for programmatic
	// error handling. Clean Rooms ML will not change the status code for existing
	// error conditions.
	StatusCode *string

	noSmithyDocumentSerde
}

// The criteria used to stop model training.
type StoppingCondition struct {

	// The maximum amount of time, in seconds, that model training can run before it
	// is terminated.
	MaxRuntimeInSeconds *int32

	noSmithyDocumentSerde
}

// Information about the output of the trained model export job.
type TrainedModelExportOutputConfiguration struct {

	// The members that will received the exported trained model output.
	//
	// This member is required.
	Members []TrainedModelExportReceiverMember

	noSmithyDocumentSerde
}

// Provides information about the member who will receive trained model exports.
type TrainedModelExportReceiverMember struct {

	// The account ID of the member who will receive trained model exports.
	//
	// This member is required.
	AccountId *string

	noSmithyDocumentSerde
}

// Information about how the trained model exports are configured.
type TrainedModelExportsConfigurationPolicy struct {

	// The files that are exported during the trained model export job.
	//
	// This member is required.
	FilesToExport []TrainedModelExportFileType

	// The maximum size of the data that can be exported.
	//
	// This member is required.
	MaxSize *TrainedModelExportsMaxSize

	noSmithyDocumentSerde
}

// The maximum size of the trained model metrics that can be exported. If the
// trained model metrics dataset is larger than this value, it will not be
// exported.
type TrainedModelExportsMaxSize struct {

	// The unit of measurement for the data size.
	//
	// This member is required.
	Unit TrainedModelExportsMaxSizeUnitType

	// The maximum size of the dataset to export.
	//
	// This member is required.
	Value *float64

	noSmithyDocumentSerde
}

// Provides configuration information for the trained model inference job.
type TrainedModelInferenceJobsConfigurationPolicy struct {

	// The logs container for the trained model inference job.
	ContainerLogs []LogsConfigurationPolicy

	// The maximum allowed size of the output of the trained model inference job.
	MaxOutputSize *TrainedModelInferenceMaxOutputSize

	noSmithyDocumentSerde
}

// Provides information about the trained model inference job.
type TrainedModelInferenceJobSummary struct {

	// The collaboration ID of the collaboration that contains the trained model
	// inference job.
	//
	// This member is required.
	CollaborationIdentifier *string

	// The time at which the trained model inference job was created.
	//
	// This member is required.
	CreateTime *time.Time

	// The membership ID of the membership that contains the trained model inference
	// job.
	//
	// This member is required.
	MembershipIdentifier *string

	// The name of the trained model inference job.
	//
	// This member is required.
	Name *string

	// The output configuration information of the trained model job.
	//
	// This member is required.
	OutputConfiguration *InferenceOutputConfiguration

	// The status of the trained model inference job.
	//
	// This member is required.
	Status TrainedModelInferenceJobStatus

	// The Amazon Resource Name (ARN) of the trained model that is used for the
	// trained model inference job.
	//
	// This member is required.
	TrainedModelArn *string

	// The Amazon Resource Name (ARN) of the trained model inference job.
	//
	// This member is required.
	TrainedModelInferenceJobArn *string

	// The most recent time at which the trained model inference job was updated.
	//
	// This member is required.
	UpdateTime *time.Time

	// The Amazon Resource Name (ARN) of the configured model algorithm association
	// that is used for the trained model inference job.
	ConfiguredModelAlgorithmAssociationArn *string

	// The description of the trained model inference job.
	Description *string

	// The log status of the trained model inference job.
	LogsStatus LogsStatus

	// Details about the log status for the trained model inference job.
	LogsStatusDetails *string

	// The metric status of the trained model inference job.
	MetricsStatus MetricsStatus

	// Details about the metrics status for the trained model inference job.
	MetricsStatusDetails *string

	noSmithyDocumentSerde
}

// Information about the maximum output size for a trained model inference job.
type TrainedModelInferenceMaxOutputSize struct {

	// The measurement unit to use.
	//
	// This member is required.
	Unit TrainedModelInferenceMaxOutputSizeUnitType

	// The maximum output size value.
	//
	// This member is required.
	Value *float64

	noSmithyDocumentSerde
}

// The configuration policy for the trained models.
type TrainedModelsConfigurationPolicy struct {

	// The container for the logs of the trained model.
	ContainerLogs []LogsConfigurationPolicy

	// The container for the metrics of the trained model.
	ContainerMetrics *MetricsConfigurationPolicy

	noSmithyDocumentSerde
}

// Summary information about the trained model.
type TrainedModelSummary struct {

	// The collaboration ID of the collaboration that contains the trained model.
	//
	// This member is required.
	CollaborationIdentifier *string

	// The Amazon Resource Name (ARN) of the configured model algorithm association
	// that was used to create this trained model.
	//
	// This member is required.
	ConfiguredModelAlgorithmAssociationArn *string

	// The time at which the trained model was created.
	//
	// This member is required.
	CreateTime *time.Time

	// The membership ID of the member that created the trained model.
	//
	// This member is required.
	MembershipIdentifier *string

	// The name of the trained model.
	//
	// This member is required.
	Name *string

	// The status of the trained model.
	//
	// This member is required.
	Status TrainedModelStatus

	// The Amazon Resource Name (ARN) of the trained model.
	//
	// This member is required.
	TrainedModelArn *string

	// The most recent time at which the trained model was updated.
	//
	// This member is required.
	UpdateTime *time.Time

	// The description of the trained model.
	Description *string

	noSmithyDocumentSerde
}

// Provides information about the training dataset.
type TrainingDatasetSummary struct {

	// The time at which the training dataset was created.
	//
	// This member is required.
	CreateTime *time.Time

	// The name of the training dataset.
	//
	// This member is required.
	Name *string

	// The status of the training dataset.
	//
	// This member is required.
	Status TrainingDatasetStatus

	// The Amazon Resource Name (ARN) of the training dataset.
	//
	// This member is required.
	TrainingDatasetArn *string

	// The most recent time at which the training dataset was updated.
	//
	// This member is required.
	UpdateTime *time.Time

	// The description of the training dataset.
	Description *string

	noSmithyDocumentSerde
}

// Configuration information about the compute workers that perform the transform
// job.
type WorkerComputeConfiguration struct {

	// The number of compute workers that are used.
	Number *int32

	// The instance type of the compute workers that are used.
	Type WorkerComputeType

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde

// UnknownUnionMember is returned when a union member is returned over the wire,
// but has an unknown tag.
type UnknownUnionMember struct {
	Tag   string
	Value []byte

	noSmithyDocumentSerde
}

func (*UnknownUnionMember) isComputeConfiguration()   {}
func (*UnknownUnionMember) isInputChannelDataSource() {}
