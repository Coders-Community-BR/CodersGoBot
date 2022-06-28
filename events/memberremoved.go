package events

import (
	"time"

	"github.com/Coders-Community-BR/CodersGoBot/util"
	"github.com/bwmarrin/discordgo"
)

func MemberRemoved(s *discordgo.Session, m *discordgo.GuildMemberRemove) {
	channel := util.GetEnv("MEMBER_UPDATE_ID")
	guild, _ := s.Guild(util.GetEnv("GUILD_ID"))

	member := discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{
			Name:    "Sistema " + guild.Name,
			IconURL: s.State.User.AvatarURL("1024"),
		},
		Color: util.GetRandomColor(),
		Title: "🖐 Adeus " + m.User.Username + "!",
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: m.User.AvatarURL("1024"),
		},
		Description: "**Foi um prazer tê-lo em nossa comunidade, mas parece que agora nos despedimos.**",
		Footer: &discordgo.MessageEmbedFooter{
			Text:    "Sistema de Recepção",
			IconURL: guild.IconURL(),
		},
		Timestamp: time.Now().Format(time.RFC3339),
	}
	s.ChannelMessageSendEmbed(channel, &member)
}
