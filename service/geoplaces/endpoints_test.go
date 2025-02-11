// Code generated by smithy-go-codegen DO NOT EDIT.

package geoplaces

import (
	"context"
	smithy "github.com/aws/smithy-go"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/ptr"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"testing"
)

// For custom endpoint with region not set and fips disabled
func TestEndpointCase0(t *testing.T) {
	var params = EndpointParameters{
		Endpoint: ptr.String("https://example.com"),
		UseFIPS:  ptr.Bool(false),
	}

	resolver := NewDefaultEndpointResolverV2()
	result, err := resolver.ResolveEndpoint(context.Background(), params)
	_, _ = result, err

	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}

	uri, _ := url.Parse("https://example.com")

	expectEndpoint := smithyendpoints.Endpoint{
		URI:        *uri,
		Headers:    http.Header{},
		Properties: smithy.Properties{},
	}

	if e, a := expectEndpoint.URI, result.URI; e != a {
		t.Errorf("expect %v URI, got %v", e, a)
	}

	if !reflect.DeepEqual(expectEndpoint.Headers, result.Headers) {
		t.Errorf("expect headers to match\n%v != %v", expectEndpoint.Headers, result.Headers)
	}

	if !reflect.DeepEqual(expectEndpoint.Properties, result.Properties) {
		t.Errorf("expect properties to match\n%v != %v", expectEndpoint.Properties, result.Properties)
	}
}

// For custom endpoint with fips enabled
func TestEndpointCase1(t *testing.T) {
	var params = EndpointParameters{
		Endpoint: ptr.String("https://example.com"),
		UseFIPS:  ptr.Bool(true),
	}

	resolver := NewDefaultEndpointResolverV2()
	result, err := resolver.ResolveEndpoint(context.Background(), params)
	_, _ = result, err

	if err == nil {
		t.Fatalf("expect error, got none")
	}
	if e, a := "Invalid Configuration: FIPS and custom endpoint are not supported", err.Error(); !strings.Contains(a, e) {
		t.Errorf("expect %v error in %v", e, a)
	}
}

// For custom endpoint with fips disabled and dualstack enabled
func TestEndpointCase2(t *testing.T) {
	var params = EndpointParameters{
		Endpoint:     ptr.String("https://example.com"),
		UseFIPS:      ptr.Bool(false),
		UseDualStack: ptr.Bool(true),
	}

	resolver := NewDefaultEndpointResolverV2()
	result, err := resolver.ResolveEndpoint(context.Background(), params)
	_, _ = result, err

	if err == nil {
		t.Fatalf("expect error, got none")
	}
	if e, a := "Invalid Configuration: Dualstack and custom endpoint are not supported", err.Error(); !strings.Contains(a, e) {
		t.Errorf("expect %v error in %v", e, a)
	}
}

// For region us-east-1 with FIPS enabled and DualStack enabled
func TestEndpointCase3(t *testing.T) {
	var params = EndpointParameters{
		Region:       ptr.String("us-east-1"),
		UseFIPS:      ptr.Bool(true),
		UseDualStack: ptr.Bool(true),
	}

	resolver := NewDefaultEndpointResolverV2()
	result, err := resolver.ResolveEndpoint(context.Background(), params)
	_, _ = result, err

	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}

	uri, _ := url.Parse("https://places.geo-fips.us-east-1.api.aws/v2")

	expectEndpoint := smithyendpoints.Endpoint{
		URI:        *uri,
		Headers:    http.Header{},
		Properties: smithy.Properties{},
	}

	if e, a := expectEndpoint.URI, result.URI; e != a {
		t.Errorf("expect %v URI, got %v", e, a)
	}

	if !reflect.DeepEqual(expectEndpoint.Headers, result.Headers) {
		t.Errorf("expect headers to match\n%v != %v", expectEndpoint.Headers, result.Headers)
	}

	if !reflect.DeepEqual(expectEndpoint.Properties, result.Properties) {
		t.Errorf("expect properties to match\n%v != %v", expectEndpoint.Properties, result.Properties)
	}
}

