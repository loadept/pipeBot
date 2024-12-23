package observer

import "github.com/bwmarrin/discordgo"

type MessageObserver interface {
	NotifyMessage(s *discordgo.Session, m *discordgo.MessageCreate)
}

type GuildObserver interface {
	NotifyGuild(s *discordgo.Session, g *discordgo.GuildMemberAdd)
}
