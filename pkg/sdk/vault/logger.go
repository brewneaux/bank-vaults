// Copyright © 2020 Banzai Cloud
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

package vault

import (
	"context"
)

// Logger is a unified interface for various logging use cases and practices, including:
// 		- leveled logging
// 		- structured logging
type Logger interface {
	// Trace logs a Trace event.
	//
	// Even more fine-grained information than Debug events.
	// Loggers not supporting this level should fall back to Debug.
	Trace(msg string, fields ...map[string]interface{})

	// Debug logs a Debug event.
	//
	// A verbose series of information events.
	// They are useful when debugging the system.
	Debug(msg string, fields ...map[string]interface{})

	// Info logs an Info event.
	//
	// General information about what's happening inside the system.
	Info(msg string, fields ...map[string]interface{})

	// Warn logs a Warn(ing) event.
	//
	// Non-critical events that should be looked at.
	Warn(msg string, fields ...map[string]interface{})

	// Error logs an Error event.
	//
	// Critical events that require immediate attention.
	// Loggers commonly provide Fatal and Panic levels above Error level,
	// but exiting and panicing is out of scope for a logging library.
	Error(msg string, fields ...map[string]interface{})
}

// LoggerContext is an optional interface that MAY be implemented by a Logger.
// It is similar to Logger, but it receives a context as the first parameter.
// An implementation MAY extract information from the context and annotate the log context with it.
//
// LoggerContext MAY honor the deadline carried by the context, but that's not a hard requirement.
type LoggerContext interface {
	// TraceContext logs a Trace event.
	//
	// Even more fine-grained information than Debug events.
	// Loggers not supporting this level should fall back to Debug.
	TraceContext(ctx context.Context, msg string, fields ...map[string]interface{})

	// DebugContext logs a Debug event.
	//
	// A verbose series of information events.
	// They are useful when debugging the system.
	DebugContext(ctx context.Context, msg string, fields ...map[string]interface{})

	// InfoContext logs an Info event.
	//
	// General information about what's happening inside the system.
	InfoContext(ctx context.Context, msg string, fields ...map[string]interface{})

	// WarnContext logs a Warn(ing) event.
	//
	// Non-critical events that should be looked at.
	WarnContext(ctx context.Context, msg string, fields ...map[string]interface{})

	// ErrorContext logs an Error event.
	//
	// Critical events that require immediate attention.
	// Loggers commonly provide Fatal and Panic levels above Error level,
	// but exiting and panicing is out of scope for a logging library.
	ErrorContext(ctx context.Context, msg string, fields ...map[string]interface{})
}

// Fields is used to define structured fields which are appended to log events.
// It can be used as a shorthand for map[string]interface{}.
type Fields map[string]interface{}

// LogFunc records a log event.
type LogFunc func(msg string, fields ...map[string]interface{})

// LogContextFunc records a log event.
type LogContextFunc func(ctx context.Context, msg string, fields ...map[string]interface{})

// noopLogger is a no-op logger that discards all received log events.
//
// It implements both Logger and LoggerContext interfaces.
type noopLogger struct{}

func (noopLogger) Trace(_ string, _ ...map[string]interface{}) {}
func (noopLogger) Debug(_ string, _ ...map[string]interface{}) {}
func (noopLogger) Info(_ string, _ ...map[string]interface{})  {}
func (noopLogger) Warn(_ string, _ ...map[string]interface{})  {}
func (noopLogger) Error(_ string, _ ...map[string]interface{}) {}

func (noopLogger) TraceContext(_ context.Context, _ string, _ ...map[string]interface{}) {}
func (noopLogger) DebugContext(_ context.Context, _ string, _ ...map[string]interface{}) {}
func (noopLogger) InfoContext(_ context.Context, _ string, _ ...map[string]interface{})  {}
func (noopLogger) WarnContext(_ context.Context, _ string, _ ...map[string]interface{})  {}
func (noopLogger) ErrorContext(_ context.Context, _ string, _ ...map[string]interface{}) {}

// LevelEnabler checks if a level is enabled in a logger.
// If the logger cannot reliably decide the correct level this method MUST return true.
type LevelEnabler interface {
	// LevelEnabled checks if a level is enabled in a logger.
	LevelEnabled(level Level) bool
}

// Level represents the same levels as defined in Logger.
type Level uint32

// Levels as defined in Logger.
const (
	// Even more fine-grained information than Debug events.
	// Loggers not supporting this level should fall back to Debug.
	Trace Level = iota

	// A verbose series of information events.
	// They are useful when debugging the system.
	Debug

	// General information about what's happening inside the system.
	Info

	// Non-critical events that should be looked at.
	Warn

	// Critical events that require immediate attention.
	Error
)
