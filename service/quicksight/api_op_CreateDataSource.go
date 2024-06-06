// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a data source.
func (c *Client) CreateDataSource(ctx context.Context, params *CreateDataSourceInput, optFns ...func(*Options)) (*CreateDataSourceOutput, error) {
	if params == nil {
		params = &CreateDataSourceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDataSource", params, optFns, c.addOperationCreateDataSourceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDataSourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDataSourceInput struct {

	// The Amazon Web Services account ID.
	//
	// This member is required.
	AwsAccountId *string

	// An ID for the data source. This ID is unique per Amazon Web Services Region for
	// each Amazon Web Services account.
	//
	// This member is required.
	DataSourceId *string

	// A display name for the data source.
	//
	// This member is required.
	Name *string

	// The type of the data source. To return a list of all data sources, use
	// ListDataSources .
	//
	// Use AMAZON_ELASTICSEARCH for Amazon OpenSearch Service.
	//
	// This member is required.
	Type types.DataSourceType

	// The credentials Amazon QuickSight that uses to connect to your underlying
	// source. Currently, only credentials based on user name and password are
	// supported.
	Credentials *types.DataSourceCredentials

	// The parameters that Amazon QuickSight uses to connect to your underlying source.
	DataSourceParameters types.DataSourceParameters

	// When you create the data source, Amazon QuickSight adds the data source to
	// these folders.
	FolderArns []string

	// A list of resource permissions on the data source.
	Permissions []types.ResourcePermission

	// Secure Socket Layer (SSL) properties that apply when Amazon QuickSight connects
	// to your underlying source.
	SslProperties *types.SslProperties

	// Contains a map of the key-value pairs for the resource tag or tags assigned to
	// the data source.
	Tags []types.Tag

	// Use this parameter only when you want Amazon QuickSight to use a VPC connection
	// when connecting to your underlying source.
	VpcConnectionProperties *types.VpcConnectionProperties

	noSmithyDocumentSerde
}

type CreateDataSourceOutput struct {

	// The Amazon Resource Name (ARN) of the data source.
	Arn *string

	// The status of creating the data source.
	CreationStatus types.ResourceStatus

	// The ID of the data source. This ID is unique per Amazon Web Services Region for
	// each Amazon Web Services account.
	DataSourceId *string

	// The Amazon Web Services request ID for this operation.
	RequestId *string

	// The HTTP status of the request.
	Status int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateDataSourceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateDataSource{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateDataSource{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateDataSource"); err != nil {
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
	if err = addOpCreateDataSourceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDataSource(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateDataSource(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateDataSource",
	}
}
