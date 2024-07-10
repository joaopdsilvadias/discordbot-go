package bot

import (
	"discordbot/bot/commands"

	"github.com/bwmarrin/discordgo"
)

func CommandHandler(discord *discordgo.Session, message *discordgo.MessageCreate) {
	switch message.Content {
	case "!ping":
		commands.GetDolarPrice(discord, message)
	}
}
