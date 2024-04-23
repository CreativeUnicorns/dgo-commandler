// commandler/interaction.go

package commandler

import (
	"github.com/bwmarrin/discordgo"
)

// InteractionCommand represents a Discord slash command
type InteractionCommand struct {
	Name                     string
	Description              string
	Options                  []*discordgo.ApplicationCommandOption
	DefaultMemberPermissions int64
	DMPermission             bool
	NSFW                     bool
	Handler                  func(s *discordgo.Session, i *discordgo.InteractionCreate)
}

var interactionCommands = make([]*InteractionCommand, 0)

// RegisterInteractionCommand registers a new slash command
// Modify the RegisterInteractionCommand function to apply middleware
func RegisterInteractionCommand(cmd *InteractionCommand, additionalMiddlewares ...Middleware) {
	// Apply default middleware, potentially including additional specified middleware
	wrappedHandler := ApplyDefaultMiddleware(cmd.Handler, additionalMiddlewares...)
	cmd.Handler = wrappedHandler

	interactionCommands = append(interactionCommands, cmd)
}

// GetInteractionCommands returns all registered commands
func GetInteractionCommands() []*InteractionCommand {
	return interactionCommands
}
