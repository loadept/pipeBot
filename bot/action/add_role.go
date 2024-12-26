package action

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/loadept/pipeBot/pkg/throwable"
	"github.com/loadept/pipeBot/pkg/util"
)

type AddRole struct{}

func (ar *AddRole) Execute(s *discordgo.Session, m *discordgo.MessageCreate) {
	content := strings.Split(m.Content, " ")
	if err := util.CheckMentions(s, m.ChannelID, content, m.Mentions); err != nil {
		throwable.SendErrorEmbed(s, m.ChannelID, err.Error())
	}

	roleContent := content[len(content)-1]

	roles, err := s.GuildRoles(m.GuildID)
	if err != nil {
		fmt.Println("Error al obtener miembro:", err)
		return
	}
	member, err := s.GuildMember(m.GuildID, m.Author.ID)
	if err != nil {
		fmt.Println("Error al obtener miembro:", err)
		return
	}

	roleMapByName := make(map[string]*discordgo.Role)
	roleMapByID := make(map[string]*discordgo.Role)
	for _, role := range roles {
		roleMapByName[role.Name] = role
		roleMapByID[role.ID] = role
	}

	if !util.IsAdmin(member, roleMapByID) {
		throwable.SendErrorEmbed(s, m.ChannelID, "You do not have permission to perform this action.")
		return
	}

	targetRole, exists := roleMapByName[roleContent]
	if !exists {
		throwable.SendErrorEmbed(s, m.ChannelID, "The role you tried to assign does not exist.")
		return
	}

	if err := s.GuildMemberRoleAdd(m.GuildID, m.Mentions[0].ID, targetRole.ID); err != nil {
		throwable.SendErrorEmbed(s, m.ChannelID, "Failed to establish role.")
		return
	}

	embed := &discordgo.MessageEmbed{
		Title:       "🔵 Member add to role",
		Description: fmt.Sprintf("Member %s has been assigned to the role %s", m.Mentions[0].Username, roleContent),
		Color:       0x00d8ff,
	}
	s.ChannelMessageSendEmbed(m.ChannelID, embed)
}
