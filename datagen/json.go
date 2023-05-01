package datagen

import "encoding/json"

type GenJSON struct {
}

func (g *GenJSON) Parse(source, destination string) error {
	var data json.Marshaler
	result, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if err = g.output(result, destination); err != nil {
		return err
	}

	return nil
}

func (g *GenJSON) output(result []byte, destination string) error {
	return nil
}