// For region us-east-1 with FIPS enabled and DualStack disabled
func TestEndpointCase4(t *testing.T) {
	var params = EndpointParameters{
		Region:       ptr.String("us-east-1"),
		UseFIPS:      ptr.Bool(true),
		UseDualStack: ptr.Bool(false),
	}

	resolver := NewDefaultEndpointResolverV2()
	result, err := resolver.ResolveEndpoint(context.Background(), params)
	_, _ = result, err

	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}

	uri, _ := url.Parse("https://places.geo-fips.us-east-1.amazonaws.com/v2")

	expectEndpoint := smithyendpoints.Endpoint{
		URI:        *uri,
		Headers:    http.Header{},
		Properties: smithy.Properties{},
	}

	if e, a := expectEndpoint.URI, result.URI; e != a {
		t.Errorf("expect %v URI, got %v", e, a)
	}

	if !reflect.DeepEqual(expectEndpoint.Headers, result.Headers) {
		t.Errorf("expect headers to match\n%v != %v", expectEndpoint.Headers, result.Headers)
	}

	if !reflect.DeepEqual(expectEndpoint.Properties, result.Properties) {
		t.Errorf("expect properties to match\n%v != %v", expectEndpoint.Properties, result.Properties)
	}
}

// For region us-east-1 with FIPS disabled and DualStack enabled
func TestEndpointCase5(t *testing.T) {
	var params = EndpointParameters{
		Region:       ptr.String("us-east-1"),
		UseFIPS:      ptr.Bool(false),
		UseDualStack: ptr.Bool(true),
	}

	resolver := NewDefaultEndpointResolverV2()
	result, err := resolver.ResolveEndpoint(context.Background(), params)
	_, _ = result, err

	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}

	uri, _ := url.Parse("https://places.geo.us-east-1.api.aws/v2")

	expectEndpoint := smithyendpoints.Endpoint{
		URI:        *uri,
		Headers:    http.Header{},
		Properties: smithy.Properties{},
	}

	if e, a := expectEndpoint.URI, result.URI; e != a {
		t.Errorf("expect %v URI, got %v", e, a)
	}

	if !reflect.DeepEqual(expectEndpoint.Headers, result.Headers) {
		t.Errorf("expect headers to match\n%v != %v", expectEndpoint.Headers, result.Headers)
	}

	if !reflect.DeepEqual(expectEndpoint.Properties, result.Properties) {
		t.Errorf("expect properties to match\n%v != %v", expectEndpoint.Properties, result.Properties)
	}
}

// For region us-east-1 with FIPS disabled and DualStack disabled
func TestEndpointCase6(t *testing.T) {
	var params = EndpointParameters{
		Region:       ptr.String("us-east-1"),
		UseFIPS:      ptr.Bool(false),
		UseDualStack: ptr.Bool(false),
	}

	resolver := NewDefaultEndpointResolverV2()
	result, err := resolver.ResolveEndpoint(context.Background(), params)
	_, _ = result, err

	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}

	uri, _ := url.Parse("https://places.geo.us-east-1.amazonaws.com/v2")

	expectEndpoint := smithyendpoints.Endpoint{
		URI:        *uri,
		Headers:    http.Header{},
		Properties: smithy.Properties{},
	}

	if e, a := expectEndpoint.URI, result.URI; e != a {
		t.Errorf("expect %v URI, got %v", e, a)
	}

	if !reflect.DeepEqual(expectEndpoint.Headers, result.Headers) {
		t.Errorf("expect headers to match\n%v != %v", expectEndpoint.Headers, result.Headers)
	}

	if !reflect.DeepEqual(expectEndpoint.Properties, result.Properties) {
		t.Errorf("expect properties to match\n%v != %v", expectEndpoint.Properties, result.Properties)
	}
}

// For region cn-northwest-1 with FIPS enabled and DualStack enabled
func TestEndpointCase7(t *testing.T) {
	var params = EndpointParameters{
		Region:       ptr.String("cn-northwest-1"),
		UseFIPS:      ptr.Bool(true),
		UseDualStack: ptr.Bool(true),
	}

	resolver := NewDefaultEndpointResolverV2()
	result, err := resolver.ResolveEndpoint(context.Background(), params)
	_, _ = result, err

	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}

	uri, _ := url.Parse("https://geo-places-fips.cn-northwest-1.api.amazonwebservices.com.cn")

	expectEndpoint := smithyendpoints.Endpoint{
		URI:        *uri,
		Headers:    http.Header{},
		Properties: smithy.Properties{},
	}

	if e, a := expectEndpoint.URI, result.URI; e != a {
		t.Errorf("expect %v URI, got %v", e, a)
	}

	if !reflect.DeepEqual(expectEndpoint.Headers, result.Headers) {
		t.Errorf("expect headers to match\n%v != %v", expectEndpoint.Headers, result.Headers)
	}

	if !reflect.DeepEqual(expectEndpoint.Properties, result.Properties) {
		t.Errorf("expect properties to match\n%v != %v", expectEndpoint.Properties, result.Properties)
	}
}

