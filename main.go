package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/rumblefrog/ChronoGG-Discord-Bot/bot"
	"github.com/rumblefrog/ChronoGG-Discord-Bot/cache"
	"github.com/sirupsen/logrus"
)

func main() {
	bot.StartBot()

	cache.Tick()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	logrus.Info("Received exit signal. Terminating.")

	bot.Session.Close()
}
