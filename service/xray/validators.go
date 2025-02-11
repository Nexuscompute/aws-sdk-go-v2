// Code generated by smithy-go-codegen DO NOT EDIT.

package xray

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/xray/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpBatchGetTraces struct {
}

func (*validateOpBatchGetTraces) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpBatchGetTraces) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*BatchGetTracesInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpBatchGetTracesInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCancelTraceRetrieval struct {
}

func (*validateOpCancelTraceRetrieval) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCancelTraceRetrieval) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CancelTraceRetrievalInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCancelTraceRetrievalInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateGroup struct {
}

func (*validateOpCreateGroup) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateGroup) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateGroupInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateGroupInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateSamplingRule struct {
}

func (*validateOpCreateSamplingRule) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateSamplingRule) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateSamplingRuleInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateSamplingRuleInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteResourcePolicy struct {
}

func (*validateOpDeleteResourcePolicy) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteResourcePolicy) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteResourcePolicyInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteResourcePolicyInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetInsightEvents struct {
}

func (*validateOpGetInsightEvents) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetInsightEvents) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetInsightEventsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetInsightEventsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetInsightImpactGraph struct {
}

func (*validateOpGetInsightImpactGraph) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetInsightImpactGraph) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetInsightImpactGraphInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetInsightImpactGraphInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetInsight struct {
}

func (*validateOpGetInsight) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetInsight) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetInsightInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetInsightInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetInsightSummaries struct {
}

func (*validateOpGetInsightSummaries) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetInsightSummaries) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetInsightSummariesInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetInsightSummariesInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetRetrievedTracesGraph struct {
}

func (*validateOpGetRetrievedTracesGraph) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetRetrievedTracesGraph) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetRetrievedTracesGraphInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetRetrievedTracesGraphInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetSamplingTargets struct {
}

func (*validateOpGetSamplingTargets) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetSamplingTargets) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetSamplingTargetsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetSamplingTargetsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetServiceGraph struct {
}

func (*validateOpGetServiceGraph) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetServiceGraph) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetServiceGraphInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetServiceGraphInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetTimeSeriesServiceStatistics struct {
}

func (*validateOpGetTimeSeriesServiceStatistics) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetTimeSeriesServiceStatistics) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetTimeSeriesServiceStatisticsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetTimeSeriesServiceStatisticsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetTraceGraph struct {
}

func (*validateOpGetTraceGraph) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetTraceGraph) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetTraceGraphInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetTraceGraphInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetTraceSummaries struct {
}

func (*validateOpGetTraceSummaries) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetTraceSummaries) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetTraceSummariesInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetTraceSummariesInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListRetrievedTraces struct {
}

func (*validateOpListRetrievedTraces) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListRetrievedTraces) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListRetrievedTracesInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListRetrievedTracesInput(input); err != nil {
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

type validateOpPutEncryptionConfig struct {
}

func (*validateOpPutEncryptionConfig) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpPutEncryptionConfig) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*PutEncryptionConfigInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpPutEncryptionConfigInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpPutResourcePolicy struct {
}

func (*validateOpPutResourcePolicy) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpPutResourcePolicy) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*PutResourcePolicyInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpPutResourcePolicyInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpPutTelemetryRecords struct {
}

func (*validateOpPutTelemetryRecords) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpPutTelemetryRecords) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*PutTelemetryRecordsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpPutTelemetryRecordsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpPutTraceSegments struct {
}

func (*validateOpPutTraceSegments) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpPutTraceSegments) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*PutTraceSegmentsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpPutTraceSegmentsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStartTraceRetrieval struct {
}

func (*validateOpStartTraceRetrieval) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStartTraceRetrieval) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StartTraceRetrievalInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStartTraceRetrievalInput(input); err != nil {
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

type validateOpUpdateIndexingRule struct {
}

func (*validateOpUpdateIndexingRule) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateIndexingRule) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateIndexingRuleInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateIndexingRuleInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateSamplingRule struct {
}

func (*validateOpUpdateSamplingRule) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateSamplingRule) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateSamplingRuleInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateSamplingRuleInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpBatchGetTracesValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpBatchGetTraces{}, middleware.After)
}

func addOpCancelTraceRetrievalValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCancelTraceRetrieval{}, middleware.After)
}

func addOpCreateGroupValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateGroup{}, middleware.After)
}

func addOpCreateSamplingRuleValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateSamplingRule{}, middleware.After)
}

func addOpDeleteResourcePolicyValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteResourcePolicy{}, middleware.After)
}

func addOpGetInsightEventsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetInsightEvents{}, middleware.After)
}

func addOpGetInsightImpactGraphValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetInsightImpactGraph{}, middleware.After)
}

func addOpGetInsightValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetInsight{}, middleware.After)
}

func addOpGetInsightSummariesValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetInsightSummaries{}, middleware.After)
}

func addOpGetRetrievedTracesGraphValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetRetrievedTracesGraph{}, middleware.After)
}

func addOpGetSamplingTargetsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetSamplingTargets{}, middleware.After)
}

func addOpGetServiceGraphValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetServiceGraph{}, middleware.After)
}

func addOpGetTimeSeriesServiceStatisticsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetTimeSeriesServiceStatistics{}, middleware.After)
}

func addOpGetTraceGraphValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetTraceGraph{}, middleware.After)
}

func addOpGetTraceSummariesValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetTraceSummaries{}, middleware.After)
}

func addOpListRetrievedTracesValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListRetrievedTraces{}, middleware.After)
}

func addOpListTagsForResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListTagsForResource{}, middleware.After)
}

func addOpPutEncryptionConfigValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpPutEncryptionConfig{}, middleware.After)
}

func addOpPutResourcePolicyValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpPutResourcePolicy{}, middleware.After)
}

func addOpPutTelemetryRecordsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpPutTelemetryRecords{}, middleware.After)
}

func addOpPutTraceSegmentsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpPutTraceSegments{}, middleware.After)
}

func addOpStartTraceRetrievalValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpStartTraceRetrieval{}, middleware.After)
}

func addOpTagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpTagResource{}, middleware.After)
}

func addOpUntagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUntagResource{}, middleware.After)
}

func addOpUpdateIndexingRuleValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateIndexingRule{}, middleware.After)
}

func addOpUpdateSamplingRuleValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateSamplingRule{}, middleware.After)
}

