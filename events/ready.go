package events

import (
	"fmt"
	"time"

	_ "github.com/Coders-Community-BR/CodersGoBot/commands/info"
	_ "github.com/Coders-Community-BR/CodersGoBot/commands/staff"
	"github.com/Coders-Community-BR/CodersGoBot/handler"
	"github.com/Coders-Community-BR/CodersGoBot/util"
	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
)

func Ready(s *discordgo.Session, m *discordgo.Ready) {
	guild, err := s.Guild(util.GetEnv("GUILD_ID"))
	if err != nil {
		logrus.Panic(err)
	}
	logrus.Info("Adding commands...")
	handler.RegisterCommands(s, guild)
	go func(s *discordgo.Session) {
		for {
			s.UpdateGameStatus(0, guild.Name)
			time.Sleep(time.Duration(5 * time.Second))
			if guild := util.GetStateGuild(s, util.GetEnv("GUILD_ID")); guild != nil {
				s.UpdateListeningStatus(fmt.Sprintf("%d membros", guild.MemberCount))
			}
		}
	}(s)
}
