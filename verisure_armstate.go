package goverisure

import (
	"encoding/json"
)

func (h *Handler) ArmState(installationID string) (ArmState, error) {
	o := ArmState{}
	data, err := h.Get("GET", urlGetArmState, h.apiURL, installationID)

	err = json.Unmarshal(data, &o)
	if err != nil {
		return ArmState{}, err
	}

	return o, nil
}
