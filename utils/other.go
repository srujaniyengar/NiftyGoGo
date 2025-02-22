package utils

import "log"
// Check logs an error message and exits the program if an error is found.
// This prevents redundant error checking throughout the code.
func Check(err error) {
    if err != nil {
        log.Println("Error:", err)
    }
}

