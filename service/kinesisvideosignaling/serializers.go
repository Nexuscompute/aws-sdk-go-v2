// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesisvideosignaling

import (
	"bytes"
	"context"
	"fmt"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/encoding/httpbinding"
	smithyjson "github.com/aws/smithy-go/encoding/json"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/tracing"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

type awsRestjson1_serializeOpGetIceServerConfig struct {
}

func (*awsRestjson1_serializeOpGetIceServerConfig) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpGetIceServerConfig) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	_, span := tracing.StartSpan(ctx, "OperationSerializer")
	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
	defer endTimer()
	defer span.End()
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*GetIceServerConfigInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/v1/get-ice-server-config")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "POST"
	var restEncoder *httpbinding.Encoder
	if request.URL.RawPath == "" {
		restEncoder, err = httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	} else {
		request.URL.RawPath = smithyhttp.JoinPath(request.URL.RawPath, opPath)
		restEncoder, err = httpbinding.NewEncoderWithRawPath(request.URL.Path, request.URL.RawPath, request.URL.RawQuery, request.Header)
	}

	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	restEncoder.SetHeader("Content-Type").String("application/json")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsRestjson1_serializeOpDocumentGetIceServerConfigInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	endTimer()
	span.End()
	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsGetIceServerConfigInput(v *GetIceServerConfigInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	return nil
}

func awsRestjson1_serializeOpDocumentGetIceServerConfigInput(v *GetIceServerConfigInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.ChannelARN != nil {
		ok := object.Key("ChannelARN")
		ok.String(*v.ChannelARN)
	}

	if v.ClientId != nil {
		ok := object.Key("ClientId")
		ok.String(*v.ClientId)
	}

	if len(v.Service) > 0 {
		ok := object.Key("Service")
		ok.String(string(v.Service))
	}

	if v.Username != nil {
		ok := object.Key("Username")
		ok.String(*v.Username)
	}

	return nil
}

type awsRestjson1_serializeOpSendAlexaOfferToMaster struct {
}

func (*awsRestjson1_serializeOpSendAlexaOfferToMaster) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpSendAlexaOfferToMaster) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	_, span := tracing.StartSpan(ctx, "OperationSerializer")
	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
	defer endTimer()
	defer span.End()
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*SendAlexaOfferToMasterInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/v1/send-alexa-offer-to-master")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "POST"
	var restEncoder *httpbinding.Encoder
	if request.URL.RawPath == "" {
		restEncoder, err = httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	} else {
		request.URL.RawPath = smithyhttp.JoinPath(request.URL.RawPath, opPath)
		restEncoder, err = httpbinding.NewEncoderWithRawPath(request.URL.Path, request.URL.RawPath, request.URL.RawQuery, request.Header)
	}

	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	restEncoder.SetHeader("Content-Type").String("application/json")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsRestjson1_serializeOpDocumentSendAlexaOfferToMasterInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	endTimer()
	span.End()
	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsSendAlexaOfferToMasterInput(v *SendAlexaOfferToMasterInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	return nil
}

func awsRestjson1_serializeOpDocumentSendAlexaOfferToMasterInput(v *SendAlexaOfferToMasterInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.ChannelARN != nil {
		ok := object.Key("ChannelARN")
		ok.String(*v.ChannelARN)
	}

	if v.MessagePayload != nil {
		ok := object.Key("MessagePayload")
		ok.String(*v.MessagePayload)
	}

	if v.SenderClientId != nil {
		ok := object.Key("SenderClientId")
		ok.String(*v.SenderClientId)
	}

	return nil
}
