package action

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/loadept/pipeBot/pkg/throwable"
	"github.com/loadept/pipeBot/pkg/util"
)

type ListRole struct{}

func (l *ListRole) Execute(s *discordgo.Session, m *discordgo.MessageCreate) {
	if err := util.CheckMentions(s, m.ChannelID, m.Mentions); err != nil {
		throwable.SendErrorEmbed(s, m.ChannelID, err.Error())
		return
	}
	roles, err := s.GuildRoles(m.GuildID)
	if err != nil {
		throwable.SendErrorEmbed(s, m.ChannelID, throwable.SomethingWentWrongRole.Error())
		return
	}

	member, err := s.GuildMember(m.GuildID, m.Mentions[0].ID)
	if err != nil {
		throwable.SendErrorEmbed(s, m.ChannelID, throwable.SomethingWentWrongMember.Error())
		return
	}

	roleMap := make(map[string]*discordgo.Role)
	for _, role := range roles {
		roleMap[role.ID] = role
	}

	if !util.IsAdmin(member, roleMap) {
		throwable.SendErrorEmbed(s, m.ChannelID, throwable.WithoutSufficientPermissions.Error())
		return
	}

	embedFields := make([]*discordgo.MessageEmbedField, len(member.Roles))
	for k, roleMember := range member.Roles {
		role, exists := roleMap[roleMember]
		if exists {
			embedFields[k] = &discordgo.MessageEmbedField{
				Name:   fmt.Sprintf("Role %d", k+1),
				Value:  fmt.Sprintf("`%s`", role.Name),
				Inline: false,
			}
		}
	}

	embed := &discordgo.MessageEmbed{
		Title:       "👥 Server Roles",
		Description: "All roles available on the server",
		Fields:      embedFields,
		Color:       0x00d8ff,
	}
	s.ChannelMessageSendEmbed(m.ChannelID, embed)
}
