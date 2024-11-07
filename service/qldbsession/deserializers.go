// Code generated by smithy-go-codegen DO NOT EDIT.

package qldbsession

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/qldbsession/types"
	smithy "github.com/aws/smithy-go"
	smithyio "github.com/aws/smithy-go/io"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/ptr"
	smithytime "github.com/aws/smithy-go/time"
	"github.com/aws/smithy-go/tracing"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"io"
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

type awsAwsjson10_deserializeOpSendCommand struct {
}

func (*awsAwsjson10_deserializeOpSendCommand) ID() string {
	return "OperationDeserializer"
}

func (m *awsAwsjson10_deserializeOpSendCommand) HandleDeserialize(ctx context.Context, in middleware.DeserializeInput, next middleware.DeserializeHandler) (
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
		return out, metadata, awsAwsjson10_deserializeOpErrorSendCommand(response, &metadata)
	}
	output := &SendCommandOutput{}
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

	err = awsAwsjson10_deserializeOpDocumentSendCommandOutput(&output, shape)
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

func awsAwsjson10_deserializeOpErrorSendCommand(response *smithyhttp.Response, metadata *middleware.Metadata) error {
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
	case strings.EqualFold("BadRequestException", errorCode):
		return awsAwsjson10_deserializeErrorBadRequestException(response, errorBody)

	case strings.EqualFold("CapacityExceededException", errorCode):
		return awsAwsjson10_deserializeErrorCapacityExceededException(response, errorBody)

	case strings.EqualFold("InvalidSessionException", errorCode):
		return awsAwsjson10_deserializeErrorInvalidSessionException(response, errorBody)

	case strings.EqualFold("LimitExceededException", errorCode):
		return awsAwsjson10_deserializeErrorLimitExceededException(response, errorBody)

	case strings.EqualFold("OccConflictException", errorCode):
		return awsAwsjson10_deserializeErrorOccConflictException(response, errorBody)

	case strings.EqualFold("RateExceededException", errorCode):
		return awsAwsjson10_deserializeErrorRateExceededException(response, errorBody)

	default:
		genericError := &smithy.GenericAPIError{
			Code:    errorCode,
			Message: errorMessage,
		}
		return genericError

	}
}

