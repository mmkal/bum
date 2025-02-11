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

// Package deps holds a set of interfaces
package deps

import (
	"github.com/Netflix/chaosbum/v2"
	"github.com/Netflix/chaosbum/v2/clock"
	"github.com/Netflix/chaosbum/v2/config"
	"github.com/Netflix/chaosbum/v2/deploy"
	"github.com/Netflix/chaosbum/v2/schedule"
)

var (
	// GetTrackers returns a list of trackers
	// This variable must be set in the init() method of a module imported by
	// the main module.
	GetTrackers func(*config.Bum) ([]chaosbum.Tracker, error)

	// GetErrorCounter returns an error counter
	GetErrorCounter func(*config.Bum) (chaosbum.ErrorCounter, error)

	// GetDecryptor returns a decryptor
	GetDecryptor func(*config.Bum) (chaosbum.Decryptor, error)

	// GetEnv returns info about the deployed environment
	GetEnv func(*config.Bum) (chaosbum.Env, error)

	// GetOutage returns an interface for checking if there is an outage
	GetOutage func(*config.Bum) (chaosbum.Outage, error)

	// GetConstrainer returns an interface for constraining the schedule
	GetConstrainer func(*config.Bum) (schedule.Constrainer, error)
)

// Deps are a common set of external dependencies
type Deps struct {
	BumCfg  *config.Bum
	Checker    chaosbum.Checker
	ConfGetter chaosbum.AppConfigGetter
	Cl         clock.Clock
	Dep        deploy.Deployment
	T          chaosbum.Terminator
	Trackers   []chaosbum.Tracker
	Ou         chaosbum.Outage
	ErrCounter chaosbum.ErrorCounter
	Env        chaosbum.Env
}
