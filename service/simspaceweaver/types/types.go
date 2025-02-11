// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// The Amazon CloudWatch Logs log group for the simulation. For more information
// about log groups, see [Working with log groups and log streams]in the Amazon CloudWatch Logs User Guide.
//
// [Working with log groups and log streams]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/Working-with-log-groups-and-streams.html
type CloudWatchLogsLogGroup struct {

	// The Amazon Resource Name (ARN) of the Amazon CloudWatch Logs log group for the
	// simulation. For more information about ARNs, see [Amazon Resource Names (ARNs)]in the Amazon Web Services
	// General Reference. For more information about log groups, see [Working with log groups and log streams]in the Amazon
	// CloudWatch Logs User Guide.
	//
	// [Working with log groups and log streams]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/Working-with-log-groups-and-streams.html
	// [Amazon Resource Names (ARNs)]: https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html
	LogGroupArn *string

	noSmithyDocumentSerde
}

// A collection of app instances that run the same executable app code and have
// the same launch options and commands.
//
// For more information about domains, see [Key concepts: Domains] in the SimSpace Weaver User Guide.
//
// [Key concepts: Domains]: https://docs.aws.amazon.com/simspaceweaver/latest/userguide/what-is_key-concepts.html#what-is_key-concepts_domains
type Domain struct {

	// The type of lifecycle management for apps in the domain. Indicates whether apps
	// in this domain are managed (SimSpace Weaver starts and stops the apps) or
	// unmanaged (you must start and stop the apps).
	//
	// Lifecycle types
	//
	//   - PerWorker – Managed: SimSpace Weaver starts one app on each worker.
	//
	//   - BySpatialSubdivision – Managed: SimSpace Weaver starts one app for each
	//   spatial partition.
	//
	//   - ByRequest – Unmanaged: You use the StartApp API to start the apps and use
	//   the StopApp API to stop the apps.
	Lifecycle LifecycleManagementStrategy

	// The name of the domain.
	Name *string

	noSmithyDocumentSerde
}

// Options that apply when the app starts. These options override default behavior.
type LaunchOverrides struct {

	// App launch commands and command line parameters that override the launch
	// command configured in the simulation schema.
	LaunchCommands []string

	noSmithyDocumentSerde
}

// A collection of additional state information, such as domain and clock
// configuration.
type LiveSimulationState struct {

	// A list of simulation clocks.
	//
	// At this time, a simulation has only one clock.
	Clocks []SimulationClock

	// A list of domains for the simulation. For more information about domains, see [Key concepts: Domains]
	// in the SimSpace Weaver User Guide.
	//
	// [Key concepts: Domains]: https://docs.aws.amazon.com/simspaceweaver/latest/userguide/what-is_key-concepts.html#what-is_key-concepts_domains
	Domains []Domain

	noSmithyDocumentSerde
}

// The location where SimSpace Weaver sends simulation log data.
type LogDestination struct {

	// An Amazon CloudWatch Logs log group that stores simulation log data. For more
	// information about log groups, see [Working with log groups and log streams]in the Amazon CloudWatch Logs User Guide.
	//
	// [Working with log groups and log streams]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/Working-with-log-groups-and-streams.html
	CloudWatchLogsLogGroup *CloudWatchLogsLogGroup

	noSmithyDocumentSerde
}

// The logging configuration for a simulation.
type LoggingConfiguration struct {

	// A list of the locations where SimSpace Weaver sends simulation log data.
	Destinations []LogDestination

	noSmithyDocumentSerde
}

// An Amazon S3 bucket and optional folder (object key prefix) where SimSpace
// Weaver creates a file.
type S3Destination struct {

	// The name of an Amazon S3 bucket. For more information about buckets, see [Creating, configuring, and working with Amazon S3 buckets] in
	// the Amazon Simple Storage Service User Guide.
	//
	// [Creating, configuring, and working with Amazon S3 buckets]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/creating-buckets-s3.html
	//
	// This member is required.
	BucketName *string

	// A string prefix for an Amazon S3 object key. It's usually a folder name. For
	// more information about folders in Amazon S3, see [Organizing objects in the Amazon S3 console using folders]in the Amazon Simple Storage
	// Service User Guide.
	//
	// [Organizing objects in the Amazon S3 console using folders]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/using-folders.html
	ObjectKeyPrefix *string

	noSmithyDocumentSerde
}

