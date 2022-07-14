package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	var whatIsIt string

	secret := "aWFuZ25vVzpOQU06RU5JTDp0YTpzdTpuaW9K"
	sd, _ := base64.StdEncoding.DecodeString(secret)

	for index := len(sd) - 1; index >= 0; index-- {
		whatIsIt += string(sd[index])
	}

	fmt.Println(whatIsIt)
}