func awsAwsjson10_deserializeErrorBadRequestException(response *smithyhttp.Response, errorBody *bytes.Reader) error {
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

	output := &types.BadRequestException{}
	err := awsAwsjson10_deserializeDocumentBadRequestException(&output, shape)

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

func awsAwsjson10_deserializeErrorCapacityExceededException(response *smithyhttp.Response, errorBody *bytes.Reader) error {
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

	output := &types.CapacityExceededException{}
	err := awsAwsjson10_deserializeDocumentCapacityExceededException(&output, shape)

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

func awsAwsjson10_deserializeErrorInvalidSessionException(response *smithyhttp.Response, errorBody *bytes.Reader) error {
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

	output := &types.InvalidSessionException{}
	err := awsAwsjson10_deserializeDocumentInvalidSessionException(&output, shape)

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

func awsAwsjson10_deserializeErrorLimitExceededException(response *smithyhttp.Response, errorBody *bytes.Reader) error {
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

	output := &types.LimitExceededException{}
	err := awsAwsjson10_deserializeDocumentLimitExceededException(&output, shape)

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

func awsAwsjson10_deserializeErrorOccConflictException(response *smithyhttp.Response, errorBody *bytes.Reader) error {
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

	output := &types.OccConflictException{}
	err := awsAwsjson10_deserializeDocumentOccConflictException(&output, shape)

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

func awsAwsjson10_deserializeErrorRateExceededException(response *smithyhttp.Response, errorBody *bytes.Reader) error {
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

	output := &types.RateExceededException{}
	err := awsAwsjson10_deserializeDocumentRateExceededException(&output, shape)

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

func awsAwsjson10_deserializeDocumentAbortTransactionResult(v **types.AbortTransactionResult, value interface{}) error {
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

	var sv *types.AbortTransactionResult
	if *v == nil {
		sv = &types.AbortTransactionResult{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "TimingInformation":
			if err := awsAwsjson10_deserializeDocumentTimingInformation(&sv.TimingInformation, value); err != nil {
				return err
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsAwsjson10_deserializeDocumentBadRequestException(v **types.BadRequestException, value interface{}) error {
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

	var sv *types.BadRequestException
	if *v == nil {
		sv = &types.BadRequestException{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "Code":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected ErrorCode to be of type string, got %T instead", value)
				}
				sv.Code = ptr.String(jtv)
			}

		case "message", "Message":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected ErrorMessage to be of type string, got %T instead", value)
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

func awsAwsjson10_deserializeDocumentCapacityExceededException(v **types.CapacityExceededException, value interface{}) error {
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

	var sv *types.CapacityExceededException
	if *v == nil {
		sv = &types.CapacityExceededException{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "message", "Message":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected ErrorMessage to be of type string, got %T instead", value)
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

func awsAwsjson10_deserializeDocumentCommitTransactionResult(v **types.CommitTransactionResult, value interface{}) error {
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

	var sv *types.CommitTransactionResult
	if *v == nil {
		sv = &types.CommitTransactionResult{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "CommitDigest":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected CommitDigest to be []byte, got %T instead", value)
				}
				dv, err := base64.StdEncoding.DecodeString(jtv)
				if err != nil {
					return fmt.Errorf("failed to base64 decode CommitDigest, %w", err)
				}
				sv.CommitDigest = dv
			}

		case "ConsumedIOs":
			if err := awsAwsjson10_deserializeDocumentIOUsage(&sv.ConsumedIOs, value); err != nil {
				return err
			}

		case "TimingInformation":
			if err := awsAwsjson10_deserializeDocumentTimingInformation(&sv.TimingInformation, value); err != nil {
				return err
			}

		case "TransactionId":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected TransactionId to be of type string, got %T instead", value)
				}
				sv.TransactionId = ptr.String(jtv)
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsAwsjson10_deserializeDocumentEndSessionResult(v **types.EndSessionResult, value interface{}) error {
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

	var sv *types.EndSessionResult
	if *v == nil {
		sv = &types.EndSessionResult{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "TimingInformation":
			if err := awsAwsjson10_deserializeDocumentTimingInformation(&sv.TimingInformation, value); err != nil {
				return err
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsAwsjson10_deserializeDocumentExecuteStatementResult(v **types.ExecuteStatementResult, value interface{}) error {
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

	var sv *types.ExecuteStatementResult
	if *v == nil {
		sv = &types.ExecuteStatementResult{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "ConsumedIOs":
			if err := awsAwsjson10_deserializeDocumentIOUsage(&sv.ConsumedIOs, value); err != nil {
				return err
			}

		case "FirstPage":
			if err := awsAwsjson10_deserializeDocumentPage(&sv.FirstPage, value); err != nil {
				return err
			}

		case "TimingInformation":
			if err := awsAwsjson10_deserializeDocumentTimingInformation(&sv.TimingInformation, value); err != nil {
				return err
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsAwsjson10_deserializeDocumentFetchPageResult(v **types.FetchPageResult, value interface{}) error {
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

	var sv *types.FetchPageResult
	if *v == nil {
		sv = &types.FetchPageResult{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "ConsumedIOs":
			if err := awsAwsjson10_deserializeDocumentIOUsage(&sv.ConsumedIOs, value); err != nil {
				return err
			}

		case "Page":
			if err := awsAwsjson10_deserializeDocumentPage(&sv.Page, value); err != nil {
				return err
			}

		case "TimingInformation":
			if err := awsAwsjson10_deserializeDocumentTimingInformation(&sv.TimingInformation, value); err != nil {
				return err
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsAwsjson10_deserializeDocumentInvalidSessionException(v **types.InvalidSessionException, value interface{}) error {
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

	var sv *types.InvalidSessionException
	if *v == nil {
		sv = &types.InvalidSessionException{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "Code":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected ErrorCode to be of type string, got %T instead", value)
				}
				sv.Code = ptr.String(jtv)
			}

		case "message", "Message":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected ErrorMessage to be of type string, got %T instead", value)
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

func awsAwsjson10_deserializeDocumentIOUsage(v **types.IOUsage, value interface{}) error {
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

	var sv *types.IOUsage
	if *v == nil {
		sv = &types.IOUsage{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "ReadIOs":
			if value != nil {
				jtv, ok := value.(json.Number)
				if !ok {
					return fmt.Errorf("expected ReadIOs to be json.Number, got %T instead", value)
				}
				i64, err := jtv.Int64()
				if err != nil {
					return err
				}
				sv.ReadIOs = i64
			}

		case "WriteIOs":
			if value != nil {
				jtv, ok := value.(json.Number)
				if !ok {
					return fmt.Errorf("expected WriteIOs to be json.Number, got %T instead", value)
				}
				i64, err := jtv.Int64()
				if err != nil {
					return err
				}
				sv.WriteIOs = i64
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsAwsjson10_deserializeDocumentLimitExceededException(v **types.LimitExceededException, value interface{}) error {
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

	var sv *types.LimitExceededException
	if *v == nil {
		sv = &types.LimitExceededException{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "message", "Message":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected ErrorMessage to be of type string, got %T instead", value)
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

func awsAwsjson10_deserializeDocumentOccConflictException(v **types.OccConflictException, value interface{}) error {
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

	var sv *types.OccConflictException
	if *v == nil {
		sv = &types.OccConflictException{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "message", "Message":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected ErrorMessage to be of type string, got %T instead", value)
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

func awsAwsjson10_deserializeDocumentPage(v **types.Page, value interface{}) error {
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

	var sv *types.Page
	if *v == nil {
		sv = &types.Page{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "NextPageToken":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected PageToken to be of type string, got %T instead", value)
				}
				sv.NextPageToken = ptr.String(jtv)
			}

		case "Values":
			if err := awsAwsjson10_deserializeDocumentValueHolders(&sv.Values, value); err != nil {
				return err
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsAwsjson10_deserializeDocumentRateExceededException(v **types.RateExceededException, value interface{}) error {
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

	var sv *types.RateExceededException
	if *v == nil {
		sv = &types.RateExceededException{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "message", "Message":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected ErrorMessage to be of type string, got %T instead", value)
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

func awsAwsjson10_deserializeDocumentStartSessionResult(v **types.StartSessionResult, value interface{}) error {
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

	var sv *types.StartSessionResult
	if *v == nil {
		sv = &types.StartSessionResult{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "SessionToken":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected SessionToken to be of type string, got %T instead", value)
				}
				sv.SessionToken = ptr.String(jtv)
			}

		case "TimingInformation":
			if err := awsAwsjson10_deserializeDocumentTimingInformation(&sv.TimingInformation, value); err != nil {
				return err
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsAwsjson10_deserializeDocumentStartTransactionResult(v **types.StartTransactionResult, value interface{}) error {
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

	var sv *types.StartTransactionResult
	if *v == nil {
		sv = &types.StartTransactionResult{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "TimingInformation":
			if err := awsAwsjson10_deserializeDocumentTimingInformation(&sv.TimingInformation, value); err != nil {
				return err
			}

		case "TransactionId":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected TransactionId to be of type string, got %T instead", value)
				}
				sv.TransactionId = ptr.String(jtv)
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsAwsjson10_deserializeDocumentTimingInformation(v **types.TimingInformation, value interface{}) error {
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

	var sv *types.TimingInformation
	if *v == nil {
		sv = &types.TimingInformation{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "ProcessingTimeMilliseconds":
			if value != nil {
				jtv, ok := value.(json.Number)
				if !ok {
					return fmt.Errorf("expected ProcessingTimeMilliseconds to be json.Number, got %T instead", value)
				}
				i64, err := jtv.Int64()
				if err != nil {
					return err
				}
				sv.ProcessingTimeMilliseconds = i64
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsAwsjson10_deserializeDocumentValueHolder(v **types.ValueHolder, value interface{}) error {
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

	var sv *types.ValueHolder
	if *v == nil {
		sv = &types.ValueHolder{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "IonBinary":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected IonBinary to be []byte, got %T instead", value)
				}
				dv, err := base64.StdEncoding.DecodeString(jtv)
				if err != nil {
					return fmt.Errorf("failed to base64 decode IonBinary, %w", err)
				}
				sv.IonBinary = dv
			}

		case "IonText":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected IonText to be of type string, got %T instead", value)
				}
				sv.IonText = ptr.String(jtv)
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsAwsjson10_deserializeDocumentValueHolders(v *[]types.ValueHolder, value interface{}) error {
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

	var cv []types.ValueHolder
	if *v == nil {
		cv = []types.ValueHolder{}
	} else {
		cv = *v
	}

	for _, value := range shape {
		var col types.ValueHolder
		destAddr := &col
		if err := awsAwsjson10_deserializeDocumentValueHolder(&destAddr, value); err != nil {
			return err
		}
		col = *destAddr
		cv = append(cv, col)

	}
	*v = cv
	return nil
}

func awsAwsjson10_deserializeOpDocumentSendCommandOutput(v **SendCommandOutput, value interface{}) error {
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

	var sv *SendCommandOutput
	if *v == nil {
		sv = &SendCommandOutput{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "AbortTransaction":
			if err := awsAwsjson10_deserializeDocumentAbortTransactionResult(&sv.AbortTransaction, value); err != nil {
				return err
			}

		case "CommitTransaction":
			if err := awsAwsjson10_deserializeDocumentCommitTransactionResult(&sv.CommitTransaction, value); err != nil {
				return err
			}

		case "EndSession":
			if err := awsAwsjson10_deserializeDocumentEndSessionResult(&sv.EndSession, value); err != nil {
				return err
			}

		case "ExecuteStatement":
			if err := awsAwsjson10_deserializeDocumentExecuteStatementResult(&sv.ExecuteStatement, value); err != nil {
				return err
			}

		case "FetchPage":
			if err := awsAwsjson10_deserializeDocumentFetchPageResult(&sv.FetchPage, value); err != nil {
				return err
			}

		case "StartSession":
			if err := awsAwsjson10_deserializeDocumentStartSessionResult(&sv.StartSession, value); err != nil {
				return err
			}

		case "StartTransaction":
			if err := awsAwsjson10_deserializeDocumentStartTransactionResult(&sv.StartTransaction, value); err != nil {
				return err
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
