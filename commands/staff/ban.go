package staff

import (
	"fmt"
	"time"

	"github.com/Coders-Community-BR/CodersGoBot/handler"
	"github.com/Coders-Community-BR/CodersGoBot/util"
	"github.com/bwmarrin/discordgo"
)


func Ban(cx handler.Context){
	channel := util.GetEnv("PUNISHMENT_CHANNEL_ID")
	if channel == "" {
		channel = cx.Event.ChannelID
	}

	options := cx.Event.ApplicationCommandData().Options
	optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
	for _, opt := range options {
		optionMap[opt.Name] = opt
	}

	// TODO: Check if dias exists on the optionMap, if not dias = 0.
	var dias = int(optionMap["dias"].IntValue())
	if dias > 7 || dias < 0 {
		cx.ReplyText("❌ Você deve informar de 0 a 7 dias.")
	}
	
	if optionMap["usuario"].Value == cx.Event.Member.User.ID {
		cx.ReplyText("Você não pode se banir.")
		return
	}
	if optionMap["usuario"].Value == cx.Session.State.User.ID {
		cx.ReplyText("Eu não posso me banir.")
		return
	}
	bannedUser, _ := cx.Session.User(optionMap["usuario"].Value.(string))
	err := cx.Session.GuildBanCreateWithReason(cx.Event.GuildID, optionMap["usuario"].Value.(string), optionMap["motivo"].StringValue(), dias)
	if err != nil {
		cx.ReplyText("Não foi possível banir o usuário: " + err.Error())
		return
	}

	cx.ReplyText("Usuário banido.")

	embed := discordgo.MessageEmbed{
			Author: &discordgo.MessageEmbedAuthor{
			Name: fmt.Sprintf("Sistema %s", cx.Session.State.User.Username),
			IconURL:  cx.Session.State.User.AvatarURL("1024"),
		},
		Color: 0x0000CD,
		Timestamp: time.Now().Format(time.RFC3339),
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: bannedUser.AvatarURL("1024"),
		},
		Title: "🚫 Novo Usuário Banido",
		Description: fmt.Sprintf("> **Usuário:** `%s` \n> **Razão:** `%s` \n> **Dias de Mensagens Apagadas:** `%d`",
	bannedUser.Username, optionMap["motivo"].StringValue(), dias),
	Footer: &discordgo.MessageEmbedFooter{
		Text: "Solicitado por " + util.GetUserNameByContext(&cx),
		IconURL: cx.Event.Member.AvatarURL("1024"),
	},
	}
	cx.Session.ChannelMessageSendComplex("", &discordgo.MessageSend{
		Embed: &embed,
	})
}


var defaultMemberPermissions int64 = discordgo.PermissionBanMembers

func init(){
	handler.AddCommand(Ban, &discordgo.ApplicationCommand{
		Name:	"ban",
		Description: "Bane um membro do servidor.",
		DefaultMemberPermissions: &defaultMemberPermissions,
		Options: []*discordgo.ApplicationCommandOption{

			{
				Type: discordgo.ApplicationCommandOptionUser,
				Name: "usuario",
				Description: "Usuário a ser banido.",
				Required: true,
			},

			{
				Type: discordgo.ApplicationCommandOptionString,
				Name: "motivo",
				Description: "Motivo pelo qual o usuário está sendo banido.",
				Required: true,
			},

			{
				Type: discordgo.ApplicationCommandOptionInteger,
				Name: "dias",
				Description: "Quantos dias de mensagens do histórico serão apagadas.",
				Required: true,
			},
		},
	})
}

