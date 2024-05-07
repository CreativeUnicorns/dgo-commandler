// Package utils provides utility functions to assist with tasks common in the commandler package,
// especially those involving interactions with the Discord API.
package utils

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

// IsUserBot checks whether the user identified by userID is a bot.
// It returns true if the user is a bot, or false along with an error if the user information could not be retrieved.
func IsUserBot(session *discordgo.Session, userID string) (bool, error) {
	user, err := session.User(userID)
	if err != nil {
		Logger.Error("Error retrieving user", "userID", userID, "error", err)
		return false, fmt.Errorf("error retrieving user: %w", err)
	}

	return user.Bot, nil
}

// SendDirectMessage sends a message to the direct message channel of the specified user.
// It first checks whether the user is a bot, to avoid unnecessary messages to bots.
func SendDirectMessage(session *discordgo.Session, userID string, message string) error {
	isBot, err := IsUserBot(session, userID)
	if err != nil {
		Logger.Error("Error checking if user is a bot", "userID", userID, "error", err)
		return err
	}
	Logger.Info("Checked bot status", "userID", userID, "isBot", isBot)

	channel, err := session.UserChannelCreate(userID)
	if err != nil {
		Logger.Error("Error creating DM channel", "userID", userID, "error", err)
		return fmt.Errorf("error creating DM channel: %w", err)
	}

	_, err = session.ChannelMessageSend(channel.ID, message)
	if err != nil {
		Logger.Error("Error sending message", "channelID", channel.ID, "message", message, "error", err)
		return fmt.Errorf("error sending message: %w", err)
	}

	return nil
}

// ComesFromDM checks if a message was sent from a direct message channel.
// It returns true if the message originates from a DM, or false and an error if the channel details could not be confirmed.
func ComesFromDM(s *discordgo.Session, m *discordgo.MessageCreate) (bool, error) {
	channel, err := s.State.Channel(m.ChannelID)
	if err != nil {
		if channel, err = s.Channel(m.ChannelID); err != nil {
			Logger.Error("Failed to get channel details", "channelID", m.ChannelID, "error", err)
			return false, err
		}
	}

	return channel.Type == discordgo.ChannelTypeDM, nil
}

// GetChannelName retrieves the name of a channel from its ID, utilizing cached data if available.
// It returns the channel name or an empty string and an error if the channel details could not be retrieved.
func GetChannelName(s *discordgo.Session, channelID string) (string, error) {
	channel, err := s.State.Channel(channelID)
	if err != nil {
		Logger.Error("Failed to get channel details", "channelID", channelID, "error", err)
		return "", err
	}
	return channel.Name, nil
}

// GetGuildName retrieves the name of a guild from its ID, using cached data if available.
// It returns the guild name or an empty string and an error if the guild details could not be retrieved.
func GetGuildName(s *discordgo.Session, guildID string) (string, error) {
	guild, err := s.State.Guild(guildID)
	if err != nil {
		Logger.Error("Failed to get guild details", "guildID", guildID, "error", err)
		return "", err
	}
	return guild.Name, nil
}
