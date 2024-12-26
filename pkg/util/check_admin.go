package util

import (
	"github.com/bwmarrin/discordgo"
)

func IsAdmin(member *discordgo.Member, roleMap map[string]*discordgo.Role) bool {
	if member.Permissions&discordgo.PermissionAdministrator != 0 {
		return true
	}
	for _, roleID := range member.Roles {
		role, ok := roleMap[roleID]
		if !ok {
			continue
		}
		if role.Permissions&discordgo.PermissionAdministrator != 0 {
			return true
		}
	}
	return false
}
