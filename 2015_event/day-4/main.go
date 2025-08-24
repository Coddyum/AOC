package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

func main() {

	inputKey := "bgvyzdsv"
	aditionalValue := 1

	for {
		newKey := inputKey + strconv.Itoa(aditionalValue)
		trueValue := encodingMD5Hash2(newKey)

		if strings.HasPrefix(trueValue, "000000") {
			fmt.Println("Found at:", aditionalValue)
			fmt.Println("Hash:", trueValue)
			break
		} else {
			aditionalValue++
		}
	}

}

func encodingMD5Hash2(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
