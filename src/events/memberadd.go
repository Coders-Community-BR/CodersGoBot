package events

import (
	"github.com/Coders-Community-BR/CodersGoBot/src/util"
	"github.com/bwmarrin/discordgo"
)

func MemberAdd(s *discordgo.Session, m *discordgo.GuildMemberAdd) {
	channel := util.GetEnv("MEMBER_UPDATE_ID")
	guild, _ := s.Guild(util.GetEnv("GUILD_ID"))
	rules, _ := s.Channel(util.GetEnv("RULES_ID"))
	register, _ := s.Channel(util.GetEnv("REGISTER_ID"))
	channelsGuide, _ := s.Channel(util.GetEnv("CHANNELS_GUIDE_ID"))

	member := discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{
			IconURL: s.State.User.AvatarURL("1024"),
			Name:    "Sistema " + guild.Name,
		},
		Color: 0x0066ff,
		Title: "👋 Seja bem-vindo(a) " + m.User.Username + "!",
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: m.User.AvatarURL("1024"),
		},
		Description: "**Antes de continuar " + m.User.Username + ", consulte os canais:**\n🧾 _**Regras:**_ " + rules.Mention() + "\n📖 _**Registre-se:**_" + register.Mention() + "\n🔎 _**Guia de Canais:**_" + channelsGuide.Mention(),
		Footer: &discordgo.MessageEmbedFooter{
			Text:    "Sistema de Recepção",
			IconURL: guild.IconURL(),
		},
	}
	s.ChannelMessageSendComplex(channel, &discordgo.MessageSend{
		Content: m.User.Mention(),
		Embed:   &member,
	})
}
