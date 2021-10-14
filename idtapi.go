package idtapi

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
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

func GetComplexityScore(sequences []string, username string, password string, clientId string, clientSecret string, urlPath string, urlToken string) [][]Problem {
	var sequencesInput []Sequence
	for i, sequence := range sequences {
		sequencesInput = append(sequencesInput, Sequence{"#" + strconv.Itoa(i), sequence})
	}

	auth := GetToken(username, password, clientId, clientSecret, urlToken)
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

func GetToken(username string, password string, clientId string, clientSecret string, urlToken string) Authentication {
	data := url.Values{}
	data.Set("grant_type", "password")
	data.Set("username", username)
	data.Set("password", password)
	data.Set("scope", "test")

	req, _ := http.NewRequest("POST", urlToken, strings.NewReader(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth(clientId, clientSecret)

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
	} else {
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}

		log.Fatalln(string(b))
	}

	return auth

}
