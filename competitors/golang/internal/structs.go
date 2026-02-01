package internal

import (
	"encoding/json"
	"errors"
	"strconv"
)

type Request struct {
	ID    string `json:"request_id"`
	Items []struct {
		ID    string   `json:"id"`
		Value Value    `json:"value"`
		Tags  []string `json:"tags"`
	} `json:"items"`
}

type Value float64

var _ json.Unmarshaler = new(Value)

func (v *Value) UnmarshalJSON(data []byte) error {
	res, err := strconv.ParseFloat(string(data), 64) //TODO: seems like bottleneck
	if errors.Is(err, strconv.ErrRange) {
		res = 0
		err = nil
	}

	*v = Value(res)

	return err
}

type Response struct {
	ID       string                  `json:"request_id"`
	Checksum []byte                  `json:"checksum,omitempty"`
	Stats    map[string]ResponseStat `json:"stats"`
}

type ResponseStat struct {
	Count uint64  `json:"count"`
	Sum   float64 `json:"sum"`
}
