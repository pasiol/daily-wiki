package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"log"
	"net/http"
	"os"
)

type payload struct {
	Task string `json:"task"`
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Print("Reading environment failed.")
	}
	r, err := http.Get("https://en.wikipedia.org/wiki/Special:Random")
	if err != nil {
		log.Fatalf("getting wiki url failed: %s", err)
	}

	task := fmt.Sprintf("%s://%s%s", r.Request.URL.Scheme, r.Request.URL.Host, r.Request.URL.Path)

	data := payload{
		Task: task,
	}
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("generating payload failed: %s", err)
	}

	body := bytes.NewReader(payloadBytes)
	url := fmt.Sprintf(
		"http://%s:%s/todos",
		os.Getenv("TODO_HOST"),
		os.Getenv("TODO_PORT"),
	)
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		// handle err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("request failed: %s", err)
	}
	if resp != nil {
		log.Printf("request statuscode %d %s", resp.StatusCode, resp.Status)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalf("closing request failed: %s", err)
		}
	}(resp.Body)

}
