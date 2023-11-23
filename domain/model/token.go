package model

import "time"

type Token struct {
	Id           int64
	SystemUser   SystemUser
	Guild        Guild
	AccessToken  string
	TokenType    string
	RefreshToken string
	Expiry       time.Time
}

func NewToken(id int64, systemUser SystemUser, guild Guild, accessToken, tokenType, refreshToken string, expiry time.Time) (*Token, error) {
	return &Token{
		Id:           id,
		SystemUser:   systemUser,
		Guild:        guild,
		AccessToken:  accessToken,
		TokenType:    tokenType,
		RefreshToken: refreshToken,
		Expiry:       expiry,
	}, nil
}

func (t *Token) Set(id int64, systemUser SystemUser, guild Guild, accessToken, tokenType, refreshToken string, expiry time.Time) error {
	t.Id = id
	t.SystemUser = systemUser
	t.Guild = guild
	t.AccessToken = accessToken
	t.TokenType = tokenType
	t.RefreshToken = refreshToken
	t.Expiry = expiry
	return nil
}
