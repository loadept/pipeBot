package action

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type AddRole struct{}

func (ar *AddRole) Execute(s *discordgo.Session, m *discordgo.MessageCreate) {
	content := strings.Split(m.Content, " ")
	roleContent := content[len(content)-1]

	roles, err := s.GuildRoles(m.GuildID)
	if err != nil {
		fmt.Println("Error al obtener miembro:", err)
		return
	}

	var roleID string
	for _, role := range roles {
		if strings.EqualFold(roleContent, role.Name) {
			roleID = role.ID
			break
		}
	}

	if err := s.GuildMemberRoleAdd(m.GuildID, m.Mentions[0].ID, roleID); err != nil {
		fmt.Println(err)
		embed := &discordgo.MessageEmbed{
			Title:       "🔴 Invalid action",
			Description: "Not recognized action",
			Color:       0xff0000,
		}
		s.ChannelMessageSendEmbed(m.ChannelID, embed)
		return
	}

	embed := &discordgo.MessageEmbed{
		Title:       "🔵 Member add to role",
		Description: fmt.Sprintf("Member %s has been assigned to the role %s", m.Mentions[0].Username, roleContent),
		Color:       0x00d8ff,
	}
	s.ChannelMessageSendEmbed(m.ChannelID, embed)
}