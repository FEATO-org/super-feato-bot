package interfaces

import (
	"strings"

	"github.com/FEATO-org/support-feato-system/domain/model"
	"github.com/FEATO-org/support-feato-system/usecase"
	"github.com/bwmarrin/discordgo"
)

type CharacterInterfaces = CommandInterface

type characterInterfaces struct {
	characterUsecase usecase.CharacterUsecase
}

// BuildCommands implements CommandInterface
func (*characterInterfaces) BuildCommands() []*discordgo.ApplicationCommand {
	return []*discordgo.ApplicationCommand{
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

// BuildHandlers implements CommandInterface
func (ci *characterInterfaces) BuildHandlers() map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"generate-character": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			var character *model.Character
			var err error

			options := i.ApplicationCommandData().Options
			optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
			for _, opt := range options {
				optionMap[opt.Name] = opt
			}
			if option, ok := optionMap["gender"]; ok {
				character, err = ci.characterUsecase.Generate(option.StringValue())
			} else {
				character, err = ci.characterUsecase.Generate("")
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

func NewCharacterInterfaces(characterUsecase usecase.CharacterUsecase) CharacterInterfaces {
	return &characterInterfaces{
		characterUsecase: characterUsecase,
	}
}
