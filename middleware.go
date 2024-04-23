// File: middleware.go
// Package: commandler
// Description: Middleware for command handlers.
package commandler

import (
	"github.com/CreativeUnicorns/dgo-commandler/utils"
	"github.com/bwmarrin/discordgo"
)

// CommandHandler is the type for functions that handle Discord interaction commands.
type CommandHandler func(s *discordgo.Session, i *discordgo.InteractionCreate)

// Middleware is the type for middleware functions.
type Middleware func(CommandHandler) CommandHandler

// ApplyDefaultMiddleware sets up the default middleware chain for command handlers,
// with the option to include additional middleware.
func ApplyDefaultMiddleware(handler CommandHandler, additionalMiddlewares ...Middleware) CommandHandler {
	// Start with a default set of middleware (here, just LoggerMiddleware)
	middlewares := []Middleware{LoggerMiddleware}

	// Append any additional middleware provided
	middlewares = append(middlewares, additionalMiddlewares...)

	// Chain all middlewares
	return ChainMiddlewares(handler, middlewares...)
}

// LoggerMiddleware logs information about the command execution before passing execution to the actual handler.
func LoggerMiddleware(next CommandHandler) CommandHandler {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		logCommandExecution(s, i)
		next(s, i)
	}
}

// logCommandExecution handles the creation of log entries for command executions.
func logCommandExecution(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// Structured logging of command execution details.
	commandName := i.ApplicationCommandData().Name
	user := i.Member.User
	userID := user.ID
	userName := user.Username
	userNick := i.Member.Nick

	// Check if it's in a guild or a DM.
	if i.GuildID == "" {
		utils.Logger.Info("Executing command in DMs", "command", commandName, "userID", userID, "userName", userName, "userNick", userNick)
	} else {
		channelName, errChannel := utils.GetChannelName(s, i.ChannelID)
		guildName, errGuild := utils.GetGuildName(s, i.GuildID)
		if errChannel != nil || errGuild != nil {
			utils.Logger.Error("Error retrieving channel or guild name", "channelError", errChannel, "guildError", errGuild)
		} else {
			utils.Logger.Info("Executing command in guild",
				"command", commandName, "userName", userName, "guildName", guildName, "channelName", channelName,
				"userID", userID, "guildID", i.GuildID, "channelID", i.ChannelID, "userNick", userNick)
		}
	}
}

// ChainMiddlewares creates a new CommandHandler by chaining multiple middlewares.
func ChainMiddlewares(handler CommandHandler, middlewares ...Middleware) CommandHandler {
	for _, middleware := range middlewares {
		handler = middleware(handler)
	}
	return handler
}
