module github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression

go 1.21

require (
	github.com/aws/aws-sdk-go-v2 v1.32.0
	github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue v1.15.10
	github.com/aws/aws-sdk-go-v2/service/dynamodb v1.36.0
)

require (
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.3.19 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.6.19 // indirect
	github.com/aws/aws-sdk-go-v2/service/dynamodbstreams v1.24.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.12.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/endpoint-discovery v1.10.0 // indirect
	github.com/aws/smithy-go v1.22.0 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
)

replace github.com/aws/aws-sdk-go-v2 => ../../../

replace github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue => ../../../feature/dynamodb/attributevalue/

replace github.com/aws/aws-sdk-go-v2/internal/configsources => ../../../internal/configsources/

replace github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 => ../../../internal/endpoints/v2/

replace github.com/aws/aws-sdk-go-v2/service/dynamodb => ../../../service/dynamodb/

replace github.com/aws/aws-sdk-go-v2/service/dynamodbstreams => ../../../service/dynamodbstreams/

replace github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding => ../../../service/internal/accept-encoding/

replace github.com/aws/aws-sdk-go-v2/service/internal/endpoint-discovery => ../../../service/internal/endpoint-discovery/
