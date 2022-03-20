package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)
type Info struct {
	Name string 'json: "name" '
	Bio string  'json: "bio"'
	Repos int  'json:"public_repos"'
}
func main() {

	// client := &http.Client
	req, err := http.Get("http://api.github.com/users/brodiep21")
	if err != nil {
		log.Fatal(err)
	}

	defer req.Body.Close()

	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Fatal(err)
	}
	s := string(body)
	fmt.Printf("%q\n", s)
}
