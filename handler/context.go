package handler

import "github.com/bwmarrin/discordgo"

type Context struct {
	Session *discordgo.Session
	Event   discordgo.InteractionCreate
}

func (context *Context) ReplyText(content string) error {
	err := context.Session.InteractionRespond(context.Event.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: content,
		},
	})
	return err
}

func (context *Context) ReplyEmbed(embed *discordgo.MessageEmbed) error {
	err := context.Session.InteractionRespond(context.Event.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{embed},
		},
	})
	return err
}
