package omake

import (
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

// msg 正規化されたメッセージ(string)を期待する
func JudgeDice(message string) bool {
	return strings.Contains(message, "!d")
}

func ResponseDiceRoll(message string, session *discordgo.Session, event *discordgo.MessageCreate) {
	array := strings.Split(message, " ")
	if len(array) > 2 {
		session.ChannelMessageSendReply(event.ChannelID, "Error!　コマンド指定が正しくありません", event.MessageReference)
		return
	}

	// すべて加算の前提のためダイスや足す数の区切りの判別に使う
	array = strings.Split(array[1], "+")
	rand.Seed(time.Now().UnixNano())
	// 各ダイスの結果を格納する
	var calcArray []int64

	for _, val := range array {
		// ダイスであるか、足す数かどうか判別する
		// ダイスであればダイスを振り次のループに行く
		if strings.Contains(val, "d") {
			roll := strings.Split(val, "d")
			if len(roll) > 2 {
				session.ChannelMessageSendReply(event.ChannelID, "Error!　dの数が多いです", event.MessageReference)
				return
			}
			dice, err2 := strconv.Atoi(roll[1])
			count, err1 := strconv.Atoi(roll[0])
			if err1 != nil || err2 != nil {
				session.ChannelMessageSendReply(event.ChannelID, "Error!　数字以外のものが指定されました", event.MessageReference)
				return
			}
			for i := 0; i < count; i++ {
				calcArray = append(calcArray, rand.Int63n(int64(dice))+1)
			}
			continue
		}
		// ダイスでなければcalcArrayに数字を加算する
		sum, err := strconv.Atoi(val)
		if err != nil {
			session.ChannelMessageSendReply(event.ChannelID, "Error!　数字以外のものが指定されました", event.MessageReference)
			return
		}
		calcArray = append(calcArray, int64(sum))
	}
	response := "( "
	var total int64
	for index, calc := range calcArray {
		total = total + calc
		response = response + strconv.Itoa(int(calc))
		if (index + 1) != len(calcArray) {
			response = response + " + "
		} else {
			response = response + " )"
		}
	}
	session.ChannelMessageSendReply(event.ChannelID, strconv.Itoa(int(total))+" "+response, event.MessageReference)
}
