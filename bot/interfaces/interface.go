package interfaces

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

type CommandInterface interface {
	BuildCommands() []*discordgo.ApplicationCommand
	BuildHandlers() map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate)
}

// エラーをlogに流した上でDiscordに返答する
func ServerErrorInteractionRespond(err error, s *discordgo.Session, i *discordgo.InteractionCreate) {
	log.Println(err)
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: err.Error(),
		},
	})
}
