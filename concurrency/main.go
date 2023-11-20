package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"runtime"
	"strings"
	"sync"
)

// func f1() {
// 	fmt.Println("f1 goroutine started")
// 	for i := 0; i < 5; i++ {
// 		fmt.Println("f1, i=", i)
// 	}
// 	fmt.Println("f1 goroutine ended")
// }

// func f2(wg *sync.WaitGroup) {
// 	fmt.Println("f2 goroutine started")
// 	for i := 5; i < 9; i++ {
// 		fmt.Println("f2 i=", i)
// 		//time.Sleep(time.Second) // indicating a time consuming execution
// 	}
// 	fmt.Println("f2 goroutine ended")
// 	wg.Done() // indicate that groutine is done
// }

func checkAndSaveBody(url string, wg *sync.WaitGroup) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("%s is down!, error: %s\n", url, err)
	} else {
		defer resp.Body.Close()
		if resp.StatusCode == http.StatusOK {
			bodyBytes, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Fatal("error reading webpage body:", err)
			}
			urlName := strings.Split(url, ".")[0]
			mainName := strings.Split(urlName, "//")[1]

			mainName += ".txt"

			fmt.Printf("writing response body to %s\n", mainName)

			writeErr := os.WriteFile(mainName, bodyBytes, 0644)

			if writeErr != nil {
				log.Fatal("error writing to body.html:", writeErr)
			}
		}
	}
	wg.Done()
}

func main() {
	// CONCURRENCY.
	// Spawning Goroutines

	// Printing out some system information

	// fmt.Println("No of Cores", runtime.NumCPU())
	// fmt.Println("No of Goroutines", runtime.NumGoroutine())
	// fmt.Println("OS", runtime.GOOS)
	// fmt.Println("Architecture", runtime.GOARCH)
	// fmt.Println("No of OS threads in parallel", runtime.GOMAXPROCS(0))

	// A gotoutine is a function called with go keyword before it name.

	// f1()
	// go f2()
	// fmt.Println("No of Goroutines", runtime.NumGoroutine())

	// time.Sleep(time.Millisecond) // But this method isn't logical in real world application.

	//
	//
	//

	// **WAIT GROUPS**//

	// STEP 1: instantiate sync.WaitGroup to an instance
	//wg := sync.WaitGroup{}
	// OR//
	//var waitG sync.WaitGroup

	//STEP 2: call the add of n method before attempting to execute the goroutine
	//wg.Add(1) // to specify the number of goroutines to wait for

	// STEP 3: change the function definition of the goroutine to indicate to the wait group

	// STEP 4: pass the wait group address to the goroutine
	// go f2(&wg)

	// f1()
	// STEP 5: call the Wait method on the wait group wg.
	//wg.Wait()S

	// websiteFile, err := os.OpenFile("HTML_File.html", os.O_WRONLY|os.O_CREATE, 0644)
	// if err != nil {
	// 	log.Fatal("Error oepning file: ", err)
	// }
	// defer websiteFile.Close()

	//** FOR THE EXERCISE **//

	urls := []string{"https://ibedc.com", "https://golang.org", "https://google.com", "http://pay.ibedc.com"}

	wg := sync.WaitGroup{}

	wg.Add(len(urls))

	for i, url := range urls {
		fmt.Println("Starting the execution")
		fmt.Println("No : ", i)
		go checkAndSaveBody(url, &wg)
		fmt.Printf("Execution of %s ended \n", url)
	}

	fmt.Println("No of Goroutines", runtime.NumGoroutine())

	wg.Wait()

	fmt.Println("main execution finished")
}
