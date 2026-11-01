package database

import (
	"discord-bot/discord"
	"encoding/json"

	"github.com/bwmarrin/discordgo"
)

func SyncChannelsDb(dg *discordgo.Session, serverId string) {
	channels, err := discord.GetChannels(dg, serverId)
	if err != nil {
		return
	}
	if channels == nil {
		return
	}
	for _, channel := range channels {
		body, err := json.Marshal(struct {
			ID       string `json:"id"`
			Name     string `json:"name"`
			ServerID string `json:"serverId"`
		}{
			ID:       channel.ID,
			Name:     channel.Name,
			ServerID: serverId,
		})
	}
}
