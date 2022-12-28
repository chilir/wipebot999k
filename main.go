package main

import (
	"fmt"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
	dotenv "github.com/joho/godotenv"
)

func main() {
	err := dotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file %w", err)
	}

	discord, err := discordgo.New("Bot " + os.Getenv("BOT_TOKEN"))
	if err != nil {
		log.Fatal("Error creating Discord session %w", err)
	}

	// Open a websocket connection to Discord and begin listening.
	err = discord.Open()
	if err != nil {
		log.Fatal("Error opening connection %w", err)
	}

	fmt.Println("Bot is running! Bot will close after this message.")

	// Cleanly close down the Discord session.
	discord.Close()
}
