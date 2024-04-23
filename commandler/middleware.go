// Package commandler provides middleware support for handling Discord interaction commands.
package commandler

import (
	"github.com/CreativeUnicorns/dgo-commandler/utils"
	"github.com/bwmarrin/discordgo"
)

// CommandHandler defines a function type that processes a Discord interaction event.
type CommandHandler func(s *discordgo.Session, i *discordgo.InteractionCreate)

// Middleware defines a function type for middleware, which modifies or extends the behavior of a CommandHandler.
type Middleware func(CommandHandler) CommandHandler

// ApplyDefaultMiddleware sets up a default middleware chain for a given CommandHandler.
// Additional middlewares can be appended to the chain.
func ApplyDefaultMiddleware(handler CommandHandler, additionalMiddlewares ...Middleware) CommandHandler {
	// Start with a default set of middleware (here, just LoggerMiddleware)
	middlewares := []Middleware{LoggerMiddleware}

	// Append any additional middleware provided
	middlewares = append(middlewares, additionalMiddlewares...)

	// Chain all middlewares
	return ChainMiddlewares(handler, middlewares...)
}

// LoggerMiddleware logs information about the execution of Discord commands before the command handler is invoked.
func LoggerMiddleware(next CommandHandler) CommandHandler {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		logCommandExecution(s, i)
		next(s, i)
	}
}

// logCommandExecution logs detailed information about a command execution, such as command name,
// user ID, user name, and other relevant details.
func logCommandExecution(s *discordgo.Session, i *discordgo.InteractionCreate) {
	commandName := i.ApplicationCommandData().Name
	user := i.Member.User
	userID := user.ID
	userName := user.Username
	userNick := i.Member.Nick

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

// ChainMiddlewares returns a new CommandHandler that is the result of chaining multiple Middleware functions.
// Each Middleware wraps the CommandHandler passed to it, potentially modifying its execution or adding new behaviors.
func ChainMiddlewares(handler CommandHandler, middlewares ...Middleware) CommandHandler {
	for _, middleware := range middlewares {
		handler = middleware(handler)
	}
	return handler
}
