package commands

import (
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
)

func CreateTestingCommands(dg *discordgo.Session) {
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

func Ping(dg *discordgo.Session, i *discordgo.InteractionCreate) error {

	err := dg.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Pong!",
		},
	})
	return err
}
