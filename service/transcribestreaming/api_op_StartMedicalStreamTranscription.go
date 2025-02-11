// Code generated by smithy-go-codegen DO NOT EDIT.

package transcribestreaming

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream/eventstreamapi"
	"github.com/aws/aws-sdk-go-v2/service/transcribestreaming/types"
	"github.com/aws/smithy-go/middleware"
	smithysync "github.com/aws/smithy-go/sync"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"sync"
	"time"
)

// Starts a bidirectional HTTP/2 or WebSocket stream where audio is streamed to
// Amazon Transcribe Medical and the transcription results are streamed to your
// application.
//
// The following parameters are required:
//
//   - language-code
//
//   - media-encoding
//
//   - sample-rate
//
// For more information on streaming with Amazon Transcribe Medical, see [Transcribing streaming audio].
//
// [Transcribing streaming audio]: https://docs.aws.amazon.com/transcribe/latest/dg/streaming.html
func (c *Client) StartMedicalStreamTranscription(ctx context.Context, params *StartMedicalStreamTranscriptionInput, optFns ...func(*Options)) (*StartMedicalStreamTranscriptionOutput, error) {
	if params == nil {
		params = &StartMedicalStreamTranscriptionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartMedicalStreamTranscription", params, optFns, c.addOperationStartMedicalStreamTranscriptionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartMedicalStreamTranscriptionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartMedicalStreamTranscriptionInput struct {

	// Specify the language code that represents the language spoken in your audio.
	//
	// Amazon Transcribe Medical only supports US English ( en-US ).
	//
	// This member is required.
	LanguageCode types.LanguageCode

	// Specify the encoding used for the input audio. Supported formats are:
	//
	//   - FLAC
	//
	//   - OPUS-encoded audio in an Ogg container
	//
	//   - PCM (only signed 16-bit little-endian audio formats, which does not include
	//   WAV)
	//
	// For more information, see [Media formats].
	//
	// [Media formats]: https://docs.aws.amazon.com/transcribe/latest/dg/how-input.html#how-input-audio
	//
	// This member is required.
	MediaEncoding types.MediaEncoding

	// The sample rate of the input audio (in hertz). Amazon Transcribe Medical
	// supports a range from 16,000 Hz to 48,000 Hz. Note that the sample rate you
	// specify must match that of your audio.
	//
	// This member is required.
	MediaSampleRateHertz *int32

	// Specify the medical specialty contained in your audio.
	//
	// This member is required.
	Specialty types.Specialty

	// Specify the type of input audio. For example, choose DICTATION for a provider
	// dictating patient notes and CONVERSATION for a dialogue between a patient and a
	// medical professional.
	//
	// This member is required.
	Type types.Type

	// Labels all personal health information (PHI) identified in your transcript.
	//
	// Content identification is performed at the segment level; PHI is flagged upon
	// complete transcription of an audio segment.
	//
	// For more information, see [Identifying personal health information (PHI) in a transcription].
	//
	// [Identifying personal health information (PHI) in a transcription]: https://docs.aws.amazon.com/transcribe/latest/dg/phi-id.html
	ContentIdentificationType types.MedicalContentIdentificationType

	// Enables channel identification in multi-channel audio.
	//
	// Channel identification transcribes the audio on each channel independently,
	// then appends the output for each channel into one transcript.
	//
	// If you have multi-channel audio and do not enable channel identification, your
	// audio is transcribed in a continuous manner and your transcript is not separated
	// by channel.
	//
	// If you include EnableChannelIdentification in your request, you must also
	// include NumberOfChannels .
	//
	// For more information, see [Transcribing multi-channel audio].
	//
	// [Transcribing multi-channel audio]: https://docs.aws.amazon.com/transcribe/latest/dg/channel-id.html
	EnableChannelIdentification bool

	// Specify the number of channels in your audio stream. This value must be 2 , as
	// only two channels are supported. If your audio doesn't contain multiple
	// channels, do not include this parameter in your request.
	//
	// If you include NumberOfChannels in your request, you must also include
	// EnableChannelIdentification .
	NumberOfChannels *int32

	// Specify a name for your transcription session. If you don't include this
	// parameter in your request, Amazon Transcribe Medical generates an ID and returns
	// it in the response.
	SessionId *string

	// Enables speaker partitioning (diarization) in your transcription output.
	// Speaker partitioning labels the speech from individual speakers in your media
	// file.
	//
	// For more information, see [Partitioning speakers (diarization)].
	//
	// [Partitioning speakers (diarization)]: https://docs.aws.amazon.com/transcribe/latest/dg/diarization.html
	ShowSpeakerLabel bool

	// Specify the name of the custom vocabulary that you want to use when processing
	// your transcription. Note that vocabulary names are case sensitive.
	VocabularyName *string

	noSmithyDocumentSerde
}

type StartMedicalStreamTranscriptionOutput struct {

	// Shows whether content identification was enabled for your transcription.
	ContentIdentificationType types.MedicalContentIdentificationType

	// Shows whether channel identification was enabled for your transcription.
	EnableChannelIdentification bool

	// Provides the language code that you specified in your request. This must be
	// en-US .
	LanguageCode types.LanguageCode

	// Provides the media encoding you specified in your request.
	MediaEncoding types.MediaEncoding

	// Provides the sample rate that you specified in your request.
	MediaSampleRateHertz *int32

	// Provides the number of channels that you specified in your request.
	NumberOfChannels *int32

	// Provides the identifier for your streaming request.
	RequestId *string

	// Provides the identifier for your transcription session.
	SessionId *string

	// Shows whether speaker partitioning was enabled for your transcription.
	ShowSpeakerLabel bool

	// Provides the medical specialty that you specified in your request.
	Specialty types.Specialty

	// Provides the type of audio you specified in your request.
	Type types.Type

	// Provides the name of the custom vocabulary that you specified in your request.
	VocabularyName *string

	eventStream *StartMedicalStreamTranscriptionEventStream

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

// GetStream returns the type to interact with the event stream.
func (o *StartMedicalStreamTranscriptionOutput) GetStream() *StartMedicalStreamTranscriptionEventStream {
	return o.eventStream
}

func (c *Client) addOperationStartMedicalStreamTranscriptionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartMedicalStreamTranscription{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartMedicalStreamTranscription{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "StartMedicalStreamTranscription"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addEventStreamStartMedicalStreamTranscriptionMiddleware(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddRequireMinimumProtocol(stack, 2, 0); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addStreamingEventsPayload(stack); err != nil {
		return err
	}
	if err = addContentSHA256Header(stack); err != nil {
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
	if err = addSpanRetryLoop(stack, options); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = eventstreamapi.AddInitializeStreamWriter(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = addOpStartMedicalStreamTranscriptionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartMedicalStreamTranscription(options.Region), middleware.Before); err != nil {
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
	if err = addSpanInitializeStart(stack); err != nil {
		return err
	}
	if err = addSpanInitializeEnd(stack); err != nil {
		return err
	}
	if err = addSpanBuildRequestStart(stack); err != nil {
		return err
	}
	if err = addSpanBuildRequestEnd(stack); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opStartMedicalStreamTranscription(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "StartMedicalStreamTranscription",
	}
}

// StartMedicalStreamTranscriptionEventStream provides the event stream handling for the StartMedicalStreamTranscription operation.
//
// For testing and mocking the event stream this type should be initialized via
// the NewStartMedicalStreamTranscriptionEventStream constructor function. Using the functional options
// to pass in nested mock behavior.
type StartMedicalStreamTranscriptionEventStream struct {
	// AudioStreamWriter is the EventStream writer for the AudioStream events. This
	// value is automatically set by the SDK when the API call is made Use this member
	// when unit testing your code with the SDK to mock out the EventStream Writer.
	//
	// Must not be nil.
	Writer AudioStreamWriter

	// MedicalTranscriptResultStreamReader is the EventStream reader for the
	// MedicalTranscriptResultStream events. This value is automatically set by the SDK
	// when the API call is made Use this member when unit testing your code with the
	// SDK to mock out the EventStream Reader.
	//
	// Must not be nil.
	Reader MedicalTranscriptResultStreamReader

	done      chan struct{}
	closeOnce sync.Once
	err       *smithysync.OnceErr
}

// NewStartMedicalStreamTranscriptionEventStream initializes an StartMedicalStreamTranscriptionEventStream.
// This function should only be used for testing and mocking the StartMedicalStreamTranscriptionEventStream
// stream within your application.
//
// The Writer member must be set before writing events to the stream.
//
// The Reader member must be set before reading events from the stream.
func NewStartMedicalStreamTranscriptionEventStream(optFns ...func(*StartMedicalStreamTranscriptionEventStream)) *StartMedicalStreamTranscriptionEventStream {
	es := &StartMedicalStreamTranscriptionEventStream{
		done: make(chan struct{}),
		err:  smithysync.NewOnceErr(),
	}
	for _, fn := range optFns {
		fn(es)
	}
	return es
}

// Send writes the event to the stream blocking until the event is written.
// Returns an error if the event was not written.
func (es *StartMedicalStreamTranscriptionEventStream) Send(ctx context.Context, event types.AudioStream) error {
	return es.Writer.Send(ctx, event)
}

// Events returns a channel to read events from.
func (es *StartMedicalStreamTranscriptionEventStream) Events() <-chan types.MedicalTranscriptResultStream {
	return es.Reader.Events()
}

// Close closes the stream. This will also cause the stream to be closed.
// Close must be called when done using the stream API. Not calling Close
// may result in resource leaks.
//
// Will close the underlying EventStream writer and reader, and no more events can be
// sent or received.
func (es *StartMedicalStreamTranscriptionEventStream) Close() error {
	es.closeOnce.Do(es.safeClose)
	return es.Err()
}

func (es *StartMedicalStreamTranscriptionEventStream) safeClose() {
	close(es.done)

	t := time.NewTicker(time.Second)
	defer t.Stop()
	writeCloseDone := make(chan error)
	go func() {
		if err := es.Writer.Close(); err != nil {
			es.err.SetError(err)
		}
		close(writeCloseDone)
	}()
	select {
	case <-t.C:
	case <-writeCloseDone:
	}

	es.Reader.Close()
}

// Err returns any error that occurred while reading or writing EventStream Events
// from the service API's response. Returns nil if there were no errors.
func (es *StartMedicalStreamTranscriptionEventStream) Err() error {
	if err := es.err.Err(); err != nil {
		return err
	}

	if err := es.Writer.Err(); err != nil {
		return err
	}

	if err := es.Reader.Err(); err != nil {
		return err
	}

	return nil
}

func (es *StartMedicalStreamTranscriptionEventStream) waitStreamClose() {
	type errorSet interface {
		ErrorSet() <-chan struct{}
	}

	var inputErrCh <-chan struct{}
	if v, ok := es.Writer.(errorSet); ok {
		inputErrCh = v.ErrorSet()
	}

	var outputErrCh <-chan struct{}
	if v, ok := es.Reader.(errorSet); ok {
		outputErrCh = v.ErrorSet()
	}
	var outputClosedCh <-chan struct{}
	if v, ok := es.Reader.(interface{ Closed() <-chan struct{} }); ok {
		outputClosedCh = v.Closed()
	}

	select {
	case <-es.done:
	case <-inputErrCh:
		es.err.SetError(es.Writer.Err())
		es.Close()

	case <-outputErrCh:
		es.err.SetError(es.Reader.Err())
		es.Close()

	case <-outputClosedCh:
		if err := es.Reader.Err(); err != nil {
			es.err.SetError(es.Reader.Err())
		}
		es.Close()

	}
}
