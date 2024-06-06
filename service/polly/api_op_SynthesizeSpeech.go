// Code generated by smithy-go-codegen DO NOT EDIT.

package polly

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/polly/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"io"
)

// Synthesizes UTF-8 input, plain text or SSML, to a stream of bytes. SSML input
// must be valid, well-formed SSML. Some alphabets might not be available with all
// the voices (for example, Cyrillic might not be read at all by English voices)
// unless phoneme mapping is used. For more information, see [How it Works].
//
// [How it Works]: https://docs.aws.amazon.com/polly/latest/dg/how-text-to-speech-works.html
func (c *Client) SynthesizeSpeech(ctx context.Context, params *SynthesizeSpeechInput, optFns ...func(*Options)) (*SynthesizeSpeechOutput, error) {
	if params == nil {
		params = &SynthesizeSpeechInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SynthesizeSpeech", params, optFns, c.addOperationSynthesizeSpeechMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SynthesizeSpeechOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SynthesizeSpeechInput struct {

	//  The format in which the returned output will be encoded. For audio stream,
	// this will be mp3, ogg_vorbis, or pcm. For speech marks, this will be json.
	//
	// When pcm is used, the content returned is audio/pcm in a signed 16-bit, 1
	// channel (mono), little-endian format.
	//
	// This member is required.
	OutputFormat types.OutputFormat

	//  Input text to synthesize. If you specify ssml as the TextType , follow the SSML
	// format for the input text.
	//
	// This member is required.
	Text *string

	//  Voice ID to use for the synthesis. You can get a list of available voice IDs
	// by calling the [DescribeVoices]operation.
	//
	// [DescribeVoices]: https://docs.aws.amazon.com/polly/latest/dg/API_DescribeVoices.html
	//
	// This member is required.
	VoiceId types.VoiceId

	// Specifies the engine ( standard , neural , long-form , or generative ) for
	// Amazon Polly to use when processing input text for speech synthesis. Provide an
	// engine that is supported by the voice you select. If you don't provide an
	// engine, the standard engine is selected by default. If a chosen voice isn't
	// supported by the standard engine, this will result in an error. For information
	// on Amazon Polly voices and which voices are available for each engine, see [Available Voices].
	//
	// Type: String
	//
	// Valid Values: standard | neural | long-form | generative
	//
	// Required: Yes
	//
	// [Available Voices]: https://docs.aws.amazon.com/polly/latest/dg/voicelist.html
	Engine types.Engine

	// Optional language code for the Synthesize Speech request. This is only
	// necessary if using a bilingual voice, such as Aditi, which can be used for
	// either Indian English (en-IN) or Hindi (hi-IN).
	//
	// If a bilingual voice is used and no language code is specified, Amazon Polly
	// uses the default language of the bilingual voice. The default language for any
	// voice is the one returned by the [DescribeVoices]operation for the LanguageCode parameter. For
	// example, if no language code is specified, Aditi will use Indian English rather
	// than Hindi.
	//
	// [DescribeVoices]: https://docs.aws.amazon.com/polly/latest/dg/API_DescribeVoices.html
	LanguageCode types.LanguageCode

	// List of one or more pronunciation lexicon names you want the service to apply
	// during synthesis. Lexicons are applied only if the language of the lexicon is
	// the same as the language of the voice. For information about storing lexicons,
	// see [PutLexicon].
	//
	// [PutLexicon]: https://docs.aws.amazon.com/polly/latest/dg/API_PutLexicon.html
	LexiconNames []string

	// The audio frequency specified in Hz.
	//
	// The valid values for mp3 and ogg_vorbis are "8000", "16000", "22050", and
	// "24000". The default value for standard voices is "22050". The default value for
	// neural voices is "24000". The default value for long-form voices is "24000". The
	// default value for generative voices is "24000".
	//
	// Valid values for pcm are "8000" and "16000" The default value is "16000".
	SampleRate *string

	// The type of speech marks returned for the input text.
	SpeechMarkTypes []types.SpeechMarkType

	//  Specifies whether the input text is plain text or SSML. The default value is
	// plain text. For more information, see [Using SSML].
	//
	// [Using SSML]: https://docs.aws.amazon.com/polly/latest/dg/ssml.html
	TextType types.TextType

	noSmithyDocumentSerde
}

type SynthesizeSpeechOutput struct {

	//  Stream containing the synthesized speech.
	AudioStream io.ReadCloser

	//  Specifies the type audio stream. This should reflect the OutputFormat
	// parameter in your request.
	//
	//   - If you request mp3 as the OutputFormat , the ContentType returned is
	//   audio/mpeg.
	//
	//   - If you request ogg_vorbis as the OutputFormat , the ContentType returned is
	//   audio/ogg.
	//
	//   - If you request pcm as the OutputFormat , the ContentType returned is
	//   audio/pcm in a signed 16-bit, 1 channel (mono), little-endian format.
	//
	//   - If you request json as the OutputFormat , the ContentType returned is
	//   application/x-json-stream.
	ContentType *string

	// Number of characters synthesized.
	RequestCharacters int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSynthesizeSpeechMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpSynthesizeSpeech{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpSynthesizeSpeech{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "SynthesizeSpeech"); err != nil {
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
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addOpSynthesizeSpeechValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSynthesizeSpeech(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opSynthesizeSpeech(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "SynthesizeSpeech",
	}
}

// PresignSynthesizeSpeech is used to generate a presigned HTTP Request which
// contains presigned URL, signed headers and HTTP method used.
func (c *PresignClient) PresignSynthesizeSpeech(ctx context.Context, params *SynthesizeSpeechInput, optFns ...func(*PresignOptions)) (*v4.PresignedHTTPRequest, error) {
	if params == nil {
		params = &SynthesizeSpeechInput{}
	}
	options := c.options.copy()
	for _, fn := range optFns {
		fn(&options)
	}
	clientOptFns := append(options.ClientOptions, withNopHTTPClientAPIOption)

	result, _, err := c.client.invokeOperation(ctx, "SynthesizeSpeech", params, clientOptFns,
		c.client.addOperationSynthesizeSpeechMiddlewares,
		presignConverter(options).convertToPresignMiddleware,
	)
	if err != nil {
		return nil, err
	}

	out := result.(*v4.PresignedHTTPRequest)
	return out, nil
}