// For region cn-northwest-1 with FIPS enabled and DualStack disabled
func TestEndpointCase8(t *testing.T) {
	var params = EndpointParameters{
		Region:       ptr.String("cn-northwest-1"),
		UseFIPS:      ptr.Bool(true),
		UseDualStack: ptr.Bool(false),
	}

	resolver := NewDefaultEndpointResolverV2()
	result, err := resolver.ResolveEndpoint(context.Background(), params)
	_, _ = result, err

	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}

	uri, _ := url.Parse("https://geo-places-fips.cn-northwest-1.amazonaws.com.cn")

	expectEndpoint := smithyendpoints.Endpoint{
		URI:        *uri,
		Headers:    http.Header{},
		Properties: smithy.Properties{},
	}

	if e, a := expectEndpoint.URI, result.URI; e != a {
		t.Errorf("expect %v URI, got %v", e, a)
	}

	if !reflect.DeepEqual(expectEndpoint.Headers, result.Headers) {
		t.Errorf("expect headers to match\n%v != %v", expectEndpoint.Headers, result.Headers)
	}

	if !reflect.DeepEqual(expectEndpoint.Properties, result.Properties) {
		t.Errorf("expect properties to match\n%v != %v", expectEndpoint.Properties, result.Properties)
	}
}

// For region cn-northwest-1 with FIPS disabled and DualStack enabled
func TestEndpointCase9(t *testing.T) {
	var params = EndpointParameters{
		Region:       ptr.String("cn-northwest-1"),
		UseFIPS:      ptr.Bool(false),
		UseDualStack: ptr.Bool(true),
	}

	resolver := NewDefaultEndpointResolverV2()
	result, err := resolver.ResolveEndpoint(context.Background(), params)
	_, _ = result, err

	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}

	uri, _ := url.Parse("https://geo-places.cn-northwest-1.api.amazonwebservices.com.cn")

	expectEndpoint := smithyendpoints.Endpoint{
		URI:        *uri,
		Headers:    http.Header{},
		Properties: smithy.Properties{},
	}

	if e, a := expectEndpoint.URI, result.URI; e != a {
		t.Errorf("expect %v URI, got %v", e, a)
	}

	if !reflect.DeepEqual(expectEndpoint.Headers, result.Headers) {
		t.Errorf("expect headers to match\n%v != %v", expectEndpoint.Headers, result.Headers)
	}

	if !reflect.DeepEqual(expectEndpoint.Properties, result.Properties) {
		t.Errorf("expect properties to match\n%v != %v", expectEndpoint.Properties, result.Properties)
	}
}

// For region cn-northwest-1 with FIPS disabled and DualStack disabled
func TestEndpointCase10(t *testing.T) {
	var params = EndpointParameters{
		Region:       ptr.String("cn-northwest-1"),
		UseFIPS:      ptr.Bool(false),
		UseDualStack: ptr.Bool(false),
	}

	resolver := NewDefaultEndpointResolverV2()
	result, err := resolver.ResolveEndpoint(context.Background(), params)
	_, _ = result, err

	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}

	uri, _ := url.Parse("https://geo-places.cn-northwest-1.amazonaws.com.cn")

	expectEndpoint := smithyendpoints.Endpoint{
		URI:        *uri,
		Headers:    http.Header{},
		Properties: smithy.Properties{},
	}

	if e, a := expectEndpoint.URI, result.URI; e != a {
		t.Errorf("expect %v URI, got %v", e, a)
	}

	if !reflect.DeepEqual(expectEndpoint.Headers, result.Headers) {
		t.Errorf("expect headers to match\n%v != %v", expectEndpoint.Headers, result.Headers)
	}

	if !reflect.DeepEqual(expectEndpoint.Properties, result.Properties) {
		t.Errorf("expect properties to match\n%v != %v", expectEndpoint.Properties, result.Properties)
	}
}

