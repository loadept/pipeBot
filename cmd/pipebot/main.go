package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"github.com/loadept/pipeBot/bot/handler"
	"github.com/loadept/pipeBot/internal/message"
)

var (
	token string
)

func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Printf(".env file not found, using system variables: %v\n", err)
	}
	token = os.Getenv("TOKEN")
	if len(token) == 0 {
		fmt.Println("Error: Please set your TOKEN environment variable")
		os.Exit(1)
	}
}

func main() {
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Printf("Error at: %v", err)
	}

	messageHandler := &message.MessageHandler{}
	messageHandler.SubscribeObserver(&handler.MusicChannel{})
	messageHandler.SubscribeObserver(&handler.WallpaperChannel{})
	messageHandler.SubscribeObserver(&handler.Commands{})
	messageHandler.SubscribeObserver(&handler.NewMember{})

	dg.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		messageHandler.NotifyObservers(s, m)
	})
	dg.AddHandler(func(s *discordgo.Session, g *discordgo.GuildMemberAdd) {
		messageHandler.NotifyObservers(s, g)
	})

	// dg.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	go handler.AsyncEvent(dg)

	err = dg.Open()
	if err != nil {
		fmt.Printf("Error to open connection with discord: %v\n", err)
		return
	}
	defer dg.Close()

	fmt.Println("Bot is running. Press Ctrl+C to exit.")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-stop
}
