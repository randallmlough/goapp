package command

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/randallmlough/goapp/command/template"
	"github.com/randallmlough/goapp/internal/projectpath"
	"github.com/urfave/cli/v2"
	"path"
)

var NewProjectCmd = &cli.Command{
	Name:    "project",
	Aliases: []string{"p"},
	Usage:   "create a project",
	Flags: []cli.Flag{
		&cli.StringFlag{Name: "name, n", Usage: "name of the project"},
		&cli.StringFlag{Name: "template, t", Usage: "project template"},
	},
	Action: newProjectAction,
}

func newProjectAction(ctx *cli.Context) error {

	var projectName string
	if err := survey.Ask([]*survey.Question{&projectNameQuestion}, &projectName); err != nil {
		return surveyError(err)
	}
	project := NewProject(projectName)

	templateDir := path.Join(projectpath.Root, "templates")
	t, err := askTemplate(templateDir)
	if err != nil {
		return err
	}
	project.Template = t

	mods := make(map[string]*template.Package)
	for pkgName, pkg := range t.Packages {
		selections, err := askModules(pkg.AvailableModules(), pkg.PreSelectedModules())
		if err != nil {
			return err
		}
		mods[pkgName] = pkg.SelectModules(selections...)
	}
	t.Packages = mods

	if err := project.Build(); err != nil {
		return err
	}
	return nil
}

type Project struct {
	ImportPath  string
	ProjectName string
	Template    *template.Template
}

func (p *Project) Import(packageName string) string {
	return path.Join(p.ImportPath, packageName)
}

func NewProject(projectName string) *Project {
	return &Project{
		ImportPath:  "github.com/randallmlough/" + projectName,
		ProjectName: projectName,
	}
}
