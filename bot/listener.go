package bot

import (
	"github.com/rumblefrog/ChronoGG-Discord-Bot/config"
	"github.com/rumblefrog/discordgo"
)

func (c *ChronoBot) Listen() {
	for {
		select {
		case sale := <-c.Router:
			c.Session.ChannelMessageSendEmbed(
				config.Conf.ChannelID,
				&discordgo.MessageEmbed{
					Color:       0x220F33,
					Description: "\u200B",
					Title:       sale.GetHeader(),
					URL:         "https://chrono.gg",
					Timestamp:   sale.StartDate.String(),
					Footer: &discordgo.MessageEmbedFooter{
						IconURL: "https://cdn.chrono.gg/assets/images/favicon/favicon-32x32.png",
						Text:    "ChronoGG",
					},
					Image: &discordgo.MessageEmbedImage{
						URL: sale.PromoImage,
					},
					Fields: []*discordgo.MessageEmbedField{
						&discordgo.MessageEmbedField{
							Name:  "Normal Price",
							Value: sale.GetNormalPrice(),
						},
						&discordgo.MessageEmbedField{
							Name:  "Sale Price",
							Value: sale.GetSalePrice(),
						},
					},
				},
			)
		}
	}
}
