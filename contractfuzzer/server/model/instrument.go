package model

type InstrumentRequest struct {
	Name         string   `json:"name"`
	Input        string   `json:"input"`
	Instructions []uint64 `json:"instructions"`
}
