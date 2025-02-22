package utils

import "log"

// Check handles errors by logging and panicking if an error is encountered.
func Check(err error) {
	if err != nil {
		log.Panic(err)
	}
}

