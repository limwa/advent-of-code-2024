package util

import (
	"strings"
)

func NormalizeInput(input *string) {
	*input = strings.ReplaceAll(*input, "\r\n", "\n")
	*input = strings.Trim(*input, "\n")
	
}
