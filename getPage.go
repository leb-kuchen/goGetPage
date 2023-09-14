package goGetPage

import (
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
)

func MustGetAndReadPage(url string, sessionId string) string {
	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Fatal("error while creating cookie jar")
	}
	client := http.Client{
		Jar: jar,
	}
	cookie := &http.Cookie{
		Name:   "session",
		Value:  sessionId,
		MaxAge: 300,
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("bad request '%v'", err)
	}
	req.AddCookie(cookie)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error occured")
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("could not parse body")
	}
	return string(body)
}
