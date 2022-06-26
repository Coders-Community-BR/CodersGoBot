package events

import (

	"github.com/bwmarrin/discordgo"
)

func Ready(s *discordgo.Session, m *discordgo.Ready) {
	buf := "Em desenvolvimento!"
	s.UpdateGameStatus(0, buf)
}
