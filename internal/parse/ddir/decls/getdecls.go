package decls

import (
	"os"
	"path/filepath"
)

func GetDecls(ddir string) (Decls, error) {
	var decls Decls
	if declfiles, err := os.ReadDir(filepath.Join(ddir, "decls")); err != nil {
		return nil, err
	} else {
		for _, f := range declfiles {
			if f.IsDir() {
				// is decl
				var d Decl
				d.Load(filepath.Join(ddir, "decls", f.Name()))
				decls = append(decls, d)
			}
		}
	}
	return decls, nil
}
