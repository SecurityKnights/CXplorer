package modules

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

var BaseURL string
var wg sync.WaitGroup

func CheckURL(URL string) int {
	resp, err := http.Get(URL)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	return resp.StatusCode
}

func DirectoryScan(URL string, words []string, fileExtensions []string, output bool, outputFile string) {
	defer wg.Done()

	for i := 0; i < len(words); i++ {
		tempURL := URL + words[i]

		res := CheckURL(tempURL)

		if res == 200 || res == 204 || res == 301 || res == 302 || res == 307 || res == 403 {
			if output {
				data := "FOUND: " + tempURL + "\t[" + strconv.Itoa(res) + "]\n" + "Scanning Sub Directories in " +
					tempURL + "\n"

				f, err := os.OpenFile(outputFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)

				if err != nil {
					log.Fatal(err)
				} else {
					if _, err := f.WriteString(data); err != nil {
						log.Fatal(err)
					} else {
						f.Close()
					}
				}
			} else {
				fmt.Printf("FOUND: %s \t [%d]\n", tempURL, res)
				fmt.Printf("\nScanning Sub Directories in %s\n", tempURL)
			}

			if tempURL[len(tempURL)-1:] == "/" {
				go DirectoryScan(tempURL, words, fileExtensions, output, outputFile)
			} else {
				go DirectoryScan(tempURL+"/", words, fileExtensions, output, outputFile)
			}
		}
	}

	if BaseURL != URL {
		FileScan(URL, words, fileExtensions, output, outputFile)
	}
}

func FileScan(baseURL string, words []string, fileExtensions []string, output bool, outputFile string) {
	defer wg.Done()

	for i := 0; i < len(words); i++ {
		for j := 0; j < len(fileExtensions); j++ {
			tempURL := baseURL + words[i] + "." + fileExtensions[j]

			res := CheckURL(tempURL)

			if res == 200 || res == 204 || res == 301 || res == 302 || res == 307 || res == 403 {
				if output {
					data := "FOUND: " + tempURL + "\t[" + strconv.Itoa(res) + "]\n"

					f, err := os.OpenFile(outputFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)

					if err != nil {
						log.Fatal(err)
					} else {
						if _, err := f.WriteString(data); err != nil {
							log.Fatal(err)
						} else {
							f.Close()
						}
					}
				} else {
					fmt.Printf("FOUND: %s \t [%d]\n", tempURL, res)
				}
			}
		}
	}
}

func InitDirectoryScan(baseURL string, wordlist string, extensions string, output bool, outputFile string) {
	BaseURL = baseURL

	file, err := ioutil.ReadFile(wordlist)

	if err != nil {
		log.Fatal(err)
	}

	words := strings.Split(string(file), "\n")
	fileExtensions := strings.Split(extensions, ",")

	for i := range fileExtensions {
		fileExtensions[i] = strings.TrimSpace(fileExtensions[i])
	}

	fmt.Println("Target:\t\t\t", BaseURL)
	fmt.Println("File Extensions:\t", fileExtensions)
	fmt.Println("Wordlist: \t\t", wordlist)
	fmt.Println("HTTP Status Codes: \t\t[200,204,301,302,307,403]")
	fmt.Println("OS:\t\t\t", runtime.GOOS)
	fmt.Println("ARCHITECTURE:\t\t", runtime.GOARCH)
	fmt.Println("CPUs:\t\t\t", runtime.NumCPU())

	fmt.Println("=============================================================")

	fmt.Println("Scanning Started at", time.Now().Format("2006-01-02 3:4:5"))

	fmt.Println("=============================================================")

	if output {
		fmt.Println("Saving Data to File")

		data := "Target:\t\t\t" + BaseURL + "\n" + "File Extensions:\t" + strings.Join(fileExtensions, ",") + "\n" +
			"Wordlist: \t\t" + wordlist + "\n" + "HTTP Status Codes: \t\t[200,204,301,302,307,403]" + "\n" + "OS:\t\t\t" +
			runtime.GOOS + "\n" + "ARCHITECTURE:\t\t" + runtime.GOARCH + "\n" + "CPUs:\t\t\t" +
			strconv.Itoa(runtime.NumCPU()) + "\n" + "===========================================" +
			"==================" + "\n" + "Scanning Started at" + time.Now().Format("2006-01-02 3:4:5") + "\n" +
			"=============================================================\n"

		f, err := os.OpenFile(outputFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)

		if err != nil {
			log.Fatal(err)
		} else {
			if _, er := f.WriteString(data); er != nil {
				log.Fatal(er)
			}

			f.Close()
		}
	}

	wg.Add(2)
	timeStart := time.Now()

	go FileScan(BaseURL, words, fileExtensions, output, outputFile)

	go DirectoryScan(BaseURL, words, fileExtensions, output, outputFile)

	wg.Wait()

	timeElapsed := time.Since(timeStart)

	fmt.Println("=============================================================")

	fmt.Println("Scanning Done")
	fmt.Printf("Time Taken: %s\n", timeElapsed)

	if output {
		fmt.Println("Saving Data to File")

		data := "=============================================================" + "\n" + "Scanning Done" + "\n" +
			"Time Taken: " + strconv.Itoa(int(timeElapsed)) + "\n" +
			"=============================================================\n"

		f, err := os.OpenFile(outputFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)

		if err != nil {
			log.Fatal(err)
		} else {
			if _, er := f.WriteString(data); er != nil {
				log.Fatal(er)
			}

			f.Close()
		}
	}
}
