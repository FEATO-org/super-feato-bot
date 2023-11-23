package model

import "errors"

type SystemUser struct {
	Id        int64
	DiscordID string
	Guilds    []Guild
}

func NewSystemUser(id int64, discordID string, guilds []Guild) (*SystemUser, error) {
	if discordID == "" {
		return nil, errors.New("DiscordIDが指定されてません")
	}
	return &SystemUser{
		Id:        id,
		DiscordID: discordID,
		Guilds:    guilds,
	}, nil
}

func (su *SystemUser) Set(id int64, discordID string, guilds []Guild) error {
	if discordID == "" {
		return errors.New("DiscordIDが指定されてません")
	}
	su.Id = id
	su.DiscordID = discordID
	su.Guilds = guilds
	return nil
}
