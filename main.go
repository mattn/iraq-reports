package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
)

type report struct {
	File      string `json:"file"`
	Page      string `json:"page"`
	Matchtext string `json:"matchtext"`
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s [word]\n", os.Args[0])
		os.Exit(1)
	}
	params := url.Values{}
	params.Set("q", os.Args[1])
	resp, err := http.Get(`https://iraq-reports.herokuapp.com/search/keyword?` + params.Encode())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var reports []report

	err = json.NewDecoder(resp.Body).Decode(&reports)
	if err != nil {
		log.Fatal(err)
	}
	for _, r := range reports {
		fmt.Printf("https://storage.googleapis.com/iraq-report-pdfs/%s (%s)\n  %s\n", r.File, r.Page, r.Matchtext)
	}
}
