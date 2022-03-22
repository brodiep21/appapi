package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type Info struct {
	Name  string `json:"name"`
	Bio   string `json:"bio"`
	Repos int    `json:"public_repos"`
}

func main() {
	apiContact("http://api.github.com/users/brodiep21")
}

func apiContact(url string) {
	var s Info

	client := &http.Client{Timeout: 5 * time.Second}
	req, err := client.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer req.Body.Close()

	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(body, &s)

	fmt.Println(s)
}