func validateIndexingRuleValueUpdate(v types.IndexingRuleValueUpdate) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "IndexingRuleValueUpdate"}
	switch uv := v.(type) {
	case *types.IndexingRuleValueUpdateMemberProbabilistic:
		if err := validateProbabilisticRuleValueUpdate(&uv.Value); err != nil {
			invalidParams.AddNested("[Probabilistic]", err.(smithy.InvalidParamsError))
		}

	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateProbabilisticRuleValueUpdate(v *types.ProbabilisticRuleValueUpdate) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ProbabilisticRuleValueUpdate"}
	if v.DesiredSamplingPercentage == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DesiredSamplingPercentage"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateSamplingRule(v *types.SamplingRule) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "SamplingRule"}
	if v.ResourceARN == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceARN"))
	}
	if v.Priority == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Priority"))
	}
	if v.ServiceName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ServiceName"))
	}
	if v.ServiceType == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ServiceType"))
	}
	if v.Host == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Host"))
	}
	if v.HTTPMethod == nil {
		invalidParams.Add(smithy.NewErrParamRequired("HTTPMethod"))
	}
	if v.URLPath == nil {
		invalidParams.Add(smithy.NewErrParamRequired("URLPath"))
	}
	if v.Version == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Version"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateSamplingStatisticsDocument(v *types.SamplingStatisticsDocument) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "SamplingStatisticsDocument"}
	if v.RuleName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RuleName"))
	}
	if v.ClientID == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ClientID"))
	}
	if v.Timestamp == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Timestamp"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateSamplingStatisticsDocumentList(v []types.SamplingStatisticsDocument) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "SamplingStatisticsDocumentList"}
	for i := range v {
		if err := validateSamplingStatisticsDocument(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
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

func validateTelemetryRecord(v *types.TelemetryRecord) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TelemetryRecord"}
	if v.Timestamp == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Timestamp"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateTelemetryRecordList(v []types.TelemetryRecord) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TelemetryRecordList"}
	for i := range v {
		if err := validateTelemetryRecord(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpBatchGetTracesInput(v *BatchGetTracesInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "BatchGetTracesInput"}
	if v.TraceIds == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TraceIds"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCancelTraceRetrievalInput(v *CancelTraceRetrievalInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CancelTraceRetrievalInput"}
	if v.RetrievalToken == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RetrievalToken"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateGroupInput(v *CreateGroupInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateGroupInput"}
	if v.GroupName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("GroupName"))
	}
	if v.Tags != nil {
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

func validateOpCreateSamplingRuleInput(v *CreateSamplingRuleInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateSamplingRuleInput"}
	if v.SamplingRule == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SamplingRule"))
	} else if v.SamplingRule != nil {
		if err := validateSamplingRule(v.SamplingRule); err != nil {
			invalidParams.AddNested("SamplingRule", err.(smithy.InvalidParamsError))
		}
	}
	if v.Tags != nil {
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

func validateOpDeleteResourcePolicyInput(v *DeleteResourcePolicyInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteResourcePolicyInput"}
	if v.PolicyName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PolicyName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetInsightEventsInput(v *GetInsightEventsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetInsightEventsInput"}
	if v.InsightId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("InsightId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetInsightImpactGraphInput(v *GetInsightImpactGraphInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetInsightImpactGraphInput"}
	if v.InsightId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("InsightId"))
	}
	if v.StartTime == nil {
		invalidParams.Add(smithy.NewErrParamRequired("StartTime"))
	}
	if v.EndTime == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EndTime"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetInsightInput(v *GetInsightInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetInsightInput"}
	if v.InsightId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("InsightId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetInsightSummariesInput(v *GetInsightSummariesInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetInsightSummariesInput"}
	if v.StartTime == nil {
		invalidParams.Add(smithy.NewErrParamRequired("StartTime"))
	}
	if v.EndTime == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EndTime"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetRetrievedTracesGraphInput(v *GetRetrievedTracesGraphInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetRetrievedTracesGraphInput"}
	if v.RetrievalToken == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RetrievalToken"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetSamplingTargetsInput(v *GetSamplingTargetsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetSamplingTargetsInput"}
	if v.SamplingStatisticsDocuments == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SamplingStatisticsDocuments"))
	} else if v.SamplingStatisticsDocuments != nil {
		if err := validateSamplingStatisticsDocumentList(v.SamplingStatisticsDocuments); err != nil {
			invalidParams.AddNested("SamplingStatisticsDocuments", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetServiceGraphInput(v *GetServiceGraphInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetServiceGraphInput"}
	if v.StartTime == nil {
		invalidParams.Add(smithy.NewErrParamRequired("StartTime"))
	}
	if v.EndTime == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EndTime"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetTimeSeriesServiceStatisticsInput(v *GetTimeSeriesServiceStatisticsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetTimeSeriesServiceStatisticsInput"}
	if v.StartTime == nil {
		invalidParams.Add(smithy.NewErrParamRequired("StartTime"))
	}
	if v.EndTime == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EndTime"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetTraceGraphInput(v *GetTraceGraphInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetTraceGraphInput"}
	if v.TraceIds == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TraceIds"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetTraceSummariesInput(v *GetTraceSummariesInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetTraceSummariesInput"}
	if v.StartTime == nil {
		invalidParams.Add(smithy.NewErrParamRequired("StartTime"))
	}
	if v.EndTime == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EndTime"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListRetrievedTracesInput(v *ListRetrievedTracesInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListRetrievedTracesInput"}
	if v.RetrievalToken == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RetrievalToken"))
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

func validateOpPutEncryptionConfigInput(v *PutEncryptionConfigInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PutEncryptionConfigInput"}
	if len(v.Type) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Type"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpPutResourcePolicyInput(v *PutResourcePolicyInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PutResourcePolicyInput"}
	if v.PolicyName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PolicyName"))
	}
	if v.PolicyDocument == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PolicyDocument"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpPutTelemetryRecordsInput(v *PutTelemetryRecordsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PutTelemetryRecordsInput"}
	if v.TelemetryRecords == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TelemetryRecords"))
	} else if v.TelemetryRecords != nil {
		if err := validateTelemetryRecordList(v.TelemetryRecords); err != nil {
			invalidParams.AddNested("TelemetryRecords", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpPutTraceSegmentsInput(v *PutTraceSegmentsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PutTraceSegmentsInput"}
	if v.TraceSegmentDocuments == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TraceSegmentDocuments"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpStartTraceRetrievalInput(v *StartTraceRetrievalInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StartTraceRetrievalInput"}
	if v.TraceIds == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TraceIds"))
	}
	if v.StartTime == nil {
		invalidParams.Add(smithy.NewErrParamRequired("StartTime"))
	}
	if v.EndTime == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EndTime"))
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

func validateOpUpdateIndexingRuleInput(v *UpdateIndexingRuleInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateIndexingRuleInput"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if v.Rule == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Rule"))
	} else if v.Rule != nil {
		if err := validateIndexingRuleValueUpdate(v.Rule); err != nil {
			invalidParams.AddNested("Rule", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateSamplingRuleInput(v *UpdateSamplingRuleInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateSamplingRuleInput"}
	if v.SamplingRuleUpdate == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SamplingRuleUpdate"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
