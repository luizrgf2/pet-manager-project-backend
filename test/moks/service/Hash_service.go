package test

import "fmt"

type HashServiceInMemory struct {
}

func (H HashServiceInMemory) Hash(data string) (*string, error) {
	toReturn := fmt.Sprintf("validpass")
	return &toReturn, nil

}

func (H HashServiceInMemory) Compare(data string, encrypedData string) (bool, error) {
	if data == "validpass" {
		return true, nil
	} else {
		return false, nil
	}
}
