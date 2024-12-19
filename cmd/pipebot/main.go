package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"github.com/loadept/pipeBot/bot"
)

func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Printf(".env file not found, using system variables: %v\n", err)
	}
}

func main() {
	token := os.Getenv("TOKEN")
	if len(token) == 0 {
		fmt.Println("Error: Please set your TOKEN environment variable")
		os.Exit(1)
	}

	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Printf("Error at: %v", err)
	}

	dg.AddHandler(bot.MessageCreate)

	err = dg.Open()
	if err != nil {
		fmt.Printf("Error to open connection with discord: %v\n", err)
		return
	}

	dg.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	fmt.Println("Bot is running. Press Ctrl+C to exit.")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-stop

	dg.Close()
}
