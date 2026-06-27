# Eywa Starter — Agent Instructions

A production-shaped starter for building agents on [Eywa](https://github.com/wmulabs/eywa). Hexagonal,
one concern per file. Keep the structure when extending.

## Layout

```
cmd/api/                          entrypoint — thin: bootstrap.Initialize → start REST server
internal/
  application/                    YOUR domain plugged into Eywa
    actions/                      Action implementations (tools the Oracle calls)
                                    time_action.go        — no-dependency tool
                                    catalog_lookup_action.go — tool that calls a domain port
    scouts/                       Scout implementations (context enrichment)
    converters/                   Receptor implementations (payload → Pulse)
  domain/
    services/                     domain ports the actions depend on (e.g. CatalogService)
  infrastructure/
    bootstrap/                    wiring, ONE FILE PER CONCERN
      application.go              Application container + Initialize + Shutdown
      database.go                 Mongo + Redis connections
      repositories.go            Eywa repositories (incl. app-token repo)
      services.go                 domain services <- driven adapters
      engine.go                  builds the Weave; seeds Spirits; registers actions/scouts/converters/events
      actions.go / scouts.go     register the Action / Scout registries
      inbound_converters.go      register Receptors
      events.go                  register Links (event key → Receptor + Scouts + Spirit)
      spirits.go                 seed default Spirit definitions
    driven/                       outbound adapters implementing domain ports (catalog/ = in-memory stub)
    config/                       typed Config + Load()
    driver/rest/                  inbound HTTP — single RegisterRoutes (open events + authed management)
prompts/                          reference copies of Spirit prompts (not loaded at runtime)
deployments/ docker/ cloudbuild.yaml   ops
```

**Hexagonal flow to copy:** an Action (application) depends on a port in `domain/services`, which is
implemented by an adapter in `infrastructure/driven` and wired in `bootstrap/services.go`. The Action
never imports the concrete adapter. See `catalog_lookup_action.go` → `services.CatalogService` →
`driven/catalog`.

## Eywa vocabulary (use these names)

Weave (engine) · Spirit (agent) · Action (tool) · Scout (enricher) · Voice (outbound) · Receptor
(inbound) · Bond (lock) · Link (event config) · Keeper (scheduler) · Oracle (LLM). Never invent
synonyms ("agent" for Spirit, "tool" for Action, etc.).

## How to extend

- **New tool (no deps)** → add an `eywa.Action` in `application/actions/` (copy `time_action.go`),
  register in `bootstrap/actions.go`, list it in the Spirit's `AllowedActions` (`bootstrap/spirits.go`).
- **New tool that calls the outside world** → define a port in `domain/services/`, implement it in
  `infrastructure/driven/`, wire it in `bootstrap/services.go`, inject into the Action
  (copy `catalog_lookup_action.go`). Keep the Action depending on the port, not the adapter.
- **New Scout** → add to `application/scouts/`, register in `bootstrap/scouts.go`, opt in via
  `WithScouts(...)` on the Link (`events.go`).
- **New agent** → add a Spirit in `bootstrap/spirits.go` and a `Link` in `bootstrap/events.go`.
- **New inbound event** → add a `Receptor` (`application/converters/` + `bootstrap/inbound_converters.go`)
  and a `Link` (`events.go`). Served at `POST /api/v1/events/{key}` (+ `/stream`, `/async`).
- **More of the reasoning loop** (stall detection, reflection, grounding, plan, durable execution) →
  builder options in `bootstrap/engine.go`. See Eywa docs/reasoning.md.
- **Auth** → management behind `ADMIN_API_KEY`; event ingestion open by default, or require an app
  token (`REQUIRE_EVENT_TOKEN=true`) / verify channel signatures (`EventVerifiers` in `driver/rest`).
  See Eywa docs/authentication.md.

## Conventions

- Comments only for non-obvious WHY. No godoc on self-explanatory functions.
- Config from env via `config.Load()` — no `os.Getenv` scattered in business code.
- Errors wrapped with context (`fmt.Errorf("...: %w", err)`).
- Security: event/health routes are open (webhook-style); the management API is behind `ADMIN_API_KEY`.
  Never expose Spirit-mutation routes unauthenticated — that's an infra-level prompt-injection surface.
