package action

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/loadept/pipeBot/pkg/throwable"
	"github.com/loadept/pipeBot/pkg/util"
)

type Ban struct{}

func (b *Ban) Execute(s *discordgo.Session, m *discordgo.MessageCreate) {
	roles, err := s.GuildRoles(m.GuildID)
	if err != nil {
		throwable.SendErrorEmbed(s, m.ChannelID, throwable.SomethingWentWrongRole.Error())
		return
	}

	member, err := s.GuildMember(m.GuildID, m.Author.ID)
	if err != nil {
		throwable.SendErrorEmbed(s, m.ChannelID, throwable.SomethingWentWrongMember.Error())
		return
	}

	roleMapByID := make(map[string]*discordgo.Role)
	for _, role := range roles {
		roleMapByID[role.ID] = role
	}

	if !util.IsAdmin(member, roleMapByID) {
		throwable.SendErrorEmbed(s, m.ChannelID, throwable.WithoutSufficientPermissions.Error())
		return
	}

	err = s.GuildBanCreate(m.GuildID, m.Mentions[0].ID, 1)
	if err != nil {
		throwable.SendErrorEmbed(s, m.ChannelID, "Error banning member.")
		return
	}

	embed := &discordgo.MessageEmbed{
		Title:       "🟠 Banned member",
		Description: fmt.Sprintf("Member %s has been banned by %s", m.Mentions[0].Username, m.Author.Username),
		Color:       0xeff6400,
	}
	s.ChannelMessageSendEmbed(m.ChannelID, embed)
}
