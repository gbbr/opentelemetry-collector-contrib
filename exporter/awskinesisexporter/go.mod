module github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awskinesisexporter

go 1.17

require (
	github.com/aws/aws-sdk-go v1.41.14
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/jaeger v0.66.0
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/zipkin v0.38.0
	github.com/stretchr/testify v1.8.1
	go.opentelemetry.io/collector v0.66.0
	go.opentelemetry.io/collector/model v0.38.1-0.20211103011348-c24dfeb047a8
	go.uber.org/zap v1.24.0
)

require (
	github.com/gogo/protobuf v1.3.2
	github.com/google/uuid v1.3.0
	github.com/jaegertracing/jaeger v1.40.0
	go.opentelemetry.io/collector/component v0.66.0
	go.opentelemetry.io/collector/consumer v0.66.0
	go.uber.org/multierr v1.8.0
)

require (
	github.com/apache/thrift v0.17.0 // indirect
	github.com/benbjohnson/clock v1.3.0 // indirect
	github.com/cenkalti/backoff/v4 v4.2.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/knadh/koanf v1.4.4 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal v0.66.0 // indirect
	github.com/opentracing/opentracing-go v1.2.0 // indirect
	github.com/openzipkin/zipkin-go v0.3.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/uber/jaeger-client-go v2.30.0+incompatible // indirect
	github.com/uber/jaeger-lib v2.4.1+incompatible // indirect
	go.opencensus.io v0.24.0 // indirect
	go.opentelemetry.io/collector/featuregate v0.65.0 // indirect
	go.opentelemetry.io/collector/pdata v0.66.0 // indirect
	go.opentelemetry.io/otel v1.11.2 // indirect
	go.opentelemetry.io/otel/metric v0.34.0 // indirect
	go.opentelemetry.io/otel/trace v1.11.2 // indirect
	go.uber.org/atomic v1.10.0 // indirect
	golang.org/x/net v0.0.0-20221014081412-f15817d10f9b // indirect
	golang.org/x/sys v0.3.0 // indirect
	golang.org/x/text v0.4.0 // indirect
	google.golang.org/genproto v0.0.0-20221027153422-115e99e71e1c // indirect
	google.golang.org/grpc v1.51.0 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/jaeger => ../../pkg/translator/jaeger

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/zipkin => ../../pkg/translator/zipkin

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal => ../../internal/coreinternal

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/opencensus => ../../pkg/translator/opencensus
