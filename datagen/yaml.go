package datagen

import "gopkg.in/yaml.v3"

type GenYAML struct {
}

func (g *GenYAML) Parse(source, destination string) error {
	var data yaml.Marshaler
	result, err := yaml.Marshal(data)
	if err != nil {
		return err
	}

	if err = g.output(result, destination); err != nil {
		return err
	}

	return nil
}

func (g *GenYAML) output(result []byte, destination string) error {
	return nil
}
