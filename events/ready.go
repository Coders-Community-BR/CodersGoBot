package events

import (
	"log"
	"time"

	"github.com/Coders-Community-BR/CodersGoBot/commandHandler"
	_ "github.com/Coders-Community-BR/CodersGoBot/commands/info"
	"github.com/Coders-Community-BR/CodersGoBot/util"
	"github.com/bwmarrin/discordgo"
)

func Ready(s *discordgo.Session, m *discordgo.Ready) {

	guild, err := s.Guild(util.GetEnv("GUILD_ID"))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Adding commands...")
	commandHandler.RegisterCommands(s, guild)
	go func(s *discordgo.Session) {
		for {
			s.UpdateGameStatus(0, guild.Name)
			time.Sleep(20)
			s.UpdateGameStatus(0, "0 membros")
		}
	}(s)
}
