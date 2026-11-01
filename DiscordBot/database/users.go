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

func SyncUsersDb(dg *discordgo.Session, guildID string) {
	members, err := discord.GetMembers(dg, guildID)
	if err != nil {
		log.Println(err)
		return
	}

	for _, member := range members {
		SaveUser(member)
	}
}

func SaveUser(member *discordgo.Member) {

	if member.User == nil {
		return
	}
	log.Println("Saving user:", member.User.Username)

	body, err := json.Marshal(struct {
		ID       string `json:"id"`
		Username string `json:"username"`
	}{
		ID:       member.User.ID,
		Username: member.User.Username,
	})
	if err != nil {
		log.Println(err)
		return
	}

	resp, err := http.Post(os.Getenv("BACKEND_URL")+"/api/Users", "application/json", bytes.NewReader(body))
	if err != nil {
		log.Println(err)
		return
	}
	resp.Body.Close()
}
