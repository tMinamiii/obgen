package datagen

type DataFormat string

const (
	JSON DataFormat = "json"
	YAML DataFormat = "yaml"
)

func (d DataFormat) String() string {
	return string(d)
}

type DataGen interface {
	Parse(source, destination string) error
}

func NewDataGen(format DataFormat) DataGen {
	switch format {
	case JSON:
		return &GenJSON{}
	case YAML:
		return &GenYAML{}
	default:
		return nil
	}

}
