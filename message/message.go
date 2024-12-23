package message

import (
	"github.com/bwmarrin/discordgo"
	"github.com/loadept/pipeBot/observer"
)

type MessageHandler struct {
	handlers []any
}

func (mh *MessageHandler) SubscribeObserver(observer any) {
	mh.handlers = append(mh.handlers, observer)
}

func (mh *MessageHandler) NotifyObservers(s *discordgo.Session, event any) {
	for _, v := range mh.handlers {
		switch e := event.(type) {
		case *discordgo.MessageCreate:
			if msgObserver, ok := v.(observer.MessageObserver); ok {
				msgObserver.NotifyMessage(s, e)
			}
		case *discordgo.GuildMemberAdd:
			if guildObserver, ok := v.(observer.GuildObserver); ok {
				guildObserver.NotifyGuild(s, e)
			}
		}
	}
}
