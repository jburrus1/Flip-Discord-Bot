package database

import (
	"bytes"
	"discord-bot/discord"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/bwmarrin/discordgo"
)

func InitMessagesDb(dg *discordgo.Session, channelID string) {
	messages, err := discord.GetMessages(dg, channelID)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(os.Getenv("BACKEND_URL") + "/api/Messages")

	for _, message := range messages {
		if message.Author == nil {
			continue
		}

		body, err := json.Marshal(struct {
			ID        string `json:"id"`
			AuthorID  string `json:"authorId"`
			ChannelID string `json:"channelId"`
			Content   string `json:"content"`
		}{
			ID:        message.ID,
			AuthorID:  message.Author.ID,
			ChannelID: message.ChannelID,
			Content:   message.Content,
		})
		if err != nil {
			log.Println(err)
			continue
		}
		resp, err := http.Post(os.Getenv("BACKEND_URL")+"/api/Messages", "application/json", bytes.NewReader(body))
		if err != nil {
			log.Println(err)
			continue
		}
		resp.Body.Close()
	}
}
