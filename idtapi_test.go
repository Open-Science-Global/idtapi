package idtapi

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func ExampleGetComplexityScore() {
	fmt.Println(GetComplexityScore([]string{"TGGTACGAAAATTAGGGGATCTACCTAGAAAGCCACAAGGCGAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAATAGGTCAAGCTTAAAGAACCCTTACATGGATCTTACAGATTCTGAAAGTAAAGAAACAACAGAGGTTAAACAAACAGAACCAAAAAGAAAAAAAGCATTGTTGAAAACAATGAAAGTTGATGTTTCAATCCATAATAAGATTAAATCGCTGCACGAAATTCTGGCAGCATCCGAAGGAAAAA"}))
	//Output: H
}

func ExampleGetToken() {
	err := godotenv.Load(".env")

	if err != nil {
		panic("Error loading .env file")
	}

	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	encondedAuth := os.Getenv("ENCONDED_AUTH")

	auth := getToken(username, password, encondedAuth)
	fmt.Println(auth)
	//Output: Hello
}
