// Code generated by smithy-go-codegen DO NOT EDIT.

package types_test

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/bedrockagentruntime/document"
	"github.com/aws/aws-sdk-go-v2/service/bedrockagentruntime/types"
)

func ExampleActionGroupExecutor_outputUsage() {
	var union types.ActionGroupExecutor
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.ActionGroupExecutorMemberCustomControl:
		_ = v.Value // Value is types.CustomControlMethod

	case *types.ActionGroupExecutorMemberLambda:
		_ = v.Value // Value is string

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *string
var _ types.CustomControlMethod

func ExampleAPISchema_outputUsage() {
	var union types.APISchema
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.APISchemaMemberPayload:
		_ = v.Value // Value is string

	case *types.APISchemaMemberS3:
		_ = v.Value // Value is types.S3Identifier

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.S3Identifier
var _ *string

func ExampleCaller_outputUsage() {
	var union types.Caller
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.CallerMemberAgentAliasArn:
		_ = v.Value // Value is string

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *string

func ExampleContentBlock_outputUsage() {
	var union types.ContentBlock
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.ContentBlockMemberText:
		_ = v.Value // Value is string

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *string

func ExampleFlowInputContent_outputUsage() {
	var union types.FlowInputContent
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.FlowInputContentMemberDocument:
		_ = v.Value // Value is document.Interface

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ document.Interface

func ExampleFlowMultiTurnInputContent_outputUsage() {
	var union types.FlowMultiTurnInputContent
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.FlowMultiTurnInputContentMemberDocument:
		_ = v.Value // Value is document.Interface

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ document.Interface

func ExampleFlowOutputContent_outputUsage() {
	var union types.FlowOutputContent
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.FlowOutputContentMemberDocument:
		_ = v.Value // Value is document.Interface

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ document.Interface

func ExampleFlowResponseStream_outputUsage() {
	var union types.FlowResponseStream
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.FlowResponseStreamMemberFlowCompletionEvent:
		_ = v.Value // Value is types.FlowCompletionEvent

	case *types.FlowResponseStreamMemberFlowMultiTurnInputRequestEvent:
		_ = v.Value // Value is types.FlowMultiTurnInputRequestEvent

	case *types.FlowResponseStreamMemberFlowOutputEvent:
		_ = v.Value // Value is types.FlowOutputEvent

	case *types.FlowResponseStreamMemberFlowTraceEvent:
		_ = v.Value // Value is types.FlowTraceEvent

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.FlowOutputEvent
var _ *types.FlowTraceEvent
var _ *types.FlowCompletionEvent
var _ *types.FlowMultiTurnInputRequestEvent

func ExampleFlowTrace_outputUsage() {
	var union types.FlowTrace
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.FlowTraceMemberConditionNodeResultTrace:
		_ = v.Value // Value is types.FlowTraceConditionNodeResultEvent

	case *types.FlowTraceMemberNodeInputTrace:
		_ = v.Value // Value is types.FlowTraceNodeInputEvent

	case *types.FlowTraceMemberNodeOutputTrace:
		_ = v.Value // Value is types.FlowTraceNodeOutputEvent

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.FlowTraceNodeInputEvent
var _ *types.FlowTraceConditionNodeResultEvent
var _ *types.FlowTraceNodeOutputEvent

func ExampleFlowTraceNodeInputContent_outputUsage() {
	var union types.FlowTraceNodeInputContent
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.FlowTraceNodeInputContentMemberDocument:
		_ = v.Value // Value is document.Interface

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ document.Interface

func ExampleFlowTraceNodeOutputContent_outputUsage() {
	var union types.FlowTraceNodeOutputContent
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.FlowTraceNodeOutputContentMemberDocument:
		_ = v.Value // Value is document.Interface

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ document.Interface

func ExampleFunctionSchema_outputUsage() {
	var union types.FunctionSchema
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.FunctionSchemaMemberFunctions:
		_ = v.Value // Value is []types.FunctionDefinition

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ []types.FunctionDefinition

func ExampleInlineAgentResponseStream_outputUsage() {
	var union types.InlineAgentResponseStream
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.InlineAgentResponseStreamMemberChunk:
		_ = v.Value // Value is types.InlineAgentPayloadPart

	case *types.InlineAgentResponseStreamMemberFiles:
		_ = v.Value // Value is types.InlineAgentFilePart

	case *types.InlineAgentResponseStreamMemberReturnControl:
		_ = v.Value // Value is types.InlineAgentReturnControlPayload

	case *types.InlineAgentResponseStreamMemberTrace:
		_ = v.Value // Value is types.InlineAgentTracePart

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.InlineAgentTracePart
var _ *types.InlineAgentPayloadPart
var _ *types.InlineAgentReturnControlPayload
var _ *types.InlineAgentFilePart

func ExampleInputPrompt_outputUsage() {
	var union types.InputPrompt
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.InputPromptMemberTextPrompt:
		_ = v.Value // Value is types.TextPrompt

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.TextPrompt

func ExampleInvocationInputMember_outputUsage() {
	var union types.InvocationInputMember
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.InvocationInputMemberMemberApiInvocationInput:
		_ = v.Value // Value is types.ApiInvocationInput

	case *types.InvocationInputMemberMemberFunctionInvocationInput:
		_ = v.Value // Value is types.FunctionInvocationInput

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.ApiInvocationInput
var _ *types.FunctionInvocationInput

func ExampleInvocationResultMember_outputUsage() {
	var union types.InvocationResultMember
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.InvocationResultMemberMemberApiResult:
		_ = v.Value // Value is types.ApiResult

	case *types.InvocationResultMemberMemberFunctionResult:
		_ = v.Value // Value is types.FunctionResult

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.ApiResult
var _ *types.FunctionResult

func ExampleMemory_outputUsage() {
	var union types.Memory
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.MemoryMemberSessionSummary:
		_ = v.Value // Value is types.MemorySessionSummary

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.MemorySessionSummary

func ExampleOptimizedPrompt_outputUsage() {
	var union types.OptimizedPrompt
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.OptimizedPromptMemberTextPrompt:
		_ = v.Value // Value is types.TextPrompt

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.TextPrompt

func ExampleOptimizedPromptStream_outputUsage() {
	var union types.OptimizedPromptStream
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.OptimizedPromptStreamMemberAnalyzePromptEvent:
		_ = v.Value // Value is types.AnalyzePromptEvent

	case *types.OptimizedPromptStreamMemberOptimizedPromptEvent:
		_ = v.Value // Value is types.OptimizedPromptEvent

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.OptimizedPromptEvent
var _ *types.AnalyzePromptEvent

func ExampleOrchestrationTrace_outputUsage() {
	var union types.OrchestrationTrace
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.OrchestrationTraceMemberInvocationInput:
		_ = v.Value // Value is types.InvocationInput

	case *types.OrchestrationTraceMemberModelInvocationInput:
		_ = v.Value // Value is types.ModelInvocationInput

	case *types.OrchestrationTraceMemberModelInvocationOutput:
		_ = v.Value // Value is types.OrchestrationModelInvocationOutput

	case *types.OrchestrationTraceMemberObservation:
		_ = v.Value // Value is types.Observation

	case *types.OrchestrationTraceMemberRationale:
		_ = v.Value // Value is types.Rationale

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.InvocationInput
var _ *types.ModelInvocationInput
var _ *types.OrchestrationModelInvocationOutput
var _ *types.Rationale
var _ *types.Observation

func ExamplePostProcessingTrace_outputUsage() {
	var union types.PostProcessingTrace
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.PostProcessingTraceMemberModelInvocationInput:
		_ = v.Value // Value is types.ModelInvocationInput

	case *types.PostProcessingTraceMemberModelInvocationOutput:
		_ = v.Value // Value is types.PostProcessingModelInvocationOutput

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.ModelInvocationInput
var _ *types.PostProcessingModelInvocationOutput

func ExamplePreProcessingTrace_outputUsage() {
	var union types.PreProcessingTrace
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.PreProcessingTraceMemberModelInvocationInput:
		_ = v.Value // Value is types.ModelInvocationInput

	case *types.PreProcessingTraceMemberModelInvocationOutput:
		_ = v.Value // Value is types.PreProcessingModelInvocationOutput

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.ModelInvocationInput
var _ *types.PreProcessingModelInvocationOutput

func ExampleRerankingMetadataSelectiveModeConfiguration_outputUsage() {
	var union types.RerankingMetadataSelectiveModeConfiguration
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.RerankingMetadataSelectiveModeConfigurationMemberFieldsToExclude:
		_ = v.Value // Value is []types.FieldForReranking

	case *types.RerankingMetadataSelectiveModeConfigurationMemberFieldsToInclude:
		_ = v.Value // Value is []types.FieldForReranking

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ []types.FieldForReranking

func ExampleResponseStream_outputUsage() {
	var union types.ResponseStream
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.ResponseStreamMemberChunk:
		_ = v.Value // Value is types.PayloadPart

	case *types.ResponseStreamMemberFiles:
		_ = v.Value // Value is types.FilePart

	case *types.ResponseStreamMemberReturnControl:
		_ = v.Value // Value is types.ReturnControlPayload

	case *types.ResponseStreamMemberTrace:
		_ = v.Value // Value is types.TracePart

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.ReturnControlPayload
var _ *types.PayloadPart
var _ *types.TracePart
var _ *types.FilePart

func ExampleRetrievalFilter_outputUsage() {
	var union types.RetrievalFilter
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.RetrievalFilterMemberAndAll:
		_ = v.Value // Value is []types.RetrievalFilter

	case *types.RetrievalFilterMemberEquals:
		_ = v.Value // Value is types.FilterAttribute

	case *types.RetrievalFilterMemberGreaterThan:
		_ = v.Value // Value is types.FilterAttribute

	case *types.RetrievalFilterMemberGreaterThanOrEquals:
		_ = v.Value // Value is types.FilterAttribute

	case *types.RetrievalFilterMemberIn:
		_ = v.Value // Value is types.FilterAttribute

	case *types.RetrievalFilterMemberLessThan:
		_ = v.Value // Value is types.FilterAttribute

	case *types.RetrievalFilterMemberLessThanOrEquals:
		_ = v.Value // Value is types.FilterAttribute

	case *types.RetrievalFilterMemberListContains:
		_ = v.Value // Value is types.FilterAttribute

	case *types.RetrievalFilterMemberNotEquals:
		_ = v.Value // Value is types.FilterAttribute

	case *types.RetrievalFilterMemberNotIn:
		_ = v.Value // Value is types.FilterAttribute

	case *types.RetrievalFilterMemberOrAll:
		_ = v.Value // Value is []types.RetrievalFilter

	case *types.RetrievalFilterMemberStartsWith:
		_ = v.Value // Value is types.FilterAttribute

	case *types.RetrievalFilterMemberStringContains:
		_ = v.Value // Value is types.FilterAttribute

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ []types.RetrievalFilter
var _ *types.FilterAttribute

func ExampleRetrieveAndGenerateStreamResponseOutput_outputUsage() {
	var union types.RetrieveAndGenerateStreamResponseOutput
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.RetrieveAndGenerateStreamResponseOutputMemberCitation:
		_ = v.Value // Value is types.CitationEvent

	case *types.RetrieveAndGenerateStreamResponseOutputMemberGuardrail:
		_ = v.Value // Value is types.GuardrailEvent

	case *types.RetrieveAndGenerateStreamResponseOutputMemberOutput:
		_ = v.Value // Value is types.RetrieveAndGenerateOutputEvent

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.GuardrailEvent
var _ *types.CitationEvent
var _ *types.RetrieveAndGenerateOutputEvent

func ExampleRoutingClassifierTrace_outputUsage() {
	var union types.RoutingClassifierTrace
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.RoutingClassifierTraceMemberInvocationInput:
		_ = v.Value // Value is types.InvocationInput

	case *types.RoutingClassifierTraceMemberModelInvocationInput:
		_ = v.Value // Value is types.ModelInvocationInput

	case *types.RoutingClassifierTraceMemberModelInvocationOutput:
		_ = v.Value // Value is types.RoutingClassifierModelInvocationOutput

	case *types.RoutingClassifierTraceMemberObservation:
		_ = v.Value // Value is types.Observation

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.RoutingClassifierModelInvocationOutput
var _ *types.InvocationInput
var _ *types.ModelInvocationInput
var _ *types.Observation

func ExampleTrace_outputUsage() {
	var union types.Trace
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.TraceMemberCustomOrchestrationTrace:
		_ = v.Value // Value is types.CustomOrchestrationTrace

	case *types.TraceMemberFailureTrace:
		_ = v.Value // Value is types.FailureTrace

	case *types.TraceMemberGuardrailTrace:
		_ = v.Value // Value is types.GuardrailTrace

	case *types.TraceMemberOrchestrationTrace:
		_ = v.Value // Value is types.OrchestrationTrace

	case *types.TraceMemberPostProcessingTrace:
		_ = v.Value // Value is types.PostProcessingTrace

	case *types.TraceMemberPreProcessingTrace:
		_ = v.Value // Value is types.PreProcessingTrace

	case *types.TraceMemberRoutingClassifierTrace:
		_ = v.Value // Value is types.RoutingClassifierTrace

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.FailureTrace
var _ *types.GuardrailTrace
var _ *types.CustomOrchestrationTrace
var _ types.PreProcessingTrace
var _ types.PostProcessingTrace
var _ types.OrchestrationTrace
var _ types.RoutingClassifierTrace
