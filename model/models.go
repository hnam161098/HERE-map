package model

const (
	Key = "Skg1vxohQl5I32LLn_7tWUg6FOxpmzx56Vnqnw5TOnA"
	URL = "https://discover.search.hereapi.com/v1/discover"
)

type Items struct {
	Items []Item `json:"items"`
}

type Item struct {
	Title    string    `json:"title"`
	Addr     Address   `json:"address"`
	Position Positions `json:"position"`
}

type Address struct {
	Label       string `json:"label"`
	CountryCode string `json:"countryCode"`
	CountryName string `json:"countryName"`
	Country     string `json:"county"`
	City        string `json:"city"`
	District    string `json:"district"`
	Street      string `json:"street"`
	HouseNumer  string `json:"houseNumber"`
}

type Positions struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}
