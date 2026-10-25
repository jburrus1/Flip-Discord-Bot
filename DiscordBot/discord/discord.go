package discord

import "github.com/bwmarrin/discordgo"

func GetMembers(s *discordgo.Session, guildID string) ([]*discordgo.Member, error) {
	return s.GuildMembers(guildID, "", 1000)
}

func GetChannels(s *discordgo.Session, guildID string) ([]*discordgo.Channel, error) {
	channels, err := s.GuildChannels(guildID)
	if err != nil {
		return nil, err
	}
	var textChannels []*discordgo.Channel
	for _, channel := range channels {
		if channel.Type == discordgo.ChannelTypeGuildText {
			textChannels = append(textChannels, channel)
		}
	}
	return textChannels, nil
}

func GetMessages(s *discordgo.Session, channelID string) ([]*discordgo.Message, error) {
	return s.ChannelMessages(channelID, 100, "", "", "")
}
