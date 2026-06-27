package bootstrap

import (
	"context"
	"log"
	"time"

	eywa "github.com/wmulabs/eywa"
)

// seedSpirits creates the agent definitions this app ships with. Edit the Spirit to make it yours:
// its prompt, model, and the Actions it may call. Re-creating an existing Spirit is expected on
// restart and is logged, not fatal.
func seedSpirits(ctx context.Context, repo eywa.SpiritRepository) error {
	assistant := &eywa.Spirit{
		Name:         "assistant",
		Description:  "Starter assistant",
		SystemPrompt: "You are a concise, helpful assistant built with Eywa. Use get_time for the current time and catalog_lookup to search the product catalog.",
		ModelConfig: eywa.SpiritModel{
			Provider:    "gemini",
			Model:       "gemini-2.5-flash",
			Temperature: 0.7,
			MaxTokens:   500,
		},
		AllowedActions: []eywa.AllowedAction{{Name: "get_time"}, {Name: "catalog_lookup"}},
		IsActive:       true,
		CreatedAt:      time.Now(),
	}

	if err := repo.Create(ctx, assistant); err != nil {
		log.Printf("seed spirit %q: %v (continuing — likely already exists)", assistant.Name, err)
	}
	return nil
}
