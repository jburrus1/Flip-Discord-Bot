package commands

import (
	"discord-bot/database"

	"github.com/bwmarrin/discordgo"
)

func ResetDb(dg *discordgo.Session, i *discordgo.InteractionCreate) error {
	option := i.ApplicationCommandData().GetOption("database")
	if option == nil {
		return dg.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "No database was selected.",
			},
		})
	}

	database := option.StringValue()
	var target string
	switch database {
	case "all":
		target = "all databases"
	case "users":
		target = "the users database"
	case "messages":
		target = "the messages database"
	default:
		target = "an unknown database"
	}

	return dg.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Reset requested for " + target + ", but database reset is not implemented yet.",
		},
	})
}

func InitDb(dg *discordgo.Session, i *discordgo.InteractionCreate) error {
	var target string
	success := false
	//database.SyncUsersDb(dg)
	database.SyncMessagesDb(dg, i.ChannelID)

	if success {
		return dg.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Initialization requested for " + target + " completed successfully.",
			},
		})
	}
	return dg.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Initialization requested for " + target + ", but database initialization is not implemented yet.",
		},
	})
}
