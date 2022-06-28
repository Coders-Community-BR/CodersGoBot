package info

import (
	"github.com/Coders-Community-BR/CodersGoBot/commandHandler"
	"github.com/bwmarrin/discordgo"
)

func Ping(cx commandHandler.Context) {
	cx.ReplyText("Pong!")
}
func init() {
	commandHandler.AddCommand(Ping, &discordgo.ApplicationCommand{
		Name:        "ping",
		Description: "Pong!",
	})
}
