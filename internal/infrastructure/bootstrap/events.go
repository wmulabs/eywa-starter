package bootstrap

import eywa "github.com/wmulabs/eywa"

// registerEvents maps inbound event keys to the Spirit that handles them. Add a Link per event your
// app accepts; the REST server routes POST /events/{key} here.
func registerEvents(weave *eywa.Weave) {
	weave.RegisterEventConfiguration(
		eywa.NewLink("chat").
			WithInboundConverter("chat_inbound").
			WithScouts("business_hours").
			WithDefaultSpirit("assistant").
			Build(),
	)
}
