package interfaces

import (
	"context"
	"errors"
	"net/http"

	"golang.org/x/oauth2"
)

type Oauth2Interfaces interface {
	GetAuthUrl() string
	GetClient(token *oauth2.Token, authCode string) (*http.Client, *oauth2.Token, error)
}

type oauth2Interfaces struct {
	context context.Context
	config  *oauth2.Config
}

// GetAuthUrl implements GoogleApisInterfaces.
func (gi *oauth2Interfaces) GetAuthUrl() string {
	return gi.config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
}

// GetClient implements GoogleApisInterfaces.
func (gi *oauth2Interfaces) GetClient(token *oauth2.Token, authCode string) (*http.Client, *oauth2.Token, error) {
	if authCode == "" {
		return gi.config.Client(gi.context, token), token, nil
	} else if token == nil {
		token, err := gi.config.Exchange(context.TODO(), authCode)
		if err != nil {
			return nil, nil, err
		}
		return gi.config.Client(gi.context, token), token, nil
	} else {
		return nil, nil, errors.New("正しく引数が指定されていません")
	}
}

func NewOauth2Interfaces(context context.Context, config *oauth2.Config) Oauth2Interfaces {
	return &oauth2Interfaces{
		context: context,
		config:  config,
	}
}
