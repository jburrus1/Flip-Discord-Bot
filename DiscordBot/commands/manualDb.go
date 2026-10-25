package commands

import (
	"discord-bot/database"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
)

func CreateManualDbCommands(dg *discordgo.Session) {
	_, err := dg.ApplicationCommandCreate(
		dg.State.User.ID,
		os.Getenv("GUILD_ID"),
		&discordgo.ApplicationCommand{
			Name:        "reset-db",
			Description: "Reset database(s)",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "database",
					Description: "Database to reset",
					Required:    true,
					Choices: []*discordgo.ApplicationCommandOptionChoice{
						{
							Name:  "All",
							Value: "all",
						},
						{
							Name:  "Users",
							Value: "users",
						},
						{
							Name:  "Messages",
							Value: "messages",
						},
					},
				},
			},
		})
	if err != nil {
		log.Fatal(err)
	}

	_, err = dg.ApplicationCommandCreate(
		dg.State.User.ID,
		os.Getenv("GUILD_ID"),
		&discordgo.ApplicationCommand{
			Name:        "init-db",
			Description: "Initialize database(s)",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "database",
					Description: "Database to initialize",
					Required:    true,
					Choices: []*discordgo.ApplicationCommandOptionChoice{
						{
							Name:  "All",
							Value: "all",
						},
						{
							Name:  "Users",
							Value: "users",
						},
						{
							Name:  "Messages",
							Value: "messages",
						},
					},
				},
			},
		})
	if err != nil {
		log.Fatal(err)
	}
}

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
	option := i.ApplicationCommandData().GetOption("database")
	if option == nil {
		return dg.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "No database was selected.",
			},
		})
	}

	db := option.StringValue()
	var target string
	success := false
	switch db {
	case "all":
		target = "all databases"
		database.InitUsersDb(dg)
		database.InitMessagesDb(dg, i.ChannelID)
	case "users":
		target = "the users database"
		database.InitUsersDb(dg)
	case "messages":
		target = "the messages database"
		database.InitMessagesDb(dg, i.ChannelID)
		success = true
	default:
		target = "an unknown database"
	}

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
