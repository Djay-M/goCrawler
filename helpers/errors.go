package helpers

import "log"

func HandleErrors(funcName string, err error) {
	if err != nil {
		log.Fatal("Error in: ", funcName, "\n ERROR :::", err)
	}
}
