package util

import (
	"github.com/bwmarrin/discordgo"
	"github.com/loadept/pipeBot/pkg/throwable"
)

func ValidateRole(s *discordgo.Session, m *discordgo.MessageCreate) (*discordgo.Role, error) {
	if err := CheckMentionsRole(s, m.ChannelID, m.MentionRoles, m.Mentions); err != nil {
		return nil, err
	}

	roleContent := m.MentionRoles[0]

	roles, err := s.GuildRoles(m.GuildID)
	if err != nil {
		return nil, throwable.SomethingWentWrongRole
	}

	member, err := s.GuildMember(m.GuildID, m.Author.ID)
	if err != nil {
		return nil, throwable.SomethingWentWrongMember
	}

	roleMapByName := make(map[string]*discordgo.Role)
	roleMapByID := make(map[string]*discordgo.Role)
	for _, role := range roles {
		roleMapByName[role.Name] = role
		roleMapByID[role.ID] = role
	}

	if !IsAdmin(member, roleMapByID) {
		return nil, throwable.WithoutSufficientPermissions
	}

	targetRole, exists := roleMapByID[roleContent]
	if !exists {
		return nil, throwable.RoleDoesNotExists
	}

	return targetRole, nil
}
