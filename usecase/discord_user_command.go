package usecase

import (
	"errors"
	"strings"

	"github.com/FEATO-org/support-feato-system/domain/model"
	"github.com/FEATO-org/support-feato-system/domain/repository"
	"golang.org/x/text/width"
)

type DiscordUserCommandUsecase interface {
	DiceRoll(query string) (*model.Dice, error)
	CharacterGenerate(gender string) (*model.Character, error)
	CreateSystemUser(userDiscordID, guildDiscordID, sheetId, guildName string) (*model.SystemUser, error)
}

type discordUserCommandUsecase struct {
	diceRepository       repository.DiceRepository
	characterRepository  repository.CharacterRepository
	systemUserRepository repository.SystemUserRepository
}

// CreateSystemUser implements DiscordUserCommandUsecase.
func (du *discordUserCommandUsecase) CreateSystemUser(userDiscordID, guildDiscordID, sheetId, guildName string) (*model.SystemUser, error) {
	guildModel, err := model.NewGuild(0, guildName, guildDiscordID, sheetId)
	if err != nil {
		return nil, err
	}
	model, err := model.NewSystemUser(0, userDiscordID, []model.Guild{*guildModel})
	if err != nil {
		return nil, err
	}

	return du.systemUserRepository.Create(model)
}

func (du discordUserCommandUsecase) DiceRoll(query string) (*model.Dice, error) {
	// 正規化
	normalizeQuery := strings.ToLower(query)
	normalizeQuery = width.Narrow.String(normalizeQuery)
	// バリデーション
	slice := strings.Split(normalizeQuery, "")
	len := len(slice)
	for i := 0; i < len; i++ {
		if !strings.ContainsAny(slice[i], "0123456789d+") {
			return nil, errors.New("想定外の文字がダイスに指定されました")
		}
	}
	diceModel, err := model.NewDice(normalizeQuery, "", 0)
	if err != nil {
		return nil, err
	}
	dice, err := du.diceRepository.Roll(diceModel)
	if err != nil {
		return nil, err
	}
	return dice, nil
}

func (du *discordUserCommandUsecase) CharacterGenerate(gender string) (*model.Character, error) {
	character, err := model.NewCharacter(gender, "", "", "", "")
	if err != nil {
		return nil, err
	}
	return du.characterRepository.Create(character)
}

func NewDiscordUserCommand(diceRepository repository.DiceRepository, characterRepository repository.CharacterRepository, systemUserRepository repository.SystemUserRepository) DiscordUserCommandUsecase {
	return &discordUserCommandUsecase{
		diceRepository:       diceRepository,
		characterRepository:  characterRepository,
		systemUserRepository: systemUserRepository,
	}
}
