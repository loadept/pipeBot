package handler

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/loadept/pipeBot/bot/action"
	"github.com/loadept/pipeBot/internal/invoker"
)

type Commands struct{}

func (c *Commands) NotifyMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	content := strings.Split(m.Content, " ")
	command := content[0]
	if !strings.HasPrefix(command, "|") {
		return
	}

	bot := invoker.NewBot()

	bot.SetCommand("|ban", &action.Ban{})
	bot.SetCommand("|role+", &action.AddRole{})
	bot.SetCommand("|role-", &action.RemoveRole{})
	bot.SetCommand("|role", &action.ListRole{})

	bot.Invoker(command, s, m)
}
