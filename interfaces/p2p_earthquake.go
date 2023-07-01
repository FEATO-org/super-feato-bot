package interfaces

import (
	"context"
	"fmt"
	"log"

	"github.com/FEATO-org/support-feato-system/config"
	"github.com/FEATO-org/support-feato-system/usecase"
	"github.com/bwmarrin/discordgo"
	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

type P2PEarthquakeInterfaces interface {
	ReceiveEEWToDiscord(s *discordgo.Session) error
}

type p2pEarthquakeInterfaces struct {
	systemWSIncomingUsecase usecase.SystemWSIncomingUsecase
	p2pEarthquakeConfig     config.P2PEarthquake
	context                 context.Context
	cancel                  context.CancelFunc
	discordConfig           config.DiscordConfig
}

// ReceiveEEW implements P2PEqInterfaces.
func (pi *p2pEarthquakeInterfaces) ReceiveEEWToDiscord(s *discordgo.Session) error {
	defer pi.cancel()

	conn, _, err := websocket.Dial(pi.context, pi.p2pEarthquakeConfig.WebSocketURL, &websocket.DialOptions{
		HTTPClient: pi.p2pEarthquakeConfig.Client,
	})
	if err != nil {
		log.Fatal("WebSocket connection error:", err)
	}
	defer conn.Close(websocket.StatusInternalError, "Connection closed")

	// メッセージの受信ループ
	for {
		var message interface{}
		err := wsjson.Read(pi.context, conn, &message)
		if err != nil {
			log.Fatal("Error reading message:", err)
		}

		// メッセージの処理
		fmt.Println("Received message:", message)
		pi.systemWSIncomingUsecase.ReceiveEEW(message, false)
		_, err = s.ChannelMessageSend(pi.discordConfig.NotifyChannelID, interfaceToString(message))
		if err != nil {
			log.Fatal("Error sending message:", err)
		}
	}
}

func NewP2PEarthquakeInterfaces(systemWSIncomingUsecase usecase.SystemWSIncomingUsecase, p2pEarthquakeConfig config.P2PEarthquake, context context.Context, cancel context.CancelFunc, discordConfig config.DiscordConfig) P2PEarthquakeInterfaces {
	return &p2pEarthquakeInterfaces{
		systemWSIncomingUsecase: systemWSIncomingUsecase,
		p2pEarthquakeConfig:     p2pEarthquakeConfig,
		context:                 context,
		cancel:                  cancel,
		discordConfig:           discordConfig,
	}
}
