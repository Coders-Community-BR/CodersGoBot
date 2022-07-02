package info

import (
	"github.com/Coders-Community-BR/CodersGoBot/handler"
	"github.com/bwmarrin/discordgo"
)

func Ping(cx handler.Context) {
	cx.ReplyText("Pong!")
}
func init() {
	handler.AddCommand(Ping, &discordgo.ApplicationCommand{
		Name:        "ping",
		Description: "Pong!",
	})
}
