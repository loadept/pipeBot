package invoker

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/loadept/pipeBot/bot/action"
	"github.com/loadept/pipeBot/internal/command"
)

type Bot struct {
	commands map[string]command.Command
}

func NewBot() *Bot {
	bot := &Bot{commands: make(map[string]command.Command)}
	bot.commands["|ban"] = &action.Ban{}
	bot.commands["|role+"] = &action.AddRole{}
	bot.commands["|role-"] = &action.RemoveRole{}
	bot.commands["|role"] = &action.ListRole{}

	return bot
}

func (b *Bot) Invoker(command string, s *discordgo.Session, m *discordgo.MessageCreate) {
	if cmd, exists := b.commands[command]; exists {
		cmd.Execute(s, m)
	} else {
		fmt.Println("Invalid Command")
	}
}
