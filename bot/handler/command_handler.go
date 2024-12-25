package handler

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/loadept/pipeBot/internal/invoker"
)

type Commands struct{}

func (c *Commands) NotifyMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	content := strings.Split(m.Content, " ")
	command := content[0]

	bot := invoker.NewBot()

	bot.Invoker(command, s, m)
}
