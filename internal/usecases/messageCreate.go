package usecases

import (
	"log"
	"strings"

	"github.com/AlecSmith96/dnd-scheduler-bot/internal/entities"
	"github.com/bwmarrin/discordgo"
)

// MessageCreate is invoked every time a user sends a message to a channel that the bot is in
func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
    // This isn't required in this specific example but it's a good practice.
    if m.Author.ID == s.State.User.ID {
        return
    }

	messageInputs := strings.Split(m.Content, " ")
	switch messageInputs[0] {
	case "!hello":
		_, err := s.ChannelMessageSend(m.ChannelID, "hello to you")
		if err != nil {
			log.Printf("unable to send message to channel: %v", err)
		}
	case "!schedule":
		_, err := s.ChannelMessageSend(m.ChannelID, "hello to you")
		if err != nil {
			log.Printf("unable to send message to channel: %v", err)
		}
	case "!help":
		_, err := s.ChannelMessageSend(m.ChannelID, entities.HelpMessage)
		if err != nil {
			log.Printf("unable to send message to channel: %v", err)
		}
	}
}