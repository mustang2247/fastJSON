//

package fastjson

import (
	"encoding/json"
	"log"
)

func ParseObject(jsonStr string) (*Value, error) {
	var p Parser
	v, err := p.Parse(jsonStr)
	if err != nil {
		log.Fatalf("cannot parse json: %s", err)
		return nil, err
	}

	return v, nil
}

func ParseArray(jsonStr string) ([]*Value, error) {
	var p Parser
	v, err := p.Parse(jsonStr)
	if err != nil {
		log.Fatalf("cannot parse json: %s", err)
		return nil, err
	}

	return v.GetArray(), nil
}

func ParseObjectByRawMessage(jsonStr json.RawMessage) (*Value, error) {
	var p Parser
	v, err := p.ParseBytes(jsonStr)
	if err != nil {
		log.Fatalf("cannot parse json: %s", err)
		return nil, err
	}

	return v, nil
}

func ParseArrayRawMessage(jsonStr json.RawMessage) ([]*Value, error) {
	var p Parser
	v, err := p.ParseBytes(jsonStr)
	if err != nil {
		log.Fatalf("cannot parse json: %s", err)
		return nil, err
	}

	return v.GetArray(), nil
}
