package mathematic

import (
	"errors"
)

func FindElementString(element string, arrayFind []string) (string , error) {
	for _, v := range arrayFind {
		if v == element {
			return element , nil
		}
	}
	return "" , errors.New("Not found")
}