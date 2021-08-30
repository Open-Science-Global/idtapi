package idtapi

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func ExampleGetComplexityScore() {
	err := godotenv.Load(".env")

	if err != nil {
		panic("Error loading .env file")
	}

	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	encodedAuth := os.Getenv("ENCODED_AUTH")
	urlPath := os.Getenv("COMPLEXITY_URL")
	urlToken := os.Getenv("TOKEN_URL")
	sequences := []string{"TGGTACGAAAATTAGGGGATCTACCTAGAAAGCCACAAGGCGAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAATAGGTCAAGCTTAAAGAACCCTTACATGGATCTTACAGATTCTGAAAGTAAAGAAACAACAGAGGTTAAACAAACAGAACCAAAAAGAAAAAAAGCATTGTTGAAAACAATGAAAGTTGATGTTTCAATCCATAATAAGATTAAATCGCTGCACGAAATTCTGGCAGCATCCGAAGGAAAAA"}
	fmt.Println(GetComplexityScore(sequences, username, password, encodedAuth, urlPath, urlToken))
	//Output: H
}

func ExampleGetToken() {
	err := godotenv.Load(".env")

	if err != nil {
		panic("Error loading .env file")
	}

	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	encondedAuth := os.Getenv("ENCODED_AUTH")
	urlToken := os.Getenv("TOKEN_URL")

	auth := getToken(username, password, encondedAuth, urlToken)
	fmt.Println(auth)
	//Output: Hello
}
