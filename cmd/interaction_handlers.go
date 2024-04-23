package cmd

import (
	"github.com/CreativeUnicorns/dgo-commandler/internal/logger"
	"github.com/CreativeUnicorns/dgo-commandler/pkg"
	"github.com/bwmarrin/discordgo"
)

// InteractionHandler manages command executions and middleware applications.
type InteractionHandler struct {
	Session *discordgo.Session
}

// HandleCommand parses and dispatches commands to their respective handlers.
func (h *InteractionHandler) HandleCommand(i *discordgo.InteractionCreate) {
	cmdName := i.ApplicationCommandData().Name

	cmd := pkg.GetCommandByName(cmdName)
	if cmd == nil {
		logger.LogInfo("Command not found: ", cmdName)
		return
	}

	cmd.Handler(h.Session, i)
}
