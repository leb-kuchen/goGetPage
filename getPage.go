package goGetPage

import (
	"io"
	"log"
	"net/http"
)

func MustGetAndReadPage(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("bad request '%v'", resp.StatusCode)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("could not parse body")
	}
	return string(body)
}
