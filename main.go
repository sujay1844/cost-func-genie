package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/huh"
)

var Theme = huh.ThemeCatppuccin()

func main() {
	metric := GetMetric()
	DisplayMetric(metric)
}

func GetMetric() string {
	Intro()

	switch QuestionWhatToPredict() {
	case "class-labels":

		switch QuestionAreBothClassesEquallyImportant() {
		case true:

			switch QuestionDoMostSamplesBelongToMajorityClass() {
			case true:
				return "Accuracy"
			case false:
				return "Geometric mean"
			}

		case false:
			switch QuestionWhichFalsesAreMoreCostly() {
			case "false-positives":
				return "F0.5 Score"
			case "false-negatives":
				return "F2 Score"
			case "both":
				return "F1 Score"
			}

		}

	case "probabilities":
		switch QuestionProbabilitiesOrClassLabels() {

		case "probabilities":
			return "Brier Score"

		case "class-labels":
			switch QuestionWhichClassesAreMoreImportant() {

			case "positive":
				return "PR AUC"
			case "both":
				return "ROC AUC"
			}
		}
	}

	return ""

}

func DisplayMetric(metric string) {
	huh.NewNote().
		Title("Recommended Metric: " + metric).
		Description("[Press any key to exit]").
		WithTheme(huh.ThemeCatppuccin()).
		Run()

	fmt.Println("Recommended Metric: " + metric)
}

func Intro() {
	fmt.Println()
	err := huh.NewNote().
		Title("Hi! I'm CostFunc Genie.").
		Description(`I'm here to help you choose the right performance metric
for your imbalanced binary classification problem.
[Press any key to continue]`).
		WithTheme(Theme).
		Run()

	HandleError(err)
}

func QuestionWhatToPredict() string {
	var value string
	err := huh.NewSelect[string]().
		Title("What do you want to predict?").
		Options(
			huh.NewOption("Class labels", "class-labels"),
			huh.NewOption("Probabilities", "probabilities"),
		).
		Value(&value).
		WithTheme(Theme).
		Run()

	HandleError(err)

	return value
}

func QuestionAreBothClassesEquallyImportant() bool {
	var value bool
	err := huh.NewSelect[bool]().
		Title("Are both classes equally important?").
		Options(
			huh.NewOption("Yes", true),
			huh.NewOption("No", false),
		).
		Value(&value).
		WithTheme(Theme).
		Run()

	HandleError(err)

	return value
}

func QuestionDoMostSamplesBelongToMajorityClass() bool {
	var value bool
	err := huh.NewSelect[bool]().
		Title("Do most samples belong to the majority class?").
		Options(
			huh.NewOption("Yes", true),
			huh.NewOption("No", false),
		).
		Value(&value).
		WithTheme(Theme).
		Run()

	HandleError(err)

	return value
}

func QuestionWhichFalsesAreMoreCostly() string {
	var value string
	err := huh.NewSelect[string]().
		Title("Which type of false is more costly?").
		Options(
			huh.NewOption("False positives", "false-positives"),
			huh.NewOption("False negatives", "false-negatives"),
			huh.NewOption("Both are equally costly", "both"),
		).
		Value(&value).
		WithTheme(Theme).
		Run()

	HandleError(err)

	return value
}

func QuestionProbabilitiesOrClassLabels() string {
	var value string
	err := huh.NewSelect[string]().
		Title("What do you have?").
		Options(
			huh.NewOption("Class labels", "class-labels"),
			huh.NewOption("Probabilities", "probabilities"),
		).
		Value(&value).
		WithTheme(Theme).
		Run()

	HandleError(err)

	return value
}

func QuestionWhichClassesAreMoreImportant() string {
	var value string
	err := huh.NewSelect[string]().
		Title("Which classes are more important?").
		Options(
			huh.NewOption("Positive class", "positive"),
			huh.NewOption("Both are equally important", "both"),
		).
		Value(&value).
		WithTheme(Theme).
		Run()

	HandleError(err)

	return value
}

func HandleError(err error) {
	if err == huh.ErrUserAborted {
		os.Exit(0)
	} else if err != nil {
		fmt.Println("An error occurred:", err)
		os.Exit(1)
	}
}
