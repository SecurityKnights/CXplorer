
# CXplorer

CXplorer is an reconnaissance tool written in Go which is capable of scanning directories, sub domains, open ports and sniff network traffic. For scanning it uses inbuilt net library and for network sniffing it uses gopacket. It is meant to find potential vulnerabilities and loopholes to help the user in improving web and network security.

## Installation

	
    git clone https://github.com/SecurityKnights/CXplorer
    cd CXplorer
    ./INSTALL.sh
    
Note: If you get any error related to 'Unable to Create Directories', do the same manually and then run the './INSTALL.sh' command again. This can happen due to lack of permissions. 

## Usage

    CXplorer <options> <target>
    
## Features

 - Directory and sub directory scanning
 - Port Scanning
 - Sub domain scanning
 - Network Sniffing

## Help Menu

    Usage: CXplorer <options> <target URL|IP>
    
    Commom Flags
    -h			Print Help Menu
    -o			Get Output to a file
    
    Scanning Directories
    -u			Target URL
				[Add the http:// or https:// as prefix and /  as suffix]
			        [You can enter the IP address or domain list]
	-w			Wordlist to use
				[Default wordlist is loaded]
	-f			File extensions to specify seperated by comma
				[Default: html,php,txt]
	
	Scanning SubDomains
	-s			Domain to scan
				[Eg: http://www.google.com/ or https://www.google.com]
	-w			Wordlist to use
				[Default wordlist is loaded]
	
	Scanning Ports
	-p			Target IP Address
	-protocol		Specify the protocol
				[Default: tcp]
	-ports			Specify the start and end port with hyphen(-)
				[Default:1-1000]
				[Eg: -ports 1-1000]

	Network Sniffer
	-net			Mandatory flag for network sniffing
				[Usage: -net <options>]
	
	Options
	list			List all devices
	cap			To start capturing
	write			To capture and save data in a pcap file
	read[file]		To read a pcap file
	filter			To capture data with given filter
	filterWrite		To capture data and save data in pcap file with filter
	packetInfo		To print data of packets at Ethernet,TCP,IPv4/v6 layers
	create			To create your own packet and send it
	decode			To print source and destination data of packets

## License
You can check the License [here](https://github.com/SecurityKnights/CXplorer/blob/main/LICENSE)

## Issues and Bugs
If you find any bugs or want to add some features to improve the code. Feel free to raise an issue regarding the same.

Happy Hacking :)
