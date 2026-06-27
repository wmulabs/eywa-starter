// Package actions holds this app's Action implementations (tools the Oracle can call). Copy the shape
// of TimeAction for your own: name, description, JSON-schema parameters, and an Execute that returns a
// string the model reads back.
package actions

import (
	"context"
	"time"

	eywa "github.com/wmulabs/eywa"
)

type TimeAction struct{}

func NewTimeAction() eywa.Action { return &TimeAction{} }

func (a *TimeAction) GetName() string               { return "get_time" }
func (a *TimeAction) GetDescription() string        { return "Returns the current UTC time in RFC3339." }
func (a *TimeAction) GetParameters() map[string]any { return map[string]any{"type": "object"} }
func (a *TimeAction) IsCritical() bool              { return false }
func (a *TimeAction) GetCategory() eywa.ActionCategory {
	return eywa.ActionRetrieval
}
func (a *TimeAction) Validate(map[string]any) error { return nil }
func (a *TimeAction) Execute(context.Context, map[string]any) (string, error) {
	return time.Now().UTC().Format(time.RFC3339), nil
}
