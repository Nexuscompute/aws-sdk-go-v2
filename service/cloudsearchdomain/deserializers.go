// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudsearchdomain

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/cloudsearchdomain/types"
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

type awsRestjson1_deserializeOpSearch struct {
}

func (*awsRestjson1_deserializeOpSearch) ID() string {
	return "OperationDeserializer"
}

func (m *awsRestjson1_deserializeOpSearch) HandleDeserialize(ctx context.Context, in middleware.DeserializeInput, next middleware.DeserializeHandler) (
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
		return out, metadata, awsRestjson1_deserializeOpErrorSearch(response, &metadata)
	}
	output := &SearchOutput{}
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

	err = awsRestjson1_deserializeOpDocumentSearchOutput(&output, shape)
	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		return out, metadata, &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body with invalid JSON, %w", err),
			Snapshot: snapshot.Bytes(),
		}
	}

	span.End()
	return out, metadata, err
}

func awsRestjson1_deserializeOpErrorSearch(response *smithyhttp.Response, metadata *middleware.Metadata) error {
	var errorBuffer bytes.Buffer
	if _, err := io.Copy(&errorBuffer, response.Body); err != nil {
		return &smithy.DeserializationError{Err: fmt.Errorf("failed to copy error response body, %w", err)}
	}
	errorBody := bytes.NewReader(errorBuffer.Bytes())

	errorCode := "UnknownError"
	errorMessage := errorCode

	headerCode := response.Header.Get("X-Amzn-ErrorType")
	if len(headerCode) != 0 {
		errorCode = restjson.SanitizeErrorCode(headerCode)
	}

	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	jsonCode, message, err := restjson.GetErrorInfo(decoder)
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
	if len(headerCode) == 0 && len(jsonCode) != 0 {
		errorCode = restjson.SanitizeErrorCode(jsonCode)
	}
	if len(message) != 0 {
		errorMessage = message
	}

	switch {
	case strings.EqualFold("SearchException", errorCode):
		return awsRestjson1_deserializeErrorSearchException(response, errorBody)

	default:
		genericError := &smithy.GenericAPIError{
			Code:    errorCode,
			Message: errorMessage,
		}
		return genericError

	}
}

func awsRestjson1_deserializeOpDocumentSearchOutput(v **SearchOutput, value interface{}) error {
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

	var sv *SearchOutput
	if *v == nil {
		sv = &SearchOutput{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "facets":
			if err := awsRestjson1_deserializeDocumentFacets(&sv.Facets, value); err != nil {
				return err
			}
		case "hits":
			if err := awsRestjson1_deserializeDocumentHits(&sv.Hits, value); err != nil {
				return err
			}
		case "stats":
			if err := awsRestjson1_deserializeDocumentStats(&sv.Stats, value); err != nil {
				return err
			}
		case "status":
			if err := awsRestjson1_deserializeDocumentSearchStatus(&sv.Status, value); err != nil {
				return err
			}
		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

type awsRestjson1_deserializeOpSuggest struct {
}

func (*awsRestjson1_deserializeOpSuggest) ID() string {
	return "OperationDeserializer"
}

func (m *awsRestjson1_deserializeOpSuggest) HandleDeserialize(ctx context.Context, in middleware.DeserializeInput, next middleware.DeserializeHandler) (
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
		return out, metadata, awsRestjson1_deserializeOpErrorSuggest(response, &metadata)
	}
	output := &SuggestOutput{}
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

	err = awsRestjson1_deserializeOpDocumentSuggestOutput(&output, shape)
	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		return out, metadata, &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body with invalid JSON, %w", err),
			Snapshot: snapshot.Bytes(),
		}
	}

	span.End()
	return out, metadata, err
}

