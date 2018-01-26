/*********************************************************************************************************************
* Cryptopals: 	Set 01 Challenge 01
* Title: 		Convert hex to base64
* Input: 		49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d
* Output: 		SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t
* Note: 		Always operate on raw bytes, never on encoded strings. Only use hex and base64 for pretty-printing.
*********************************************************************************************************************/
package main

import (
	b64 "encoding/base64"
	hex "encoding/hex"
	"fmt"
	"log"
)

func main() {

	var inputString string = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"

	hexEnc, err := hex.DecodeString(inputString)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("The String is:\t\t" + string(hexEnc))
	fmt.Println("In Hex is:\t\t" + hex.EncodeToString(hexEnc))
	base64Encode := b64.StdEncoding.EncodeToString(hexEnc)

	fmt.Println("Base64 Output is:\t" + base64Encode)
}
