package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func checkAndSaveBody(url string, ch chan string) {
	resp, err := http.Get(url)
	if err != nil {
		s := fmt.Sprintf("%v is Down!, ", url)
		s += fmt.Sprintf(" Cause is %v\n", err)
		ch <- s
	} else {
		defer resp.Body.Close()
		s := fmt.Sprintf("%s - Status Code: %d\n", url, resp.StatusCode)
		if resp.StatusCode == http.StatusOK {
			bodyBytes, err := io.ReadAll(resp.Body)
			if err != nil {
				s += "Error reading body\n"
			}
			urlName := strings.Split(url, ".")[0]
			mainName := strings.Split(urlName, "//")[1]

			mainName += ".txt"

			s += fmt.Sprintf("writing response body to %s\n", mainName)

			writeErr := os.WriteFile(mainName, bodyBytes, 0644)

			if writeErr != nil {
				s += "Error writing file\n"
				ch <- s
			}
		}
		s += fmt.Sprintf("%v is up and running\n", url)
		ch <- s
	}

}

func main() {
	ch := make(chan string)
	defer close(ch)

	urls := []string{"https://ibedc.com", "https://golang.org", "https://google.com", "http://pay.ibedc.com"}

	for i, url := range urls {
		fmt.Println("Starting the execution")
		fmt.Println("No : ", i)
		go checkAndSaveBody(url, ch)
		fmt.Printf("Execution of %s ended \n", url)
	}

	for i := 0; i < len(urls); i++ {
		fmt.Println(<-ch)
	}
	fmt.Println("main execution finished")
}
