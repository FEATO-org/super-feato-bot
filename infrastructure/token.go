package infrastructure

import (
	"context"
	"database/sql"

	"github.com/FEATO-org/support-feato-system/db/sqlc"
	"github.com/FEATO-org/support-feato-system/domain/model"
	"github.com/FEATO-org/support-feato-system/domain/repository"
)

type TokenRepository struct {
	db  *sql.DB
	ctx context.Context
}

// Create implements repository.TokenRepository.
func (tr *TokenRepository) Create(token *model.Token) (*model.Token, error) {
	queries := sqlc.New(tr.db)

	isSystemUser, isGuild := true, true
	if token.SystemUser.DiscordID == "" {
		systemUser, err := queries.FindByDiscordIDSystemUser(tr.ctx, token.SystemUser.DiscordID)
		if err != nil {
			return nil, err
		}
		isSystemUser = false
		token.Set(token.Id, model.SystemUser{
			Id:        systemUser.ID,
			DiscordID: systemUser.DiscordID,
			Guilds:    []model.Guild{},
		}, token.Guild, token.AccessToken, token.TokenType, token.RefreshToken, token.Expiry)
	}

	if token.Guild.DiscordID == "" {
		guild, err := queries.FindByDiscordIDGuild(tr.ctx, token.Guild.DiscordID)
		if err != nil {
			return nil, err
		}
		isGuild = false
		token.Set(token.Id, token.SystemUser, model.Guild{
			Id:        guild.ID,
			Name:      guild.Name,
			DiscordID: guild.DiscordID,
			SheetID:   guild.SheetID.String,
		}, token.AccessToken, token.TokenType, token.RefreshToken, token.Expiry)
	}

	result, err := queries.CreateToken(tr.ctx, sqlc.CreateTokenParams{
		SystemUserID: sql.NullInt64{
			Int64: token.SystemUser.Id,
			Valid: isSystemUser,
		},
		GuildID: sql.NullInt64{
			Int64: token.Guild.Id,
			Valid: isGuild,
		},
		AccessToken:  token.AccessToken,
		TokenType:    token.TokenType,
		RefreshToken: token.RefreshToken,
		Expiry:       token.Expiry,
	})
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	err = token.Set(id, token.SystemUser, token.Guild, token.AccessToken, token.TokenType, token.RefreshToken, token.Expiry)
	if err != nil {
		return nil, err
	}

	return token, nil
}

// Delete implements repository.TokenRepository.
func (*TokenRepository) Delete(token *model.Token) error {
	panic("unimplemented")
}

// FindByID implements repository.TokenRepository.
func (*TokenRepository) FindByID(id int64) (*model.Token, error) {
	panic("unimplemented")
}

// Update implements repository.TokenRepository.
func (*TokenRepository) Update(token *model.Token) (*model.Token, error) {
	panic("unimplemented")
}

func NewTokenRepository(db *sql.DB, ctx context.Context) repository.TokenRepository {
	return &TokenRepository{
		db:  db,
		ctx: ctx,
	}
}
