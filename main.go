package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/cookiejar"
)

var client http.Client

func init() {
	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Fatalf("Error %s", err.Error())
	}
	client = http.Client{
		Jar: jar,
	}
}
func main() {
	abcd := &http.Cookie{
		Name:   "NEW",
		Value:  "15",
		MaxAge: 125,
	}
	xyz := &http.Cookie{
		Name:   "CLICK",
		Value:  "true",
		MaxAge: 125,
	}
	req, err := http.NewRequest("GET", "http://golang.org", nil)
	if err != nil {
		log.Fatalf("Error %s", err.Error())
	}
	req.AddCookie(abcd)
	req.AddCookie(xyz)
	for _, c := range req.Cookies() {
		fmt.Println(c)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("error occured.error is %s", err.Error())
	}
	defer resp.Body.Close()
	fmt.Printf("StatusCode: %d\n", resp.StatusCode)
}
