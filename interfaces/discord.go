package interfaces

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

type DiscordInterfaces interface {
	CreateApplicationCommand(s *discordgo.Session)
	AddCommandHandler(s *discordgo.Session)
	AddMessageHandler(s *discordgo.Session)
	DeleteApplicationCommand(s *discordgo.Session)
}

type discordInterfaces struct {
	diceInterfaces      DiceInterfaces
	characterInterfaces CharacterInterfaces
	systemInterfaces    SystemInterfaces
	guildIDs            []string
	commands            map[string][]*discordgo.ApplicationCommand
}

func NewDiscordInterfaces(diceInterfaces DiceInterfaces, characterInterfaces CharacterInterfaces, systemInterfaces SystemInterfaces, guildIDs []string) DiscordInterfaces {
	return &discordInterfaces{
		diceInterfaces:      diceInterfaces,
		characterInterfaces: characterInterfaces,
		systemInterfaces:    systemInterfaces,
		guildIDs:            guildIDs,
		commands:            map[string][]*discordgo.ApplicationCommand{},
	}
}

func (di *discordInterfaces) AddMessageHandler(s *discordgo.Session) {
	s.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}
		// If the message is "ping" reply with "Pong!"
		if m.Content == "ping" {
			s.ChannelMessageSend(m.ChannelID, "Pong!")
			return
		}

		// If the message is "pong" reply with "Ping!"
		if m.Content == "pong" {
			s.ChannelMessageSend(m.ChannelID, "Ping!")
			return
		}
	})

}

func (di *discordInterfaces) CreateApplicationCommand(s *discordgo.Session) {
	commands := di.diceInterfaces.BuildCommands()
	commands = append(commands, di.characterInterfaces.BuildCommands()...)
	commands = append(commands, di.systemInterfaces.BuildCommands()...)
	for _, guildID := range di.guildIDs {
		registeredCommands := make([]*discordgo.ApplicationCommand, len(commands))
		for i, v := range commands {
			cmd, err := s.ApplicationCommandCreate(s.State.User.ID, guildID, v)
			if err != nil {
				log.Panicf("Cannot create '%v' command: %v", v.Name, err)
			}
			registeredCommands[i] = cmd
		}
		di.commands[guildID] = registeredCommands
	}
	log.Println("Completed create application command.")
}

func (di *discordInterfaces) AddCommandHandler(s *discordgo.Session) {
	commandHandlers := margeCommandHandlerMap(di.diceInterfaces.BuildHandlers(), di.characterInterfaces.BuildHandlers())
	commandHandlers = margeCommandHandlerMap(commandHandlers, di.systemInterfaces.BuildHandlers())
	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})
}

func (di *discordInterfaces) DeleteApplicationCommand(s *discordgo.Session) {
	for k, v := range di.commands {
		for _, v := range v {
			err := s.ApplicationCommandDelete(s.State.User.ID, k, v.ID)
			if err != nil {
				log.Panicf("Cannot delete '%v' command: %v", v.Name, err)
			}
		}
	}
	log.Println("Completed application command delete.")
}

func margeCommandHandlerMap(baseMap, appendMap map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate)) map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	marge := map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){}
	for k, v := range baseMap {
		marge[k] = v
	}
	for k, v := range appendMap {
		marge[k] = v
	}
	return (marge)
}
