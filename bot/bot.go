package bot

import (
	"github.com/rumblefrog/ChronoGG-Discord-Bot/config"
	"github.com/rumblefrog/ChronoGG-Discord-Bot/sale"
	"github.com/rumblefrog/discordgo"
	"github.com/sirupsen/logrus"
)

var Chrono *ChronoBot

func StartBot() {
	session, err := discordgo.New("Bot " + config.Conf.Token)

	if err != nil {
		logrus.WithField("error", err).Fatal("Unable to initiate bot session")
	}

	Chrono = &ChronoBot{
		Session: session,
		Router:  make(chan *sale.Sale),
	}

	session.AddHandler(ready)

	err = session.Open()

	if err != nil {
		logrus.WithField("error", err).Fatal("Unable to open bot session")
	}

	Chrono.Listen()
}

func ready(s *discordgo.Session, event *discordgo.Ready) {
	logrus.WithFields(logrus.Fields{
		"Username":    event.User.Username,
		"Guild Count": len(event.Guilds),
	}).Info("Bot is now running")
}
