package commandHandler

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

var commandRegistry = map[*discordgo.ApplicationCommand]func(Context){}
var registredCommands []*discordgo.ApplicationCommand

func AddCommand(handler func(Context), command *discordgo.ApplicationCommand) {
	commandRegistry[command] = handler
}
func Get(session *discordgo.Session, event *discordgo.InteractionCreate) {
	for command, handler := range commandRegistry {
		if command.Name == event.ApplicationCommandData().Name {
			handler(Context{session, *event})
		}
	}
}
func RegisterCommands(session *discordgo.Session, guild *discordgo.Guild) {
	for command := range commandRegistry {
		cmd, err := session.ApplicationCommandCreate(session.State.User.ID, guild.ID, command)
		if err != nil {
			log.Fatalf("Error creating command: %s", err)
		} else {
			log.Printf("Created command: %s", command.Name)
		}
		registredCommands = append(registredCommands, cmd)
	}
}
func RemoveCommands(session *discordgo.Session, guild *discordgo.Guild) {
	for _, command := range registredCommands {
		err := session.ApplicationCommandDelete(session.State.User.ID, guild.ID, command.ID)
		if err != nil {
			log.Panicf("Cannot delete '%v' command: %v", command.Name, err)
		}
		log.Println("Removed command:", command.Name)
	}
}
