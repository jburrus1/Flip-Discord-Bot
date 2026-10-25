package discord

import (
	"fmt"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
)

func GetMembersTest(dg *discordgo.Session) {
	members, err := GetMembers(dg, os.Getenv("GUILD_ID"))
	if err != nil {
		log.Fatal(err)
	}

	for _, member := range members {
		fmt.Printf("ID: %s, Username: %s, Nickname: %s\n",
			member.User.ID,
			member.User.Username,
			member.Nick,
		)
	}
}

func GetChannelsTest(dg *discordgo.Session) {
	channels, err := GetChannels(dg, os.Getenv("GUILD_ID"))
	if err != nil {
		log.Fatal(err)
	}

	for _, channel := range channels {
		fmt.Printf("ID: %s, Name: %s\n",
			channel.ID,
			channel.Name,
		)
	}
}

func GetMessagesTest(dg *discordgo.Session) {
	channels, err := GetChannels(dg, os.Getenv("GUILD_ID"))
	if err != nil {
		log.Fatal(err)
	}
	id := channels[0].ID
	messages, err := GetMessages(dg, id)
	if err != nil {
		log.Fatal(err)
	}
	for _, message := range messages {
		fmt.Printf("ID: %s, Content: %s\n",
			message.ID,
			message.Content,
		)
	}
}
