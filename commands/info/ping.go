package info

import (
	"github.com/Coders-Community-BR/CodersGoBot/commandHandler"
	"github.com/bwmarrin/discordgo"
)

func ping(cx commandHandler.Context) {
	cx.ReplyText("Pong!")
}
func init() {
	commandHandler.AddCommand(ping, &discordgo.ApplicationCommand{
		Name:        "ping",
		Description: "Pong!",
	})
}