func awsRestjson1_deserializeOpErrorSuggest(response *smithyhttp.Response, metadata *middleware.Metadata) error {
	var errorBuffer bytes.Buffer
	if _, err := io.Copy(&errorBuffer, response.Body); err != nil {
		return &smithy.DeserializationError{Err: fmt.Errorf("failed to copy error response body, %w", err)}
	}
	errorBody := bytes.NewReader(errorBuffer.Bytes())

	errorCode := "UnknownError"
	errorMessage := errorCode

	headerCode := response.Header.Get("X-Amzn-ErrorType")
	if len(headerCode) != 0 {
		errorCode = restjson.SanitizeErrorCode(headerCode)
	}

	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	jsonCode, message, err := restjson.GetErrorInfo(decoder)
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
	if len(headerCode) == 0 && len(jsonCode) != 0 {
		errorCode = restjson.SanitizeErrorCode(jsonCode)
	}
	if len(message) != 0 {
		errorMessage = message
	}

	switch {
	case strings.EqualFold("SearchException", errorCode):
		return awsRestjson1_deserializeErrorSearchException(response, errorBody)

	default:
		genericError := &smithy.GenericAPIError{
			Code:    errorCode,
			Message: errorMessage,
		}
		return genericError

	}
}

func awsRestjson1_deserializeOpDocumentSuggestOutput(v **SuggestOutput, value interface{}) error {
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

	var sv *SuggestOutput
	if *v == nil {
		sv = &SuggestOutput{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "status":
			if err := awsRestjson1_deserializeDocumentSuggestStatus(&sv.Status, value); err != nil {
				return err
			}
		case "suggest":
			if err := awsRestjson1_deserializeDocumentSuggestModel(&sv.Suggest, value); err != nil {
				return err
			}
		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

type awsRestjson1_deserializeOpUploadDocuments struct {
}

func (*awsRestjson1_deserializeOpUploadDocuments) ID() string {
	return "OperationDeserializer"
}

func (m *awsRestjson1_deserializeOpUploadDocuments) HandleDeserialize(ctx context.Context, in middleware.DeserializeInput, next middleware.DeserializeHandler) (
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
		return out, metadata, awsRestjson1_deserializeOpErrorUploadDocuments(response, &metadata)
	}
	output := &UploadDocumentsOutput{}
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

	err = awsRestjson1_deserializeOpDocumentUploadDocumentsOutput(&output, shape)
	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		return out, metadata, &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body with invalid JSON, %w", err),
			Snapshot: snapshot.Bytes(),
		}
	}

	span.End()
	return out, metadata, err
}

func awsRestjson1_deserializeOpErrorUploadDocuments(response *smithyhttp.Response, metadata *middleware.Metadata) error {
	var errorBuffer bytes.Buffer
	if _, err := io.Copy(&errorBuffer, response.Body); err != nil {
		return &smithy.DeserializationError{Err: fmt.Errorf("failed to copy error response body, %w", err)}
	}
	errorBody := bytes.NewReader(errorBuffer.Bytes())

	errorCode := "UnknownError"
	errorMessage := errorCode

	headerCode := response.Header.Get("X-Amzn-ErrorType")
	if len(headerCode) != 0 {
		errorCode = restjson.SanitizeErrorCode(headerCode)
	}

	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	jsonCode, message, err := restjson.GetErrorInfo(decoder)
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
	if len(headerCode) == 0 && len(jsonCode) != 0 {
		errorCode = restjson.SanitizeErrorCode(jsonCode)
	}
	if len(message) != 0 {
		errorMessage = message
	}

	switch {
	case strings.EqualFold("DocumentServiceException", errorCode):
		return awsRestjson1_deserializeErrorDocumentServiceException(response, errorBody)

	default:
		genericError := &smithy.GenericAPIError{
			Code:    errorCode,
			Message: errorMessage,
		}
		return genericError

	}
}

