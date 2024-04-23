// File: utils/helpers.go
// Package: utils

package utils

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

// IsUserBot checks if the user with the given userID is a bot.
func IsUserBot(session *discordgo.Session, userID string) (bool, error) {
	// Retrieve the user information
	user, err := session.User(userID)
	if err != nil {
		return false, fmt.Errorf("error retrieving user: %w", err)
	}

	// Return true if the user is a bot
	return user.Bot, nil
}

func SendDirectMessage(session *discordgo.Session, userID string, message string) error {
	// Check if the user is a bot
	isBot, err := IsUserBot(session, userID)
	if err != nil {
		log.Printf("Error checking if user is a bot: %s\n", err)
	} else {
		log.Printf("Is the user a bot? %t\n", isBot)
	}
	// Create a DM channel with this user
	channel, err := session.UserChannelCreate(userID)
	if err != nil {
		return fmt.Errorf("error creating DM channel: %w", err)
	}

	// Send a message to this channel
	_, err = session.ChannelMessageSend(channel.ID, message)
	if err != nil {
		return fmt.Errorf("error sending message: %w", err)
	}

	return nil
}

// ComesFromDM returns true if a message comes from a DM channel
func ComesFromDM(s *discordgo.Session, m *discordgo.MessageCreate) (bool, error) {
	channel, err := s.State.Channel(m.ChannelID)
	if err != nil {
		if channel, err = s.Channel(m.ChannelID); err != nil {
			return false, err
		}
	}

	return channel.Type == discordgo.ChannelTypeDM, nil
}

// GetChannelName retrieves the name of a channel by its ID.
func GetChannelName(session *discordgo.Session, channelID string) (string, error) {
	channel, err := session.State.Channel(channelID)
	if err != nil {
		return "", err
	}
	return channel.Name, nil
}

// GetGuildName retrieves the name of a guild by its ID.
func GetGuildName(session *discordgo.Session, guildID string) (string, error) {
	guild, err := session.State.Guild(guildID)
	if err != nil {
		return "", err
	}
	return guild.Name, nil
}
