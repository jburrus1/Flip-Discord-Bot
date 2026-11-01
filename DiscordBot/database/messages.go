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

func SyncMessagesDb(dg *discordgo.Session, serverId string) {
	channels, err := discord.GetChannels(dg, serverId)

	recent := GetMostRecentMessage(dg, serverId)
	afterId := ""
	if recent != nil {
		afterId = recent.ID
	}

	for _, channel := range channels {
		messages, err := GetMessagesAfter(dg, channel.ID, afterId)
		if err != nil {
			continue
		}
		if messages == nil {
			continue
		}
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
			if resp.StatusCode < 200 || resp.StatusCode >= 300 {
				responseBody, _ := io.ReadAll(resp.Body)

				log.Printf(
					"Failed to save message %s: %s - %s",
					message.ID,
					resp.Status,
					string(responseBody),
				)
			}
			resp.Body.Close()
		}
	}
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(os.Getenv("BACKEND_URL") + "/api/Messages")
}

func GetMostRecentMessage(dg *discordgo.Session, serverID string) *discordgo.Message {
	resp, err := http.Get(os.Getenv("BACKEND_URL") + "/api/messages/recent/" + serverID)
	if err != nil {
		log.Println(err)
		return nil
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		responseBody, _ := io.ReadAll(resp.Body)
		log.Printf(
			"Failed to get most recent message for server %s: %s - %s",
			serverID,
			resp.Status,
			string(responseBody),
		)
		return nil
	}

	var message discordgo.Message
	if err := json.NewDecoder(resp.Body).Decode(&message); err != nil {
		log.Println(err)
		return nil
	}

	return &message
}

func GetMessagesAfter(s *discordgo.Session, channelID string, afterID string) ([]*discordgo.Message, error) {
	var allMessages []*discordgo.Message

	for {
		messages, err := s.ChannelMessages(
			channelID, 100, "", afterID, "",
		)
		if err != nil {
			return nil, err
		}

		if len(messages) == 0 {
			break
		}

		allMessages = append(allMessages, messages...)

		// Discord returns messages newest first.
		// Advance to the newest message in this batch.
		afterID = messages[0].ID

		if len(messages) < 100 {
			break
		}
	}

	return allMessages, nil
}
