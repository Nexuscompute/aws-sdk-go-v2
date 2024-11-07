// Code generated by smithy-go-codegen DO NOT EDIT.

package freetier

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/freetier/types"
	smithy "github.com/aws/smithy-go"
	smithyio "github.com/aws/smithy-go/io"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/ptr"
	smithytime "github.com/aws/smithy-go/time"
	"github.com/aws/smithy-go/tracing"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"io"
	"math"
	"strings"
	"time"
)

func deserializeS3Expires(v string) (*time.Time, error) {
	t, err := smithytime.ParseHTTPDate(v)
	if err != nil {
		return nil, nil
	}
	return &t, nil
}

type awsAwsjson10_deserializeOpGetFreeTierUsage struct {
}

func (*awsAwsjson10_deserializeOpGetFreeTierUsage) ID() string {
	return "OperationDeserializer"
}

func (m *awsAwsjson10_deserializeOpGetFreeTierUsage) HandleDeserialize(ctx context.Context, in middleware.DeserializeInput, next middleware.DeserializeHandler) (
	out middleware.DeserializeOutput, metadata middleware.Metadata, err error,
) {
	out, metadata, err = next.HandleDeserialize(ctx, in)
	if err != nil {
		return out, metadata, err
	}

	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
	defer endTimer()
	defer span.End()
	response, ok := out.RawResponse.(*smithyhttp.Response)
	if !ok {
		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
	}

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return out, metadata, awsAwsjson10_deserializeOpErrorGetFreeTierUsage(response, &metadata)
	}
	output := &GetFreeTierUsageOutput{}
	out.Result = output

	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(response.Body, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	var shape interface{}
	if err := decoder.Decode(&shape); err != nil && err != io.EOF {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return out, metadata, err
	}

	err = awsAwsjson10_deserializeOpDocumentGetFreeTierUsageOutput(&output, shape)
	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return out, metadata, err
	}

	return out, metadata, err
}

func awsAwsjson10_deserializeOpErrorGetFreeTierUsage(response *smithyhttp.Response, metadata *middleware.Metadata) error {
	var errorBuffer bytes.Buffer
	if _, err := io.Copy(&errorBuffer, response.Body); err != nil {
		return &smithy.DeserializationError{Err: fmt.Errorf("failed to copy error response body, %w", err)}
	}
	errorBody := bytes.NewReader(errorBuffer.Bytes())

	errorCode := "UnknownError"
	errorMessage := errorCode

	headerCode := response.Header.Get("X-Amzn-ErrorType")

	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	bodyInfo, err := getProtocolErrorInfo(decoder)
	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	errorBody.Seek(0, io.SeekStart)
	if typ, ok := resolveProtocolErrorType(headerCode, bodyInfo); ok {
		errorCode = restjson.SanitizeErrorCode(typ)
	}
	if len(bodyInfo.Message) != 0 {
		errorMessage = bodyInfo.Message
	}
	switch {
	case strings.EqualFold("InternalServerException", errorCode):
		return awsAwsjson10_deserializeErrorInternalServerException(response, errorBody)

	case strings.EqualFold("ThrottlingException", errorCode):
		return awsAwsjson10_deserializeErrorThrottlingException(response, errorBody)

	case strings.EqualFold("ValidationException", errorCode):
		return awsAwsjson10_deserializeErrorValidationException(response, errorBody)

	default:
		genericError := &smithy.GenericAPIError{
			Code:    errorCode,
			Message: errorMessage,
		}
		return genericError

	}
}

