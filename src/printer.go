package src

import (
	"fmt"
	"strings"
)

var width int = 50

var rstring = strings.Repeat("=", width)

func PrintSummary(successes, errors, hardErrors, avarageDoneTime int) {
	println(rstring + "\n" + "Successful Requests: " + fmt.Sprint(successes) + "\n" + "Errors (not matching status code): " + fmt.Sprint(errors) + "\n" + "Hard errors (failed requests): " + fmt.Sprint(hardErrors) + "\n" + "Avarage thread done time: " + fmt.Sprint(avarageDoneTime) + "seconds" + "\n" + rstring)
}
