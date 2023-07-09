package config

import (
	"net/http"
	"time"
)

var BASE_URL = "api.p2pquake.net"

type P2PEarthquake struct {
	WebSocketURL string
	Client       *http.Client
}

func NewP2PEarthquake() *P2PEarthquake {
	webSocketURL := "ws://" + BASE_URL + "/v2/ws"
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	return &P2PEarthquake{
		WebSocketURL: webSocketURL,
		Client:       client,
	}
}
