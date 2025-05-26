package utils

import "fmt"

func Must(err error) {
	if err != nil {
		panic(err)
	}
}

func MustMsg(err error, msg string) {
	if err != nil {
		panic(fmt.Errorf("%s: %w", msg, err))
	}
}
