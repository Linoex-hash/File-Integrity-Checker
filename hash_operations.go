package main

import (
	"crypto/sha256"
	"fmt"
	"os"
)

// Hashes the file and returns the sha 256 digest as a string. Errors when the file specified by filename can't be opened.
func hash_file(filename string) (string, error) {
	file, err := os.ReadFile(filename)

	if err != nil {
		return "", err
	}

	hasher := sha256.New()
	hasher.Write(file)
	bs := hasher.Sum(nil)

	return fmt.Sprintf("%x", bs), nil

}
