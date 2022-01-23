package omake

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/mattn/go-gimei"
)

// msg 正規化されたメッセージ(string)を期待する
func JudgeRandomName(message string) bool {
	return strings.Contains(message, "sfs rn")
}

func ResponseRandomName(message string, session *discordgo.Session, event *discordgo.MessageCreate) {
	messageArray := strings.Split(message, " ")
	var name *gimei.Name
	var response string

	// 男性か女性か取るためにコマンド長さから判定
	// 1つならばランダムな名前を返す、2つならば男女、3つ以上はエラーとする
	// ex) 「sfs rn male」「sfs rn」
	if len(messageArray) == 2 {
		name = gimei.NewName()
	} else if len(messageArray) == 3 {
		if messageArray[2] == "male" {
			name = gimei.NewMale()
		} else if messageArray[2] == "female" {
			name = gimei.NewFemale()
		} else {
			session.ChannelMessageSendReply(event.ChannelID, "Error!　コマンド指定が正しくありません。maleかfemaleを指定してください", event.MessageReference)
			return
		}
	} else {
		session.ChannelMessageSendReply(event.ChannelID, "Error!　コマンド指定が正しくありません。引数が多いです", event.MessageReference)
		return
	}
	response = "名前：" + name.Kanji() + "\nふりがな：" + name.Hiragana() + "\n性別：" + name.Sex.String()

	session.ChannelMessageSendReply(event.ChannelID, response, event.MessageReference)
}
