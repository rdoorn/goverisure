package goverisure

import "encoding/json"

type Installation struct {
	Giid            string `json:"giid"`
	FirmwareVersion int    `json:"firmwareVersion"`
	RoutingGroup    string `json:"routingGroup"`
	Shard           int    `json:"shard"`
	Locale          string `json:"locale"`
	CountryID       string `json:"countryId"`
	SignalFilterID  int    `json:"signalFilterId"`
	Deleted         bool   `json:"deleted"`
	Cid             string `json:"cid"`
	Street          string `json:"street"`
	StreetNo1       string `json:"streetNo1"`
	StreetNo2       string `json:"streetNo2"`
	Alias           string `json:"alias"`
	CustomerType    string `json:"customerType"`
}

func (h *Handler) Installations() ([]Installation, error) {
	var i []Installation
	data, err := h.Get("GET", urlGetInstallations, h.apiURL, h.username)

	err = json.Unmarshal(data, &i)
	if err != nil {
		return nil, err
	}

	return i, nil
}