func awsRestjson1_deserializeOpDocumentUploadDocumentsOutput(v **UploadDocumentsOutput, value interface{}) error {
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

	var sv *UploadDocumentsOutput
	if *v == nil {
		sv = &UploadDocumentsOutput{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "adds":
			if value != nil {
				jtv, ok := value.(json.Number)
				if !ok {
					return fmt.Errorf("expected Adds to be json.Number, got %T instead", value)
				}
				i64, err := jtv.Int64()
				if err != nil {
					return err
				}
				sv.Adds = i64
			}
		case "deletes":
			if value != nil {
				jtv, ok := value.(json.Number)
				if !ok {
					return fmt.Errorf("expected Deletes to be json.Number, got %T instead", value)
				}
				i64, err := jtv.Int64()
				if err != nil {
					return err
				}
				sv.Deletes = i64
			}
		case "status":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected String to be of type string, got %T instead", value)
				}
				sv.Status = ptr.String(jtv)
			}
		case "warnings":
			if err := awsRestjson1_deserializeDocumentDocumentServiceWarnings(&sv.Warnings, value); err != nil {
				return err
			}
		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsRestjson1_deserializeErrorDocumentServiceException(response *smithyhttp.Response, errorBody *bytes.Reader) error {
	output := &types.DocumentServiceException{}
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

	err := awsRestjson1_deserializeDocumentDocumentServiceException(&output, shape)

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

func awsRestjson1_deserializeErrorSearchException(response *smithyhttp.Response, errorBody *bytes.Reader) error {
	output := &types.SearchException{}
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

	err := awsRestjson1_deserializeDocumentSearchException(&output, shape)

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

func awsRestjson1_deserializeDocumentBucket(v **types.Bucket, value interface{}) error {
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

	var sv *types.Bucket
	if *v == nil {
		sv = &types.Bucket{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "count":
			if value != nil {
				jtv, ok := value.(json.Number)
				if !ok {
					return fmt.Errorf("expected Long to be json.Number, got %T instead", value)
				}
				i64, err := jtv.Int64()
				if err != nil {
					return err
				}
				sv.Count = i64
			}
		case "value":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected String to be of type string, got %T instead", value)
				}
				sv.Value = ptr.String(jtv)
			}
		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsRestjson1_deserializeDocumentBucketInfo(v **types.BucketInfo, value interface{}) error {
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

	var sv *types.BucketInfo
	if *v == nil {
		sv = &types.BucketInfo{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "buckets":
			if err := awsRestjson1_deserializeDocumentBucketList(&sv.Buckets, value); err != nil {
				return err
			}
		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsRestjson1_deserializeDocumentBucketList(v *[]types.Bucket, value interface{}) error {
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

	var cv []types.Bucket
	if *v == nil {
		cv = []types.Bucket{}
	} else {
		cv = *v
	}

	for _, value := range shape {
		var col types.Bucket
		destAddr := &col
		if err := awsRestjson1_deserializeDocumentBucket(&destAddr, value); err != nil {
			return err
		}
		col = *destAddr
		cv = append(cv, col)

	}
	*v = cv
	return nil
}

func awsRestjson1_deserializeDocumentDocumentServiceException(v **types.DocumentServiceException, value interface{}) error {
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

	var sv *types.DocumentServiceException
	if *v == nil {
		sv = &types.DocumentServiceException{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "message", "Message":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected String to be of type string, got %T instead", value)
				}
				sv.Message = ptr.String(jtv)
			}
		case "status":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected String to be of type string, got %T instead", value)
				}
				sv.Status = ptr.String(jtv)
			}
		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsRestjson1_deserializeDocumentDocumentServiceWarning(v **types.DocumentServiceWarning, value interface{}) error {
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

	var sv *types.DocumentServiceWarning
	if *v == nil {
		sv = &types.DocumentServiceWarning{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "message":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected String to be of type string, got %T instead", value)
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

func awsRestjson1_deserializeDocumentDocumentServiceWarnings(v *[]types.DocumentServiceWarning, value interface{}) error {
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

	var cv []types.DocumentServiceWarning
	if *v == nil {
		cv = []types.DocumentServiceWarning{}
	} else {
		cv = *v
	}

	for _, value := range shape {
		var col types.DocumentServiceWarning
		destAddr := &col
		if err := awsRestjson1_deserializeDocumentDocumentServiceWarning(&destAddr, value); err != nil {
			return err
		}
		col = *destAddr
		cv = append(cv, col)

	}
	*v = cv
	return nil
}

func awsRestjson1_deserializeDocumentExprs(v *map[string]string, value interface{}) error {
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

	var mv map[string]string
	if *v == nil {
		mv = map[string]string{}
	} else {
		mv = *v
	}

	for key, value := range shape {
		var parsedVal string
		if value != nil {
			jtv, ok := value.(string)
			if !ok {
				return fmt.Errorf("expected String to be of type string, got %T instead", value)
			}
			parsedVal = jtv
		}
		mv[key] = parsedVal

	}
	*v = mv
	return nil
}

func awsRestjson1_deserializeDocumentFacets(v *map[string]types.BucketInfo, value interface{}) error {
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

	var mv map[string]types.BucketInfo
	if *v == nil {
		mv = map[string]types.BucketInfo{}
	} else {
		mv = *v
	}

	for key, value := range shape {
		var parsedVal types.BucketInfo
		mapVar := parsedVal
		destAddr := &mapVar
		if err := awsRestjson1_deserializeDocumentBucketInfo(&destAddr, value); err != nil {
			return err
		}
		parsedVal = *destAddr
		mv[key] = parsedVal

	}
	*v = mv
	return nil
}

func awsRestjson1_deserializeDocumentFields(v *map[string][]string, value interface{}) error {
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

	var mv map[string][]string
	if *v == nil {
		mv = map[string][]string{}
	} else {
		mv = *v
	}

	for key, value := range shape {
		var parsedVal []string
		mapVar := parsedVal
		if err := awsRestjson1_deserializeDocumentFieldValue(&mapVar, value); err != nil {
			return err
		}
		parsedVal = mapVar
		mv[key] = parsedVal

	}
	*v = mv
	return nil
}

func awsRestjson1_deserializeDocumentFieldStats(v **types.FieldStats, value interface{}) error {
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

	var sv *types.FieldStats
	if *v == nil {
		sv = &types.FieldStats{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "count":
			if value != nil {
				jtv, ok := value.(json.Number)
				if !ok {
					return fmt.Errorf("expected Long to be json.Number, got %T instead", value)
				}
				i64, err := jtv.Int64()
				if err != nil {
					return err
				}
				sv.Count = i64
			}
		case "max":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected String to be of type string, got %T instead", value)
				}
				sv.Max = ptr.String(jtv)
			}
		case "mean":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected String to be of type string, got %T instead", value)
				}
				sv.Mean = ptr.String(jtv)
			}
		case "min":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected String to be of type string, got %T instead", value)
				}
				sv.Min = ptr.String(jtv)
			}
		case "missing":
			if value != nil {
				jtv, ok := value.(json.Number)
				if !ok {
					return fmt.Errorf("expected Long to be json.Number, got %T instead", value)
				}
				i64, err := jtv.Int64()
				if err != nil {
					return err
				}
				sv.Missing = i64
			}
		case "stddev":
			if value != nil {
				switch jtv := value.(type) {
				case json.Number:
					f64, err := jtv.Float64()
					if err != nil {
						return err
					}
					sv.Stddev = f64

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
					sv.Stddev = f64

				default:
					return fmt.Errorf("expected Double to be a JSON Number, got %T instead", value)

				}
			}
		case "sum":
			if value != nil {
				switch jtv := value.(type) {
				case json.Number:
					f64, err := jtv.Float64()
					if err != nil {
						return err
					}
					sv.Sum = f64

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
					sv.Sum = f64

				default:
					return fmt.Errorf("expected Double to be a JSON Number, got %T instead", value)

				}
			}
		case "sumOfSquares":
			if value != nil {
				switch jtv := value.(type) {
				case json.Number:
					f64, err := jtv.Float64()
					if err != nil {
						return err
					}
					sv.SumOfSquares = f64

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
					sv.SumOfSquares = f64

				default:
					return fmt.Errorf("expected Double to be a JSON Number, got %T instead", value)

				}
			}
		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsRestjson1_deserializeDocumentFieldValue(v *[]string, value interface{}) error {
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

	var cv []string
	if *v == nil {
		cv = []string{}
	} else {
		cv = *v
	}

	for _, value := range shape {
		var col string
		if value != nil {
			jtv, ok := value.(string)
			if !ok {
				return fmt.Errorf("expected String to be of type string, got %T instead", value)
			}
			col = jtv
		}
		cv = append(cv, col)

	}
	*v = cv
	return nil
}

func awsRestjson1_deserializeDocumentHighlights(v *map[string]string, value interface{}) error {
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

	var mv map[string]string
	if *v == nil {
		mv = map[string]string{}
	} else {
		mv = *v
	}

	for key, value := range shape {
		var parsedVal string
		if value != nil {
			jtv, ok := value.(string)
			if !ok {
				return fmt.Errorf("expected String to be of type string, got %T instead", value)
			}
			parsedVal = jtv
		}
		mv[key] = parsedVal

	}
	*v = mv
	return nil
}

func awsRestjson1_deserializeDocumentHit(v **types.Hit, value interface{}) error {
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

	var sv *types.Hit
	if *v == nil {
		sv = &types.Hit{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "exprs":
			if err := awsRestjson1_deserializeDocumentExprs(&sv.Exprs, value); err != nil {
				return err
			}
		case "fields":
			if err := awsRestjson1_deserializeDocumentFields(&sv.Fields, value); err != nil {
				return err
			}
		case "highlights":
			if err := awsRestjson1_deserializeDocumentHighlights(&sv.Highlights, value); err != nil {
				return err
			}
		case "id":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected String to be of type string, got %T instead", value)
				}
				sv.Id = ptr.String(jtv)
			}
		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsRestjson1_deserializeDocumentHitList(v *[]types.Hit, value interface{}) error {
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

	var cv []types.Hit
	if *v == nil {
		cv = []types.Hit{}
	} else {
		cv = *v
	}

	for _, value := range shape {
		var col types.Hit
		destAddr := &col
		if err := awsRestjson1_deserializeDocumentHit(&destAddr, value); err != nil {
			return err
		}
		col = *destAddr
		cv = append(cv, col)

	}
	*v = cv
	return nil
}

func awsRestjson1_deserializeDocumentHits(v **types.Hits, value interface{}) error {
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

	var sv *types.Hits
	if *v == nil {
		sv = &types.Hits{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "cursor":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected String to be of type string, got %T instead", value)
				}
				sv.Cursor = ptr.String(jtv)
			}
		case "found":
			if value != nil {
				jtv, ok := value.(json.Number)
				if !ok {
					return fmt.Errorf("expected Long to be json.Number, got %T instead", value)
				}
				i64, err := jtv.Int64()
				if err != nil {
					return err
				}
				sv.Found = i64
			}
		case "hit":
			if err := awsRestjson1_deserializeDocumentHitList(&sv.Hit, value); err != nil {
				return err
			}
		case "start":
			if value != nil {
				jtv, ok := value.(json.Number)
				if !ok {
					return fmt.Errorf("expected Long to be json.Number, got %T instead", value)
				}
				i64, err := jtv.Int64()
				if err != nil {
					return err
				}
				sv.Start = i64
			}
		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsRestjson1_deserializeDocumentSearchException(v **types.SearchException, value interface{}) error {
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

	var sv *types.SearchException
	if *v == nil {
		sv = &types.SearchException{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "message", "Message":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected String to be of type string, got %T instead", value)
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

func awsRestjson1_deserializeDocumentSearchStatus(v **types.SearchStatus, value interface{}) error {
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

	var sv *types.SearchStatus
	if *v == nil {
		sv = &types.SearchStatus{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "rid":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected String to be of type string, got %T instead", value)
				}
				sv.Rid = ptr.String(jtv)
			}
		case "timems":
			if value != nil {
				jtv, ok := value.(json.Number)
				if !ok {
					return fmt.Errorf("expected Long to be json.Number, got %T instead", value)
				}
				i64, err := jtv.Int64()
				if err != nil {
					return err
				}
				sv.Timems = i64
			}
		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsRestjson1_deserializeDocumentStats(v *map[string]types.FieldStats, value interface{}) error {
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

	var mv map[string]types.FieldStats
	if *v == nil {
		mv = map[string]types.FieldStats{}
	} else {
		mv = *v
	}

	for key, value := range shape {
		var parsedVal types.FieldStats
		mapVar := parsedVal
		destAddr := &mapVar
		if err := awsRestjson1_deserializeDocumentFieldStats(&destAddr, value); err != nil {
			return err
		}
		parsedVal = *destAddr
		mv[key] = parsedVal

	}
	*v = mv
	return nil
}

func awsRestjson1_deserializeDocumentSuggestionMatch(v **types.SuggestionMatch, value interface{}) error {
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

	var sv *types.SuggestionMatch
	if *v == nil {
		sv = &types.SuggestionMatch{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "id":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected String to be of type string, got %T instead", value)
				}
				sv.Id = ptr.String(jtv)
			}
		case "score":
			if value != nil {
				jtv, ok := value.(json.Number)
				if !ok {
					return fmt.Errorf("expected Long to be json.Number, got %T instead", value)
				}
				i64, err := jtv.Int64()
				if err != nil {
					return err
				}
				sv.Score = i64
			}
		case "suggestion":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected String to be of type string, got %T instead", value)
				}
				sv.Suggestion = ptr.String(jtv)
			}
		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsRestjson1_deserializeDocumentSuggestions(v *[]types.SuggestionMatch, value interface{}) error {
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

	var cv []types.SuggestionMatch
	if *v == nil {
		cv = []types.SuggestionMatch{}
	} else {
		cv = *v
	}

	for _, value := range shape {
		var col types.SuggestionMatch
		destAddr := &col
		if err := awsRestjson1_deserializeDocumentSuggestionMatch(&destAddr, value); err != nil {
			return err
		}
		col = *destAddr
		cv = append(cv, col)

	}
	*v = cv
	return nil
}

func awsRestjson1_deserializeDocumentSuggestModel(v **types.SuggestModel, value interface{}) error {
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

	var sv *types.SuggestModel
	if *v == nil {
		sv = &types.SuggestModel{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "found":
			if value != nil {
				jtv, ok := value.(json.Number)
				if !ok {
					return fmt.Errorf("expected Long to be json.Number, got %T instead", value)
				}
				i64, err := jtv.Int64()
				if err != nil {
					return err
				}
				sv.Found = i64
			}
		case "query":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected String to be of type string, got %T instead", value)
				}
				sv.Query = ptr.String(jtv)
			}
		case "suggestions":
			if err := awsRestjson1_deserializeDocumentSuggestions(&sv.Suggestions, value); err != nil {
				return err
			}
		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsRestjson1_deserializeDocumentSuggestStatus(v **types.SuggestStatus, value interface{}) error {
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

	var sv *types.SuggestStatus
	if *v == nil {
		sv = &types.SuggestStatus{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "rid":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected String to be of type string, got %T instead", value)
				}
				sv.Rid = ptr.String(jtv)
			}
		case "timems":
			if value != nil {
				jtv, ok := value.(json.Number)
				if !ok {
					return fmt.Errorf("expected Long to be json.Number, got %T instead", value)
				}
				i64, err := jtv.Int64()
				if err != nil {
					return err
				}
				sv.Timems = i64
			}
		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}
