//go:build ignore
// +build ignore

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"text/template"

	"github.com/aws/aws-sdk-go-v2/internal/endpoints/awsrulesfn"
)

var tmpl = template.Must(template.New("generate").
	Funcs(map[string]interface{}{
		"quote": func(v string) string {
			return strconv.Quote(v)
		},
		"strOrNil": func(v *string) string {
			if v == nil {
				return "nil"
			}
			return fmt.Sprintf("ptr.String(%q)", *v)
		},
		"boolOrNil": func(v *bool) string {
			if v == nil {
				return "nil"
			}
			return fmt.Sprintf("ptr.Bool(%t)", *v)
		},
	}).
	Parse(`
{{- block "root" $ -}}
// Code generated by endpoint/awsrulesfn/internal/partition. DO NOT EDIT.

package awsrulesfn


// GetPartition returns an AWS [Partition] for the region provided. If the
// partition cannot be determined nil will be returned.
func GetPartition(region string) *PartitionConfig {
	return getPartition(partitions, region)
}

var partitions = []Partition {

	{{- range $_, $partition := $.Partitions }}
			Partition {
				ID: {{ quote $partition.ID }},
				RegionRegex: {{ quote $partition.RegionRegex }},
				DefaultConfig: PartitionConfig{
					Name:                 {{ quote $partition.DefaultConfig.Name }},
					DnsSuffix:            {{ quote $partition.DefaultConfig.DnsSuffix }},
					DualStackDnsSuffix:   {{ quote $partition.DefaultConfig.DualStackDnsSuffix }},
					SupportsFIPS:         {{ $partition.DefaultConfig.SupportsFIPS }},
					SupportsDualStack:    {{ $partition.DefaultConfig.SupportsDualStack }},
					ImplicitGlobalRegion: {{ quote $partition.DefaultConfig.ImplicitGlobalRegion }},
				},
				Regions: map[string]RegionOverrides {
				{{- range $region, $config := $partition.Regions }}
					{{ quote $region }}: RegionOverrides{
						Name:               {{ strOrNil $config.Name }},
						DnsSuffix:          {{ strOrNil $config.DnsSuffix }},
						DualStackDnsSuffix: {{ strOrNil $config.DualStackDnsSuffix }},
						SupportsFIPS:       {{ boolOrNil $config.SupportsFIPS }},
						SupportsDualStack:  {{ boolOrNil $config.SupportsDualStack }},
					},
				{{- end }}
				},
			},
	{{- end }}
}
{{- end }}
`))

type PartitionModel struct {
	Partitions []awsrulesfn.Partition `json:"partitions"`
}

func main() {
	var modelFilename, outputFilename string
	flag.StringVar(&modelFilename, "model", "partitions.json", "The `file` providing the partition model metadata.")
	flag.StringVar(&outputFilename, "output", "partitions.go", "The `file` to write the source to.")
	flag.Parse()

	modelFile, err := os.Open(modelFilename)
	if err != nil {
		log.Fatalf("failed to open model file, %v", err)
	}

	var model PartitionModel
	if err = json.NewDecoder(modelFile).Decode(&model); err != nil {
		log.Fatalf("failed to unmarshal model file, %v", err)
	}
	modelFile.Close()

	file, err := os.Create(outputFilename)
	if err != nil {
		log.Fatalf("failed to create output file, %v", err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Fatalf("failed to close output file, %v", err)
		}
	}()

	err = tmpl.Execute(file, struct {
		Partitions []awsrulesfn.Partition
	}{
		Partitions: model.Partitions,
	})
	if err != nil {
		log.Fatalf("failed to render partition Go file, %v", err)
	}
}
