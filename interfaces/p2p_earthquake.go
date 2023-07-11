package interfaces

import (
	"context"
	"fmt"
	"log"
	"time"

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

type eew struct {
	Id         string `json:"id"`
	Code       int32  `json:"code"`
	Time       string `json:"time"`
	Test       bool   `json:"test"`
	Earthquake struct {
		OriginTime  string `json:"originTime"`
		ArrivalTime string `json:"arrivalTime"`
		Condition   string `json:"condition"`
		Hypocenter  struct {
			Name       string `json:"name"`
			ReduceName string `json:"reduceName"`
			Latitude   int32  `json:"latitude"`
			Longitude  int32  `json:"longitude"`
			Depth      int32  `json:"depth"`
			Magnitude  int32  `json:"magnitude"`
		}
	}
	Issue struct {
		Time    string `json:"time"`
		EventId string `json:"eventId"`
		Serial  string `json:"serial"`
	}
	Cancelled bool `json:"cancelled"`
	Area      struct {
		Pref        string `json:"pref"`
		Name        string `json:"name"`
		ScaleFrom   int32  `json:"scaleFrom"`
		ScaleTo     int32  `json:"scaleTo"`
		KindCode    string `json:"kindCode"`
		ArrivalTime string `json:"arrivalTime"`
	}
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
		var message eew
		err := wsjson.Read(pi.context, conn, &message)
		if err != nil {
			log.Fatal("Error reading message:", err)
		}

		// メッセージの処理
		fmt.Println("Received message:", message)
		// Areapeers Userquake UserquakeEvaluationを無視する
		if message.Code == 555 || message.Code == 561 || message.Code == 9611 {
			return nil
		}
		if message.Test {
			pi.systemWSIncomingUsecase.ReceiveEEW(message, true)
			return nil
		}
		pi.systemWSIncomingUsecase.ReceiveEEW(message, false)
		jst := time.FixedZone("Asia/Tokyo", 9*60*60)
		t, err := time.ParseInLocation("2006/01/02 15:04:05.000", message.Time, jst)
		if err != nil {
			return err
		}
		title := "地震情報を受信しました（実験版）"
		if message.Cancelled {
			title = "地震情報取消報を受信しました（実験版）"
		}
		_, err = s.ChannelMessageSendEmbed(pi.discordConfig.NotifyChannelID, &discordgo.MessageEmbed{
			Title:       title,
			Description: interfaceToString(message),
			Timestamp:   t.UTC().Format(time.RFC3339),
			Color:       0xffff00,
			Provider: &discordgo.MessageEmbedProvider{
				URL:  "https://www.p2pquake.net/app/web/",
				Name: "P2P 地震情報",
			},
			Author: &discordgo.MessageEmbedAuthor{
				URL:     "https://www.p2pquake.net/app/web/",
				Name:    "P2P 地震情報",
				IconURL: "https://www.p2pquake.net/images/favicon.png",
			},
		})
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
