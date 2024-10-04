module github.com/aws/aws-sdk-go-v2/service/opsworkscm

go 1.21

require (
	github.com/aws/aws-sdk-go-v2 v1.32.0
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.3.19
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.6.19
	github.com/aws/smithy-go v1.22.0
	github.com/jmespath/go-jmespath v0.4.0
)

replace github.com/aws/aws-sdk-go-v2 => ../../

replace github.com/aws/aws-sdk-go-v2/internal/configsources => ../../internal/configsources/

replace github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 => ../../internal/endpoints/v2/
