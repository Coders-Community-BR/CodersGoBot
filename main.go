package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Coders-Community-BR/CodersGoBot/commandHandler"
	"github.com/Coders-Community-BR/CodersGoBot/events"
	"github.com/Coders-Community-BR/CodersGoBot/util"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func discordGoLogger(level, caller int, format string, args ...interface{}) {
	switch level {
	case discordgo.LogError:
		logrus.Errorf(format, args...)
	case discordgo.LogWarning:
		logrus.Warnf(format, args...)
	case discordgo.LogInformational:
		logrus.Infof(format, args...)
	case discordgo.LogDebug:
		logrus.Debugf(format, args...)
	}
}

func main() {
	discordgo.Logger = discordGoLogger
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	if err := godotenv.Load(); err != nil {
		logrus.Panicf("Error loading .env file: %v", err)
	}
	token := util.GetEnv("TOKEN")
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		logrus.Fatalf("Error creating Discord session: %s", err)
		return
	}
	dg.LogLevel = discordgo.LogWarning
	dg.StateEnabled = true

	dg.AddHandler(events.Ready)
	dg.AddHandler(events.MessageCreate)
	dg.AddHandler(events.MemberAdd)
	dg.AddHandler(events.MemberRemoved)
	dg.AddHandler(commandHandler.Get)
	dg.Identify.Intents = discordgo.IntentsAll

	coders, err := dg.Guild(os.Getenv("GUILD_ID"))
	if err != nil {
		log.Fatal(err)
	}
	if err := dg.Open(); err != nil {
		log.Fatal(err)
	}
	defer dg.Close()
	defer commandHandler.RemoveCommands(dg, coders)

	fmt.Printf(
		"\033[0;34m| CODERS COMMUNITY BOT |\033[0m\n| BOT ONLINE | ID: %s |\n| PRESSIONE CTRL-C PARA PARAR |\n",
		dg.State.User.ID,
	)
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}
