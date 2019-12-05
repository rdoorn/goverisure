package goverisure

import "encoding/json"

/*
{"data":{"count":1,"telemetries":[{"date":"2019-12-01 11:04:10","totalActivePower":695.0,"dcVoltage":748.125,"groundFaultResistance":7629.17,"powerLimit":100.0,"totalEnergy":10112.0,"temperature":35.0047,"inverterMode":"MPPT","operationMode":0,
"vL1To2":395.531,"vL2To3":394.562,"vL3To1":398.688,"L1Data":{"acCurrent":1.13281,"acVoltage":230.219,"acFrequency":50.0101,"apparentPower":261.0,"activePower":234.0,"reactivePower":-115.0,"cosPhi":1.0},
"L2Data":{"acCurrent":1.10938,"acVoltage":228.188,"acFrequency":50.0082,"apparentPower":260.0,"activePower":228.0,"reactivePower":-107.0,"cosPhi":1.0},
"L3Data":{"acCurrent":1.13281,"acVoltage":228.141,"acFrequency":50.0098,"apparentPower":252.0,"activePower":233.0,"reactivePower":-115.0,"cosPhi":1.0}}]}}
*/

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
