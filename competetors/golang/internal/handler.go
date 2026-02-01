package internal

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"

	"github.com/gofiber/fiber/v3"
)

func Handler(c fiber.Ctx) error {
	request := new(Request)
	err := c.Bind().Body(request)
	if err != nil {
		return err
	}

	response := &Response{
		ID:       request.ID,
		Stats:    make(map[string]ResponseStat),
		Checksum: make([]byte, 64),
	}

	for _, item := range request.Items {
		for _, tag := range item.Tags {
			if stat, ok := response.Stats[tag]; ok {
				stat.Count++
				stat.Sum += float64(item.Value)
				response.Stats[tag] = stat
			} else {
				stat = ResponseStat{
					Count: 1,
					Sum:   float64(item.Value),
				}
				response.Stats[tag] = stat
			}
		}
	}

	respJson, err := json.Marshal(response)
	if err != nil {
		return err
	}

	h := sha256.New()
	h.Write(respJson)
	hash := h.Sum(nil)
	hex.Encode(response.Checksum, hash)

	return c.JSON(response)
}
