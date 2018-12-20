package bot

import (
	"github.com/rumblefrog/ChronoGG-Discord-Bot/sale"
	"github.com/rumblefrog/discordgo"
)

type ChronoBot struct {
	Session *discordgo.Session
	Router  chan *sale.Sale
}
