package config

import "os"

var NOTIFY_CHANNEL_ID string

type DiscordConfig struct {
	NotifyChannelID string
}

func init() {
	NOTIFY_CHANNEL_ID = os.Getenv("NOTIFY_CHANNEL_ID")
}

func NewDiscordConfig() *DiscordConfig {
	return &DiscordConfig{
		NotifyChannelID: NOTIFY_CHANNEL_ID,
	}
}
