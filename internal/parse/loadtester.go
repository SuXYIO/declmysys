package parse

import (
	"reflect"
	"testing"

	"github.com/suxyio/declmysys/internal/parse/subs"
)

type Loadable interface {
	Load([]byte, subs.SubsDef) error
}

type TomlLoadRet struct {
	Result    Loadable
	ExpectErr bool
}

type TomlLoadTests map[string]TomlLoadRet

func RunTomlLoadTest[L Loadable](t *testing.T, tests TomlLoadTests, typeinst L) {
	// NOTE: the typeinst is a instance for the type to be loaded,
	// not really used though, only to determine type

	ltype := reflect.TypeOf(typeinst).Elem()

	// NOTE: Will be tested with default (empty) subsdef
	sd := subs.SubsDef{}

	for in, out := range tests {
		inst := reflect.New(ltype).Interface().(L)
		err := inst.Load([]byte(in), sd)

		// check
		if out.ExpectErr && err == nil {
			// won't check output if error is expected, since output might be meaningless
			t.Errorf("expected error for case \n%v\n, got error %v with result \n%v\n", in, err, inst)
		}
		if !out.ExpectErr {
			// not expect error
			if err != nil {
				t.Errorf("unexpected error for case \n%v\n, got error \"%v\" with result %v", in, err, inst)
			}

			// check if types match
			_, ok := out.Result.(L)
			if !ok {
				t.Errorf("types don't match, got typeinst of type %T and Result of type %T", typeinst, out.Result)
			}
			// check val
			if !reflect.DeepEqual(inst, out.Result) {
				t.Errorf("expected \n%v\n for case \n%v\n, got result \n%v\n", out.Result, in, inst)
			}
		}
	}
}
