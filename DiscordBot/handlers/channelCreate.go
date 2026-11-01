package handlers

import (
	"discord-bot/database"

	"github.com/bwmarrin/discordgo"
)

func ChannelCreate(s *discordgo.Session, g *discordgo.ChannelCreate) {
	database.SaveChannel(g.Channel, g.GuildID)
}
