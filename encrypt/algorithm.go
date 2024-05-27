package encrypt

func Nimbus(str string) string {
	encryptedStr := ""
	for _, c := range str {
		asciiCode := int(c)
		char := string(asciiCode + 3)
		encryptedStr += char
	}
	return encryptedStr
}
