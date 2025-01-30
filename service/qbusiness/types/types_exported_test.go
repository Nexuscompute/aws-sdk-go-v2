// Code generated by smithy-go-codegen DO NOT EDIT.

package types_test

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/qbusiness/types"
	"time"
)

func ExampleAPISchema_outputUsage() {
	var union types.APISchema
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.APISchemaMemberPayload:
		_ = v.Value // Value is string

	case *types.APISchemaMemberS3:
		_ = v.Value // Value is types.S3

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *string
var _ *types.S3

func ExampleChatInputStream_outputUsage() {
	var union types.ChatInputStream
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.ChatInputStreamMemberActionExecutionEvent:
		_ = v.Value // Value is types.ActionExecutionEvent

	case *types.ChatInputStreamMemberAttachmentEvent:
		_ = v.Value // Value is types.AttachmentInputEvent

	case *types.ChatInputStreamMemberAuthChallengeResponseEvent:
		_ = v.Value // Value is types.AuthChallengeResponseEvent

	case *types.ChatInputStreamMemberConfigurationEvent:
		_ = v.Value // Value is types.ConfigurationEvent

	case *types.ChatInputStreamMemberEndOfInputEvent:
		_ = v.Value // Value is types.EndOfInputEvent

	case *types.ChatInputStreamMemberTextEvent:
		_ = v.Value // Value is types.TextInputEvent

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.EndOfInputEvent
var _ *types.ActionExecutionEvent
var _ *types.TextInputEvent
var _ *types.AttachmentInputEvent
var _ *types.ConfigurationEvent
var _ *types.AuthChallengeResponseEvent

func ExampleChatModeConfiguration_outputUsage() {
	var union types.ChatModeConfiguration
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.ChatModeConfigurationMemberPluginConfiguration:
		_ = v.Value // Value is types.PluginConfiguration

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.PluginConfiguration

func ExampleChatOutputStream_outputUsage() {
	var union types.ChatOutputStream
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.ChatOutputStreamMemberActionReviewEvent:
		_ = v.Value // Value is types.ActionReviewEvent

	case *types.ChatOutputStreamMemberAuthChallengeRequestEvent:
		_ = v.Value // Value is types.AuthChallengeRequestEvent

	case *types.ChatOutputStreamMemberFailedAttachmentEvent:
		_ = v.Value // Value is types.FailedAttachmentEvent

	case *types.ChatOutputStreamMemberMetadataEvent:
		_ = v.Value // Value is types.MetadataEvent

	case *types.ChatOutputStreamMemberTextEvent:
		_ = v.Value // Value is types.TextOutputEvent

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.AuthChallengeRequestEvent
var _ *types.FailedAttachmentEvent
var _ *types.ActionReviewEvent
var _ *types.MetadataEvent
var _ *types.TextOutputEvent

func ExampleContentSource_outputUsage() {
	var union types.ContentSource
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.ContentSourceMemberRetriever:
		_ = v.Value // Value is types.RetrieverContentSource

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.RetrieverContentSource

func ExampleCopyFromSource_outputUsage() {
	var union types.CopyFromSource
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.CopyFromSourceMemberConversation:
		_ = v.Value // Value is types.ConversationSource

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.ConversationSource

func ExampleDocumentAttributeBoostingConfiguration_outputUsage() {
	var union types.DocumentAttributeBoostingConfiguration
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.DocumentAttributeBoostingConfigurationMemberDateConfiguration:
		_ = v.Value // Value is types.DateAttributeBoostingConfiguration

	case *types.DocumentAttributeBoostingConfigurationMemberNumberConfiguration:
		_ = v.Value // Value is types.NumberAttributeBoostingConfiguration

	case *types.DocumentAttributeBoostingConfigurationMemberStringConfiguration:
		_ = v.Value // Value is types.StringAttributeBoostingConfiguration

	case *types.DocumentAttributeBoostingConfigurationMemberStringListConfiguration:
		_ = v.Value // Value is types.StringListAttributeBoostingConfiguration

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.StringListAttributeBoostingConfiguration
var _ *types.DateAttributeBoostingConfiguration
var _ *types.StringAttributeBoostingConfiguration
var _ *types.NumberAttributeBoostingConfiguration

func ExampleDocumentAttributeValue_outputUsage() {
	var union types.DocumentAttributeValue
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.DocumentAttributeValueMemberDateValue:
		_ = v.Value // Value is time.Time

	case *types.DocumentAttributeValueMemberLongValue:
		_ = v.Value // Value is int64

	case *types.DocumentAttributeValueMemberStringListValue:
		_ = v.Value // Value is []string

	case *types.DocumentAttributeValueMemberStringValue:
		_ = v.Value // Value is string

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *string
var _ *int64
var _ []string
var _ *time.Time

func ExampleDocumentContent_outputUsage() {
	var union types.DocumentContent
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.DocumentContentMemberBlob:
		_ = v.Value // Value is []byte

	case *types.DocumentContentMemberS3:
		_ = v.Value // Value is types.S3

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.S3
var _ []byte

func ExampleIdentityProviderConfiguration_outputUsage() {
	var union types.IdentityProviderConfiguration
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.IdentityProviderConfigurationMemberOpenIDConnectConfiguration:
		_ = v.Value // Value is types.OpenIDConnectProviderConfiguration

	case *types.IdentityProviderConfigurationMemberSamlConfiguration:
		_ = v.Value // Value is types.SamlProviderConfiguration

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.OpenIDConnectProviderConfiguration
var _ *types.SamlProviderConfiguration

func ExamplePluginAuthConfiguration_outputUsage() {
	var union types.PluginAuthConfiguration
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.PluginAuthConfigurationMemberBasicAuthConfiguration:
		_ = v.Value // Value is types.BasicAuthConfiguration

	case *types.PluginAuthConfigurationMemberIdcAuthConfiguration:
		_ = v.Value // Value is types.IdcAuthConfiguration

	case *types.PluginAuthConfigurationMemberNoAuthConfiguration:
		_ = v.Value // Value is types.NoAuthConfiguration

	case *types.PluginAuthConfigurationMemberOAuth2ClientCredentialConfiguration:
		_ = v.Value // Value is types.OAuth2ClientCredentialConfiguration

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.BasicAuthConfiguration
var _ *types.IdcAuthConfiguration
var _ *types.OAuth2ClientCredentialConfiguration
var _ *types.NoAuthConfiguration

func ExamplePrincipal_outputUsage() {
	var union types.Principal
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.PrincipalMemberGroup:
		_ = v.Value // Value is types.PrincipalGroup

	case *types.PrincipalMemberUser:
		_ = v.Value // Value is types.PrincipalUser

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.PrincipalUser
var _ *types.PrincipalGroup

func ExampleRetrieverConfiguration_outputUsage() {
	var union types.RetrieverConfiguration
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.RetrieverConfigurationMemberKendraIndexConfiguration:
		_ = v.Value // Value is types.KendraIndexConfiguration

	case *types.RetrieverConfigurationMemberNativeIndexConfiguration:
		_ = v.Value // Value is types.NativeIndexConfiguration

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.NativeIndexConfiguration
var _ *types.KendraIndexConfiguration

func ExampleRuleConfiguration_outputUsage() {
	var union types.RuleConfiguration
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.RuleConfigurationMemberContentBlockerRule:
		_ = v.Value // Value is types.ContentBlockerRule

	case *types.RuleConfigurationMemberContentRetrievalRule:
		_ = v.Value // Value is types.ContentRetrievalRule

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.ContentRetrievalRule
var _ *types.ContentBlockerRule

func ExampleSubscriptionPrincipal_outputUsage() {
	var union types.SubscriptionPrincipal
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.SubscriptionPrincipalMemberGroup:
		_ = v.Value // Value is string

	case *types.SubscriptionPrincipalMemberUser:
		_ = v.Value // Value is string

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *string
var _ *string

func ExampleWebExperienceAuthConfiguration_outputUsage() {
	var union types.WebExperienceAuthConfiguration
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.WebExperienceAuthConfigurationMemberSamlConfiguration:
		_ = v.Value // Value is types.SamlConfiguration

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.SamlConfiguration
