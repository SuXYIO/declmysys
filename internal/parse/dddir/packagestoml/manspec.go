package packagestoml

import (
	"fmt"

	"github.com/suxyio/declmysys/internal/parse/cmdtype"
)

// designed for custom multi-type unmarshal
type manSpec struct {
	Preset    string
	CustomCmd cmdtype.Cmd
}

func (m *manSpec) UnmarshalTOML(data any) error {
	switch v := data.(type) {
	case string:
		// preset
		m.Preset = v
		m.CustomCmd = nil
	case []any:
		if len(v) == 0 {
			return fmt.Errorf("cmd must not be empty list")
		}
		// might be cmd list
		for _, p := range v {
			s, ok := p.(string)
			if !ok {
				return fmt.Errorf("if using cmd, elements must be of type string, got %v of type %T", p, p)
			}
			m.CustomCmd = append(m.CustomCmd, s)
		}
		m.Preset = ""
	default:
		// not recognized
		return fmt.Errorf("manager must be of type string or []string, got %v of type %T", data, data)
	}
	return nil
}
