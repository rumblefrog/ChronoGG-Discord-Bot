package cache

import "time"

const CacheFile = "chrono.cache"

var Ticker *time.Ticker

func Tick() {
	Ticker = time.NewTicker(time.Hour)

	go func() {
		for range Ticker.C {
			Poll()
		}
	}()
}
