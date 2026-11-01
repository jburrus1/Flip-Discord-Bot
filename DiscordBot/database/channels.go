package database

import (
	"bytes"
	"discord-bot/discord"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/bwmarrin/discordgo"
)

func SyncChannelsDb(dg *discordgo.Session, serverId string) {
	channels, err := discord.GetChannels(dg, serverId)
	if err != nil {
		log.Println(err)
		return
	}
	if channels == nil {
		return
	}
	for _, channel := range channels {
		if channel.Type != discordgo.ChannelTypeGuildText {
			continue
		}
		body, err := json.Marshal(struct {
			ID       string `json:"id"`
			Name     string `json:"name"`
			ServerID string `json:"serverId"`
		}{
			ID:       channel.ID,
			Name:     channel.Name,
			ServerID: serverId,
		})
		if err != nil {
			log.Println(err)
			continue
		}
		resp, err := http.Post(os.Getenv("BACKEND_URL")+"/api/Channels", "application/json", bytes.NewReader(body))
		if err != nil {
			log.Println(err)
			continue
		}
		if resp.StatusCode < 200 || resp.StatusCode >= 300 {
			responseBody, _ := io.ReadAll(resp.Body)

			log.Printf(
				"Failed to save channel %s: %s - %s",
				channel.ID,
				resp.Status,
				string(responseBody),
			)
		}
		resp.Body.Close()
	}
}
