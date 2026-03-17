package cmdtype

import (
	"fmt"
)

// to cmds from any to []cmdtype.Cmd
func ToCmds(incmds any) ([]Cmd, error) {
	// and watch these error msgs, if you don't know whats "verbose" lmao
	if cmds, ok := incmds.([]Cmd); ok {
		return cmds, nil
	}

	var res []Cmd

	// first layer: any slice
	slice1, ok := incmds.([]any)
	if !ok {
		return nil, fmt.Errorf("cmds must be of type []any for preset \"cmds\", got cmds %v of type %T", incmds, incmds)
	}
	for _, v := range slice1 {
		cmd, err := ToCmd(v)
		if err != nil {
			return nil, err
		}

		res = append(res, cmd)
	}

	return res, nil
}

func ToCmd(v any) (Cmd, error) {
	if cmd, ok := v.(Cmd); ok {
		return cmd, nil
	}

	// second layer: slice of any slice
	slice2, ok := v.([]any)
	if !ok {
		return nil, fmt.Errorf("cmds[i] must be of type []any for preset \"cmds\", got %v of type %T", v, v)
	}

	var cmd Cmd
	for _, t := range slice2 {
		// third layer: slice of string slice
		cmdElem, ok := t.(string)
		if !ok {
			return nil, fmt.Errorf("cmds[i][j] must be of type string for preset \"cmds\", got %v of type %T", t, t)
		}
		cmd = append(cmd, cmdElem)
	}

	return cmd, nil
}
