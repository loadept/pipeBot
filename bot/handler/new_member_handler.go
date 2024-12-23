package handler

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

type NewMember struct{}

func (nm *NewMember) NotifyGuild(s *discordgo.Session, g *discordgo.GuildMemberAdd) {
	roles, _ := s.GuildRoles(g.GuildID)
	fmt.Println(roles)
}
