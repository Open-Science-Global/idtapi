package idtapi

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type Authentication struct {
	Token   string `json:"access_token"`
	Expires int    `json:"expires_in"`
	Type    string `json:"token_type"`
}

type Sequence struct {
	Name     string `json:"Name"`
	Sequence string `json:"Sequence"`
}

type Problem struct {
	ComplexityScore float64 `json:"Score"`
	Name            string  `json:"Name"`
	StartIndex      int     `json:"StartIndex"`
}

func GetComplexityScore(sequences []string) [][]Problem {
	var sequencesInput []Sequence
	for i, sequence := range sequences {
		sequencesInput = append(sequencesInput, Sequence{"#" + strconv.Itoa(i), sequence})
	}
	err := godotenv.Load(".env")

	if err != nil {
		panic("Error loading .env file")
	}

	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	encondedAuth := os.Getenv("ENCONDED_AUTH")
	urlPath := os.Getenv("COMPLEXITY_URL")

	auth := getToken(username, password, encondedAuth)
	requestByte, _ := json.Marshal(sequencesInput)
	req, _ := http.NewRequest("POST", urlPath, bytes.NewReader(requestByte))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", auth.Type+" "+auth.Token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var gBlockAnalyzed [][]Problem
	if resp.Status == "200 OK" {
		err := json.NewDecoder(resp.Body).Decode(&gBlockAnalyzed)

		if err != nil {
			panic(err)
		}

		return gBlockAnalyzed
	}
	return gBlockAnalyzed

}
func getToken(username string, password string, encondedAuth string) Authentication {
	urlPath := os.Getenv("TOKEN_URL")

	data := url.Values{}
	data.Set("grant_type", "password")
	data.Set("username", username)
	data.Set("password", password)
	data.Set("scope", "test")

	req, err := http.NewRequest("POST", urlPath, strings.NewReader(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", encondedAuth)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	auth := Authentication{}
	if resp.Status == "200 OK" {
		err := json.NewDecoder(resp.Body).Decode(&auth)

		if err != nil {
			panic(err)
		}

		return auth
	}
	return auth
}
