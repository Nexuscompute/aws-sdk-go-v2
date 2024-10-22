// Code generated by smithy-go-codegen DO NOT EDIT.

package timestreamquery

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/timestreamquery/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpCancelQuery struct {
}

func (*validateOpCancelQuery) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCancelQuery) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CancelQueryInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCancelQueryInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateScheduledQuery struct {
}

func (*validateOpCreateScheduledQuery) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateScheduledQuery) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateScheduledQueryInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateScheduledQueryInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteScheduledQuery struct {
}

func (*validateOpDeleteScheduledQuery) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteScheduledQuery) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteScheduledQueryInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteScheduledQueryInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeScheduledQuery struct {
}

func (*validateOpDescribeScheduledQuery) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeScheduledQuery) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeScheduledQueryInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeScheduledQueryInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpExecuteScheduledQuery struct {
}

func (*validateOpExecuteScheduledQuery) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpExecuteScheduledQuery) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ExecuteScheduledQueryInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpExecuteScheduledQueryInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListTagsForResource struct {
}

func (*validateOpListTagsForResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListTagsForResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListTagsForResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListTagsForResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpPrepareQuery struct {
}

func (*validateOpPrepareQuery) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpPrepareQuery) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*PrepareQueryInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpPrepareQueryInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpQuery struct {
}

func (*validateOpQuery) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpQuery) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*QueryInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpQueryInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpTagResource struct {
}

func (*validateOpTagResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpTagResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*TagResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpTagResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUntagResource struct {
}

func (*validateOpUntagResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUntagResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UntagResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUntagResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateScheduledQuery struct {
}

func (*validateOpUpdateScheduledQuery) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateScheduledQuery) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateScheduledQueryInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateScheduledQueryInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpCancelQueryValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCancelQuery{}, middleware.After)
}

func addOpCreateScheduledQueryValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateScheduledQuery{}, middleware.After)
}

func addOpDeleteScheduledQueryValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteScheduledQuery{}, middleware.After)
}

func addOpDescribeScheduledQueryValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeScheduledQuery{}, middleware.After)
}

func addOpExecuteScheduledQueryValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpExecuteScheduledQuery{}, middleware.After)
}

func addOpListTagsForResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListTagsForResource{}, middleware.After)
}

func addOpPrepareQueryValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpPrepareQuery{}, middleware.After)
}

func addOpQueryValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpQuery{}, middleware.After)
}

func addOpTagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpTagResource{}, middleware.After)
}

func addOpUntagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUntagResource{}, middleware.After)
}

func addOpUpdateScheduledQueryValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateScheduledQuery{}, middleware.After)
}

