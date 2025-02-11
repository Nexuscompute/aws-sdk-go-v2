// Code generated by smithy-go-codegen DO NOT EDIT.

package query

import (
	"bytes"
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/protocoltest/query/types"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/ptr"
	smithyrand "github.com/aws/smithy-go/rand"
	smithytesting "github.com/aws/smithy-go/testing"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestClient_XmlMapsXmlName_awsAwsqueryDeserialize(t *testing.T) {
	cases := map[string]struct {
		StatusCode    int
		Header        http.Header
		BodyMediaType string
		Body          []byte
		ExpectResult  *XmlMapsXmlNameOutput
	}{
		// Serializes XML lists
		"QueryQueryXmlMapsXmlName": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"text/xml"},
			},
			BodyMediaType: "application/xml",
			Body: []byte(`<XmlMapsXmlNameResponse xmlns="https://example.com/">
			    <XmlMapsXmlNameResult>
			        <myMap>
			            <entry>
			                <Attribute>foo</Attribute>
			                <Setting>
			                    <hi>there</hi>
			                </Setting>
			            </entry>
			            <entry>
			                <Attribute>baz</Attribute>
			                <Setting>
			                    <hi>bye</hi>
			                </Setting>
			            </entry>
			        </myMap>
			    </XmlMapsXmlNameResult>
			</XmlMapsXmlNameResponse>
			`),
			ExpectResult: &XmlMapsXmlNameOutput{
				MyMap: map[string]types.GreetingStruct{
					"foo": {
						Hi: ptr.String("there"),
					},
					"baz": {
						Hi: ptr.String("bye"),
					},
				},
			},
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			serverURL := "http://localhost:8888/"
			client := New(Options{
				HTTPClient: smithyhttp.ClientDoFunc(func(r *http.Request) (*http.Response, error) {
					headers := http.Header{}
					for k, vs := range c.Header {
						for _, v := range vs {
							headers.Add(k, v)
						}
					}
					if len(c.BodyMediaType) != 0 && len(headers.Values("Content-Type")) == 0 {
						headers.Set("Content-Type", c.BodyMediaType)
					}
					response := &http.Response{
						StatusCode: c.StatusCode,
						Header:     headers,
						Request:    r,
					}
					if len(c.Body) != 0 {
						response.ContentLength = int64(len(c.Body))
						response.Body = ioutil.NopCloser(bytes.NewReader(c.Body))
					} else {

						response.Body = http.NoBody
					}
					return response, nil
				}),
				APIOptions: []func(*middleware.Stack) error{
					func(s *middleware.Stack) error {
						s.Finalize.Clear()
						s.Initialize.Remove(`OperationInputValidation`)
						return nil
					},
				},
				EndpointResolver: EndpointResolverFunc(func(region string, options EndpointResolverOptions) (e aws.Endpoint, err error) {
					e.URL = serverURL
					e.SigningRegion = "us-west-2"
					return e, err
				}),
				IdempotencyTokenProvider: smithyrand.NewUUIDIdempotencyToken(&smithytesting.ByteLoop{}),
				Region:                   "us-west-2",
			})
			var params XmlMapsXmlNameInput
			result, err := client.XmlMapsXmlName(context.Background(), &params)
			if err != nil {
				t.Fatalf("expect nil err, got %v", err)
			}
			if result == nil {
				t.Fatalf("expect not nil result")
			}
			if err := smithytesting.CompareValues(c.ExpectResult, result); err != nil {
				t.Errorf("expect c.ExpectResult value match:\n%v", err)
			}
		})
	}
}
