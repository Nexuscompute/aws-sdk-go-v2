// Code generated by smithy-go-codegen DO NOT EDIT.

package computeoptimizer

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/computeoptimizer/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the recommendation preferences that are in effect for a given resource,
// such as enhanced infrastructure metrics. Considers all applicable preferences
// that you might have set at the resource, account, and organization level.
//
// When you create a recommendation preference, you can set its status to Active
// or Inactive . Use this action to view the recommendation preferences that are in
// effect, or Active .
func (c *Client) GetEffectiveRecommendationPreferences(ctx context.Context, params *GetEffectiveRecommendationPreferencesInput, optFns ...func(*Options)) (*GetEffectiveRecommendationPreferencesOutput, error) {
	if params == nil {
		params = &GetEffectiveRecommendationPreferencesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetEffectiveRecommendationPreferences", params, optFns, c.addOperationGetEffectiveRecommendationPreferencesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetEffectiveRecommendationPreferencesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetEffectiveRecommendationPreferencesInput struct {

	// The Amazon Resource Name (ARN) of the resource for which to confirm effective
	// recommendation preferences. Only EC2 instance and Auto Scaling group ARNs are
	// currently supported.
	//
	// This member is required.
	ResourceArn *string

	noSmithyDocumentSerde
}

type GetEffectiveRecommendationPreferencesOutput struct {

	// The status of the enhanced infrastructure metrics recommendation preference.
	// Considers all applicable preferences that you might have set at the resource,
	// account, and organization level.
	//
	// A status of Active confirms that the preference is applied in the latest
	// recommendation refresh, and a status of Inactive confirms that it's not yet
	// applied to recommendations.
	//
	// To validate whether the preference is applied to your last generated set of
	// recommendations, review the effectiveRecommendationPreferences value in the
	// response of the GetAutoScalingGroupRecommendationsand GetEC2InstanceRecommendations actions.
	//
	// For more information, see [Enhanced infrastructure metrics] in the Compute Optimizer User Guide.
	//
	// [Enhanced infrastructure metrics]: https://docs.aws.amazon.com/compute-optimizer/latest/ug/enhanced-infrastructure-metrics.html
	EnhancedInfrastructureMetrics types.EnhancedInfrastructureMetrics

	// The provider of the external metrics recommendation preference. Considers all
	// applicable preferences that you might have set at the account and organization
	// level.
	//
	// If the preference is applied in the latest recommendation refresh, an object
	// with a valid source value appears in the response. If the preference isn't
	// applied to the recommendations already, then this object doesn't appear in the
	// response.
	//
	// To validate whether the preference is applied to your last generated set of
	// recommendations, review the effectiveRecommendationPreferences value in the
	// response of the GetEC2InstanceRecommendationsactions.
	//
	// For more information, see [Enhanced infrastructure metrics] in the Compute Optimizer User Guide.
	//
	// [Enhanced infrastructure metrics]: https://docs.aws.amazon.com/compute-optimizer/latest/ug/external-metrics-ingestion.html
	ExternalMetricsPreference *types.ExternalMetricsPreference

	//  The number of days the utilization metrics of the Amazon Web Services resource
	// are analyzed.
	//
	// To validate that the preference is applied to your last generated set of
	// recommendations, review the effectiveRecommendationPreferences value in the
	// response of the GetAutoScalingGroupRecommendations or
	// GetEC2InstanceRecommendations actions.
	LookBackPeriod types.LookBackPeriodPreference

	//  The resource type values that are considered as candidates when generating
	// rightsizing recommendations. This object resolves any wildcard expressions and
	// returns the effective list of candidate resource type values. It also considers
	// all applicable preferences that you set at the resource, account, and
	// organization level.
	//
	// To validate that the preference is applied to your last generated set of
	// recommendations, review the effectiveRecommendationPreferences value in the
	// response of the GetAutoScalingGroupRecommendations or
	// GetEC2InstanceRecommendations actions.
	PreferredResources []types.EffectivePreferredResource

	//  The resource’s CPU and memory utilization preferences, such as threshold and
	// headroom, that were used to generate rightsizing recommendations. It considers
	// all applicable preferences that you set at the resource, account, and
	// organization level.
	//
	// To validate that the preference is applied to your last generated set of
	// recommendations, review the effectiveRecommendationPreferences value in the
	// response of the GetAutoScalingGroupRecommendations or
	// GetEC2InstanceRecommendations actions.
	UtilizationPreferences []types.UtilizationPreference

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetEffectiveRecommendationPreferencesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpGetEffectiveRecommendationPreferences{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpGetEffectiveRecommendationPreferences{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetEffectiveRecommendationPreferences"); err != nil {
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
	if err = addOpGetEffectiveRecommendationPreferencesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetEffectiveRecommendationPreferences(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetEffectiveRecommendationPreferences(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetEffectiveRecommendationPreferences",
	}
}
