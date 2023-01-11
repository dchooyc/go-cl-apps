package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
)

func main() {
	url := "https://edition.cnn.com/"
	html := retrieveHTML(url)
	f, err := os.Create("index.html")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err2 := f.WriteString(string(html))

	if err2 != nil {
		log.Fatal(err2)
	}

	re := regexp.MustCompile(`"headline":".*?"`)
	headlines := re.FindAll(html, -1)

	for _, headline := range headlines {
		n := len(headline)
		if n > 0 {
			fmt.Println(string(headline[12 : n-1]))
		}
	}
}

func retrieveHTML(url string) []byte {
	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	html, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	return html
}
