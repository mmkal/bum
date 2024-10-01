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

package command

import (
	"github.com/kardianos/osext"
	"os"
)

// ChaosbumExecutable is a representation of Chaosbum executable
type ChaosbumExecutable struct {
}

// ExecutablePath implements command.CurrentExecutable.ExecutablePath
func (e ChaosbumExecutable) ExecutablePath() (string, error) {
	return osext.Executable()
}

// EnsureFileAbsent ensures that a file is absent, returning an error otherwise
func EnsureFileAbsent(path string) error {
	err := os.Remove(path)

	// If it's an IsNotExist error, we can ignore it, since it
	// satisfies the contract of the file being absent
	if os.IsNotExist(err) {
		return nil
	}

	return err
}
