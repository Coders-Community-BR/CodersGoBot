package info

import (
	"fmt"
	"time"

	"github.com/Coders-Community-BR/CodersGoBot/commandHandler"
	"github.com/Coders-Community-BR/CodersGoBot/util"
	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
)

func ServerInfo(context commandHandler.Context) {
	guild := util.GetStateGuild(context.Session, context.Event.GuildID)
	if guild == nil {
		context.ReplyText("Não consegui encontrar o servidor.")
		return
	}
	owner, err := context.Session.User(guild.OwnerID)
	if err != nil {
		logrus.Errorf("Error getting owner: %s", err)
		context.ReplyText("Não consegui encontrar o dono do servidor.")
		return
	}
	rawCreatedAt, err := discordgo.SnowflakeTimestamp(guild.ID)
	createdAt := rawCreatedAt.Format("2006-01-02 ás 15:04:05")
	embed := &discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{
			Name:    fmt.Sprintf("Sistema %s", context.Session.State.User.Username),
			IconURL: context.Session.State.User.AvatarURL("1024"),
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text:    fmt.Sprintf("Solicitado por %s", util.GetUserNameByContext(&context)),
			IconURL: context.Event.Member.AvatarURL("1024"),
		},
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: guild.IconURL(),
		},
		Image: &discordgo.MessageEmbedImage{
			URL: guild.BannerURL(),
		},
		Timestamp:   time.Now().Format(time.RFC3339),
		Description: guild.Description,
		Title:       fmt.Sprintf("Informações de %s", guild.Name),
		Color:       util.GetRandomColor(),
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "🚀 Nome do Servidor",
				Value:  util.CodeBlock(guild.Name),
				Inline: true,
			},
			{
				Name:   "🗄 ID do Servidor",
				Value:  util.CodeBlock(guild.ID),
				Inline: true,
			},
			{
				Name:   "📅 Criado em",
				Value:  util.CodeBlock(createdAt),
				Inline: true,
			},
			{
				Name:   "🥇 Criado Por",
				Value:  util.CodeBlock(owner.String()),
				Inline: true,
			},
			{
				Name:   "👥 Número de Membros",
				Value:  util.CodeBlock(fmt.Sprintf("%d", guild.MemberCount)),
				Inline: true,
			},
			{
				Name:   "📚 Número de Canais",
				Value:  util.CodeBlock(fmt.Sprintf("%d", len(guild.Channels))),
				Inline: true,
			},
			{
				Name:   "🇧🇷 Linguagem Utilizada",
				Value:  util.CodeBlock(guild.PreferredLocale),
				Inline: true,
			},
			{
				Name:   "🌠 Nível de Impulsão",
				Value:  util.CodeBlock(fmt.Sprintf("%d", guild.PremiumTier)),
				Inline: true,
			},
			{
				Name:   "🦸‍♂️ Quantidade de Impulsionadores",
				Value:  util.CodeBlock(fmt.Sprintf("%d", guild.PremiumSubscriptionCount)),
				Inline: true,
			},
		},
	}
	err = context.ReplyEmbed(embed)
	if err != nil {
		logrus.Errorf("Error sending message: %v\n", err)
	}
}

func init() {
	commandHandler.AddCommand(ServerInfo, &discordgo.ApplicationCommand{
		Name:        "server",
		Description: "mostra as informações do servidor",
	})
}
