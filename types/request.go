package types

type Request struct {
	ID        string                 `json:"id"`
	Residence map[string]interface{} `json:"residence"`
	Activity  map[string]interface{} `json:"activity"`
	Info      map[string]interface{} `json:"info"`
}