// For region us-gov-west-1 with FIPS enabled and DualStack enabled
func TestEndpointCase11(t *testing.T) {
	var params = EndpointParameters{
		Region:       ptr.String("us-gov-west-1"),
		UseFIPS:      ptr.Bool(true),
		UseDualStack: ptr.Bool(true),
	}

	resolver := NewDefaultEndpointResolverV2()
	result, err := resolver.ResolveEndpoint(context.Background(), params)
	_, _ = result, err

	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}

	uri, _ := url.Parse("https://places.geo-fips.us-gov-west-1.us-gov.api.aws/v2")

	expectEndpoint := smithyendpoints.Endpoint{
		URI:        *uri,
		Headers:    http.Header{},
		Properties: smithy.Properties{},
	}

	if e, a := expectEndpoint.URI, result.URI; e != a {
		t.Errorf("expect %v URI, got %v", e, a)
	}

	if !reflect.DeepEqual(expectEndpoint.Headers, result.Headers) {
		t.Errorf("expect headers to match\n%v != %v", expectEndpoint.Headers, result.Headers)
	}

	if !reflect.DeepEqual(expectEndpoint.Properties, result.Properties) {
		t.Errorf("expect properties to match\n%v != %v", expectEndpoint.Properties, result.Properties)
	}
}

// For region us-gov-west-1 with FIPS enabled and DualStack disabled
func TestEndpointCase12(t *testing.T) {
	var params = EndpointParameters{
		Region:       ptr.String("us-gov-west-1"),
		UseFIPS:      ptr.Bool(true),
		UseDualStack: ptr.Bool(false),
	}

	resolver := NewDefaultEndpointResolverV2()
	result, err := resolver.ResolveEndpoint(context.Background(), params)
	_, _ = result, err

	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}

	uri, _ := url.Parse("https://places.geo-fips.us-gov-west-1.us-gov.amazonaws.com/v2")

	expectEndpoint := smithyendpoints.Endpoint{
		URI:        *uri,
		Headers:    http.Header{},
		Properties: smithy.Properties{},
	}

	if e, a := expectEndpoint.URI, result.URI; e != a {
		t.Errorf("expect %v URI, got %v", e, a)
	}

	if !reflect.DeepEqual(expectEndpoint.Headers, result.Headers) {
		t.Errorf("expect headers to match\n%v != %v", expectEndpoint.Headers, result.Headers)
	}

	if !reflect.DeepEqual(expectEndpoint.Properties, result.Properties) {
		t.Errorf("expect properties to match\n%v != %v", expectEndpoint.Properties, result.Properties)
	}
}

// For region us-gov-west-1 with FIPS disabled and DualStack enabled
func TestEndpointCase13(t *testing.T) {
	var params = EndpointParameters{
		Region:       ptr.String("us-gov-west-1"),
		UseFIPS:      ptr.Bool(false),
		UseDualStack: ptr.Bool(true),
	}

	resolver := NewDefaultEndpointResolverV2()
	result, err := resolver.ResolveEndpoint(context.Background(), params)
	_, _ = result, err

	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}

	uri, _ := url.Parse("https://places.geo.us-gov-west-1.us-gov.api.aws/v2")

	expectEndpoint := smithyendpoints.Endpoint{
		URI:        *uri,
		Headers:    http.Header{},
		Properties: smithy.Properties{},
	}

	if e, a := expectEndpoint.URI, result.URI; e != a {
		t.Errorf("expect %v URI, got %v", e, a)
	}

	if !reflect.DeepEqual(expectEndpoint.Headers, result.Headers) {
		t.Errorf("expect headers to match\n%v != %v", expectEndpoint.Headers, result.Headers)
	}

	if !reflect.DeepEqual(expectEndpoint.Properties, result.Properties) {
		t.Errorf("expect properties to match\n%v != %v", expectEndpoint.Properties, result.Properties)
	}
}