func validateDimensionMapping(v *types.DimensionMapping) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DimensionMapping"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if len(v.DimensionValueType) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("DimensionValueType"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateDimensionMappingList(v []types.DimensionMapping) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DimensionMappingList"}
	for i := range v {
		if err := validateDimensionMapping(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateErrorReportConfiguration(v *types.ErrorReportConfiguration) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ErrorReportConfiguration"}
	if v.S3Configuration == nil {
		invalidParams.Add(smithy.NewErrParamRequired("S3Configuration"))
	} else if v.S3Configuration != nil {
		if err := validateS3Configuration(v.S3Configuration); err != nil {
			invalidParams.AddNested("S3Configuration", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateMixedMeasureMapping(v *types.MixedMeasureMapping) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "MixedMeasureMapping"}
	if len(v.MeasureValueType) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("MeasureValueType"))
	}
	if v.MultiMeasureAttributeMappings != nil {
		if err := validateMultiMeasureAttributeMappingList(v.MultiMeasureAttributeMappings); err != nil {
			invalidParams.AddNested("MultiMeasureAttributeMappings", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateMixedMeasureMappingList(v []types.MixedMeasureMapping) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "MixedMeasureMappingList"}
	for i := range v {
		if err := validateMixedMeasureMapping(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateMultiMeasureAttributeMapping(v *types.MultiMeasureAttributeMapping) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "MultiMeasureAttributeMapping"}
	if v.SourceColumn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SourceColumn"))
	}
	if len(v.MeasureValueType) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("MeasureValueType"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateMultiMeasureAttributeMappingList(v []types.MultiMeasureAttributeMapping) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "MultiMeasureAttributeMappingList"}
	for i := range v {
		if err := validateMultiMeasureAttributeMapping(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateMultiMeasureMappings(v *types.MultiMeasureMappings) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "MultiMeasureMappings"}
	if v.MultiMeasureAttributeMappings == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MultiMeasureAttributeMappings"))
	} else if v.MultiMeasureAttributeMappings != nil {
		if err := validateMultiMeasureAttributeMappingList(v.MultiMeasureAttributeMappings); err != nil {
			invalidParams.AddNested("MultiMeasureAttributeMappings", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateNotificationConfiguration(v *types.NotificationConfiguration) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "NotificationConfiguration"}
	if v.SnsConfiguration == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SnsConfiguration"))
	} else if v.SnsConfiguration != nil {
		if err := validateSnsConfiguration(v.SnsConfiguration); err != nil {
			invalidParams.AddNested("SnsConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateQueryInsights(v *types.QueryInsights) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "QueryInsights"}
	if len(v.Mode) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Mode"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateS3Configuration(v *types.S3Configuration) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "S3Configuration"}
	if v.BucketName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("BucketName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateScheduleConfiguration(v *types.ScheduleConfiguration) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ScheduleConfiguration"}
	if v.ScheduleExpression == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ScheduleExpression"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateScheduledQueryInsights(v *types.ScheduledQueryInsights) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ScheduledQueryInsights"}
	if len(v.Mode) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Mode"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateSnsConfiguration(v *types.SnsConfiguration) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "SnsConfiguration"}
	if v.TopicArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TopicArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateTag(v *types.Tag) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Tag"}
	if v.Key == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Key"))
	}
	if v.Value == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Value"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateTagList(v []types.Tag) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TagList"}
	for i := range v {
		if err := validateTag(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateTargetConfiguration(v *types.TargetConfiguration) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TargetConfiguration"}
	if v.TimestreamConfiguration == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TimestreamConfiguration"))
	} else if v.TimestreamConfiguration != nil {
		if err := validateTimestreamConfiguration(v.TimestreamConfiguration); err != nil {
			invalidParams.AddNested("TimestreamConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateTimestreamConfiguration(v *types.TimestreamConfiguration) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TimestreamConfiguration"}
	if v.DatabaseName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DatabaseName"))
	}
	if v.TableName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TableName"))
	}
	if v.TimeColumn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TimeColumn"))
	}
	if v.DimensionMappings == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DimensionMappings"))
	} else if v.DimensionMappings != nil {
		if err := validateDimensionMappingList(v.DimensionMappings); err != nil {
			invalidParams.AddNested("DimensionMappings", err.(smithy.InvalidParamsError))
		}
	}
	if v.MultiMeasureMappings != nil {
		if err := validateMultiMeasureMappings(v.MultiMeasureMappings); err != nil {
			invalidParams.AddNested("MultiMeasureMappings", err.(smithy.InvalidParamsError))
		}
	}
	if v.MixedMeasureMappings != nil {
		if err := validateMixedMeasureMappingList(v.MixedMeasureMappings); err != nil {
			invalidParams.AddNested("MixedMeasureMappings", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCancelQueryInput(v *CancelQueryInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CancelQueryInput"}
	if v.QueryId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("QueryId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateScheduledQueryInput(v *CreateScheduledQueryInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateScheduledQueryInput"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if v.QueryString == nil {
		invalidParams.Add(smithy.NewErrParamRequired("QueryString"))
	}
	if v.ScheduleConfiguration == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ScheduleConfiguration"))
	} else if v.ScheduleConfiguration != nil {
		if err := validateScheduleConfiguration(v.ScheduleConfiguration); err != nil {
			invalidParams.AddNested("ScheduleConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if v.NotificationConfiguration == nil {
		invalidParams.Add(smithy.NewErrParamRequired("NotificationConfiguration"))
	} else if v.NotificationConfiguration != nil {
		if err := validateNotificationConfiguration(v.NotificationConfiguration); err != nil {
			invalidParams.AddNested("NotificationConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if v.TargetConfiguration != nil {
		if err := validateTargetConfiguration(v.TargetConfiguration); err != nil {
			invalidParams.AddNested("TargetConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if v.ScheduledQueryExecutionRoleArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ScheduledQueryExecutionRoleArn"))
	}
	if v.Tags != nil {
		if err := validateTagList(v.Tags); err != nil {
			invalidParams.AddNested("Tags", err.(smithy.InvalidParamsError))
		}
	}
	if v.ErrorReportConfiguration == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ErrorReportConfiguration"))
	} else if v.ErrorReportConfiguration != nil {
		if err := validateErrorReportConfiguration(v.ErrorReportConfiguration); err != nil {
			invalidParams.AddNested("ErrorReportConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteScheduledQueryInput(v *DeleteScheduledQueryInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteScheduledQueryInput"}
	if v.ScheduledQueryArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ScheduledQueryArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeScheduledQueryInput(v *DescribeScheduledQueryInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeScheduledQueryInput"}
	if v.ScheduledQueryArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ScheduledQueryArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpExecuteScheduledQueryInput(v *ExecuteScheduledQueryInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ExecuteScheduledQueryInput"}
	if v.ScheduledQueryArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ScheduledQueryArn"))
	}
	if v.InvocationTime == nil {
		invalidParams.Add(smithy.NewErrParamRequired("InvocationTime"))
	}
	if v.QueryInsights != nil {
		if err := validateScheduledQueryInsights(v.QueryInsights); err != nil {
			invalidParams.AddNested("QueryInsights", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListTagsForResourceInput(v *ListTagsForResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListTagsForResourceInput"}
	if v.ResourceARN == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceARN"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpPrepareQueryInput(v *PrepareQueryInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PrepareQueryInput"}
	if v.QueryString == nil {
		invalidParams.Add(smithy.NewErrParamRequired("QueryString"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpQueryInput(v *QueryInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "QueryInput"}
	if v.QueryString == nil {
		invalidParams.Add(smithy.NewErrParamRequired("QueryString"))
	}
	if v.QueryInsights != nil {
		if err := validateQueryInsights(v.QueryInsights); err != nil {
			invalidParams.AddNested("QueryInsights", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpTagResourceInput(v *TagResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TagResourceInput"}
	if v.ResourceARN == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceARN"))
	}
	if v.Tags == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Tags"))
	} else if v.Tags != nil {
		if err := validateTagList(v.Tags); err != nil {
			invalidParams.AddNested("Tags", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUntagResourceInput(v *UntagResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UntagResourceInput"}
	if v.ResourceARN == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceARN"))
	}
	if v.TagKeys == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TagKeys"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateScheduledQueryInput(v *UpdateScheduledQueryInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateScheduledQueryInput"}
	if v.ScheduledQueryArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ScheduledQueryArn"))
	}
	if len(v.State) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("State"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
