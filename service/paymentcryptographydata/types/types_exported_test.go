// Code generated by smithy-go-codegen DO NOT EDIT.

package types_test

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/paymentcryptographydata/types"
)

func ExampleCardGenerationAttributes_outputUsage() {
	var union types.CardGenerationAttributes
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.CardGenerationAttributesMemberAmexCardSecurityCodeVersion1:
		_ = v.Value // Value is types.AmexCardSecurityCodeVersion1

	case *types.CardGenerationAttributesMemberAmexCardSecurityCodeVersion2:
		_ = v.Value // Value is types.AmexCardSecurityCodeVersion2

	case *types.CardGenerationAttributesMemberCardHolderVerificationValue:
		_ = v.Value // Value is types.CardHolderVerificationValue

	case *types.CardGenerationAttributesMemberCardVerificationValue1:
		_ = v.Value // Value is types.CardVerificationValue1

	case *types.CardGenerationAttributesMemberCardVerificationValue2:
		_ = v.Value // Value is types.CardVerificationValue2

	case *types.CardGenerationAttributesMemberDynamicCardVerificationCode:
		_ = v.Value // Value is types.DynamicCardVerificationCode

	case *types.CardGenerationAttributesMemberDynamicCardVerificationValue:
		_ = v.Value // Value is types.DynamicCardVerificationValue

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.AmexCardSecurityCodeVersion2
var _ *types.AmexCardSecurityCodeVersion1
var _ *types.CardHolderVerificationValue
var _ *types.DynamicCardVerificationCode
var _ *types.CardVerificationValue2
var _ *types.CardVerificationValue1
var _ *types.DynamicCardVerificationValue

func ExampleCardVerificationAttributes_outputUsage() {
	var union types.CardVerificationAttributes
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.CardVerificationAttributesMemberAmexCardSecurityCodeVersion1:
		_ = v.Value // Value is types.AmexCardSecurityCodeVersion1

	case *types.CardVerificationAttributesMemberAmexCardSecurityCodeVersion2:
		_ = v.Value // Value is types.AmexCardSecurityCodeVersion2

	case *types.CardVerificationAttributesMemberCardHolderVerificationValue:
		_ = v.Value // Value is types.CardHolderVerificationValue

	case *types.CardVerificationAttributesMemberCardVerificationValue1:
		_ = v.Value // Value is types.CardVerificationValue1

	case *types.CardVerificationAttributesMemberCardVerificationValue2:
		_ = v.Value // Value is types.CardVerificationValue2

	case *types.CardVerificationAttributesMemberDiscoverDynamicCardVerificationCode:
		_ = v.Value // Value is types.DiscoverDynamicCardVerificationCode

	case *types.CardVerificationAttributesMemberDynamicCardVerificationCode:
		_ = v.Value // Value is types.DynamicCardVerificationCode

	case *types.CardVerificationAttributesMemberDynamicCardVerificationValue:
		_ = v.Value // Value is types.DynamicCardVerificationValue

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.AmexCardSecurityCodeVersion2
var _ *types.AmexCardSecurityCodeVersion1
var _ *types.CardHolderVerificationValue
var _ *types.DynamicCardVerificationCode
var _ *types.DiscoverDynamicCardVerificationCode
var _ *types.CardVerificationValue2
var _ *types.CardVerificationValue1
var _ *types.DynamicCardVerificationValue

func ExampleCryptogramAuthResponse_outputUsage() {
	var union types.CryptogramAuthResponse
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.CryptogramAuthResponseMemberArpcMethod1:
		_ = v.Value // Value is types.CryptogramVerificationArpcMethod1

	case *types.CryptogramAuthResponseMemberArpcMethod2:
		_ = v.Value // Value is types.CryptogramVerificationArpcMethod2

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.CryptogramVerificationArpcMethod2
var _ *types.CryptogramVerificationArpcMethod1

func ExampleDerivationMethodAttributes_outputUsage() {
	var union types.DerivationMethodAttributes
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.DerivationMethodAttributesMemberAmex:
		_ = v.Value // Value is types.AmexAttributes

	case *types.DerivationMethodAttributesMemberEmv2000:
		_ = v.Value // Value is types.Emv2000Attributes

	case *types.DerivationMethodAttributesMemberEmvCommon:
		_ = v.Value // Value is types.EmvCommonAttributes

	case *types.DerivationMethodAttributesMemberMastercard:
		_ = v.Value // Value is types.MasterCardAttributes

	case *types.DerivationMethodAttributesMemberVisa:
		_ = v.Value // Value is types.VisaAttributes

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.Emv2000Attributes
var _ *types.MasterCardAttributes
var _ *types.AmexAttributes
var _ *types.EmvCommonAttributes
var _ *types.VisaAttributes

func ExampleEncryptionDecryptionAttributes_outputUsage() {
	var union types.EncryptionDecryptionAttributes
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.EncryptionDecryptionAttributesMemberAsymmetric:
		_ = v.Value // Value is types.AsymmetricEncryptionAttributes

	case *types.EncryptionDecryptionAttributesMemberDukpt:
		_ = v.Value // Value is types.DukptEncryptionAttributes

	case *types.EncryptionDecryptionAttributesMemberEmv:
		_ = v.Value // Value is types.EmvEncryptionAttributes

	case *types.EncryptionDecryptionAttributesMemberSymmetric:
		_ = v.Value // Value is types.SymmetricEncryptionAttributes

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.EmvEncryptionAttributes
var _ *types.SymmetricEncryptionAttributes
var _ *types.AsymmetricEncryptionAttributes
var _ *types.DukptEncryptionAttributes

func ExampleMacAttributes_outputUsage() {
	var union types.MacAttributes
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.MacAttributesMemberAlgorithm:
		_ = v.Value // Value is types.MacAlgorithm

	case *types.MacAttributesMemberDukptCmac:
		_ = v.Value // Value is types.MacAlgorithmDukpt

	case *types.MacAttributesMemberDukptIso9797Algorithm1:
		_ = v.Value // Value is types.MacAlgorithmDukpt

	case *types.MacAttributesMemberDukptIso9797Algorithm3:
		_ = v.Value // Value is types.MacAlgorithmDukpt

	case *types.MacAttributesMemberEmvMac:
		_ = v.Value // Value is types.MacAlgorithmEmv

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.MacAlgorithmDukpt
var _ types.MacAlgorithm
var _ *types.MacAlgorithmEmv

func ExamplePinData_outputUsage() {
	var union types.PinData
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.PinDataMemberPinOffset:
		_ = v.Value // Value is string

	case *types.PinDataMemberVerificationValue:
		_ = v.Value // Value is string

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *string
var _ *string

func ExamplePinGenerationAttributes_outputUsage() {
	var union types.PinGenerationAttributes
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.PinGenerationAttributesMemberIbm3624NaturalPin:
		_ = v.Value // Value is types.Ibm3624NaturalPin

	case *types.PinGenerationAttributesMemberIbm3624PinFromOffset:
		_ = v.Value // Value is types.Ibm3624PinFromOffset

	case *types.PinGenerationAttributesMemberIbm3624PinOffset:
		_ = v.Value // Value is types.Ibm3624PinOffset

	case *types.PinGenerationAttributesMemberIbm3624RandomPin:
		_ = v.Value // Value is types.Ibm3624RandomPin

	case *types.PinGenerationAttributesMemberVisaPin:
		_ = v.Value // Value is types.VisaPin

	case *types.PinGenerationAttributesMemberVisaPinVerificationValue:
		_ = v.Value // Value is types.VisaPinVerificationValue

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.Ibm3624RandomPin
var _ *types.Ibm3624PinOffset
var _ *types.VisaPinVerificationValue
var _ *types.VisaPin
var _ *types.Ibm3624NaturalPin
var _ *types.Ibm3624PinFromOffset

func ExamplePinVerificationAttributes_outputUsage() {
	var union types.PinVerificationAttributes
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.PinVerificationAttributesMemberIbm3624Pin:
		_ = v.Value // Value is types.Ibm3624PinVerification

	case *types.PinVerificationAttributesMemberVisaPin:
		_ = v.Value // Value is types.VisaPinVerification

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.Ibm3624PinVerification
var _ *types.VisaPinVerification

func ExampleReEncryptionAttributes_outputUsage() {
	var union types.ReEncryptionAttributes
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.ReEncryptionAttributesMemberDukpt:
		_ = v.Value // Value is types.DukptEncryptionAttributes

	case *types.ReEncryptionAttributesMemberSymmetric:
		_ = v.Value // Value is types.SymmetricEncryptionAttributes

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.SymmetricEncryptionAttributes
var _ *types.DukptEncryptionAttributes

func ExampleSessionKeyDerivation_outputUsage() {
	var union types.SessionKeyDerivation
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.SessionKeyDerivationMemberAmex:
		_ = v.Value // Value is types.SessionKeyAmex

	case *types.SessionKeyDerivationMemberEmv2000:
		_ = v.Value // Value is types.SessionKeyEmv2000

	case *types.SessionKeyDerivationMemberEmvCommon:
		_ = v.Value // Value is types.SessionKeyEmvCommon

	case *types.SessionKeyDerivationMemberMastercard:
		_ = v.Value // Value is types.SessionKeyMastercard

	case *types.SessionKeyDerivationMemberVisa:
		_ = v.Value // Value is types.SessionKeyVisa

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.SessionKeyVisa
var _ *types.SessionKeyEmvCommon
var _ *types.SessionKeyEmv2000
var _ *types.SessionKeyAmex
var _ *types.SessionKeyMastercard

func ExampleSessionKeyDerivationValue_outputUsage() {
	var union types.SessionKeyDerivationValue
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.SessionKeyDerivationValueMemberApplicationCryptogram:
		_ = v.Value // Value is string

	case *types.SessionKeyDerivationValueMemberApplicationTransactionCounter:
		_ = v.Value // Value is string

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *string
var _ *string

func ExampleTranslationIsoFormats_outputUsage() {
	var union types.TranslationIsoFormats
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.TranslationIsoFormatsMemberIsoFormat0:
		_ = v.Value // Value is types.TranslationPinDataIsoFormat034

	case *types.TranslationIsoFormatsMemberIsoFormat1:
		_ = v.Value // Value is types.TranslationPinDataIsoFormat1

	case *types.TranslationIsoFormatsMemberIsoFormat3:
		_ = v.Value // Value is types.TranslationPinDataIsoFormat034

	case *types.TranslationIsoFormatsMemberIsoFormat4:
		_ = v.Value // Value is types.TranslationPinDataIsoFormat034

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.TranslationPinDataIsoFormat1
var _ *types.TranslationPinDataIsoFormat034

func ExampleWrappedKeyMaterial_outputUsage() {
	var union types.WrappedKeyMaterial
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.WrappedKeyMaterialMemberDiffieHellmanSymmetricKey:
		_ = v.Value // Value is types.EcdhDerivationAttributes

	case *types.WrappedKeyMaterialMemberTr31KeyBlock:
		_ = v.Value // Value is string

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *string
var _ *types.EcdhDerivationAttributes