// For region us-gov-west-1 with FIPS disabled and DualStack disabled
func TestEndpointCase14(t *testing.T) {
	var params = EndpointParameters{
		Region:       ptr.String("us-gov-west-1"),
		UseFIPS:      ptr.Bool(false),
		UseDualStack: ptr.Bool(false),
	}

	resolver := NewDefaultEndpointResolverV2()
	result, err := resolver.ResolveEndpoint(context.Background(), params)
	_, _ = result, err

	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}

	uri, _ := url.Parse("https://places.geo.us-gov-west-1.us-gov.amazonaws.com/v2")

	expectEndpoint := smithyendpoints.Endpoint{
		URI:        *uri,
		Headers:    http.Header{},
		Properties: smithy.Properties{},
	}

	if e, a := expectEndpoint.URI, result.URI; e != a {
		t.Errorf("expect %v URI, got %v", e, a)
	}

	if !reflect.DeepEqual(expectEndpoint.Headers, result.Headers) {
		t.Errorf("expect headers to match\n%v != %v", expectEndpoint.Headers, result.Headers)
	}

	if !reflect.DeepEqual(expectEndpoint.Properties, result.Properties) {
		t.Errorf("expect properties to match\n%v != %v", expectEndpoint.Properties, result.Properties)
	}
}

// For region us-iso-east-1 with FIPS enabled and DualStack enabled
func TestEndpointCase15(t *testing.T) {
	var params = EndpointParameters{
		Region:       ptr.String("us-iso-east-1"),
		UseFIPS:      ptr.Bool(true),
		UseDualStack: ptr.Bool(true),
	}

	resolver := NewDefaultEndpointResolverV2()
	result, err := resolver.ResolveEndpoint(context.Background(), params)
	_, _ = result, err

	if err == nil {
		t.Fatalf("expect error, got none")
	}
	if e, a := "FIPS and DualStack are enabled, but this partition does not support one or both", err.Error(); !strings.Contains(a, e) {
		t.Errorf("expect %v error in %v", e, a)
	}
}

// For region us-iso-east-1 with FIPS enabled and DualStack disabled
func TestEndpointCase16(t *testing.T) {
	var params = EndpointParameters{
		Region:       ptr.String("us-iso-east-1"),
		UseFIPS:      ptr.Bool(true),
		UseDualStack: ptr.Bool(false),
	}

	resolver := NewDefaultEndpointResolverV2()
	result, err := resolver.ResolveEndpoint(context.Background(), params)
	_, _ = result, err

	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}

	uri, _ := url.Parse("https://geo-places-fips.us-iso-east-1.c2s.ic.gov")

	expectEndpoint := smithyendpoints.Endpoint{
		URI:        *uri,
		Headers:    http.Header{},
		Properties: smithy.Properties{},
	}

	if e, a := expectEndpoint.URI, result.URI; e != a {
		t.Errorf("expect %v URI, got %v", e, a)
	}

	if !reflect.DeepEqual(expectEndpoint.Headers, result.Headers) {
		t.Errorf("expect headers to match\n%v != %v", expectEndpoint.Headers, result.Headers)
	}

	if !reflect.DeepEqual(expectEndpoint.Properties, result.Properties) {
		t.Errorf("expect properties to match\n%v != %v", expectEndpoint.Properties, result.Properties)
	}
}

// For region us-iso-east-1 with FIPS disabled and DualStack enabled
func TestEndpointCase17(t *testing.T) {
	var params = EndpointParameters{
		Region:       ptr.String("us-iso-east-1"),
		UseFIPS:      ptr.Bool(false),
		UseDualStack: ptr.Bool(true),
	}

	resolver := NewDefaultEndpointResolverV2()
	result, err := resolver.ResolveEndpoint(context.Background(), params)
	_, _ = result, err

	if err == nil {
		t.Fatalf("expect error, got none")
	}
	if e, a := "DualStack is enabled but this partition does not support DualStack", err.Error(); !strings.Contains(a, e) {
		t.Errorf("expect %v error in %v", e, a)
	}
}

// For region us-iso-east-1 with FIPS disabled and DualStack disabled
func TestEndpointCase18(t *testing.T) {
	var params = EndpointParameters{
		Region:       ptr.String("us-iso-east-1"),
		UseFIPS:      ptr.Bool(false),
		UseDualStack: ptr.Bool(false),
	}

	resolver := NewDefaultEndpointResolverV2()
	result, err := resolver.ResolveEndpoint(context.Background(), params)
	_, _ = result, err

	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}

	uri, _ := url.Parse("https://geo-places.us-iso-east-1.c2s.ic.gov")

	expectEndpoint := smithyendpoints.Endpoint{
		URI:        *uri,
		Headers:    http.Header{},
		Properties: smithy.Properties{},
	}

	if e, a := expectEndpoint.URI, result.URI; e != a {
		t.Errorf("expect %v URI, got %v", e, a)
	}

	if !reflect.DeepEqual(expectEndpoint.Headers, result.Headers) {
		t.Errorf("expect headers to match\n%v != %v", expectEndpoint.Headers, result.Headers)
	}

	if !reflect.DeepEqual(expectEndpoint.Properties, result.Properties) {
		t.Errorf("expect properties to match\n%v != %v", expectEndpoint.Properties, result.Properties)
	}
}

