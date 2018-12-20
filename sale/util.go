package sale

func (s *Sale) GetHeader() string {
	return s.Name + " for " + s.SalePrice + " " + s.Currency + " (" + s.Discount + " off)!"
}

func (s *Sale) GetNormalPrice() string {
	return s.NormalPrice + " " + s.Currency
}

func (s *Sale) GetSalePrice() string {
	return s.SalePrice + " " + s.Currency
}
