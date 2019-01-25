package utils

import "fmt"

func CheckError(err error, desc interface{}) bool {
	if err != nil {
		fmt.Printf("%s,%s", err, desc)
		return false
	}

	return true
}