// A location in Amazon Simple Storage Service (Amazon S3) where SimSpace Weaver
// stores simulation data, such as your app .zip files and schema file. For more
// information about Amazon S3, see the [Amazon Simple Storage Service User Guide].
//
// [Amazon Simple Storage Service User Guide]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/Welcome.html
type S3Location struct {

	// The name of an Amazon S3 bucket. For more information about buckets, see [Creating, configuring, and working with Amazon S3 buckets] in
	// the Amazon Simple Storage Service User Guide.
	//
	// [Creating, configuring, and working with Amazon S3 buckets]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/creating-buckets-s3.html
	//
	// This member is required.
	BucketName *string

	// The key name of an object in Amazon S3. For more information about Amazon S3
	// objects and object keys, see [Uploading, downloading, and working with objects in Amazon S3]in the Amazon Simple Storage Service User Guide.
	//
	// [Uploading, downloading, and working with objects in Amazon S3]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/uploading-downloading-objects.html
	//
	// This member is required.
	ObjectKey *string

	noSmithyDocumentSerde
}

// Information about the network endpoint that you can use to connect to your
// custom or service app. For more information about SimSpace Weaver apps, see [Key concepts: Apps]in
// the SimSpace Weaver User Guide..
//
// [Key concepts: Apps]: https://docs.aws.amazon.com/simspaceweaver/latest/userguide/what-is_key-concepts.html#what-is_key-concepts_apps
type SimulationAppEndpointInfo struct {

	// The IP address of the app. SimSpace Weaver dynamically assigns this IP address
	// when the app starts.
	Address *string

	// The inbound TCP/UDP port numbers of the app. The combination of an IP address
	// and a port number form a network endpoint.
	IngressPortMappings []SimulationAppPortMapping

	noSmithyDocumentSerde
}

// A collection of metadata about the app.
type SimulationAppMetadata struct {

	// The domain of the app. For more information about domains, see [Key concepts: Domains] in the SimSpace
	// Weaver User Guide.
	//
	// [Key concepts: Domains]: https://docs.aws.amazon.com/simspaceweaver/latest/userguide/what-is_key-concepts.html#what-is_key-concepts_domains
	Domain *string

	// The name of the app.
	Name *string

	// The name of the simulation of the app.
	Simulation *string

	// The current status of the app.
	Status SimulationAppStatus

	// The desired status of the app.
	TargetStatus SimulationAppTargetStatus

	noSmithyDocumentSerde
}

// A collection of TCP/UDP ports for a custom or service app.
type SimulationAppPortMapping struct {

	// The TCP/UDP port number of the running app. SimSpace Weaver dynamically assigns
	// this port number when the app starts. SimSpace Weaver maps the Declared port to
	// the Actual port. Clients connect to the app using the app's IP address and the
	// Actual port number.
	Actual *int32

	// The TCP/UDP port number of the app, declared in the simulation schema. SimSpace
	// Weaver maps the Declared port to the Actual port. The source code for the app
	// should bind to the Declared port.
	Declared *int32

	noSmithyDocumentSerde
}

// Status information about the simulation clock.
type SimulationClock struct {

	// The current status of the simulation clock.
	Status ClockStatus

	// The desired status of the simulation clock.
	TargetStatus ClockTargetStatus

	noSmithyDocumentSerde
}

// A collection of data about the simulation.
type SimulationMetadata struct {

	// The Amazon Resource Name (ARN) of the simulation. For more information about
	// ARNs, see [Amazon Resource Names (ARNs)]in the Amazon Web Services General Reference.
	//
	// [Amazon Resource Names (ARNs)]: https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html
	Arn *string

	// The time when the simulation was created, expressed as the number of seconds
	// and milliseconds in UTC since the Unix epoch (0:0:0.000, January 1, 1970).
	CreationTime *time.Time

	// The name of the simulation.
	Name *string

	// The current status of the simulation.
	Status SimulationStatus

	// The desired status of the simulation.
	TargetStatus SimulationTargetStatus

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
