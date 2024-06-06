// Code generated by smithy-go-codegen DO NOT EDIT.

package datazone

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/datazone/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Updates the specified data source in Amazon DataZone.
func (c *Client) UpdateDataSource(ctx context.Context, params *UpdateDataSourceInput, optFns ...func(*Options)) (*UpdateDataSourceOutput, error) {
	if params == nil {
		params = &UpdateDataSourceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateDataSource", params, optFns, c.addOperationUpdateDataSourceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateDataSourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateDataSourceInput struct {

	// The identifier of the domain in which to update a data source.
	//
	// This member is required.
	DomainIdentifier *string

	// The identifier of the data source to be updated.
	//
	// This member is required.
	Identifier *string

	// The asset forms to be updated as part of the UpdateDataSource action.
	AssetFormsInput []types.FormInput

	// The configuration to be updated as part of the UpdateDataSource action.
	Configuration types.DataSourceConfigurationInput

	// The description to be updated as part of the UpdateDataSource action.
	Description *string

	// The enable setting to be updated as part of the UpdateDataSource action.
	EnableSetting types.EnableSetting

	// The name to be updated as part of the UpdateDataSource action.
	Name *string

	// The publish on import setting to be updated as part of the UpdateDataSource
	// action.
	PublishOnImport *bool

	// The recommendation to be updated as part of the UpdateDataSource action.
	Recommendation *types.RecommendationConfiguration

	// The schedule to be updated as part of the UpdateDataSource action.
	Schedule *types.ScheduleConfiguration

	noSmithyDocumentSerde
}

type UpdateDataSourceOutput struct {

	// The identifier of the Amazon DataZone domain in which a data source is to be
	// updated.
	//
	// This member is required.
	DomainId *string

	// The identifier of the environment in which a data source is to be updated.
	//
	// This member is required.
	EnvironmentId *string

	// The identifier of the data source to be updated.
	//
	// This member is required.
	Id *string

	// The name to be updated as part of the UpdateDataSource action.
	//
	// This member is required.
	Name *string

	// The identifier of the project where data source is to be updated.
	//
	// This member is required.
	ProjectId *string

	// The asset forms to be updated as part of the UpdateDataSource action.
	AssetFormsOutput []types.FormOutput

	// The configuration to be updated as part of the UpdateDataSource action.
	Configuration types.DataSourceConfigurationOutput

	// The timestamp of when the data source was updated.
	CreatedAt *time.Time

	// The description to be updated as part of the UpdateDataSource action.
	Description *string

	// The enable setting to be updated as part of the UpdateDataSource action.
	EnableSetting types.EnableSetting

	// Specifies the error message that is returned if the operation cannot be
	// successfully completed.
	ErrorMessage *types.DataSourceErrorMessage

	// The timestamp of when the data source was last run.
	LastRunAt *time.Time

	// The last run error message of the data source.
	LastRunErrorMessage *types.DataSourceErrorMessage

	// The last run status of the data source.
	LastRunStatus types.DataSourceRunStatus

	// The publish on import setting to be updated as part of the UpdateDataSource
	// action.
	PublishOnImport *bool

	// The recommendation to be updated as part of the UpdateDataSource action.
	Recommendation *types.RecommendationConfiguration

	// The schedule to be updated as part of the UpdateDataSource action.
	Schedule *types.ScheduleConfiguration

	// The status to be updated as part of the UpdateDataSource action.
	Status types.DataSourceStatus

	// The type to be updated as part of the UpdateDataSource action.
	Type *string

	// The timestamp of when the data source was updated.
	UpdatedAt *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateDataSourceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateDataSource{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateDataSource{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateDataSource"); err != nil {
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
	if err = addOpUpdateDataSourceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateDataSource(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateDataSource(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateDataSource",
	}
}
