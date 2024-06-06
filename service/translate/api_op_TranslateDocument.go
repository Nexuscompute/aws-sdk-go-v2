// Code generated by smithy-go-codegen DO NOT EDIT.

package translate

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/translate/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Translates the input document from the source language to the target language.
// This synchronous operation supports text, HTML, or Word documents as the input
// document.
//
// TranslateDocument supports translations from English to any supported language,
// and from any supported language to English. Therefore, specify either the source
// language code or the target language code as “en” (English).
//
// If you set the Formality parameter, the request will fail if the target
// language does not support formality. For a list of target languages that support
// formality, see [Setting formality].
//
// [Setting formality]: https://docs.aws.amazon.com/translate/latest/dg/customizing-translations-formality.html
func (c *Client) TranslateDocument(ctx context.Context, params *TranslateDocumentInput, optFns ...func(*Options)) (*TranslateDocumentOutput, error) {
	if params == nil {
		params = &TranslateDocumentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "TranslateDocument", params, optFns, c.addOperationTranslateDocumentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*TranslateDocumentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type TranslateDocumentInput struct {

	// The content and content type for the document to be translated. The document
	// size must not exceed 100 KB.
	//
	// This member is required.
	Document *types.Document

	// The language code for the language of the source text. For a list of supported
	// language codes, see [Supported languages].
	//
	// To have Amazon Translate determine the source language of your text, you can
	// specify auto in the SourceLanguageCode field. If you specify auto , Amazon
	// Translate will call [Amazon Comprehend]to determine the source language.
	//
	// If you specify auto , you must send the TranslateDocument request in a region
	// that supports Amazon Comprehend. Otherwise, the request returns an error
	// indicating that autodetect is not supported.
	//
	// [Supported languages]: https://docs.aws.amazon.com/translate/latest/dg/what-is-languages.html
	// [Amazon Comprehend]: https://docs.aws.amazon.com/comprehend/latest/dg/comprehend-general.html
	//
	// This member is required.
	SourceLanguageCode *string

	// The language code requested for the translated document. For a list of
	// supported language codes, see [Supported languages].
	//
	// [Supported languages]: https://docs.aws.amazon.com/translate/latest/dg/what-is-languages.html
	//
	// This member is required.
	TargetLanguageCode *string

	// Settings to configure your translation output. You can configure the following
	// options:
	//
	//   - Brevity: not supported.
	//
	//   - Formality: sets the formality level of the output text.
	//
	//   - Profanity: masks profane words and phrases in your translation output.
	Settings *types.TranslationSettings

	// The name of a terminology list file to add to the translation job. This file
	// provides source terms and the desired translation for each term. A terminology
	// list can contain a maximum of 256 terms. You can use one custom terminology
	// resource in your translation request.
	//
	// Use the ListTerminologies operation to get the available terminology lists.
	//
	// For more information about custom terminology lists, see [Custom terminology].
	//
	// [Custom terminology]: https://docs.aws.amazon.com/translate/latest/dg/how-custom-terminology.html
	TerminologyNames []string

	noSmithyDocumentSerde
}

type TranslateDocumentOutput struct {

	// The language code of the source document.
	//
	// This member is required.
	SourceLanguageCode *string

	// The language code of the translated document.
	//
	// This member is required.
	TargetLanguageCode *string

	// The document containing the translated content. The document format matches the
	// source document format.
	//
	// This member is required.
	TranslatedDocument *types.TranslatedDocument

	// Settings to configure your translation output. You can configure the following
	// options:
	//
	//   - Brevity: reduces the length of the translation output for most
	//   translations. Available for TranslateText only.
	//
	//   - Formality: sets the formality level of the translation output.
	//
	//   - Profanity: masks profane words and phrases in the translation output.
	AppliedSettings *types.TranslationSettings

	// The names of the custom terminologies applied to the input text by Amazon
	// Translate to produce the translated text document.
	AppliedTerminologies []types.AppliedTerminology

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationTranslateDocumentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpTranslateDocument{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpTranslateDocument{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "TranslateDocument"); err != nil {
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
	if err = addOpTranslateDocumentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opTranslateDocument(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opTranslateDocument(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "TranslateDocument",
	}
}
