package create

import (
	"fmt"
	"time"

	"github.com/everitosan/snimm-scrapper/internal/app/consult"
	"github.com/manifoldco/promptui"
)

func getOptionsPrompt(label string, options []string) (index int, result string, err error) {

	prompt := promptui.Select{
		Label: label,
		Items: options,
	}

	index, result, err = prompt.Run()

	return index, result, err
}

func validateDate(date string) error {
	if date == consult.Now {
		return nil
	}

	_, err := time.Parse("02/01/2006", date)
	if err != nil {
		return fmt.Errorf("inavlid date %v", err)
	}
	return nil
}

func getDatePrompt(label string) (result string, err error) {
	prompt := promptui.Prompt{
		Label:    label,
		Validate: validateDate,
	}

	result, err = prompt.Run()

	return result, err

}
