package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
)

func OnPage(link string) string {
	res, err := http.Get(link)
	if err != nil {
		log.Fatal(err)
	}
	content, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}

func countOccurrenceOnPage(url string, searchString string) (res result) {
	pageContent := OnPage(url)
	res.Path = url
	res.Count = strings.Count(pageContent, searchString)
	res.Err = nil
	return
}

func countOccurrenceInFile(filename string, searchString string) (res result) {
	res.Path = filename

	var bytesBody []byte
	bytesBody, res.Err = ioutil.ReadFile(filename)
	if res.Err != nil {
		panic(fmt.Sprintf("cannot read file with name %s", filename))
	}
	res.Count = strings.Count(string(bytesBody), searchString)
	return
}

type result struct {
	Path  string
	Count int
	Err   error
}

func Process(path string, searchString string, pathType string) result {
	switch pathType {

	// USAGE: echo -e 'https://golang.org\nhttps://golang.org\nhttps://golang.org' | go run quick.go -type url
	case "url":
		return countOccurrenceOnPage(path, searchString)

	// USAGE: echo -e '/etc/passwd\n/etc/hosts' | go run quick.go - type file
	case "file":
		return countOccurrenceInFile(path, searchString)

	default:
		return result{path, 0, errors.New("incorrect path type")}
	}
}

func main() {
	const searchString = "Go"

	typePtr := flag.String("type", "url", "url|file")
	flag.Parse()

	scanner := bufio.NewScanner(os.Stdin)

	var (
		wg       sync.WaitGroup
		resultCh = make(chan result)
		goroutines = make(chan struct{}, 5)
		done = make(chan struct{})
	)

	go func(resultCh <-chan result, done chan<- struct{}) {
		for res := range resultCh {
			fmt.Printf("Count for %s: %d\n", res.Path, res.Count)
		}
		done <- struct{}{}
	}(resultCh, done)

	for scanner.Scan() {
		path := scanner.Text()
		goroutines <- struct{}{}

		wg.Add(1)
		go func(path *string, resultCh chan<- result, goroutines <-chan struct{}, wg *sync.WaitGroup) {
			defer wg.Done()
			resultCh <- Process(*path, searchString, *typePtr)
			<- goroutines
		}(&path, resultCh, goroutines, &wg)
	}

	wg.Wait()
	close(resultCh)
	close(goroutines)

	<- done

}
