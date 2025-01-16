package invoker

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/loadept/pipeBot/internal/command"
)

type Bot struct {
	commands map[string]command.Command
}

func NewBot() *Bot {
	bot := &Bot{commands: make(map[string]command.Command)}

	return bot
}

func (b *Bot) SetCommand(action string, command command.Command) {
	b.commands[action] = command
}

func (b *Bot) Invoker(command string, s *discordgo.Session, m *discordgo.MessageCreate) {
	if cmd, exists := b.commands[command]; exists {
		cmd.Execute(s, m)
	} else {
		fmt.Println("Invalid Command")
	}
}
