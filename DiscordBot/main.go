package main

import (
	"discord-bot/discord"
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

	discord.GetMessagesTest(dg)

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
	log.Println("Registering /ping")
	_, err := dg.ApplicationCommandCreate(
		dg.State.User.ID,
		os.Getenv("GUILD_ID"),
		&discordgo.ApplicationCommand{
			Name:        "ping",
			Description: "Check if bot is responding",
		})

	if err != nil {
		log.Fatal(err)
	}
}

func interactionCreate(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Type != discordgo.InteractionApplicationCommand {
		return
	}

	command := i.ApplicationCommandData()

	log.Printf("Handling command: %v\n", command.Name)
	switch command.Name {
	case "ping":
		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Pong!",
			},
		})

		if err != nil {
			log.Println(err)
		}
	}
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}
}
