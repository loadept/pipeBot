package handler

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

type NewMember struct{}

func (nm *NewMember) NotifyGuild(s *discordgo.Session, g *discordgo.GuildMemberAdd) {
	roles, _ := s.GuildRoles("1092232803866914926")
	fmt.Println(roles)
}
