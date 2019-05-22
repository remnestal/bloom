package bloom

import (
	"encoding/json"
)

// translate is a transition struct for encoding/decoding of Filter data.
// The purpose of this struct is to avoid exposing the internals of the Filter
// just to enable marshaling.
type translate struct {
	Data    []uint64 `json:"data"`
	Lookups int      `json:"lookups"`
	Count   int64    `json:"count"`
}

// EncodeJSON returns the JSON representation of a Bloom filter
func (f *Filter) EncodeJSON() ([]byte, error) {
	return json.Marshal(translate{f.data, f.lookups, f.count})
}

// DecodeJSON creates a new Bloom filter from a JSON byte string
func DecodeJSON(data []byte) (*Filter, error) {
	var t translate
	err := json.Unmarshal(data, &t)
	return &Filter{t.Data, t.Lookups, t.Count}, err
}
