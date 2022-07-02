package util

import (
	"fmt"

	"github.com/Coders-Community-BR/CodersGoBot/handler"
	"github.com/bwmarrin/discordgo"
)

func CodeBlock(code interface{}) string {
	return fmt.Sprintf("`%s`", code)
}

func GetUserNameByContext(context *handler.Context) string {
	member := context.Event.Member
	user := context.Event.Member.User
	if member.Nick != "" {
		return member.Nick
	} else {
		return user.Username
	}
}

func GetStateGuild(session *discordgo.Session, guildID string) *discordgo.Guild {
	for _, guild := range session.State.Guilds {
		if guild.ID == guildID {
			return guild
		}
	}
	return nil
}
