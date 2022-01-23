package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/FEATO-org/support-feato-system/src/omake"
	"github.com/FEATO-org/support-feato-system/src/utility"
	"github.com/bwmarrin/discordgo"
)

func main() {

	token := os.Getenv("discordToken")
	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// In this example, we only care about receiving message events.
	// dg.Identify.Intents = discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, syscall.SIGTERM)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}
	// If the message is "ping" reply with "Pong!"
	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
		return
	}

	// If the message is "pong" reply with "Ping!"
	if m.Content == "pong" {
		s.ChannelMessageSend(m.ChannelID, "Ping!")
		return
	}

	msg := beforeMessageNormalization(m.Content)

	if diceJudge(msg) {
		diceRoll(msg, s, m)
	}
	if omake.JudgeRandomName(msg) {
		omake.ResponseRandomName(msg, s, m)
		return
	}
}

func beforeMessageNormalization(text string) string {
	var rules []string
	targetStrings := []string{"d", "!", "+"}
	for _, str := range targetStrings {
		rules = append(rules, buildnewReplacerRules(str)...)
	}
	rules = append(rules, "　", " ")
	rep := strings.NewReplacer(rules...)
	response := rep.Replace(text)
	return response
}

// 半角英小文字を与え、対象の半角小文字を全角半角小文字大文字の変換結果をNewReplacerで置き換え可能な形にする配列を返す
func buildnewReplacerRules(text string) []string {
	var rules []string
	replaceStrings := utility.BuildReplaceString(text)
	for _, replace := range replaceStrings {
		rules = append(rules, replace, text)
	}
	return rules
}

func diceJudge(message string) bool {
	return strings.Contains(message, "!d")
}

func diceRoll(message string, session *discordgo.Session, event *discordgo.MessageCreate) {
	array := strings.Split(message, " ")
	if len(array) > 2 {
		session.ChannelMessageSendReply(event.ChannelID, "Error!　コマンド指定が正しくありません", event.MessageReference)
		return
	}
	array = strings.Split(array[1], "+")
	rand.Seed(time.Now().UnixNano())
	var calcArray []int64
	for _, val := range array {
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
