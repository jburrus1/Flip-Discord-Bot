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

func InitUsersDb(dg *discordgo.Session) {
	members, err := discord.GetMembers(dg, os.Getenv("GUILD_ID"))
	if err != nil {
		log.Println(err)
		return
	}

	for _, member := range members {
		if member.User == nil {
			continue
		}

		body, err := json.Marshal(struct {
			ID       string `json:"id"`
			Username string `json:"username"`
		}{
			ID:       member.User.ID,
			Username: member.User.Username,
		})
		if err != nil {
			log.Println(err)
			continue
		}

		resp, err := http.Post(os.Getenv("BACKEND_URL")+"/api/Users", "application/json", bytes.NewReader(body))
		if err != nil {
			log.Println(err)
			continue
		}
		resp.Body.Close()
	}
}
