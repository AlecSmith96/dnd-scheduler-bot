package usecases

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
    // This isn't required in this specific example but it's a good practice.
    if m.Author.ID == s.State.User.ID {
        return
    }

	if m.Content == "!hello" {
		// Send a text message with the list of Gophers
		_, err := s.ChannelMessageSend(m.ChannelID, "hello to you")
		if err != nil {
			log.Printf("unable to send message to channel: %v", err)
		}
	}
}