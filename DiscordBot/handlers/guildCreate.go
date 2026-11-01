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
	database.SyncServersDb(s, g.ID)
	database.SyncChannelsDb(s, g.ID)
	database.SyncUsersDb(s, g.ID)
	database.SyncMessagesDb(s, g.ID)
}
