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

	if URL[:7] == "http://" {
		tempURL = URL[:7]
		k = 8
	} else if URL[:8] == "https://" {
		tempURL = URL[:8]
		k = 9
	}

	for i := 0; i < len(words); i++ {
		tempURL = tempURL + words[i] + "." + URL[k + 3:]

		resp, err := http.Get(tempURL)

		if err != nil {
			log.Fatal(err)
		} else {
			if resp.StatusCode == 200 || resp.StatusCode == 301 {
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