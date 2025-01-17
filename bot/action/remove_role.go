package action

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/loadept/pipeBot/pkg/throwable"
	"github.com/loadept/pipeBot/pkg/util"
)

type RemoveRole struct{}

func (ar *RemoveRole) Execute(s *discordgo.Session, m *discordgo.MessageCreate) {
	logsChannel := util.GetGuildChannel(s, m.GuildID, ".*logs?")
	if len(logsChannel) == 0 {
		throwable.SendErrorEmbed(s, logsChannel, throwable.SomethingWentWrong.Error())
	}

	targetRole, err := util.ValidateRole(s, m)
	if err != nil {
		throwable.SendErrorEmbed(s, logsChannel, err.Error())
		return
	}

	if err := s.GuildMemberRoleRemove(m.GuildID, m.Mentions[0].ID, targetRole.ID); err != nil {
		throwable.SendErrorEmbed(s, logsChannel, "Failed to remove role.")
		return
	}

	embed := &discordgo.MessageEmbed{
		Title:       "🟡 Member removed to role",
		Description: fmt.Sprintf("Member **%s** has been removed from role **%s**", m.Mentions[0].Username, targetRole.Name),
		Color:       0xe8ff00,
	}
	s.ChannelMessageSendEmbed(logsChannel, embed)
}
