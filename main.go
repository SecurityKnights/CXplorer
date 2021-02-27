package main

import (
	directoryScanner "CXplorer/modules/DirectoryScanner"
	helpMenu "CXplorer/modules/Help"
	networkSniffer "CXplorer/modules/NetworkSniffer"
	portScanner "CXplorer/modules/PortScanner"
	subdomainScanner "CXplorer/modules/SubDomainScanner"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// Filter Validation as filter will be used in single colons

func main() {
	var baseURL string
	var wordlist string
	var extensions string
	var domain string
	var ipAddress string
	var protocol string
	var ports string
	var outputFile string
	var help bool
	var net string

	flag.StringVar(&baseURL, "u", "", "Enter the target URL as http://[domain or ip]/ "+
		"\n (You can use https as well)")
	flag.StringVar(&wordlist, " w", "/usr/local/CXplorer/wordlist.txt", "Enter the wordlist path")
	flag.StringVar(&extensions, "f", "txt,html,php", "Enter the file extensions to be searched "+
		"separated via comma(,) ")
	flag.StringVar(&domain, "s", "", "Enter the domain you need to scan for subdomains")
	flag.StringVar(&ipAddress, "p", "", "Enter the target IP Address")
	flag.StringVar(&protocol, "protocol", "tcp", "Enter the protocol for which Ports need to scanned")
	flag.StringVar(&ports, "ports", "1-1000", "Enter the port range to be scanned in the format"+
		"\nstart-end")
	flag.StringVar(&outputFile, "o", "", "To Print Result in a given Output File")
	flag.BoolVar(&help, "help", false, "For Detailed Help")
	flag.StringVar(&net, "net", "", "For Network Sniffing")

	flag.Parse()

	fmt.Println("=============================================================")
	fmt.Println("CXplorer v0.1")
	fmt.Println("=============================================================")

	// Flag Validations
	if len(os.Args) < 2 {
		fmt.Println("ERROR: Use -h or -help flag for help.")
		fmt.Println("=============================================================")
		os.Exit(1)
	} else if baseURL != "" && domain != "" {
		fmt.Println("ERROR: You cannot scan directories and subdomains together. \n" +
			"This is restricted due to loss in performance.")
		fmt.Println("=============================================================")
		os.Exit(1)
	} else if baseURL != "" && domain != "" && ipAddress != "" {
		fmt.Println("ERROR: You cannot scan directories, subdomains and ports together.")
		fmt.Println("=============================================================")
		os.Exit(1)
	} else if baseURL != "" && net != "" {
		fmt.Println("ERROR: You cannot scan directories and capture packets together.")
		fmt.Println("=============================================================")
		os.Exit(1)
	} else if domain != "" && net != "" {
		fmt.Println("ERROR: You cannot scan subdomains and capture packets together.")
		fmt.Println("=============================================================")
		os.Exit(1)
	} else if baseURL != "" && domain != "" && net != "" {
		fmt.Println("ERROR: You cannot scan directories, subdomains and capture packets together.")
		fmt.Println("=============================================================")
		os.Exit(1)
	} else if baseURL != "" && domain != "" && ipAddress != "" && net != "" {
		fmt.Println("ERROR: You cannot use all the functions together.")
		fmt.Println("=============================================================")
		os.Exit(1)
	}

	// Input Validations
	if baseURL != "" {
		if ("http://" != baseURL[:7]) || (baseURL[len(baseURL)-1:] != "/") {
			fmt.Println("ERROR: Invalid URL.")
			fmt.Println("ERROR: Refer the help menu using -h or -help flag")
			fmt.Println("=============================================================")
			os.Exit(1)
		} else {
			resp, err := http.Get(baseURL)

			if err != nil {
				log.Fatal(err)
			} else if resp.StatusCode != 200 {
				fmt.Println("ERROR: Target Unreachable.")
				fmt.Println("=============================================================")
				os.Exit(1)
			}
		}
	} else if domain != "" {
		resp, err := http.Get(domain)

		if err != nil {
			log.Fatal(err)
		} else if resp.StatusCode != 200 {
			fmt.Println("ERROR: Target Unreachable.")
			fmt.Println("=============================================================")
			os.Exit(1)
		}
	} else if ipAddress != "" {
		octets := strings.Split(ipAddress, ".")

		if len(octets) != 4 {
			fmt.Println("ERROR: Invalid IP Address")
			fmt.Println("=============================================================")
			os.Exit(1)
		}

		for o := range octets {
			t, _ := strconv.Atoi(octets[o])

			if t > 255 || t < 0 {
				fmt.Println("ERROR: Invalid IP Address")
				fmt.Println("=============================================================")
				os.Exit(1)
			}
		}
	}

	output := false

	if outputFile != "" {
		output = true

		data := "=============================================================\n" +
			"CXplorer v0.1" + "=============================================================\n"

		f, err := os.OpenFile(outputFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)

		if err != nil {
			log.Fatal(err)
		}

		if _, err := f.WriteString(data); err != nil {
			log.Fatal(err)
		}

		f.Close()
	}

	if help {
		helpMenu.Help()
	} else if baseURL != "" {
		directoryScanner.InitDirectoryScan(baseURL, wordlist, extensions, output, outputFile)
	} else if domain != "" {
		subdomainScanner.InitSubdomainScanner(domain, wordlist)
	} else if ipAddress != "" {
		portScanner.InitPortScanner(ipAddress, protocol, ports)
	} else if net != "" {
		if net == "read" {
			data := strings.TrimSpace(os.Args[3])
			networkSniffer.InitNetworkSniffer(net, data)
		} else {
			networkSniffer.InitNetworkSniffer(net, "")
		}
	}

	fmt.Println("=============================================================")

	if output {
		data := "=============================================================\n"

		f, err := os.OpenFile(outputFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)

		if err != nil {
			log.Fatal(err)
		}

		defer f.Close()

		if _, err := f.WriteString(data); err != nil {
			log.Fatal(err)
		}
	}
}