// For region us-isob-east-1 with FIPS enabled and DualStack enabled
func TestEndpointCase19(t *testing.T) {
	var params = EndpointParameters{
		Region:       ptr.String("us-isob-east-1"),
		UseFIPS:      ptr.Bool(true),
		UseDualStack: ptr.Bool(true),
	}

	resolver := NewDefaultEndpointResolverV2()
	result, err := resolver.ResolveEndpoint(context.Background(), params)
	_, _ = result, err

	if err == nil {
		t.Fatalf("expect error, got none")
	}
	if e, a := "FIPS and DualStack are enabled, but this partition does not support one or both", err.Error(); !strings.Contains(a, e) {
		t.Errorf("expect %v error in %v", e, a)
	}
}

// For region us-isob-east-1 with FIPS enabled and DualStack disabled
func TestEndpointCase20(t *testing.T) {
	var params = EndpointParameters{
		Region:       ptr.String("us-isob-east-1"),
		UseFIPS:      ptr.Bool(true),
		UseDualStack: ptr.Bool(false),
	}

	resolver := NewDefaultEndpointResolverV2()
	result, err := resolver.ResolveEndpoint(context.Background(), params)
	_, _ = result, err

	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}

	uri, _ := url.Parse("https://geo-places-fips.us-isob-east-1.sc2s.sgov.gov")

	expectEndpoint := smithyendpoints.Endpoint{
		URI:        *uri,
		Headers:    http.Header{},
		Properties: smithy.Properties{},
	}

	if e, a := expectEndpoint.URI, result.URI; e != a {
		t.Errorf("expect %v URI, got %v", e, a)
	}

	if !reflect.DeepEqual(expectEndpoint.Headers, result.Headers) {
		t.Errorf("expect headers to match\n%v != %v", expectEndpoint.Headers, result.Headers)
	}

	if !reflect.DeepEqual(expectEndpoint.Properties, result.Properties) {
		t.Errorf("expect properties to match\n%v != %v", expectEndpoint.Properties, result.Properties)
	}
}

// For region us-isob-east-1 with FIPS disabled and DualStack enabled
func TestEndpointCase21(t *testing.T) {
	var params = EndpointParameters{
		Region:       ptr.String("us-isob-east-1"),
		UseFIPS:      ptr.Bool(false),
		UseDualStack: ptr.Bool(true),
	}

	resolver := NewDefaultEndpointResolverV2()
	result, err := resolver.ResolveEndpoint(context.Background(), params)
	_, _ = result, err

	if err == nil {
		t.Fatalf("expect error, got none")
	}
	if e, a := "DualStack is enabled but this partition does not support DualStack", err.Error(); !strings.Contains(a, e) {
		t.Errorf("expect %v error in %v", e, a)
	}
}

// For region us-isob-east-1 with FIPS disabled and DualStack disabled
func TestEndpointCase22(t *testing.T) {
	var params = EndpointParameters{
		Region:       ptr.String("us-isob-east-1"),
		UseFIPS:      ptr.Bool(false),
		UseDualStack: ptr.Bool(false),
	}

	resolver := NewDefaultEndpointResolverV2()
	result, err := resolver.ResolveEndpoint(context.Background(), params)
	_, _ = result, err

	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}

	uri, _ := url.Parse("https://geo-places.us-isob-east-1.sc2s.sgov.gov")

	expectEndpoint := smithyendpoints.Endpoint{
		URI:        *uri,
		Headers:    http.Header{},
		Properties: smithy.Properties{},
	}

	if e, a := expectEndpoint.URI, result.URI; e != a {
		t.Errorf("expect %v URI, got %v", e, a)
	}

	if !reflect.DeepEqual(expectEndpoint.Headers, result.Headers) {
		t.Errorf("expect headers to match\n%v != %v", expectEndpoint.Headers, result.Headers)
	}

	if !reflect.DeepEqual(expectEndpoint.Properties, result.Properties) {
		t.Errorf("expect properties to match\n%v != %v", expectEndpoint.Properties, result.Properties)
	}
}

