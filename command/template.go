package command

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/Masterminds/sprig/v3"
	"github.com/pkg/errors"
	"github.com/randallmlough/goapp/adapters/colors"
	"github.com/randallmlough/goapp/command/template"
	"github.com/randallmlough/goapp/internal/utils"
	"github.com/randallmlough/gogen"
	"os"
	"os/exec"
	"path"
	"path/filepath"
)

var funcs = sprig.TxtFuncMap()

func getWorkingPath() (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return wd, nil
}

func askTemplate(templateDir string) (*template.Template, error) {

	templates, err := gatherAvailableTemplates(templateDir)
	if err != nil {
		return nil, err
	}

	var templateChoice string
	if err := survey.Ask([]*survey.Question{
		{
			Name: "templates",
			Prompt: &survey.Select{
				Message: "Choose a template",
				Options: templates,
			},
		},
	}, &templateChoice); err != nil {
		return nil, err
	}

	Println("loading template...", templateChoice)

	templatePath := path.Join(templateDir, templateChoice)
	t, err := template.New(templatePath)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func gatherAvailableTemplates(templateDir string) ([]string, error) {
	templates := []string{}
	err := filepath.Walk(templateDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if templateDir != path && info.IsDir() {
			templates = append(templates, info.Name())
			return filepath.SkipDir
		}

		return nil
	})
	if err != nil {
		return nil, errors.Wrap(err, "locating templates")
	}
	return templates, nil
}

func askModules(choices, selectedByDefault template.Modules) ([]string, error) {
	answers := []string{}
	if err := survey.Ask([]*survey.Question{
		{
			Name: "modules",
			Prompt: &survey.MultiSelect{
				Message: "What modules would you like to use?",
				Options: choices.Slice(),
				Default: selectedByDefault.Slice(),
			},
		},
	}, &answers); err != nil {
		return nil, err
	}

	return answers, nil
}

func (p *Project) Build() error {

	// build current dir
	wd, err := getWorkingPath()
	if err != nil {
		return err
	}
	projectDir := path.Join(wd, p.ProjectName)
	if err := utils.MakeDirAndChDir(projectDir); err != nil {
		return err
	}

	for pkgName, pkg := range p.Template.Packages {
		if err := utils.MakeDirAndChDir(pkgName); err != nil {
			return err
		}
		pkg.Deps = append(pkg.Deps, p.ImportPath)
		if pkg.GoModEnabled {
			cmd := exec.Command("go", "mod", "init", p.ImportPath)
			if err := cmd.Run(); err != nil {
				return fmt.Errorf("failed to initalize gomod %w", err)
			}
		}
		fmt.Printf("%s %s %s\n", colors.Yellow("[gogen]"), colors.Magenta("Generating Package"), pkgName)
		for modName, module := range pkg.Modules {
			fmt.Printf("%s %s %s\n", colors.Yellow("[gogen]"), colors.Green("Generating Module"), modName)
			for _, component := range module.Components {
				fmt.Printf("%s %s %s\n", colors.Yellow("[gogen]"), colors.Cyan("Generating File"), component.Name)
				file, err := component.File(path.Join(p.Template.LocalLocation, pkg.TemplateDir), pkg.Deps...)
				if err != nil {
					return err
				}
				if err := generateFile(file, p); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
func generateFile(file gogen.File, templateData interface{}) error {
	if err := gogen.Generate(file, gogen.SetGlobalTemplateData(templateData), gogen.SetGlobalFuncMap(funcs)); err != nil {
		fmt.Printf("%s %s %s\n", colors.Yellow("[gogen]"), colors.Black(" ERROR ").BgRed(), err)
		return err
	}
	return nil
}

func generateFiles(files []gogen.File, templateData interface{}) error {
	for _, file := range files {
		if err := generateFile(file, templateData); err != nil {
			return err
		}
	}
	return nil
}
