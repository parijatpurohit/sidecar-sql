package utils

import (
	"io/ioutil"
	"log"
)

func GetFileBytes(path string) ([]byte, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("could not read file, err: %v", err)
		return nil, err
	}
	return content, nil
}