// For region eu-isoe-west-1 with FIPS enabled and DualStack enabled
func TestEndpointCase23(t *testing.T) {
	var params = EndpointParameters{
		Region:       ptr.String("eu-isoe-west-1"),
		UseFIPS:      ptr.Bool(true),
		UseDualStack: ptr.Bool(true),
	}

	resolver := NewDefaultEndpointResolverV2()
	result, err := resolver.ResolveEndpoint(context.Background(), params)
	_, _ = result, err

	if err == nil {
		t.Fatalf("expect error, got none")
	}
	if e, a := "FIPS and DualStack are enabled, but this partition does not support one or both", err.Error(); !strings.Contains(a, e) {
		t.Errorf("expect %v error in %v", e, a)
	}
}

// For region eu-isoe-west-1 with FIPS enabled and DualStack disabled
func TestEndpointCase24(t *testing.T) {
	var params = EndpointParameters{
		Region:       ptr.String("eu-isoe-west-1"),
		UseFIPS:      ptr.Bool(true),
		UseDualStack: ptr.Bool(false),
	}

	resolver := NewDefaultEndpointResolverV2()
	result, err := resolver.ResolveEndpoint(context.Background(), params)
	_, _ = result, err

	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}

	uri, _ := url.Parse("https://geo-places-fips.eu-isoe-west-1.cloud.adc-e.uk")

	expectEndpoint := smithyendpoints.Endpoint{
		URI:        *uri,
		Headers:    http.Header{},
		Properties: smithy.Properties{},
	}

	if e, a := expectEndpoint.URI, result.URI; e != a {
		t.Errorf("expect %v URI, got %v", e, a)
	}

	if !reflect.DeepEqual(expectEndpoint.Headers, result.Headers) {
		t.Errorf("expect headers to match\n%v != %v", expectEndpoint.Headers, result.Headers)
	}

	if !reflect.DeepEqual(expectEndpoint.Properties, result.Properties) {
		t.Errorf("expect properties to match\n%v != %v", expectEndpoint.Properties, result.Properties)
	}
}

// For region eu-isoe-west-1 with FIPS disabled and DualStack enabled
func TestEndpointCase25(t *testing.T) {
	var params = EndpointParameters{
		Region:       ptr.String("eu-isoe-west-1"),
		UseFIPS:      ptr.Bool(false),
		UseDualStack: ptr.Bool(true),
	}

	resolver := NewDefaultEndpointResolverV2()
	result, err := resolver.ResolveEndpoint(context.Background(), params)
	_, _ = result, err

	if err == nil {
		t.Fatalf("expect error, got none")
	}
	if e, a := "DualStack is enabled but this partition does not support DualStack", err.Error(); !strings.Contains(a, e) {
		t.Errorf("expect %v error in %v", e, a)
	}
}

// For region eu-isoe-west-1 with FIPS disabled and DualStack disabled
func TestEndpointCase26(t *testing.T) {
	var params = EndpointParameters{
		Region:       ptr.String("eu-isoe-west-1"),
		UseFIPS:      ptr.Bool(false),
		UseDualStack: ptr.Bool(false),
	}

	resolver := NewDefaultEndpointResolverV2()
	result, err := resolver.ResolveEndpoint(context.Background(), params)
	_, _ = result, err

	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}

	uri, _ := url.Parse("https://geo-places.eu-isoe-west-1.cloud.adc-e.uk")

	expectEndpoint := smithyendpoints.Endpoint{
		URI:        *uri,
		Headers:    http.Header{},
		Properties: smithy.Properties{},
	}

	if e, a := expectEndpoint.URI, result.URI; e != a {
		t.Errorf("expect %v URI, got %v", e, a)
	}

	if !reflect.DeepEqual(expectEndpoint.Headers, result.Headers) {
		t.Errorf("expect headers to match\n%v != %v", expectEndpoint.Headers, result.Headers)
	}

	if !reflect.DeepEqual(expectEndpoint.Properties, result.Properties) {
		t.Errorf("expect properties to match\n%v != %v", expectEndpoint.Properties, result.Properties)
	}
}

