package template

import (
	"github.com/randallmlough/gogen"
	"path"
)

type Package struct {
	OutputDir    string            `yaml:"output_dir"`
	TemplateDir  string            `yaml:"template_dir"`
	GoModEnabled bool              `yaml:"go_mod"`
	Deps         []string          `yaml:"deps"`
	Modules      map[string]Module `yaml:"modules"`
}

func (pkg *Package) Files(baseDir string) ([]gogen.File, error) {
	files := []gogen.File{}
	for _, module := range pkg.Modules {
		ff, err := module.Files(path.Join(baseDir, pkg.TemplateDir))
		if err != nil {
			return nil, err
		}
		files = append(files, ff...)
	}
	return files, nil
}

func (pkg *Package) SelectModules(modules ...string) *Package {
	p := pkg
	mods := make(map[string]Module)
	for _, module := range modules {
		if mod, ok := pkg.Modules[module]; ok {
			mods[module] = mod
		}
	}
	p.Modules = mods
	return p
}

func (pkg *Package) RequiredModules() Modules {
	var required []Module
	for name, module := range pkg.Modules {
		if module.Required {
			module.Name = name
			required = append(required, module)
		}
	}
	return required
}

func (pkg *Package) AvailableModules() Modules {
	var available []Module
	for name, module := range pkg.Modules {
		if !module.Required {
			module.Name = name
			available = append(available, module)
		}
	}
	return available
}

func (pkg *Package) PreSelectedModules() Modules {
	var preSelected []Module
	for name, module := range pkg.Modules {
		if !module.Required && module.Default {
			module.Name = name
			preSelected = append(preSelected, module)
		}
	}
	return preSelected
}

func (pkg *Package) HasModule(module string) bool {
	if _, ok := pkg.Modules[module]; ok {
		return true
	}
	return false
}
