package database

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/bwmarrin/discordgo"
)

func SyncServersDb(dg *discordgo.Session, serverId string) {
	server, err := dg.Guild(serverId)
	if err != nil {
		log.Println(err)
		return
	}

	body, err := json.Marshal(struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}{
		ID:   server.ID,
		Name: server.Name,
	})
	if err != nil {
		log.Println(err)
		return
	}

	resp, err := http.Post(os.Getenv("BACKEND_URL")+"/api/Servers", "application/json", bytes.NewReader(body))
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		responseBody, _ := io.ReadAll(resp.Body)
		log.Printf("Failed to save server %s: %s - %s", server.ID, resp.Status, string(responseBody))
	}
}
