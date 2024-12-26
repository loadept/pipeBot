package util

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/loadept/pipeBot/pkg/throwable"
)

func ValidateRole(s *discordgo.Session, m *discordgo.MessageCreate) (*discordgo.Role, string, error) {
	content := strings.Split(m.Content, " ")
	if err := CheckMentions(s, m.ChannelID, content, m.Mentions); err != nil {
		return nil, "", err
	}

	roleContent := content[len(content)-1]

	roles, err := s.GuildRoles(m.GuildID)
	if err != nil {
		return nil, "", throwable.SomethingWentWrongRole
	}

	member, err := s.GuildMember(m.GuildID, m.Author.ID)
	if err != nil {
		return nil, "", throwable.SomethingWentWrongMember
	}

	roleMapByName := make(map[string]*discordgo.Role)
	roleMapByID := make(map[string]*discordgo.Role)
	for _, role := range roles {
		roleMapByName[role.Name] = role
		roleMapByID[role.ID] = role
	}

	if !IsAdmin(member, roleMapByID) {
		return nil, "", throwable.WithoutSufficientPermissions
	}

	targetRole, exists := roleMapByName[roleContent]
	if !exists {
		return nil, "", throwable.RoleDoesNotExists
	}

	return targetRole, roleContent, nil
}
