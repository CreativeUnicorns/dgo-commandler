// File: commandler/middleware.go
// Description: Middleware for command handlers.
package commandler

import (
	"fmt"
	"log"
	"log/slog"

	"github.com/CreativeUnicorns/dgo-commandler/utils"
	"github.com/bwmarrin/discordgo"
)

// CommandHandlerMiddleware defines the type for middleware functions wrapping command handlers.
type CommandHandlerMiddleware func(handler func(s *discordgo.Session, i *discordgo.InteractionCreate)) func(s *discordgo.Session, i *discordgo.InteractionCreate)

// LoggerMiddleware logs information about the command execution before passing execution to the actual handler.
func LoggerMiddleware(handler func(s *discordgo.Session, i *discordgo.InteractionCreate)) func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		logCommandExecution(s, i)
		handler(s, i)
	}
}

// logCommandExecution handles the creation of log entries for command executions.
func logCommandExecution(s *discordgo.Session, i *discordgo.InteractionCreate) {
	commandName := i.ApplicationCommandData().Name
	user := i.Member.User
	userID := user.ID
	userName := user.Username
	userNick := i.Member.Nick

	if i.GuildID == "" {
		log.Printf("[middleware] Executing command: /%s by %s in DMs (userID='%s', userName='%s', userNick='%s')", commandName, userName, userID, userName, userNick)
	} else {
		channelName, errChannel := utils.GetChannelName(s, i.ChannelID)
		guildName, errGuild := utils.GetGuildName(s, i.GuildID)
		if errChannel != nil || errGuild != nil {
			slog.Error(fmt.Sprintf("[middleware] Error retrieving channel or guild name: %v %v", errChannel, errGuild))
		}
		slog.Info(fmt.Sprintf("[middleware] Executing command: /%s by %s in guild %s (%s) in %s (%s) (userID='%s', userName='%s', userNick='%s')", commandName, userName, guildName, i.GuildID, channelName, i.ChannelID, userID, userName, userNick))
	}
}
