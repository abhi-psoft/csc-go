package main

type Country struct {
	Name      string `json:"name"`
	IsoCode   string `json:"isoCode"`
	Flag      string `json:"flag"`
	Phonecode string `json:"phonecode"`
	Currency  string `json:"currency"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

type State struct {
	Name        string `json:"name"`
	IsoCode     string `json:"isoCode"`
	CountryCode string `json:"countryCode"`
	Latitude    string `json:"latitude"`
	Longitude   string `json:"longitude"`
}

type City struct {
	Name        string `json:"name"`
	CountryCode string `json:"countryCode"`
	StateCode   string `json:"stateCode"`
	Latitude    string `json:"latitude"`
	Longitude   string `json:"longitude"`
}
