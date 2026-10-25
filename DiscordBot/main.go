package main

import (
	"discord-bot/commands"
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
)

func main() {
	token := os.Getenv("DISCORD_TOKEN")
	if token == "" {
		log.Fatal("DISCORD_TOKEN is not set")
	}

	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal(err)
	}

	dg.Identify.Intents = discordgo.IntentsGuildMessages |
		discordgo.IntentsMessageContent |
		discordgo.IntentsGuildMembers

	if err := dg.Open(); err != nil {
		log.Fatal(err)
	}
	defer dg.Close()

	addHandlers(dg)
	addCommands(dg)

	log.Println("Bot is running!")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop
}

func addHandlers(dg *discordgo.Session) {
	dg.AddHandler(messageCreate)
	dg.AddHandler(interactionCreate)
}

func addCommands(dg *discordgo.Session) {
	commands.CreateTestingCommands(dg)
	commands.CreateManualDbCommands(dg)
}

func interactionCreate(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Type != discordgo.InteractionApplicationCommand {
		return
	}

	command := i.ApplicationCommandData()

	log.Printf("Handling command: %v\n", command.Name)
	var err error
	switch command.Name {
	case "ping":
		err = commands.Ping(s, i)
	case "reset-db":
		err = commands.ResetDb(s, i)
	case "init-db":
		err = commands.InitDb(s, i)
	}

	if err != nil {
		log.Println(err)
	}
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}
}
