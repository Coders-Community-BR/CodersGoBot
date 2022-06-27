package main

import (
	"github.com/Coders-Community-BR/CodersGoBot/events"
	"github.com/Coders-Community-BR/CodersGoBot/util"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"github.com/bwmarrin/discordgo"
)

func main() {

	token := util.GetEnv("TOKEN")

	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	dg.AddHandler(events.Ready)
	dg.AddHandler(events.MessageCreate)
	dg.AddHandler(events.MemberAdd)
	dg.AddHandler(events.MemberRemoved)

	dg.Identify.Intents = discordgo.IntentsAll

	err = dg.Open()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("\033[0;34m| CODERS COMMUNITY BOT |\033[0m\n| BOT ONLINE | ID: %s |\n| PRESSIONE CTRL-C PARA PARAR |\n", dg.State.User.ID)
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()
}
