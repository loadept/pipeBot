package action

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/loadept/pipeBot/pkg/throwable"
	"github.com/loadept/pipeBot/pkg/util"
)

type AddRole struct{}

func (ar *AddRole) Execute(s *discordgo.Session, m *discordgo.MessageCreate) {
	targetRole, err := util.ValidateRole(s, m)
	if err != nil {
		throwable.SendErrorEmbed(s, m.ChannelID, err.Error())
		return
	}

	if err := s.GuildMemberRoleAdd(m.GuildID, m.Mentions[0].ID, targetRole.ID); err != nil {
		throwable.SendErrorEmbed(s, m.ChannelID, "Failed to establish role.")
		return
	}

	embed := &discordgo.MessageEmbed{
		Title:       "🔵 Member add to role",
		Description: fmt.Sprintf("Member **%s** has been assigned to the role **%s**", m.Mentions[0].Username, targetRole.Name),
		Color:       0x00d8ff,
	}
	s.ChannelMessageSendEmbed(m.ChannelID, embed)
}
