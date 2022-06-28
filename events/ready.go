package events

import (
	"github.com/Coders-Community-BR/CodersGoBot/util"
	"github.com/bwmarrin/discordgo"
	"time"
)

func Ready(s *discordgo.Session, m *discordgo.Ready) {
	guild, _ := s.Guild(util.GetEnv("GUILD_ID"))
	go func(s *discordgo.Session) {
		for {
			s.UpdateGameStatus(0, guild.Name)
			time.Sleep(20)
			s.UpdateGameStatus(0, "0 membros")
		}
	}(s)
}
