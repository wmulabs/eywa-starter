package bootstrap

import (
	eywa "github.com/wmulabs/eywa"
	"github.com/wmulabs/eywa-starter/internal/application/converters"
)

// registerInboundConverters wires the Receptors that turn HTTP payloads into Pulses. A Link references
// one by name via WithInboundConverter (see events.go).
func registerInboundConverters(weave *eywa.Weave) {
	weave.RegisterReceptor("chat_inbound", converters.NewChatReceptor())
}
