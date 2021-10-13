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
	urlPath := os.Getenv("COMPLEXITY_URL")
	urlToken := os.Getenv("TOKEN_URL")
	sequences := []string{"TGGTACGAAAATTAGGGGATCTACCTAGAAAGCCACAAGGCGAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAATAGGTCAAGCTTAAAGAACCCTTACATGGATCTTACAGATTCTGAAAGTAAAGAAACAACAGAGGTTAAACAAACAGAACCAAAAAGAAAAAAAGCATTGTTGAAAACAATGAAAGTTGATGTTTCAATCCATAATAAGATTAAATCGCTGCACGAAATTCTGGCAGCATCCGAAGGAAAAA"}
	fmt.Println(GetComplexityScore(sequences, username, password, "friendzymes001", "66d054be-f056-4a62-846a-e4b1f84baba7", urlPath, urlToken))

	//Output: H
}

func ExampleGetToken() {
	err := godotenv.Load(".env")

	if err != nil {
		panic("Error loading .env file")
	}

	username := "friendzymes"
	password := "Open3000"

	fmt.Println("USERNAME: ", username)
	fmt.Println("PASSWORD: ", password)

	token := GetToken(username, password, "friendzymes001", "66d054be-f056-4a62-846a-e4b1f84baba7", "https://www.idtdna.com/Identityserver/connect/token")
	fmt.Println(token)
	//Output: Hello
}
