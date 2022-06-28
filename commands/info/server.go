package info

import (
	"strconv"
	"time"

	"github.com/Coders-Community-BR/CodersGoBot/commandHandler"
	"github.com/Coders-Community-BR/CodersGoBot/util"
	"github.com/bwmarrin/discordgo"
)

func codesnippet(content string) string {
	return "`" + content + "`"
}

func serverInfo(cx commandHandler.Context) {
	guild, _ := cx.Session.Guild(util.GetEnv("GUILD_ID"))

	info := discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{
			Name:    "Sistema " + guild.Name,
			IconURL: cx.Session.State.User.AvatarURL("1024"),
		},
		Color:     util.GetRandomColor(),
		Title:     "Informações para " + guild.Name,
		Timestamp: time.Now().Format(time.RFC3339),
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "🚀 Nome do Servidor",
				Value:  codesnippet(guild.Name),
				Inline: true,
			},
			{
				Name:   "🔖 Id do Servidor",
				Value:  codesnippet(guild.ID),
				Inline: true,
			},
			{
				Name:   "🥇 Criado Por",
				Value:  codesnippet("<@" + guild.OwnerID + ">"),
				Inline: true,
			},
			{
				Name:   "👥 Número de Membros",
				Value:  codesnippet(strconv.Itoa(guild.MemberCount)),
				Inline: true,
			},
			{
				Name:   "🎫 Total de Canais",
				Value:  codesnippet(strconv.Itoa(len(guild.Channels))),
				Inline: true,
			},
			{
				Name:   "🇧🇷 Linguagem Utilizada",
				Value:  codesnippet(guild.PreferredLocale),
				Inline: true,
			},
			{
				Name:   "🌠 Nível de Impulsão",
				Value:  codesnippet(strconv.Itoa(int(guild.PremiumTier))),
				Inline: true,
			},
			{
				Name:   "🦸‍♂️ Quantidade de Impulsionadores",
				Value:  codesnippet(strconv.Itoa(guild.PremiumSubscriptionCount)),
				Inline: true,
			},
		},
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: guild.IconURL(),
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text:    "Solicitado por" + cx.Session.State.User.Username,
			IconURL: cx.Session.State.User.AvatarURL("1024"),
		},
	}

	cx.ReplyEmbed(&info)
}
func init() {
	commandHandler.AddCommand(serverInfo, &discordgo.ApplicationCommand{
		Name:        "server",
		Description: "〔❓ Info〕 Retorna informações a respeito do servidor da Coders Community",
	})
}
