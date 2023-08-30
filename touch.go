package main

import (
	"fmt"
	"os"
	"time"
)

func touch(filename string) error {
	// Check if file exists
	if _, err := os.Stat(filename); err != nil {
		if os.IsNotExist(err) {
			// If file does not exist, create a new empty file
			file, err := os.Create(filename)
			if err != nil {
				return err
			}
			defer file.Close()
		} else {
			// Some other error occurred while trying to stat the file
			return err
		}
	} else {
		// If file exists, update the modification time
		currentTime := time.Now()
		err := os.Chtimes(filename, currentTime, currentTime)
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: touch <filename>")
		return
	}

	err := touch(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
	}
}
