// Copyright The OpenTelemetry Authors
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

// Package log provides a unified logger.
//
// Deprecated: This package is no longer supported.
package log

import (
	"github.com/go-logr/logr"
	"github.com/go-logr/zapr"
	"go.uber.org/zap"
)

// Logger is a global logger used for internal logging.
var Logger logr.Logger

// Init initializes [Logger] to be a production Zap logger.
func Init() error {
	zapLog, err := zap.NewProduction()
	if err != nil {
		return err
	}

	Logger = zapr.NewLogger(zapLog)
	return nil
}
