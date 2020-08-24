package sensitive

import (
	"bytes"
	"danmu/utils/files"
	"strings"
)

func init() {
	var allSensitiveWords []string

	if sensitiveWordPath, err := files.SensitiveWordPath(); err == nil {
		if fileContentByte, err := files.ReadFileBytes(sensitiveWordPath); err == nil {
			splitByte := bytes.Split(fileContentByte, []byte{'\n'})
			for _, sensitiveByte := range splitByte {
				allSensitiveWords = append(allSensitiveWords, string(sensitiveByte))
			}
		}
	}

	acAutoMachine = NewAcAutoMachine()
	for _, word := range allSensitiveWords {
		trimmedWord := strings.Trim(word, " ")

		if len(trimmedWord) > 0 {
			acAutoMachine.AddPattern(word)
		}
	}
	acAutoMachine.Build()
}
