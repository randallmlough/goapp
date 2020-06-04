package template

import "github.com/randallmlough/gogen"

type Module struct {
	Name       string      `yaml:"name"`
	Required   bool        `yaml:"required"`
	Default    bool        `yaml:"default"`
	Components []Component `yaml:"components"`
}
type Modules []Module

func (mods Modules) Slice() []string {
	m := make([]string, 0, len(mods))
	for _, mod := range mods {
		m = append(m, mod.Name)
	}
	return m
}

func (mod *Module) Files(templateDir string) ([]gogen.File, error) {
	files := make([]gogen.File, 0, len(mod.Components))
	for _, component := range mod.Components {
		file, err := component.File(templateDir)
		if err != nil {
			return nil, err
		}
		files = append(files, file)
	}
	return files, nil
}
