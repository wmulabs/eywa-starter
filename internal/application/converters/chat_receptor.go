// Package converters holds Receptor implementations: they turn a raw inbound payload (the HTTP body)
// into one or more Pulses for the Weave. Add one per channel/event shape your app accepts.
package converters

import (
	"context"
	"fmt"

	eywa "github.com/wmulabs/eywa"
)

type ChatReceptor struct{}

func NewChatReceptor() eywa.Receptor { return &ChatReceptor{} }

func (c *ChatReceptor) GetName() string { return "chat_inbound" }

// Convert maps a JSON body of {"user": "...", "message": "..."} into a Pulse.
func (c *ChatReceptor) Convert(_ context.Context, _ string, payload map[string]any) ([]*eywa.Pulse, error) {
	user, _ := payload["user"].(string)
	if user == "" {
		return nil, fmt.Errorf("payload requires a non-empty \"user\"")
	}
	message, _ := payload["message"].(string)

	pulse := eywa.NewPulse(eywa.MemoryKey{Channel: "api", User: user}).
		WithUserMessage(message).
		WithSource("api").
		Build()

	return []*eywa.Pulse{pulse}, nil
}
