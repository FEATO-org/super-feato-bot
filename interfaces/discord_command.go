package interfaces

import (
	"errors"
	"strings"

	"github.com/FEATO-org/support-feato-system/domain/model"
	"github.com/FEATO-org/support-feato-system/usecase"
	"github.com/bwmarrin/discordgo"
)

type DiscordCommandInterfaces interface {
	BuildCommands() []*discordgo.ApplicationCommand
	BuildHandlers() map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate)
}

type discordCommandInterfaces struct {
	discordUserCommandUsecase usecase.DiscordUserCommandUsecase
}

// BuildCommands implements DiscordCommandInterfaces.
func (di *discordCommandInterfaces) BuildCommands() []*discordgo.ApplicationCommand {
	return []*discordgo.ApplicationCommand{
		{
			Type:        discordgo.ChatApplicationCommand,
			Name:        "dice",
			Description: "dice roll",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "dice-option",
					Description: "[dice]d[men](+[dice]d[men])...",
					Type:        discordgo.ApplicationCommandOptionString,
					Required:    true,
				},
			},
		},
		{
			Type:        discordgo.ChatApplicationCommand,
			Name:        "generate-character",
			Description: "Generate character",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "gender",
					Description: "set gender",
					Required:    false,
					Type:        discordgo.ApplicationCommandOptionString,
					Choices: []*discordgo.ApplicationCommandOptionChoice{
						{
							Name:  "male",
							Value: "male",
						},
						{
							Name:  "female",
							Value: "female",
						},
					},
				},
			},
		},
	}
}

// BuildHandlers implements DiscordCommandInterfaces.
func (di *discordCommandInterfaces) BuildHandlers() map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"dice": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			var dice *model.Dice
			var err error

			options := i.ApplicationCommandData().Options
			optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
			for _, opt := range options {
				optionMap[opt.Name] = opt
			}
			if option, ok := optionMap["dice-option"]; ok {
				dice, err = di.discordUserCommandUsecase.DiceRoll(option.StringValue())
			} else {
				err = errors.New("ダイスのオプショナルが指定されませんでした")
			}
			if err != nil {
				ServerErrorInteractionRespond(err, s, i)
				return
			}

			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: dice.GetResult(),
				},
			},
			)
		},
		"generate-character": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			var character *model.Character
			var err error

			options := i.ApplicationCommandData().Options
			optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
			for _, opt := range options {
				optionMap[opt.Name] = opt
			}
			if option, ok := optionMap["gender"]; ok {
				character, err = di.discordUserCommandUsecase.CharacterGenerate(option.StringValue())
			} else {
				character, err = di.discordUserCommandUsecase.CharacterGenerate("")
			}
			if err != nil {
				ServerErrorInteractionRespond(err, s, i)
			}

			var messageBuilder strings.Builder
			messageBuilder.WriteString("名前：")
			messageBuilder.WriteString(character.GetName())
			messageBuilder.WriteString("\nふりがな：")
			messageBuilder.WriteString(character.GetNameHiragana())
			messageBuilder.WriteString("\n性別：")
			messageBuilder.WriteString(character.GetGender())

			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: messageBuilder.String(),
				},
			})
		},
	}
}

func NewDiscordCommandInterfaces(discordUserCommandUsecase usecase.DiscordUserCommandUsecase) DiscordCommandInterfaces {
	return &discordCommandInterfaces{
		discordUserCommandUsecase: discordUserCommandUsecase,
	}
}
