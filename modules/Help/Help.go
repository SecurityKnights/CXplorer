package module

import "fmt"

func Help() {

	fmt.Println("\nCommon Flags")
	fmt.Println("-h\t\tPrint Help Menu")
	fmt.Println("-o\t\tGet Output to a file")
	fmt.Println("  \t\t[Eg: -o <output-file>]")

	fmt.Println("\nScanning Directories")
	fmt.Println("-u\t\tTarget URL.")
	fmt.Println("  \t\t[Add the http:// or https:// as prefix and / as suffix]")
	fmt.Println("  \t\t[You can enter the IP Address or Domain Name")
	fmt.Println("-w\t\tWordlist to Use")
	fmt.Println("  \t\t[Default Wordlist is loaded]")
	fmt.Println("-f\t\tFile Extensions to Specify separated by comma")
	fmt.Println("  \t\t[Default: html,php,txt]")

	fmt.Println("\nScanning SubDomains")
	fmt.Println("-s\t\tDomain to Scan")
	fmt.Println("  \t\t[Eg: http://www.google.com/ or https://www.google.com/]")
	fmt.Println("-w\t\tWordlist to Use")
	fmt.Println("  \t\t[Default Wordlist is loaded]")

	fmt.Println("\nScanning Ports")
	fmt.Println("-p\t\tTarget IP Address")
	fmt.Println("-protocol\tSpecify the Protocol")
	fmt.Println("  \t\t[Default: tcp]")
	fmt.Println("-ports\t\tSpecify the Start and End Port with a hyphen")
	fmt.Println("  \t\t[Default: 1-1000]")
	fmt.Println("  \t\t[Example: -ports 1-1000")

	fmt.Println("=============================================================")
	fmt.Println("For any bugs contact: bugs@s3curityKnights.com")
}
