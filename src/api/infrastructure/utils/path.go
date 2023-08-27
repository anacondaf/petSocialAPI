package utils

import (
	"errors"
	"os"
)

func PathIsNotExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil && errors.Is(err, os.ErrNotExist) {
		return true
	}

	return false
}