// For region us-isof-south-1 with FIPS enabled and DualStack enabled
func TestEndpointCase27(t *testing.T) {
	var params = EndpointParameters{
		Region:       ptr.String("us-isof-south-1"),
		UseFIPS:      ptr.Bool(true),
		UseDualStack: ptr.Bool(true),
	}

	resolver := NewDefaultEndpointResolverV2()
	result, err := resolver.ResolveEndpoint(context.Background(), params)
	_, _ = result, err

	if err == nil {
		t.Fatalf("expect error, got none")
	}
	if e, a := "FIPS and DualStack are enabled, but this partition does not support one or both", err.Error(); !strings.Contains(a, e) {
		t.Errorf("expect %v error in %v", e, a)
	}
}

// For region us-isof-south-1 with FIPS enabled and DualStack disabled
func TestEndpointCase28(t *testing.T) {
	var params = EndpointParameters{
		Region:       ptr.String("us-isof-south-1"),
		UseFIPS:      ptr.Bool(true),
		UseDualStack: ptr.Bool(false),
	}

	resolver := NewDefaultEndpointResolverV2()
	result, err := resolver.ResolveEndpoint(context.Background(), params)
	_, _ = result, err

	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}

	uri, _ := url.Parse("https://geo-places-fips.us-isof-south-1.csp.hci.ic.gov")

	expectEndpoint := smithyendpoints.Endpoint{
		URI:        *uri,
		Headers:    http.Header{},
		Properties: smithy.Properties{},
	}

	if e, a := expectEndpoint.URI, result.URI; e != a {
		t.Errorf("expect %v URI, got %v", e, a)
	}

	if !reflect.DeepEqual(expectEndpoint.Headers, result.Headers) {
		t.Errorf("expect headers to match\n%v != %v", expectEndpoint.Headers, result.Headers)
	}

	if !reflect.DeepEqual(expectEndpoint.Properties, result.Properties) {
		t.Errorf("expect properties to match\n%v != %v", expectEndpoint.Properties, result.Properties)
	}
}

// For region us-isof-south-1 with FIPS disabled and DualStack enabled
func TestEndpointCase29(t *testing.T) {
	var params = EndpointParameters{
		Region:       ptr.String("us-isof-south-1"),
		UseFIPS:      ptr.Bool(false),
		UseDualStack: ptr.Bool(true),
	}

	resolver := NewDefaultEndpointResolverV2()
	result, err := resolver.ResolveEndpoint(context.Background(), params)
	_, _ = result, err

	if err == nil {
		t.Fatalf("expect error, got none")
	}
	if e, a := "DualStack is enabled but this partition does not support DualStack", err.Error(); !strings.Contains(a, e) {
		t.Errorf("expect %v error in %v", e, a)
	}
}

// For region us-isof-south-1 with FIPS disabled and DualStack disabled
func TestEndpointCase30(t *testing.T) {
	var params = EndpointParameters{
		Region:       ptr.String("us-isof-south-1"),
		UseFIPS:      ptr.Bool(false),
		UseDualStack: ptr.Bool(false),
	}

	resolver := NewDefaultEndpointResolverV2()
	result, err := resolver.ResolveEndpoint(context.Background(), params)
	_, _ = result, err

	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}

	uri, _ := url.Parse("https://geo-places.us-isof-south-1.csp.hci.ic.gov")

	expectEndpoint := smithyendpoints.Endpoint{
		URI:        *uri,
		Headers:    http.Header{},
		Properties: smithy.Properties{},
	}

	if e, a := expectEndpoint.URI, result.URI; e != a {
		t.Errorf("expect %v URI, got %v", e, a)
	}

	if !reflect.DeepEqual(expectEndpoint.Headers, result.Headers) {
		t.Errorf("expect headers to match\n%v != %v", expectEndpoint.Headers, result.Headers)
	}

	if !reflect.DeepEqual(expectEndpoint.Properties, result.Properties) {
		t.Errorf("expect properties to match\n%v != %v", expectEndpoint.Properties, result.Properties)
	}
}

// Missing region
func TestEndpointCase31(t *testing.T) {
	var params = EndpointParameters{}

	resolver := NewDefaultEndpointResolverV2()
	result, err := resolver.ResolveEndpoint(context.Background(), params)
	_, _ = result, err

	if err == nil {
		t.Fatalf("expect error, got none")
	}
	if e, a := "Invalid Configuration: Missing Region", err.Error(); !strings.Contains(a, e) {
		t.Errorf("expect %v error in %v", e, a)
	}
}
