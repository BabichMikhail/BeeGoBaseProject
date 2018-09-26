package utils

import "errors"

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func ThrowError(message string) {
	panic(errors.New(message))
}
