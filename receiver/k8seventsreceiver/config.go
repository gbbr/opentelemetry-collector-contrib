// Copyright  OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package k8seventsreceiver

import (
	"go.opentelemetry.io/collector/config"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/k8sconfig"
)

// Config defines configuration for kubernetes events receiver.
type Config struct {
	config.ReceiverSettings `mapstructure:",squash"`
	k8sconfig.APIConfig     `mapstructure:",squash"`

	// List of ‘namespaces’ to collect events from.
	Namespaces []string `mapstructure:"namespaces"`
}

func (cfg *Config) Validate() error {
	if err := cfg.ReceiverSettings.Validate(); err != nil {
		return err
	}
	if err := cfg.APIConfig.Validate(); err != nil {
		return err
	}
	return nil
}
