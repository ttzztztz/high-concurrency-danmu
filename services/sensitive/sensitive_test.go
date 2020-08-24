package sensitive

import (
	"strings"
	"testing"
)

func TestAllowPublish(t *testing.T) {
	acAutoMachine = NewAcAutoMachine()

	allSensitiveWords := []string{
		"sb",
		"fuck",
	}
	for _, word := range allSensitiveWords {
		trimmedWord := strings.Trim(word, " ")

		if len(trimmedWord) > 0 {
			acAutoMachine.AddPattern(word)
		}
	}
	acAutoMachine.Build()

	if AllowPublish("you are an sb") != false {
		t.Fail()
	}

	if AllowPublish("fack") != true {
		t.Fail()
	}
}
