package action

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

type Ban struct{}

func (b *Ban) Execute(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.GuildMember(m.GuildID, m.Author.ID)
	member, err := s.GuildMember(m.GuildID, m.Author.ID)
	if err != nil {
		fmt.Println("Error al obtener miembro:", err)
		return
	}

	for _, v := range member.Roles {
		role, err := s.State.Role(m.GuildID, v)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(role.Name)
	}
	err = s.GuildBanCreate(m.GuildID, m.Mentions[0].ID, 1)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
}
