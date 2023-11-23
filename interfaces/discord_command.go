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
	BuildCommandHandlers() map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate)
	BuildMessageComponentHandlers() map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate)
}

type discordCommandInterfaces struct {
	discordUserCommandUsecase usecase.DiscordUserCommandUsecase
}

// BuildMessageHandlers implements DiscordCommandInterfaces.
func (di *discordCommandInterfaces) BuildMessageComponentHandlers() map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){}
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
		{
			Type:        discordgo.ChatApplicationCommand,
			Name:        "create-system-user",
			Description: "管理者ユーザーを登録する",
		},
	}
}

// BuildHandlers implements DiscordCommandInterfaces.
func (di *discordCommandInterfaces) BuildCommandHandlers() map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate) {
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

			err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: dice.GetResult(),
				},
			})
			if err != nil {
				ServerErrorInteractionRespond(err, s, i)
			}
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

			err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: messageBuilder.String(),
				},
			})
			if err != nil {
				ServerErrorInteractionRespond(err, s, i)
			}
		},
		"create-system-user": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			guild, err := s.State.Guild(i.GuildID)
			if err != nil {
				ServerErrorInteractionRespond(err, s, i)
				return
			}
			_, err = di.discordUserCommandUsecase.CreateSystemUser(getUserID(i), guild.ID, "", guild.Name)
			if err != nil {
				ServerErrorInteractionRespond(err, s, i)
				return
			}
			err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "登録完了",
					Flags:   discordgo.MessageFlagsEphemeral,
				},
			})
			if err != nil {
				ServerErrorInteractionRespond(err, s, i)
			}
		},
	}
}

func NewDiscordCommandInterfaces(discordUserCommandUsecase usecase.DiscordUserCommandUsecase) DiscordCommandInterfaces {
	return &discordCommandInterfaces{
		discordUserCommandUsecase: discordUserCommandUsecase,
	}
}

func getUserID(i *discordgo.InteractionCreate) string {
	return i.Member.User.ID
}

func getGuildID(i *discordgo.InteractionCreate) string {
	return i.GuildID
}
