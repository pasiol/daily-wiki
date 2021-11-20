package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Todo struct {
	task string
}

func main() {
	r, err := http.Get("https://en.wikipedia.org/wiki/Special:Random")
	if err != nil {
		log.Fatalf("getting wiki url failed: %s", err)
	}
	dailyWiki := Todo{task:	fmt.Sprintf("%s://%s%s", r.Request.URL.Scheme, r.Request.URL.Host, r.Request.URL.Path)}

	body, err := json.Marshal(dailyWiki)
	if err != nil {
		log.Fatalf("marshalling dailyWiki failed: %s", err)
	}
	http.Post(fmt.Sprintf("%s/todo",os.Getenv("TODO_HOST")),"application/json", bytes.NewBuffer(body))
	log.Printf("created task for wikiarticle: %s", dailyWiki)
}
