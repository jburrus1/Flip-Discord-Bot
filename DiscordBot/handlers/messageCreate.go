package handlers

import (
	"discord-bot/database"

	"github.com/bwmarrin/discordgo"
)

func MessageCreate(s *discordgo.Session, g *discordgo.MessageCreate) {
	database.SaveMessage(g.Message)
}
