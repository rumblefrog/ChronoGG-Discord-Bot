package cache

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/rumblefrog/ChronoGG-Discord-Bot/bot"
	"github.com/rumblefrog/ChronoGG-Discord-Bot/sale"
	"github.com/sirupsen/logrus"
)

func Poll() {
	res, err := http.Get("https://api.chrono.gg/sale")

	if err != nil {
		logrus.WithField("error", err).Error("Failed to poll sale")
		return
	}

	defer res.Body.Close()

	bytes, err := ioutil.ReadAll(res.Body)

	if err != nil {
		logrus.WithField("error", err).Error("Failed to read body to bytes")
		return
	}

	sale := &sale.Sale{}

	json.Unmarshal(bytes, sale)

	new, err := CompareSale(sale)

	if err != nil {
		logrus.WithField("error", err).Error("Unable to read sale cache")

		return
	}

	if new {
		bot.Chrono.Router <- sale
		WriteSale(sale)
	}
}
