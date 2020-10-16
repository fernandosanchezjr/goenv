package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func Usage() {
	fmt.Println("Usage: goenv [destination_folder]")
}

func EnsurePathExists(path string) {
	if err := os.MkdirAll(path, 0755); err != nil {
		fmt.Printf("Error creating path %s: %s", path, err)
		os.Exit(-1)
	}
}

func CheckIfExists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		return !os.IsNotExist(err)
	} else {
		return true
	}
}

func CreateGoEnv(path string) {
	if CheckIfExists(path) {
		fmt.Println("Refreshing existing goenv in", path)
	} else {
		fmt.Println("Creating new goenv in", path)
	}
	EnsurePathExists(path)
	binPath := fmt.Sprintf("%s/bin", path)
	EnsurePathExists(binPath)
	WriteActivateScript(binPath)
	EnsurePathExists(fmt.Sprintf("%s/pkg", path))
	EnsurePathExists(fmt.Sprintf("%s/src", path))
}

func WriteActivateScript(binPath string) {
	if err := ioutil.WriteFile(fmt.Sprintf("%s/activate", binPath),
		[]byte(activateScript), 0755); err != nil {
		fmt.Printf("Error writing %s: %s", binPath, err)
		os.Exit(-1)
	}
}
