package template

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path"
)

type Template struct {
	LocalLocation string              `yaml:"-"`
	Name          string              `yaml:"name"`
	Description   string              `yaml:"description"`
	Version       string              `yaml:"version"`
	Source        string              `yaml:"source"`
	Packages      map[string]*Package `yaml:"packages"`
}

func New(template string) (*Template, error) {
	t, err := loadConfig(template)
	if err != nil {
		return nil, err
	}
	t.LocalLocation = template
	return t, nil
}

func loadConfig(template string) (*Template, error) {
	b, err := ioutil.ReadFile(path.Join(template, "template.yml"))
	if err != nil {
		return nil, err
	}

	var t Template
	if err := yaml.Unmarshal(b, &t); err != nil {
		return nil, err
	}
	return &t, nil
}

func (t *Template) Package(packageName string) *Package {
	if pkg, ok := t.Packages[packageName]; ok {
		return pkg
	}
	return nil
}

func (t *Template) RemovePackage(packageName string) {
	delete(t.Packages, packageName)
}
