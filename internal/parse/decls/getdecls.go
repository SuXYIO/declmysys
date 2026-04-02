package decls

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/suxyio/declmysys/internal/parse/metadata"
)

func GetDecls(ddir string) (Decls, error) {
	var decls Decls
	if declfiles, err := os.ReadDir(ddir); err != nil {
		return nil, err
	} else {
		for _, f := range declfiles {
			if f.IsDir() {
				// ignore
				if is, err := metadata.IsExcluded([]byte(f.Name())); err != nil {
					return decls, err
				} else if is {
					continue
				}

				// is decl
				var d Decl
				if err := d.Load(filepath.Join(ddir, f.Name())); err != nil {
					return nil, fmt.Errorf("%v in decl %s", err, f.Name())
				}
				decls = append(decls, d)
			}
		}
	}
	return decls, nil
}
