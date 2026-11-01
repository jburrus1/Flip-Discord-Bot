package handlers

import (
	"discord-bot/database"
	"log"

	"github.com/bwmarrin/discordgo"
)

func ServerCreate(s *discordgo.Session, g *discordgo.GuildCreate) {
	log.Printf("Server available: %s (%s)", g.Name, g.ID)
	syncDatabase(s, g)
}

func syncDatabase(s *discordgo.Session, g *discordgo.GuildCreate) {
	database.SyncUsersDb(s, g.ID)
	for _, channel := range g.Channels {
		if channel.Type != discordgo.ChannelTypeGuildText {
			continue
		}
		database.SyncMessagesDb(s, channel.ID)
	}
}
