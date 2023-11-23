package model

import "errors"

type Guild struct {
	Id        int64
	Name      string
	DiscordID string
	SheetID   string
}

func NewGuild(id int64, name, discordID, sheetID string) (*Guild, error) {
	if discordID == "" {
		return nil, errors.New("DiscordIdが指定されてません")
	}
	return &Guild{
		Id:        id,
		Name:      name,
		DiscordID: discordID,
		SheetID:   sheetID,
	}, nil
}

func (g *Guild) Set(id int64, name, discordID, sheetID string) error {
	if discordID == "" {
		return errors.New("DiscordIDが指定されてません")
	}
	g.Id = id
	g.Name = name
	g.DiscordID = discordID
	g.SheetID = sheetID
	return nil
}
