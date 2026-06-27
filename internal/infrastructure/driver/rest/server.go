// Package rest is the inbound HTTP adapter. A single registrar (eywafiber.RegisterRoutes) mounts the
// whole surface: health + event ingestion are open (webhook-style); the management API and app-token
// management sit behind API-key auth when an admin key is configured. Event ingestion can optionally
// require an app token.
package rest

import (
	"context"

	fiberlib "github.com/gofiber/fiber/v2"
	eywa "github.com/wmulabs/eywa"
	eywafiber "github.com/wmulabs/eywa/fiber"
	// Example: WhatsApp webhook signature verification (add channels/whatsapp dep to use):
	// twilio "github.com/wmulabs/eywa/channels/whatsapp/twilio"
)

// Deps is what the HTTP layer needs. Echo/Chronicle/AppToken are the read/management ports.
type Deps struct {
	Weave             *eywa.Weave
	Port              string
	AdminAPIKey       string
	RequireEventToken bool
	Echo              eywa.EchoRepository
	EchoQuery         eywa.EchoQueryRepository
	Chronicle         eywa.ChronicleQueryRepository
	AppToken          eywa.AppTokenRepository
	// TwilioAuthToken string // uncomment to enable the Twilio signature verifier below
}

type Server struct {
	app  *fiberlib.App
	port string
}

func NewServer(d Deps) (*Server, error) {
	app := fiberlib.New(fiberlib.Config{DisableStartupMessage: true})

	deps := eywafiber.RouteDeps{}

	// Management API (Spirit CRUD is NOT exposed — Spirits live in code). Mounted behind the API key
	// only when ADMIN_API_KEY is set; each group needs its repo.
	if d.AdminAPIKey != "" {
		deps.APIKeys = map[string]string{d.AdminAPIKey: "admin"}
		deps.ChronicleQueryRepo = d.Chronicle
		deps.EchoRepo = d.Echo
		deps.EchoQueryRepo = d.EchoQuery
		deps.AppTokenRepo = d.AppToken // exposes POST/GET/DELETE /api/v1/app-tokens

		// Optionally require an app token on the event endpoint (tokens minted via /api/v1/app-tokens).
		if d.RequireEventToken {
			deps.EventAuth = []eywa.TokenValidator{eywa.NewAppTokenValidator(d.AppToken)}
		}
	}

	// Example — verify Twilio webhook signatures on events (no token needed; Twilio signs automatically):
	//   deps.EventVerifiers = []eywa.RequestVerifier{twilio.NewSignatureVerifier(d.TwilioAuthToken)}

	if err := eywafiber.RegisterRoutes(app, d.Weave, deps); err != nil {
		return nil, err
	}

	return &Server{app: app, port: d.Port}, nil
}

func (s *Server) Start() error {
	return s.app.Listen(":" + s.port)
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.app.ShutdownWithContext(ctx)
}
