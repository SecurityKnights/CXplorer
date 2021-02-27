package modules

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
	"strings"
	"sync"
	"time"
)

var wg sync.WaitGroup

func SubdomainScanner(URL string, words []string) {
	defer wg.Done()

	var tempURL = URL
	k := 1

	for i := 0; i < len(words); i++ {
		if URL[:7] == "http://" {
			tempURL = URL[:7]
			k = 8
		} else if URL[:8] == "https://" {
			tempURL = URL[:8]
			k = 9
		}

		tempURL = tempURL + words[i] + "." + URL[k+3:]

		resp, err := http.Get(tempURL)

		if err == nil {
			if resp.StatusCode == 200 || resp.StatusCode == 204 || resp.StatusCode == 301 || resp.StatusCode == 302 ||
				resp.StatusCode == 307 || resp.StatusCode == 403 {
				fmt.Printf("FOUND: %s \t [%d]\n", tempURL, resp.StatusCode)
			}
		}
	}
}

func InitSubdomainScanner(baseURL string, wordlist string) {
	file, err := ioutil.ReadFile(wordlist)

	if err != nil {
		log.Fatal(err)
	}

	words := strings.Split(string(file), "\n")

	fmt.Println("Target:\t\t\t", baseURL)
	fmt.Println("Wordlist: \t\t", wordlist)
	fmt.Println("HTTP Status Codes: \t [200,204,301,302,307,403]")
	fmt.Println("OS:\t\t\t", runtime.GOOS)
	fmt.Println("ARCHITECTURE:\t\t", runtime.GOARCH)
	fmt.Println("CPUs:\t\t\t", runtime.NumCPU())

	fmt.Println("=============================================================")

	fmt.Println("Scanning Started at", time.Now().Format("2006-01-02 3:4:5"))

	fmt.Println("=============================================================")

	wg.Add(1)

	timeStart := time.Now()

	go SubdomainScanner(baseURL, words)

	wg.Wait()

	timeElapsed := time.Since(timeStart)

	fmt.Println("=============================================================")

	fmt.Println("Scanning Done")
	fmt.Printf("Time Taken: %s\n", timeElapsed)
}
