package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/FEATO-org/support-feato-system/config"
	"github.com/FEATO-org/support-feato-system/infrastructure"
	"github.com/FEATO-org/support-feato-system/interfaces"
	"github.com/FEATO-org/support-feato-system/usecase"
	"github.com/bwmarrin/discordgo"
	"golang.org/x/oauth2"
)

var (
	discordToken string
	guildIDList  []string
	ctx          context.Context
	dbtx         *sql.DB
	oauthConfig  *oauth2.Config
)

func init() {
	discordToken = os.Getenv("DISCORD_TOKEN")
	// 仮置き
	guildIDList = strings.Split(os.Getenv("GUILD_IDS"), ",")

	ctx = context.Background()
	dbtx = config.NewDB()
	oauthConfig = config.NewOauth2()
	log.SetFlags(log.Llongfile)
}

func main() {
	// dice
	diceUsecase := usecase.NewDiceUsecase(infrastructure.NewDiceRepository())
	diceInterface := interfaces.NewDiceInterfaces(diceUsecase)
	// character
	characterUsecase := usecase.NewCharacterUsecase(infrastructure.NewCharacterRepository())
	characterInterfaces := interfaces.NewCharacterInterfaces(characterUsecase)

	discordInterfaces := interfaces.NewDiscordInterfaces(diceInterface, characterInterfaces, guildIDList)

	// discordへの接続と初期化処理
	dg, err := discordgo.New("Bot " + discordToken)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}
	dg.AddHandlerOnce(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
	})
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}
	fmt.Println("Bot open connection.  Press CTRL-C to exit.")

	// コマンドやハンドラーの登録
	discordInterfaces.CreateApplicationCommand(dg)
	discordInterfaces.AddCommandHandler(dg)
	discordInterfaces.AddMessageHandler(dg)

	fmt.Println("Bot is now running.")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, syscall.SIGTERM)
	<-sc
	discordInterfaces.DeleteApplicationCommand(dg)
	dg.Close()
	fmt.Println("Bot is shutdown.")
}
