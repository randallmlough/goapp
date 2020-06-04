package command

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/AlecAivazis/survey/v2/terminal"
	"os"
	"strings"
)

const (
	fullStackApp = "FullStack App"
	apiOnly      = "API Only"
)

const (
	featureGraphQL            = "GraphQL"
	featureREST               = "REST"
	featureJWT                = "JWT"
	featureSecureCookies      = "Secure Cookies"
	featureAuthentication     = "Authentication"
	featureDatabaseMigrations = "Database Migrations"
	featureLogging            = "Logging"
)

const (
	frameworkReact = "React"
	frameworkHTML  = "HTML"
)

var (
	projectNameQuestion = survey.Question{
		Name:     "project_name",
		Prompt:   &survey.Input{Message: "Project name?"},
		Validate: survey.Required,
		Transform: survey.TransformString(func(s string) string {
			return strings.ReplaceAll(strings.ToLower(s), " ", "-")
		}),
	}
	appTypeQuestion = survey.Question{
		Name: "app_type",
		Prompt: &survey.Select{
			Message: "What kind of project is this?",
			Options: []string{fullStackApp, apiOnly},
			Default: fullStackApp,
		},
	}

	serverFeaturesQuestion = survey.Question{
		Name: "features",
		Prompt: &survey.MultiSelect{
			Message: "What features would you like to support?",
			Options: []string{featureGraphQL, featureREST},
			Default: []string{featureGraphQL},
		},
	}

	frameworkQuestion = survey.Question{
		Name: "framework",
		Prompt: &survey.Select{
			Message: "What framework would you like to use?",
			Options: []string{frameworkReact, frameworkHTML},
			Default: frameworkReact,
		},
	}

	initQuestion = survey.Question{
		Name: "init_git",
		Prompt: &survey.Confirm{
			Message: "Use Git?",
			Default: true,
		},
	}

	downloadDepsQuestion = survey.Question{
		Name: "download_deps",
		Prompt: &survey.Confirm{
			Message: "Download dependencies?",
			Default: true,
		},
	}
)

func surveyError(err error) error {
	if err == terminal.InterruptErr {
		os.Exit(0)
	}
	return err
}
