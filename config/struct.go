package config

type Config struct {
	Token     string `toml:"Token"`
	ChannelID int64  `toml:"ChannelID"`
}
