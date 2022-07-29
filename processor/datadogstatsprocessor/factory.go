// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package datadogstatsprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/datadogstatsprocessor"

import (
	"context"
	"errors"
	"fmt"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.uber.org/zap"
)

const typeStr = "datadogstats"

func NewFactory() component.ProcessorFactory {
	return component.NewProcessorFactory(
		typeStr,
		createDefaultConfig,
		component.WithTracesProcessor(createTracesProcessor),
	)
}

func createDefaultConfig() config.Processor {
	return &Config{
		ProcessorSettings: config.NewProcessorSettings(config.NewComponentID(typeStr)),
		Exporter:          "datadog",
	}
}

type Config struct {
	config.ProcessorSettings `mapstructure:",squash"` // squash ensures fields are correctly decoded in embedded struct

	Exporter string `mapstructure:"exporter"`
}

// Validate validates the configuration and returns an error if invalid.
func (cfg *Config) Validate() error {
	// TODO
	return nil
}

func createTracesProcessor(_ context.Context, params component.ProcessorCreateSettings, cfg config.Processor, next consumer.Traces) (component.TracesProcessor, error) {
	return &processor{
		logger: params.Logger,
		cfg:    cfg.(*Config),
		next:   next,
	}, nil
}

var _ component.TracesProcessor = (*processor)(nil)

type processor struct {
	logger *zap.Logger
	cfg    *Config
	next   consumer.Traces
	out    component.MetricsExporter
}

func log(p *processor, msg string, params ...interface{}) {
	p.logger.Info(fmt.Sprintf("DEBUG LINE "+msg, params...))
}

func (p *processor) Start(ctx context.Context, host component.Host) error {
	mexps, ok := host.GetExporters()[config.MetricsDataType]
	if !ok {
		return errors.New("stats can not be sent if no metrics exporters are registered")
	}
	for cid, exp := range mexps {
		if cid.String() == p.cfg.Exporter {
			p.out, ok = exp.(component.MetricsExporter)
			if !ok {
				continue
			}
			break
		}
	}
	if p.out == nil {
		return fmt.Errorf("exporter named %q not found", p.cfg.Exporter)
	}
	return nil
}

func (p *processor) Shutdown(ctx context.Context) error {
	return nil
}

func (p *processor) Capabilities() consumer.Capabilities {
	return consumer.Capabilities{MutatesData: false}
}

func (p *processor) ConsumeTraces(ctx context.Context, traces ptrace.Traces) error {
	log(p, "Got %d spans.", traces.SpanCount())

	data := pmetric.NewMetrics()
	rm := data.ResourceMetrics().AppendEmpty()
	rm.Resource().Attributes().InsertBool("_dd.apmstats", true)
	m := rm.ScopeMetrics().AppendEmpty().Metrics().AppendEmpty()
	m.SetName("test.metric.from.processor")
	m.SetDataType(pmetric.MetricDataTypeSum)
	m.Sum().DataPoints().AppendEmpty().SetIntVal(5)

	if err := p.out.ConsumeMetrics(ctx, data); err != nil {
		log(p, "Error sending metrics: %v", err)
	}
	return p.next.ConsumeTraces(ctx, traces)
}
