# Copying Custom Modules
sudo mkdir /usr/local/go/src/CXplorer
sudo cp -r modules/ /usr/local/go/src/CXplorer

# Installing Required C Libraries
sudo apt-get install libpcap-dev

# Installing Required Go Packages
go get github.com/google/gopacket
go get golang.org/x/net/bpf
go get golang.org/x/sys/unix

# Copying Wordlist
sudo mkdir /usr/local/CXplorer
sudo cp wordlist.txt /usr/local/CXplorer/

# Building Tool
go build -o CXplorer

# Moving Executable to Bin
sudo cp CXplorer /usr/local/bin/