func awsAwsjson10_deserializeErrorInternalServerException(response *smithyhttp.Response, errorBody *bytes.Reader) error {
	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	var shape interface{}
	if err := decoder.Decode(&shape); err != nil && err != io.EOF {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	output := &types.InternalServerException{}
	err := awsAwsjson10_deserializeDocumentInternalServerException(&output, shape)

	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	errorBody.Seek(0, io.SeekStart)
	return output
}

func awsAwsjson10_deserializeErrorThrottlingException(response *smithyhttp.Response, errorBody *bytes.Reader) error {
	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	var shape interface{}
	if err := decoder.Decode(&shape); err != nil && err != io.EOF {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	output := &types.ThrottlingException{}
	err := awsAwsjson10_deserializeDocumentThrottlingException(&output, shape)

	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	errorBody.Seek(0, io.SeekStart)
	return output
}

func awsAwsjson10_deserializeErrorValidationException(response *smithyhttp.Response, errorBody *bytes.Reader) error {
	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	var shape interface{}
	if err := decoder.Decode(&shape); err != nil && err != io.EOF {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	output := &types.ValidationException{}
	err := awsAwsjson10_deserializeDocumentValidationException(&output, shape)

	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	errorBody.Seek(0, io.SeekStart)
	return output
}

func awsAwsjson10_deserializeDocumentFreeTierUsage(v **types.FreeTierUsage, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var sv *types.FreeTierUsage
	if *v == nil {
		sv = &types.FreeTierUsage{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "actualUsageAmount":
			if value != nil {
				switch jtv := value.(type) {
				case json.Number:
					f64, err := jtv.Float64()
					if err != nil {
						return err
					}
					sv.ActualUsageAmount = f64

				case string:
					var f64 float64
					switch {
					case strings.EqualFold(jtv, "NaN"):
						f64 = math.NaN()

					case strings.EqualFold(jtv, "Infinity"):
						f64 = math.Inf(1)

					case strings.EqualFold(jtv, "-Infinity"):
						f64 = math.Inf(-1)

					default:
						return fmt.Errorf("unknown JSON number value: %s", jtv)

					}
					sv.ActualUsageAmount = f64

				default:
					return fmt.Errorf("expected GenericDouble to be a JSON Number, got %T instead", value)

				}
			}

		case "description":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected GenericString to be of type string, got %T instead", value)
				}
				sv.Description = ptr.String(jtv)
			}

		case "forecastedUsageAmount":
			if value != nil {
				switch jtv := value.(type) {
				case json.Number:
					f64, err := jtv.Float64()
					if err != nil {
						return err
					}
					sv.ForecastedUsageAmount = f64

				case string:
					var f64 float64
					switch {
					case strings.EqualFold(jtv, "NaN"):
						f64 = math.NaN()

					case strings.EqualFold(jtv, "Infinity"):
						f64 = math.Inf(1)

					case strings.EqualFold(jtv, "-Infinity"):
						f64 = math.Inf(-1)

					default:
						return fmt.Errorf("unknown JSON number value: %s", jtv)

					}
					sv.ForecastedUsageAmount = f64

				default:
					return fmt.Errorf("expected GenericDouble to be a JSON Number, got %T instead", value)

				}
			}

		case "freeTierType":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected GenericString to be of type string, got %T instead", value)
				}
				sv.FreeTierType = ptr.String(jtv)
			}

		case "limit":
			if value != nil {
				switch jtv := value.(type) {
				case json.Number:
					f64, err := jtv.Float64()
					if err != nil {
						return err
					}
					sv.Limit = f64

				case string:
					var f64 float64
					switch {
					case strings.EqualFold(jtv, "NaN"):
						f64 = math.NaN()

					case strings.EqualFold(jtv, "Infinity"):
						f64 = math.Inf(1)

					case strings.EqualFold(jtv, "-Infinity"):
						f64 = math.Inf(-1)

					default:
						return fmt.Errorf("unknown JSON number value: %s", jtv)

					}
					sv.Limit = f64

				default:
					return fmt.Errorf("expected GenericDouble to be a JSON Number, got %T instead", value)

				}
			}

		case "operation":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected GenericString to be of type string, got %T instead", value)
				}
				sv.Operation = ptr.String(jtv)
			}

		case "region":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected GenericString to be of type string, got %T instead", value)
				}
				sv.Region = ptr.String(jtv)
			}

		case "service":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected GenericString to be of type string, got %T instead", value)
				}
				sv.Service = ptr.String(jtv)
			}

		case "unit":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected GenericString to be of type string, got %T instead", value)
				}
				sv.Unit = ptr.String(jtv)
			}

		case "usageType":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected GenericString to be of type string, got %T instead", value)
				}
				sv.UsageType = ptr.String(jtv)
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsAwsjson10_deserializeDocumentFreeTierUsages(v *[]types.FreeTierUsage, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.([]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var cv []types.FreeTierUsage
	if *v == nil {
		cv = []types.FreeTierUsage{}
	} else {
		cv = *v
	}

	for _, value := range shape {
		var col types.FreeTierUsage
		destAddr := &col
		if err := awsAwsjson10_deserializeDocumentFreeTierUsage(&destAddr, value); err != nil {
			return err
		}
		col = *destAddr
		cv = append(cv, col)

	}
	*v = cv
	return nil
}

func awsAwsjson10_deserializeDocumentInternalServerException(v **types.InternalServerException, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var sv *types.InternalServerException
	if *v == nil {
		sv = &types.InternalServerException{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "message", "Message":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected GenericString to be of type string, got %T instead", value)
				}
				sv.Message = ptr.String(jtv)
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsAwsjson10_deserializeDocumentThrottlingException(v **types.ThrottlingException, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var sv *types.ThrottlingException
	if *v == nil {
		sv = &types.ThrottlingException{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "message", "Message":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected GenericString to be of type string, got %T instead", value)
				}
				sv.Message = ptr.String(jtv)
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsAwsjson10_deserializeDocumentValidationException(v **types.ValidationException, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var sv *types.ValidationException
	if *v == nil {
		sv = &types.ValidationException{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "message", "Message":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected GenericString to be of type string, got %T instead", value)
				}
				sv.Message = ptr.String(jtv)
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsAwsjson10_deserializeOpDocumentGetFreeTierUsageOutput(v **GetFreeTierUsageOutput, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var sv *GetFreeTierUsageOutput
	if *v == nil {
		sv = &GetFreeTierUsageOutput{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "freeTierUsages":
			if err := awsAwsjson10_deserializeDocumentFreeTierUsages(&sv.FreeTierUsages, value); err != nil {
				return err
			}

		case "nextToken":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected NextPageToken to be of type string, got %T instead", value)
				}
				sv.NextToken = ptr.String(jtv)
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

type protocolErrorInfo struct {
	Type    string `json:"__type"`
	Message string
	Code    any // nonstandard for awsjson but some services do present the type here
}

func getProtocolErrorInfo(decoder *json.Decoder) (protocolErrorInfo, error) {
	var errInfo protocolErrorInfo
	if err := decoder.Decode(&errInfo); err != nil {
		if err == io.EOF {
			return errInfo, nil
		}
		return errInfo, err
	}

	return errInfo, nil
}

func resolveProtocolErrorType(headerType string, bodyInfo protocolErrorInfo) (string, bool) {
	if len(headerType) != 0 {
		return headerType, true
	} else if len(bodyInfo.Type) != 0 {
		return bodyInfo.Type, true
	} else if code, ok := bodyInfo.Code.(string); ok && len(code) != 0 {
		return code, true
	}
	return "", false
}
