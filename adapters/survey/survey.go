package survey

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/AlecAivazis/survey/v2/terminal"
	"os"
)

type Question = survey.Question

var Ask = func(qs []*Question, response interface{}, opts ...survey.AskOpt) error {
	if err := survey.Ask(qs, response, opts...); err != nil {
		return surveyError(err)
	}
	return nil
}

func surveyError(err error) error {
	if err == terminal.InterruptErr {
		os.Exit(0)
	}
	return err
}
