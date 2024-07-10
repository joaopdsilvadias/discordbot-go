package bot

import (
	"discordbot/config"
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
)

func Start(token string) {

	discord, err := discordgo.New("Bot " + config.GetEnv(token))

	if err != nil {
		log.Println("Error creating Discord session: ", err)
	}

	err = discord.Open()

	if err != nil {
		log.Println("Error opening connection: ", err)
	}

	defer discord.Close()

	discord.AddHandler(CommandHandler)

	log.Println("Bot is running!")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}
