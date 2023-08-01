package types

type Request struct {
	ID        string         `json:"id"`
	Residence map[string]any `json:"residence"`
	Activity  map[string]any `json:"activity"`
	Info      map[string]any `json:"info"`
}

type InfoRequest struct {
	Info map[string]any `json:"info"`
}
