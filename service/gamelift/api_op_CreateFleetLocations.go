// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/gamelift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

//	This operation has been expanded to use with the Amazon GameLift containers
//
// feature, which is currently in public preview.
//
// Adds remote locations to an EC2 or container fleet and begins populating the
// new locations with instances. The new instances conform to the fleet's instance
// type, auto-scaling, and other configuration settings.
//
// You can't add remote locations to a fleet that resides in an Amazon Web
// Services Region that doesn't support multiple locations. Fleets created prior to
// March 2021 can't support multiple locations.
//
// To add fleet locations, specify the fleet to be updated and provide a list of
// one or more locations.
//
// If successful, this operation returns the list of added locations with their
// status set to NEW . Amazon GameLift initiates the process of starting an
// instance in each added location. You can track the status of each new location
// by monitoring location creation events using [DescribeFleetEvents].
//
// # Learn more
//
// [Setting up fleets]
//
// [Update fleet locations]
//
// [Amazon GameLift service locations]for managed hosting.
//
// [DescribeFleetEvents]: https://docs.aws.amazon.com/gamelift/latest/apireference/API_DescribeFleetEvents.html
// [Amazon GameLift service locations]: https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-regions.html
// [Update fleet locations]: https://docs.aws.amazon.com/gamelift/latest/developerguide/fleets-editing.html#fleets-update-locations
// [Setting up fleets]: https://docs.aws.amazon.com/gamelift/latest/developerguide/fleets-intro.html
func (c *Client) CreateFleetLocations(ctx context.Context, params *CreateFleetLocationsInput, optFns ...func(*Options)) (*CreateFleetLocationsOutput, error) {
	if params == nil {
		params = &CreateFleetLocationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateFleetLocations", params, optFns, c.addOperationCreateFleetLocationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateFleetLocationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateFleetLocationsInput struct {

	// A unique identifier for the fleet to add locations to. You can use either the
	// fleet ID or ARN value.
	//
	// This member is required.
	FleetId *string

	// A list of locations to deploy additional instances to and manage as part of the
	// fleet. You can add any Amazon GameLift-supported Amazon Web Services Region as a
	// remote location, in the form of an Amazon Web Services Region code such as
	// us-west-2 .
	//
	// This member is required.
	Locations []types.LocationConfiguration

	noSmithyDocumentSerde
}

type CreateFleetLocationsOutput struct {

	// The Amazon Resource Name ([ARN] ) that is assigned to a Amazon GameLift fleet
	// resource and uniquely identifies it. ARNs are unique across all Regions. Format
	// is arn:aws:gamelift:::fleet/fleet-a1234567-b8c9-0d1e-2fa3-b45c6d7e8912 .
	//
	// [ARN]: https://docs.aws.amazon.com/AmazonS3/latest/dev/s3-arn-format.html
	FleetArn *string

	// A unique identifier for the fleet that was updated with new locations.
	FleetId *string

	// The remote locations that are being added to the fleet, and the life-cycle
	// status of each location. For new locations, the status is set to NEW . During
	// location creation, Amazon GameLift updates each location's status as instances
	// are deployed there and prepared for game hosting. This list does not include the
	// fleet home Region or any remote locations that were already added to the fleet.
	LocationStates []types.LocationState

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateFleetLocationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateFleetLocations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateFleetLocations{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateFleetLocations"); err != nil {
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
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = addOpCreateFleetLocationsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateFleetLocations(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateFleetLocations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateFleetLocations",
	}
}
