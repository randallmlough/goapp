package template

import (
	"github.com/randallmlough/gogen"
	"path"
)

type Component struct {
	Name         string `yaml:"name"`
	Type         string `yaml:"type"`
	OutputPath   string `yaml:"output_path"`
	TemplatePath string `yaml:"template_path"`
}

func (c *Component) File(baseTemplateDir string, imports ...string) (gogen.File, error) {
	var file gogen.File
	outputLoc := c.OutputPath
	templatePath := path.Join(baseTemplateDir, c.TemplatePath)
	if c.Type == "dir" {
		file = &gogen.Dir{
			OutputDir:   outputLoc,
			TemplateDir: templatePath,
		}
	} else {
		template, err := gogen.LoadTemplate(templatePath)
		if err != nil {
			return nil, err
		}

		if c.Type == "code" {
			file = &gogen.Go{
				Template: template,
				Filename: outputLoc,
				Imports:  imports,
			}

		} else {
			file = &gogen.Doc{
				Template: template,
				Filename: outputLoc,
			}
		}
	}
	return file, nil
}
