package main

import (
	"os"
	"fmt"
	"net/http"
	"io/ioutil"
	"flag"
	"strings"
)

type Headers []string

func (h *Headers) String() string {
	return fmt.Sprintf("string")
}

func (h *Headers) Set(s string) error {
	*h = append(*h, s)
	return nil
}

var h Headers

func main() {
	flag.Var(&h, "header", "Header setting")

	flag.Parse()
	if len(flag.Args()) == 0 {
		fmt.Fprintf(os.Stderr, "%s: Specify the URL", os.Args[0])
		os.Exit(1)
	}

	url := flag.Args()[0]
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "NewRequest: %v", err)
		os.Exit(1)
	}
	for _, header := range h {
		spHeader := strings.Split(header, ": ")
	}

	client := new(http.Client)
	if resp, err := client.Do(req); err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	} else {
		defer resp.Body.Close()
		contents, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		}
		fmt.Printf("%s", string(contents))
	}
}
