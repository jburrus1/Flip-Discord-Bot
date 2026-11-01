package handlers

import (
	"discord-bot/database"

	"github.com/bwmarrin/discordgo"
)

func UserJoin(s *discordgo.Session, g *discordgo.GuildMemberAdd) {
	database.SaveUser(g.Member)
}
