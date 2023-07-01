package interfaces

import (
	"fmt"
	"runtime"

	"github.com/bwmarrin/discordgo"
)

// エラーをlogに流した上でDiscordに返答する
func ServerErrorInteractionRespond(err error, s *discordgo.Session, i *discordgo.InteractionCreate) {
	printStackTrace()
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: err.Error(),
		},
	})
}

func printStackTrace() {
	i := 0
	for {
		_, file, line, ok := runtime.Caller(i)
		if !ok {
			// 取得できなくなったら終了
			break
		}
		fmt.Printf("%s:%d, \n", file, line)
		i += 1
	}
}

func interfaceToString(source interface{}) string {
	return fmt.Sprintf("%v", source)
}
