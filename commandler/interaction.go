// Package commandler provides structures and functions to create and manage Discord slash commands.
package commandler

import (
	"github.com/bwmarrin/discordgo"
)

// InteractionCommand represents a Discord slash command. It contains the command's name,
// description, options, permissions, and the handler function to execute when the command is triggered.
type InteractionCommand struct {
	Name                     string
	Description              string
	Options                  []*discordgo.ApplicationCommandOption
	DefaultMemberPermissions int64
	DMPermission             bool
	NSFW                     bool
	Handler                  func(s *discordgo.Session, i *discordgo.InteractionCreate)
}

// interactionCommands holds a list of all registered InteractionCommand instances.
var interactionCommands = make([]*InteractionCommand, 0)

// RegisterInteractionCommand adds a new InteractionCommand to the list of registered commands.
// It wraps the command's Handler with default and any additional provided middlewares.
func RegisterInteractionCommand(cmd *InteractionCommand, additionalMiddlewares ...Middleware) {
	wrappedHandler := ApplyDefaultMiddleware(cmd.Handler, additionalMiddlewares...)
	cmd.Handler = wrappedHandler

	interactionCommands = append(interactionCommands, cmd)
}

// GetInteractionCommands returns a slice containing all registered InteractionCommands.
func GetInteractionCommands() []*InteractionCommand {
	return interactionCommands
}
