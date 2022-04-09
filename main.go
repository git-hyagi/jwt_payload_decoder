/*
 Print the payload from a jwt token
*/
package main

import (
	b64 "encoding/base64"
	"fmt"
	"os"
	"regexp"
)

func main() {
	var encodedPayload, token string

	checkInput()
	token = os.Args[1]

	// try to parse jwt token
	r := regexp.MustCompile(`(\w+).(\w+).(.*)`)
	match := r.FindStringSubmatch(token)

	if len(match) != 4 {
		fmt.Println("Invalid jwt token!")
		os.Exit(2)
	}

	// padding the input length to follow base64 standard
	if len(match[2])%3 == 1 {
		encodedPayload = match[2] + "=="
	} else if len(match[2])%3 == 2 {
		encodedPayload = match[2] + "="
	} else {
		encodedPayload = match[2]
	}

	decodedPayload, _ := b64.StdEncoding.DecodeString(encodedPayload)
	fmt.Println(string(decodedPayload))

	// not working (tried prettyOutput like jq)
	//prettyOutput, _ := json.MarshalIndent(string(decodedPayload), "", "  ")
	//fmt.Println(string(prettyOutput))

}

func checkInput() {
	if len(os.Args) != 2 {
		fmt.Printf(`usage:
  %s <token>
`, os.Args[0])
		os.Exit(1)
	}

}
