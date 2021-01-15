package main

import (
	"fmt"
	"strconv"

	"github.com/pkg/errors"

	"github.com/manifoldco/promptui"
)

// NOTE (About How to implement Yes/No prompt using promptui):
//   The method using "IsConfirm: true" is not suitable in this case.
//   Because any case where the input is not "y", non-nil error will be returned.
//   ref: https://github.com/manifoldco/promptui/issues/81
func receiveOUIInteractively() (oui, error) {
	prompt := promptui.Select{
		Label: "Do you want to specify OUI? [Yes/No]",
		Items: []string{"Yes", "No"},
	}
	_, result, err := prompt.Run()
	if err != nil {
		return "", errors.Wrap(err, "prompt failed: ")
	}
	if result == "No" {
		return "", nil
	}

	promptSecond := promptui.Prompt{
		Label: "Enter the OUI to specify",
		Validate: func(input string) error {
			if isValidOUI(input) {
				return nil
			}
			return errors.New("invalid input. it's must be valid OUI format")
		},
	}

	result, err = promptSecond.Run()
	if err != nil {
		return "", errors.Wrap(err, "prompt failed: ")
	}
	return oui(result), nil
}

func receiveFormatInteractively() (format, error) {
	type item struct {
		Name   string
		Value  format
		Sample string
	}
	items := []item{
		{Name: none.String(), Value: none, Sample: none.Sample()},
		{Name: colon.String(), Value: colon, Sample: colon.Sample()},
		{Name: hyphen.String(), Value: hyphen, Sample: hyphen.Sample()},
	}
	prompt := promptui.Select{
		Label: "Select a format for the address you want to generate",
		Items: items,
		Templates: &promptui.SelectTemplates{
			// Ref: https://github.com/manifoldco/promptui/blob/981a3cab68f6f3481bf42c6a98521af7fbd14fae/select.go#L421
			Label:    fmt.Sprintf("%s {{.Name}}: ", promptui.IconInitial),
			Active:   fmt.Sprintf("%s {{ .Name | underline }}", promptui.IconSelect),
			Inactive: "  {{.Name}}",
			Selected: fmt.Sprintf(`{{ "%s" | green }} {{ .Name | faint }}`, promptui.IconGood),
			Details: `
--------- Selected format ----------
{{ "Name:" | faint }}	{{ .Name }}
{{ "Sample:" | faint }}	{{ .Sample }}`,
		},
		Size: 4,
	}

	idx, _, err := prompt.Run()
	if err != nil {
		return 0, errors.Wrap(err, "prompt failed: ")
	}
	return items[idx].Value, nil
}

func receiveLettercaseInteractively() (lettercase, error) {
	type item struct {
		Name   string
		Value  lettercase
		Sample string
	}
	items := []item{
		{Name: upper.String(), Value: upper, Sample: upper.Sample()},
		{Name: lower.String(), Value: lower, Sample: lower.Sample()},
	}
	prompt := promptui.Select{
		Label: "Select a lettercase for the address you want to generate",
		Items: items,
		Templates: &promptui.SelectTemplates{
			Label:    fmt.Sprintf("%s {{.Name}}: ", promptui.IconInitial),
			Active:   fmt.Sprintf("%s {{ .Name | underline }}", promptui.IconSelect),
			Inactive: "  {{.Name}}",
			Selected: fmt.Sprintf(`{{ "%s" | green }} {{ .Name | faint }}`, promptui.IconGood),
			Details: `
--------- Selected lettercase ----------
{{ "Name:" | faint }}	{{ .Name }}
{{ "Sample:" | faint }}	{{ .Sample }}`,
		},
		Size: 3,
	}

	idx, _, err := prompt.Run()
	if err != nil {
		return 0, errors.Wrap(err, "prompt failed: ")
	}
	return items[idx].Value, nil
}

func receiveQuantityInteractively() (int, error) {
	prompt := promptui.Prompt{
		Label: "How many do you want to generate? (Recommend: 5)",
		Validate: func(input string) error {
			if isOnlyInteger(input) {
				return nil
			}
			return errors.New("invalid input. it's must be integer")
		},
	}

	result, err := prompt.Run()
	if err != nil {
		return 0, errors.Wrap(err, "prompt failed: ")
	}

	resultInt, err := strconv.Atoi(result)
	if err != nil {
		return 0, errors.Wrap(err, "type convert failed: ")
	}
	return resultInt, nil
}

func (c *config) receiveConfigsInteractively() error {
	o, err := receiveOUIInteractively()
	if err != nil {
		return err
	}
	p := []byte(o) // TODO: Improve
	c.p = p

	f, err := receiveFormatInteractively()
	if err != nil {
		return err
	}
	c.f = f

	l, err := receiveLettercaseInteractively()
	if err != nil {
		return err
	}
	c.l = l // Update config data

	q, err := receiveQuantityInteractively()
	if err != nil {
		return err
	}
	c.q = q

	return nil
}
