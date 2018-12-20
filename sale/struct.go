package sale

import "time"

type Sale struct {
	Name        string
	URL         string
	UniqueURL   string
	SteamURL    string
	OGImage     string
	Platforms   []string
	PromoImage  string
	NormalPrice string
	Discount    string
	SalePrice   string
	StartDate   time.Time
	Items       []*Item
	EndDate     time.Time
	Currency    string
}

type Item struct {
	Type string
	ID   string
}
