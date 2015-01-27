package arbreader

import (
	"encoding/json"
	"io"
	"strings"

	"github.com/juju/errgo"
)

type Message struct {
  Key string
  Type string
  Extended map[string]string
  Description string
}

func Read(reader io.Reader) ([]*Message, error) {
	data := map[string]interface{}{}
	if err := json.NewDecoder(reader).Decode(&data); err != nil {
		return nil, errgo.Mask(err)
	}

	messages := []*Message{}
	for key := range data {
		if key[0] == '@' {
			continue
		}

		vv := data["@" + key].(map[string]interface{})

		extended := map[string]string{}
		for ekey, evalue := range vv {
			if !strings.HasPrefix(ekey, "x-") {
				continue
			}

			extended[ekey] = evalue.(string)
		}

		messages = append(messages, &Message{
			Key: key,
			Type: vv["type"].(string),
			Description: vv["description"].(string),
			Extended: extended,
		})
	}

	return messages, nil
}
