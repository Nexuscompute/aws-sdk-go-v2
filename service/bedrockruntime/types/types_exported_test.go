// Code generated by smithy-go-codegen DO NOT EDIT.

package types_test

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime/document"
	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime/types"
)

func ExampleAsyncInvokeOutputDataConfig_outputUsage() {
	var union types.AsyncInvokeOutputDataConfig
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.AsyncInvokeOutputDataConfigMemberS3OutputDataConfig:
		_ = v.Value // Value is types.AsyncInvokeS3OutputDataConfig

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.AsyncInvokeS3OutputDataConfig

func ExampleContentBlock_outputUsage() {
	var union types.ContentBlock
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.ContentBlockMemberDocument:
		_ = v.Value // Value is types.DocumentBlock

	case *types.ContentBlockMemberGuardContent:
		_ = v.Value // Value is types.GuardrailConverseContentBlock

	case *types.ContentBlockMemberImage:
		_ = v.Value // Value is types.ImageBlock

	case *types.ContentBlockMemberText:
		_ = v.Value // Value is string

	case *types.ContentBlockMemberToolResult:
		_ = v.Value // Value is types.ToolResultBlock

	case *types.ContentBlockMemberToolUse:
		_ = v.Value // Value is types.ToolUseBlock

	case *types.ContentBlockMemberVideo:
		_ = v.Value // Value is types.VideoBlock

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.DocumentBlock
var _ *string
var _ types.GuardrailConverseContentBlock
var _ *types.ImageBlock
var _ *types.ToolResultBlock
var _ *types.VideoBlock
var _ *types.ToolUseBlock

func ExampleContentBlockDelta_outputUsage() {
	var union types.ContentBlockDelta
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.ContentBlockDeltaMemberText:
		_ = v.Value // Value is string

	case *types.ContentBlockDeltaMemberToolUse:
		_ = v.Value // Value is types.ToolUseBlockDelta

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.ToolUseBlockDelta
var _ *string

func ExampleContentBlockStart_outputUsage() {
	var union types.ContentBlockStart
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.ContentBlockStartMemberToolUse:
		_ = v.Value // Value is types.ToolUseBlockStart

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.ToolUseBlockStart

func ExampleConverseOutput_outputUsage() {
	var union types.ConverseOutput
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.ConverseOutputMemberMessage:
		_ = v.Value // Value is types.Message

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.Message

func ExampleConverseStreamOutput_outputUsage() {
	var union types.ConverseStreamOutput
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.ConverseStreamOutputMemberContentBlockDelta:
		_ = v.Value // Value is types.ContentBlockDeltaEvent

	case *types.ConverseStreamOutputMemberContentBlockStart:
		_ = v.Value // Value is types.ContentBlockStartEvent

	case *types.ConverseStreamOutputMemberContentBlockStop:
		_ = v.Value // Value is types.ContentBlockStopEvent

	case *types.ConverseStreamOutputMemberMessageStart:
		_ = v.Value // Value is types.MessageStartEvent

	case *types.ConverseStreamOutputMemberMessageStop:
		_ = v.Value // Value is types.MessageStopEvent

	case *types.ConverseStreamOutputMemberMetadata:
		_ = v.Value // Value is types.ConverseStreamMetadataEvent

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.MessageStartEvent
var _ *types.ContentBlockStopEvent
var _ *types.MessageStopEvent
var _ *types.ContentBlockDeltaEvent
var _ *types.ContentBlockStartEvent
var _ *types.ConverseStreamMetadataEvent

func ExampleDocumentSource_outputUsage() {
	var union types.DocumentSource
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.DocumentSourceMemberBytes:
		_ = v.Value // Value is []byte

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ []byte

func ExampleGuardrailContentBlock_outputUsage() {
	var union types.GuardrailContentBlock
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.GuardrailContentBlockMemberText:
		_ = v.Value // Value is types.GuardrailTextBlock

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.GuardrailTextBlock

func ExampleGuardrailConverseContentBlock_outputUsage() {
	var union types.GuardrailConverseContentBlock
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.GuardrailConverseContentBlockMemberText:
		_ = v.Value // Value is types.GuardrailConverseTextBlock

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.GuardrailConverseTextBlock

func ExampleImageSource_outputUsage() {
	var union types.ImageSource
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.ImageSourceMemberBytes:
		_ = v.Value // Value is []byte

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ []byte

func ExamplePromptVariableValues_outputUsage() {
	var union types.PromptVariableValues
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.PromptVariableValuesMemberText:
		_ = v.Value // Value is string

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *string

func ExampleResponseStream_outputUsage() {
	var union types.ResponseStream
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.ResponseStreamMemberChunk:
		_ = v.Value // Value is types.PayloadPart

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.PayloadPart

func ExampleSystemContentBlock_outputUsage() {
	var union types.SystemContentBlock
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.SystemContentBlockMemberGuardContent:
		_ = v.Value // Value is types.GuardrailConverseContentBlock

	case *types.SystemContentBlockMemberText:
		_ = v.Value // Value is string

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *string
var _ types.GuardrailConverseContentBlock

func ExampleTool_outputUsage() {
	var union types.Tool
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.ToolMemberToolSpec:
		_ = v.Value // Value is types.ToolSpecification

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.ToolSpecification

func ExampleToolChoice_outputUsage() {
	var union types.ToolChoice
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.ToolChoiceMemberAny:
		_ = v.Value // Value is types.AnyToolChoice

	case *types.ToolChoiceMemberAuto:
		_ = v.Value // Value is types.AutoToolChoice

	case *types.ToolChoiceMemberTool:
		_ = v.Value // Value is types.SpecificToolChoice

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.AutoToolChoice
var _ *types.SpecificToolChoice
var _ *types.AnyToolChoice

func ExampleToolInputSchema_outputUsage() {
	var union types.ToolInputSchema
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.ToolInputSchemaMemberJson:
		_ = v.Value // Value is document.Interface

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ document.Interface

func ExampleToolResultContentBlock_outputUsage() {
	var union types.ToolResultContentBlock
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.ToolResultContentBlockMemberDocument:
		_ = v.Value // Value is types.DocumentBlock

	case *types.ToolResultContentBlockMemberImage:
		_ = v.Value // Value is types.ImageBlock

	case *types.ToolResultContentBlockMemberJson:
		_ = v.Value // Value is document.Interface

	case *types.ToolResultContentBlockMemberText:
		_ = v.Value // Value is string

	case *types.ToolResultContentBlockMemberVideo:
		_ = v.Value // Value is types.VideoBlock

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.DocumentBlock
var _ *string
var _ document.Interface
var _ *types.ImageBlock
var _ *types.VideoBlock

func ExampleVideoSource_outputUsage() {
	var union types.VideoSource
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.VideoSourceMemberBytes:
		_ = v.Value // Value is []byte

	case *types.VideoSourceMemberS3Location:
		_ = v.Value // Value is types.S3Location

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.S3Location
var _ []byte
