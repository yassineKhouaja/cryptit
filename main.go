package main

import (
	"fmt"

	"github.com/yassinekhouaja/cryptit/decrypt"
	"github.com/yassinekhouaja/cryptit/encrypt"
)

func main() {
	enc := encrypt.Nimbus("Kodekloud")
	fmt.Println(enc)
	fmt.Println(decrypt.Nimbus(enc))
}
