// Copyright 2016 Netflix, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mock

import "github.com/Netflix/chaosbum/v2"

// ConfigGetter implements chaosbum.Getter
type ConfigGetter struct {
	Config chaosbum.AppConfig
}

// NewConfigGetter returns a mock config getter that always returns the specified config
func NewConfigGetter(config chaosbum.AppConfig) ConfigGetter {
	return ConfigGetter{Config: config}
}

// DefaultConfigGetter returns a mock config getter that always returns the same config
func DefaultConfigGetter() ConfigGetter {
	return ConfigGetter{
		Config: chaosbum.AppConfig{
			Enabled:                        true,
			RegionsAreIndependent:          true,
			MeanTimeBetweenKillsInWorkDays: 5,
			MinTimeBetweenKillsInWorkDays:  1,
			Grouping:                       chaosbum.Cluster,
			Exceptions:                     nil,
		},
	}
}

// Get implements chaosbum.Getter.Get
func (c ConfigGetter) Get(app string) (*chaosbum.AppConfig, error) {
	return &c.Config, nil
}
