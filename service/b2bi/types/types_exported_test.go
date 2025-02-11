// Code generated by smithy-go-codegen DO NOT EDIT.

package types_test

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/b2bi/types"
)

func ExampleCapabilityConfiguration_outputUsage() {
	var union types.CapabilityConfiguration
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.CapabilityConfigurationMemberEdi:
		_ = v.Value // Value is types.EdiConfiguration

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.EdiConfiguration

func ExampleConversionTargetFormatDetails_outputUsage() {
	var union types.ConversionTargetFormatDetails
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.ConversionTargetFormatDetailsMemberX12:
		_ = v.Value // Value is types.X12Details

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.X12Details

func ExampleEdiType_outputUsage() {
	var union types.EdiType
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.EdiTypeMemberX12Details:
		_ = v.Value // Value is types.X12Details

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.X12Details

func ExampleFormatOptions_outputUsage() {
	var union types.FormatOptions
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.FormatOptionsMemberX12:
		_ = v.Value // Value is types.X12Details

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.X12Details

func ExampleInputFileSource_outputUsage() {
	var union types.InputFileSource
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.InputFileSourceMemberFileContent:
		_ = v.Value // Value is string

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *string

func ExampleOutboundEdiOptions_outputUsage() {
	var union types.OutboundEdiOptions
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.OutboundEdiOptionsMemberX12:
		_ = v.Value // Value is types.X12Envelope

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.X12Envelope

func ExampleOutputSampleFileSource_outputUsage() {
	var union types.OutputSampleFileSource
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.OutputSampleFileSourceMemberFileLocation:
		_ = v.Value // Value is types.S3Location

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.S3Location

func ExampleTemplateDetails_outputUsage() {
	var union types.TemplateDetails
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.TemplateDetailsMemberX12:
		_ = v.Value // Value is types.X12Details

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.X12Details
