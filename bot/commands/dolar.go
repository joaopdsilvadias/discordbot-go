package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/bwmarrin/discordgo"
)

type CurrencyData struct {
	Code       string `json:"code"`
	Codein     string `json:"codein"`
	Name       string `json:"name"`
	High       string `json:"high"`
	Low        string `json:"low"`
	VarBid     string `json:"varBid"`
	PctChange  string `json:"pctChange"`
	Bid        string `json:"bid"`
	Ask        string `json:"ask"`
	Timestamp  string `json:"timestamp"`
	CreateDate string `json:"create_date"`
}

func GetDolarPrice(discord *discordgo.Session, message *discordgo.MessageCreate) {

	resp, err := http.Get("https://economia.awesomeapi.com.br/json/last/USD")

	if err != nil {
		log.Println("Error getting dolar price: ", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Println("Error reading response body: ", err)
	}

	var data map[string]CurrencyData

	err = json.Unmarshal([]byte(body), &data)
	if err != nil {
		fmt.Println("Erro ao fazer parse do JSON:", err)
		return
	}

	for _, value := range data {
		discord.ChannelMessageSendEmbed(message.ChannelID, &discordgo.MessageEmbed{
			Title: "Dolar Price",
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:  "Bid",
					Value: value.Bid,
				},
				{
					Name:  "Ask",
					Value: value.Ask,
				},
				{
					Name:  "High",
					Value: value.High,
				},
				{
					Name:  "Low",
					Value: value.Low,
				},
			},
		})
	}

}
