package interfaces

import (
	"errors"

	"github.com/FEATO-org/support-feato-system/domain/model"
	"github.com/FEATO-org/support-feato-system/usecase"
	"github.com/bwmarrin/discordgo"
)

type DiceInterfaces = CommandInterface

type diceInterfaces struct {
	diceUsecase usecase.DiceUsecase
}

func NewDiceInterfaces(diceUsecase usecase.DiceUsecase) DiceInterfaces {
	return &diceInterfaces{
		diceUsecase: diceUsecase,
	}
}

func (di *diceInterfaces) BuildCommands() []*discordgo.ApplicationCommand {
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
	}
}

func (di *diceInterfaces) BuildHandlers() map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate) {
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
				dice, err = di.diceUsecase.Roll(option.StringValue())
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
	}
}
